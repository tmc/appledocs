// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class CSUserQueryContext */


/* debug [class_header]: Header for CSUserQueryContext */
// The class instance for the [CSUserQueryContext] class.
var (
	CSUserQueryContextClass     _CSUserQueryContextClass
	CSUserQueryContextClassOnce sync.Once
)

func getCSUserQueryContextClass() _CSUserQueryContextClass {
	CSUserQueryContextClassOnce.Do(func() {
		CSUserQueryContextClass = _CSUserQueryContextClass{objc.GetClass("CSUserQueryContext")}
	})
	return CSUserQueryContextClass
}

type _CSUserQueryContextClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CSUserQueryContext */
// An interface definition for the [CSUserQueryContext] class.
type ICSUserQueryContext interface {
	ICSSearchQueryContext
	
/* debug [class_interface_properties]: Properties for CSUserQueryContext */
	// properties:
	DisableSemanticSearch() bool
	SetDisableSemanticSearch(value bool)
	EnableRankedResults() bool
	SetEnableRankedResults(value bool)
	MaxRankedResultCount() int
	SetMaxRankedResultCount(value int)
	MaxResultCount() int
	SetMaxResultCount(value int)
	MaxSuggestionCount() int
	SetMaxSuggestionCount(value int)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CSUserQueryContext */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CSUserQueryContext */
// Alloc allocates a new instance without initialization.
func (cc _CSUserQueryContextClass) Alloc() CSUserQueryContext {
	rv := objc.Send[CSUserQueryContext](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CSUserQueryContextClass) New() CSUserQueryContext {
	rv := objc.Send[CSUserQueryContext](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CSUserQueryContext) Init() CSUserQueryContext {
	rv := objc.Send[CSUserQueryContext](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CSUserQueryContext) Autorelease() CSUserQueryContext {
	rv := objc.Send[CSUserQueryContext](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCSUserQueryContext creates a new CSUserQueryContext instance.
func NewCSUserQueryContext() CSUserQueryContext {
	return getCSUserQueryContextClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CSUserQueryContext */
// The configuration details to apply to a user query.
//
// Use an instance of to configure the search parameters for a object. This object stores configuration details that the query uses to modify the search results it delivers. For example, use this object to specify the maximum number of results or suggestions you want the query to return. You can also use it to enable or disable the ranking of results by Spotlight. For information about search filters and other configurable query parameters, see the parent class .


// The configuration details to apply to a user query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSUserQueryContext
type CSUserQueryContext struct {
	CSSearchQueryContext
}

// CSUserQueryContextFrom constructs a [CSUserQueryContext] from an unsafe.Pointer.
//
// The configuration details to apply to a user query.
func CSUserQueryContextFrom(ptr unsafe.Pointer) CSUserQueryContext {
	return CSUserQueryContext{
		CSSearchQueryContext: CSSearchQueryContextFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CSUserQueryContext */

// Creates a new query context object with an optional suggested search string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSUserQueryContext/init(currentSuggestion:)
func NewCSUserQueryContextWithCurrentSuggestion(currentSuggestion ICSSuggestion) CSUserQueryContext {
	rv := objc.Send[CSUserQueryContext](objc.ID(getCSUserQueryContextClass().class), objc.Sel("userQueryContextWithCurrentSuggestion:"), currentSuggestion)
	return rv
}/* debug [class_init_methods/constructor]: NewCSUserQueryContextWithCurrentSuggestion */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CSUserQueryContext */

// Creates a new query context object with an optional suggested search string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSUserQueryContext/init(currentSuggestion:)
func (cc _CSUserQueryContextClass) UserQueryContextWithCurrentSuggestion(currentSuggestion ICSSuggestion) CSUserQueryContext {
	rv := objc.Send[CSUserQueryContext](objc.ID(cc.class), objc.Sel("userQueryContextWithCurrentSuggestion:"), currentSuggestion)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=UserQueryContextWithCurrentSuggestion) */


// Returns the current behavior configuration for the user query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSUserQueryContext/userQueryContext
func (cc _CSUserQueryContextClass) UserQueryContext() CSUserQueryContext {
	rv := objc.Send[CSUserQueryContext](objc.ID(cc.class), objc.Sel("userQueryContext"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=UserQueryContext) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CSUserQueryContext */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CSUserQueryContext */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CSUserQueryContext */

// A Boolean value that indicates whether to exclude semantic-based search results from the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSUserQueryContext/disableSemanticSearch
func (c_ CSUserQueryContext) DisableSemanticSearch() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("disableSemanticSearch"))
	return rv
}/* debug [instance_properties/getter]: disableSemanticSearch */


// A Boolean value that indicates whether to exclude semantic-based search results from the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSUserQueryContext/disableSemanticSearch
func (c_ CSUserQueryContext) SetDisableSemanticSearch(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDisableSemanticSearch:"), value)
}/* debug [instance_properties/setter]: disableSemanticSearch */


// A Boolean value that indicates whether the query sorts results by their relevance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSUserQueryContext/enableRankedResults
func (c_ CSUserQueryContext) EnableRankedResults() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("enableRankedResults"))
	return rv
}/* debug [instance_properties/getter]: enableRankedResults */


// A Boolean value that indicates whether the query sorts results by their relevance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSUserQueryContext/enableRankedResults
func (c_ CSUserQueryContext) SetEnableRankedResults(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEnableRankedResults:"), value)
}/* debug [instance_properties/setter]: enableRankedResults */


// The maximum number of ranked results to return during the query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSUserQueryContext/maxRankedResultCount
func (c_ CSUserQueryContext) MaxRankedResultCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("maxRankedResultCount"))
	return rv
}/* debug [instance_properties/getter]: maxRankedResultCount */


// The maximum number of ranked results to return during the query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSUserQueryContext/maxRankedResultCount
func (c_ CSUserQueryContext) SetMaxRankedResultCount(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxRankedResultCount:"), value)
}/* debug [instance_properties/setter]: maxRankedResultCount */


// The maximum number of search results for the query to return.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSUserQueryContext/maxResultCount
func (c_ CSUserQueryContext) MaxResultCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("maxResultCount"))
	return rv
}/* debug [instance_properties/getter]: maxResultCount */


// The maximum number of search results for the query to return.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSUserQueryContext/maxResultCount
func (c_ CSUserQueryContext) SetMaxResultCount(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxResultCount:"), value)
}/* debug [instance_properties/setter]: maxResultCount */


// The maximum number of suggested text completions for the query to return.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSUserQueryContext/maxSuggestionCount
func (c_ CSUserQueryContext) MaxSuggestionCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("maxSuggestionCount"))
	return rv
}/* debug [instance_properties/getter]: maxSuggestionCount */


// The maximum number of suggested text completions for the query to return.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSUserQueryContext/maxSuggestionCount
func (c_ CSUserQueryContext) SetMaxSuggestionCount(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxSuggestionCount:"), value)
}/* debug [instance_properties/setter]: maxSuggestionCount */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CSUserQueryContext */


