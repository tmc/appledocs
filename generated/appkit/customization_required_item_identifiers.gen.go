
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [customizationRequiredItemIdentifiers] class.
var customizationRequiredItemIdentifiersClass _customizationRequiredItemIdentifiersClass

func init() {
	customizationRequiredItemIdentifiersClass = _customizationRequiredItemIdentifiersClass{objc.GetClass("customizationRequiredItemIdentifiers")}
}

type _customizationRequiredItemIdentifiersClass struct {
	objc.Class
}

// An interface definition for the [customizationRequiredItemIdentifiers] class.
type IcustomizationRequiredItemIdentifiers interface {
	ID() objc.ID
}

type customizationRequiredItemIdentifiers struct {
	id objc.ID
}

func customizationRequiredItemIdentifiersFrom(ptr unsafe.Pointer) customizationRequiredItemIdentifiers {
	return customizationRequiredItemIdentifiers{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ customizationRequiredItemIdentifiers) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _customizationRequiredItemIdentifiersClass) Alloc() customizationRequiredItemIdentifiers {
	rv := objc.Send[customizationRequiredItemIdentifiers](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _customizationRequiredItemIdentifiersClass) New() customizationRequiredItemIdentifiers {
	rv := objc.Send[customizationRequiredItemIdentifiers](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcustomizationRequiredItemIdentifiers creates and returns a new initialized instance.
func NewcustomizationRequiredItemIdentifiers() customizationRequiredItemIdentifiers {
	return customizationRequiredItemIdentifiersClass.New()
}

// Init initializes the instance.
func (c_ customizationRequiredItemIdentifiers) Init() customizationRequiredItemIdentifiers {
	rv := objc.Send[customizationRequiredItemIdentifiers](c_.ID(), selInit)
	return rv
}
