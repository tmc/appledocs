// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextLineFragment] class.
var textLineFragmentClass = _TextLineFragmentClass{objc.GetClass("NSTextLineFragment")}

type _TextLineFragmentClass struct {
	class objc.Class
}

// An interface definition for the [TextLineFragment] class.
type ITextLineFragment interface {
	objectivec.IObject
}

// A class that represents a line fragment as a single textual layout and rendering unit inside a text layout fragment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLineFragment

type TextLineFragment struct {
	objectivec.Object
}

// TextLineFragmentFrom constructs a [TextLineFragment] from an unsafe.Pointer.
//
// A class that represents a line fragment as a single textual layout and rendering unit inside a text layout fragment.
func TextLineFragmentFrom(ptr unsafe.Pointer) TextLineFragment {
	return TextLineFragment{objectivec.Object{objc.ID(ptr)}}
}



