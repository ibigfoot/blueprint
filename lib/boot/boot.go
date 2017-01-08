// Package boot handles the initialization of the web components.
package boot

import (
	"log"

	"github.com/ibigfoot/blueprint/controller"
	"github.com/ibigfoot/blueprint/lib/env"
	"github.com/ibigfoot/blueprint/lib/flight"
	"github.com/ibigfoot/blueprint/viewfunc/link"
	"github.com/ibigfoot/blueprint/viewfunc/noescape"
	"github.com/ibigfoot/blueprint/viewfunc/prettytime"
	"github.com/ibigfoot/blueprint/viewmodify/authlevel"
	"github.com/ibigfoot/blueprint/viewmodify/flash"
	"github.com/ibigfoot/blueprint/viewmodify/uri"

	"github.com/blue-jay/core/form"
	"github.com/blue-jay/core/xsrf"
)

// RegisterServices sets up all the web components.
func RegisterServices(config *env.Info) {
	// Set up the session cookie store
	err := config.Session.SetupConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the MySQL database
	//mysqlDB, _ := config.MySQL.Connect(true)

	// connect to the PostgresDB
	pgsqlDB, _ := config.Postgresql.Connect(true)

	// Load the controller routes
	controller.LoadRoutes()

	// Set up the views
	config.View.SetTemplates(config.Template.Root, config.Template.Children)

	// Set up the functions for the views
	config.View.SetFuncMaps(
		config.Asset.Map(config.View.BaseURI),
		link.Map(config.View.BaseURI),
		noescape.Map(),
		prettytime.Map(),
		form.Map(),
	)

	// Set up the variables and modifiers for the views
	config.View.SetModifiers(
		authlevel.Modify,
		uri.Modify,
		xsrf.Token,
		flash.Modify,
	)

	// Store the variables in flight
	flight.StoreConfig(*config)

	// Store the database connection in flight
	flight.StoreDB(pgsqlDB)

	// Store the csrf information
	flight.StoreXsrf(xsrf.Info{
		AuthKey: config.Session.CSRFKey,
		Secure:  config.Session.Options.Secure,
	})
}
