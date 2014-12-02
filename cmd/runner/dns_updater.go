package main

import (
    "flag"
    "strings"
    "log"
    "github.com/brandonparsons/go_cloudflare_dyndns"
)

func main() {
    cloudflareApiKey := flag.String("api-key", "", "Your Cloudflare API key/token")
    cloudflareAccountEmail := flag.String("api-email", "", "Your Cloudflare user email")

    forceUpdate := flag.Bool("force", false, "Force update, even if cached IP accurate")

    domainsFlag := flag.String("domains", "", "comma-separated list of domains to test")
    serviceLevelsFlag := flag.String("servicelevels", "", "comma-separated list of service levels to use for domains, in same order")

    flag.Parse()

    domainsToCheck := strings.Split(*domainsFlag, ",")
    serviceLevels := strings.Split(*serviceLevelsFlag, ",")

    if *cloudflareApiKey == "" {
        log.Fatal("A Cloudflare API key/token is required.")
    }

    if *cloudflareAccountEmail == "" {
        log.Fatal("A Cloudflare user email is required.")
    }

    if domainsToCheck[0] == "" {
        log.Fatal("A list of domains to update is required.")
    }

    if serviceLevels[0] == "" || len(serviceLevels) != len(domainsToCheck) {
        log.Fatal("A service level (0 or 1) is required for each domain.")
    }

    log.Println("Running DNS update....")
    go_cloudflare_dyndns.Run(*cloudflareApiKey, *cloudflareAccountEmail, domainsToCheck, serviceLevels, *forceUpdate)
    log.Println("Finished!")
}
