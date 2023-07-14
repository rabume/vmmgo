package vmmgo

import "errors"

var (
	ERR_BAD_VERSION = errors.New("Bad version")
	ERR_CALL        = errors.New("Function returns 0")
)
