package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/jawher/mow.cli"
	"github.com/pteich/consul-kv-search/search"
	"log"
	"os"
	"text/tabwriter"
)

var Version string

func main() {

	app := cli.App("consul-kv-search", "CLI tool to search for data in Consul K/V store. https://github.com/pteich/consul-kv-search")
	app.Version("v version", Version)
	app.Spec = "[-a] [-t] [-p] [-g|-r] QUERY"

	var (
		configConsulAddr = app.StringOpt("a address", "127.0.0.1:8500", "Address and port of Consul server")
		configToken      = app.StringOpt("t token", "", "Consul ACL token to use")
		configPath       = app.StringOpt("p path", "/", "K/V path to start search")
		configQueryGlob  = app.BoolOpt("g glob", true, "Query interpreted as glob pattern")
		configQueryRegex = app.BoolOpt("r regex", false, "Query interpreted as regular expression")
		configQuery      = app.StringArg("QUERY", "*", "Search query")
	)

	app.Action = func() {
		fmt.Println(*configConsulAddr, *configToken, *configPath, *configQueryGlob, *configQueryRegex, *configQuery)

		consulConfig := api.DefaultConfig()
		consulConfig.Address = *configConsulAddr
		consulConfig.Token = *configToken

		consulClient, err := api.NewClient(consulConfig)
		if err != nil {
			fmt.Printf("Could not connect to Consul at %s - %v\n", *configConsulAddr, err)
			os.Exit(1)
		}

		consulSearch := search.NewConsulSearch(consulClient)
		var foundPairs []search.ResultPair

		if *configQueryRegex {
			foundPairs, err = consulSearch.SearchRegex(*configQuery, *configPath)
		} else {
			foundPairs, err = consulSearch.SearchGlob(*configQuery, *configPath)
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		w := new(tabwriter.Writer)
		w.Init(os.Stdout, 0, 8, 0, '\t', 0)
		for _, element := range foundPairs {
			fmt.Fprintf(w, "%s\t%s\n", element.Key, element.Value)
			fmt.Fprintln(w)
		}
		w.Flush()

	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
