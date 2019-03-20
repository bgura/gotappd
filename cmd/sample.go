package main

import (
	"fmt"
	"github.com/bgura/gotappd/pkg/gotappd"
)
func main() {
	clientid     := "<CLIENT_ID>"
	clientsecret := "<CLIENT_SECRET>"

	if clientid == "<CLIENT_ID>" {
		panic("Must set <CLIENT_ID>")
	}

	if clientsecret == "<CLIENT_SECRET>" {
		panic("Must set <CLIENT_SECRET>")
	}
	client := gotappd.NewClient(clientid, clientsecret)

	res, err := client.GetBeerInfo(16630, &gotappd.BeerInfoParams{Compact: true})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}