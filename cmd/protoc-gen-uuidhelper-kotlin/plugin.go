package main

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"path"
	"strings"
	"xiam.li/uuidhelper/core"
)

type backend struct {
}

var protoFiles map[protoreflect.FileDescriptor]*protogen.File

func main() {
	core.Main(&backend{})
}

func (b backend) Setup(gen *protogen.Plugin) {
	protoFiles = make(map[protoreflect.FileDescriptor]*protogen.File, len(gen.Files))
	for _, file := range gen.Files {
		protoFiles[file.Desc] = file
	}
}

func (b backend) OpenFile(gen *protogen.Plugin, file *protogen.File) core.UUIDFileWriter {
	// Retrieve Java package and class options
	javaPackage := file.Proto.GetOptions().GetJavaPackage()
	if javaPackage == "" {
		if protoPkg := file.Proto.GetPackage(); protoPkg != "" {
			javaPackage = protoPkg
		} else {
			javaPackage = "pb" // Fallback if no java_package is defined
		}
	}
	javaClassname := file.Proto.GetOptions().GetJavaOuterClassname()
	if javaClassname == "" {
		//goland:noinspection GoDeprecation
		javaClassname = strings.Title(file.GeneratedFilenamePrefix) // Generate a classname if not defined
	}
	javaClassname += "UUIDHelperKt"

	// Define the output filename
	packagePath := strings.ReplaceAll(javaPackage, ".", "/")
	outputFileName := path.Join(packagePath, javaClassname+".kt")

	g := gen.NewGeneratedFile(outputFileName, "")
	return &kotlinFileWriter{gen, file, g}
}
