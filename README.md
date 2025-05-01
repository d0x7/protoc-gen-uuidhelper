# Protobuf UUID Helper Code Generator

This project provides a collection of `protoc` plugins that generate helper methods to access UUID fields in Protobuf messages across multiple languages (currently Kotlin and Go). These helpers simplify UUID parsing by providing idiomatic accessors for fields that represent UUIDs as byte arrays or strings.

## Features

- **Language Support**:
    - ✅ Go: Generates methods like `func (m *OnlinePlayer) GetPlayerUUID() uuid.UUID`
    - ✅ Kotlin: Generates extension functions like `player.getPlayerUUID()`
  - 🧪 Extensible: Easily add support for more languages via a clean plugin structure

- **Field Detection**:
  - Automatically detects UUID fields based on the naming pattern `*_uuid` and protobuf type `bytes`

## Example

Given this Protobuf message:

```proto
message Player {
  bytes internal_uuid = 1;
  string username     = 2;
  bytes session_uuid  = 3;
  string  string_uuid = 4;
}
```

The Kotlin generator outputs:

```kotlin
var PlayerKt.Dsl.InternalUUID: UUID
    get() = byteStringToUUID(this.internalUuid)
    set(value) {
        this.internalUuid = uuidToByteString(value)
    }

fun Test.Player.InternalUUID(): UUID = byteStringToUUID(this.internalUuid)

var PlayerKt.Dsl.SessionUUID: UUID
    get() = byteStringToUUID(this.sessionUuid)
    set(value) {
        this.sessionUuid = uuidToByteString(value)
    }

fun Test.Player.SessionUUID(): UUID = byteStringToUUID(this.sessionUuid)
```

The Go generator outputs:

```go
func (m *Player) GetInternalUUID() uuid.UUID {
return uuid.Must(uuid.FromBytes(m.GetInternalUuid()))
}

func (m *Player) SetInternalUUID(u uuid.UUID) {
m.InternalUuid = u[:]
}

func (m *Player) GetSessionUUID() uuid.UUID {
return uuid.Must(uuid.FromBytes(m.GetSessionUuid()))
}

func (m *Player) SetSessionUUID(u uuid.UUID) {
m.SessionUuid = u[:]
}

```

## Installation

Each language-specific plugin is located in its own folder under `cmd/` with a `README.md` file.

Follow the instructions in the respective folder to build and use the plugin:

- [Go Plugin](cmd/protoc-gen-uuidhelper-go/README.md)
- [Kotlin Plugin](cmd/protoc-gen-uuidhelper-kotlin/README.md)

## Adding a New Language

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
      - The writer has to implement what should be generated for the file header (e.g. package, imports, etc.) as well as what methods to generate for each UUID field.

    Note: The writer will only be called for those fields that actually are considered "UUID-Fields", therefore you don't have to filter the fields yourself, as that's already done by the core implementation.

    For generating the methods, there are various helper methods in the `core` package that can be used e.g. convert a fields name from Protobuf' snake_case to the target language's camelCase or PascalCase.

    Refer to the Go implementation in `cmd/protoc-gen-uuidhelper-go/plugin.go` for an example of how to structure the plugin.


3. **Build the Plugin**  
   Build the plugin using the following command:
   ```bash
   go build -o protoc-gen-uuidhelper-<lang> ./cmd/protoc-gen-uuidhelper-<lang>
    ```

## License

MIT License. See [LICENSE](LICENSE) for details.
