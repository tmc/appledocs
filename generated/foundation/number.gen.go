// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Number] class.
var numberClass = _NumberClass{objc.GetClass("NSNumber")}

type _NumberClass struct {
	class objc.Class
}

// An object wrapper for primitive scalar numeric values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber

type Number struct {
	Value
}

// NumberFrom constructs a [Number] from an unsafe.Pointer.
//
// An object wrapper for primitive scalar numeric values.
func NumberFrom(ptr unsafe.Pointer) Number {
	return Number{
		Value: ValueFrom(ptr),
	}
}

// Returns a string that represents the contents of the number object for a given locale. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/description(withLocale:)
func (n_ Number) DescriptionWithLocale(locale objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("descriptionWithLocale:"), locale)
	return rv
}


