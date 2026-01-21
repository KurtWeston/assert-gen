# assert-gen

Generate Go test assertions from actual runtime values by recording function outputs during manual testing

## Features

- Record function calls with input parameters and return values to JSON
- Generate standard Go test assertions (if/else with t.Errorf)
- Generate table-driven test format with test cases array
- Support for primitive types (int, string, bool, float)
- Support for structs with nested fields and proper formatting
- Support for slices, maps, and pointer types
- Pretty-print generated code with proper indentation
- CLI flag to specify output format (standard vs table-driven)
- CLI flag to specify test function name
- Copy-paste ready test code output
- JSON schema validation for recorded sessions
- Handle nil values and zero values correctly

## How to Use

Use this project when you need to:

- Quickly solve problems related to assert-gen
- Integrate go functionality into your workflow
- Learn how go handles common patterns

## Installation

```bash
# Clone the repository
git clone https://github.com/KurtWeston/assert-gen.git
cd assert-gen

# Install dependencies
go build
```

## Usage

```bash
./main
```

## Built With

- go

## Dependencies

- `github.com/spf13/cobra`
- `github.com/stretchr/testify`

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
