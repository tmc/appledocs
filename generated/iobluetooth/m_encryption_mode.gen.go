// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mEncryptionMode] class.
var (
	MEncryptionModeClass     _mEncryptionModeClass
	MEncryptionModeClassOnce sync.Once
)

func getmEncryptionModeClass() _mEncryptionModeClass {
	MEncryptionModeClassOnce.Do(func() {
		MEncryptionModeClass = _mEncryptionModeClass{objc.GetClass("mEncryptionMode")}
	})
	return MEncryptionModeClass
}

type _mEncryptionModeClass struct {
	class objc.Class
}

// An interface definition for the [mEncryptionMode] class.
type ImEncryptionMode interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevice/mEncryptionMode
type mEncryptionMode struct {
	objectivec.Object
}

// mEncryptionModeFrom constructs a [mEncryptionMode] from an unsafe.Pointer.
func mEncryptionModeFrom(ptr unsafe.Pointer) mEncryptionMode {
	return mEncryptionMode{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mEncryptionModeClass) Alloc() mEncryptionMode {
	rv := objc.Send[mEncryptionMode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mEncryptionModeClass) New() mEncryptionMode {
	rv := objc.Send[mEncryptionMode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mEncryptionMode) Init() mEncryptionMode {
	rv := objc.Send[mEncryptionMode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mEncryptionMode) Autorelease() mEncryptionMode {
	rv := objc.Send[mEncryptionMode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmEncryptionMode creates a new mEncryptionMode instance.
func NewmEncryptionMode() mEncryptionMode {
	return getmEncryptionModeClass().New()
}




