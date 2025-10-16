
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isIndeterminate] class.
var isIndeterminateClass _isIndeterminateClass

func init() {
	isIndeterminateClass = _isIndeterminateClass{objc.GetClass("isIndeterminate")}
}

type _isIndeterminateClass struct {
	objc.Class
}

// An interface definition for the [isIndeterminate] class.
type IisIndeterminate interface {
	ID() objc.ID
}

type isIndeterminate struct {
	id objc.ID
}

func isIndeterminateFrom(ptr unsafe.Pointer) isIndeterminate {
	return isIndeterminate{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isIndeterminate) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isIndeterminateClass) Alloc() isIndeterminate {
	rv := objc.Send[isIndeterminate](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isIndeterminateClass) New() isIndeterminate {
	rv := objc.Send[isIndeterminate](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisIndeterminate creates and returns a new initialized instance.
func NewisIndeterminate() isIndeterminate {
	return isIndeterminateClass.New()
}

// Init initializes the instance.
func (i_ isIndeterminate) Init() isIndeterminate {
	rv := objc.Send[isIndeterminate](i_.ID(), selInit)
	return rv
}
