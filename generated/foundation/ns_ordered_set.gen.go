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
	OrderedSetClass     _OrderedSetClass
	OrderedSetClassOnce sync.Once
)

func getOrderedSetClass() _OrderedSetClass {
	OrderedSetClassOnce.Do(func() {
		OrderedSetClass = _OrderedSetClass{objc.GetClass("NSOrderedSet")}
	})
	return OrderedSetClass
}

type _OrderedSetClass struct {
	class objc.Class
}

// An interface definition for the [OrderedSet] class.
type IOrderedSet interface {
	objectivec.IObject
	DescriptionWithLocaleIndent(locale objc.ID, level uint) string
	EnumerateObjectsWithOptionsUsingBlock(opts unsafe.Pointer, block unsafe.Pointer)
	IndexOfObjectPassingTest(predicate unsafe.Pointer) uint
	ObjectAtIndex(idx uint) unsafe.Pointer
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

// Returns a string that represents the contents of the ordered set, formatted as a property list.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/description(withLocale:indent:)
func (o_ OrderedSet) DescriptionWithLocaleIndent(locale objc.ID, level uint) string {
	rv := objc.Send[string](o_.ID, objc.Sel("descriptionWithLocale:indent:"), locale, level)
	return rv
}

// Executes a given block using each object in the set, using the specified enumeration options.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/enumerateObjects(options:using:)
func (o_ OrderedSet) EnumerateObjectsWithOptionsUsingBlock(opts unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("enumerateObjectsWithOptions:usingBlock:"), opts, block)
}

// Returns the index of the object in the ordered set that passes a test in a given block.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/index(ofObjectPassingTest:)
func (o_ OrderedSet) IndexOfObjectPassingTest(predicate unsafe.Pointer) uint {
	rv := objc.Send[uint](o_.ID, objc.Sel("indexOfObjectPassingTest:"), predicate)
	return rv
}

// Returns the object at the specified index of the set.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/object(at:)
func (o_ OrderedSet) ObjectAtIndex(idx uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("objectAtIndex:"), idx)
	return rv
}

// The last object in the ordered set.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/lastObject
func (o_ OrderedSet) LastObject() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("lastObject"))
	return rv
}
