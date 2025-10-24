// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PurgeableData] class.
var (
	PurgeableDataClass     _PurgeableDataClass
	PurgeableDataClassOnce sync.Once
)

func getPurgeableDataClass() _PurgeableDataClass {
	PurgeableDataClassOnce.Do(func() {
		PurgeableDataClass = _PurgeableDataClass{objc.GetClass("NSPurgeableData")}
	})
	return PurgeableDataClass
}

type _PurgeableDataClass struct {
	class objc.Class
}

// An interface definition for the [PurgeableData] class.
type IPurgeableData interface {
	IMutableData
	// properties:
	// methods:
}

// A mutable data object containing bytes that can be discarded when they’re no longer needed.
//
// objects inherit their creation methods from their superclass, while providing a default implementation of the protocol. All objects begin “accessed” to ensure that they are not instantly discarded. The method marks the object’s bytes as “accessed,” thus protecting them from being discarded, and must be called before accessing the object, or else an exception will be raised. This method returns if the bytes have not been discarded and if they have been successfully marked as “accessed”. Any method that directly or indirectly accesses these bytes or their length when they are not “accessed” will raise an exception. When you are done with the data, call to allow them to be discarded in order to quickly free up memory. You may use these objects by themselves, and do not necessarily have to use them in conjunction with to get the purging behavior. The class incorporates a caching mechanism with some auto-removal policies to ensure that its memory footprint does not get too large. objects should not be used as keys in hashing-based collections, because the value of the bytes pointer can change after every mutation of the data.


// A mutable data object containing bytes that can be discarded when they’re no longer needed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPurgeableData
type PurgeableData struct {
	MutableData
}

// PurgeableDataFrom constructs a [PurgeableData] from an unsafe.Pointer.
//
// A mutable data object containing bytes that can be discarded when they’re no longer needed.
func PurgeableDataFrom(ptr unsafe.Pointer) PurgeableData {
	return PurgeableData{
		MutableData: MutableDataFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PurgeableDataClass) Alloc() PurgeableData {
	rv := objc.Send[PurgeableData](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PurgeableDataClass) New() PurgeableData {
	rv := objc.Send[PurgeableData](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PurgeableData) Init() PurgeableData {
	rv := objc.Send[PurgeableData](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PurgeableData) Autorelease() PurgeableData {
	rv := objc.Send[PurgeableData](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPurgeableData creates a new PurgeableData instance.
func NewPurgeableData() PurgeableData {
	return getPurgeableDataClass().New()
}




