
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isRotatedFromBase] class.
var isRotatedFromBaseClass _isRotatedFromBaseClass

func init() {
	isRotatedFromBaseClass = _isRotatedFromBaseClass{objc.GetClass("isRotatedFromBase")}
}

type _isRotatedFromBaseClass struct {
	objc.Class
}

// An interface definition for the [isRotatedFromBase] class.
type IisRotatedFromBase interface {
	ID() objc.ID
}

type isRotatedFromBase struct {
	id objc.ID
}

func isRotatedFromBaseFrom(ptr unsafe.Pointer) isRotatedFromBase {
	return isRotatedFromBase{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isRotatedFromBase) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isRotatedFromBaseClass) Alloc() isRotatedFromBase {
	rv := objc.Send[isRotatedFromBase](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isRotatedFromBaseClass) New() isRotatedFromBase {
	rv := objc.Send[isRotatedFromBase](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisRotatedFromBase creates and returns a new initialized instance.
func NewisRotatedFromBase() isRotatedFromBase {
	return isRotatedFromBaseClass.New()
}

// Init initializes the instance.
func (i_ isRotatedFromBase) Init() isRotatedFromBase {
	rv := objc.Send[isRotatedFromBase](i_.ID(), selInit)
	return rv
}
