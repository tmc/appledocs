
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [removeAllToolTips] class.
var removeAllToolTipsClass _removeAllToolTipsClass

func init() {
	removeAllToolTipsClass = _removeAllToolTipsClass{objc.GetClass("removeAllToolTips")}
}

type _removeAllToolTipsClass struct {
	objc.Class
}

// An interface definition for the [removeAllToolTips] class.
type IremoveAllToolTips interface {
	ID() objc.ID
}

type removeAllToolTips struct {
	id objc.ID
}

func removeAllToolTipsFrom(ptr unsafe.Pointer) removeAllToolTips {
	return removeAllToolTips{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ removeAllToolTips) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _removeAllToolTipsClass) Alloc() removeAllToolTips {
	rv := objc.Send[removeAllToolTips](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _removeAllToolTipsClass) New() removeAllToolTips {
	rv := objc.Send[removeAllToolTips](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewremoveAllToolTips creates and returns a new initialized instance.
func NewremoveAllToolTips() removeAllToolTips {
	return removeAllToolTipsClass.New()
}

// Init initializes the instance.
func (r_ removeAllToolTips) Init() removeAllToolTips {
	rv := objc.Send[removeAllToolTips](r_.ID(), selInit)
	return rv
}
