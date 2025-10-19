// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DocumentController] class.
var (
	documentControllerClass     _DocumentControllerClass
	documentControllerClassOnce sync.Once
)

func getDocumentControllerClass() _DocumentControllerClass {
	documentControllerClassOnce.Do(func() {
		documentControllerClass = _DocumentControllerClass{objc.GetClass("NSDocumentController")}
	})
	return documentControllerClass
}

type _DocumentControllerClass struct {
	class objc.Class
}

// An interface definition for the [DocumentController] class.
type IDocumentController interface {
	objectivec.IObject
}

// An object that manages an app’s documents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController

type DocumentController struct {
	objectivec.Object
}

// DocumentControllerFrom constructs a [DocumentController] from an unsafe.Pointer.
//
// An object that manages an app’s documents.
func DocumentControllerFrom(ptr unsafe.Pointer) DocumentController {
	return DocumentController{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (dc _DocumentControllerClass) Alloc() DocumentController {
	rv := objc.Send[DocumentController](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DocumentControllerClass) New() DocumentController {
	rv := objc.Send[DocumentController](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DocumentController) Init() DocumentController {
	rv := objc.Send[DocumentController](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DocumentController) Autorelease() DocumentController {
	rv := objc.Send[DocumentController](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDocumentController creates a new DocumentController instance.
func NewDocumentController() DocumentController {
	return getDocumentControllerClass().New()
}




