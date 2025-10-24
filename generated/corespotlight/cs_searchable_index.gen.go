// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/uniformtypeidentifiers"
)

/* debug [class.gen.go]: Generating class CSSearchableIndex */


/* debug [class_header]: Header for CSSearchableIndex */
// The class instance for the [CSSearchableIndex] class.
var (
	CSSearchableIndexClass     _CSSearchableIndexClass
	CSSearchableIndexClassOnce sync.Once
)

func getCSSearchableIndexClass() _CSSearchableIndexClass {
	CSSearchableIndexClassOnce.Do(func() {
		CSSearchableIndexClass = _CSSearchableIndexClass{objc.GetClass("CSSearchableIndex")}
	})
	return CSSearchableIndexClass
}

type _CSSearchableIndexClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CSSearchableIndex */
// An interface definition for the [CSSearchableIndex] class.
type ICSSearchableIndex interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CSSearchableIndex */
	// properties:
	IndexDelegate() unsafe.Pointer
	SetIndexDelegate(value unsafe.Pointer)
	IsEligibleForSearch() bool
	SetIsEligibleForSearch(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CSSearchableIndex */
	// methods:
	BeginIndexBatch()
	DeleteAllSearchableItemsWithCompletionHandler(completionHandler unsafe.Pointer)
	DeleteSearchableItemsWithDomainIdentifiersCompletionHandler(domainIdentifiers []string, completionHandler unsafe.Pointer)
	DeleteSearchableItemsWithIdentifiersCompletionHandler(identifiers []string, completionHandler unsafe.Pointer)
	EndIndexBatchWithClientStateCompletionHandler(clientState objc.IObject /* cross-framework: NSData */, completionHandler unsafe.Pointer)
	EndIndexBatchWithExpectedClientStateNewClientStateCompletionHandler(expectedClientState objc.IObject /* cross-framework: NSData */, newClientState objc.IObject /* cross-framework: NSData */, completionHandler unsafe.Pointer)
	FetchDataForBundleIdentifierItemIdentifierContentTypeCompletionHandler(bundleIdentifier objc.IObject /* cross-framework: NSString */, itemIdentifier objc.IObject /* cross-framework: NSString */, contentType uniformtypeidentifiers.UTType, completionHandler unsafe.Pointer)
	FetchLastClientStateWithCompletionHandler(completionHandler unsafe.Pointer)
	IndexSearchableItemsCompletionHandler(items []CSSearchableItem, completionHandler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CSSearchableIndex */
// Alloc allocates a new instance without initialization.
func (cc _CSSearchableIndexClass) Alloc() CSSearchableIndex {
	rv := objc.Send[CSSearchableIndex](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CSSearchableIndexClass) New() CSSearchableIndex {
	rv := objc.Send[CSSearchableIndex](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CSSearchableIndex) Init() CSSearchableIndex {
	rv := objc.Send[CSSearchableIndex](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CSSearchableIndex) Autorelease() CSSearchableIndex {
	rv := objc.Send[CSSearchableIndex](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCSSearchableIndex creates a new CSSearchableIndex instance.
func NewCSSearchableIndex() CSSearchableIndex {
	return getCSSearchableIndexClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CSSearchableIndex */
// An on-device index for your app’s searchable content.
//
// A object manages an on-device index for your app’s searchable content. To make your app’s content searchable, package it in one or more objects and add them to the index. You can create as many searchable indexes as you need to manage your content, and you can apply different levels of encryption to protect the content in each index. When you execute a query, Core Spotlight searches your app’s indexes for the requested information and returns the results to your code. Put your content into a custom that you create. Custom indexes support batch operations and additional levels of data protection. Place sensitive personal information in protected indexes to encrypt that content, and prevent its disclosure without proper authorization from the owner of the device. Although you can put content into the default index, you can’t encrypt the content in that index or perform batch operations to add content to it. When adding large amounts of data to the index, consider adding it in batches to minimize risk. Batch-based updates make it easier to handle errors that might occur during the indexing process. For each batch, you provide client-state information to identify the current batch. If your app or extension crashes while a batch operation is in progress, you can use that state information to determine where to start indexing again later. Modify custom objects only on one thread or task at a time. It’s a programming error to access a custom index from multiple threads simultaneously. When performing batch updates on an index, start each new batch operation only after calling the or method of the previous batch operation.


// An on-device index for your app’s searchable content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndex
type CSSearchableIndex struct {
	objectivec.Object
}

// CSSearchableIndexFrom constructs a [CSSearchableIndex] from an unsafe.Pointer.
//
// An on-device index for your app’s searchable content.
func CSSearchableIndexFrom(ptr unsafe.Pointer) CSSearchableIndex {
	return CSSearchableIndex{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CSSearchableIndex */

// Returns an on-device index with the specified name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndex/init(name:)
func NewCSSearchableIndexWithName(name objc.IObject /* cross-framework: NSString */) CSSearchableIndex {
	instance := getCSSearchableIndexClass().Alloc()
	rv := objc.Send[CSSearchableIndex](instance.ID, objc.Sel("initWithName:"), name)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCSSearchableIndexWithName */


// Returns an on-device index with the specified name and data protection class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndex/init(name:protectionClass:)
func NewCSSearchableIndexWithNameProtectionClass(name objc.IObject /* cross-framework: NSString */, protectionClass FileProtectionType /* not a class type */) CSSearchableIndex {
	instance := getCSSearchableIndexClass().Alloc()
	rv := objc.Send[CSSearchableIndex](instance.ID, objc.Sel("initWithName:protectionClass:"), name, protectionClass)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCSSearchableIndexWithNameProtectionClass */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CSSearchableIndex */

// Returns the default on-device index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndex/default()
func (cc _CSSearchableIndexClass) DefaultSearchableIndex() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("defaultSearchableIndex"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DefaultSearchableIndex) */


// Returns a Boolean value that indicates whether indexing is available on the current device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndex/isIndexingAvailable()
func (cc _CSSearchableIndexClass) IsIndexingAvailable() bool {
	rv := objc.Send[bool](objc.ID(cc.class), objc.Sel("isIndexingAvailable"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=IsIndexingAvailable) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CSSearchableIndex */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CSSearchableIndex */

// Begins a batch of updates to an index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndex/beginBatch()
func (c_ CSSearchableIndex) BeginIndexBatch() {
	objc.Send[objc.ID](c_.ID, objc.Sel("beginIndexBatch"))
}/* debug [instance_methods/method]: BeginIndexBatch */


// Deletes all searchable items from the index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndex/deleteAllSearchableItems(completionHandler:)
func (c_ CSSearchableIndex) DeleteAllSearchableItemsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("deleteAllSearchableItemsWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: DeleteAllSearchableItemsWithCompletionHandler */


// Removes from the index all searchable items associated with the specified domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndex/deleteSearchableItems(withDomainIdentifiers:completionHandler:)
func (c_ CSSearchableIndex) DeleteSearchableItemsWithDomainIdentifiersCompletionHandler(domainIdentifiers []string, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("deleteSearchableItemsWithDomainIdentifiers:completionHandler:"), domainIdentifiers, completionHandler)
}/* debug [instance_methods/method]: DeleteSearchableItemsWithDomainIdentifiersCompletionHandler */


// Removes from the index all items with the specified identifiers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndex/deleteSearchableItems(withIdentifiers:completionHandler:)
func (c_ CSSearchableIndex) DeleteSearchableItemsWithIdentifiersCompletionHandler(identifiers []string, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("deleteSearchableItemsWithIdentifiers:completionHandler:"), identifiers, completionHandler)
}/* debug [instance_methods/method]: DeleteSearchableItemsWithIdentifiersCompletionHandler */


// Ends a batch of index updates and stores the specified state information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndex/endBatch(withClientState:completionHandler:)
func (c_ CSSearchableIndex) EndIndexBatchWithClientStateCompletionHandler(clientState objc.IObject /* cross-framework: NSData */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("endIndexBatchWithClientState:completionHandler:"), clientState, completionHandler)
}/* debug [instance_methods/method]: EndIndexBatchWithClientStateCompletionHandler */


// Ends a batch of index updates and stores the specified state information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndex/endIndexBatch(expectedClientState:newClientState:completionHandler:)
func (c_ CSSearchableIndex) EndIndexBatchWithExpectedClientStateNewClientStateCompletionHandler(expectedClientState objc.IObject /* cross-framework: NSData */, newClientState objc.IObject /* cross-framework: NSData */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("endIndexBatchWithExpectedClientState:newClientState:completionHandler:"), expectedClientState, newClientState, completionHandler)
}/* debug [instance_methods/method]: EndIndexBatchWithExpectedClientStateNewClientStateCompletionHandler */


// Fetches data from an external provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndex/fetchData(forBundleIdentifier:itemIdentifier:contentType:completionHandler:)
func (c_ CSSearchableIndex) FetchDataForBundleIdentifierItemIdentifierContentTypeCompletionHandler(bundleIdentifier objc.IObject /* cross-framework: NSString */, itemIdentifier objc.IObject /* cross-framework: NSString */, contentType uniformtypeidentifiers.UTType, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("fetchDataForBundleIdentifier:itemIdentifier:contentType:completionHandler:"), bundleIdentifier, itemIdentifier, contentType, completionHandler)
}/* debug [instance_methods/method]: FetchDataForBundleIdentifierItemIdentifierContentTypeCompletionHandler */


// Fetches the app’s most recent client state information asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndex/fetchLastClientState(completionHandler:)
func (c_ CSSearchableIndex) FetchLastClientStateWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("fetchLastClientStateWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: FetchLastClientStateWithCompletionHandler */


// Adds or updates items in the index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndex/indexSearchableItems(_:completionHandler:)
func (c_ CSSearchableIndex) IndexSearchableItemsCompletionHandler(items []CSSearchableItem, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("indexSearchableItems:completionHandler:"), items, completionHandler)
}/* debug [instance_methods/method]: IndexSearchableItemsCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CSSearchableIndex */

// The delegate object that can handle index-management tasks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndex/indexDelegate
func (c_ CSSearchableIndex) IndexDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("indexDelegate"))
	return rv
}/* debug [instance_properties/getter]: indexDelegate */


// The delegate object that can handle index-management tasks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndex/indexDelegate
func (c_ CSSearchableIndex) SetIndexDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIndexDelegate:"), value)
}/* debug [instance_properties/setter]: indexDelegate */


// A Boolean value that indicates whether the activity should be added to the on-device index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/isEligibleForSearch
func (c_ CSSearchableIndex) IsEligibleForSearch() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEligibleForSearch"))
	return rv
}/* debug [instance_properties/getter]: isEligibleForSearch */


// A Boolean value that indicates whether the activity should be added to the on-device index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserActivity/isEligibleForSearch
func (c_ CSSearchableIndex) SetIsEligibleForSearch(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsEligibleForSearch:"), value)
}/* debug [instance_properties/setter]: isEligibleForSearch */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CSSearchableIndex */


