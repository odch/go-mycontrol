package mycontrol

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type FlightsListOptions struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type FlightsList struct {
	Flights []FlightId `json:"flights"`
	Links   Links      `json:"links"`
}

type FlightId struct {
	Id string `json:"id,omitempty"`
}

type Place struct {
	Name    string `json:"name"`
	Outside bool   `json:"outside"`
}

type Time string

type ArrDep struct {
	Place Place `json:"place"`
	Time  Time  `json:"time"`
}
type Flight struct {
	FlightId
	Aircraft  Aircraft `json:"aircraft"`
	Arrival   ArrDep   `json:"arrival"`
	Departure ArrDep   `json:"departure"`
	PIC       string   `json:"pic"`
	Landings  Landings `json:"landings"`
}
type Aircraft struct {
	Registration string `json:"registration"`
	Type         string `json:"type,omitempty"`
}
type Landings struct {
	Day *int `json:"day"`
}

type Links struct {
	Next  LinkHref `json:"next"`
	Prev  LinkHref `json:"prev"`
	First LinkHref `json:"first"`
	Last  LinkHref `json:"last"`
}

type LinkHref struct {
	Href string `json:"href"`
}

func (c *Client) GetFlights(options *FlightsListOptions) (*FlightsList, error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/flights?limit=%d&page=%d", c.baseURL, limit, page), nil)
	if err != nil {
		return nil, err
	}

	res := FlightsList{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetFlight(id string) (*Flight, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/flights/%s", c.baseURL, id), nil)
	if err != nil {
		return nil, err
	}

	res := Flight{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) AddFlight(flight *Flight) (*Flight, error) {
	w := bytes.NewBuffer(nil)
	if err := json.NewEncoder(w).Encode(flight); err != nil {
		return nil, err
	}

	fmt.Println(string(w.Bytes()))
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/flights", c.baseURL), w)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	res := Flight{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
