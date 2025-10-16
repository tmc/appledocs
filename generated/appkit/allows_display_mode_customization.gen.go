
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [allowsDisplayModeCustomization] class.
var allowsDisplayModeCustomizationClass _allowsDisplayModeCustomizationClass

func init() {
	allowsDisplayModeCustomizationClass = _allowsDisplayModeCustomizationClass{objc.GetClass("allowsDisplayModeCustomization")}
}

type _allowsDisplayModeCustomizationClass struct {
	objc.Class
}

// An interface definition for the [allowsDisplayModeCustomization] class.
type IallowsDisplayModeCustomization interface {
	ID() objc.ID
}

type allowsDisplayModeCustomization struct {
	id objc.ID
}

func allowsDisplayModeCustomizationFrom(ptr unsafe.Pointer) allowsDisplayModeCustomization {
	return allowsDisplayModeCustomization{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ allowsDisplayModeCustomization) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _allowsDisplayModeCustomizationClass) Alloc() allowsDisplayModeCustomization {
	rv := objc.Send[allowsDisplayModeCustomization](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _allowsDisplayModeCustomizationClass) New() allowsDisplayModeCustomization {
	rv := objc.Send[allowsDisplayModeCustomization](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewallowsDisplayModeCustomization creates and returns a new initialized instance.
func NewallowsDisplayModeCustomization() allowsDisplayModeCustomization {
	return allowsDisplayModeCustomizationClass.New()
}

// Init initializes the instance.
func (a_ allowsDisplayModeCustomization) Init() allowsDisplayModeCustomization {
	rv := objc.Send[allowsDisplayModeCustomization](a_.ID(), selInit)
	return rv
}
