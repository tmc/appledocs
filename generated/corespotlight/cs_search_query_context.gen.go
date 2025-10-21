// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [CSSearchQueryContext] class.
type ICSSearchQueryContext interface {
	objectivec.IObject
}

// The behavior configuration to use for a search query.
//
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

// Alloc allocates a new instance without initialization.
func (cc _CSSearchQueryContextClass) Alloc() CSSearchQueryContext {
	rv := objc.Send[CSSearchQueryContext](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The attributes the system fetches for the searchable items.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQueryContext/fetchAttributes
func (c_ CSSearchQueryContext) FetchAttributes() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("fetchAttributes"))
	return rv
}


// SetFetchAttributes sets the value of the fetchAttributes property.
// The attributes the system fetches for the searchable items.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQueryContext/fetchAttributes
func (c_ CSSearchQueryContext) SetFetchAttributes(value []string) {
	// Convert Go slice to NSArray
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
}

// The query string used to filter the results.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQueryContext/filterQueries
func (c_ CSSearchQueryContext) FilterQueries() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("filterQueries"))
	return rv
}


// SetFilterQueries sets the value of the filterQueries property.
// The query string used to filter the results.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQueryContext/filterQueries
func (c_ CSSearchQueryContext) SetFilterQueries(value []string) {
	// Convert Go slice to NSArray
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
}

// The language used for the query.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQueryContext/keyboardLanguage
func (c_ CSSearchQueryContext) KeyboardLanguage() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("keyboardLanguage"))
	return rv
}


// SetKeyboardLanguage sets the value of the keyboardLanguage property.
// The language used for the query.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQueryContext/keyboardLanguage
func (c_ CSSearchQueryContext) SetKeyboardLanguage(value appkit.string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKeyboardLanguage:"), value)
}

// The query source options to allow or deny Mail messages in the search.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQueryContext/sourceOptions-swift.property
func (c_ CSSearchQueryContext) SourceOptions() CSSearchQuerySourceOptions {
	rv := objc.Send[CSSearchQuerySourceOptions](c_.ID, objc.Sel("sourceOptions"))
	return rv
}


// SetSourceOptions sets the value of the sourceOptions property.
// The query source options to allow or deny Mail messages in the search.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQueryContext/sourceOptions-swift.property
func (c_ CSSearchQueryContext) SetSourceOptions(value CSSearchQuerySourceOptions) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSourceOptions:"), value)
}



