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

	alreadyGenerated := make(map[string]struct{})
	var generateMessage func(writer UUIDFileWriter, msg *protogen.Message)
	generateMessage = func(writer UUIDFileWriter, msg *protogen.Message) {
		if _, already := alreadyGenerated[string(msg.Desc.Name())]; already {
			return
		}
		for _, field := range msg.Fields {
			if isUUIDField(field) {
				writer.GenerateSingleField(msg, field)
			} else if isUUIDsField(field) {
				writer.GenerateListField(msg, field)
			} else if isUUIDMap(field) {
				writer.GenerateMapField(msg, field)
			} else if isEmbeddedUUIDField(field) {
				generateMessage(writer, field.Message)
			}
		}
		alreadyGenerated[string(msg.Desc.Name())] = struct{}{}
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
				generateMessage(writer, msg)
			}

			writer.Close()
		}
		return nil
	})
}

func isUUIDMap(field *protogen.Field) bool {
	return (strings.HasSuffix(string(field.Desc.Name()), "_uuid") ||
		strings.HasSuffix(string(field.Desc.Name()), "_uuids")) &&
		field.Desc.Kind() == protoreflect.MessageKind &&
		field.Desc.IsMap() &&
		field.Desc.MapValue().Kind() == protoreflect.BytesKind
}

func isEmbeddedUUIDField(field *protogen.Field) bool {
	return field.Desc.Kind() == protoreflect.MessageKind && field.Message != nil
}

func isUUIDsField(field *protogen.Field) bool {
	return field.Desc.Kind() == protoreflect.BytesKind && strings.HasSuffix(string(field.Desc.Name()), "_uuids") && field.Desc.IsList()
}

func isUUIDField(field *protogen.Field) bool {
	return field.Desc.Kind() == protoreflect.BytesKind && strings.HasSuffix(string(field.Desc.Name()), "_uuid")
}
