
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [verticalScrollElasticity] class.
var verticalScrollElasticityClass _verticalScrollElasticityClass

func init() {
	verticalScrollElasticityClass = _verticalScrollElasticityClass{objc.GetClass("verticalScrollElasticity")}
}

type _verticalScrollElasticityClass struct {
	objc.Class
}

// An interface definition for the [verticalScrollElasticity] class.
type IverticalScrollElasticity interface {
	ID() objc.ID
}

type verticalScrollElasticity struct {
	id objc.ID
}

func verticalScrollElasticityFrom(ptr unsafe.Pointer) verticalScrollElasticity {
	return verticalScrollElasticity{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (v_ verticalScrollElasticity) ID() objc.ID {
	return v_.id
}

// Alloc allocates a new instance without initialization.
func (vc _verticalScrollElasticityClass) Alloc() verticalScrollElasticity {
	rv := objc.Send[verticalScrollElasticity](objc.ID(vc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (vc _verticalScrollElasticityClass) New() verticalScrollElasticity {
	rv := objc.Send[verticalScrollElasticity](objc.ID(vc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewverticalScrollElasticity creates and returns a new initialized instance.
func NewverticalScrollElasticity() verticalScrollElasticity {
	return verticalScrollElasticityClass.New()
}

// Init initializes the instance.
func (v_ verticalScrollElasticity) Init() verticalScrollElasticity {
	rv := objc.Send[verticalScrollElasticity](v_.ID(), selInit)
	return rv
}
