
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [canSelectHiddenExtension] class.
var canSelectHiddenExtensionClass _canSelectHiddenExtensionClass

func init() {
	canSelectHiddenExtensionClass = _canSelectHiddenExtensionClass{objc.GetClass("canSelectHiddenExtension")}
}

type _canSelectHiddenExtensionClass struct {
	objc.Class
}

// An interface definition for the [canSelectHiddenExtension] class.
type IcanSelectHiddenExtension interface {
	ID() objc.ID
}

type canSelectHiddenExtension struct {
	id objc.ID
}

func canSelectHiddenExtensionFrom(ptr unsafe.Pointer) canSelectHiddenExtension {
	return canSelectHiddenExtension{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ canSelectHiddenExtension) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _canSelectHiddenExtensionClass) Alloc() canSelectHiddenExtension {
	rv := objc.Send[canSelectHiddenExtension](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _canSelectHiddenExtensionClass) New() canSelectHiddenExtension {
	rv := objc.Send[canSelectHiddenExtension](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcanSelectHiddenExtension creates and returns a new initialized instance.
func NewcanSelectHiddenExtension() canSelectHiddenExtension {
	return canSelectHiddenExtensionClass.New()
}

// Init initializes the instance.
func (c_ canSelectHiddenExtension) Init() canSelectHiddenExtension {
	rv := objc.Send[canSelectHiddenExtension](c_.ID(), selInit)
	return rv
}
