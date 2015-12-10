package main

import (
	"fmt"

	"log"

	"github.com/davecgh/go-spew/spew"
	apiclient "github.com/jawspeak/go-jira-restclient/client"
	"github.com/jawspeak/go-jira-restclient/client/operations"
	"github.com/jawspeak/go-jira-restclient/config"
)

func main() {
	config.SetupConfig() // CRITICAL!
	// Note you can also set the transport in there to Debug for tracing.

	exampleGettingA400()

	exampleSuccessfulQuery()
}

func exampleSuccessfulQuery() {
	// Working example.
	resp1, err := apiclient.Default.Operations.GetSearch(operations.GetSearchParams{
		// You need to edit this query to play with the API.
		Jql:        `project = TEST and "Epic Link" = "My Epic" order by rank desc`,
		MaxResults: 1,
	})
	fatalIfErrUnless404(err)
	spew.Dump(resp1)
	fmt.Println("--------------")
	fmt.Printf("%#v\n\n", resp1.Payload)
}

func exampleGettingA400() {
	// More detailed example to handle a 400.
	_, err := apiclient.Default.Operations.GetSearch(operations.GetSearchParams{
		Jql:        "bogus jql, invalid query",
		MaxResults: 1,
	})
	if err == nil {
		panic("expecting error for 400 bad request")
	} else {
		if apiErr, ok := err.(operations.APIError); !ok {
			panic("... But err not an APIError")
		} else {
			if apiErr.Code != 400 {
				panic("expecting 400")
			}

			if resp400, ok := apiErr.Response.(*operations.GetSearchBadRequest); !ok {
				panic(".. err response not GetSearchBadRequest")
			} else {
				log.Println("this err's payload: ", resp400.Payload)
			}
		}
	}

}

func fatalIfErrUnless404(err error) {
	if err != nil {
		if apiErr, ok := err.(operations.APIError); ok {
			if apiErr.Code == 404 {
				log.Print("Not found - skipping", apiErr)
				return
			}
		}
		log.Fatal(err)
		spew.Dump(err)
	}
}
