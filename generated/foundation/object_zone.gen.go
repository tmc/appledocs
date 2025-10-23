// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [objectZone] class.
var (
	ObjectZoneClass     _objectZoneClass
	ObjectZoneClassOnce sync.Once
)

func getobjectZoneClass() _objectZoneClass {
	ObjectZoneClassOnce.Do(func() {
		ObjectZoneClass = _objectZoneClass{objc.GetClass("objectZone")}
	})
	return ObjectZoneClass
}

type _objectZoneClass struct {
	class objc.Class
}

// An interface definition for the [objectZone] class.
type IobjectZone interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnarchiver/objectZone-c.ivar
type objectZone struct {
	objectivec.Object
}

// objectZoneFrom constructs a [objectZone] from an unsafe.Pointer.
func objectZoneFrom(ptr unsafe.Pointer) objectZone {
	return objectZone{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (oc _objectZoneClass) Alloc() objectZone {
	rv := objc.Send[objectZone](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _objectZoneClass) New() objectZone {
	rv := objc.Send[objectZone](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ objectZone) Init() objectZone {
	rv := objc.Send[objectZone](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ objectZone) Autorelease() objectZone {
	rv := objc.Send[objectZone](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewobjectZone creates a new objectZone instance.
func NewobjectZone() objectZone {
	return getobjectZoneClass().New()
}




