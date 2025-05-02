package main

import (
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"path"
	"xiam.li/uuidhelper/core"
)

func getProtoFile(method *protogen.Message) *protogen.File {
	file, ok := protoFiles[method.Desc.ParentFile()]
	if !ok {
		panic(fmt.Sprintf("File not found for method %v", method.Desc.FullName()))
	}
	return file
}

func toJavaImport(message *protogen.Message) JavaImport {
	protoFile := getProtoFile(message)
	/*pkg,*/ class, subClass := getJavaImport(protoFile, message.Desc)
	return JavaImport{ /*Package: JavaImportPackage(pkg),*/ Class: class, SubClass: subClass, KtSubClass: subClass + "Kt"}
}

func getJavaImport(file *protogen.File, desc protoreflect.MessageDescriptor) ( /*pkg,*/ class, subClass string) {
	/*pkg = file.Proto.GetOptions().GetJavaPackage()
	if pkg == "" {
		if protoPkg := file.Proto.GetPackage(); protoPkg != "" {
			pkg = protoPkg
		} else {
			pkg = "proto"
		}
	}*/

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
