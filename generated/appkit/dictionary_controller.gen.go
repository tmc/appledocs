
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [DictionaryController] class.
var DictionaryControllerClass _DictionaryControllerClass

func init() {
	DictionaryControllerClass = _DictionaryControllerClass{objc.GetClass("NSDictionaryController")}
}

type _DictionaryControllerClass struct {
	objc.Class
}

// An interface definition for the [DictionaryController] class.
type IDictionaryController interface {
	ID() objc.ID
}

type DictionaryController struct {
	id objc.ID
}

func DictionaryControllerFrom(ptr unsafe.Pointer) DictionaryController {
	return DictionaryController{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ DictionaryController) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _DictionaryControllerClass) Alloc() DictionaryController {
	rv := objc.Send[DictionaryController](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _DictionaryControllerClass) New() DictionaryController {
	rv := objc.Send[DictionaryController](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewDictionaryController creates and returns a new initialized instance.
func NewDictionaryController() DictionaryController {
	return DictionaryControllerClass.New()
}

// Init initializes the instance.
func (d_ DictionaryController) Init() DictionaryController {
	rv := objc.Send[DictionaryController](d_.ID(), selInit)
	return rv
}
