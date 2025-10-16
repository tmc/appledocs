
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextViewportLayoutController] class.
var TextViewportLayoutControllerClass _TextViewportLayoutControllerClass

func init() {
	TextViewportLayoutControllerClass = _TextViewportLayoutControllerClass{objc.GetClass("NSTextViewportLayoutController")}
}

type _TextViewportLayoutControllerClass struct {
	objc.Class
}

// An interface definition for the [TextViewportLayoutController] class.
type ITextViewportLayoutController interface {
	ID() objc.ID
}

type TextViewportLayoutController struct {
	id objc.ID
}

func TextViewportLayoutControllerFrom(ptr unsafe.Pointer) TextViewportLayoutController {
	return TextViewportLayoutController{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TextViewportLayoutController) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TextViewportLayoutControllerClass) Alloc() TextViewportLayoutController {
	rv := objc.Send[TextViewportLayoutController](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TextViewportLayoutControllerClass) New() TextViewportLayoutController {
	rv := objc.Send[TextViewportLayoutController](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTextViewportLayoutController creates and returns a new initialized instance.
func NewTextViewportLayoutController() TextViewportLayoutController {
	return TextViewportLayoutControllerClass.New()
}

// Init initializes the instance.
func (t_ TextViewportLayoutController) Init() TextViewportLayoutController {
	rv := objc.Send[TextViewportLayoutController](t_.ID(), selInit)
	return rv
}
