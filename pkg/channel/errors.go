package channel

import (
	"encoding/json"
	"strings"

	"github.com/govinda-attal/kiss-lib/pkg/core/status"
	"github.com/govinda-attal/kiss-lib/pkg/core/status/codes"
)

//ErrClientInitFailed ...
func ErrClientInitFailed() status.ErrServiceStatus {
	return status.ErrServiceStatus{
		ServiceStatus: status.ServiceStatus{Code: codes.ErrInternal, Message: "Failed to initialize fabric channel client"},
	}
}

//ErrRqMarshalFailed ...
func ErrRqMarshalFailed() status.ErrServiceStatus {
	return status.ErrServiceStatus{
		ServiceStatus: status.ServiceStatus{Code: codes.ErrInternal, Message: "Failed to marshal request message"},
	}
}

//ErrRsUnMarshalFailed ...
func ErrRsUnMarshalFailed() status.ErrServiceStatus {
	return status.ErrServiceStatus{
		ServiceStatus: status.ServiceStatus{Code: codes.ErrInternal, Message: "Failed to unmarshal response message"},
	}
}

//ErrOperationFailed ...
func ErrOperationFailed() status.ErrServiceStatus {
	return status.ErrServiceStatus{
		ServiceStatus: status.ServiceStatus{Code: codes.ErrInternal, Message: "Failed to invoke/query on fabric channel"},
	}
}

func extractAppErr(err error) error {
	errStr := err.Error()
	i := strings.Index(errStr, "{")
	j := strings.LastIndex(errStr, "}")
	if i > -1 && j > -1 && i < j {
		b := []byte(errStr)[i : j+1]
		var errSvc status.ErrServiceStatus
		json.Unmarshal(b, &errSvc)
		if errSvc.Code != 0 {
			return errSvc
		}
	}
	return ErrOperationFailed().WithError(err)
}
