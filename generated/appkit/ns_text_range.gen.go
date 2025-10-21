// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextRange] class.
var (
	TextRangeClass     _TextRangeClass
	TextRangeClassOnce sync.Once
)

func getTextRangeClass() _TextRangeClass {
	TextRangeClassOnce.Do(func() {
		TextRangeClass = _TextRangeClass{objc.GetClass("NSTextRange")}
	})
	return TextRangeClass
}

type _TextRangeClass struct {
	class objc.Class
}

// An interface definition for the [TextRange] class.
type ITextRange interface {
	objectivec.IObject
}

// A class that represents a contiguous range between two locations inside document contents.
//
// An consists of the starting and terminating locations. There the two basic properties: and , respectively. The terminating , , is directly following the last location in the range. For example, a location contains a range if is .
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getTextRangeClass().New()
}




