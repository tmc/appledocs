// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MutableSet] class.
var (
	MutableSetClass     _MutableSetClass
	MutableSetClassOnce sync.Once
)

func getMutableSetClass() _MutableSetClass {
	MutableSetClassOnce.Do(func() {
		MutableSetClass = _MutableSetClass{objc.GetClass("NSMutableSet")}
	})
	return MutableSetClass
}

type _MutableSetClass struct {
	class objc.Class
}

// An interface definition for the [MutableSet] class.
type IMutableSet interface {
	ISet
	// properties:
	// methods:
}

// A dynamic unordered collection of unique objects.
//
// You can use this type in Swift instead of a in cases that require reference semantics. The class declares the programmatic interface to a mutable, unordered collection of distinct objects. The class, which is a concrete subclass of , supports mutable sets that can contain multiple instances of the same element. The class supports creating and managing immutable sets. NSMutableSet is “toll-free bridged” with its Core Foundation counterpart, . See for more information.


// A dynamic unordered collection of unique objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableSet
type MutableSet struct {
	Set
}

// MutableSetFrom constructs a [MutableSet] from an unsafe.Pointer.
//
// A dynamic unordered collection of unique objects.
func MutableSetFrom(ptr unsafe.Pointer) MutableSet {
	return MutableSet{
		Set: SetFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MutableSetClass) Alloc() MutableSet {
	rv := objc.Send[MutableSet](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MutableSetClass) New() MutableSet {
	rv := objc.Send[MutableSet](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MutableSet) Init() MutableSet {
	rv := objc.Send[MutableSet](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MutableSet) Autorelease() MutableSet {
	rv := objc.Send[MutableSet](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMutableSet creates a new MutableSet instance.
func NewMutableSet() MutableSet {
	return getMutableSetClass().New()
}




