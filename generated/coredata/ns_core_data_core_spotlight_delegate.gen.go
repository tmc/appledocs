// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CoreDataCoreSpotlightDelegate] class.
var (
	CoreDataCoreSpotlightDelegateClass     _CoreDataCoreSpotlightDelegateClass
	CoreDataCoreSpotlightDelegateClassOnce sync.Once
)

func getCoreDataCoreSpotlightDelegateClass() _CoreDataCoreSpotlightDelegateClass {
	CoreDataCoreSpotlightDelegateClassOnce.Do(func() {
		CoreDataCoreSpotlightDelegateClass = _CoreDataCoreSpotlightDelegateClass{objc.GetClass("NSCoreDataCoreSpotlightDelegate")}
	})
	return CoreDataCoreSpotlightDelegateClass
}

type _CoreDataCoreSpotlightDelegateClass struct {
	class objc.Class
}

// An interface definition for the [CoreDataCoreSpotlightDelegate] class.
type ICoreDataCoreSpotlightDelegate interface {
	objectivec.IObject
	// properties:
	IsIndexingEnabled() bool /* primitive/slice/pointer. */
	SetIsIndexingEnabled(value bool /* primitive/slice/pointer. */)
	NSCoreDataCoreSpotlightExporter() string /* primitive/slice/pointer. */
	// methods:
}

// A set of methods that enable integration with Core Spotlight.


// A set of methods that enable integration with Core Spotlight.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSCoreDataCoreSpotlightDelegate
type CoreDataCoreSpotlightDelegate struct {
	objectivec.Object
}

// CoreDataCoreSpotlightDelegateFrom constructs a [CoreDataCoreSpotlightDelegate] from an unsafe.Pointer.
//
// A set of methods that enable integration with Core Spotlight.
func CoreDataCoreSpotlightDelegateFrom(ptr unsafe.Pointer) CoreDataCoreSpotlightDelegate {
	return CoreDataCoreSpotlightDelegate{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CoreDataCoreSpotlightDelegateClass) Alloc() CoreDataCoreSpotlightDelegate {
	rv := objc.Send[CoreDataCoreSpotlightDelegate](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CoreDataCoreSpotlightDelegateClass) New() CoreDataCoreSpotlightDelegate {
	rv := objc.Send[CoreDataCoreSpotlightDelegate](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CoreDataCoreSpotlightDelegate) Init() CoreDataCoreSpotlightDelegate {
	rv := objc.Send[CoreDataCoreSpotlightDelegate](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CoreDataCoreSpotlightDelegate) Autorelease() CoreDataCoreSpotlightDelegate {
	rv := objc.Send[CoreDataCoreSpotlightDelegate](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCoreDataCoreSpotlightDelegate creates a new CoreDataCoreSpotlightDelegate instance.
func NewCoreDataCoreSpotlightDelegate() CoreDataCoreSpotlightDelegate {
	return getCoreDataCoreSpotlightDelegateClass().New()
}



// A Boolean value that indicates whether Core Data is currently updating the Core Spotlight index with the persistent store’s entities.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nscoredatacorespotlightdelegate/isindexingenabled
func (c_ CoreDataCoreSpotlightDelegate) IsIndexingEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isIndexingEnabled"))
	return rv
}


// A Boolean value that indicates whether Core Data is currently updating the Core Spotlight index with the persistent store’s entities.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nscoredatacorespotlightdelegate/isindexingenabled
func (c_ CoreDataCoreSpotlightDelegate) SetIsIndexingEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsIndexingEnabled:"), value)
}


// The key you use to specify your Core Spotlight delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nscoredatacorespotlightexporter
func (c_ CoreDataCoreSpotlightDelegate) NSCoreDataCoreSpotlightExporter() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("NSCoreDataCoreSpotlightExporter"))
	return rv
}



