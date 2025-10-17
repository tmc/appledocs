// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ATSTypesetter] class.
var aTSTypesetterClass = _ATSTypesetterClass{objc.GetClass("NSATSTypesetter")}

type _ATSTypesetterClass struct {
	class objc.Class
}

// A concrete typesetter object that places glyphs during the text layout process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSATSTypesetter

type ATSTypesetter struct {
	Typesetter
}

// ATSTypesetterFrom constructs a [ATSTypesetter] from an unsafe.Pointer.
//
// A concrete typesetter object that places glyphs during the text layout process.
func ATSTypesetterFrom(ptr unsafe.Pointer) ATSTypesetter {
	return ATSTypesetter{
		Typesetter: TypesetterFrom(ptr),
	}
}



