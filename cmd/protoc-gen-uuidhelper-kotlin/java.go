package main

import (
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"path"
	"xiam.li/uuidhelper/core"
)

type JavaImport struct {
	Class      string
	SubClass   string
	KtSubClass string
}

func getProtoFile(method *protogen.Message) *protogen.File {
	file, ok := protoFiles[method.Desc.ParentFile()]
	if !ok {
		panic(fmt.Sprintf("File not found for method %v", method.Desc.FullName()))
	}
	return file
}

func toJavaImport(message *protogen.Message) JavaImport {
	protoFile := getProtoFile(message)
	class, subClass := getJavaImport(protoFile, message.Desc)
	return JavaImport{Class: class, SubClass: subClass, KtSubClass: subClass + "Kt"}
}

func getJavaImport(file *protogen.File, desc protoreflect.MessageDescriptor) (class, subClass string) {
	javaClassname := file.GeneratedFilenamePrefix
	if file.Proto.GetOptions().GetJavaOuterClassname() != "" {
		javaClassname = file.Proto.GetOptions().GetJavaOuterClassname()
	}

	if file.Proto.GetOptions().GetJavaMultipleFiles() {
		class = string(desc.Name())
	} else {
		class = core.SnakeToPascalCase(path.Base(javaClassname))
		subClass = string(desc.Name())
	}

	return
}
