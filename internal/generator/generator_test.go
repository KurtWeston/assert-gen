package generator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerateFromFile(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.json")
	jsonData := `{"function":"Add","test_cases":[{"name":"positive","inputs":{"a":1,"b":2},"output":3}]}`

	if err := os.WriteFile(testFile, []byte(jsonData), 0644); err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	t.Run("standard format", func(t *testing.T) {
		gen := New("standard", "TestAdd")
		code, err := gen.GenerateFromFile(testFile)
		if err != nil {
			t.Fatalf("GenerateFromFile() error = %v", err)
		}
		if !strings.Contains(code, "func TestAdd(t *testing.T)") {
			t.Errorf("Generated code missing test function")
		}
		if !strings.Contains(code, "t.Run") {
			t.Errorf("Generated code missing t.Run")
		}
	})

	t.Run("table format", func(t *testing.T) {
		gen := New("table", "TestAdd")
		code, err := gen.GenerateFromFile(testFile)
		if err != nil {
			t.Fatalf("GenerateFromFile() error = %v", err)
		}
		if !strings.Contains(code, "tests := []struct") {
			t.Errorf("Generated code missing table structure")
		}
		if !strings.Contains(code, "for _, tt := range tests") {
			t.Errorf("Generated code missing table loop")
		}
	})

	t.Run("file not found", func(t *testing.T) {
		gen := New("standard", "TestAdd")
		_, err := gen.GenerateFromFile("nonexistent.json")
		if err == nil {
			t.Error("GenerateFromFile() expected error for missing file")
		}
	})

	t.Run("invalid json", func(t *testing.T) {
		invalidFile := filepath.Join(tmpDir, "invalid.json")
		os.WriteFile(invalidFile, []byte(`{invalid}`), 0644)
		gen := New("standard", "TestAdd")
		_, err := gen.GenerateFromFile(invalidFile)
		if err == nil {
			t.Error("GenerateFromFile() expected error for invalid JSON")
		}
	})
}
