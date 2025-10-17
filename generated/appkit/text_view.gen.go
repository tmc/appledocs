
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextView] class.
var TextViewClass _TextViewClass

func init() {
	TextViewClass = _TextViewClass{objc.GetClass("NSTextView")}
}

type _TextViewClass struct {
	objc.Class
}

// An interface definition for the [TextView] class.
type ITextView interface {
	ID() objc.ID
}

type TextView struct {
	id objc.ID
}

func TextViewFrom(ptr unsafe.Pointer) TextView {
	return TextView{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TextView) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TextViewClass) Alloc() TextView {
	rv := objc.Send[TextView](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TextViewClass) New() TextView {
	rv := objc.Send[TextView](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTextView creates and returns a new initialized instance.
func NewTextView() TextView {
	return TextViewClass.New()
}

// Init initializes the instance.
func (t_ TextView) Init() TextView {
	rv := objc.Send[TextView](t_.ID(), selInit)
	return rv
}
// The layout manager that lays out text for the receiver’s text container. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTextView/layoutManager
func (t_ TextView) LayoutManager() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("layoutManager"))
	return rv
}
// The receiver’s text container. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTextView/textContainer
func (t_ TextView) TextContainer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("textContainer"))
	return rv
}
// SetTextContainer sets the value of the textContainer property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTextView/textContainer
func (t_ TextView) SetTextContainer(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setTextContainer:"), value)
}
// The receiver’s text storage object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTextView/textStorage
func (t_ TextView) TextStorage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("textStorage"))
	return rv
}
