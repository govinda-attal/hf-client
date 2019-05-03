package channel

import (
	"github.com/golang/protobuf/proto"
	"context"
	"encoding/json"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

// Client ...
type Client interface {
	//Execute runs a given function-transaction within chaincode on fabric channel. This transaction may result in blocks on the chain.
	Execute(ctx context.Context, ccID string, fcn string, args []interface{}, resp interface{}) (interface{}, error)
	//Query runs a given query function within the chaincode on the fabric channel. This transaction doesn't result in blocks on the chain.
	Query(ctx context.Context, ccID string, fcn string, args []interface{}, resp interface{}) (interface{}, error)
	//ExecuteWithTransArgs runs a given function-transaction within chaincode on fabric channel with transient arguments if necessary. This transaction may result in blocks on the chain.
	ExecuteWithTransArgs(ctx context.Context, ccID string, fcn string, args []interface{}, transDataMap map[string]interface{}, resp interface{}) (interface{}, error)
	//QueryWithTransArgs runs a given query function within the chaincode on the fabric channel with transient arguments if necessary. This transaction doesn't result in blocks on the chain.
	QueryWithTransArgs(ctx context.Context, ccID string, fcn string, args []interface{}, transDataMap map[string]interface{}, resp interface{}) (interface{}, error)
}

//New ...
func New(sdk *fabsdk.FabricSDK, channelID string, orgName string, userName string) (*FabChnClient, error) {

	chnlProvider := sdk.ChannelContext(channelID, fabsdk.WithUser(userName), fabsdk.WithOrg(orgName))

	chClient, err := channel.New(chnlProvider)
	if err != nil {
		return nil, ErrClientInitFailed().WithError(err)
	}
	return &FabChnClient{chnClient: chClient}, nil
}

//NewWithSignID ...
func NewWithSignID(sdk *fabsdk.FabricSDK, channelID string, orgName string, signID msp.SigningIdentity) (*FabChnClient, error) {
	clientChannelContext := sdk.ChannelContext(channelID, fabsdk.WithUser(orgName), fabsdk.WithIdentity(signID))
	chClient, err := channel.New(clientChannelContext)
	if err != nil {
		return nil, ErrClientInitFailed().WithError(err)
	}
	return &FabChnClient{chnClient: chClient}, nil
}

type FabChnClient struct {
	chnClient *channel.Client
}

func (fc *FabChnClient) Execute(ctx context.Context, ccID string, fcn string, req []interface{}, resp interface{}) (interface{}, error) {
	return fc.ExecuteWithTransArgs(ctx, ccID, fcn, req, nil, resp)
}

func (fc *FabChnClient) Query(ctx context.Context, ccID string, fcn string, req []interface{}, resp interface{}) (interface{}, error) {
	return fc.QueryWithTransArgs(ctx, ccID, fcn, req, nil, resp)
}

func (fc *FabChnClient) ExecuteWithTransArgs(ctx context.Context, ccID string, fcn string, req []interface{}, transDataMap map[string]interface{}, resp interface{}) (interface{}, error) {
	args, err := getArrayDataAsBytes(req)
	if err != nil {
		return nil, ErrRqMarshalFailed().WithError(err)
	}
	trans, err := getMapDataAsBytes(transDataMap)
	if err != nil {
		return nil, ErrRqMarshalFailed().WithError(err)
	}
	ccRs, err := fc.chnClient.Execute(channel.Request{
		ChaincodeID:  ccID,
		Fcn:          fcn,
		Args:         *args,
		TransientMap: *trans,
	})
	if err != nil {
		return nil, extractAppErr(err)
	}
	
	if _, ok := resp.(proto.Message); ok {
		err = proto.Unmarshal(ccRs.Payload, resp)	
	}
	else {
		err = json.Unmarshal(ccRs.Payload, resp)		
	}
	
	if err != nil {
		return nil, ErrRsUnMarshalFailed().WithError(err)
	}
	return resp, nil
}

func (fc *FabChnClient) QueryWithTransArgs(ctx context.Context, ccID string, fcn string, req []interface{}, transDataMap map[string]interface{}, resp interface{}) (interface{}, error) {
	args, err := getArrayDataAsBytes(req)
	if err != nil {
		return nil, ErrRqMarshalFailed().WithError(err)
	}
	trans, err := getMapDataAsBytes(transDataMap)
	if err != nil {
		return nil, ErrRqMarshalFailed().WithError(err)
	}
	ccRs, err := fc.chnClient.Query(channel.Request{
		ChaincodeID:  ccID,
		Fcn:          fcn,
		Args:         *args,
		TransientMap: *trans,
	}, channel.WithRetry(retry.DefaultChannelOpts))
	if err != nil {
		return nil, extractAppErr(err)
	}
	if _, ok := resp.(proto.Message); ok {
		err = proto.Unmarshal(ccRs.Payload, resp)	
	}
	else {
		err = json.Unmarshal(ccRs.Payload, resp)		
	}
	if err != nil {
		return nil, ErrRsUnMarshalFailed().WithError(err)
	}
	return resp, nil
}
