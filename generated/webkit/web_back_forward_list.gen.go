// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WebBackForwardList */

/* debug [class_header]: Header for WebBackForwardList */
// The class instance for the [WebBackForwardList] class.
var (
	WebBackForwardListClass     _WebBackForwardListClass
	WebBackForwardListClassOnce sync.Once
)

func getWebBackForwardListClass() _WebBackForwardListClass {
	WebBackForwardListClassOnce.Do(func() {
		WebBackForwardListClass = _WebBackForwardListClass{objc.GetClass("WebBackForwardList")}
	})
	return WebBackForwardListClass
}

type _WebBackForwardListClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for WebBackForwardList */
// An interface definition for the [WebBackForwardList] class.
type IWebBackForwardList interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for WebBackForwardList */
	// properties:
	BackItem() IWebHistoryItem
	BackListCount() int
	Capacity() int
	SetCapacity(value int)
	CurrentItem() IWebHistoryItem
	ForwardItem() IWebHistoryItem
	ForwardListCount() int
	OrderedLastVisitedDays() unsafe.Pointer
	SetOrderedLastVisitedDays(value unsafe.Pointer)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for WebBackForwardList */
	// methods:
	PageCacheSize() uint
	SetPageCacheSize(size uint)
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for WebBackForwardList */
// Alloc allocates a new instance without initialization.
func (wc _WebBackForwardListClass) Alloc() WebBackForwardList {
	rv := objc.Send[WebBackForwardList](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (wc _WebBackForwardListClass) New() WebBackForwardList {
	rv := objc.Send[WebBackForwardList](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WebBackForwardList) Init() WebBackForwardList {
	rv := objc.Send[WebBackForwardList](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WebBackForwardList) Autorelease() WebBackForwardList {
	rv := objc.Send[WebBackForwardList](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWebBackForwardList creates a new WebBackForwardList instance.
func NewWebBackForwardList() WebBackForwardList {
	return getWebBackForwardListClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for WebBackForwardList */
// A object maintains a list of visited pages used to go back and forward to the most recent page. A object maintains only the list data—it does not perform actual page loads (in other words, it does not make any client requests). If you need to perform a page load, see the method in to find out how to do this.
//
// Items are typically inserted in a back-forward list in the order they are visited. A object also maintains the notion of the current item (which is always at index ), the preceding item (which is at index ), and the following item (which is at index ). The and methods move the current item backward or forward by one. The method sets the current item to the specified item. All other methods that return objects do not change the value of the current item, they just return the requested item or items. You can also limit the number of history items stored in the back-forward list using . objects also control the number of pages cached. You can turn page caching off by setting the page cache size to using the method, or limit the number of pages cached by passing a value greater than 0.

// A object maintains a list of visited pages used to go back and forward to the most recent page. A object maintains only the list data—it does not perform actual page loads (in other words, it does not make any client requests). If you need to perform a page load, see the method in to find out how to do this.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebBackForwardList
type WebBackForwardList struct {
	objectivec.Object
}

// WebBackForwardListFrom constructs a [WebBackForwardList] from an unsafe.Pointer.
//
// A object maintains a list of visited pages used to go back and forward to the most recent page. A object maintains only the list data—it does not perform actual page loads (in other words, it does not make any client requests). If you need to perform a page load, see the method in to find out how to do this.
func WebBackForwardListFrom(ptr unsafe.Pointer) WebBackForwardList {
	return WebBackForwardList{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for WebBackForwardList */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for WebBackForwardList */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for WebBackForwardList */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for WebBackForwardList */

// Returns the maximum number of pages that the receiver can cache.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebBackForwardList/pageCacheSize()
func (w_ WebBackForwardList) PageCacheSize() uint {
	rv := objc.Send[uint](w_.ID, objc.Sel("pageCacheSize"))
	return rv
} /* debug [instance_methods/method]: PageCacheSize */

// Sets the maximum number of pages the receiver can cache.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebBackForwardList/setPageCacheSize(_:)
func (w_ WebBackForwardList) SetPageCacheSize(size uint) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPageCacheSize:"), size)
} /* debug [instance_methods/method]: SetPageCacheSize */

/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for WebBackForwardList */

// The item that precedes the current item in the back-forward list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebBackForwardList/backItem
func (w_ WebBackForwardList) BackItem() IWebHistoryItem {
	rv := objc.Send[WebHistoryItem](w_.ID, objc.Sel("backItem"))
	return rv
} /* debug [instance_properties/getter]: backItem */

// The number of items that precede the current item in the back-forward list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebBackForwardList/backListCount
func (w_ WebBackForwardList) BackListCount() int {
	rv := objc.Send[int](w_.ID, objc.Sel("backListCount"))
	return rv
} /* debug [instance_properties/getter]: backListCount */

// The maximum number of items that the back-forward list can contain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebBackForwardList/capacity
func (w_ WebBackForwardList) Capacity() int {
	rv := objc.Send[int](w_.ID, objc.Sel("capacity"))
	return rv
} /* debug [instance_properties/getter]: capacity */

// The maximum number of items that the back-forward list can contain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebBackForwardList/capacity
func (w_ WebBackForwardList) SetCapacity(value int) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCapacity:"), value)
} /* debug [instance_properties/setter]: capacity */

// The current item in the back-forward list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebBackForwardList/currentItem
func (w_ WebBackForwardList) CurrentItem() IWebHistoryItem {
	rv := objc.Send[WebHistoryItem](w_.ID, objc.Sel("currentItem"))
	return rv
} /* debug [instance_properties/getter]: currentItem */

// The item that follows the current item in the back-forward list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebBackForwardList/forwardItem
func (w_ WebBackForwardList) ForwardItem() IWebHistoryItem {
	rv := objc.Send[WebHistoryItem](w_.ID, objc.Sel("forwardItem"))
	return rv
} /* debug [instance_properties/getter]: forwardItem */

// The number of items that follow the current item in the back-forward list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebBackForwardList/forwardListCount
func (w_ WebBackForwardList) ForwardListCount() int {
	rv := objc.Send[int](w_.ID, objc.Sel("forwardListCount"))
	return rv
} /* debug [instance_properties/getter]: forwardListCount */

// An array of all calendar days represented in the web history.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webhistory/orderedlastvisiteddays
func (w_ WebBackForwardList) OrderedLastVisitedDays() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("orderedLastVisitedDays"))
	return rv
} /* debug [instance_properties/getter]: orderedLastVisitedDays */

// An array of all calendar days represented in the web history.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webhistory/orderedlastvisiteddays
func (w_ WebBackForwardList) SetOrderedLastVisitedDays(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setOrderedLastVisitedDays:"), value)
} /* debug [instance_properties/setter]: orderedLastVisitedDays */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class WebBackForwardList */
