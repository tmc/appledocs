// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [SortDescriptor] class.
type ISortDescriptor interface {
	objectivec.IObject
	CompareObjectToObject(object1 objc.ID, object2 objc.ID) unsafe.Pointer
}

// An immutable description of how to order a collection of objects according to a property common to all the objects.
//
// You construct instances of by specifying the key path of the property to compare and the order of the sort (ascending or descending). Optionally, you can also specify a selector to use to perform the comparison, which allows you to specify other comparison selectors, such as and . Sorting raises an exception if the objects don’t respond to the sort descriptor’s comparison selector. You can use sort descriptors for the following: Sorting an array (an instance of or — see and ) Comparing two objects directly (see ) Specifying the order of objects that return from a Core Data fetch request (see )
//
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

// Alloc allocates a new instance without initialization.
func (sc _SortDescriptorClass) Alloc() SortDescriptor {
	rv := objc.Send[SortDescriptor](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates a sort descriptor by decoding from the coder you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSortDescriptor/init(coder:)
func NewSortDescriptorWithCoder(coder unsafe.Pointer) SortDescriptor {
	instance := getSortDescriptorClass().Alloc()
	rv := objc.Send[SortDescriptor](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}



// Creates a sort descriptor with a specified string key path and sort order.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSortDescriptor/init(key:ascending:)
func NewSortDescriptorWithKeyAscending(key string, ascending bool) SortDescriptor {
	instance := getSortDescriptorClass().Alloc()
	rv := objc.Send[SortDescriptor](instance.ID, objc.Sel("initWithKey:ascending:"), objc.String(key), ascending)
	rv.Autorelease()
	return rv
}



// Creates a sort descriptor with a specified string key path and ordering, and a comparator block.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSortDescriptor/init(key:ascending:comparator:)
func NewSortDescriptorWithKeyAscendingComparator(key string, ascending bool, cmptr unsafe.Pointer) SortDescriptor {
	instance := getSortDescriptorClass().Alloc()
	rv := objc.Send[SortDescriptor](instance.ID, objc.Sel("initWithKey:ascending:comparator:"), objc.String(key), ascending, cmptr)
	rv.Autorelease()
	return rv
}



// Creates a sort descriptor with a specified string key path, ordering, and comparison selector.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSortDescriptor/init(key:ascending:selector:)
func NewSortDescriptorWithKeyAscendingSelector(key string, ascending bool, selector objc.SEL) SortDescriptor {
	instance := getSortDescriptorClass().Alloc()
	rv := objc.Send[SortDescriptor](instance.ID, objc.Sel("initWithKey:ascending:selector:"), objc.String(key), ascending, selector)
	rv.Autorelease()
	return rv
}


// Creates and returns a sort descriptor initialized with the specified key path and ordering, and a comparator block.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSortDescriptor/sortDescriptorWithKey:ascending:comparator:
func (sc _SortDescriptorClass) SortDescriptorWithKeyAscendingComparator(key string, ascending bool, cmptr unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("sortDescriptorWithKey:ascending:comparator:"), objc.String(key), ascending, cmptr)
	return rv
}

// Returns a comparison result value that indicates the sort order of two objects.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSortDescriptor/compare(_:to:)
func (s_ SortDescriptor) CompareObjectToObject(object1 objc.ID, object2 objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("compareObject:toObject:"), object1, object2)
	return rv
}

// The sort descriptors of the fetch request.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/sortDescriptors
func (s_ SortDescriptor) SortDescriptors() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("sortDescriptors"))
	return rv
}


// SetSortDescriptors sets the value of the sortDescriptors property.
// The sort descriptors of the fetch request.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchRequest/sortDescriptors
func (s_ SortDescriptor) SetSortDescriptors(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSortDescriptors:"), value)
}

// Returns a sort descriptor that reverses the sort order.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssortdescriptor/reversedsortdescriptor
func (s_ SortDescriptor) ReversedSortDescriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("reversedSortDescriptor"))
	return rv
}


// SetReversedSortDescriptor sets the value of the reversedSortDescriptor property.
// Returns a sort descriptor that reverses the sort order.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssortdescriptor/reversedsortdescriptor
func (s_ SortDescriptor) SetReversedSortDescriptor(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setReversedSortDescriptor:"), value)
}

// The comparator for the sort descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssortdescriptor/comparator
func (s_ SortDescriptor) Comparator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("comparator"))
	return rv
}


// SetComparator sets the value of the comparator property.
// The comparator for the sort descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssortdescriptor/comparator
func (s_ SortDescriptor) SetComparator(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setComparator:"), value)
}

// The key path that specifies the property to compare during sorting.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssortdescriptor/keypath
func (s_ SortDescriptor) KeyPath() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("keyPath"))
	return rv
}


// SetKeyPath sets the value of the keyPath property.
// The key path that specifies the property to compare during sorting.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssortdescriptor/keypath
func (s_ SortDescriptor) SetKeyPath(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setKeyPath:"), value)
}

// A Boolean value that indicates whether the receiver specifies sorting in ascending order.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssortdescriptor/ascending
func (s_ SortDescriptor) Ascending() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("ascending"))
	return rv
}


// SetAscending sets the value of the ascending property.
// A Boolean value that indicates whether the receiver specifies sorting in ascending order.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssortdescriptor/ascending
func (s_ SortDescriptor) SetAscending(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAscending:"), value)
}

// The key that specifies the property to compare during sorting.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSortDescriptor/key
func (s_ SortDescriptor) Key() string {
	rv := objc.Send[string](s_.ID, objc.Sel("key"))
	return rv
}

// The selector for comparing objects.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSortDescriptor/selector
func (s_ SortDescriptor) Selector() objc.SEL {
	rv := objc.Send[objc.SEL](s_.ID, objc.Sel("selector"))
	return rv
}


