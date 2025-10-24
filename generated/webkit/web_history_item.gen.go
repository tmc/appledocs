// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WebHistoryItem */


/* debug [class_header]: Header for WebHistoryItem */
// The class instance for the [WebHistoryItem] class.
var (
	WebHistoryItemClass     _WebHistoryItemClass
	WebHistoryItemClassOnce sync.Once
)

func getWebHistoryItemClass() _WebHistoryItemClass {
	WebHistoryItemClassOnce.Do(func() {
		WebHistoryItemClass = _WebHistoryItemClass{objc.GetClass("WebHistoryItem")}
	})
	return WebHistoryItemClass
}

type _WebHistoryItemClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for WebHistoryItem */
// An interface definition for the [WebHistoryItem] class.
type IWebHistoryItem interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for WebHistoryItem */
	// properties:
	AlternateTitle() objc.IObject /* cross-framework: NSString */
	SetAlternateTitle(value objc.IObject /* cross-framework: NSString */)
	Icon() appkit.Image
	LastVisitedTimeInterval() float64
	OriginalURLString() objc.IObject /* cross-framework: NSString */
	Title() objc.IObject /* cross-framework: NSString */
	URLString() objc.IObject /* cross-framework: NSString */
	OrderedLastVisitedDays() objectivec.IObject
	SetOrderedLastVisitedDays(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for WebHistoryItem */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for WebHistoryItem */
// Alloc allocates a new instance without initialization.
func (wc _WebHistoryItemClass) Alloc() WebHistoryItem {
	rv := objc.Send[WebHistoryItem](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (wc _WebHistoryItemClass) New() WebHistoryItem {
	rv := objc.Send[WebHistoryItem](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WebHistoryItem) Init() WebHistoryItem {
	rv := objc.Send[WebHistoryItem](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WebHistoryItem) Autorelease() WebHistoryItem {
	rv := objc.Send[WebHistoryItem](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWebHistoryItem creates a new WebHistoryItem instance.
func NewWebHistoryItem() WebHistoryItem {
	return getWebHistoryItemClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for WebHistoryItem */
// WebHistoryItem objects encapsulate information about visiting a page so that users can return to that page. WebHistory and WebBackForwardList objects manage lists of WebHistoryItem objects. WebHistoryItem objects are created and added to these lists automatically when loading pages, so you do not need to create WebHistoryItem objects directly.


// WebHistoryItem objects encapsulate information about visiting a page so that users can return to that page. WebHistory and WebBackForwardList objects manage lists of WebHistoryItem objects. WebHistoryItem objects are created and added to these lists automatically when loading pages, so you do not need to create WebHistoryItem objects directly.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebHistoryItem
type WebHistoryItem struct {
	objectivec.Object
}

// WebHistoryItemFrom constructs a [WebHistoryItem] from an unsafe.Pointer.
//
// WebHistoryItem objects encapsulate information about visiting a page so that users can return to that page. WebHistory and WebBackForwardList objects manage lists of WebHistoryItem objects. WebHistoryItem objects are created and added to these lists automatically when loading pages, so you do not need to create WebHistoryItem objects directly.
func WebHistoryItemFrom(ptr unsafe.Pointer) WebHistoryItem {
	return WebHistoryItem{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for WebHistoryItem */

// Initializes the receiver with a URL, , a title specified by and the last time this item was visited specified by title, and time last visited.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebHistoryItem/init(urlString:title:lastVisitedTimeInterval:)
func NewWebHistoryItemWithURLStringTitleLastVisitedTimeInterval(URLString objc.IObject /* cross-framework: NSString */, title objc.IObject /* cross-framework: NSString */, time float64) WebHistoryItem {
	instance := getWebHistoryItemClass().Alloc()
	rv := objc.Send[WebHistoryItem](instance.ID, objc.Sel("initWithURLString:title:lastVisitedTimeInterval:"), URLString, title, time)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewWebHistoryItemWithURLStringTitleLastVisitedTimeInterval */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for WebHistoryItem */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for WebHistoryItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for WebHistoryItem */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for WebHistoryItem */

// An alternate title that may be used in place of the receiver’s page title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebHistoryItem/alternateTitle
func (w_ WebHistoryItem) AlternateTitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("alternateTitle"))
	return rv
}/* debug [instance_properties/getter]: alternateTitle */


// An alternate title that may be used in place of the receiver’s page title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebHistoryItem/alternateTitle
func (w_ WebHistoryItem) SetAlternateTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAlternateTitle:"), value)
}/* debug [instance_properties/setter]: alternateTitle */


// The icon for the receiver’s page, or if none exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebHistoryItem/icon
func (w_ WebHistoryItem) Icon() appkit.Image {
	rv := objc.Send[appkit.Image](w_.ID, objc.Sel("icon"))
	return rv
}/* debug [instance_properties/getter]: icon */


// The last time and date the receiver’s page was visited.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebHistoryItem/lastVisitedTimeInterval
func (w_ WebHistoryItem) LastVisitedTimeInterval() float64 {
	rv := objc.Send[float64](w_.ID, objc.Sel("lastVisitedTimeInterval"))
	return rv
}/* debug [instance_properties/getter]: lastVisitedTimeInterval */


// The string representation of the original URL for the receiver’s page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebHistoryItem/originalURLString
func (w_ WebHistoryItem) OriginalURLString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("originalURLString"))
	return rv
}/* debug [instance_properties/getter]: originalURLString */


// The receiver’s original page title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebHistoryItem/title
func (w_ WebHistoryItem) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// The string representation of the URL for the receiver’s page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebHistoryItem/urlString
func (w_ WebHistoryItem) URLString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("URLString"))
	return rv
}/* debug [instance_properties/getter]: URLString */


// An array of all calendar days represented in the web history.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webhistory/orderedlastvisiteddays
func (w_ WebHistoryItem) OrderedLastVisitedDays() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](w_.ID, objc.Sel("orderedLastVisitedDays"))
	return rv
}/* debug [instance_properties/getter]: orderedLastVisitedDays */


// An array of all calendar days represented in the web history.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webhistory/orderedlastvisiteddays
func (w_ WebHistoryItem) SetOrderedLastVisitedDays(value objectivec.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setOrderedLastVisitedDays:"), value)
}/* debug [instance_properties/setter]: orderedLastVisitedDays */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class WebHistoryItem */


