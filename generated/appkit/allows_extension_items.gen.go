
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [allowsExtensionItems] class.
var allowsExtensionItemsClass _allowsExtensionItemsClass

func init() {
	allowsExtensionItemsClass = _allowsExtensionItemsClass{objc.GetClass("allowsExtensionItems")}
}

type _allowsExtensionItemsClass struct {
	objc.Class
}

// An interface definition for the [allowsExtensionItems] class.
type IallowsExtensionItems interface {
	ID() objc.ID
}

type allowsExtensionItems struct {
	id objc.ID
}

func allowsExtensionItemsFrom(ptr unsafe.Pointer) allowsExtensionItems {
	return allowsExtensionItems{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ allowsExtensionItems) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _allowsExtensionItemsClass) Alloc() allowsExtensionItems {
	rv := objc.Send[allowsExtensionItems](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _allowsExtensionItemsClass) New() allowsExtensionItems {
	rv := objc.Send[allowsExtensionItems](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewallowsExtensionItems creates and returns a new initialized instance.
func NewallowsExtensionItems() allowsExtensionItems {
	return allowsExtensionItemsClass.New()
}

// Init initializes the instance.
func (a_ allowsExtensionItems) Init() allowsExtensionItems {
	rv := objc.Send[allowsExtensionItems](a_.ID(), selInit)
	return rv
}
