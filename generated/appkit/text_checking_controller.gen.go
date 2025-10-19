// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextCheckingController] class.
var (
	textCheckingControllerClass     _TextCheckingControllerClass
	textCheckingControllerClassOnce sync.Once
)

func getTextCheckingControllerClass() _TextCheckingControllerClass {
	textCheckingControllerClassOnce.Do(func() {
		textCheckingControllerClass = _TextCheckingControllerClass{objc.GetClass("NSTextCheckingController")}
	})
	return textCheckingControllerClass
}

type _TextCheckingControllerClass struct {
	class objc.Class
}

// An interface definition for the [TextCheckingController] class.
type ITextCheckingController interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextCheckingController

type TextCheckingController struct {
	objectivec.Object
}

// TextCheckingControllerFrom constructs a [TextCheckingController] from an unsafe.Pointer.
func TextCheckingControllerFrom(ptr unsafe.Pointer) TextCheckingController {
	return TextCheckingController{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (tc _TextCheckingControllerClass) Alloc() TextCheckingController {
	rv := objc.Send[TextCheckingController](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (tc _TextCheckingControllerClass) New() TextCheckingController {
	rv := objc.Send[TextCheckingController](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextCheckingController) Init() TextCheckingController {
	rv := objc.Send[TextCheckingController](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextCheckingController) Autorelease() TextCheckingController {
	rv := objc.Send[TextCheckingController](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextCheckingController creates a new TextCheckingController instance.
func NewTextCheckingController() TextCheckingController {
	return getTextCheckingControllerClass().New()
}




