package decoder

import (
	"testing"

	"github.com/bacsorg/problem/go/bacs/problem"
	"github.com/bacsorg/problem_single/go/bacs/problem/single"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSingleDecoder(t *testing.T) {
	testData := []struct {
		pb           proto.Message
		expectedText string
		decoder      HumanDecoder
	}{
		{
			pb: &single.Task{
				System: &problem.System{
					ProblemType: "problem-type",
					Package:     "package",
				},
			},
			expectedText: `system: <
  problem_type: "problem-type"
  package: "package"
>
`,
			decoder: SingleTaskDecoder,
		},
		{
			pb: &single.Result{
				System: &problem.SystemResult{
					Status: problem.SystemResult_INVALID_REVISION,
				},
			},
			expectedText: `system: <
  status: INVALID_REVISION
>
`,
			decoder: SingleResultDecoder,
		},
	}
	for _, tt := range testData {
		data, err := proto.Marshal(tt.pb)
		require.NoError(t, err)
		text, err := tt.decoder.DecodeToText(data)
		if assert.NoError(t, err) {
			assert.Equal(t, tt.expectedText, text)
		}
	}
}
