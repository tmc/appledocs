// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSSortDescriptor */


/* debug [class_header]: Header for NSSortDescriptor */
// The class instance for the [SortDescriptor] class.
var (
	SortDescriptorClass     _SortDescriptorClass
	SortDescriptorClassOnce sync.Once
)

func getSortDescriptorClass() _SortDescriptorClass {
	SortDescriptorClassOnce.Do(func() {
		SortDescriptorClass = _SortDescriptorClass{objc.GetClass("NSSortDescriptor")}
	})
	return SortDescriptorClass
}

type _SortDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SortDescriptor */
// An interface definition for the [SortDescriptor] class.
type ISortDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SortDescriptor */
	// properties:
	Ascending() bool
	Comparator() Comparator /* not a class type */
	Key() IString
	ReversedSortDescriptor() objc.ID
	Selector() objc.SEL
	SortDescriptors() ISortDescriptor
	SetSortDescriptors(value ISortDescriptor)
	KeyPath() objectivec.IObject
	SetKeyPath(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SortDescriptor */
	// methods:
	AllowEvaluation()
	CompareObjectToObject(object1 objc.IObject, object2 objc.IObject) ComparisonResult
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SortDescriptor */
// Alloc allocates a new instance without initialization.
func (sc _SortDescriptorClass) Alloc() SortDescriptor {
	rv := objc.Send[SortDescriptor](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SortDescriptorClass) New() SortDescriptor {
	rv := objc.Send[SortDescriptor](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SortDescriptor) Init() SortDescriptor {
	rv := objc.Send[SortDescriptor](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SortDescriptor) Autorelease() SortDescriptor {
	rv := objc.Send[SortDescriptor](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSortDescriptor creates a new SortDescriptor instance.
func NewSortDescriptor() SortDescriptor {
	return getSortDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SortDescriptor */
// An immutable description of how to order a collection of objects according to a property common to all the objects.
//
// You construct instances of by specifying the key path of the property to compare and the order of the sort (ascending or descending). Optionally, you can also specify a selector to use to perform the comparison, which allows you to specify other comparison selectors, such as and . Sorting raises an exception if the objects don’t respond to the sort descriptor’s comparison selector. You can use sort descriptors for the following: Sorting an array (an instance of or — see and ) Comparing two objects directly (see ) Specifying the order of objects that return from a Core Data fetch request (see )


// An immutable description of how to order a collection of objects according to a property common to all the objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSortDescriptor
type SortDescriptor struct {
	objectivec.Object
}

// SortDescriptorFrom constructs a [SortDescriptor] from an unsafe.Pointer.
//
// An immutable description of how to order a collection of objects according to a property common to all the objects.
func SortDescriptorFrom(ptr unsafe.Pointer) SortDescriptor {
	return SortDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SortDescriptor */

// Creates a sort descriptor by decoding from the coder you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSortDescriptor/init(coder:)
func NewSortDescriptorWithCoder(coder ICoder) SortDescriptor {
	instance := getSortDescriptorClass().Alloc()
	rv := objc.Send[SortDescriptor](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSortDescriptorWithCoder */


// Creates a sort descriptor with a specified string key path and sort order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSortDescriptor/init(key:ascending:)
func NewSortDescriptorWithKeyAscending(key IString, ascending bool) SortDescriptor {
	instance := getSortDescriptorClass().Alloc()
	rv := objc.Send[SortDescriptor](instance.ID, objc.Sel("initWithKey:ascending:"), key, ascending)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSortDescriptorWithKeyAscending */


// Creates a sort descriptor with a specified string key path and ordering, and a comparator block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSortDescriptor/init(key:ascending:comparator:)
func NewSortDescriptorWithKeyAscendingComparator(key IString, ascending bool, cmptr Comparator /* not a class type */) SortDescriptor {
	instance := getSortDescriptorClass().Alloc()
	rv := objc.Send[SortDescriptor](instance.ID, objc.Sel("initWithKey:ascending:comparator:"), key, ascending, cmptr)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSortDescriptorWithKeyAscendingComparator */


// Creates a sort descriptor with a specified string key path, ordering, and comparison selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSortDescriptor/init(key:ascending:selector:)
func NewSortDescriptorWithKeyAscendingSelector(key IString, ascending bool, selector objc.SEL) SortDescriptor {
	instance := getSortDescriptorClass().Alloc()
	rv := objc.Send[SortDescriptor](instance.ID, objc.Sel("initWithKey:ascending:selector:"), key, ascending, selector)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSortDescriptorWithKeyAscendingSelector */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SortDescriptor */

// Creates and returns a sort descriptor with the specified key path and ordering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSortDescriptor/sortDescriptorWithKey:ascending:
func (sc _SortDescriptorClass) SortDescriptorWithKeyAscending(key IString, ascending bool) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("sortDescriptorWithKey:ascending:"), key, ascending)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SortDescriptorWithKeyAscending) */


// Creates and returns a sort descriptor initialized with the specified key path and ordering, and a comparator block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSortDescriptor/sortDescriptorWithKey:ascending:comparator:
func (sc _SortDescriptorClass) SortDescriptorWithKeyAscendingComparator(key IString, ascending bool, cmptr Comparator /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("sortDescriptorWithKey:ascending:comparator:"), key, ascending, cmptr)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SortDescriptorWithKeyAscendingComparator) */


// Creates a sort descriptor with the specified key path, ordering, and comparison selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSortDescriptor/sortDescriptorWithKey:ascending:selector:
func (sc _SortDescriptorClass) SortDescriptorWithKeyAscendingSelector(key IString, ascending bool, selector objc.SEL) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("sortDescriptorWithKey:ascending:selector:"), key, ascending, selector)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SortDescriptorWithKeyAscendingSelector) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SortDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SortDescriptor */

// Forces a securely decoded sort descriptor to allow evaluation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSortDescriptor/allowEvaluation()
func (s_ SortDescriptor) AllowEvaluation() {
	objc.Send[objc.ID](s_.ID, objc.Sel("allowEvaluation"))
}/* debug [instance_methods/method]: AllowEvaluation */


// Returns a comparison result value that indicates the sort order of two objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSortDescriptor/compare(_:to:)
func (s_ SortDescriptor) CompareObjectToObject(object1 objc.IObject, object2 objc.IObject) ComparisonResult {
	rv := objc.Send[ComparisonResult](s_.ID, objc.Sel("compareObject:toObject:"), object1, object2)
	return rv
}/* debug [instance_methods/method]: CompareObjectToObject */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SortDescriptor */

// A Boolean value that indicates whether the receiver specifies sorting in ascending order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSortDescriptor/ascending
func (s_ SortDescriptor) Ascending() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("ascending"))
	return rv
}/* debug [instance_properties/getter]: ascending */


// The comparator for the sort descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSortDescriptor/comparator
func (s_ SortDescriptor) Comparator() Comparator /* not a class type */ {
	rv := objc.Send[Comparator](s_.ID, objc.Sel("comparator"))
	return rv
}/* debug [instance_properties/getter]: comparator */


// The key that specifies the property to compare during sorting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSortDescriptor/key
func (s_ SortDescriptor) Key() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("key"))
	return rv
}/* debug [instance_properties/getter]: key */


// Returns a sort descriptor that reverses the sort order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSortDescriptor/reversedSortDescriptor
func (s_ SortDescriptor) ReversedSortDescriptor() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("reversedSortDescriptor"))
	return rv
}/* debug [instance_properties/getter]: reversedSortDescriptor */


// The selector for comparing objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSortDescriptor/selector
func (s_ SortDescriptor) Selector() objc.SEL {
	rv := objc.Send[objc.SEL](s_.ID, objc.Sel("selector"))
	return rv
}/* debug [instance_properties/getter]: selector */


// The sort descriptors of the fetch request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/sortDescriptors
func (s_ SortDescriptor) SortDescriptors() ISortDescriptor {
	rv := objc.Send[SortDescriptor](s_.ID, objc.Sel("sortDescriptors"))
	return rv
}/* debug [instance_properties/getter]: sortDescriptors */


// The sort descriptors of the fetch request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/sortDescriptors
func (s_ SortDescriptor) SetSortDescriptors(value ISortDescriptor) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSortDescriptors:"), value)
}/* debug [instance_properties/setter]: sortDescriptors */


// The key path that specifies the property to compare during sorting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssortdescriptor/keypath
func (s_ SortDescriptor) KeyPath() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("keyPath"))
	return rv
}/* debug [instance_properties/getter]: keyPath */


// The key path that specifies the property to compare during sorting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssortdescriptor/keypath
func (s_ SortDescriptor) SetKeyPath(value objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setKeyPath:"), value)
}/* debug [instance_properties/setter]: keyPath */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSSortDescriptor */


