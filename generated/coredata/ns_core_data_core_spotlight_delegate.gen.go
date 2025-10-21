// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corespotlight"
	"github.com/tmc/appledocs/generated/foundation"
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
	AttributeSetForObject(object IManagedObject) corespotlight.CSSearchableItemAttributeSet
	DeleteSpotlightIndexWithCompletionHandler(completionHandler unsafe.Pointer)
	DomainIdentifier() foundation.String
	IndexName() foundation.String
	SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler(searchableIndex corespotlight.ICSSearchableIndex, acknowledgementHandler unsafe.Pointer)
	SearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler(searchableIndex corespotlight.ICSSearchableIndex, identifiers []string, acknowledgementHandler unsafe.Pointer)
	StartSpotlightIndexing()
	StopSpotlightIndexing()
}

// A set of methods that enable integration with Core Spotlight.
//
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




// Creates a Core Spotlight delegate with the specified store description and coordinator.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSCoreDataCoreSpotlightDelegate/init(forStoreWith:coordinator:)
func NewCoreDataCoreSpotlightDelegateForStoreWithDescriptionCoordinator(description IPersistentStoreDescription, psc IPersistentStoreCoordinator) CoreDataCoreSpotlightDelegate {
	instance := getCoreDataCoreSpotlightDelegateClass().Alloc()
	rv := objc.Send[CoreDataCoreSpotlightDelegate](instance.ID, objc.Sel("initForStoreWithDescription:coordinator:"), description, psc)
	rv.Autorelease()
	return rv
}



// Creates a Core Spotlight delegate with the specified store description and managed object model.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSCoreDataCoreSpotlightDelegate/init(forStoreWith:model:)
func NewCoreDataCoreSpotlightDelegateForStoreWithDescriptionModel(description IPersistentStoreDescription, model IManagedObjectModel) CoreDataCoreSpotlightDelegate {
	instance := getCoreDataCoreSpotlightDelegateClass().Alloc()
	rv := objc.Send[CoreDataCoreSpotlightDelegate](instance.ID, objc.Sel("initForStoreWithDescription:model:"), description, model)
	rv.Autorelease()
	return rv
}


// Returns the searchable attributes for the specified managed object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSCoreDataCoreSpotlightDelegate/attributeSet(for:)
func (c_ CoreDataCoreSpotlightDelegate) AttributeSetForObject(object IManagedObject) corespotlight.CSSearchableItemAttributeSet {
	rv := objc.Send[corespotlight.CSSearchableItemAttributeSet](c_.ID, objc.Sel("attributeSetForObject:"), object)
	return rv
}

// Deletes all searchable items from the configured index.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSCoreDataCoreSpotlightDelegate/deleteSpotlightIndex(completionHandler:)
func (c_ CoreDataCoreSpotlightDelegate) DeleteSpotlightIndexWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("deleteSpotlightIndexWithCompletionHandler:"), completionHandler)
}

// Returns the domain identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSCoreDataCoreSpotlightDelegate/domainIdentifier()
func (c_ CoreDataCoreSpotlightDelegate) DomainIdentifier() foundation.String {
	rv := objc.Send[foundation.String](c_.ID, objc.Sel("domainIdentifier"))
	return rv
}

// Returns the index’s name.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSCoreDataCoreSpotlightDelegate/indexName()
func (c_ CoreDataCoreSpotlightDelegate) IndexName() foundation.String {
	rv := objc.Send[foundation.String](c_.ID, objc.Sel("indexName"))
	return rv
}

// Reindexes all searchable items and clears any local state.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSCoreDataCoreSpotlightDelegate/searchableIndex(_:reindexAllSearchableItemsWithAcknowledgementHandler:)
func (c_ CoreDataCoreSpotlightDelegate) SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler(searchableIndex corespotlight.ICSSearchableIndex, acknowledgementHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("searchableIndex:reindexAllSearchableItemsWithAcknowledgementHandler:"), searchableIndex, acknowledgementHandler)
}

// Reindexes the searchable items for the specified identifiers.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSCoreDataCoreSpotlightDelegate/searchableIndex(_:reindexSearchableItemsWithIdentifiers:acknowledgementHandler:)
func (c_ CoreDataCoreSpotlightDelegate) SearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler(searchableIndex corespotlight.ICSSearchableIndex, identifiers []string, acknowledgementHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("searchableIndex:reindexSearchableItemsWithIdentifiers:acknowledgementHandler:"), searchableIndex, identifiers, acknowledgementHandler)
}

// Starts the indexing of the store’s entities.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSCoreDataCoreSpotlightDelegate/startSpotlightIndexing()
func (c_ CoreDataCoreSpotlightDelegate) StartSpotlightIndexing() {
	objc.Send[objc.ID](c_.ID, objc.Sel("startSpotlightIndexing"))
}

// Stops the indexing of the store’s entities.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSCoreDataCoreSpotlightDelegate/stopSpotlightIndexing()
func (c_ CoreDataCoreSpotlightDelegate) StopSpotlightIndexing() {
	objc.Send[objc.ID](c_.ID, objc.Sel("stopSpotlightIndexing"))
}

// A Boolean value that indicates whether Core Data is currently updating the Core Spotlight index with the persistent store’s entities.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSCoreDataCoreSpotlightDelegate/isIndexingEnabled
func (c_ CoreDataCoreSpotlightDelegate) IndexingEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("indexingEnabled"))
	return rv
}

// A Boolean value that indicates whether Core Data is currently updating the Core Spotlight index with the persistent store’s entities.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nscoredatacorespotlightdelegate/isindexingenabled
func (c_ CoreDataCoreSpotlightDelegate) IsIndexingEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isIndexingEnabled"))
	return rv
}


// SetIsIndexingEnabled sets the value of the isIndexingEnabled property.
// A Boolean value that indicates whether Core Data is currently updating the Core Spotlight index with the persistent store’s entities.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nscoredatacorespotlightdelegate/isindexingenabled
func (c_ CoreDataCoreSpotlightDelegate) SetIsIndexingEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsIndexingEnabled:"), value)
}

// The key you use to specify your Core Spotlight delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nscoredatacorespotlightexporter
func (c_ CoreDataCoreSpotlightDelegate) NSCoreDataCoreSpotlightExporter() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("NSCoreDataCoreSpotlightExporter"))
	return rv
}


