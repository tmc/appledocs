// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [numBytes] class.
var numBytesClass = _numBytesClass{objc.GetClass("numBytes")}

type _numBytesClass struct {
	class objc.Class
}

// An interface definition for the [numBytes] class.
type InumBytes interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSimpleCString/numBytes

type numBytes struct {
	objectivec.Object
}

// numBytesFrom constructs a [numBytes] from an unsafe.Pointer.
func numBytesFrom(ptr unsafe.Pointer) numBytes {
	return numBytes{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (nc _numBytesClass) Alloc() numBytes {
	rv := objc.Send[numBytes](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (nc _numBytesClass) New() numBytes {
	rv := objc.Send[numBytes](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ numBytes) Init() numBytes {
	rv := objc.Send[numBytes](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ numBytes) Autorelease() numBytes {
	rv := objc.Send[numBytes](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewnumBytes creates a new numBytes instance.
func NewnumBytes() numBytes {
	return numBytesClass.New()
}




