
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [contentResizeIncrements] class.
var contentResizeIncrementsClass _contentResizeIncrementsClass

func init() {
	contentResizeIncrementsClass = _contentResizeIncrementsClass{objc.GetClass("contentResizeIncrements")}
}

type _contentResizeIncrementsClass struct {
	objc.Class
}

// An interface definition for the [contentResizeIncrements] class.
type IcontentResizeIncrements interface {
	ID() objc.ID
}

type contentResizeIncrements struct {
	id objc.ID
}

func contentResizeIncrementsFrom(ptr unsafe.Pointer) contentResizeIncrements {
	return contentResizeIncrements{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ contentResizeIncrements) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _contentResizeIncrementsClass) Alloc() contentResizeIncrements {
	rv := objc.Send[contentResizeIncrements](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _contentResizeIncrementsClass) New() contentResizeIncrements {
	rv := objc.Send[contentResizeIncrements](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcontentResizeIncrements creates and returns a new initialized instance.
func NewcontentResizeIncrements() contentResizeIncrements {
	return contentResizeIncrementsClass.New()
}

// Init initializes the instance.
func (c_ contentResizeIncrements) Init() contentResizeIncrements {
	rv := objc.Send[contentResizeIncrements](c_.ID(), selInit)
	return rv
}
