// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLQueryItem] class.
var (
	URLQueryItemClass     _URLQueryItemClass
	URLQueryItemClassOnce sync.Once
)

func getURLQueryItemClass() _URLQueryItemClass {
	URLQueryItemClassOnce.Do(func() {
		URLQueryItemClass = _URLQueryItemClass{objc.GetClass("NSURLQueryItem")}
	})
	return URLQueryItemClass
}

type _URLQueryItemClass struct {
	class objc.Class
}

// An interface definition for the [URLQueryItem] class.
type IURLQueryItem interface {
	objectivec.IObject
}

// An object representing a single name/value pair for an item in the query portion of a URL.
//
// In Swift, this object bridges to ; use when you need reference semantics or other Foundation-specific behavior. You use query items with the property of an object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLQueryItem
type URLQueryItem struct {
	objectivec.Object
}

// URLQueryItemFrom constructs a [URLQueryItem] from an unsafe.Pointer.
//
// An object representing a single name/value pair for an item in the query portion of a URL.
func URLQueryItemFrom(ptr unsafe.Pointer) URLQueryItem {
	return URLQueryItem{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _URLQueryItemClass) Alloc() URLQueryItem {
	rv := objc.Send[URLQueryItem](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _URLQueryItemClass) New() URLQueryItem {
	rv := objc.Send[URLQueryItem](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLQueryItem) Init() URLQueryItem {
	rv := objc.Send[URLQueryItem](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLQueryItem) Autorelease() URLQueryItem {
	rv := objc.Send[URLQueryItem](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLQueryItem creates a new URLQueryItem instance.
func NewURLQueryItem() URLQueryItem {
	return getURLQueryItemClass().New()
}


// The value for the query item.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLQueryItem/value
func (u_ URLQueryItem) Value() appkit.string {
	rv := objc.Send[appkit.string](u_.ID, objc.Sel("value"))
	return rv
}

// The query URL component as an array of name/value pairs.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/queryitems
func (u_ URLQueryItem) QueryItems() URLQueryItem {
	rv := objc.Send[URLQueryItem](u_.ID, objc.Sel("queryItems"))
	return rv
}


// SetQueryItems sets the value of the queryItems property.
// The query URL component as an array of name/value pairs.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/queryitems
func (u_ URLQueryItem) SetQueryItems(value IURLQueryItem) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setQueryItems:"), value)
}

// The name of the query item.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlqueryitem/name
func (u_ URLQueryItem) Name() appkit.string {
	rv := objc.Send[appkit.string](u_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
// The name of the query item.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlqueryitem/name
func (u_ URLQueryItem) SetName(value appkit.string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setName:"), value)
}



