
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isFlushWindowDisabled] class.
var isFlushWindowDisabledClass _isFlushWindowDisabledClass

func init() {
	isFlushWindowDisabledClass = _isFlushWindowDisabledClass{objc.GetClass("isFlushWindowDisabled")}
}

type _isFlushWindowDisabledClass struct {
	objc.Class
}

// An interface definition for the [isFlushWindowDisabled] class.
type IisFlushWindowDisabled interface {
	ID() objc.ID
}

type isFlushWindowDisabled struct {
	id objc.ID
}

func isFlushWindowDisabledFrom(ptr unsafe.Pointer) isFlushWindowDisabled {
	return isFlushWindowDisabled{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isFlushWindowDisabled) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isFlushWindowDisabledClass) Alloc() isFlushWindowDisabled {
	rv := objc.Send[isFlushWindowDisabled](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isFlushWindowDisabledClass) New() isFlushWindowDisabled {
	rv := objc.Send[isFlushWindowDisabled](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisFlushWindowDisabled creates and returns a new initialized instance.
func NewisFlushWindowDisabled() isFlushWindowDisabled {
	return isFlushWindowDisabledClass.New()
}

// Init initializes the instance.
func (i_ isFlushWindowDisabled) Init() isFlushWindowDisabled {
	rv := objc.Send[isFlushWindowDisabled](i_.ID(), selInit)
	return rv
}
