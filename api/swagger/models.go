// Package swagger provides models for swagger-docs
package swagger

import (
	"sea-battle/api/rest/v1"
)

// WithSuccess - successful request response model.
type WithSuccess struct {
	Success bool           `json:"success" example:"true"`
	Result  interface{}    `json:"result"`
	Errors  []v1.HTTPError `json:"errors"`
}

// WithError - failed request response model.
type WithError struct {
	Success bool           `json:"success" example:"false"`
	Result  interface{}    `json:"result"`
	Errors  []v1.HTTPError `json:"errors"`
}
