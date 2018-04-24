// Package sdk ...
package sdk

import (
	"errors"
	"log"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

var (
	// ErrFabSDKInitFailed ...
	ErrFabSDKInitFailed = errors.New("Failed to initialize fabric channel client")
)

// NewFabSDK ...
func NewFabSDK(configFile string) (*fabsdk.FabricSDK, error) {
	cp := config.FromFile(configFile)
	sdk, err := fabsdk.New(cp)
	if err != nil {
		log.Println("Failed to read fabric SDK config file: ", err)
		return nil, ErrFabSDKInitFailed
	}
	return sdk, nil
}
