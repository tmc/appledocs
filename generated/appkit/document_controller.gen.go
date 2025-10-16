
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [DocumentController] class.
var DocumentControllerClass _DocumentControllerClass

func init() {
	DocumentControllerClass = _DocumentControllerClass{objc.GetClass("NSDocumentController")}
}

type _DocumentControllerClass struct {
	objc.Class
}

// An interface definition for the [DocumentController] class.
type IDocumentController interface {
	ID() objc.ID
}

type DocumentController struct {
	id objc.ID
}

func DocumentControllerFrom(ptr unsafe.Pointer) DocumentController {
	return DocumentController{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ DocumentController) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _DocumentControllerClass) Alloc() DocumentController {
	rv := objc.Send[DocumentController](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _DocumentControllerClass) New() DocumentController {
	rv := objc.Send[DocumentController](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewDocumentController creates and returns a new initialized instance.
func NewDocumentController() DocumentController {
	return DocumentControllerClass.New()
}

// Init initializes the instance.
func (d_ DocumentController) Init() DocumentController {
	rv := objc.Send[DocumentController](d_.ID(), selInit)
	return rv
}
