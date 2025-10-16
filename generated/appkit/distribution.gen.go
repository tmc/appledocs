
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [distribution] class.
var distributionClass _distributionClass

func init() {
	distributionClass = _distributionClass{objc.GetClass("distribution")}
}

type _distributionClass struct {
	objc.Class
}

// An interface definition for the [distribution] class.
type Idistribution interface {
	ID() objc.ID
}

type distribution struct {
	id objc.ID
}

func distributionFrom(ptr unsafe.Pointer) distribution {
	return distribution{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ distribution) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _distributionClass) Alloc() distribution {
	rv := objc.Send[distribution](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _distributionClass) New() distribution {
	rv := objc.Send[distribution](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newdistribution creates and returns a new initialized instance.
func Newdistribution() distribution {
	return distributionClass.New()
}

// Init initializes the instance.
func (d_ distribution) Init() distribution {
	rv := objc.Send[distribution](d_.ID(), selInit)
	return rv
}
