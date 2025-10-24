// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WKBackForwardList */

/* debug [class_header]: Header for WKBackForwardList */
// The class instance for the [BackForwardList] class.
var (
	BackForwardListClass     _BackForwardListClass
	BackForwardListClassOnce sync.Once
)

func getBackForwardListClass() _BackForwardListClass {
	BackForwardListClassOnce.Do(func() {
		BackForwardListClass = _BackForwardListClass{objc.GetClass("WKBackForwardList")}
	})
	return BackForwardListClass
}

type _BackForwardListClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for BackForwardList */
// An interface definition for the [BackForwardList] class.
type IBackForwardList interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for BackForwardList */
	// properties:
	BackItem() IWKBackForwardListItem
	BackList() []BackForwardListItem
	CurrentItem() IWKBackForwardListItem
	ForwardItem() IWKBackForwardListItem
	ForwardList() []BackForwardListItem
	BackForwardList() IWKBackForwardList
	SetBackForwardList(value IWKBackForwardList)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for BackForwardList */
	// methods:
	ItemAtIndex(index int) IBackForwardListItem
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for BackForwardList */
// Alloc allocates a new instance without initialization.
func (bc _BackForwardListClass) Alloc() BackForwardList {
	rv := objc.Send[BackForwardList](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _BackForwardListClass) New() BackForwardList {
	rv := objc.Send[BackForwardList](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BackForwardList) Init() BackForwardList {
	rv := objc.Send[BackForwardList](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BackForwardList) Autorelease() BackForwardList {
	rv := objc.Send[BackForwardList](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBackForwardList creates a new BackForwardList instance.
func NewBackForwardList() BackForwardList {
	return getBackForwardListClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for BackForwardList */
// An object that manages the list of previously loaded webpages, which the web view uses for forward and backward navigation.
//
// Use a object to retrieve a web view’s previously loaded pages. Typically, you don’t create objects directly. Each web view creates one automatically and uses it to store the history of all loaded pages. Fetch this object from your web view’s property and use its contents to facilitate programmatic navigation.

// An object that manages the list of previously loaded webpages, which the web view uses for forward and backward navigation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKBackForwardList
type BackForwardList struct {
	objectivec.Object
}

// BackForwardListFrom constructs a [BackForwardList] from an unsafe.Pointer.
//
// An object that manages the list of previously loaded webpages, which the web view uses for forward and backward navigation.
func BackForwardListFrom(ptr unsafe.Pointer) BackForwardList {
	return BackForwardList{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for BackForwardList */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for BackForwardList */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for BackForwardList */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for BackForwardList */

// Returns the item at the relative offset from the current item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKBackForwardList/item(at:)
func (b_ BackForwardList) ItemAtIndex(index int) IBackForwardListItem {
	rv := objc.Send[BackForwardListItem](b_.ID, objc.Sel("itemAtIndex:"), index)
	return rv
} /* debug [instance_methods/method]: ItemAtIndex */

/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for BackForwardList */

// The item immediately preceding the current item, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKBackForwardList/backItem
func (b_ BackForwardList) BackItem() IWKBackForwardListItem {
	rv := objc.Send[BackForwardListItem](b_.ID, objc.Sel("backItem"))
	return rv
} /* debug [instance_properties/getter]: backItem */

// The array of items that precede the current item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKBackForwardList/backList
func (b_ BackForwardList) BackList() []BackForwardListItem {
	rv := objc.Send[[]BackForwardListItem](b_.ID, objc.Sel("backList"))
	return rv
} /* debug [instance_properties/getter]: backList */

// The current item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKBackForwardList/currentItem
func (b_ BackForwardList) CurrentItem() IWKBackForwardListItem {
	rv := objc.Send[BackForwardListItem](b_.ID, objc.Sel("currentItem"))
	return rv
} /* debug [instance_properties/getter]: currentItem */

// The item immediately following the current item, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKBackForwardList/forwardItem
func (b_ BackForwardList) ForwardItem() IWKBackForwardListItem {
	rv := objc.Send[BackForwardListItem](b_.ID, objc.Sel("forwardItem"))
	return rv
} /* debug [instance_properties/getter]: forwardItem */

// The array of items that follow the current item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKBackForwardList/forwardList
func (b_ BackForwardList) ForwardList() []BackForwardListItem {
	rv := objc.Send[[]BackForwardListItem](b_.ID, objc.Sel("forwardList"))
	return rv
} /* debug [instance_properties/getter]: forwardList */

// The web view’s back-forward list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/backforwardlist
func (b_ BackForwardList) BackForwardList() IWKBackForwardList {
	rv := objc.Send[BackForwardList](b_.ID, objc.Sel("backForwardList"))
	return rv
} /* debug [instance_properties/getter]: backForwardList */

// The web view’s back-forward list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/backforwardlist
func (b_ BackForwardList) SetBackForwardList(value IWKBackForwardList) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBackForwardList:"), value)
} /* debug [instance_properties/setter]: backForwardList */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class WKBackForwardList */
