
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [centeredItemIdentifiers] class.
var centeredItemIdentifiersClass _centeredItemIdentifiersClass

func init() {
	centeredItemIdentifiersClass = _centeredItemIdentifiersClass{objc.GetClass("centeredItemIdentifiers")}
}

type _centeredItemIdentifiersClass struct {
	objc.Class
}

// An interface definition for the [centeredItemIdentifiers] class.
type IcenteredItemIdentifiers interface {
	ID() objc.ID
}

type centeredItemIdentifiers struct {
	id objc.ID
}

func centeredItemIdentifiersFrom(ptr unsafe.Pointer) centeredItemIdentifiers {
	return centeredItemIdentifiers{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ centeredItemIdentifiers) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _centeredItemIdentifiersClass) Alloc() centeredItemIdentifiers {
	rv := objc.Send[centeredItemIdentifiers](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _centeredItemIdentifiersClass) New() centeredItemIdentifiers {
	rv := objc.Send[centeredItemIdentifiers](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcenteredItemIdentifiers creates and returns a new initialized instance.
func NewcenteredItemIdentifiers() centeredItemIdentifiers {
	return centeredItemIdentifiersClass.New()
}

// Init initializes the instance.
func (c_ centeredItemIdentifiers) Init() centeredItemIdentifiers {
	rv := objc.Send[centeredItemIdentifiers](c_.ID(), selInit)
	return rv
}
