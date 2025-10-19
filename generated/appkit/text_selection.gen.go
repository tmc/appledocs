// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextSelection] class.
var (
	textSelectionClass     _TextSelectionClass
	textSelectionClassOnce sync.Once
)

func getTextSelectionClass() _TextSelectionClass {
	textSelectionClassOnce.Do(func() {
		textSelectionClass = _TextSelectionClass{objc.GetClass("NSTextSelection")}
	})
	return textSelectionClass
}

type _TextSelectionClass struct {
	class objc.Class
}

// An interface definition for the [TextSelection] class.
type ITextSelection interface {
	objectivec.IObject
}

// A class that represents a single logical selection context that corresponds to an insertion point.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelection
type TextSelection struct {
	objectivec.Object
}

// TextSelectionFrom constructs a [TextSelection] from an unsafe.Pointer.
//
// A class that represents a single logical selection context that corresponds to an insertion point.
func TextSelectionFrom(ptr unsafe.Pointer) TextSelection {
	return TextSelection{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TextSelectionClass) Alloc() TextSelection {
	rv := objc.Send[TextSelection](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TextSelectionClass) New() TextSelection {
	rv := objc.Send[TextSelection](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextSelection) Init() TextSelection {
	rv := objc.Send[TextSelection](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextSelection) Autorelease() TextSelection {
	rv := objc.Send[TextSelection](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextSelection creates a new TextSelection instance.
func NewTextSelection() TextSelection {
	return getTextSelectionClass().New()
}




