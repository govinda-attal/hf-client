package ca

import (
	"log"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

// Client ...
type Client interface {
	Register(rq *msp.RegistrationRequest) (string, error)
	Enroll(usrName string, secret string) error
}

const (
	// IdentityTypeUser ...
	IdentityTypeUser = "User"
	// IdentityTypePeer ...
	IdentityTypePeer = "Peer"
	// IdentityTypeApp ...
	IdentityTypeApp = "App"
)

type client struct {
	caClient *msp.Client
}

// New ...
func New(sdk *fabsdk.FabricSDK, orgName string) (Client, error) {
	ctxProvider := sdk.Context()
	mspClient, err := msp.New(ctxProvider, msp.WithOrg(orgName))
	if err != nil {
		log.Println("Failed to create new ca client: ", err)
		return nil, ErrFabricCAClientInitFailed
	}
	return &client{caClient: mspClient}, nil
}

func (c *client) Register(rq *msp.RegistrationRequest) (string, error) {
	secret, err := c.caClient.Register(rq)
	if err != nil {
		log.Println("Failed to register user: ", err)
		return "", ErrFabricCAOperationFailed
	}
	return secret, nil
}

func (c *client) Enroll(identityName string, secret string) error {
	err := c.caClient.Enroll(identityName, msp.WithSecret(secret))
	if err != nil {
		log.Println("Failed to enroll user: ", err)
		return ErrFabricCAOperationFailed
	}
	return nil
}
