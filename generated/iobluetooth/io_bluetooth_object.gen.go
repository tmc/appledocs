// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [BluetoothObject] class.
var (
	BluetoothObjectClass     _BluetoothObjectClass
	BluetoothObjectClassOnce sync.Once
)

func getBluetoothObjectClass() _BluetoothObjectClass {
	BluetoothObjectClassOnce.Do(func() {
		BluetoothObjectClass = _BluetoothObjectClass{objc.GetClass("IOBluetoothObject")}
	})
	return BluetoothObjectClass
}

type _BluetoothObjectClass struct {
	class objc.Class
}

// An interface definition for the [BluetoothObject] class.
type IBluetoothObject interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothObject
type BluetoothObject struct {
	objectivec.Object
}

// BluetoothObjectFrom constructs a [BluetoothObject] from an unsafe.Pointer.
func BluetoothObjectFrom(ptr unsafe.Pointer) BluetoothObject {
	return BluetoothObject{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (bc _BluetoothObjectClass) Alloc() BluetoothObject {
	rv := objc.Send[BluetoothObject](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BluetoothObjectClass) New() BluetoothObject {
	rv := objc.Send[BluetoothObject](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BluetoothObject) Init() BluetoothObject {
	rv := objc.Send[BluetoothObject](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BluetoothObject) Autorelease() BluetoothObject {
	rv := objc.Send[BluetoothObject](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBluetoothObject creates a new BluetoothObject instance.
func NewBluetoothObject() BluetoothObject {
	return getBluetoothObjectClass().New()
}




