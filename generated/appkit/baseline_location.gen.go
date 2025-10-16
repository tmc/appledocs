
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [baselineLocation] class.
var baselineLocationClass _baselineLocationClass

func init() {
	baselineLocationClass = _baselineLocationClass{objc.GetClass("baselineLocation")}
}

type _baselineLocationClass struct {
	objc.Class
}

// An interface definition for the [baselineLocation] class.
type IbaselineLocation interface {
	ID() objc.ID
}

type baselineLocation struct {
	id objc.ID
}

func baselineLocationFrom(ptr unsafe.Pointer) baselineLocation {
	return baselineLocation{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (b_ baselineLocation) ID() objc.ID {
	return b_.id
}

// Alloc allocates a new instance without initialization.
func (bc _baselineLocationClass) Alloc() baselineLocation {
	rv := objc.Send[baselineLocation](objc.ID(bc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (bc _baselineLocationClass) New() baselineLocation {
	rv := objc.Send[baselineLocation](objc.ID(bc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewbaselineLocation creates and returns a new initialized instance.
func NewbaselineLocation() baselineLocation {
	return baselineLocationClass.New()
}

// Init initializes the instance.
func (b_ baselineLocation) Init() baselineLocation {
	rv := objc.Send[baselineLocation](b_.ID(), selInit)
	return rv
}
