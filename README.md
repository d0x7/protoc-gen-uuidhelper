# Protobuf UUID Helper Plugin ![GitHub License](https://img.shields.io/github/license/d0x7/protoc-gen-uuidhelper?color=blue) ![GitHub branch status](https://img.shields.io/github/check-runs/d0x7/protoc-gen-uuidhelper/main) ![GitHub release](https://img.shields.io/github/v/tag/d0x7/protoc-gen-uuidhelper?label=release) [![Go Reference](https://img.shields.io/badge/-reference-blue?style=flat&logo=go&logoColor=white&labelColor=gray&color=%23027D9C)](https://pkg.go.dev/xiam.li/uuidhelper)

This project provides a collection of `protoc` plugins that generate helper methods to access UUID fields in Protobuf messages across multiple languages (currently Kotlin and Go). These helpers simplify UUID parsing by providing idiomatic accessors for fields that represent UUIDs as byte arrays or strings.

## Features

- **Language Support**:
    - ✅ Go: Generates methods like `func (m *OnlinePlayer) GetSessionUUID() uuid.UUID`
    - ✅ Kotlin: Generates DSL extensions and functions like `player.getSessionUUID()`
  - 🧪 Extensible: Easily add support for more languages via a clean plugin structure

- **Field Detection**:
  - Automatically detects UUID fields based on the naming pattern `*_uuid` and protobuf type `bytes`
  - Also supports `repeated bytes *_uuids` fields

To see an example of what code is generated, see either the [Go](cmd/protoc-gen-uuidhelper-go/README.md) or [Kotlin](cmd/protoc-gen-uuidhelper-kotlin/README.md) plugin's README.

## Installation

Each language-specific plugin is located in its own folder under `cmd/` with a `README.md` file.

Follow the instructions in the respective folder to build and use the plugin:

- [Go Plugin](cmd/protoc-gen-uuidhelper-go/README.md)
- [Kotlin Plugin](cmd/protoc-gen-uuidhelper-kotlin/README.md)

All plugins can usually be built using the same command:

```bash
go build -o protoc-gen-uuidhelper-<lang> ./cmd/protoc-gen-uuidhelper-<lang>
```

It's recommended to add the plugin to your `$PATH` so you can invoke it directly.

To make this easier, there is a `Taskfile.yaml`, which requires [Task](https://taskfile.dev) to be installed. Then you can do

```bash
task install-<lang>
```

to install the plugin to your `$PATH`, or use the `install` task to install all available plugins.

## Adding a new language

To add support for a new language, follow these steps:

1. **Create a New Plugin Directory**

   Create a new folder under `cmd/` named `protoc-gen-uuidhelper-<lang>` (e.g., `protoc-gen-uuidhelper-python`).

2. **Implement the Plugin**

   Inside the new folder, create a `main.go` file that implements the following:
    - A backend struct the [`UUIDHelperBackend`](core/generator.go) interface.
      - Please check out existing implementations for [Go](cmd/protoc-gen-uuidhelper-go/plugin.go) and [Kotlin](cmd/protoc-gen-uuidhelper-kotlin/plugin.go) for reference.
      - The backend has to implement how the file is named and then return a new writer for this file.
    - A `main` function that then calls the `Main` method on the [`core`](core/generator.go) package with the backend struct.
    - A writer struct that implements the [`UUIDHelperWriter`](core/generator.go) interface.
      - Please check out existing implementations for [Go](cmd/protoc-gen-uuidhelper-go/writer.go) and [Kotlin](cmd/protoc-gen-uuidhelper-kotlin/writer.go) for reference.
      - The writer has to implement what should be generated for the file header (e.g., package, imports, etc.) as well as what methods to generate for each UUID field.

    For generating the methods, there are various helper methods in the `core` package that can be used e.g., convert a fields name from Protobuf' snake_case to the target language's camelCase or PascalCase.

    Refer to the Go implementation in `cmd/protoc-gen-uuidhelper-go/plugin.go` for a simple example (the kotlin plugin is a lot more complex) of how to structure the plugin.

> [!TIP]
> The writer will only be called for those fields that actually are considered "UUID-Fields" — therefore you don't have to filter the fields yourself, as that's already done by the core implementation.

3. **Build the Plugin**

   Build the plugin using the following command:
   ```bash
   go build -o protoc-gen-uuidhelper-<lang> ./cmd/protoc-gen-uuidhelper-<lang>
   ```

## License

MIT License. See [LICENSE](LICENSE) for details.
