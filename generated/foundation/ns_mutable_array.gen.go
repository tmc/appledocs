// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MutableArray] class.
var (
	MutableArrayClass     _MutableArrayClass
	MutableArrayClassOnce sync.Once
)

func getMutableArrayClass() _MutableArrayClass {
	MutableArrayClassOnce.Do(func() {
		MutableArrayClass = _MutableArrayClass{objc.GetClass("NSMutableArray")}
	})
	return MutableArrayClass
}

type _MutableArrayClass struct {
	class objc.Class
}

// An interface definition for the [MutableArray] class.
type IMutableArray interface {
	IArray
	SortUsingDescriptors(sortDescriptors []SortDescriptor)
}

// A dynamic ordered collection of objects.
//
// You can use this type in Swift instead of an variable in cases that require reference semantics. The class declares the programmatic interface to objects that manage a modifiable array of objects. This class adds insertion and deletion operations to the basic array-handling behavior inherited from . NSMutableArray is “toll-free bridged” with its Core Foundation counterpart, . See for more information.


// A dynamic ordered collection of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray
type MutableArray struct {
	Array
}

// MutableArrayFrom constructs a [MutableArray] from an unsafe.Pointer.
//
// A dynamic ordered collection of objects.
func MutableArrayFrom(ptr unsafe.Pointer) MutableArray {
	return MutableArray{
		Array: ArrayFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MutableArrayClass) Alloc() MutableArray {
	rv := objc.Send[MutableArray](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MutableArrayClass) New() MutableArray {
	rv := objc.Send[MutableArray](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MutableArray) Init() MutableArray {
	rv := objc.Send[MutableArray](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MutableArray) Autorelease() MutableArray {
	rv := objc.Send[MutableArray](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMutableArray creates a new MutableArray instance.
func NewMutableArray() MutableArray {
	return getMutableArrayClass().New()
}



// Sorts the receiver using a given array of sort descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/sort(using:)-4eh07
func (m_ MutableArray) SortUsingDescriptors(sortDescriptors []SortDescriptor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("sortUsingDescriptors:"), sortDescriptors)
}



