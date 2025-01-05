package models

import "net/http"

type Categories struct {
	Name        string `json:"categories_name"`
	Description string `json:"categories_description"`
}

func (c *Categories) Bind(r *http.Request) error {

	return nil
}
