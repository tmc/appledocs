// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CSSearchQuery */


/* debug [class_header]: Header for CSSearchQuery */
// The class instance for the [CSSearchQuery] class.
var (
	CSSearchQueryClass     _CSSearchQueryClass
	CSSearchQueryClassOnce sync.Once
)

func getCSSearchQueryClass() _CSSearchQueryClass {
	CSSearchQueryClassOnce.Do(func() {
		CSSearchQueryClass = _CSSearchQueryClass{objc.GetClass("CSSearchQuery")}
	})
	return CSSearchQueryClass
}

type _CSSearchQueryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CSSearchQuery */
// An interface definition for the [CSSearchQuery] class.
type ICSSearchQuery interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CSSearchQuery */
	// properties:
	CompletionHandler() unsafe.Pointer
	SetCompletionHandler(value unsafe.Pointer)
	FoundItemCount() uint
	FoundItemsHandler() unsafe.Pointer
	SetFoundItemsHandler(value unsafe.Pointer)
	Cancelled() bool
	ProtectionClasses() []string
	SetProtectionClasses(value []string)
	CSQueryContinuationActionType() objc.IObject /* cross-framework: NSString */
	IsCancelled() bool
	SetIsCancelled(value bool)
	Results() objectivec.IObject
	SetResults(value objectivec.IObject)
	CSSearchQueryString() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CSSearchQuery */
	// methods:
	Cancel()
	Start()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CSSearchQuery */
// Alloc allocates a new instance without initialization.
func (cc _CSSearchQueryClass) Alloc() CSSearchQuery {
	rv := objc.Send[CSSearchQuery](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CSSearchQueryClass) New() CSSearchQuery {
	rv := objc.Send[CSSearchQuery](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CSSearchQuery) Init() CSSearchQuery {
	rv := objc.Send[CSSearchQuery](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CSSearchQuery) Autorelease() CSSearchQuery {
	rv := objc.Send[CSSearchQuery](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCSSearchQuery creates a new CSSearchQuery instance.
func NewCSSearchQuery() CSSearchQuery {
	return getCSSearchQueryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CSSearchQuery */
// A type you use to programmatically search the indexed app content.
//
// Use a object to search your app’s indexed content using a formatted search string. To perform a search, build a predicate string to specify the indexed attributes you want to search and the value you want them to match. After you start the query, you receive batches of results in the handlers you provide. Each object you create performs a single search operation and delivers the results back to your code. Build each predicate with an attribute name, one or more values, and either a comparison operator or the operator. Your predicate string takes one of the following forms: Queries search all of your app’s indexes by default. If your app encrypts some of its indexed data, you can limit your search to one or more of the encrypted indexes by updating the query’s property. The query must have access to the protected index to search it. For more information about how to construct predicate strings for your query, see .


// A type you use to programmatically search the indexed app content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQuery
type CSSearchQuery struct {
	objectivec.Object
}

// CSSearchQueryFrom constructs a [CSSearchQuery] from an unsafe.Pointer.
//
// A type you use to programmatically search the indexed app content.
func CSSearchQueryFrom(ptr unsafe.Pointer) CSSearchQuery {
	return CSSearchQuery{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CSSearchQuery */

// Initializes and returns a query object with the specified query string and item attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQuery/init(queryString:attributes:)
func NewCSSearchQueryWithQueryStringAttributes(queryString objc.IObject /* cross-framework: NSString */, attributes []string) CSSearchQuery {
	instance := getCSSearchQueryClass().Alloc()
	rv := objc.Send[CSSearchQuery](instance.ID, objc.Sel("initWithQueryString:attributes:"), queryString, attributes)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCSSearchQueryWithQueryStringAttributes */


// Initializes and returns a query object with the specified query string and query context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQuery/init(queryString:queryContext:)
func NewCSSearchQueryWithQueryStringQueryContext(queryString objc.IObject /* cross-framework: NSString */, queryContext ICSSearchQueryContext) CSSearchQuery {
	instance := getCSSearchQueryClass().Alloc()
	rv := objc.Send[CSSearchQuery](instance.ID, objc.Sel("initWithQueryString:queryContext:"), queryString, queryContext)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCSSearchQueryWithQueryStringQueryContext */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CSSearchQuery */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CSSearchQuery */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CSSearchQuery */

// Cancels the current query operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQuery/cancel()
func (c_ CSSearchQuery) Cancel() {
	objc.Send[objc.ID](c_.ID, objc.Sel("cancel"))
}/* debug [instance_methods/method]: Cancel */


// Starts searching the index for items that match the current query string and parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQuery/start()
func (c_ CSSearchQuery) Start() {
	objc.Send[objc.ID](c_.ID, objc.Sel("start"))
}/* debug [instance_methods/method]: Start */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CSSearchQuery */

// The block to execute when the query finishes delivering all results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQuery/completionHandler
func (c_ CSSearchQuery) CompletionHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("completionHandler"))
	return rv
}/* debug [instance_properties/getter]: completionHandler */


// The block to execute when the query finishes delivering all results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQuery/completionHandler
func (c_ CSSearchQuery) SetCompletionHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCompletionHandler:"), value)
}/* debug [instance_properties/setter]: completionHandler */


// The number of matching items found for the given query string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQuery/foundItemCount
func (c_ CSSearchQuery) FoundItemCount() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("foundItemCount"))
	return rv
}/* debug [instance_properties/getter]: foundItemCount */


// The block to execute when the query delivers a new batch of matching items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQuery/foundItemsHandler
func (c_ CSSearchQuery) FoundItemsHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("foundItemsHandler"))
	return rv
}/* debug [instance_properties/getter]: foundItemsHandler */


// The block to execute when the query delivers a new batch of matching items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQuery/foundItemsHandler
func (c_ CSSearchQuery) SetFoundItemsHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFoundItemsHandler:"), value)
}/* debug [instance_properties/setter]: foundItemsHandler */


// A Boolean value that indicates whether the current query is no longer running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQuery/isCancelled
func (c_ CSSearchQuery) Cancelled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("cancelled"))
	return rv
}/* debug [instance_properties/getter]: cancelled */


// The protection types of the indexes you want to search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQuery/protectionClasses
func (c_ CSSearchQuery) ProtectionClasses() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("protectionClasses"))
	return rv
}/* debug [instance_properties/getter]: protectionClasses */


// The protection types of the indexes you want to search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQuery/protectionClasses
func (c_ CSSearchQuery) SetProtectionClasses(value []string) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setProtectionClasses:"), nsArray)
}/* debug [instance_properties/setter]: protectionClasses */


// Indicates that the activity type to continue is a search or query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/csquerycontinuationactiontype
func (c_ CSSearchQuery) CSQueryContinuationActionType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CSQueryContinuationActionType"))
	return rv
}/* debug [instance_properties/getter]: CSQueryContinuationActionType */


// A Boolean value that indicates whether the current query is no longer running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchquery/iscancelled
func (c_ CSSearchQuery) IsCancelled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCancelled"))
	return rv
}/* debug [instance_properties/getter]: isCancelled */


// A Boolean value that indicates whether the current query is no longer running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchquery/iscancelled
func (c_ CSSearchQuery) SetIsCancelled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCancelled:"), value)
}/* debug [instance_properties/setter]: isCancelled */


// The results that match the current query string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchquery/results-swift.property
func (c_ CSSearchQuery) Results() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("results"))
	return rv
}/* debug [instance_properties/getter]: results */


// The results that match the current query string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchquery/results-swift.property
func (c_ CSSearchQuery) SetResults(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setResults:"), value)
}/* debug [instance_properties/setter]: results */


// Provides the key for the current query in the info dictionary of the user activity object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchquerystring
func (c_ CSSearchQuery) CSSearchQueryString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CSSearchQueryString"))
	return rv
}/* debug [instance_properties/getter]: CSSearchQueryString */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CSSearchQuery */


