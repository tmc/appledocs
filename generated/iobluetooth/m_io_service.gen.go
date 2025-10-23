// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mIOService] class.
var (
	MIOServiceClass     _mIOServiceClass
	MIOServiceClassOnce sync.Once
)

func getmIOServiceClass() _mIOServiceClass {
	MIOServiceClassOnce.Do(func() {
		MIOServiceClass = _mIOServiceClass{objc.GetClass("mIOService")}
	})
	return MIOServiceClass
}

type _mIOServiceClass struct {
	class objc.Class
}

// An interface definition for the [mIOService] class.
type ImIOService interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothObject/mIOService
type mIOService struct {
	objectivec.Object
}

// mIOServiceFrom constructs a [mIOService] from an unsafe.Pointer.
func mIOServiceFrom(ptr unsafe.Pointer) mIOService {
	return mIOService{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mIOServiceClass) Alloc() mIOService {
	rv := objc.Send[mIOService](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mIOServiceClass) New() mIOService {
	rv := objc.Send[mIOService](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mIOService) Init() mIOService {
	rv := objc.Send[mIOService](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mIOService) Autorelease() mIOService {
	rv := objc.Send[mIOService](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmIOService creates a new mIOService instance.
func NewmIOService() mIOService {
	return getmIOServiceClass().New()
}




