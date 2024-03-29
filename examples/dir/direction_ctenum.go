// Code generated by "ctenum create enum -package dir -output dir/direction_ctenum.go direction.yaml" version dev; DO NOT EDIT.

package dir

import (
	"encoding"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// InvalidDirectionEnumError is returned by [ToDirectionEnum] if the passed in
// string isn't one of the enum values.
type InvalidDirectionEnumError struct {
	s string
}

// Error exposes the invalid string used and shows all the possible values.
func (e InvalidDirectionEnumError) Error() string {
	return strconv.Quote(e.s) + " does not belong to DirectionEnum enum; " +
		"valid values are: " + strings.Join(DirectionEnumStrings(), " ")
}

// Is tells you if target can be unwrapped to be an InvalidDirectionEnumError.
func (e InvalidDirectionEnumError) Is(err error) bool {
	return errors.As(err, &e)
}

// DirectionEnum is a true enum, i.e. using this as a parameter in a function
// declaration requires the use of one of the constants defined in this package.
type DirectionEnum interface {
	isDirectionEnum()
	fmt.Stringer
	encoding.TextMarshaler
}

type _DirectionEnum string

// isDirectionEnum implements DirectionEnum to make it a part of the compile-time
// safe enum.
func (_DirectionEnum) isDirectionEnum() {}

// String converts the private enum value to a primitive string.
func (x _DirectionEnum) String() string {
	return string(x)
}

const (
	DirectionEnumN _DirectionEnum = "North"
	DirectionEnumS _DirectionEnum = "South"
	DirectionEnumE _DirectionEnum = "East"
	DirectionEnumW _DirectionEnum = "West"
)

var _DirectionEnumMapping = map[string]_DirectionEnum{
	"North": DirectionEnumN,
	"South": DirectionEnumS,
	"East":  DirectionEnumE,
	"West":  DirectionEnumW,
}

// ToDirectionEnum retrieves an enum value from the constants defined in this
// package. It tries it's best to find s in the enum by using the original-case,
// uppercase and lowercase.
func ToDirectionEnum(s string) (DirectionEnum, error) {
	if val, ok := _DirectionEnumMapping[s]; ok {
		return val, nil
	}
	if val, ok := _DirectionEnumMapping[strings.ToLower(s)]; ok {
		return val, nil
	}
	if val, ok := _DirectionEnumMapping[strings.ToUpper(s)]; ok {
		return val, nil
	}
	return nil, InvalidDirectionEnumError{s: s}
}

// DirectionEnumValues returns a slice of values of the DirectionEnum enum.
// Each enum value matches the index of [DirectionEnumStrings].
func DirectionEnumValues() []DirectionEnum {
	return []DirectionEnum{
		DirectionEnumN,
		DirectionEnumS,
		DirectionEnumE,
		DirectionEnumW,
	}
}

// DirectionEnumStrings returns a slice of all enum values as strings. This could
// be useful for validation with something like [slices.Contains] or quickly
// seeing all values of an enum.
// Each string matches the index of [DirectionEnumValues].
func DirectionEnumStrings() []string {
	return []string{
		"North",
		"South",
		"East",
		"West",
	}
}

// MarshalText implements the [encoding.TextMarshaler] for DirectionEnum.
func (x _DirectionEnum) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

type DirectionEnumUnmarshaler string

func (x DirectionEnumUnmarshaler) Extract() DirectionEnum {
	if x == "" {
		return nil
	}
	v, err := ToDirectionEnum(string(x))
	if err != nil {
		panic(fmt.Sprintf("Incorrect usage of %T!"+
			"Used for unmarshalling into the application", x))
	}
	return v
}

func (x *DirectionEnumUnmarshaler) UnmarshalText(text []byte) error {
	v, err := ToDirectionEnum(string(text))
	if err != nil {
		return err
	}
	*x = DirectionEnumUnmarshaler(any(v).(_DirectionEnum))
	return nil
}
