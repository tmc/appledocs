
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [accessBehavior] class.
var accessBehaviorClass _accessBehaviorClass

func init() {
	accessBehaviorClass = _accessBehaviorClass{objc.GetClass("accessBehavior")}
}

type _accessBehaviorClass struct {
	objc.Class
}

// An interface definition for the [accessBehavior] class.
type IaccessBehavior interface {
	ID() objc.ID
}

type accessBehavior struct {
	id objc.ID
}

func accessBehaviorFrom(ptr unsafe.Pointer) accessBehavior {
	return accessBehavior{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ accessBehavior) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _accessBehaviorClass) Alloc() accessBehavior {
	rv := objc.Send[accessBehavior](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _accessBehaviorClass) New() accessBehavior {
	rv := objc.Send[accessBehavior](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewaccessBehavior creates and returns a new initialized instance.
func NewaccessBehavior() accessBehavior {
	return accessBehaviorClass.New()
}

// Init initializes the instance.
func (a_ accessBehavior) Init() accessBehavior {
	rv := objc.Send[accessBehavior](a_.ID(), selInit)
	return rv
}
