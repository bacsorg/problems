package decoder

import (
	"github.com/bunsanorg/broker/go/bunsan/broker/rabbit"
	"github.com/golang/protobuf/proto"
)

type brokerTaskDecoder struct {
	payloadDecoder HumanDecoder
}
type brokerResultDecoder struct {
	payloadDecoder HumanDecoder
}

type brokerStatusDecoder struct{}

func NewBrokerTaskDecoder(payloadDecoder HumanDecoder) HumanDecoder {
	return &brokerTaskDecoder{payloadDecoder}
}
func NewBrokerResultDecoder(payloadDecoder HumanDecoder) HumanDecoder {
	return &brokerResultDecoder{payloadDecoder}
}

var BrokerStatusDecoder brokerStatusDecoder

func (dec brokerTaskDecoder) DecodeToText(data []byte) (string, error) {
	var task rabbit.RabbitTask
	err := proto.Unmarshal(data, &task)
	if err != nil {
		return "", err
	}
	payload := ""
	if task.Task != nil {
		payload, err = dec.payloadDecoder.DecodeToText(task.Task.Data)
		if err != nil {
			return "", err
		}
		task.Task.Data = nil
	}
	return proto.MarshalTextString(&task) + "\n" + payload, nil
}

func (dec brokerResultDecoder) DecodeToText(data []byte) (string, error) {
	var result rabbit.RabbitResult
	err := proto.Unmarshal(data, &result)
	if err != nil {
		return "", err
	}
	payload := ""
	if result.Result != nil {
		payload, err = dec.payloadDecoder.DecodeToText(result.Result.Data)
		if err != nil {
			return "", err
		}
		result.Result.Data = nil
	}
	return proto.MarshalTextString(&result) + "\n" + payload, nil
}

func (_ brokerStatusDecoder) DecodeToText(data []byte) (string, error) {
	var status rabbit.RabbitStatus
	err := proto.Unmarshal(data, &status)
	if err != nil {
		return "", err
	}
	return proto.MarshalTextString(&status), nil
}

func (d brokerTaskDecoder) DecodeBase64ToText(data string) (string, error) {
	return decodeBase64ToText(d, data)
}

func (d brokerResultDecoder) DecodeBase64ToText(data string) (string, error) {
	return decodeBase64ToText(d, data)
}

func (d brokerStatusDecoder) DecodeBase64ToText(data string) (string, error) {
	return decodeBase64ToText(d, data)
}
