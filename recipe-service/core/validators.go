package core

import (
	"fmt"
	"strings"
	"reflect"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/model"
)

const tagName = "validate"

type Validator interface {
	// Validate method performs validation and returns result and optional error.
	Validate(interface{}) error
}

// DefaultValidator does not perform any validations.
type DefaultValidator struct {
}

func (v DefaultValidator) Validate(val interface{})  error {
	return  nil
}

type DifficultyValidator struct {
	Min int
	Max int
}

func (v DifficultyValidator) Validate(val interface{}) error {
	l,ok:=val.(model.Difficulty)
	if(!ok){
		return  fmt.Errorf(" is not a Difficulty struct")
	}
	if int(l) < v.Min {
		return  fmt.Errorf("difficulty should be at least %v ", v.Min)
	}

	if v.Max >= v.Min && int(l) > v.Max {
		return  fmt.Errorf("difficulty should be less than %v", v.Max)
	}

	return  nil
}

// gets tag from a struct
func getValidatorFromTag(tag string) Validator {
	args := strings.Split(tag, ",")

	switch args[0] {
	case "difficulty":
		validator := DifficultyValidator{}
		fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &validator.Min, &validator.Max)
		return validator
	default:
		return DefaultValidator{}
	}
}


// Performs actual data validation using validator definitions on the struct
func ValidateStruct(s interface{}) []string {
	errs := []string{}

	// ValueOf returns a Value representing the run-time data
	v := reflect.ValueOf(s)

	for i := 0; i < v.NumField(); i++ {
		// Get the field tag value
		tag := v.Type().Field(i).Tag.Get(tagName)

		// Skip if tag is not defined or ignored
		if tag == "" || tag == "-" {
			continue
		}

		// Get a validator that corresponds to a tag
		validator := getValidatorFromTag(tag)

		// Perform validation
		 err := validator.Validate(v.Field(i).Interface())

		// Append error to results
		if err != nil {
			errs = append(errs, fmt.Sprintf("%s %s", v.Type().Field(i).Name, err.Error()))
		}
	}

	return errs
}



