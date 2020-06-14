package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"abv",
		"POST",
		"/abv",
		CalculateAbv,
	},
	Route{
		"hydrometer-adjustment",
		"POST",
		"/hydrometer-adjustment",
		CalculateHydrometerAdjustment,
	},
	Route{
		"srm",
		"POST",
		"/srm",
		CalculateSrm,
	},
	Route{
		"ibu",
		"POST",
		"/ibu",
		CalculateTotalIbu,
	},
	Route{
		"fermentables",
		"GET",
		"/fermentables",
		GetAllFermentables,
	},
	Route{
		"fermentable_name",
		"GET",
		"/fermentables/{name}",
		GetFermentableByName,
	},
	Route{
		"fermentable_id",
		"GET",
		"/fermentables/id/{id}",
		GetFermentableById,
	},
}