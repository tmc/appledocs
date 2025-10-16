
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextCheckingController] class.
var TextCheckingControllerClass _TextCheckingControllerClass

func init() {
	TextCheckingControllerClass = _TextCheckingControllerClass{objc.GetClass("NSTextCheckingController")}
}

type _TextCheckingControllerClass struct {
	objc.Class
}

// An interface definition for the [TextCheckingController] class.
type ITextCheckingController interface {
	ID() objc.ID
}

type TextCheckingController struct {
	id objc.ID
}

func TextCheckingControllerFrom(ptr unsafe.Pointer) TextCheckingController {
	return TextCheckingController{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TextCheckingController) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TextCheckingControllerClass) Alloc() TextCheckingController {
	rv := objc.Send[TextCheckingController](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TextCheckingControllerClass) New() TextCheckingController {
	rv := objc.Send[TextCheckingController](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTextCheckingController creates and returns a new initialized instance.
func NewTextCheckingController() TextCheckingController {
	return TextCheckingControllerClass.New()
}

// Init initializes the instance.
func (t_ TextCheckingController) Init() TextCheckingController {
	rv := objc.Send[TextCheckingController](t_.ID(), selInit)
	return rv
}
