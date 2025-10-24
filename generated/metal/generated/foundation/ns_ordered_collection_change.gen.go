// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [OrderedCollectionChange] class.
var (
	OrderedCollectionChangeClass     _OrderedCollectionChangeClass
	OrderedCollectionChangeClassOnce sync.Once
)

func getOrderedCollectionChangeClass() _OrderedCollectionChangeClass {
	OrderedCollectionChangeClassOnce.Do(func() {
		OrderedCollectionChangeClass = _OrderedCollectionChangeClass{objc.GetClass("NSOrderedCollectionChange")}
	})
	return OrderedCollectionChangeClass
}

type _OrderedCollectionChangeClass struct {
	class objc.Class
}

// An interface definition for the [OrderedCollectionChange] class.
type IOrderedCollectionChange interface {
	objectivec.IObject
	// properties:
	AssociatedIndex() uint
	ChangeType() CollectionChangeType /* not a class type */
	Index() uint
	GetObject() unsafe.Pointer
	NSNotFound() int
	HasChanges() bool
	SetHasChanges(value bool)
	Insertions() IOrderedCollectionChange
	SetInsertions(value IOrderedCollectionChange)
	Removals() IOrderedCollectionChange
	SetRemovals(value IOrderedCollectionChange)
	// methods:
}

// An object that represents an indexed change within an ordered collection.
//
// An ordered collection change represents changes by adding, removing, or moving objects within an ordered collection. Changes with an associated index indicate a move within the collection.


// An object that represents an indexed change within an ordered collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionChange
type OrderedCollectionChange struct {
	objectivec.Object
}

// OrderedCollectionChangeFrom constructs a [OrderedCollectionChange] from an unsafe.Pointer.
//
// An object that represents an indexed change within an ordered collection.
func OrderedCollectionChangeFrom(ptr unsafe.Pointer) OrderedCollectionChange {
	return OrderedCollectionChange{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (oc _OrderedCollectionChangeClass) Alloc() OrderedCollectionChange {
	rv := objc.Send[OrderedCollectionChange](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _OrderedCollectionChangeClass) New() OrderedCollectionChange {
	rv := objc.Send[OrderedCollectionChange](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OrderedCollectionChange) Init() OrderedCollectionChange {
	rv := objc.Send[OrderedCollectionChange](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OrderedCollectionChange) Autorelease() OrderedCollectionChange {
	rv := objc.Send[OrderedCollectionChange](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOrderedCollectionChange creates a new OrderedCollectionChange instance.
func NewOrderedCollectionChange() OrderedCollectionChange {
	return getOrderedCollectionChangeClass().New()
}



// Creates a change object that represents inserting or removing an object from an ordered collection at a specific index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionChange/init(object:type:index:)
func NewOrderedCollectionChangeWithObjectTypeIndex(anObject unsafe.Pointer, type_ CollectionChangeType /* not a class type */, index uint) OrderedCollectionChange {
	instance := getOrderedCollectionChangeClass().Alloc()
	rv := objc.Send[OrderedCollectionChange](instance.ID, objc.Sel("initWithObject:type:index:"), anObject, type_, index)
	rv.Autorelease()
	return rv
}


// Creates a change object that represents inserting, removing, or moving an object from an ordered collection at a specific index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionChange/init(object:type:index:associatedIndex:)
func NewOrderedCollectionChangeWithObjectTypeIndexAssociatedIndex(anObject unsafe.Pointer, type_ CollectionChangeType /* not a class type */, index uint, associatedIndex uint) OrderedCollectionChange {
	instance := getOrderedCollectionChangeClass().Alloc()
	rv := objc.Send[OrderedCollectionChange](instance.ID, objc.Sel("initWithObject:type:index:associatedIndex:"), anObject, type_, index, associatedIndex)
	rv.Autorelease()
	return rv
}



// Creates an change object that represents inserting or removing an object from an ordered collection at a specific index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionChange/changeWithObject:type:index:
func (oc _OrderedCollectionChangeClass) ChangeWithObjectTypeIndex(anObject unsafe.Pointer, type_ CollectionChangeType /* not a class type */, index uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("changeWithObject:type:index:"), anObject, type_, index)
	return rv
}


// Creates an change object that represents inserting or removing an object from an ordered collection at a specific index, matched with an associated location that infers a move within the collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionChange/changeWithObject:type:index:associatedIndex:
func (oc _OrderedCollectionChangeClass) ChangeWithObjectTypeIndexAssociatedIndex(anObject unsafe.Pointer, type_ CollectionChangeType /* not a class type */, index uint, associatedIndex uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("changeWithObject:type:index:associatedIndex:"), anObject, type_, index, associatedIndex)
	return rv
}


// When this property is set to a value other than , the receiver is one half of a move, and this value is the index of the change’s counterpart of the opposite type in the diff.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionChange/associatedIndex
func (o_ OrderedCollectionChange) AssociatedIndex() uint {
	rv := objc.Send[uint](o_.ID, objc.Sel("associatedIndex"))
	return rv
}


// The type of change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionChange/changeType
func (o_ OrderedCollectionChange) ChangeType() CollectionChangeType /* not a class type */ {
	rv := objc.Send[CollectionChangeType](o_.ID, objc.Sel("changeType"))
	return rv
}


// The index location of the change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionChange/index
func (o_ OrderedCollectionChange) Index() uint {
	rv := objc.Send[uint](o_.ID, objc.Sel("index"))
	return rv
}


// An object the change inserts or removes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionChange/object
func (o_ OrderedCollectionChange) GetObject() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("object"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotfound-9t5v2
func (o_ OrderedCollectionChange) NSNotFound() int {
	rv := objc.Send[int](o_.ID, objc.Sel("NSNotFound"))
	return rv
}


// A Boolean value that indicates if the difference has changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedcollectiondifference/haschanges
func (o_ OrderedCollectionChange) HasChanges() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("hasChanges"))
	return rv
}


// A Boolean value that indicates if the difference has changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedcollectiondifference/haschanges
func (o_ OrderedCollectionChange) SetHasChanges(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setHasChanges:"), value)
}


// A collection of insertion change objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedcollectiondifference/insertions
func (o_ OrderedCollectionChange) Insertions() IOrderedCollectionChange {
	rv := objc.Send[OrderedCollectionChange](o_.ID, objc.Sel("insertions"))
	return rv
}


// A collection of insertion change objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedcollectiondifference/insertions
func (o_ OrderedCollectionChange) SetInsertions(value IOrderedCollectionChange) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setInsertions:"), value)
}


// A collection of removal change objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedcollectiondifference/removals
func (o_ OrderedCollectionChange) Removals() IOrderedCollectionChange {
	rv := objc.Send[OrderedCollectionChange](o_.ID, objc.Sel("removals"))
	return rv
}


// A collection of removal change objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedcollectiondifference/removals
func (o_ OrderedCollectionChange) SetRemovals(value IOrderedCollectionChange) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setRemovals:"), value)
}


