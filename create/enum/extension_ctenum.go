// Code generated by "ctenum create enum build.go" version dev; DO NOT EDIT.

package enum

import (
	"encoding"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// InvalidExtensionEnumError is returned by [ToExtensionEnum] if the passed in
// string isn't one of the enum values.
type InvalidExtensionEnumError struct {
	s string
}

// Error exposes the invalid string used and shows all the possible values.
func (e InvalidExtensionEnumError) Error() string {
	return strconv.Quote(e.s) + " does not belong to ExtensionEnum enum; " +
		"valid values are: " + strings.Join(ExtensionEnumStrings(), " ")
}

// Is tells you if target can be unwrapped to be an InvalidExtensionEnumError.
func (e InvalidExtensionEnumError) Is(err error) bool {
	return errors.As(err, &e)
}

// ExtensionEnum is a true enum, i.e. using this as a parameter in a function
// declaration requires the use of one of the constants defined in this package.
type ExtensionEnum interface {
	isExtensionEnum()
	fmt.Stringer
	encoding.TextMarshaler
}

type _ExtensionEnum string

// isExtensionEnum implements ExtensionEnum to make it a part of the compile-time
// safe enum.
func (_ExtensionEnum) isExtensionEnum() {}

// String converts the private enum value to a primitive string.
func (x _ExtensionEnum) String() string {
	return string(x)
}

const (
	ExtensionEnumYml  _ExtensionEnum = ".yml"
	ExtensionEnumYAML _ExtensionEnum = ".yaml"
	ExtensionEnumJSON _ExtensionEnum = ".json"
	ExtensionEnumGo   _ExtensionEnum = ".go"
)

var _ExtensionEnumMapping = map[string]_ExtensionEnum{
	".yml":  ExtensionEnumYml,
	".yaml": ExtensionEnumYAML,
	".json": ExtensionEnumJSON,
	".go":   ExtensionEnumGo,
}

// ToExtensionEnum retrieves an enum value from the constants defined in this
// package. It tries it's best to find s in the enum by using the original-case,
// uppercase and lowercase.
func ToExtensionEnum(s string) (ExtensionEnum, error) {
	if val, ok := _ExtensionEnumMapping[s]; ok {
		return val, nil
	}
	if val, ok := _ExtensionEnumMapping[strings.ToLower(s)]; ok {
		return val, nil
	}
	if val, ok := _ExtensionEnumMapping[strings.ToUpper(s)]; ok {
		return val, nil
	}
	return nil, InvalidExtensionEnumError{s: s}
}

// ExtensionEnumValues returns a slice of values of the ExtensionEnum enum.
// Each enum value matches the index of [ExtensionEnumStrings].
func ExtensionEnumValues() []ExtensionEnum {
	return []ExtensionEnum{
		ExtensionEnumYml,
		ExtensionEnumYAML,
		ExtensionEnumJSON,
		ExtensionEnumGo,
	}
}

// ExtensionEnumStrings returns a slice of all enum values as strings. This could
// be useful for validation with something like [slices.Contains] or quickly
// seeing all values of an enum.
// Each string matches the index of [ExtensionEnumValues].
func ExtensionEnumStrings() []string {
	return []string{
		".yml",
		".yaml",
		".json",
		".go",
	}
}

// MarshalText implements the [encoding.TextMarshaler] for ExtensionEnum.
func (x _ExtensionEnum) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

type ExtensionEnumUnmarshaler string

func (x ExtensionEnumUnmarshaler) Extract() ExtensionEnum {
	if x == "" {
		return nil
	}
	v, err := ToExtensionEnum(string(x))
	if err != nil {
		panic(fmt.Sprintf("Incorrect usage of %T!"+
			"Used for unmarshalling into the application", x))
	}
	return v
}

func (x *ExtensionEnumUnmarshaler) UnmarshalText(text []byte) error {
	v, err := ToExtensionEnum(string(text))
	if err != nil {
		return err
	}
	*x = ExtensionEnumUnmarshaler(any(v).(_ExtensionEnum))
	return nil
}
