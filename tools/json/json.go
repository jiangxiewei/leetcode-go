package json

import "encoding/json"

func ToString(obj any) string {
	bytes, _ := json.Marshal(obj)
	return string(bytes)
}
