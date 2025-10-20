// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [FetchedPropertyDescription] class.
var (
	fetchedPropertyDescriptionClass     _FetchedPropertyDescriptionClass
	fetchedPropertyDescriptionClassOnce sync.Once
)

func getFetchedPropertyDescriptionClass() _FetchedPropertyDescriptionClass {
	fetchedPropertyDescriptionClassOnce.Do(func() {
		fetchedPropertyDescriptionClass = _FetchedPropertyDescriptionClass{objc.GetClass("NSFetchedPropertyDescription")}
	})
	return fetchedPropertyDescriptionClass
}

type _FetchedPropertyDescriptionClass struct {
	class objc.Class
}

// An interface definition for the [FetchedPropertyDescription] class.
type IFetchedPropertyDescription interface {
	IPropertyDescription
}

// A description object used to define which properties are fetched from Core Data.
//
// An example might be a iTunes playlist, if expressed as a property of a containing object. Songs don’t belong to a particular playlist, especially in the case that they’re on a remote server. The playlist may remain even after the songs have been deleted, or the remote server has become inaccessible. Note, however, that unlike a playlist a fetched property is static—it does not dynamically update itself as objects in the destination entity change. The effect of a fetched property is similar to executing a fetch request yourself and placing the results in a transient attribute, although with the framework managing the details. In particular, a fetched property is not fetched until it is requested, and the results are then cached until the object is turned into a fault. You use ( ) to manually refresh the properties—this causes the fetch request associated with this property to be executed again when the object fault is next fired. Unlike other relationships, which are all sets, fetched properties are represented by an ordered object just as if you executed the fetch request yourself. The fetch request associated with the property can have a sort ordering. The value for a fetched property of a managed object does not support .
//
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
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchedPropertyDescription/fetchRequest
func (f_ FetchedPropertyDescription) FetchRequest() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("fetchRequest"))
	return rv
}

// SetFetchRequest sets the value of the fetchRequest property.
// The fetch request of the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSFetchedPropertyDescription/fetchRequest
func (f_ FetchedPropertyDescription) SetFetchRequest(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFetchRequest:"), value)
}


