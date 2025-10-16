
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [customizationAllowedItemIdentifiers] class.
var customizationAllowedItemIdentifiersClass _customizationAllowedItemIdentifiersClass

func init() {
	customizationAllowedItemIdentifiersClass = _customizationAllowedItemIdentifiersClass{objc.GetClass("customizationAllowedItemIdentifiers")}
}

type _customizationAllowedItemIdentifiersClass struct {
	objc.Class
}

// An interface definition for the [customizationAllowedItemIdentifiers] class.
type IcustomizationAllowedItemIdentifiers interface {
	ID() objc.ID
}

type customizationAllowedItemIdentifiers struct {
	id objc.ID
}

func customizationAllowedItemIdentifiersFrom(ptr unsafe.Pointer) customizationAllowedItemIdentifiers {
	return customizationAllowedItemIdentifiers{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ customizationAllowedItemIdentifiers) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _customizationAllowedItemIdentifiersClass) Alloc() customizationAllowedItemIdentifiers {
	rv := objc.Send[customizationAllowedItemIdentifiers](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _customizationAllowedItemIdentifiersClass) New() customizationAllowedItemIdentifiers {
	rv := objc.Send[customizationAllowedItemIdentifiers](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcustomizationAllowedItemIdentifiers creates and returns a new initialized instance.
func NewcustomizationAllowedItemIdentifiers() customizationAllowedItemIdentifiers {
	return customizationAllowedItemIdentifiersClass.New()
}

// Init initializes the instance.
func (c_ customizationAllowedItemIdentifiers) Init() customizationAllowedItemIdentifiers {
	rv := objc.Send[customizationAllowedItemIdentifiers](c_.ID(), selInit)
	return rv
}
