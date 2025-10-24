// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [classVersions] class.
var (
	ClassVersionsClass     _classVersionsClass
	ClassVersionsClassOnce sync.Once
)

func getclassVersionsClass() _classVersionsClass {
	ClassVersionsClassOnce.Do(func() {
		ClassVersionsClass = _classVersionsClass{objc.GetClass("classVersions")}
	})
	return ClassVersionsClass
}

type _classVersionsClass struct {
	class objc.Class
}

// An interface definition for the [classVersions] class.
type IclassVersions interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnarchiver/classVersions
type classVersions struct {
	objectivec.Object
}

// classVersionsFrom constructs a [classVersions] from an unsafe.Pointer.
func classVersionsFrom(ptr unsafe.Pointer) classVersions {
	return classVersions{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _classVersionsClass) Alloc() classVersions {
	rv := objc.Send[classVersions](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _classVersionsClass) New() classVersions {
	rv := objc.Send[classVersions](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ classVersions) Init() classVersions {
	rv := objc.Send[classVersions](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ classVersions) Autorelease() classVersions {
	rv := objc.Send[classVersions](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewclassVersions creates a new classVersions instance.
func NewclassVersions() classVersions {
	return getclassVersionsClass().New()
}




