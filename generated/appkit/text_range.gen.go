// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextRange] class.
var textRangeClass = _TextRangeClass{objc.GetClass("NSTextRange")}

type _TextRangeClass struct {
	class objc.Class
}

// An interface definition for the [TextRange] class.
type ITextRange interface {
	objectivec.IObject
}

// A class that represents a contiguous range between two locations inside document contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextRange

type TextRange struct {
	objectivec.Object
}

// TextRangeFrom constructs a [TextRange] from an unsafe.Pointer.
//
// A class that represents a contiguous range between two locations inside document contents.
func TextRangeFrom(ptr unsafe.Pointer) TextRange {
	return TextRange{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (tc _TextRangeClass) Alloc() TextRange {
	rv := objc.Send[TextRange](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (tc _TextRangeClass) New() TextRange {
	rv := objc.Send[TextRange](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextRange) Init() TextRange {
	rv := objc.Send[TextRange](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextRange) Autorelease() TextRange {
	rv := objc.Send[TextRange](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextRange creates a new TextRange instance.
func NewTextRange() TextRange {
	return textRangeClass.New()
}




