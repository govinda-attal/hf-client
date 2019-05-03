package channel

import (
	"encoding/json"

	"github.com/golang/protobuf/proto"
)

func getArrayDataAsBytes(arr []interface{}) (*[][]byte, error) {
	var args [][]byte
	for _, data := range arr {
		switch v := data.(type) {
		case string:
			args = append(args, []byte(v))
		case proto.Message:
			b, err := proto.Marshal(data.(proto.Message))
			if err != nil {

			}
			args = append(args, b)
		default:
			b, err := json.Marshal(data)
			if err != nil {

			}
			args = append(args, b)
		}

	}
	return &args, nil
}

func getMapDataAsBytes(md map[string]interface{}) (*map[string][]byte, error) {
	var mb map[string][]byte
	for tk, td := range md {
		switch v := td.(type) {
		case string:
			mb[tk] = []byte(v)
		case proto.Message:
			b, err := proto.Marshal(td.(proto.Message))
			if err != nil {

			}
			mb[tk] = b
		default:
			b, err := json.Marshal(td)
			if err != nil {

			}
			mb[tk] = b
		}
	}
	return &mb, nil
}

func unmarshalRs(b []byte, rs interface{}) (interface{}, error) {
	if prs, ok := resp.(proto.Message); ok {
		if err := proto.Unmarshal(ccRs.Payload, prs); err != nil {
			return nil, err
		}
	} else {
		if err := json.Unmarshal(ccRs.Payload, rs); err != nil {
			return nil, err
		}
	}
	return rs, nil
}
