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
		"hydrometer_adjustment",
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
		Name:        "dilution_boil_off",
		Method:      "POST",
		Pattern:     "/dilution-boil-off",
		HandlerFunc: CalculateDilutionBoilOff,
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
	Route{
		"hops",
		"GET",
		"/hops",
		GetAllHops,
	},
	Route{
		"hop_id",
		"GET",
		"/hops/{hopId}",
		GetHopByName,
	},
}