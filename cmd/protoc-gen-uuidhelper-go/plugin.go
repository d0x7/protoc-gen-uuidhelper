package main

import (
	"google.golang.org/protobuf/compiler/protogen"
	"xiam.li/uuidhelper/core"
)

type backend struct{}

func main() {
	core.Main(backend{})
}

func (backend) OpenFile(gen *protogen.Plugin, file *protogen.File) core.UUIDFileWriter {
	filename := file.GeneratedFilenamePrefix + "_uuidhelper.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	return &goFileWriter{gen, file, g}
}
