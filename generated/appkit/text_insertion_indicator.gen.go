// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TextInsertionIndicator] class.
var (
	textInsertionIndicatorClass     _TextInsertionIndicatorClass
	textInsertionIndicatorClassOnce sync.Once
)

func getTextInsertionIndicatorClass() _TextInsertionIndicatorClass {
	textInsertionIndicatorClassOnce.Do(func() {
		textInsertionIndicatorClass = _TextInsertionIndicatorClass{objc.GetClass("NSTextInsertionIndicator")}
	})
	return textInsertionIndicatorClass
}

type _TextInsertionIndicatorClass struct {
	class objc.Class
}

// An interface definition for the [TextInsertionIndicator] class.
type ITextInsertionIndicator interface {
	IView
}

// A view that represents the insertion indicator in text.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInsertionIndicator
type TextInsertionIndicator struct {
	View
}

// TextInsertionIndicatorFrom constructs a [TextInsertionIndicator] from an unsafe.Pointer.
//
// A view that represents the insertion indicator in text.
func TextInsertionIndicatorFrom(ptr unsafe.Pointer) TextInsertionIndicator {
	return TextInsertionIndicator{
		View: ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc _TextInsertionIndicatorClass) Alloc() TextInsertionIndicator {
	rv := objc.Send[TextInsertionIndicator](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TextInsertionIndicatorClass) New() TextInsertionIndicator {
	rv := objc.Send[TextInsertionIndicator](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextInsertionIndicator) Init() TextInsertionIndicator {
	rv := objc.Send[TextInsertionIndicator](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextInsertionIndicator) Autorelease() TextInsertionIndicator {
	rv := objc.Send[TextInsertionIndicator](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextInsertionIndicator creates a new TextInsertionIndicator instance.
func NewTextInsertionIndicator() TextInsertionIndicator {
	return getTextInsertionIndicatorClass().New()
}




