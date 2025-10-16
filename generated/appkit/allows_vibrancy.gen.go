
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [allowsVibrancy] class.
var allowsVibrancyClass _allowsVibrancyClass

func init() {
	allowsVibrancyClass = _allowsVibrancyClass{objc.GetClass("allowsVibrancy")}
}

type _allowsVibrancyClass struct {
	objc.Class
}

// An interface definition for the [allowsVibrancy] class.
type IallowsVibrancy interface {
	ID() objc.ID
}

type allowsVibrancy struct {
	id objc.ID
}

func allowsVibrancyFrom(ptr unsafe.Pointer) allowsVibrancy {
	return allowsVibrancy{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ allowsVibrancy) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _allowsVibrancyClass) Alloc() allowsVibrancy {
	rv := objc.Send[allowsVibrancy](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _allowsVibrancyClass) New() allowsVibrancy {
	rv := objc.Send[allowsVibrancy](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewallowsVibrancy creates and returns a new initialized instance.
func NewallowsVibrancy() allowsVibrancy {
	return allowsVibrancyClass.New()
}

// Init initializes the instance.
func (a_ allowsVibrancy) Init() allowsVibrancy {
	rv := objc.Send[allowsVibrancy](a_.ID(), selInit)
	return rv
}
