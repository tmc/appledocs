
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [items] class.
var itemsClass _itemsClass

func init() {
	itemsClass = _itemsClass{objc.GetClass("items")}
}

type _itemsClass struct {
	objc.Class
}

// An interface definition for the [items] class.
type Iitems interface {
	ID() objc.ID
}

type items struct {
	id objc.ID
}

func itemsFrom(ptr unsafe.Pointer) items {
	return items{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ items) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _itemsClass) Alloc() items {
	rv := objc.Send[items](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _itemsClass) New() items {
	rv := objc.Send[items](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newitems creates and returns a new initialized instance.
func Newitems() items {
	return itemsClass.New()
}

// Init initializes the instance.
func (i_ items) Init() items {
	rv := objc.Send[items](i_.ID(), selInit)
	return rv
}
