// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [map] class.
var (
	MapClass     _mapClass
	MapClassOnce sync.Once
)

func getmapClass() _mapClass {
	MapClassOnce.Do(func() {
		MapClass = _mapClass{objc.GetClass("map")}
	})
	return MapClass
}

type _mapClass struct {
	class objc.Class
}

// An interface definition for the [map] class.
type Imap interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArchiver/map
type map struct {
	objectivec.Object
}

// mapFrom constructs a [map] from an unsafe.Pointer.
func mapFrom(ptr unsafe.Pointer) map {
	return map{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mapClass) Alloc() map {
	rv := objc.Send[map](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mapClass) New() map {
	rv := objc.Send[map](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ map) Init() map {
	rv := objc.Send[map](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ map) Autorelease() map {
	rv := objc.Send[map](m_.ID, objc.Sel("autorelease"))
	return rv
}

// Newmap creates a new map instance.
func Newmap() map {
	return getmapClass().New()
}




