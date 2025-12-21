package arrays_hashing_01

import (
	"reflect"
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	solution := &SolutionEncodedDecodedString{}

	testCases := []struct {
		name string
		data []string
	}{
		{
			name: "Example 1",
			data: []string{"neet", "code", "love", "you"},
		},
		{
			name: "Example 2",
			data: []string{"we", "say", ":", "yes"},
		},
		{
			name: "Single empty string",
			data: []string{""},
		},
		{
			name: "Multiple empty strings",
			data: []string{"", ""},
		},
		{
			name: "Some string + emoji",
			data: []string{"some string + 🍟", "🚀 string"},
		},
		{
			name: "Empty slice",
			data: []string{},
		},
		{
			name: "String with delimiter",
			data: []string{"hello#world", "123#456"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			encoded := solution.Encode(tc.data)
			decoded := solution.Decode(encoded)

			if !reflect.DeepEqual(tc.data, decoded) {
				t.Errorf("Test %s failed: \nInput: %v\nEncoded: %s\nDecoded: %v", tc.name, tc.data, encoded, decoded)
			}
		})
	}
}
