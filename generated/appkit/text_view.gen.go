// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TextView] class.
var (
	textViewClass     _TextViewClass
	textViewClassOnce sync.Once
)

func getTextViewClass() _TextViewClass {
	textViewClassOnce.Do(func() {
		textViewClass = _TextViewClass{objc.GetClass("NSTextView")}
	})
	return textViewClass
}

type _TextViewClass struct {
	class objc.Class
}

// An interface definition for the [TextView] class.
type ITextView interface {
	IText
}

// A view that draws text and handles user interactions with that text.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView
type TextView struct {
	Text
}

// TextViewFrom constructs a [TextView] from an unsafe.Pointer.
//
// A view that draws text and handles user interactions with that text.
func TextViewFrom(ptr unsafe.Pointer) TextView {
	return TextView{
		Text: TextFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc _TextViewClass) Alloc() TextView {
	rv := objc.Send[TextView](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TextViewClass) New() TextView {
	rv := objc.Send[TextView](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextView) Init() TextView {
	rv := objc.Send[TextView](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextView) Autorelease() TextView {
	rv := objc.Send[TextView](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextView creates a new TextView instance.
func NewTextView() TextView {
	return getTextViewClass().New()
}




