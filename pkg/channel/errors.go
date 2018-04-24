package channel

import "errors"

var (
	//ErrFabricChannelClientInitFailed ...
	ErrFabricChannelClientInitFailed = errors.New("Failed to initialize fabric channel client")
	//ErrFabricChannelRqMarshalFailed ...
	ErrFabricChannelRqMarshalFailed = errors.New("Failed to marshal request message")
	//ErrFabricChannelRsUnMarshalFailed ...
	ErrFabricChannelRsUnMarshalFailed = errors.New("Failed to unmarshal response message")
	//ErrFabricChannelOperationFailed
	ErrFabricChannelOperationFailed = errors.New("Failed to invoke/query on fabric channel")
)
