// Code generated by "ctenum create enum color.yaml" version dev; DO NOT EDIT.

package examples

import (
	"encoding"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// InvalidColorEnumError is returned by [ToColorEnum] if the passed in
// string isn't one of the enum values.
type InvalidColorEnumError struct {
	s string
}

// Error exposes the invalid string used and shows all the possible values.
func (e InvalidColorEnumError) Error() string {
	return strconv.Quote(e.s) + " does not belong to ColorEnum enum; " +
		"valid values are: " + strings.Join(ColorEnumStrings(), " ")
}

// Is tells you if target can be unwrapped to be an InvalidColorEnumError.
func (e InvalidColorEnumError) Is(err error) bool {
	return errors.As(err, &e)
}

// ColorEnum is a true enum, i.e. using this as a parameter in a function
// declaration requires the use of one of the constants defined in this package.
type ColorEnum interface {
	isColorEnum()
	fmt.Stringer
	encoding.TextMarshaler
}

type _ColorEnum string

// isColorEnum implements ColorEnum to make it a part of the compile-time
// safe enum.
func (_ColorEnum) isColorEnum() {}

// String converts the private enum value to a primitive string.
func (x _ColorEnum) String() string {
	return string(x)
}

const (
	// ColorEnumRed Represents the color of fire and blood, associated with energy and passion.
	ColorEnumRed _ColorEnum = "Red"
	// ColorEnumGreen Associated with nature, growth, and harmony.
	ColorEnumGreen _ColorEnum = "Green"
	// ColorEnumBlue Often linked to the sky and the sea, symbolizing calm and tranquility.
	ColorEnumBlue _ColorEnum = "Blue"
	// ColorEnumCyan A blend of blue and green, known for its soothing properties.
	ColorEnumCyan _ColorEnum = "Cyan"
	// ColorEnumMagenta A vibrant and intense color often associated with creativity.
	ColorEnumMagenta _ColorEnum = "Magenta"
	// ColorEnumYellow Represents sunshine and happiness, associated with positivity.
	ColorEnumYellow _ColorEnum = "Yellow"
	// ColorEnumWhite Symbolizes purity and cleanliness.
	ColorEnumWhite _ColorEnum = "White"
	// ColorEnumBlack Represents mystery, sophistication, and formality.
	ColorEnumBlack _ColorEnum = "Black"
)

var _ColorEnumMapping = map[string]_ColorEnum{
	"Red":     ColorEnumRed,
	"Green":   ColorEnumGreen,
	"Blue":    ColorEnumBlue,
	"Cyan":    ColorEnumCyan,
	"Magenta": ColorEnumMagenta,
	"Yellow":  ColorEnumYellow,
	"White":   ColorEnumWhite,
	"Black":   ColorEnumBlack,
}

// ToColorEnum retrieves an enum value from the constants defined in this
// package. It tries it's best to find s in the enum by using the original-case,
// uppercase and lowercase.
func ToColorEnum(s string) (ColorEnum, error) {
	if val, ok := _ColorEnumMapping[s]; ok {
		return val, nil
	}
	if val, ok := _ColorEnumMapping[strings.ToLower(s)]; ok {
		return val, nil
	}
	if val, ok := _ColorEnumMapping[strings.ToUpper(s)]; ok {
		return val, nil
	}
	return nil, InvalidColorEnumError{s: s}
}

// ColorEnumValues returns a slice of values of the ColorEnum enum.
// Each enum value matches the index of [ColorEnumStrings].
func ColorEnumValues() []ColorEnum {
	return []ColorEnum{
		ColorEnumRed,
		ColorEnumGreen,
		ColorEnumBlue,
		ColorEnumCyan,
		ColorEnumMagenta,
		ColorEnumYellow,
		ColorEnumWhite,
		ColorEnumBlack,
	}
}

// ColorEnumStrings returns a slice of all enum values as strings. This could
// be useful for validation with something like [slices.Contains] or quickly
// seeing all values of an enum.
// Each string matches the index of [ColorEnumValues].
func ColorEnumStrings() []string {
	return []string{
		"Red",
		"Green",
		"Blue",
		"Cyan",
		"Magenta",
		"Yellow",
		"White",
		"Black",
	}
}

// MarshalText implements the [encoding.TextMarshaler] for ColorEnum.
func (x _ColorEnum) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

type ColorEnumUnmarshaler string

func (x ColorEnumUnmarshaler) Extract() ColorEnum {
	if x == "" {
		return nil
	}
	v, err := ToColorEnum(string(x))
	if err != nil {
		panic(fmt.Sprintf("Incorrect usage of %T!"+
			"Used for unmarshalling into the application", x))
	}
	return v
}

func (x *ColorEnumUnmarshaler) UnmarshalText(text []byte) error {
	v, err := ToColorEnum(string(text))
	if err != nil {
		return err
	}
	*x = ColorEnumUnmarshaler(any(v).(_ColorEnum))
	return nil
}
