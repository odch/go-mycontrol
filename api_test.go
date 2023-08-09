package mycontrol

import (
	"fmt"
	"testing"
	"time"
)

const apiToken = "FILLHERE"

func TestApi(t *testing.T) {
	c := NewClient(apiToken)
	token, err := c.GetToken()
	if err != nil {
		panic(err)
	}
	fmt.Println(token)

	c = NewClient(token)
	flights, err := c.GetFlights(nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(flights)
	if len(flights.Flights) > 0 {

		one := flights.Flights[0]
		flight, err := c.GetFlight(one.Id)
		if err != nil {
			panic(err)
		}
		fmt.Println(flight)
	}

	depTime := time.Now().Add(-time.Hour).UTC()
	arrTime := time.Now().UTC()

	flightNew := Flight{}
	flightNew.PIC = "SELF"
	flightNew.Aircraft.Registration = "HBWYC"
	//"2021-01-01T12:00:00"
	flightNew.Departure.Time = Time(depTime.Format(time.RFC3339))
	flightNew.Departure.Place.Name = "LSZT"
	flightNew.Arrival.Place.Name = "LSZT"
	flightNew.Arrival.Time = Time(arrTime.Format(time.RFC3339))
	ldg := 1
	flightNew.Landings.Day = &ldg
	flight, err := c.AddFlight(&flightNew)
	if err != nil {
		panic(err)
	}
	fmt.Println(flight)

}
