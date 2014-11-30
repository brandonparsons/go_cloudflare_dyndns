package go_cloudflare_dyndns

import (
    "os"
    "log"
    "strings"
    "net/http"
    "net/url"
    "io/ioutil"
    "github.com/bitly/go-simplejson"
)

func Run(apiKey string, apiEmail string, forceUpdate bool) {
    currentIP := getCurrentIP()
    cachedIP := getCachedIP()

    if currentIP == cachedIP {
        log.Println("IP unchanged - nothing to do.")
        if !forceUpdate {
            return
        }
    }

    log.Println("IP is outdated - updating cache and Cloudflare")
    writeToIPCache(currentIP)
    updateCloudFlare(apiKey, apiEmail, currentIP)
}

/////////////////////

func getCurrentIP() string {
    resp, err := http.Get("http://icanhazip.com")
    if err != nil {
        log.Fatal("IP source is down...")
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    check(err)

    return TrimSuffix(string(body[:]), "\n")
}

func getCachedIP() string {
    filename := iPCacheFilename()
    if fileExists(filename) {
        dat, err := ioutil.ReadFile(filename)
        check(err)
        return string(dat)
    } else {
        log.Println("Cache file did not exist - it will be created.")
        writeToIPCache("")
        return ""
    }
}

func writeToIPCache(content string) {
   ioutil.WriteFile(iPCacheFilename(), []byte(content), 0600)
}

func iPCacheFilename() string {
    return os.Getenv("HOME") + "/" + ".wan_ip-cf.txt"
}

func fileExists(filename string) bool {
    if _, err := os.Stat(filename); err == nil {
        return true
    }
    return false
}

func updateCloudFlare(apiKey string, apiEmail string, currentIP string) {
    baseURL := "https://www.cloudflare.com/api_json.html?"
    baseDomain := "brandonparsons.me"
    domain := "&z=" + url.QueryEscape(baseDomain)
    tkn := "&tkn=" + apiKey
    email := "&email=" + url.QueryEscape(apiEmail)

    cloudFlareAllRecordsURL := baseURL + "a=rec_load_all" + tkn + email + domain
    resp, err := http.Get(cloudFlareAllRecordsURL)
    if err != nil {
        log.Fatal("Cloudflare API is down...")
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    check(err)

    js, err := simplejson.NewJson(body)
    records := js.GetPath("response", "recs", "objs").MustArray()

    sitesToUpdate := []string{"pi", "vault"}
    for _, site := range sitesToUpdate {
        recordName := site + "." + baseDomain
        for _, record := range records {
            rec := record.(map[string]interface{})
            name := rec["name"]
            if name == recordName {
                recordID := rec["rec_id"]
                recordType := rec["type"]
                recordTTL := rec["ttl"]

                id := "&id=" + recordID.(string)
                typ := "&type=" + recordType.(string)
                ttl := "&ttl=" + recordTTL.(string)
                serviceMode := "&service_mode=1"
                content := "&content=" + url.QueryEscape(currentIP)
                nm := "&name=" + url.QueryEscape(recordName)

                cloudFlareEditRecordURL := baseURL + "a=rec_edit" + tkn + email + domain + id + typ + ttl + nm + serviceMode + content
                resp, err := http.Get(cloudFlareEditRecordURL)
                check(err)
                defer resp.Body.Close()
                _, err = ioutil.ReadAll(resp.Body)
            }
        }
    }
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func TrimSuffix(s, suffix string) string {
    if strings.HasSuffix(s, suffix) {
        s = s[:len(s)-len(suffix)]
    }
    return s
}
