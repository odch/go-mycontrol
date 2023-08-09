package mycontrol

import (
	"fmt"
	"net/http"
)

type Token struct {
	Token string `json:"token"`
}

func (c *client) GetToken() (string, error) {
	// clear token
	c.token = ""
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/token", c.baseURL), nil)
	if err != nil {
		return "", err
	}

	res := Token{}
	if err := c.sendRequest(req, &res); err != nil {
		return "", err
	}
	// update token
	c.token = res.Token
	return res.Token, nil
}
