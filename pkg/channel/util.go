package channel

import (
	"encoding/json"
)

func getArrayDataAsBytes(arr []interface{}) (*[][]byte, error) {
	var args [][]byte
	for _, data := range arr {
		switch v := data.(type) {
		case string:
			args = append(args, []byte(v))
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
		default:
			b, err := json.Marshal(td)
			if err != nil {

			}
			mb[tk] = b
		}
	}
	return &mb, nil
}
