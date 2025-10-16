
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [defaultDepthLimit] class.
var defaultDepthLimitClass _defaultDepthLimitClass

func init() {
	defaultDepthLimitClass = _defaultDepthLimitClass{objc.GetClass("defaultDepthLimit")}
}

type _defaultDepthLimitClass struct {
	objc.Class
}

// An interface definition for the [defaultDepthLimit] class.
type IdefaultDepthLimit interface {
	ID() objc.ID
}

type defaultDepthLimit struct {
	id objc.ID
}

func defaultDepthLimitFrom(ptr unsafe.Pointer) defaultDepthLimit {
	return defaultDepthLimit{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ defaultDepthLimit) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _defaultDepthLimitClass) Alloc() defaultDepthLimit {
	rv := objc.Send[defaultDepthLimit](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _defaultDepthLimitClass) New() defaultDepthLimit {
	rv := objc.Send[defaultDepthLimit](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewdefaultDepthLimit creates and returns a new initialized instance.
func NewdefaultDepthLimit() defaultDepthLimit {
	return defaultDepthLimitClass.New()
}

// Init initializes the instance.
func (d_ defaultDepthLimit) Init() defaultDepthLimit {
	rv := objc.Send[defaultDepthLimit](d_.ID(), selInit)
	return rv
}
