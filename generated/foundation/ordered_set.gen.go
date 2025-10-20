// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [OrderedSet] class.
var (
	orderedSetClass     _OrderedSetClass
	orderedSetClassOnce sync.Once
)

func getOrderedSetClass() _OrderedSetClass {
	orderedSetClassOnce.Do(func() {
		orderedSetClass = _OrderedSetClass{objc.GetClass("NSOrderedSet")}
	})
	return orderedSetClass
}

type _OrderedSetClass struct {
	class objc.Class
}

// An interface definition for the [OrderedSet] class.
type IOrderedSet interface {
	objectivec.IObject
}

// A static, ordered collection of unique objects.
//
// declares the programmatic interface for static sets of distinct objects. You establish a static set’s entries when it’s created, and thereafter the entries can’t be modified. , on the other hand, declares a programmatic interface for dynamic sets of distinct objects. A dynamic—or mutable—set allows the addition and deletion of entries at any time, automatically allocating memory as needed. You can use ordered sets as an alternative to arrays when the order of elements is important and performance in testing whether an object is contained in the set is a consideration—testing for membership of an array is slower than testing for membership of a set.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet
type OrderedSet struct {
	objectivec.Object
}

// OrderedSetFrom constructs a [OrderedSet] from an unsafe.Pointer.
//
// A static, ordered collection of unique objects.
func OrderedSetFrom(ptr unsafe.Pointer) OrderedSet {
	return OrderedSet{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (oc _OrderedSetClass) Alloc() OrderedSet {
	rv := objc.Send[OrderedSet](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _OrderedSetClass) New() OrderedSet {
	rv := objc.Send[OrderedSet](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OrderedSet) Init() OrderedSet {
	rv := objc.Send[OrderedSet](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OrderedSet) Autorelease() OrderedSet {
	rv := objc.Send[OrderedSet](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOrderedSet creates a new OrderedSet instance.
func NewOrderedSet() OrderedSet {
	return getOrderedSetClass().New()
}




