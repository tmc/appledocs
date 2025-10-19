// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NEHotspotHelper] class.
var (
	nEHotspotHelperClass     _NEHotspotHelperClass
	nEHotspotHelperClassOnce sync.Once
)

func getNEHotspotHelperClass() _NEHotspotHelperClass {
	nEHotspotHelperClassOnce.Do(func() {
		nEHotspotHelperClass = _NEHotspotHelperClass{objc.GetClass("NEHotspotHelper")}
	})
	return nEHotspotHelperClass
}

type _NEHotspotHelperClass struct {
	class objc.Class
}

// An interface definition for the [NEHotspotHelper] class.
type INEHotspotHelper interface {
	objectivec.IObject
}

// A class to register a hotspot helper. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHelper
type NEHotspotHelper struct {
	objectivec.Object
}

// NEHotspotHelperFrom constructs a [NEHotspotHelper] from an unsafe.Pointer.
//
// A class to register a hotspot helper.
func NEHotspotHelperFrom(ptr unsafe.Pointer) NEHotspotHelper {
	return NEHotspotHelper{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NEHotspotHelperClass) Alloc() NEHotspotHelper {
	rv := objc.Send[NEHotspotHelper](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEHotspotHelperClass) New() NEHotspotHelper {
	rv := objc.Send[NEHotspotHelper](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEHotspotHelper) Init() NEHotspotHelper {
	rv := objc.Send[NEHotspotHelper](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEHotspotHelper) Autorelease() NEHotspotHelper {
	rv := objc.Send[NEHotspotHelper](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEHotspotHelper creates a new NEHotspotHelper instance.
func NewNEHotspotHelper() NEHotspotHelper {
	return getNEHotspotHelperClass().New()
}




