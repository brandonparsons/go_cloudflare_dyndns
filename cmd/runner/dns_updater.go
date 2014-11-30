package main

import (
    "flag"
    "log"
    "github.com/brandonparsons/go_cloudflare_dyndns"
)

func main() {
    cloudflareApiKey := flag.String("api-key", "", "Your Cloudflare API key/token")
    cloudflareAccountEmail := flag.String("api-email", "", "Your Cloudflare user email")
    forceUpdate := flag.Bool("force", false, "Force update, even if cached IP accurate")
    flag.Parse()

    if *cloudflareApiKey == "" {
        log.Fatal("A Cloudflare API key/token is required.")
    }

    if *cloudflareAccountEmail == "" {
        log.Fatal("A Cloudflare user email is required.")
    }

    log.Println("Running DNS update....")
    go_cloudflare_dyndns.Run(*cloudflareApiKey, *cloudflareAccountEmail, *forceUpdate)
    log.Println("Finished!")
}
