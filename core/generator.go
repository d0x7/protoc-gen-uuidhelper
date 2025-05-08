package core

import (
	"google.golang.org/protobuf/compiler/protogen"
)

type UUIDHelperBackend interface {
	// OpenFile is called once per file that has UUID fields to generate
	OpenFile(gen *protogen.Plugin, file *protogen.File) UUIDFileWriter

	// Version returns the version of the plugin
	Version() string
}

// UUIDFileWriter will be called by the core when a file with UUIDs is detected
type UUIDFileWriter interface {
	// GenerateFileHeader will be called once at the beginning
	GenerateFileHeader()

	// GenerateSingleField will be called for each UUID field found in a message
	GenerateSingleField(msg *protogen.Message, field *protogen.Field)

	// GenerateListField will be called for each repeated UUID field found in a message
	GenerateListField(msg *protogen.Message, field *protogen.Field)

	// Close will be called after all UUIDs are handled; should return the generated file
	Close()
}

type GeneratorSetup interface {
	// Setup is called once at the beginning of the plugin
	Setup(gen *protogen.Plugin)
}
