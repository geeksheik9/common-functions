package models

// ErrorResponse
// struct representation of what gets returned by the RespondWithError function.
// swagger:model
type ErrorResponse struct {
	Error string `json:"error"`
}
