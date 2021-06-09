package di

import (
	"go.uber.org/dig"
	"log"
	"ts/adapters"
	"ts/config"
	"ts/externalAPI/rest"
	"ts/externalAPI/tradeshiftAPI"
	"ts/mapping"
	"ts/offerImport/offerReader"
	"ts/ontology"
	"ts/ontologyValidator"
	"ts/reports"
	"ts/server"
	"ts/server/http/handlers"
)

type options = []dig.ProvideOption

type entry struct {
	constructor interface{}
	opts        options
}

var diConfig = []entry{
	{constructor: config.Get},
	{constructor: mapping.GetHandler},
	{constructor: handlers.New},
	{constructor: server.New},
	{constructor: adapters.NewFileManager},
	{constructor: adapters.NewHandler},
	{constructor: ontology.NewRulesHandler},
	{constructor: offerReader.NewOfferReader},
	{constructor: ontologyValidator.NewValidator},
	{constructor: reports.NewReportsHandler},
	{constructor: rest.NewRestClient},
	{constructor: tradeshiftAPI.NewTradeshiftAPI},
	{constructor: tradeshiftAPI.NewTradeshiftHandler},
}

func BuildContainer() *dig.Container {
	container := dig.New()
	for _, entry := range diConfig {
		if err := container.Provide(entry.constructor, entry.opts...); err != nil {
			log.Fatalf("DI provider error\n%s", err)
		}
	}
	return container
}
