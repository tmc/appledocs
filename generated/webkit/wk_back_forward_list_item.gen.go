// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [BackForwardListItem] class.
var (
	BackForwardListItemClass     _BackForwardListItemClass
	BackForwardListItemClassOnce sync.Once
)

func getBackForwardListItemClass() _BackForwardListItemClass {
	BackForwardListItemClassOnce.Do(func() {
		BackForwardListItemClass = _BackForwardListItemClass{objc.GetClass("WKBackForwardListItem")}
	})
	return BackForwardListItemClass
}

type _BackForwardListItemClass struct {
	class objc.Class
}

// An interface definition for the [BackForwardListItem] class.
type IBackForwardListItem interface {
	objectivec.IObject
	// properties:
	Title() objc.IObject /* cross-framework: NSString */
	URL() objc.IObject /* cross-framework: NSURL */
	InitialURL() objc.IObject /* cross-framework: URL */
	SetInitialURL(value objc.IObject /* cross-framework: URL */)
	// methods:
}

// A representation of a webpage that the web view previously visited.
//
// Use a object to get information about previously visited webpages. This object identifies the page’s title and URL. It also identifes the URL that requested the webpage. You don’t create objects directly. Instead, a object creates them in conjunction with its associated web view when the web view loads new pages.


// A representation of a webpage that the web view previously visited.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKBackForwardListItem
type BackForwardListItem struct {
	objectivec.Object
}

// BackForwardListItemFrom constructs a [BackForwardListItem] from an unsafe.Pointer.
//
// A representation of a webpage that the web view previously visited.
func BackForwardListItemFrom(ptr unsafe.Pointer) BackForwardListItem {
	return BackForwardListItem{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (bc _BackForwardListItemClass) Alloc() BackForwardListItem {
	rv := objc.Send[BackForwardListItem](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BackForwardListItemClass) New() BackForwardListItem {
	rv := objc.Send[BackForwardListItem](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BackForwardListItem) Init() BackForwardListItem {
	rv := objc.Send[BackForwardListItem](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BackForwardListItem) Autorelease() BackForwardListItem {
	rv := objc.Send[BackForwardListItem](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBackForwardListItem creates a new BackForwardListItem instance.
func NewBackForwardListItem() BackForwardListItem {
	return getBackForwardListItemClass().New()
}



// The title of the webpage this item represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKBackForwardListItem/title
func (b_ BackForwardListItem) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](b_.ID, objc.Sel("title"))
	return rv
}


// The URL of the webpage this item represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKBackForwardListItem/url
func (b_ BackForwardListItem) URL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](b_.ID, objc.Sel("URL"))
	return rv
}


// The source URL that originally asked the web view to load this page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkbackforwardlistitem/initialurl
func (b_ BackForwardListItem) InitialURL() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](b_.ID, objc.Sel("initialURL"))
	return rv
}


// The source URL that originally asked the web view to load this page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkbackforwardlistitem/initialurl
func (b_ BackForwardListItem) SetInitialURL(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setInitialURL:"), value)
}



