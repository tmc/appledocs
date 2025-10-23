// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mPrivateOBEXSessionData] class.
var (
	MPrivateOBEXSessionDataClass     _mPrivateOBEXSessionDataClass
	MPrivateOBEXSessionDataClassOnce sync.Once
)

func getmPrivateOBEXSessionDataClass() _mPrivateOBEXSessionDataClass {
	MPrivateOBEXSessionDataClassOnce.Do(func() {
		MPrivateOBEXSessionDataClass = _mPrivateOBEXSessionDataClass{objc.GetClass("mPrivateOBEXSessionData")}
	})
	return MPrivateOBEXSessionDataClass
}

type _mPrivateOBEXSessionDataClass struct {
	class objc.Class
}

// An interface definition for the [mPrivateOBEXSessionData] class.
type ImPrivateOBEXSessionData interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/mPrivateOBEXSessionData
type mPrivateOBEXSessionData struct {
	objectivec.Object
}

// mPrivateOBEXSessionDataFrom constructs a [mPrivateOBEXSessionData] from an unsafe.Pointer.
func mPrivateOBEXSessionDataFrom(ptr unsafe.Pointer) mPrivateOBEXSessionData {
	return mPrivateOBEXSessionData{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mPrivateOBEXSessionDataClass) Alloc() mPrivateOBEXSessionData {
	rv := objc.Send[mPrivateOBEXSessionData](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mPrivateOBEXSessionDataClass) New() mPrivateOBEXSessionData {
	rv := objc.Send[mPrivateOBEXSessionData](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mPrivateOBEXSessionData) Init() mPrivateOBEXSessionData {
	rv := objc.Send[mPrivateOBEXSessionData](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mPrivateOBEXSessionData) Autorelease() mPrivateOBEXSessionData {
	rv := objc.Send[mPrivateOBEXSessionData](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmPrivateOBEXSessionData creates a new mPrivateOBEXSessionData instance.
func NewmPrivateOBEXSessionData() mPrivateOBEXSessionData {
	return getmPrivateOBEXSessionDataClass().New()
}




