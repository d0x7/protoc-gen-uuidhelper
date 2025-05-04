package core

import (
	"flag"
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
	"os"
	"path/filepath"
	"strings"
)

func Main(generator UUIDHelperBackend) {
	MainWithFlags(&flag.FlagSet{}, generator)
}

func MainWithFlags(flags *flag.FlagSet, generator UUIDHelperBackend) {
	var showVersion bool
	flag.BoolVar(&showVersion, "version", false, "print the version and exit")
	flag.BoolVar(&showVersion, "v", false, "print the version and exit")
	flag.Parse()
	if showVersion {
		fmt.Printf("%s %s\n", filepath.Base(os.Args[0]), generator.Version())
		return
	}

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		// Check if the generator implements the GeneratorSetup interface
		if generatorSetup, ok := generator.(GeneratorSetup); ok {
			generatorSetup.Setup(gen)
		}

		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL |
			pluginpb.CodeGeneratorResponse_FEATURE_SUPPORTS_EDITIONS)
		gen.SupportedEditionsMinimum = descriptorpb.Edition_EDITION_PROTO3
		gen.SupportedEditionsMaximum = descriptorpb.Edition_EDITION_2023

		for _, file := range gen.Files {
			if !file.Generate {
				continue
			}

			// Check if the file contains any *_uuid fields
			var needsGeneration bool
			for _, msg := range file.Messages {
				for _, field := range msg.Fields {
					if isUUIDField(field) {
						needsGeneration = true
						break
					} else if isUUIDsField(field) {
						needsGeneration = true
						break
					}
				}
				if needsGeneration {
					break
				}
			}

			if !needsGeneration {
				continue
			}

			writer := generator.OpenFile(gen, file)
			writer.GenerateFileHeader()

			for _, msg := range file.Messages {
				for _, field := range msg.Fields {
					if isUUIDField(field) {
						writer.GenerateUUIDHelper(msg, field)
					} else if isUUIDsField(field) {
						writer.GenerateUUIDsHelper(msg, field)
					}
				}
			}

			writer.Close()
		}
		return nil
	})
}

func isUUIDsField(field *protogen.Field) bool {
	return field.Desc.Kind() == protoreflect.BytesKind && strings.HasSuffix(string(field.Desc.Name()), "_uuids") && field.Desc.IsList()
}

func isUUIDField(field *protogen.Field) bool {
	return field.Desc.Kind() == protoreflect.BytesKind && strings.HasSuffix(string(field.Desc.Name()), "_uuid")
}
