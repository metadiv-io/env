# Env Package Documentation

## Problem Statement

The native Go library, os, retrieves environment variables with a default type of string. The env package addresses this limitation by providing a solution for retrieving environment variables in multiple data types. Additionally, it offers functionality for setting default values and checking the presence of required environment variables.

## Installation

To integrate the env package into your Go project, use the following go get command:

``` bash
go get -u github.com/metadiv-io/env
```

## Supported Types

The env package supports the following data types for environment variable retrieval:

* string
* int
* int64
* int32
* int16
* int8
* float64
* float32
* uint
* uint64
* uint32
* uint16
* uint8
* bool
* complex128
* complex64

## Usage

```go
package main

import (
	"fmt"
	"github.com/metadiv-io/env"
)

func main() {
	// Retrieve a string variable with a default value
	v := env.String("STRING_VAR", "default")

	// Retrieve a string variable and panic if not set
	v1 := env.MustString("STRING_VAR")

	// Retrieve a list of string variables with a default value
	vs := env.Strings("STRING_VARS", []string{"default"})

	// Retrieve a list of string variables and panic if not set
	vs1 := env.MustStrings("STRING_VARS")
}
```
