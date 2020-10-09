package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gosuri/uitable"
	"github.com/hashicorp/consul/api"
	"github.com/jawher/mow.cli"

	"github.com/pteich/consul-kv-search/search"
)

var Version string

func main() {
	app := cli.App("consul-kv-search", "CLI tool to search for data in Consul K/V store. https://github.com/pteich/consul-kv-search")
	app.Version("v version", Version)
	app.Spec = "[-a] [-t] [-p] [-r] [-w] [--keys|--values] QUERY"

	var (
		configConsulAddr       = app.StringOpt("a address", "127.0.0.1:8500", "Address and port of Consul server")
		configToken            = app.StringOpt("t token", "", "Consul ACL token to use")
		configPath             = app.StringOpt("p path", "/", "K/V path to start search")
		configWrap             = app.BoolOpt("w wrap", false, "Wrap text in output table")
		configQueryRegex       = app.BoolOpt("r regex", false, "Query interpreted as regular expression (instead of glob)")
		configQueryScopeKeys   = app.BoolOpt("keys", false, "Search keys only (default everyhwere)")
		configQueryScopeValues = app.BoolOpt("values", false, "Search values only (default everyhwere)")
		configQuery            = app.StringArg("QUERY", "*", "Search query")
	)

	app.Action = func() {
		consulConfig := api.DefaultConfig()
		consulConfig.Address = *configConsulAddr
		consulConfig.Token = *configToken

		consulClient, err := api.NewClient(consulConfig)
		if err != nil {
			log.Fatalf("could not connect to Consul at %s - %v", *configConsulAddr, err)
		}

		consulSearch := search.New(consulClient)
		var foundPairs []search.ResultPair

		scope := search.Everywhere
		if *configQueryScopeKeys {
			scope = search.Keys
		} else if *configQueryScopeValues {
			scope = search.Values
		}

		if *configQueryRegex {
			foundPairs, err = consulSearch.SearchRegex(*configQuery, *configPath, scope)
		} else {
			foundPairs, err = consulSearch.SearchGlob(*configQuery, *configPath, scope)
		}
		if err != nil {
			log.Fatal(err)
		}

		found := len(foundPairs)

		if found <= 0 {
			fmt.Println("0 entries found")
			return
		}

		fmt.Printf("%d entries found\n\n", found)

		table := uitable.New()
		table.MaxColWidth = 100
		table.Wrap = *configWrap
		table.Separator = " | "

		for _, element := range foundPairs {
			table.AddRow(element.Key, element.Value)
			table.AddRow("")
		}

		fmt.Println(table)
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
