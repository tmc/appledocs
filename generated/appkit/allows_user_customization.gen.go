
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [allowsUserCustomization] class.
var allowsUserCustomizationClass _allowsUserCustomizationClass

func init() {
	allowsUserCustomizationClass = _allowsUserCustomizationClass{objc.GetClass("allowsUserCustomization")}
}

type _allowsUserCustomizationClass struct {
	objc.Class
}

// An interface definition for the [allowsUserCustomization] class.
type IallowsUserCustomization interface {
	ID() objc.ID
}

type allowsUserCustomization struct {
	id objc.ID
}

func allowsUserCustomizationFrom(ptr unsafe.Pointer) allowsUserCustomization {
	return allowsUserCustomization{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ allowsUserCustomization) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _allowsUserCustomizationClass) Alloc() allowsUserCustomization {
	rv := objc.Send[allowsUserCustomization](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _allowsUserCustomizationClass) New() allowsUserCustomization {
	rv := objc.Send[allowsUserCustomization](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewallowsUserCustomization creates and returns a new initialized instance.
func NewallowsUserCustomization() allowsUserCustomization {
	return allowsUserCustomizationClass.New()
}

// Init initializes the instance.
func (a_ allowsUserCustomization) Init() allowsUserCustomization {
	rv := objc.Send[allowsUserCustomization](a_.ID(), selInit)
	return rv
}
