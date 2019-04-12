// Package sdk ...
package sdk

import (
	"net/http"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"

	"github.com/govinda-attal/kiss-lib/pkg/core/status"
)

// ErrFabSDKInitFailed represents internal server error when unable to initialize fabric channel client.
var ErrFabSDKInitFailed = status.ErrServiceStatus{
	ServiceStatus: status.ServiceStatus{Code: http.StatusInternalServerError, Message: "Failed to initialize fabric channel client"},
}

// NewFabSDK returns an instance of FabricSDK for given configuration file.
func NewFabSDK(configFile string) (*fabsdk.FabricSDK, error) {
	cp := config.FromFile(configFile)
	sdk, err := fabsdk.New(cp)
	if err != nil {
		return nil, ErrFabSDKInitFailed
	}
	return sdk, nil
}
