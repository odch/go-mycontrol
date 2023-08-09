package mycontrol

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const (
	BaseURLV1 = "https://mycontrol.aero/api/1.0"
)

type Client interface {
	GetToken() (string, error)
	GetFlight(id string) (*Flight, error)
	GetFlights(options *FlightsListOptions) (*FlightsList, error)
	AddFlight(flight *Flight) (*Flight, error)
}

type client struct {
	baseURL    string
	apiKey     string
	HTTPClient *http.Client
}

func NewClient(apiKey string) Client {
	return &client{
		baseURL: BaseURLV1,
		apiKey:  base64.StdEncoding.EncodeToString([]byte(apiKey + ":")),
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

type fieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

type errorResponse struct {
	Message string       `json:"message"`
	Code    string       `json:"code"`
	Errors  []fieldError `json:"errors"`
}

// Content-type and body should be already added to req
func (c *client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", c.apiKey))

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	// Try to unmarshall into errorResponse
	if (res.StatusCode != http.StatusOK) && (res.StatusCode != http.StatusCreated) {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Message)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	// Unmarshall and populate v
	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}

	return nil
}
