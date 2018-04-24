package channel

import (
	"encoding/json"
	"log"

	"golang.org/x/net/context"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

// Client ...
type Client interface {
	//Execute ...
	Execute(ctx context.Context, ccID string, fcn string, args []interface{}, transDataMap map[string]interface{}, resp interface{}) error
	//Query ...
	Query(ctx context.Context, ccID string, fcn string, args []interface{}, transDataMap map[string]interface{}, resp interface{}) error
}

//New ...
func New(sdk *fabsdk.FabricSDK, channelID string, orgName string, userName string) (Client, error) {
	clientChannelContext := sdk.ChannelContext(channelID, fabsdk.WithUser(userName), fabsdk.WithOrg(orgName))
	chClient, err := channel.New(clientChannelContext)
	if err != nil {
		log.Println("Failed to create new channel client: ", err)
		return nil, ErrFabricChannelClientInitFailed
	}
	return &client{chnClient: chClient}, nil
}

//NewWithSignID ...
func NewWithSignID(sdk *fabsdk.FabricSDK, channelID string, orgName string, signID msp.SigningIdentity) (Client, error) {
	clientChannelContext := sdk.ChannelContext(channelID, fabsdk.WithUser(orgName), fabsdk.WithIdentity(signID))
	chClient, err := channel.New(clientChannelContext)
	if err != nil {
		log.Println("Failed to create new channel client: ", err)
		return nil, ErrFabricChannelClientInitFailed
	}
	return &client{chnClient: chClient}, nil
}

type client struct {
	chnClient *channel.Client
}

func (c *client) Execute(ctx context.Context, ccID string, fcn string, req []interface{}, transDataMap map[string]interface{}, resp interface{}) error {
	args, err := getArrayDataAsBytes(req)
	if err != nil {
		return ErrFabricChannelRqMarshalFailed
	}
	trans, err := getMapDataAsBytes(transDataMap)
	if err != nil {
		return ErrFabricChannelRqMarshalFailed
	}
	ccRs, err := c.chnClient.Execute(channel.Request{
		ChaincodeID:  ccID,
		Fcn:          fcn,
		Args:         *args,
		TransientMap: *trans,
	})
	if err != nil {
		return ErrFabricChannelOperationFailed
	}
	json.Unmarshal(ccRs.Payload, resp)
	return nil
}

func (c *client) Query(ctx context.Context, ccID string, fcn string, req []interface{}, transDataMap map[string]interface{}, resp interface{}) error {
	args, err := getArrayDataAsBytes(req)
	if err != nil {
		return ErrFabricChannelRqMarshalFailed
	}
	trans, err := getMapDataAsBytes(transDataMap)
	if err != nil {
		return ErrFabricChannelRqMarshalFailed
	}
	ccRs, err := c.chnClient.Query(channel.Request{
		ChaincodeID:  ccID,
		Fcn:          fcn,
		Args:         *args,
		TransientMap: *trans,
	})
	if err != nil {
		return ErrFabricChannelOperationFailed
	}
	json.Unmarshal(ccRs.Payload, resp)
	return nil
}
