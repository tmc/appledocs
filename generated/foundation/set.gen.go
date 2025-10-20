// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Set] class.
var (
	setClass     _SetClass
	setClassOnce sync.Once
)

func getSetClass() _SetClass {
	setClassOnce.Do(func() {
		setClass = _SetClass{objc.GetClass("NSSet")}
	})
	return setClass
}

type _SetClass struct {
	class objc.Class
}

// An interface definition for the [Set] class.
type ISet interface {
	objectivec.IObject
}

// A static, unordered collection of unique objects.
//
// The , , and classes declare the programmatic interface to an unordered collection of objects. declares the programmatic interface for static sets of distinct objects. You establish a static set’s entries when it’s created, and can’t modify the entries after that. , on the other hand, declares a programmatic interface for dynamic sets of distinct objects. A dynamic — or mutable — set allows the addition and deletion of entries at any time, automatically allocating memory as needed. Use sets as an alternative to arrays when the order of elements isn’t important and you need to consider performance in testing whether the set contains an object. With an array, testing for membership is slower than with sets. is “toll-free bridged” with its Core Foundation counterpart, . See for more information on toll-free bridging. In Swift, use this class instead of a constant in cases where you require reference semantics.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet
type Set struct {
	objectivec.Object
}

// SetFrom constructs a [Set] from an unsafe.Pointer.
//
// A static, unordered collection of unique objects.
func SetFrom(ptr unsafe.Pointer) Set {
	return Set{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SetClass) Alloc() Set {
	rv := objc.Send[Set](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SetClass) New() Set {
	rv := objc.Send[Set](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ Set) Init() Set {
	rv := objc.Send[Set](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ Set) Autorelease() Set {
	rv := objc.Send[Set](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSet creates a new Set instance.
func NewSet() Set {
	return getSetClass().New()
}




