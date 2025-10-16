
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [hasDynamicDepthLimit] class.
var hasDynamicDepthLimitClass _hasDynamicDepthLimitClass

func init() {
	hasDynamicDepthLimitClass = _hasDynamicDepthLimitClass{objc.GetClass("hasDynamicDepthLimit")}
}

type _hasDynamicDepthLimitClass struct {
	objc.Class
}

// An interface definition for the [hasDynamicDepthLimit] class.
type IhasDynamicDepthLimit interface {
	ID() objc.ID
}

type hasDynamicDepthLimit struct {
	id objc.ID
}

func hasDynamicDepthLimitFrom(ptr unsafe.Pointer) hasDynamicDepthLimit {
	return hasDynamicDepthLimit{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (h_ hasDynamicDepthLimit) ID() objc.ID {
	return h_.id
}

// Alloc allocates a new instance without initialization.
func (hc _hasDynamicDepthLimitClass) Alloc() hasDynamicDepthLimit {
	rv := objc.Send[hasDynamicDepthLimit](objc.ID(hc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (hc _hasDynamicDepthLimitClass) New() hasDynamicDepthLimit {
	rv := objc.Send[hasDynamicDepthLimit](objc.ID(hc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewhasDynamicDepthLimit creates and returns a new initialized instance.
func NewhasDynamicDepthLimit() hasDynamicDepthLimit {
	return hasDynamicDepthLimitClass.New()
}

// Init initializes the instance.
func (h_ hasDynamicDepthLimit) Init() hasDynamicDepthLimit {
	rv := objc.Send[hasDynamicDepthLimit](h_.ID(), selInit)
	return rv
}
