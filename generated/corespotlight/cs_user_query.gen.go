// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [CSUserQuery] class.
type ICSUserQuery interface {
	ICSSearchQuery
	Cancel()
	Start()
	UserEngagedWithItemVisibleItemsUserInteractionType(item ICSSearchableItem, visibleItems []CSSearchableItem, userInteractionType ICSUserInteraction)
	UserEngagedWithSuggestionVisibleSuggestionsUserInteractionType(suggestion ICSSuggestion, visibleSuggestions []CSSuggestion, userInteractionType ICSUserInteraction)
	FoundSuggestionCount() int
	FoundSuggestionsHandler() unsafe.Pointer
	SetFoundSuggestionsHandler(value unsafe.Pointer)
	FoundItemsHandler() unsafe.Pointer
	SetFoundItemsHandler(value unsafe.Pointer)
	Responses() unsafe.Pointer
	SetResponses(value unsafe.Pointer)
	Suggestions() unsafe.Pointer
	SetSuggestions(value unsafe.Pointer)
}

// A type you use to initiate searches from your interface and offer suggested text completions.
//
// A object provides the back-end support for your app’s search features. Combine this object with your app’s search interface to perform lexical and semantic searches of human-entered search terms. You can configure a query object to return ranked or unranked results. You can also use it to get a list of suggestions to display from your search interface. When the text in your search control changes, create a query object to begin searching for results based on the current text. You use a query object only once to perform a search. If the text changes again while you a previous query is in progress, cancel the old query and execute the new one. For this reason, it’s a good idea to delay the start of each query until there is a sufficient gap between changes. Configure the query parameters using a object, which you can reuse for multiple queries. The context lets you configure the behavior for ranking results, specify the maximum number of results and suggestions, and filter the results using a predicate string. When you’re ready to start the query, choose one of the following options: Get the value of the property and iterate over the results. Configure the property and call to execute the query manually. Each query runs until Spotlight returns the requested maximum number of results. If you don’t specify the maximum number of results, Spotlight runs until it returns all results. To end a search before you receive all the results, call the method. Cancelling a query is especially important if you’re about to start a new query with an updated search string. For more information about configuring a object, see .


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

// Alloc allocates a new instance without initialization.
func (cc _CSUserQueryClass) Alloc() CSUserQuery {
	rv := objc.Send[CSUserQuery](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Creates a new user query that searches for the specified term.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSUserQuery/init(userQueryString:userQueryContext:)
func NewCSUserQueryWithUserQueryStringUserQueryContext(userQueryString string, userQueryContext ICSUserQueryContext) CSUserQuery {
	instance := getCSUserQueryClass().Alloc()
	rv := objc.Send[CSUserQuery](instance.ID, objc.Sel("initWithUserQueryString:userQueryContext:"), objc.String(userQueryString), userQueryContext)
	rv.Autorelease()
	return rv
}



// Performs one-time tasks that prepare Spotlight to search for content in all search indexes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSUserQuery/prepare()
func (cc _CSUserQueryClass) Prepare() {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("prepare"))
}


// Performs one-time tasks that prepare Spotlight to search for content in one or more protected search indexes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSUserQuery/prepareProtectionClasses(_:)
func (cc _CSUserQueryClass) PrepareProtectionClasses(protectionClasses []string) {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("prepareProtectionClasses:"), protectionClasses)
}


// Cancels the current query operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSUserQuery/cancel()
func (c_ CSUserQuery) Cancel() {
	objc.Send[objc.ID](c_.ID, objc.Sel("cancel"))
}


// Starts searching the index for items that match the current query string and parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSUserQuery/start()
func (c_ CSUserQuery) Start() {
	objc.Send[objc.ID](c_.ID, objc.Sel("start"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSUserQuery/userEngagedWithItem:visibleItems:userInteractionType:
func (c_ CSUserQuery) UserEngagedWithItemVisibleItemsUserInteractionType(item ICSSearchableItem, visibleItems []CSSearchableItem, userInteractionType ICSUserInteraction) {
	objc.Send[objc.ID](c_.ID, objc.Sel("userEngagedWithItem:visibleItems:userInteractionType:"), item, visibleItems, userInteractionType)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSUserQuery/userEngagedWithSuggestion:visibleSuggestions:userInteractionType:
func (c_ CSUserQuery) UserEngagedWithSuggestionVisibleSuggestionsUserInteractionType(suggestion ICSSuggestion, visibleSuggestions []CSSuggestion, userInteractionType ICSUserInteraction) {
	objc.Send[objc.ID](c_.ID, objc.Sel("userEngagedWithSuggestion:visibleSuggestions:userInteractionType:"), suggestion, visibleSuggestions, userInteractionType)
}


// The number of suggested items the query found so far.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSUserQuery/foundSuggestionCount
func (c_ CSUserQuery) FoundSuggestionCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("foundSuggestionCount"))
	return rv
}


// The block to execute when the query delivers a new batch of suggested items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSUserQuery/foundSuggestionsHandler
func (c_ CSUserQuery) FoundSuggestionsHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("foundSuggestionsHandler"))
	return rv
}


// The block to execute when the query delivers a new batch of suggested items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSUserQuery/foundSuggestionsHandler
func (c_ CSUserQuery) SetFoundSuggestionsHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFoundSuggestionsHandler:"), value)
}


// The block to execute when the query delivers a new batch of matching items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchquery/founditemshandler
func (c_ CSUserQuery) FoundItemsHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("foundItemsHandler"))
	return rv
}


// The block to execute when the query delivers a new batch of matching items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchquery/founditemshandler
func (c_ CSUserQuery) SetFoundItemsHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFoundItemsHandler:"), value)
}


// The matching results and suggestions for the current query string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/csuserquery/responses-swift.property
func (c_ CSUserQuery) Responses() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("responses"))
	return rv
}


// The matching results and suggestions for the current query string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/csuserquery/responses-swift.property
func (c_ CSUserQuery) SetResponses(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setResponses:"), value)
}


// An asynchronous sequence of suggested completions for the current query text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/csuserquery/suggestions-swift.property
func (c_ CSUserQuery) Suggestions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("suggestions"))
	return rv
}


// An asynchronous sequence of suggested completions for the current query text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/csuserquery/suggestions-swift.property
func (c_ CSUserQuery) SetSuggestions(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSuggestions:"), value)
}


