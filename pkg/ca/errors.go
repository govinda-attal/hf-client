package ca

import "errors"

var (
	//ErrFabricCAClientInitFailed ...
	ErrFabricCAClientInitFailed = errors.New("Failed to initialize fabric ca client")
	//ErrFabricCAOperationFailed ...
	ErrFabricCAOperationFailed = errors.New("Failed to perform operation with fabric ca")
)
