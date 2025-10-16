
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [target] class.
var targetClass _targetClass

func init() {
	targetClass = _targetClass{objc.GetClass("target")}
}

type _targetClass struct {
	objc.Class
}

// An interface definition for the [target] class.
type Itarget interface {
	ID() objc.ID
}

type target struct {
	id objc.ID
}

func targetFrom(ptr unsafe.Pointer) target {
	return target{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ target) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _targetClass) Alloc() target {
	rv := objc.Send[target](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _targetClass) New() target {
	rv := objc.Send[target](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newtarget creates and returns a new initialized instance.
func Newtarget() target {
	return targetClass.New()
}

// Init initializes the instance.
func (t_ target) Init() target {
	rv := objc.Send[target](t_.ID(), selInit)
	return rv
}
