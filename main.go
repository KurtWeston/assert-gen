package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yourusername/assert-gen/internal/generator"
)

var (
	outputFormat string
	testFuncName string
)

var rootCmd = &cobra.Command{
	Use:   "assert-gen",
	Short: "Generate Go test assertions from recorded runtime values",
	Long:  "assert-gen helps you create test assertions by recording function outputs during manual testing and generating copy-paste ready test code.",
}

var generateCmd = &cobra.Command{
	Use:   "generate [input.json]",
	Short: "Generate test code from recorded JSON",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		inputFile := args[0]
		gen := generator.New(outputFormat, testFuncName)
		code, err := gen.GenerateFromFile(inputFile)
		if err != nil {
			return fmt.Errorf("generation failed: %w", err)
		}
		fmt.Println(code)
		return nil
	},
}

func init() {
	generateCmd.Flags().StringVarP(&outputFormat, "format", "f", "standard", "Output format: standard or table")
	generateCmd.Flags().StringVarP(&testFuncName, "name", "n", "TestGenerated", "Test function name")
	rootCmd.AddCommand(generateCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
