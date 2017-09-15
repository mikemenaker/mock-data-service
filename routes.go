package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"DataIndex",
		"GET",
		"/data",
		DataIndex,
	},
	Route{
		"DataShowStats",
		"GET",
		"/data/stats/{count}",
		DataShowStats,
	},
	Route{
		"DataShowText",
		"GET",
		"/data/text",
		DataShowText,
	},
	Route{
		"DataShowTable",
		"GET",
		"/data/table",
		DataShowTable,
	},
}
