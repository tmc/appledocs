
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [depthLimit] class.
var depthLimitClass _depthLimitClass

func init() {
	depthLimitClass = _depthLimitClass{objc.GetClass("depthLimit")}
}

type _depthLimitClass struct {
	objc.Class
}

// An interface definition for the [depthLimit] class.
type IdepthLimit interface {
	ID() objc.ID
}

type depthLimit struct {
	id objc.ID
}

func depthLimitFrom(ptr unsafe.Pointer) depthLimit {
	return depthLimit{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ depthLimit) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _depthLimitClass) Alloc() depthLimit {
	rv := objc.Send[depthLimit](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _depthLimitClass) New() depthLimit {
	rv := objc.Send[depthLimit](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewdepthLimit creates and returns a new initialized instance.
func NewdepthLimit() depthLimit {
	return depthLimitClass.New()
}

// Init initializes the instance.
func (d_ depthLimit) Init() depthLimit {
	rv := objc.Send[depthLimit](d_.ID(), selInit)
	return rv
}
