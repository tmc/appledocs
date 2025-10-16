
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [exerciseAmbiguityInLayout] class.
var exerciseAmbiguityInLayoutClass _exerciseAmbiguityInLayoutClass

func init() {
	exerciseAmbiguityInLayoutClass = _exerciseAmbiguityInLayoutClass{objc.GetClass("exerciseAmbiguityInLayout")}
}

type _exerciseAmbiguityInLayoutClass struct {
	objc.Class
}

// An interface definition for the [exerciseAmbiguityInLayout] class.
type IexerciseAmbiguityInLayout interface {
	ID() objc.ID
}

type exerciseAmbiguityInLayout struct {
	id objc.ID
}

func exerciseAmbiguityInLayoutFrom(ptr unsafe.Pointer) exerciseAmbiguityInLayout {
	return exerciseAmbiguityInLayout{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (e_ exerciseAmbiguityInLayout) ID() objc.ID {
	return e_.id
}

// Alloc allocates a new instance without initialization.
func (ec _exerciseAmbiguityInLayoutClass) Alloc() exerciseAmbiguityInLayout {
	rv := objc.Send[exerciseAmbiguityInLayout](objc.ID(ec.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ec _exerciseAmbiguityInLayoutClass) New() exerciseAmbiguityInLayout {
	rv := objc.Send[exerciseAmbiguityInLayout](objc.ID(ec.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewexerciseAmbiguityInLayout creates and returns a new initialized instance.
func NewexerciseAmbiguityInLayout() exerciseAmbiguityInLayout {
	return exerciseAmbiguityInLayoutClass.New()
}

// Init initializes the instance.
func (e_ exerciseAmbiguityInLayout) Init() exerciseAmbiguityInLayout {
	rv := objc.Send[exerciseAmbiguityInLayout](e_.ID(), selInit)
	return rv
}
