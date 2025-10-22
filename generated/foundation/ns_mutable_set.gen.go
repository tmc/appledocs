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
	AddObject(object unsafe.Pointer)
	AddObjectsFromArray(array []objc.ID)
	FilterUsingPredicate(predicate IPredicate)
	IntersectSet(otherSet unsafe.Pointer)
	MinusSet(otherSet unsafe.Pointer)
	RemoveObject(object unsafe.Pointer)
	RemoveAllObjects()
	SetSet(otherSet unsafe.Pointer)
	UnionSet(otherSet unsafe.Pointer)
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





// Returns an initialized mutable set with a given initial capacity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableSet/init(capacity:)

func NewMutableSetWithCapacity(numItems uint) MutableSet {
	instance := getMutableSetClass().Alloc()
	rv := objc.Send[MutableSet](instance.ID, objc.Sel("initWithCapacity:"), numItems)
	rv.Autorelease()
	return rv
}




//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableSet/init(coder:)

func NewMutableSetWithCoder(coder ICoder) MutableSet {
	instance := getMutableSetClass().Alloc()
	rv := objc.Send[MutableSet](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}



// Creates and returns a mutable set with a given initial capacity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableSet/setWithCapacity:

func (mc _MutableSetClass) SetWithCapacity(numItems uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("setWithCapacity:"), numItems)
	return rv
}


// Adds a given object to the set, if it is not already a member.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableSet/add(_:)

func (m_ MutableSet) AddObject(object unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addObject:"), object)
}


// Adds to the set each object contained in a given array that is not already a member.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableSet/addObjects(from:)

func (m_ MutableSet) AddObjectsFromArray(array []objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addObjectsFromArray:"), array)
}


// Evaluates a given predicate against the set’s content and removes from the set those objects for which the predicate returns false.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableSet/filter(using:)

func (m_ MutableSet) FilterUsingPredicate(predicate IPredicate) {
	objc.Send[objc.ID](m_.ID, objc.Sel("filterUsingPredicate:"), predicate)
}


// Removes from the receiving set each object that isn’t a member of another given set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableSet/intersect(_:)

func (m_ MutableSet) IntersectSet(otherSet unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("intersectSet:"), otherSet)
}


// Removes each object in another given set from the receiving set, if present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableSet/minus(_:)

func (m_ MutableSet) MinusSet(otherSet unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("minusSet:"), otherSet)
}


// Removes a given object from the set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableSet/remove(_:)

func (m_ MutableSet) RemoveObject(object unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeObject:"), object)
}


// Empties the set of all of its members.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableSet/removeAllObjects()

func (m_ MutableSet) RemoveAllObjects() {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeAllObjects"))
}


// Empties the receiving set, then adds each object contained in another given set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableSet/setSet(_:)

func (m_ MutableSet) SetSet(otherSet unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSet:"), otherSet)
}


// Adds each object in another given set to the receiving set, if not present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableSet/union(_:)

func (m_ MutableSet) UnionSet(otherSet unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("unionSet:"), otherSet)
}


