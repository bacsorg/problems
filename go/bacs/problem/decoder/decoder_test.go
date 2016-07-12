package decoder

import (
	"testing"

	"github.com/bacsorg/problems/go/bacs/problem/decoder/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestDecodeBase64ToText(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	decoder := mock_decoder.NewMockHumanDecoder(ctrl)
	decoder.EXPECT().DecodeToText([]byte("hello world")).Return("result", nil)
	result, err := decodeBase64ToText(decoder, "aGVsbG8gd29ybGQ=")
	assert.NoError(t, err)
	assert.Equal(t, "result", result)
}
