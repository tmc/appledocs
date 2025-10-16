
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [viewDidChangeBackingProperties] class.
var viewDidChangeBackingPropertiesClass _viewDidChangeBackingPropertiesClass

func init() {
	viewDidChangeBackingPropertiesClass = _viewDidChangeBackingPropertiesClass{objc.GetClass("viewDidChangeBackingProperties")}
}

type _viewDidChangeBackingPropertiesClass struct {
	objc.Class
}

// An interface definition for the [viewDidChangeBackingProperties] class.
type IviewDidChangeBackingProperties interface {
	ID() objc.ID
}

type viewDidChangeBackingProperties struct {
	id objc.ID
}

func viewDidChangeBackingPropertiesFrom(ptr unsafe.Pointer) viewDidChangeBackingProperties {
	return viewDidChangeBackingProperties{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (v_ viewDidChangeBackingProperties) ID() objc.ID {
	return v_.id
}

// Alloc allocates a new instance without initialization.
func (vc _viewDidChangeBackingPropertiesClass) Alloc() viewDidChangeBackingProperties {
	rv := objc.Send[viewDidChangeBackingProperties](objc.ID(vc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (vc _viewDidChangeBackingPropertiesClass) New() viewDidChangeBackingProperties {
	rv := objc.Send[viewDidChangeBackingProperties](objc.ID(vc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewviewDidChangeBackingProperties creates and returns a new initialized instance.
func NewviewDidChangeBackingProperties() viewDidChangeBackingProperties {
	return viewDidChangeBackingPropertiesClass.New()
}

// Init initializes the instance.
func (v_ viewDidChangeBackingProperties) Init() viewDidChangeBackingProperties {
	rv := objc.Send[viewDidChangeBackingProperties](v_.ID(), selInit)
	return rv
}
