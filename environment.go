package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// New creates a new environment variable
// It will return an error if the environment variable is not set and required is true
func New(key string, required bool) *Environment {
	e := &Environment{
		Key:      key,
		Value:    os.Getenv(key),
		Required: required,
	}
	if required && e.Value == "" {
		panic(fmt.Sprintf("environment variable %s is required", key))
	}
	return e
}

type Environment struct {
	Key      string `json:"key"`
	Value    string `json:"value"`
	Required bool   `json:"required"`
}

// String returns the environment variable value as a string
func (e *Environment) String() string {
	return e.Value
}

// Strings returns the environment variable value as a slice of strings
func (e *Environment) Strings() []string {
	if e.Value == "" {
		return []string{}
	}
	return strings.Split(e.Value, ",")
}

// Int returns the environment variable value as an integer
func (e *Environment) Int() int {
	i, err := strconv.Atoi(e.Value)
	if err != nil {
		return 0
	}
	return i
}

// Bool returns the environment variable value as a boolean
func (e *Environment) Bool() bool {
	b, err := strconv.ParseBool(e.Value)
	if err != nil {
		return false
	}
	return b
}

// Float returns the environment variable value as a float
func (e *Environment) Float() float64 {
	f, err := strconv.ParseFloat(e.Value, 64)
	if err != nil {
		return 0
	}
	return f
}

// Uint returns the environment variable value as an unsigned integer
func (e *Environment) Uint() uint {
	return uint(e.Int())
}
