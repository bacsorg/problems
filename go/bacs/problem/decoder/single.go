package decoder

import (
	"github.com/bacsorg/problem_single/go/bacs/problem/single"
	"github.com/golang/protobuf/proto"
)

type singleTaskDecoder struct{}
type singleResultDecoder struct{}

func (_ singleTaskDecoder) DecodeToText(data []byte) (string, error) {
	var task single.Task
	var profileExtension single.ProfileExtension
	err := proto.Unmarshal(data, &task)
	if err != nil {
		return "", err
	}
	if task.Profile != nil && task.Profile.Extension != nil &&
		task.Profile.Extension.TypeUrl ==
			"type.googleapis.com/bacs.problem.single.ProfileExtension" {
		ext := task.Profile.Extension
		task.Profile = nil
		err = proto.Unmarshal(ext.Value, &profileExtension)
		if err != nil {
			return "", err
		}
	}
	text := proto.MarshalTextString(&task)
	profileExtensionText := proto.MarshalTextString(&profileExtension)
	if profileExtensionText != "" {
		text += "\n"
		text += profileExtensionText
	}
	return text, nil
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
