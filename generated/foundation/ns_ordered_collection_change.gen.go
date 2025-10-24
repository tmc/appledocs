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
	NSNotFound() int
	ChangeType() CollectionChangeType
	SetChangeType(value CollectionChangeType)
	Index() int
	SetIndex(value int)
	GetObject() unsafe.Pointer
	SetGetObject(value unsafe.Pointer)
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



// When this property is set to a value other than , the receiver is one half of a move, and this value is the index of the change’s counterpart of the opposite type in the diff.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionChange/associatedIndex
func (o_ OrderedCollectionChange) AssociatedIndex() uint {
	rv := objc.Send[uint](o_.ID, objc.Sel("associatedIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotfound-9t5v2
func (o_ OrderedCollectionChange) NSNotFound() int {
	rv := objc.Send[int](o_.ID, objc.Sel("NSNotFound"))
	return rv
}


// The type of change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedcollectionchange/changetype
func (o_ OrderedCollectionChange) ChangeType() CollectionChangeType {
	rv := objc.Send[CollectionChangeType](o_.ID, objc.Sel("changeType"))
	return rv
}


// The type of change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedcollectionchange/changetype
func (o_ OrderedCollectionChange) SetChangeType(value CollectionChangeType) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setChangeType:"), value)
}


// The index location of the change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedcollectionchange/index
func (o_ OrderedCollectionChange) Index() int {
	rv := objc.Send[int](o_.ID, objc.Sel("index"))
	return rv
}


// The index location of the change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedcollectionchange/index
func (o_ OrderedCollectionChange) SetIndex(value int) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setIndex:"), value)
}


// An object the change inserts or removes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedcollectionchange/object
func (o_ OrderedCollectionChange) GetObject() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("object"))
	return rv
}


// An object the change inserts or removes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedcollectionchange/object
func (o_ OrderedCollectionChange) SetGetObject(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setGetObject:"), value)
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



