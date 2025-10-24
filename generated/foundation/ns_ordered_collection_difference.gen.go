// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSOrderedCollectionDifference */


/* debug [class_header]: Header for NSOrderedCollectionDifference */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for OrderedCollectionDifference */
// An interface definition for the [OrderedCollectionDifference] class.
type IOrderedCollectionDifference interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for OrderedCollectionDifference */
	// properties:
	HasChanges() bool
	Insertions() []OrderedCollectionChange
	Removals() []OrderedCollectionChange
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for OrderedCollectionDifference */
	// methods:
	InverseDifference() objectivec.IObject
	DifferenceByTransformingChangesWithBlock(block unsafe.Pointer) unsafe.Pointer
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for OrderedCollectionDifference */
// Alloc allocates a new instance without initialization.
func (oc _OrderedCollectionDifferenceClass) Alloc() OrderedCollectionDifference {
	rv := objc.Send[OrderedCollectionDifference](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for OrderedCollectionDifference */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for OrderedCollectionDifference */

// Creates an ordered collection difference using an array of ordered collection changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionDifference/init(changes:)
func NewOrderedCollectionDifferenceWithChanges(changes []OrderedCollectionChange) OrderedCollectionDifference {
	instance := getOrderedCollectionDifferenceClass().Alloc()
	rv := objc.Send[OrderedCollectionDifference](instance.ID, objc.Sel("initWithChanges:"), changes)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewOrderedCollectionDifferenceWithChanges */


// Creates an ordered collection difference from arrays of inserted and removed objects with corresponding sets of indices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionDifference/init(insert:insertedObjects:remove:removedObjects:)
func NewOrderedCollectionDifferenceWithInsertIndexesInsertedObjectsRemoveIndexesRemovedObjects(inserts IIndexSet, insertedObjects []objc.ID, removes IIndexSet, removedObjects []objc.ID) OrderedCollectionDifference {
	instance := getOrderedCollectionDifferenceClass().Alloc()
	rv := objc.Send[OrderedCollectionDifference](instance.ID, objc.Sel("initWithInsertIndexes:insertedObjects:removeIndexes:removedObjects:"), inserts, insertedObjects, removes, removedObjects)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewOrderedCollectionDifferenceWithInsertIndexesInsertedObjectsRemoveIndexesRemovedObjects */


// Creates an ordered collection difference from arrays of inserted and removed objects with corresponding sets of indices, in addition to an array of ordered collection changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionDifference/init(insert:insertedObjects:remove:removedObjects:additionalChanges:)
func NewOrderedCollectionDifferenceWithInsertIndexesInsertedObjectsRemoveIndexesRemovedObjectsAdditionalChanges(inserts IIndexSet, insertedObjects []objc.ID, removes IIndexSet, removedObjects []objc.ID, changes []OrderedCollectionChange) OrderedCollectionDifference {
	instance := getOrderedCollectionDifferenceClass().Alloc()
	rv := objc.Send[OrderedCollectionDifference](instance.ID, objc.Sel("initWithInsertIndexes:insertedObjects:removeIndexes:removedObjects:additionalChanges:"), inserts, insertedObjects, removes, removedObjects, changes)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewOrderedCollectionDifferenceWithInsertIndexesInsertedObjectsRemoveIndexesRemovedObjectsAdditionalChanges */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for OrderedCollectionDifference */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for OrderedCollectionDifference */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for OrderedCollectionDifference */

// Calculate the difference between two objects in the reverse direction of comparison.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionDifference/inverse()
func (o_ OrderedCollectionDifference) InverseDifference() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](o_.ID, objc.Sel("inverseDifference"))
	return rv
}/* debug [instance_methods/method]: InverseDifference */


// Create a new ordered collection difference by mapping over this difference’s members, processing the change objects with the block provided.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionDifference/transformingChanges(_:)
func (o_ OrderedCollectionDifference) DifferenceByTransformingChangesWithBlock(block unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("differenceByTransformingChangesWithBlock:"), block)
	return rv
}/* debug [instance_methods/method]: DifferenceByTransformingChangesWithBlock */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for OrderedCollectionDifference */

// A Boolean value that indicates if the difference has changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionDifference/hasChanges
func (o_ OrderedCollectionDifference) HasChanges() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("hasChanges"))
	return rv
}/* debug [instance_properties/getter]: hasChanges */


// A collection of insertion change objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionDifference/insertions
func (o_ OrderedCollectionDifference) Insertions() []OrderedCollectionChange {
	rv := objc.Send[[]OrderedCollectionChange](o_.ID, objc.Sel("insertions"))
	return rv
}/* debug [instance_properties/getter]: insertions */


// A collection of removal change objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionDifference/removals
func (o_ OrderedCollectionDifference) Removals() []OrderedCollectionChange {
	rv := objc.Send[[]OrderedCollectionChange](o_.ID, objc.Sel("removals"))
	return rv
}/* debug [instance_properties/getter]: removals */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSOrderedCollectionDifference */


