package decoder

import (
	"github.com/bacsorg/problem_single/go/bacs/problem/single"
	"github.com/golang/protobuf/proto"
)

type singleTaskDecoder struct{}

func (_ singleTaskDecoder) DecodeToText(data []byte) (string, error) {
	var task single.Task
	err := proto.Unmarshal(data, &task)
	if err != nil {
		return "", err
	}
	return proto.MarshalTextString(&task), nil
}

func (d singleTaskDecoder) DecodeBase64ToText(data string) (string, error) {
	return decodeBase64ToText(d, data)
}

var SingleTaskDecoder singleTaskDecoder
