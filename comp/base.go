package comp

import js "encoding/json"

type AmisComp interface {
	~map[string]any
}

// schema 通用 model，键未指定的 map
type schema = map[string]any

type (
	Schema schema
	Data   schema
)

type Response struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   Data   `json:"data"`
}

func (r *Response) Json() []byte {
	data, _ := js.Marshal(r)
	return data
}

// RawJson create components from amis raw json
func RawJson(data []byte) schema {
	s := schema{}
	js.Unmarshal(data, &s)
	return s
}
