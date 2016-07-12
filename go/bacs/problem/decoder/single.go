package decoder

import (
	"github.com/bacsorg/problem_single/go/bacs/problem/single"
	"github.com/golang/protobuf/proto"
)

type singleTaskDecoder struct{}
type singleResultDecoder struct{}

func (_ singleTaskDecoder) DecodeToText(data []byte) (string, error) {
	var task single.Task
	err := proto.Unmarshal(data, &task)
	if err != nil {
		return "", err
	}
	return proto.MarshalTextString(&task), nil
}

func (_ singleResultDecoder) DecodeToText(data []byte) (string, error) {
	var result single.Result
	err := proto.Unmarshal(data, &result)
	if err != nil {
		return "", err
	}
	return proto.MarshalTextString(&result), nil
}

func (d singleTaskDecoder) DecodeBase64ToText(data string) (string, error) {
	return decodeBase64ToText(d, data)
}

func (d singleResultDecoder) DecodeBase64ToText(data string) (string, error) {
	return decodeBase64ToText(d, data)
}

var SingleTaskDecoder singleTaskDecoder
var SingleResultDecoder singleResultDecoder
