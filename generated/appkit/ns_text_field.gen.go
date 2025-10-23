// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextField] class.
var (
	TextFieldClass     _TextFieldClass
	TextFieldClassOnce sync.Once
)

func getTextFieldClass() _TextFieldClass {
	TextFieldClassOnce.Do(func() {
		TextFieldClass = _TextFieldClass{objc.GetClass("NSTextField")}
	})
	return TextFieldClass
}

type _TextFieldClass struct {
	class objc.Class
}

// An interface definition for the [TextField] class.
type ITextField interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A parent class referenced by other AppKit classes.


// A parent class referenced by other AppKit classes. [Full Topic]
type TextField struct {
	objectivec.Object
}

// TextFieldFrom constructs a [TextField] from an unsafe.Pointer.
//
// A parent class referenced by other AppKit classes.
func TextFieldFrom(ptr unsafe.Pointer) TextField {
	return TextField{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TextFieldClass) Alloc() TextField {
	rv := objc.Send[TextField](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TextFieldClass) New() TextField {
	rv := objc.Send[TextField](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextField) Init() TextField {
	rv := objc.Send[TextField](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextField) Autorelease() TextField {
	rv := objc.Send[TextField](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextField creates a new TextField instance.
func NewTextField() TextField {
	return getTextFieldClass().New()
}




