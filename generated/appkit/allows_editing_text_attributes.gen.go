
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [allowsEditingTextAttributes] class.
var allowsEditingTextAttributesClass _allowsEditingTextAttributesClass

func init() {
	allowsEditingTextAttributesClass = _allowsEditingTextAttributesClass{objc.GetClass("allowsEditingTextAttributes")}
}

type _allowsEditingTextAttributesClass struct {
	objc.Class
}

// An interface definition for the [allowsEditingTextAttributes] class.
type IallowsEditingTextAttributes interface {
	ID() objc.ID
}

type allowsEditingTextAttributes struct {
	id objc.ID
}

func allowsEditingTextAttributesFrom(ptr unsafe.Pointer) allowsEditingTextAttributes {
	return allowsEditingTextAttributes{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ allowsEditingTextAttributes) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _allowsEditingTextAttributesClass) Alloc() allowsEditingTextAttributes {
	rv := objc.Send[allowsEditingTextAttributes](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _allowsEditingTextAttributesClass) New() allowsEditingTextAttributes {
	rv := objc.Send[allowsEditingTextAttributes](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewallowsEditingTextAttributes creates and returns a new initialized instance.
func NewallowsEditingTextAttributes() allowsEditingTextAttributes {
	return allowsEditingTextAttributesClass.New()
}

// Init initializes the instance.
func (a_ allowsEditingTextAttributes) Init() allowsEditingTextAttributes {
	rv := objc.Send[allowsEditingTextAttributes](a_.ID(), selInit)
	return rv
}
