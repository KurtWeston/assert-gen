package generator

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/yourusername/assert-gen/internal/formatter"
	"github.com/yourusername/assert-gen/internal/recorder"
)

type Generator struct {
	format   string
	funcName string
}

func New(format, funcName string) *Generator {
	return &Generator{format: format, funcName: funcName}
}

func (g *Generator) GenerateFromFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}
	rec, err := recorder.Parse(data)
	if err != nil {
		return "", err
	}
	if g.format == "table" {
		return g.generateTableDriven(rec), nil
	}
	return g.generateStandard(rec), nil
}

func (g *Generator) generateStandard(rec *recorder.Recording) string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("func %s(t *testing.T) {\n", g.funcName))
	for _, tc := range rec.TestCases {
		buf.WriteString(fmt.Sprintf("\tt.Run(%q, func(t *testing.T) {\n", tc.Name))
		args := []string{}
		for key, val := range tc.Inputs {
			args = append(args, fmt.Sprintf("%s := %s", key, formatter.FormatValue(val)))
		}
		for _, arg := range args {
			buf.WriteString(fmt.Sprintf("\t\t%s\n", arg))
		}
		buf.WriteString(fmt.Sprintf("\t\tgot := %s(%s)\n", rec.Function, strings.Join(getKeys(tc.Inputs), ", ")))
		buf.WriteString(fmt.Sprintf("\t\twant := %s\n", formatter.FormatValue(tc.Output)))
		buf.WriteString("\t\tif got != want {\n")
		buf.WriteString("\t\t\tt.Errorf(\"got %v, want %v\", got, want)\n")
		buf.WriteString("\t\t}\n")
		buf.WriteString("\t})\n")
	}
	buf.WriteString("}\n")
	return buf.String()
}

func (g *Generator) generateTableDriven(rec *recorder.Recording) string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("func %s(t *testing.T) {\n", g.funcName))
	buf.WriteString("\ttests := []struct {\n")
	buf.WriteString("\t\tname string\n")
	if len(rec.TestCases) > 0 {
		for key := range rec.TestCases[0].Inputs {
			buf.WriteString(fmt.Sprintf("\t\t%s interface{}\n", key))
		}
	}
	buf.WriteString("\t\twant interface{}\n")
	buf.WriteString("\t}{\n")
	for _, tc := range rec.TestCases {
		buf.WriteString("\t\t{\n")
		buf.WriteString(fmt.Sprintf("\t\t\tname: %q,\n", tc.Name))
		for key, val := range tc.Inputs {
			buf.WriteString(fmt.Sprintf("\t\t\t%s: %s,\n", key, formatter.FormatValue(val)))
		}
		buf.WriteString(fmt.Sprintf("\t\t\twant: %s,\n", formatter.FormatValue(tc.Output)))
		buf.WriteString("\t\t},\n")
	}
	buf.WriteString("\t}\n\n")
	buf.WriteString("\tfor _, tt := range tests {\n")
	buf.WriteString("\t\tt.Run(tt.name, func(t *testing.T) {\n")
	if len(rec.TestCases) > 0 {
		keys := getKeys(rec.TestCases[0].Inputs)
		args := []string{}
		for _, key := range keys {
			args = append(args, fmt.Sprintf("tt.%s", key))
		}
		buf.WriteString(fmt.Sprintf("\t\t\tgot := %s(%s)\n", rec.Function, strings.Join(args, ", ")))
	}
	buf.WriteString("\t\t\tif got != tt.want {\n")
	buf.WriteString("\t\t\t\tt.Errorf(\"got %v, want %v\", got, tt.want)\n")
	buf.WriteString("\t\t\t}\n")
	buf.WriteString("\t\t})\n")
	buf.WriteString("\t}\n")
	buf.WriteString("}\n")
	return buf.String()
}

func getKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
