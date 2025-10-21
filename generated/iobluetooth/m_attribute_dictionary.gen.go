// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mAttributeDictionary] class.
var (
	MAttributeDictionaryClass     _mAttributeDictionaryClass
	MAttributeDictionaryClassOnce sync.Once
)

func getmAttributeDictionaryClass() _mAttributeDictionaryClass {
	MAttributeDictionaryClassOnce.Do(func() {
		MAttributeDictionaryClass = _mAttributeDictionaryClass{objc.GetClass("mAttributeDictionary")}
	})
	return MAttributeDictionaryClass
}

type _mAttributeDictionaryClass struct {
	class objc.Class
}

// An interface definition for the [mAttributeDictionary] class.
type ImAttributeDictionary interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/mAttributeDictionary
type mAttributeDictionary struct {
	objectivec.Object
}

// mAttributeDictionaryFrom constructs a [mAttributeDictionary] from an unsafe.Pointer.
func mAttributeDictionaryFrom(ptr unsafe.Pointer) mAttributeDictionary {
	return mAttributeDictionary{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mAttributeDictionaryClass) Alloc() mAttributeDictionary {
	rv := objc.Send[mAttributeDictionary](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mAttributeDictionaryClass) New() mAttributeDictionary {
	rv := objc.Send[mAttributeDictionary](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mAttributeDictionary) Init() mAttributeDictionary {
	rv := objc.Send[mAttributeDictionary](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mAttributeDictionary) Autorelease() mAttributeDictionary {
	rv := objc.Send[mAttributeDictionary](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmAttributeDictionary creates a new mAttributeDictionary instance.
func NewmAttributeDictionary() mAttributeDictionary {
	return getmAttributeDictionaryClass().New()
}




