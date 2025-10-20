// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [reserved2] class.
var (
	Reserved2Class     _reserved2Class
	Reserved2ClassOnce sync.Once
)

func getreserved2Class() _reserved2Class {
	Reserved2ClassOnce.Do(func() {
		Reserved2Class = _reserved2Class{objc.GetClass("reserved2")}
	})
	return Reserved2Class
}

type _reserved2Class struct {
	class objc.Class
}

// An interface definition for the [reserved2] class.
type Ireserved2 interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPortMessage/reserved2
type reserved2 struct {
	objectivec.Object
}

// reserved2From constructs a [reserved2] from an unsafe.Pointer.
func reserved2From(ptr unsafe.Pointer) reserved2 {
	return reserved2{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _reserved2Class) Alloc() reserved2 {
	rv := objc.Send[reserved2](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _reserved2Class) New() reserved2 {
	rv := objc.Send[reserved2](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ reserved2) Init() reserved2 {
	rv := objc.Send[reserved2](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ reserved2) Autorelease() reserved2 {
	rv := objc.Send[reserved2](r_.ID, objc.Sel("autorelease"))
	return rv
}

// Newreserved2 creates a new reserved2 instance.
func Newreserved2() reserved2 {
	return getreserved2Class().New()
}




