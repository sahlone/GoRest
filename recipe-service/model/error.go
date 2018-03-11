package model

import "fmt"

type Error struct {
	Code        string
	Description string
}

/*
 * Error collection for the service
 */
var ErrPatchIsVeg = Error{"IsVeg", "must be true | false"}
var ErrPatchDifficulty = Error{"Difficulty", "must be between 1 . 5 inclusive"}
var ErrorGetLimit = Error{"Limit", "must be between 1 . 100 inclusive"}
var ErrorGetSkip = Error{"Start", "must be greater than 0"}
var ErrNotFound = Error{"error.not.found", "Recipe with ID not found"}
var ErrorServiceUnavailable = Error{"error.service.unavailable", "Please try after sometime"}
var ErrTokenInvalid = Error{"authorization.token.invalid", "Auth token invalid"}
var ErrCredsInvalid = Error{"authorization.creds.invalid", "Invalid credentials"}
var ErrorInvalidRequest = Error{"error.request.invalid", "Invalid request payload/parameters"}

func (e Error) Error() string {
	return fmt.Sprintf("Code: %s Description: %v", e.Code, e.Description)
}
