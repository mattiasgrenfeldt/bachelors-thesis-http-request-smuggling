package main

import (
	"fmt"
	"os"

	"github.com/mattiasgrenfeldt/bachelors-thesis-http-request-smuggling/test-harness/harness"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s <.req-files>\n", os.Args[0])
		os.Exit(1)
	}

	fmt.Println("Loading Requests...")
	reqs, err := harness.LoadRequests(os.Args[1:])
	if err != nil {
		panic(err)
	}

	fmt.Println("Building containers...")
	proxies, servers, err := harness.BuildAll()
	if err != nil {
		panic(err)
	}

	fmt.Println("Running servers...")
	serverResult, err := harness.RunServices("server", servers, reqs)
	if err != nil {
		panic(err)
	}

	fmt.Println("Running proxies...")
	proxyResult, err := harness.RunServices("proxy", proxies, reqs)
	if err != nil {
		panic(err)
	}

	fmt.Println("Reporting...")
	if err := harness.GenerateReport("full", serverResult, proxyResult, reqs); err != nil {
		panic(err)
	}
	if err := harness.GenerateReport("mini", serverResult, proxyResult, reqs); err != nil {
		panic(err)
	}
}
