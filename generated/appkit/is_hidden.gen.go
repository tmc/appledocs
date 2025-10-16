
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isHidden] class.
var isHiddenClass _isHiddenClass

func init() {
	isHiddenClass = _isHiddenClass{objc.GetClass("isHidden")}
}

type _isHiddenClass struct {
	objc.Class
}

// An interface definition for the [isHidden] class.
type IisHidden interface {
	ID() objc.ID
}

type isHidden struct {
	id objc.ID
}

func isHiddenFrom(ptr unsafe.Pointer) isHidden {
	return isHidden{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isHidden) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isHiddenClass) Alloc() isHidden {
	rv := objc.Send[isHidden](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isHiddenClass) New() isHidden {
	rv := objc.Send[isHidden](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisHidden creates and returns a new initialized instance.
func NewisHidden() isHidden {
	return isHiddenClass.New()
}

// Init initializes the instance.
func (i_ isHidden) Init() isHidden {
	rv := objc.Send[isHidden](i_.ID(), selInit)
	return rv
}
