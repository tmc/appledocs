// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CSSearchQueryContext */


/* debug [class_header]: Header for CSSearchQueryContext */
// The class instance for the [CSSearchQueryContext] class.
var (
	CSSearchQueryContextClass     _CSSearchQueryContextClass
	CSSearchQueryContextClassOnce sync.Once
)

func getCSSearchQueryContextClass() _CSSearchQueryContextClass {
	CSSearchQueryContextClassOnce.Do(func() {
		CSSearchQueryContextClass = _CSSearchQueryContextClass{objc.GetClass("CSSearchQueryContext")}
	})
	return CSSearchQueryContextClass
}

type _CSSearchQueryContextClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CSSearchQueryContext */
// An interface definition for the [CSSearchQueryContext] class.
type ICSSearchQueryContext interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CSSearchQueryContext */
	// properties:
	FetchAttributes() []string
	SetFetchAttributes(value []string)
	FilterQueries() []string
	SetFilterQueries(value []string)
	KeyboardLanguage() objc.IObject /* cross-framework: NSString */
	SetKeyboardLanguage(value objc.IObject /* cross-framework: NSString */)
	SourceOptions() CSSearchQuerySourceOptions
	SetSourceOptions(value CSSearchQuerySourceOptions)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CSSearchQueryContext */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CSSearchQueryContext */
// Alloc allocates a new instance without initialization.
func (cc _CSSearchQueryContextClass) Alloc() CSSearchQueryContext {
	rv := objc.Send[CSSearchQueryContext](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CSSearchQueryContextClass) New() CSSearchQueryContext {
	rv := objc.Send[CSSearchQueryContext](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CSSearchQueryContext) Init() CSSearchQueryContext {
	rv := objc.Send[CSSearchQueryContext](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CSSearchQueryContext) Autorelease() CSSearchQueryContext {
	rv := objc.Send[CSSearchQueryContext](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCSSearchQueryContext creates a new CSSearchQueryContext instance.
func NewCSSearchQueryContext() CSSearchQueryContext {
	return getCSSearchQueryContextClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CSSearchQueryContext */
// The behavior configuration to use for a search query.


// The behavior configuration to use for a search query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQueryContext
type CSSearchQueryContext struct {
	objectivec.Object
}

// CSSearchQueryContextFrom constructs a [CSSearchQueryContext] from an unsafe.Pointer.
//
// The behavior configuration to use for a search query.
func CSSearchQueryContextFrom(ptr unsafe.Pointer) CSSearchQueryContext {
	return CSSearchQueryContext{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CSSearchQueryContext *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CSSearchQueryContext */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CSSearchQueryContext */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CSSearchQueryContext */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CSSearchQueryContext */

// The attributes the system fetches for the searchable items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQueryContext/fetchAttributes
func (c_ CSSearchQueryContext) FetchAttributes() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("fetchAttributes"))
	return rv
}/* debug [instance_properties/getter]: fetchAttributes */


// The attributes the system fetches for the searchable items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQueryContext/fetchAttributes
func (c_ CSSearchQueryContext) SetFetchAttributes(value []string) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setFetchAttributes:"), nsArray)
}/* debug [instance_properties/setter]: fetchAttributes */


// The query string used to filter the results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQueryContext/filterQueries
func (c_ CSSearchQueryContext) FilterQueries() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("filterQueries"))
	return rv
}/* debug [instance_properties/getter]: filterQueries */


// The query string used to filter the results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQueryContext/filterQueries
func (c_ CSSearchQueryContext) SetFilterQueries(value []string) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setFilterQueries:"), nsArray)
}/* debug [instance_properties/setter]: filterQueries */


// The language used for the query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQueryContext/keyboardLanguage
func (c_ CSSearchQueryContext) KeyboardLanguage() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("keyboardLanguage"))
	return rv
}/* debug [instance_properties/getter]: keyboardLanguage */


// The language used for the query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQueryContext/keyboardLanguage
func (c_ CSSearchQueryContext) SetKeyboardLanguage(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKeyboardLanguage:"), value)
}/* debug [instance_properties/setter]: keyboardLanguage */


// The query source options to allow or deny Mail messages in the search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQueryContext/sourceOptions-swift.property
func (c_ CSSearchQueryContext) SourceOptions() CSSearchQuerySourceOptions {
	rv := objc.Send[CSSearchQuerySourceOptions](c_.ID, objc.Sel("sourceOptions"))
	return rv
}/* debug [instance_properties/getter]: sourceOptions */


// The query source options to allow or deny Mail messages in the search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQueryContext/sourceOptions-swift.property
func (c_ CSSearchQueryContext) SetSourceOptions(value CSSearchQuerySourceOptions) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSourceOptions:"), value)
}/* debug [instance_properties/setter]: sourceOptions */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CSSearchQueryContext */



