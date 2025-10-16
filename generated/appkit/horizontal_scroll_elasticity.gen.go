
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [horizontalScrollElasticity] class.
var horizontalScrollElasticityClass _horizontalScrollElasticityClass

func init() {
	horizontalScrollElasticityClass = _horizontalScrollElasticityClass{objc.GetClass("horizontalScrollElasticity")}
}

type _horizontalScrollElasticityClass struct {
	objc.Class
}

// An interface definition for the [horizontalScrollElasticity] class.
type IhorizontalScrollElasticity interface {
	ID() objc.ID
}

type horizontalScrollElasticity struct {
	id objc.ID
}

func horizontalScrollElasticityFrom(ptr unsafe.Pointer) horizontalScrollElasticity {
	return horizontalScrollElasticity{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (h_ horizontalScrollElasticity) ID() objc.ID {
	return h_.id
}

// Alloc allocates a new instance without initialization.
func (hc _horizontalScrollElasticityClass) Alloc() horizontalScrollElasticity {
	rv := objc.Send[horizontalScrollElasticity](objc.ID(hc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (hc _horizontalScrollElasticityClass) New() horizontalScrollElasticity {
	rv := objc.Send[horizontalScrollElasticity](objc.ID(hc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewhorizontalScrollElasticity creates and returns a new initialized instance.
func NewhorizontalScrollElasticity() horizontalScrollElasticity {
	return horizontalScrollElasticityClass.New()
}

// Init initializes the instance.
func (h_ horizontalScrollElasticity) Init() horizontalScrollElasticity {
	rv := objc.Send[horizontalScrollElasticity](h_.ID(), selInit)
	return rv
}
