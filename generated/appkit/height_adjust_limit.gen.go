
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [heightAdjustLimit] class.
var heightAdjustLimitClass _heightAdjustLimitClass

func init() {
	heightAdjustLimitClass = _heightAdjustLimitClass{objc.GetClass("heightAdjustLimit")}
}

type _heightAdjustLimitClass struct {
	objc.Class
}

// An interface definition for the [heightAdjustLimit] class.
type IheightAdjustLimit interface {
	ID() objc.ID
}

type heightAdjustLimit struct {
	id objc.ID
}

func heightAdjustLimitFrom(ptr unsafe.Pointer) heightAdjustLimit {
	return heightAdjustLimit{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (h_ heightAdjustLimit) ID() objc.ID {
	return h_.id
}

// Alloc allocates a new instance without initialization.
func (hc _heightAdjustLimitClass) Alloc() heightAdjustLimit {
	rv := objc.Send[heightAdjustLimit](objc.ID(hc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (hc _heightAdjustLimitClass) New() heightAdjustLimit {
	rv := objc.Send[heightAdjustLimit](objc.ID(hc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewheightAdjustLimit creates and returns a new initialized instance.
func NewheightAdjustLimit() heightAdjustLimit {
	return heightAdjustLimitClass.New()
}

// Init initializes the instance.
func (h_ heightAdjustLimit) Init() heightAdjustLimit {
	rv := objc.Send[heightAdjustLimit](h_.ID(), selInit)
	return rv
}
