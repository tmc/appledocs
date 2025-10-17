// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Number] class.
var NumberClass objc.Class

func init() {
	NumberClass = objc.GetClass("NSNumber")
}

type Number struct {
	objc.ID
}

func NumberFrom(ptr unsafe.Pointer) Number {
	return Number{
		ID: objc.ID(ptr),
	}
}


// Returns a string that represents the contents of the number object for a given locale. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSNumber/description(withLocale:)
func (n_ Number) DescriptionWithLocale(locale objc.ID) unsafe.Pointer {
	sel := objc.RegisterName("descriptionWithLocale:")
	ret := n_.ID.Send(sel, locale)
	return unsafe.Pointer(ret)
}

