
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isRulerVisible] class.
var isRulerVisibleClass _isRulerVisibleClass

func init() {
	isRulerVisibleClass = _isRulerVisibleClass{objc.GetClass("isRulerVisible")}
}

type _isRulerVisibleClass struct {
	objc.Class
}

// An interface definition for the [isRulerVisible] class.
type IisRulerVisible interface {
	ID() objc.ID
}

type isRulerVisible struct {
	id objc.ID
}

func isRulerVisibleFrom(ptr unsafe.Pointer) isRulerVisible {
	return isRulerVisible{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isRulerVisible) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isRulerVisibleClass) Alloc() isRulerVisible {
	rv := objc.Send[isRulerVisible](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isRulerVisibleClass) New() isRulerVisible {
	rv := objc.Send[isRulerVisible](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisRulerVisible creates and returns a new initialized instance.
func NewisRulerVisible() isRulerVisible {
	return isRulerVisibleClass.New()
}

// Init initializes the instance.
func (i_ isRulerVisible) Init() isRulerVisible {
	rv := objc.Send[isRulerVisible](i_.ID(), selInit)
	return rv
}
