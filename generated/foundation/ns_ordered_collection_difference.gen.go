// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [OrderedCollectionDifference] class.
var (
	OrderedCollectionDifferenceClass     _OrderedCollectionDifferenceClass
	OrderedCollectionDifferenceClassOnce sync.Once
)

func getOrderedCollectionDifferenceClass() _OrderedCollectionDifferenceClass {
	OrderedCollectionDifferenceClassOnce.Do(func() {
		OrderedCollectionDifferenceClass = _OrderedCollectionDifferenceClass{objc.GetClass("NSOrderedCollectionDifference")}
	})
	return OrderedCollectionDifferenceClass
}

type _OrderedCollectionDifferenceClass struct {
	class objc.Class
}

// An interface definition for the [OrderedCollectionDifference] class.
type IOrderedCollectionDifference interface {
	objectivec.IObject
	// properties:
	HasChanges() bool /* primitive/slice/pointer. */
	Insertions() []OrderedCollectionChange /* primitive/slice/pointer. */
	Removals() []OrderedCollectionChange /* primitive/slice/pointer. */
	// methods:
	InverseDifference() unsafe.Pointer
	DifferenceByTransformingChangesWithBlock(block unsafe.Pointer) unsafe.Pointer
}

// An object representing the difference between two ordered collections.
//
// Use or one of its variations to get an instance of , which represents the difference between two ordered collections. For example, the following sample compares two arrays of strings to create a difference that represents the changes:


// An object representing the difference between two ordered collections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionDifference
type OrderedCollectionDifference struct {
	objectivec.Object
}

// OrderedCollectionDifferenceFrom constructs a [OrderedCollectionDifference] from an unsafe.Pointer.
//
// An object representing the difference between two ordered collections.
func OrderedCollectionDifferenceFrom(ptr unsafe.Pointer) OrderedCollectionDifference {
	return OrderedCollectionDifference{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (oc _OrderedCollectionDifferenceClass) Alloc() OrderedCollectionDifference {
	rv := objc.Send[OrderedCollectionDifference](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _OrderedCollectionDifferenceClass) New() OrderedCollectionDifference {
	rv := objc.Send[OrderedCollectionDifference](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OrderedCollectionDifference) Init() OrderedCollectionDifference {
	rv := objc.Send[OrderedCollectionDifference](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OrderedCollectionDifference) Autorelease() OrderedCollectionDifference {
	rv := objc.Send[OrderedCollectionDifference](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOrderedCollectionDifference creates a new OrderedCollectionDifference instance.
func NewOrderedCollectionDifference() OrderedCollectionDifference {
	return getOrderedCollectionDifferenceClass().New()
}



// Creates an ordered collection difference using an array of ordered collection changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionDifference/init(changes:)
func NewOrderedCollectionDifferenceWithChanges(changes []OrderedCollectionChange /* primitive/slice/pointer. */) OrderedCollectionDifference {
	instance := getOrderedCollectionDifferenceClass().Alloc()
	rv := objc.Send[OrderedCollectionDifference](instance.ID, objc.Sel("initWithChanges:"), changes)
	rv.Autorelease()
	return rv
}


// Creates an ordered collection difference from arrays of inserted and removed objects with corresponding sets of indices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionDifference/init(insert:insertedObjects:remove:removedObjects:)
func NewOrderedCollectionDifferenceWithInsertIndexesInsertedObjectsRemoveIndexesRemovedObjects(inserts IIndexSet, insertedObjects []objc.ID /* already interface */, removes IIndexSet, removedObjects []objc.ID /* already interface */) OrderedCollectionDifference {
	instance := getOrderedCollectionDifferenceClass().Alloc()
	rv := objc.Send[OrderedCollectionDifference](instance.ID, objc.Sel("initWithInsertIndexes:insertedObjects:removeIndexes:removedObjects:"), inserts, insertedObjects, removes, removedObjects)
	rv.Autorelease()
	return rv
}


// Creates an ordered collection difference from arrays of inserted and removed objects with corresponding sets of indices, in addition to an array of ordered collection changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionDifference/init(insert:insertedObjects:remove:removedObjects:additionalChanges:)
func NewOrderedCollectionDifferenceWithInsertIndexesInsertedObjectsRemoveIndexesRemovedObjectsAdditionalChanges(inserts IIndexSet, insertedObjects []objc.ID /* already interface */, removes IIndexSet, removedObjects []objc.ID /* already interface */, changes []OrderedCollectionChange /* primitive/slice/pointer. */) OrderedCollectionDifference {
	instance := getOrderedCollectionDifferenceClass().Alloc()
	rv := objc.Send[OrderedCollectionDifference](instance.ID, objc.Sel("initWithInsertIndexes:insertedObjects:removeIndexes:removedObjects:additionalChanges:"), inserts, insertedObjects, removes, removedObjects, changes)
	rv.Autorelease()
	return rv
}



// Calculate the difference between two objects in the reverse direction of comparison.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionDifference/inverse()
func (o_ OrderedCollectionDifference) InverseDifference() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("inverseDifference"))
	return rv
}


// Create a new ordered collection difference by mapping over this difference’s members, processing the change objects with the block provided.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionDifference/transformingChanges(_:)
func (o_ OrderedCollectionDifference) DifferenceByTransformingChangesWithBlock(block unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("differenceByTransformingChangesWithBlock:"), block)
	return rv
}


// A Boolean value that indicates if the difference has changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionDifference/hasChanges
func (o_ OrderedCollectionDifference) HasChanges() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("hasChanges"))
	return rv
}


// A collection of insertion change objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionDifference/insertions
func (o_ OrderedCollectionDifference) Insertions() []OrderedCollectionChange /* primitive/slice/pointer. */ {
	rv := objc.Send[[]OrderedCollectionChange](o_.ID, objc.Sel("insertions"))
	return rv
}


// A collection of removal change objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionDifference/removals
func (o_ OrderedCollectionDifference) Removals() []OrderedCollectionChange /* primitive/slice/pointer. */ {
	rv := objc.Send[[]OrderedCollectionChange](o_.ID, objc.Sel("removals"))
	return rv
}


