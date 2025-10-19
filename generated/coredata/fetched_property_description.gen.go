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

// A description object used to define which properties are fetched from Core Data. [Full Topic]
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




