// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSMutableSet */


/* debug [class_header]: Header for NSMutableSet */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MutableSet */
// An interface definition for the [MutableSet] class.
type IMutableSet interface {
	ISet
	
/* debug [class_interface_properties]: Properties for MutableSet */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MutableSet */
	// methods:
	AddObject(object objectivec.IObject)
	AddObjectsFromArray(array []objc.ID)
	FilterUsingPredicate(predicate IPredicate)
	IntersectSet(otherSet unsafe.Pointer)
	MinusSet(otherSet unsafe.Pointer)
	RemoveObject(object objectivec.IObject)
	RemoveAllObjects()
	SetSet(otherSet unsafe.Pointer)
	UnionSet(otherSet unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MutableSet */
// Alloc allocates a new instance without initialization.
func (mc _MutableSetClass) Alloc() MutableSet {
	rv := objc.Send[MutableSet](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MutableSet */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MutableSet */

// Returns an initialized mutable set with a given initial capacity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableSet/init(capacity:)
func NewMutableSetWithCapacity(numItems uint) MutableSet {
	instance := getMutableSetClass().Alloc()
	rv := objc.Send[MutableSet](instance.ID, objc.Sel("initWithCapacity:"), numItems)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMutableSetWithCapacity */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableSet/init(coder:)
func NewMutableSetWithCoder(coder ICoder) MutableSet {
	instance := getMutableSetClass().Alloc()
	rv := objc.Send[MutableSet](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMutableSetWithCoder */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MutableSet */

// Creates and returns a mutable set with a given initial capacity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableSet/setWithCapacity:
func (mc _MutableSetClass) SetWithCapacity(numItems uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("setWithCapacity:"), numItems)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SetWithCapacity) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MutableSet */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MutableSet */

// Adds a given object to the set, if it is not already a member.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableSet/add(_:)
func (m_ MutableSet) AddObject(object objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addObject:"), object)
}/* debug [instance_methods/method]: AddObject */


// Adds to the set each object contained in a given array that is not already a member.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableSet/addObjects(from:)
func (m_ MutableSet) AddObjectsFromArray(array []objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addObjectsFromArray:"), array)
}/* debug [instance_methods/method]: AddObjectsFromArray */


// Evaluates a given predicate against the set’s content and removes from the set those objects for which the predicate returns false.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableSet/filter(using:)
func (m_ MutableSet) FilterUsingPredicate(predicate IPredicate) {
	objc.Send[objc.ID](m_.ID, objc.Sel("filterUsingPredicate:"), predicate)
}/* debug [instance_methods/method]: FilterUsingPredicate */


// Removes from the receiving set each object that isn’t a member of another given set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableSet/intersect(_:)
func (m_ MutableSet) IntersectSet(otherSet unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("intersectSet:"), otherSet)
}/* debug [instance_methods/method]: IntersectSet */


// Removes each object in another given set from the receiving set, if present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableSet/minus(_:)
func (m_ MutableSet) MinusSet(otherSet unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("minusSet:"), otherSet)
}/* debug [instance_methods/method]: MinusSet */


// Removes a given object from the set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableSet/remove(_:)
func (m_ MutableSet) RemoveObject(object objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeObject:"), object)
}/* debug [instance_methods/method]: RemoveObject */


// Empties the set of all of its members.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableSet/removeAllObjects()
func (m_ MutableSet) RemoveAllObjects() {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeAllObjects"))
}/* debug [instance_methods/method]: RemoveAllObjects */


// Empties the receiving set, then adds each object contained in another given set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableSet/setSet(_:)
func (m_ MutableSet) SetSet(otherSet unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSet:"), otherSet)
}/* debug [instance_methods/method]: SetSet */


// Adds each object in another given set to the receiving set, if not present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableSet/union(_:)
func (m_ MutableSet) UnionSet(otherSet unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("unionSet:"), otherSet)
}/* debug [instance_methods/method]: UnionSet */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MutableSet */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSMutableSet */


