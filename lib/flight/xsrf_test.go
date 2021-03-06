// Package flight_test
package flight_test

import (
	"log"
	"testing"

	"github.com/ibigfoot/blueprint/lib/flight"
	"github.com/ibigfoot/core/xsrf"
)

// TestRace tests for race conditions.
func TestXsrfRace(t *testing.T) {
	for i := 0; i < 100; i++ {
		go func() {
			// Set the csrf information
			flight.StoreXsrf(xsrf.Info{
				AuthKey: "test123",
				Secure:  true,
			})
			x := flight.Xsrf()
			x.AuthKey = "test"
			log.Println(flight.Xsrf().AuthKey)
		}()
	}
}
