package main

import (
	"plugin"
	log "github.com/kataras/golog"
)

func CrawlerRun() {
	log.Debug("qqqqq")
	crawler, err := plugin.Open("/data/go/proxypool/bin/kuaidaili.so")
	if err != nil {
		panic(err)
	}

	f, err := crawler.Lookup("Hello")
	if err != nil {
		panic(err)
	}

	f.(func())()
}