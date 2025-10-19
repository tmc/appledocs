// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextInputContext] class.
var (
	textInputContextClass     _TextInputContextClass
	textInputContextClassOnce sync.Once
)

func getTextInputContextClass() _TextInputContextClass {
	textInputContextClassOnce.Do(func() {
		textInputContextClass = _TextInputContextClass{objc.GetClass("NSTextInputContext")}
	})
	return textInputContextClass
}

type _TextInputContextClass struct {
	class objc.Class
}

// An interface definition for the [TextInputContext] class.
type ITextInputContext interface {
	objectivec.IObject
}

// An object that represents the Cocoa text input system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInputContext

type TextInputContext struct {
	objectivec.Object
}

// TextInputContextFrom constructs a [TextInputContext] from an unsafe.Pointer.
//
// An object that represents the Cocoa text input system.
func TextInputContextFrom(ptr unsafe.Pointer) TextInputContext {
	return TextInputContext{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (tc _TextInputContextClass) Alloc() TextInputContext {
	rv := objc.Send[TextInputContext](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (tc _TextInputContextClass) New() TextInputContext {
	rv := objc.Send[TextInputContext](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextInputContext) Init() TextInputContext {
	rv := objc.Send[TextInputContext](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextInputContext) Autorelease() TextInputContext {
	rv := objc.Send[TextInputContext](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextInputContext creates a new TextInputContext instance.
func NewTextInputContext() TextInputContext {
	return getTextInputContextClass().New()
}




