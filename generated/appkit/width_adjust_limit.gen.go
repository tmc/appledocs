
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [widthAdjustLimit] class.
var widthAdjustLimitClass _widthAdjustLimitClass

func init() {
	widthAdjustLimitClass = _widthAdjustLimitClass{objc.GetClass("widthAdjustLimit")}
}

type _widthAdjustLimitClass struct {
	objc.Class
}

// An interface definition for the [widthAdjustLimit] class.
type IwidthAdjustLimit interface {
	ID() objc.ID
}

type widthAdjustLimit struct {
	id objc.ID
}

func widthAdjustLimitFrom(ptr unsafe.Pointer) widthAdjustLimit {
	return widthAdjustLimit{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (w_ widthAdjustLimit) ID() objc.ID {
	return w_.id
}

// Alloc allocates a new instance without initialization.
func (wc _widthAdjustLimitClass) Alloc() widthAdjustLimit {
	rv := objc.Send[widthAdjustLimit](objc.ID(wc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (wc _widthAdjustLimitClass) New() widthAdjustLimit {
	rv := objc.Send[widthAdjustLimit](objc.ID(wc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewwidthAdjustLimit creates and returns a new initialized instance.
func NewwidthAdjustLimit() widthAdjustLimit {
	return widthAdjustLimitClass.New()
}

// Init initializes the instance.
func (w_ widthAdjustLimit) Init() widthAdjustLimit {
	rv := objc.Send[widthAdjustLimit](w_.ID(), selInit)
	return rv
}
