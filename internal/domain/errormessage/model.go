package errormessage

import (
	"encoding/json"
)

type Errors struct {
	ID         json.RawMessage `json:"id"`
	StatusCode int             `json:"statusCode"`
	Errors     *[]Error        `json:"errors"`
}

type Error struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Source Source `json:"source,omitempty"`
	Errors error  `json:"errors,omitempty"`
}

type Source struct {
	Pointer   string `json:"pointer,omitempty"`
	Parameter string `json:"parameter,omitempty"`
}
