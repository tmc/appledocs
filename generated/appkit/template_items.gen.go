
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [templateItems] class.
var templateItemsClass _templateItemsClass

func init() {
	templateItemsClass = _templateItemsClass{objc.GetClass("templateItems")}
}

type _templateItemsClass struct {
	objc.Class
}

// An interface definition for the [templateItems] class.
type ItemplateItems interface {
	ID() objc.ID
}

type templateItems struct {
	id objc.ID
}

func templateItemsFrom(ptr unsafe.Pointer) templateItems {
	return templateItems{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ templateItems) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _templateItemsClass) Alloc() templateItems {
	rv := objc.Send[templateItems](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _templateItemsClass) New() templateItems {
	rv := objc.Send[templateItems](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewtemplateItems creates and returns a new initialized instance.
func NewtemplateItems() templateItems {
	return templateItemsClass.New()
}

// Init initializes the instance.
func (t_ templateItems) Init() templateItems {
	rv := objc.Send[templateItems](t_.ID(), selInit)
	return rv
}
