package core

import (
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"strings"
	"unicode"
)

func DescriptorToCamelCase(desc protoreflect.Descriptor) string {
	return SnakeToCamelCase(string(desc.Name()))
}

// SnakeToCamelCase converts a snake_case string to CamelCase.
// For example, converts "player_uuid" to "PlayerUUID"
func SnakeToCamelCase(s string) string {
	parts := strings.Split(s, "_")
	for i, p := range parts {
		//goland:noinspection GoDeprecation
		parts[i] = strings.Title(p)
	}
	return strings.Join(parts, "")
}

func DescriptorToLowerCamelCase(desc protoreflect.Descriptor) string {
	return SnakeToLowerCamelCase(string(desc.Name()))
}

func SnakeToLowerCamelCase(s string) string {
	parts := strings.Split(s, "_")
	for i := range parts {
		if i == 0 {
			continue // first part stays lowercase
		}
		//goland:noinspection GoDeprecation
		parts[i] = strings.Title(parts[i])
	}
	return strings.Join(parts, "")
}

// SnakeToPascalCase converts a snake_case string to PascalCase.
// For example, converts "player_uuid" to "PlayerUUID"
func SnakeToPascalCase(s string) string {
	var words []string
	start := 0

	for i, r := range s {
		if r == '_' {
			if start < i {
				words = append(words, Capitalize(s[start:i]))
			}
			start = i + 1
		}
	}
	if start < len(s) {
		words = append(words, Capitalize(s[start:]))
	}

	return strings.Join(words, "")
}

// Capitalize capitalizes the first letter of a string.
// For example, convert "hello" to "Hello"
func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func ProtocVersion(gen *protogen.Plugin) string {
	v := gen.Request.GetCompilerVersion()
	if v == nil {
		return "(unknown)"
	}
	var suffix string
	if s := v.GetSuffix(); s != "" {
		suffix = "-" + s
	}
	return fmt.Sprintf("%d.%d.%d%s", v.GetMajor(), v.GetMinor(), v.GetPatch(), suffix)
}
