package formatter

import (
	"testing"
)

func TestFormatValue(t *testing.T) {
	tests := []struct {
		name  string
		input interface{}
		want  string
	}{
		{"nil", nil, "nil"},
		{"string", "hello", `"hello"`},
		{"int", 42, "42"},
		{"float", 3.14, "3.14"},
		{"bool true", true, "true"},
		{"bool false", false, "false"},
		{"empty slice", []int{}, "[]int{}"},
		{"int slice", []int{1, 2, 3}, "[]int{1, 2, 3}"},
		{"empty map", map[string]int{}, "map[string]int{}"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatValue(tt.input)
			if got != tt.want {
				t.Errorf("FormatValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatValue_Structs(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	t.Run("struct", func(t *testing.T) {
		p := Person{Name: "Alice", Age: 30}
		got := FormatValue(p)
		if got != `Person{Name: "Alice", Age: 30}` {
			t.Errorf("FormatValue() = %v", got)
		}
	})

	t.Run("pointer to struct", func(t *testing.T) {
		p := &Person{Name: "Bob", Age: 25}
		got := FormatValue(p)
		if got != `&Person{Name: "Bob", Age: 25}` {
			t.Errorf("FormatValue() = %v", got)
		}
	})

	t.Run("nil pointer", func(t *testing.T) {
		var p *Person
		got := FormatValue(p)
		if got != "nil" {
			t.Errorf("FormatValue() = %v, want nil", got)
		}
	})
}
