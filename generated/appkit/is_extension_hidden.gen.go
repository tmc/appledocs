
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isExtensionHidden] class.
var isExtensionHiddenClass _isExtensionHiddenClass

func init() {
	isExtensionHiddenClass = _isExtensionHiddenClass{objc.GetClass("isExtensionHidden")}
}

type _isExtensionHiddenClass struct {
	objc.Class
}

// An interface definition for the [isExtensionHidden] class.
type IisExtensionHidden interface {
	ID() objc.ID
}

type isExtensionHidden struct {
	id objc.ID
}

func isExtensionHiddenFrom(ptr unsafe.Pointer) isExtensionHidden {
	return isExtensionHidden{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isExtensionHidden) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isExtensionHiddenClass) Alloc() isExtensionHidden {
	rv := objc.Send[isExtensionHidden](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isExtensionHiddenClass) New() isExtensionHidden {
	rv := objc.Send[isExtensionHidden](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisExtensionHidden creates and returns a new initialized instance.
func NewisExtensionHidden() isExtensionHidden {
	return isExtensionHiddenClass.New()
}

// Init initializes the instance.
func (i_ isExtensionHidden) Init() isExtensionHidden {
	rv := objc.Send[isExtensionHidden](i_.ID(), selInit)
	return rv
}
