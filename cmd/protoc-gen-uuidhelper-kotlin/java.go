package main

//type JavaImportPackage string
//
//func (j JavaImportPackage) String() string {
//	return string(j)
//}
//
//func (j JavaImportPackage) Ident(class string) JavaImport {
//	return JavaImport{Package: j, Class: class}
//}
//
//func (j JavaImportPackage) IdentSub(class, subClass string) JavaImport {
//	return JavaImport{Package: j, Class: class, SubClass: subClass, KtSubClass: subClass + "Kt"}
//}

type JavaImport struct {
	//Package    JavaImportPackage
	Class      string
	SubClass   string
	KtSubClass string
}

//func (j JavaImport) FQDN() string {
//	fqdn := string(j.Package)
//	fqdn += "." + j.Class
//	if j.SubClass != "" {
//		fqdn += "." + j.SubClass
//	}
//	return fqdn
//}

//func (j JavaImport) Is(other JavaImport) bool {
//	return j.Package == other.Package && j.Class == other.Class && j.SubClass == other.SubClass
//}
