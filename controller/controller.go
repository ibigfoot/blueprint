// Package controller loads the routes for each of the controllers.
package controller

import (
	"github.com/ibigfoot/blueprint/controller/about"
	"github.com/ibigfoot/blueprint/controller/debug"
	"github.com/ibigfoot/blueprint/controller/home"
	"github.com/ibigfoot/blueprint/controller/login"
	"github.com/ibigfoot/blueprint/controller/notepad"
	"github.com/ibigfoot/blueprint/controller/register"
	"github.com/ibigfoot/blueprint/controller/static"
	"github.com/ibigfoot/blueprint/controller/status"
)

// LoadRoutes loads the routes for each of the controllers.
func LoadRoutes() {
	about.Load()
	debug.Load()
	register.Load()
	login.Load()
	home.Load()
	static.Load()
	status.Load()
	notepad.Load()
}
