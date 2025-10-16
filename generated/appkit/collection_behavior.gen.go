
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [collectionBehavior] class.
var collectionBehaviorClass _collectionBehaviorClass

func init() {
	collectionBehaviorClass = _collectionBehaviorClass{objc.GetClass("collectionBehavior")}
}

type _collectionBehaviorClass struct {
	objc.Class
}

// An interface definition for the [collectionBehavior] class.
type IcollectionBehavior interface {
	ID() objc.ID
}

type collectionBehavior struct {
	id objc.ID
}

func collectionBehaviorFrom(ptr unsafe.Pointer) collectionBehavior {
	return collectionBehavior{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ collectionBehavior) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _collectionBehaviorClass) Alloc() collectionBehavior {
	rv := objc.Send[collectionBehavior](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _collectionBehaviorClass) New() collectionBehavior {
	rv := objc.Send[collectionBehavior](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcollectionBehavior creates and returns a new initialized instance.
func NewcollectionBehavior() collectionBehavior {
	return collectionBehaviorClass.New()
}

// Init initializes the instance.
func (c_ collectionBehavior) Init() collectionBehavior {
	rv := objc.Send[collectionBehavior](c_.ID(), selInit)
	return rv
}
