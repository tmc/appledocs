
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [usesPredominantAxisScrolling] class.
var usesPredominantAxisScrollingClass _usesPredominantAxisScrollingClass

func init() {
	usesPredominantAxisScrollingClass = _usesPredominantAxisScrollingClass{objc.GetClass("usesPredominantAxisScrolling")}
}

type _usesPredominantAxisScrollingClass struct {
	objc.Class
}

// An interface definition for the [usesPredominantAxisScrolling] class.
type IusesPredominantAxisScrolling interface {
	ID() objc.ID
}

type usesPredominantAxisScrolling struct {
	id objc.ID
}

func usesPredominantAxisScrollingFrom(ptr unsafe.Pointer) usesPredominantAxisScrolling {
	return usesPredominantAxisScrolling{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (u_ usesPredominantAxisScrolling) ID() objc.ID {
	return u_.id
}

// Alloc allocates a new instance without initialization.
func (uc _usesPredominantAxisScrollingClass) Alloc() usesPredominantAxisScrolling {
	rv := objc.Send[usesPredominantAxisScrolling](objc.ID(uc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (uc _usesPredominantAxisScrollingClass) New() usesPredominantAxisScrolling {
	rv := objc.Send[usesPredominantAxisScrolling](objc.ID(uc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewusesPredominantAxisScrolling creates and returns a new initialized instance.
func NewusesPredominantAxisScrolling() usesPredominantAxisScrolling {
	return usesPredominantAxisScrollingClass.New()
}

// Init initializes the instance.
func (u_ usesPredominantAxisScrolling) Init() usesPredominantAxisScrolling {
	rv := objc.Send[usesPredominantAxisScrolling](u_.ID(), selInit)
	return rv
}
