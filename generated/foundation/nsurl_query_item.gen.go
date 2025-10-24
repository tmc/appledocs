// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSURLQueryItem */


/* debug [class_header]: Header for NSURLQueryItem */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for URLQueryItem */
// An interface definition for the [URLQueryItem] class.
type IURLQueryItem interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for URLQueryItem */
	// properties:
	Name() IString
	Value() IString
	QueryItems() IURLQueryItem
	SetQueryItems(value IURLQueryItem)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for URLQueryItem */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for URLQueryItem */
// Alloc allocates a new instance without initialization.
func (uc _URLQueryItemClass) Alloc() URLQueryItem {
	rv := objc.Send[URLQueryItem](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for URLQueryItem */
// An object representing a single name/value pair for an item in the query portion of a URL.
//
// In Swift, this object bridges to ; use when you need reference semantics or other Foundation-specific behavior. You use query items with the property of an object.


// An object representing a single name/value pair for an item in the query portion of a URL.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for URLQueryItem */

// Initializes a newly allocated query item with the specified name and value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLQueryItem/init(name:value:)
func NewURLQueryItemWithNameValue(name IString, value IString) URLQueryItem {
	instance := getURLQueryItemClass().Alloc()
	rv := objc.Send[URLQueryItem](instance.ID, objc.Sel("initWithName:value:"), name, value)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewURLQueryItemWithNameValue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for URLQueryItem */

// Creates a new query item with the specified name and value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLQueryItem/queryItemWithName:value:
func (uc _URLQueryItemClass) QueryItemWithNameValue(name IString, value IString) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(uc.class), objc.Sel("queryItemWithName:value:"), name, value)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=QueryItemWithNameValue) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for URLQueryItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for URLQueryItem */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for URLQueryItem */

// The name of the query item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLQueryItem/name
func (u_ URLQueryItem) Name() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The value for the query item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLQueryItem/value
func (u_ URLQueryItem) Value() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */


// The query URL component as an array of name/value pairs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/queryitems
func (u_ URLQueryItem) QueryItems() IURLQueryItem {
	rv := objc.Send[URLQueryItem](u_.ID, objc.Sel("queryItems"))
	return rv
}/* debug [instance_properties/getter]: queryItems */


// The query URL component as an array of name/value pairs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlcomponents/queryitems
func (u_ URLQueryItem) SetQueryItems(value IURLQueryItem) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setQueryItems:"), value)
}/* debug [instance_properties/setter]: queryItems */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSURLQueryItem */


