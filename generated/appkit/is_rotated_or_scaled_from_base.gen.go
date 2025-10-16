
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isRotatedOrScaledFromBase] class.
var isRotatedOrScaledFromBaseClass _isRotatedOrScaledFromBaseClass

func init() {
	isRotatedOrScaledFromBaseClass = _isRotatedOrScaledFromBaseClass{objc.GetClass("isRotatedOrScaledFromBase")}
}

type _isRotatedOrScaledFromBaseClass struct {
	objc.Class
}

// An interface definition for the [isRotatedOrScaledFromBase] class.
type IisRotatedOrScaledFromBase interface {
	ID() objc.ID
}

type isRotatedOrScaledFromBase struct {
	id objc.ID
}

func isRotatedOrScaledFromBaseFrom(ptr unsafe.Pointer) isRotatedOrScaledFromBase {
	return isRotatedOrScaledFromBase{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isRotatedOrScaledFromBase) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isRotatedOrScaledFromBaseClass) Alloc() isRotatedOrScaledFromBase {
	rv := objc.Send[isRotatedOrScaledFromBase](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isRotatedOrScaledFromBaseClass) New() isRotatedOrScaledFromBase {
	rv := objc.Send[isRotatedOrScaledFromBase](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisRotatedOrScaledFromBase creates and returns a new initialized instance.
func NewisRotatedOrScaledFromBase() isRotatedOrScaledFromBase {
	return isRotatedOrScaledFromBaseClass.New()
}

// Init initializes the instance.
func (i_ isRotatedOrScaledFromBase) Init() isRotatedOrScaledFromBase {
	rv := objc.Send[isRotatedOrScaledFromBase](i_.ID(), selInit)
	return rv
}
