
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isExpanded] class.
var isExpandedClass _isExpandedClass

func init() {
	isExpandedClass = _isExpandedClass{objc.GetClass("isExpanded")}
}

type _isExpandedClass struct {
	objc.Class
}

// An interface definition for the [isExpanded] class.
type IisExpanded interface {
	ID() objc.ID
}

type isExpanded struct {
	id objc.ID
}

func isExpandedFrom(ptr unsafe.Pointer) isExpanded {
	return isExpanded{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isExpanded) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isExpandedClass) Alloc() isExpanded {
	rv := objc.Send[isExpanded](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isExpandedClass) New() isExpanded {
	rv := objc.Send[isExpanded](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisExpanded creates and returns a new initialized instance.
func NewisExpanded() isExpanded {
	return isExpandedClass.New()
}

// Init initializes the instance.
func (i_ isExpanded) Init() isExpanded {
	rv := objc.Send[isExpanded](i_.ID(), selInit)
	return rv
}
