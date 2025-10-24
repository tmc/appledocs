// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSCountedSet */


/* debug [class_header]: Header for NSCountedSet */
// The class instance for the [CountedSet] class.
var (
	CountedSetClass     _CountedSetClass
	CountedSetClassOnce sync.Once
)

func getCountedSetClass() _CountedSetClass {
	CountedSetClassOnce.Do(func() {
		CountedSetClass = _CountedSetClass{objc.GetClass("NSCountedSet")}
	})
	return CountedSetClass
}

type _CountedSetClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CountedSet */
// An interface definition for the [CountedSet] class.
type ICountedSet interface {
	IMutableSet
	
/* debug [class_interface_properties]: Properties for CountedSet */
	// properties:
	Count() int
	SetCount(value int)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CountedSet */
	// methods:
	AddObject(object objectivec.IObject)
	CountForObject(object objectivec.IObject) uint
	ObjectEnumerator() unsafe.Pointer
	RemoveObject(object objectivec.IObject)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CountedSet */
// Alloc allocates a new instance without initialization.
func (cc _CountedSetClass) Alloc() CountedSet {
	rv := objc.Send[CountedSet](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CountedSetClass) New() CountedSet {
	rv := objc.Send[CountedSet](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CountedSet) Init() CountedSet {
	rv := objc.Send[CountedSet](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CountedSet) Autorelease() CountedSet {
	rv := objc.Send[CountedSet](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCountedSet creates a new CountedSet instance.
func NewCountedSet() CountedSet {
	return getCountedSetClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CountedSet */
// A mutable, unordered collection of distinct objects that may appear more than once in the collection.
//
// Each distinct object inserted into an object has a counter associated with it. keeps track of the number of times objects are inserted and requires that objects be removed the same number of times. Thus, there is only one instance of an object in an object even if the object has been added to the set multiple times. The method defined by the superclass has special significance; it returns the number of distinct objects, not the total number of times objects are represented in the set. The and classes are provided for static and dynamic sets, respectively, whose elements are distinct. While and are not toll-free bridged, they provide similar functionality. For more information about , see the .


// A mutable, unordered collection of distinct objects that may appear more than once in the collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCountedSet
type CountedSet struct {
	MutableSet
}

// CountedSetFrom constructs a [CountedSet] from an unsafe.Pointer.
//
// A mutable, unordered collection of distinct objects that may appear more than once in the collection.
func CountedSetFrom(ptr unsafe.Pointer) CountedSet {
	return CountedSet{
		MutableSet: MutableSetFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CountedSet */

// Returns a counted set object initialized with the contents of a given array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCountedSet/init(array:)
func NewCountedSetWithArray(array []objc.ID) CountedSet {
	instance := getCountedSetClass().Alloc()
	rv := objc.Send[CountedSet](instance.ID, objc.Sel("initWithArray:"), array)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCountedSetWithArray */


// Returns a counted set object initialized with enough memory to hold a given number of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCountedSet/init(capacity:)
func NewCountedSetWithCapacity(numItems uint) CountedSet {
	instance := getCountedSetClass().Alloc()
	rv := objc.Send[CountedSet](instance.ID, objc.Sel("initWithCapacity:"), numItems)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCountedSetWithCapacity */


// Returns a counted set object initialized with the contents of a given set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCountedSet/init(set:)
func NewCountedSetWithSet(set unsafe.Pointer) CountedSet {
	instance := getCountedSetClass().Alloc()
	rv := objc.Send[CountedSet](instance.ID, objc.Sel("initWithSet:"), set)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCountedSetWithSet */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CountedSet */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CountedSet */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CountedSet */

// Adds a given object to the set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCountedSet/add(_:)
func (c_ CountedSet) AddObject(object objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addObject:"), object)
}/* debug [instance_methods/method]: AddObject */


// Returns the count associated with a given object in the set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCountedSet/count(for:)
func (c_ CountedSet) CountForObject(object objectivec.IObject) uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("countForObject:"), object)
	return rv
}/* debug [instance_methods/method]: CountForObject */


// Returns an enumerator object that lets you access each object in the set once, independent of its count.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCountedSet/objectEnumerator()
func (c_ CountedSet) ObjectEnumerator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("objectEnumerator"))
	return rv
}/* debug [instance_methods/method]: ObjectEnumerator */


// Removes a given object from the set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCountedSet/remove(_:)
func (c_ CountedSet) RemoveObject(object objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeObject:"), object)
}/* debug [instance_methods/method]: RemoveObject */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CountedSet */

// The number of members in the set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/count
func (c_ CountedSet) Count() int {
	rv := objc.Send[int](c_.ID, objc.Sel("count"))
	return rv
}/* debug [instance_properties/getter]: count */


// The number of members in the set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/count
func (c_ CountedSet) SetCount(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCount:"), value)
}/* debug [instance_properties/setter]: count */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSCountedSet */


