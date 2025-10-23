// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [FetchedPropertyDescription] class.
var (
	FetchedPropertyDescriptionClass     _FetchedPropertyDescriptionClass
	FetchedPropertyDescriptionClassOnce sync.Once
)

func getFetchedPropertyDescriptionClass() _FetchedPropertyDescriptionClass {
	FetchedPropertyDescriptionClassOnce.Do(func() {
		FetchedPropertyDescriptionClass = _FetchedPropertyDescriptionClass{objc.GetClass("NSFetchedPropertyDescription")}
	})
	return FetchedPropertyDescriptionClass
}

type _FetchedPropertyDescriptionClass struct {
	class objc.Class
}

// An interface definition for the [FetchedPropertyDescription] class.
type IFetchedPropertyDescription interface {
	IPropertyDescription
	FetchRequest() NSFetchRequest
	SetFetchRequest(value IFetchRequest)
	AffectedStores() NSPersistentStore
	SetAffectedStores(value IPersistentStore)
	FetchBatchSize() int
	SetFetchBatchSize(value int)
	FetchLimit() int
	SetFetchLimit(value int)
	FetchOffset() int
	SetFetchOffset(value int)
	Predicate() foundation.Predicate
	SetPredicate(value foundation.IPredicate)
}

// A description object used to define which properties are fetched from Core Data.
//
// An example might be a iTunes playlist, if expressed as a property of a containing object. Songs don’t belong to a particular playlist, especially in the case that they’re on a remote server. The playlist may remain even after the songs have been deleted, or the remote server has become inaccessible. Note, however, that unlike a playlist a fetched property is static—it does not dynamically update itself as objects in the destination entity change. The effect of a fetched property is similar to executing a fetch request yourself and placing the results in a transient attribute, although with the framework managing the details. In particular, a fetched property is not fetched until it is requested, and the results are then cached until the object is turned into a fault. You use ( ) to manually refresh the properties—this causes the fetch request associated with this property to be executed again when the object fault is next fired. Unlike other relationships, which are all sets, fetched properties are represented by an ordered object just as if you executed the fetch request yourself. The fetch request associated with the property can have a sort ordering. The value for a fetched property of a managed object does not support .


// A description object used to define which properties are fetched from Core Data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchedPropertyDescription
type FetchedPropertyDescription struct {
	PropertyDescription
}

// FetchedPropertyDescriptionFrom constructs a [FetchedPropertyDescription] from an unsafe.Pointer.
//
// A description object used to define which properties are fetched from Core Data.
func FetchedPropertyDescriptionFrom(ptr unsafe.Pointer) FetchedPropertyDescription {
	return FetchedPropertyDescription{
		PropertyDescription: PropertyDescriptionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (fc _FetchedPropertyDescriptionClass) Alloc() FetchedPropertyDescription {
	rv := objc.Send[FetchedPropertyDescription](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FetchedPropertyDescriptionClass) New() FetchedPropertyDescription {
	rv := objc.Send[FetchedPropertyDescription](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FetchedPropertyDescription) Init() FetchedPropertyDescription {
	rv := objc.Send[FetchedPropertyDescription](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FetchedPropertyDescription) Autorelease() FetchedPropertyDescription {
	rv := objc.Send[FetchedPropertyDescription](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFetchedPropertyDescription creates a new FetchedPropertyDescription instance.
func NewFetchedPropertyDescription() FetchedPropertyDescription {
	return getFetchedPropertyDescriptionClass().New()
}



// The fetch request of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchedPropertyDescription/fetchRequest
func (f_ FetchedPropertyDescription) FetchRequest() NSFetchRequest {
	rv := objc.Send[NSFetchRequest](f_.ID, objc.Sel("fetchRequest"))
	return rv
}


// The fetch request of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchedPropertyDescription/fetchRequest
func (f_ FetchedPropertyDescription) SetFetchRequest(value IFetchRequest) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFetchRequest:"), value)
}


// An array of persistent stores specified for the fetch request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/affectedstores
func (f_ FetchedPropertyDescription) AffectedStores() NSPersistentStore {
	rv := objc.Send[NSPersistentStore](f_.ID, objc.Sel("affectedStores"))
	return rv
}


// An array of persistent stores specified for the fetch request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/affectedstores
func (f_ FetchedPropertyDescription) SetAffectedStores(value IPersistentStore) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setAffectedStores:"), value)
}


// The batch size of the objects specified in the fetch request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/fetchbatchsize
func (f_ FetchedPropertyDescription) FetchBatchSize() int {
	rv := objc.Send[int](f_.ID, objc.Sel("fetchBatchSize"))
	return rv
}


// The batch size of the objects specified in the fetch request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/fetchbatchsize
func (f_ FetchedPropertyDescription) SetFetchBatchSize(value int) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFetchBatchSize:"), value)
}


// The fetch limit of the fetch request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/fetchlimit
func (f_ FetchedPropertyDescription) FetchLimit() int {
	rv := objc.Send[int](f_.ID, objc.Sel("fetchLimit"))
	return rv
}


// The fetch limit of the fetch request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/fetchlimit
func (f_ FetchedPropertyDescription) SetFetchLimit(value int) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFetchLimit:"), value)
}


// The fetch offset of the fetch request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/fetchoffset
func (f_ FetchedPropertyDescription) FetchOffset() int {
	rv := objc.Send[int](f_.ID, objc.Sel("fetchOffset"))
	return rv
}


// The fetch offset of the fetch request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/fetchoffset
func (f_ FetchedPropertyDescription) SetFetchOffset(value int) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFetchOffset:"), value)
}


// The predicate of the fetch request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/predicate
func (f_ FetchedPropertyDescription) Predicate() foundation.Predicate {
	rv := objc.Send[foundation.Predicate](f_.ID, objc.Sel("predicate"))
	return rv
}


// The predicate of the fetch request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/predicate
func (f_ FetchedPropertyDescription) SetPredicate(value foundation.IPredicate) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setPredicate:"), value)
}



