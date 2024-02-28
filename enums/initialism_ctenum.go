// Code generated by "ctenum create enum initialism.go" version dev; DO NOT EDIT.

package enums

import (
	"encoding"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// InvalidInitialismEnumError is returned by [ToInitialismEnum] if the passed in
// string isn't one of the enum values.
type InvalidInitialismEnumError struct {
	s string
}

// Error exposes the invalid string used and shows all the possible values.
func (e InvalidInitialismEnumError) Error() string {
	return strconv.Quote(e.s) + " does not belong to InitialismEnum enum; " +
		"valid values are: " + strings.Join(InitialismEnumStrings(), " ")
}

// Is tells you if target can be unwrapped to be an InvalidInitialismEnumError.
func (e InvalidInitialismEnumError) Is(err error) bool {
	return errors.As(err, &e)
}

// InitialismEnum is a true enum, i.e. using this as a parameter in a function
// declaration requires the use of one of the constants defined in this package.
type InitialismEnum interface {
	isInitialismEnum()
	fmt.Stringer
	encoding.TextMarshaler
}

type _InitialismEnum string

// isInitialismEnum implements InitialismEnum to make it a part of the compile-time
// safe enum.
func (_InitialismEnum) isInitialismEnum() {}

// String converts the private enum value to a primitive string.
func (x _InitialismEnum) String() string {
	return string(x)
}

const (
	InitialismEnumACL _InitialismEnum = "ACL"
	InitialismEnumAPI _InitialismEnum = "API"
	InitialismEnumASCII _InitialismEnum = "ASCII"
	InitialismEnumCPU _InitialismEnum = "CPU"
	// InitialismEnumCSS Cascading style sheets.
	InitialismEnumCSS _InitialismEnum = "CSS"
	// InitialismEnumEOF eof is End of File.
	InitialismEnumEOF _InitialismEnum = "EOF"
	InitialismEnumGPU _InitialismEnum = "GPU"
	InitialismEnumGUID _InitialismEnum = "GUID"
	// InitialismEnumHTML HyperTextMarkupLanguage.
	InitialismEnumHTML _InitialismEnum = "HTML"
	InitialismEnumHTTP _InitialismEnum = "HTTP"
	InitialismEnumHTTPS _InitialismEnum = "HTTPS"
	InitialismEnumID _InitialismEnum = "ID"
	InitialismEnumIP _InitialismEnum = "IP"
	InitialismEnumJSON _InitialismEnum = "JSON"
	InitialismEnumJWT _InitialismEnum = "JWT"
	InitialismEnumRAM _InitialismEnum = "RAM"
	InitialismEnumREST _InitialismEnum = "REST"
	InitialismEnumSQL _InitialismEnum = "SQL"
	InitialismEnumTCP _InitialismEnum = "TCP"
	InitialismEnumTLS _InitialismEnum = "TLS"
	InitialismEnumUDP _InitialismEnum = "UDP"
	InitialismEnumURI _InitialismEnum = "URI"
	InitialismEnumURL _InitialismEnum = "URL"
	InitialismEnumUTF _InitialismEnum = "UTF"
	InitialismEnumUUID _InitialismEnum = "UUID"
	InitialismEnumXML _InitialismEnum = "XML"
	InitialismEnumYAML _InitialismEnum = "YAML"
	InitialismEnumZIP _InitialismEnum = "ZIP"
)

var _InitialismEnumMapping = map[string]_InitialismEnum{
	"ACL":   InitialismEnumACL,
	"API":   InitialismEnumAPI,
	"ASCII": InitialismEnumASCII,
	"CPU":   InitialismEnumCPU,
	"CSS":   InitialismEnumCSS,
	"EOF":   InitialismEnumEOF,
	"GPU":   InitialismEnumGPU,
	"GUID":  InitialismEnumGUID,
	"HTML":  InitialismEnumHTML,
	"HTTP":  InitialismEnumHTTP,
	"HTTPS": InitialismEnumHTTPS,
	"ID":    InitialismEnumID,
	"IP":    InitialismEnumIP,
	"JSON":  InitialismEnumJSON,
	"JWT":   InitialismEnumJWT,
	"RAM":   InitialismEnumRAM,
	"REST":  InitialismEnumREST,
	"SQL":   InitialismEnumSQL,
	"TCP":   InitialismEnumTCP,
	"TLS":   InitialismEnumTLS,
	"UDP":   InitialismEnumUDP,
	"URI":   InitialismEnumURI,
	"URL":   InitialismEnumURL,
	"UTF":   InitialismEnumUTF,
	"UUID":  InitialismEnumUUID,
	"XML":   InitialismEnumXML,
	"YAML":  InitialismEnumYAML,
	"ZIP":   InitialismEnumZIP,
}

// ToInitialismEnum retrieves an enum value from the constants defined in this
// package. It tries it's best to find s in the enum by using the original-case,
// uppercase and lowercase.
func ToInitialismEnum(s string) (InitialismEnum, error) {
	if val, ok := _InitialismEnumMapping[s]; ok {
		return val, nil
	}
	if val, ok := _InitialismEnumMapping[strings.ToLower(s)]; ok {
		return val, nil
	}
	if val, ok := _InitialismEnumMapping[strings.ToUpper(s)]; ok {
		return val, nil
	}
	return nil, InvalidInitialismEnumError{s: s}
}

// InitialismEnumValues returns a slice of values of the InitialismEnum enum.
// Each enum value matches the index of [InitialismEnumStrings].
func InitialismEnumValues() []InitialismEnum {
	return []InitialismEnum{
		InitialismEnumACL,
		InitialismEnumAPI,
		InitialismEnumASCII,
		InitialismEnumCPU,
		InitialismEnumCSS,
		InitialismEnumEOF,
		InitialismEnumGPU,
		InitialismEnumGUID,
		InitialismEnumHTML,
		InitialismEnumHTTP,
		InitialismEnumHTTPS,
		InitialismEnumID,
		InitialismEnumIP,
		InitialismEnumJSON,
		InitialismEnumJWT,
		InitialismEnumRAM,
		InitialismEnumREST,
		InitialismEnumSQL,
		InitialismEnumTCP,
		InitialismEnumTLS,
		InitialismEnumUDP,
		InitialismEnumURI,
		InitialismEnumURL,
		InitialismEnumUTF,
		InitialismEnumUUID,
		InitialismEnumXML,
		InitialismEnumYAML,
		InitialismEnumZIP,
	}
}

// InitialismEnumStrings returns a slice of all enum values as strings. This could
// be useful for validation with something like [slices.Contains] or quickly
// seeing all values of an enum.
// Each string matches the index of [InitialismEnumValues].
func InitialismEnumStrings() []string {
	return []string{
		"ACL",
		"API",
		"ASCII",
		"CPU",
		"CSS",
		"EOF",
		"GPU",
		"GUID",
		"HTML",
		"HTTP",
		"HTTPS",
		"ID",
		"IP",
		"JSON",
		"JWT",
		"RAM",
		"REST",
		"SQL",
		"TCP",
		"TLS",
		"UDP",
		"URI",
		"URL",
		"UTF",
		"UUID",
		"XML",
		"YAML",
		"ZIP",
	}
}

// MarshalText implements the [encoding.TextMarshaler] for InitialismEnum.
func (x _InitialismEnum) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}
