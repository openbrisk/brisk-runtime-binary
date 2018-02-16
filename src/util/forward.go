package util

import (
	"encoding/json"
)

// FunctionResult is the result structure for the forward feature.
type FunctionResult struct {
	Result  string        `json:"result"`
	Forward ForwardResult `json:"forward"`
}

// ForwardResult is the structure that contains the forward info.
type ForwardResult struct {
	To string `json:"to"`
}

// UnmarshalFunctionResult marshalls the given json to the FunctionResult struct.
func UnmarshalFunctionResult(data []byte) (*FunctionResult, error) {
	request := FunctionResult{}
	err := json.Unmarshal(data, &request)
	return &request, err
}
