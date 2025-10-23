// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Map] class.
var (
	MapClass     _MapClass
	MapClassOnce sync.Once
)

func getMapClass() _MapClass {
	MapClassOnce.Do(func() {
		MapClass = _MapClass{objc.GetClass("map")}
	})
	return MapClass
}

type _MapClass struct {
	class objc.Class
}

// An interface definition for the [Map] class.
type IMap interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArchiver/map
type Map struct {
	objectivec.Object
}

// MapFrom constructs a [Map] from an unsafe.Pointer.
func MapFrom(ptr unsafe.Pointer) Map {
	return Map{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MapClass) Alloc() Map {
	rv := objc.Send[Map](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MapClass) New() Map {
	rv := objc.Send[Map](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ Map) Init() Map {
	rv := objc.Send[Map](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ Map) Autorelease() Map {
	rv := objc.Send[Map](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMap creates a new Map instance.
func NewMap() Map {
	return getMapClass().New()
}




