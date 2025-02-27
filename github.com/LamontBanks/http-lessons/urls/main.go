package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	url, err := url.Parse("https://wiki.guildwars2.com/wiki/Main_Page")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(url.Hostname())
}
