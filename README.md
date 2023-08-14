# Go AsyncAPI Validator CLI

A simple cli-tool to validate JSON against a provided AsyncAPI spec using JSONPath. Currently it supports AsyncAPI V2.0.0

## Installation

1. Clone the repository:

```bash
git clone https://github.com/max-weis/go-asyncapi-validator.git
```

2. Navigate to the cloned repository:

```bash
cd go-asyncapi-validator
```

3. Build the CLI tool:

```bash
go build -o aa-validator
```

4. Move the binary to a directory in your **PATH** if desired:

```bash
sudo mv aa-validator /usr/local/bin/
```

After these steps, you should see an executable named `validator` on your system.

## Usage

To use the CLI tool, you need to provide three arguments:

- **spec**: Path to the AsyncAPI specification file.
- **json**: Path to the JSON file that you want to validate against the spec.
- **jsonpath**: The JSONPath expression which points to the location of the schema inside the spec.

Here's the general usage:

```bash
aa-validator -spec [PATH_TO_SPEC] -json [PATH_TO_JSON] -jsonpath [JSONPATH_EXPRESSION]
```

### Example:

To run the example project, just run:

```bash
aa-validator -spec ./examples/spec.json -json ./examples/example.json -jsonpath $.channels.personUpdates.subscribe.message.payload
```

If the provided JSON is valid against the given schema, the tool will output:

```
the provided JSON is valid
```

Otherwise, it will print validation errors.
