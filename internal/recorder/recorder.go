package recorder

import (
	"encoding/json"
	"fmt"
)

type TestCase struct {
	Name   string                 `json:"name"`
	Inputs map[string]interface{} `json:"inputs"`
	Output interface{}            `json:"output"`
	Error  *string                `json:"error,omitempty"`
}

type Recording struct {
	Function  string     `json:"function"`
	TestCases []TestCase `json:"test_cases"`
}

func Validate(data []byte) error {
	var rec Recording
	if err := json.Unmarshal(data, &rec); err != nil {
		return fmt.Errorf("invalid JSON: %w", err)
	}
	if rec.Function == "" {
		return fmt.Errorf("function name is required")
	}
	if len(rec.TestCases) == 0 {
		return fmt.Errorf("at least one test case is required")
	}
	for i, tc := range rec.TestCases {
		if tc.Name == "" {
			return fmt.Errorf("test case %d: name is required", i)
		}
	}
	return nil
}

func Parse(data []byte) (*Recording, error) {
	if err := Validate(data); err != nil {
		return nil, err
	}
	var rec Recording
	if err := json.Unmarshal(data, &rec); err != nil {
		return nil, err
	}
	return &rec, nil
}
