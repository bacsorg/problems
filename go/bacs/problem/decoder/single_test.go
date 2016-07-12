package decoder

import (
	"testing"

	"github.com/bacsorg/problem/go/bacs/problem"
	"github.com/bacsorg/problem_single/go/bacs/problem/single"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSingleTaskDecoder(t *testing.T) {
	task := single.Task{
		System: &problem.System{
			ProblemType: "problem-type",
			Package:     "package",
		},
	}
	expectedText := `system: <
  problem_type: "problem-type"
  package: "package"
>
`
	data, err := proto.Marshal(&task)
	require.NoError(t, err)
	text, err := SingleTaskDecoder.DecodeToText(data)
	if assert.NoError(t, err) {
		assert.Equal(t, expectedText, text)
	}
}
