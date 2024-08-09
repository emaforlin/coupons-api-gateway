package main

import "errors"

var (
	ErrBindingBody     = errors.New("error binding body")
	ErrCreatingAccount = errors.New("could not create account")
)
