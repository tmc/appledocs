// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CSUserQuery */


/* debug [class_header]: Header for CSUserQuery */
// The class instance for the [CSUserQuery] class.
var (
	CSUserQueryClass     _CSUserQueryClass
	CSUserQueryClassOnce sync.Once
)

func getCSUserQueryClass() _CSUserQueryClass {
	CSUserQueryClassOnce.Do(func() {
		CSUserQueryClass = _CSUserQueryClass{objc.GetClass("CSUserQuery")}
	})
	return CSUserQueryClass
}

type _CSUserQueryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CSUserQuery */
// An interface definition for the [CSUserQuery] class.
type ICSUserQuery interface {
	ICSSearchQuery
	
/* debug [class_interface_properties]: Properties for CSUserQuery */
	// properties:
	FoundSuggestionCount() int
	FoundSuggestionsHandler() unsafe.Pointer
	SetFoundSuggestionsHandler(value unsafe.Pointer)
	FoundItemsHandler() objectivec.IObject
	SetFoundItemsHandler(value objectivec.IObject)
	Responses() objectivec.IObject
	SetResponses(value objectivec.IObject)
	Suggestions() objectivec.IObject
	SetSuggestions(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CSUserQuery */
	// methods:
	Cancel()
	Start()
	UserEngagedWithItemVisibleItemsUserInteractionType(item ICSSearchableItem, visibleItems []CSSearchableItem, userInteractionType CSUserInteraction)
	UserEngagedWithSuggestionVisibleSuggestionsUserInteractionType(suggestion ICSSuggestion, visibleSuggestions []CSSuggestion, userInteractionType CSUserInteraction)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CSUserQuery */
// Alloc allocates a new instance without initialization.
func (cc _CSUserQueryClass) Alloc() CSUserQuery {
	rv := objc.Send[CSUserQuery](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CSUserQueryClass) New() CSUserQuery {
	rv := objc.Send[CSUserQuery](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CSUserQuery) Init() CSUserQuery {
	rv := objc.Send[CSUserQuery](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CSUserQuery) Autorelease() CSUserQuery {
	rv := objc.Send[CSUserQuery](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCSUserQuery creates a new CSUserQuery instance.
func NewCSUserQuery() CSUserQuery {
	return getCSUserQueryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CSUserQuery */
// A type you use to initiate searches from your interface and offer suggested text completions.
//
// A object provides the back-end support for your app’s search features. Combine this object with your app’s search interface to perform lexical and semantic searches of human-entered search terms. You can configure a query object to return ranked or unranked results. You can also use it to get a list of suggestions to display from your search interface. When the text in your search control changes, create a query object to begin searching for results based on the current text. You use a query object only once to perform a search. If the text changes again while a previous query is in progress, cancel the old query and execute the new one. For this reason, it’s a good idea to delay the start of each query until there is a sufficient gap between changes. Configure the query parameters using a object, which you can reuse for multiple queries. The context lets you configure the behavior for ranking results, specify the maximum number of results and suggestions, and filter the results using a predicate string. When you’re ready to start the query, choose one of the following options: Get the value of the property and iterate over the results. Configure the property and call to execute the query manually. Each query runs until Spotlight returns the requested maximum number of results. If you don’t specify the maximum number of results, Spotlight runs until it returns all results. To end a search before you receive all the results, call the method. Cancelling a query is especially important if you’re about to start a new query with an updated search string. For more information about configuring a object, see .


// A type you use to initiate searches from your interface and offer suggested text completions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSUserQuery
type CSUserQuery struct {
	CSSearchQuery
}

// CSUserQueryFrom constructs a [CSUserQuery] from an unsafe.Pointer.
//
// A type you use to initiate searches from your interface and offer suggested text completions.
func CSUserQueryFrom(ptr unsafe.Pointer) CSUserQuery {
	return CSUserQuery{
		CSSearchQuery: CSSearchQueryFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CSUserQuery */

// Creates a new user query that searches for the specified term.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSUserQuery/init(userQueryString:userQueryContext:)
func NewCSUserQueryWithUserQueryStringUserQueryContext(userQueryString objc.IObject /* cross-framework: NSString */, userQueryContext ICSUserQueryContext) CSUserQuery {
	instance := getCSUserQueryClass().Alloc()
	rv := objc.Send[CSUserQuery](instance.ID, objc.Sel("initWithUserQueryString:userQueryContext:"), userQueryString, userQueryContext)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCSUserQueryWithUserQueryStringUserQueryContext */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CSUserQuery */

// Performs one-time tasks that prepare Spotlight to search for content in all search indexes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSUserQuery/prepare()
func (cc _CSUserQueryClass) Prepare() {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("prepare"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Prepare) */


// Performs one-time tasks that prepare Spotlight to search for content in one or more protected search indexes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSUserQuery/prepareProtectionClasses(_:)
func (cc _CSUserQueryClass) PrepareProtectionClasses(protectionClasses []string) {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("prepareProtectionClasses:"), protectionClasses)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PrepareProtectionClasses) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CSUserQuery */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CSUserQuery */

// Cancels the current query operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSUserQuery/cancel()
func (c_ CSUserQuery) Cancel() {
	objc.Send[objc.ID](c_.ID, objc.Sel("cancel"))
}/* debug [instance_methods/method]: Cancel */


// Starts searching the index for items that match the current query string and parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSUserQuery/start()
func (c_ CSUserQuery) Start() {
	objc.Send[objc.ID](c_.ID, objc.Sel("start"))
}/* debug [instance_methods/method]: Start */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSUserQuery/userEngagedWithItem:visibleItems:userInteractionType:
func (c_ CSUserQuery) UserEngagedWithItemVisibleItemsUserInteractionType(item ICSSearchableItem, visibleItems []CSSearchableItem, userInteractionType CSUserInteraction) {
	objc.Send[objc.ID](c_.ID, objc.Sel("userEngagedWithItem:visibleItems:userInteractionType:"), item, visibleItems, userInteractionType)
}/* debug [instance_methods/method]: UserEngagedWithItemVisibleItemsUserInteractionType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSUserQuery/userEngagedWithSuggestion:visibleSuggestions:userInteractionType:
func (c_ CSUserQuery) UserEngagedWithSuggestionVisibleSuggestionsUserInteractionType(suggestion ICSSuggestion, visibleSuggestions []CSSuggestion, userInteractionType CSUserInteraction) {
	objc.Send[objc.ID](c_.ID, objc.Sel("userEngagedWithSuggestion:visibleSuggestions:userInteractionType:"), suggestion, visibleSuggestions, userInteractionType)
}/* debug [instance_methods/method]: UserEngagedWithSuggestionVisibleSuggestionsUserInteractionType */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CSUserQuery */

// The number of suggested items the query found so far.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSUserQuery/foundSuggestionCount
func (c_ CSUserQuery) FoundSuggestionCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("foundSuggestionCount"))
	return rv
}/* debug [instance_properties/getter]: foundSuggestionCount */


// The block to execute when the query delivers a new batch of suggested items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSUserQuery/foundSuggestionsHandler
func (c_ CSUserQuery) FoundSuggestionsHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("foundSuggestionsHandler"))
	return rv
}/* debug [instance_properties/getter]: foundSuggestionsHandler */


// The block to execute when the query delivers a new batch of suggested items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSUserQuery/foundSuggestionsHandler
func (c_ CSUserQuery) SetFoundSuggestionsHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFoundSuggestionsHandler:"), value)
}/* debug [instance_properties/setter]: foundSuggestionsHandler */


// The block to execute when the query delivers a new batch of matching items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchquery/founditemshandler
func (c_ CSUserQuery) FoundItemsHandler() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("foundItemsHandler"))
	return rv
}/* debug [instance_properties/getter]: foundItemsHandler */


// The block to execute when the query delivers a new batch of matching items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchquery/founditemshandler
func (c_ CSUserQuery) SetFoundItemsHandler(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFoundItemsHandler:"), value)
}/* debug [instance_properties/setter]: foundItemsHandler */


// The matching results and suggestions for the current query string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/csuserquery/responses-swift.property
func (c_ CSUserQuery) Responses() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("responses"))
	return rv
}/* debug [instance_properties/getter]: responses */


// The matching results and suggestions for the current query string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/csuserquery/responses-swift.property
func (c_ CSUserQuery) SetResponses(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setResponses:"), value)
}/* debug [instance_properties/setter]: responses */


// An asynchronous sequence of suggested completions for the current query text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/csuserquery/suggestions-swift.property
func (c_ CSUserQuery) Suggestions() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("suggestions"))
	return rv
}/* debug [instance_properties/getter]: suggestions */


// An asynchronous sequence of suggested completions for the current query text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/csuserquery/suggestions-swift.property
func (c_ CSUserQuery) SetSuggestions(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSuggestions:"), value)
}/* debug [instance_properties/setter]: suggestions */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CSUserQuery */


