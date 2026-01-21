package recorder

import (
	"testing"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		wantErr bool
	}{
		{
			name:    "valid recording",
			json:    `{"function":"Add","test_cases":[{"name":"test1","inputs":{"a":1},"output":2}]}`,
			wantErr: false,
		},
		{
			name:    "invalid json",
			json:    `{invalid}`,
			wantErr: true,
		},
		{
			name:    "missing function name",
			json:    `{"function":"","test_cases":[{"name":"test1","inputs":{},"output":1}]}`,
			wantErr: true,
		},
		{
			name:    "no test cases",
			json:    `{"function":"Add","test_cases":[]}`,
			wantErr: true,
		},
		{
			name:    "test case missing name",
			json:    `{"function":"Add","test_cases":[{"name":"","inputs":{},"output":1}]}`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Validate([]byte(tt.json))
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParse(t *testing.T) {
	t.Run("valid recording", func(t *testing.T) {
		json := `{"function":"Add","test_cases":[{"name":"add positive","inputs":{"a":1,"b":2},"output":3}]}`
		rec, err := Parse([]byte(json))
		if err != nil {
			t.Fatalf("Parse() error = %v", err)
		}
		if rec.Function != "Add" {
			t.Errorf("Function = %v, want Add", rec.Function)
		}
		if len(rec.TestCases) != 1 {
			t.Errorf("len(TestCases) = %v, want 1", len(rec.TestCases))
		}
	})

	t.Run("invalid json", func(t *testing.T) {
		_, err := Parse([]byte(`invalid`))
		if err == nil {
			t.Error("Parse() expected error, got nil")
		}
	})
}
