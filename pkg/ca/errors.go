package ca

import (
	"github.com/govinda-attal/kiss-lib/pkg/core/status"
	"github.com/govinda-attal/kiss-lib/pkg/core/status/codes"
)

//ErrClientInitFailed ...
func ErrClientInitFailed() status.ErrServiceStatus {
	return status.ErrServiceStatus{
		ServiceStatus: status.ServiceStatus{Code: codes.ErrInternal, Message: "Failed to initialize fabric ca client"},
	}
}

//ErrOperationFailed ...
func ErrOperationFailed() status.ErrServiceStatus {
	return status.ErrServiceStatus{
		ServiceStatus: status.ServiceStatus{Code: codes.ErrInternal, Message: "Failed to perform operation with fabric ca"},
	}
}
