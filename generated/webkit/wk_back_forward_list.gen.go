// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [BackForwardList] class.
type IBackForwardList interface {
	objectivec.IObject
	// properties:
	ForwardList() []IBackForwardListItem
	BackItem() IWKBackForwardListItem
	SetBackItem(value IWKBackForwardListItem)
	BackList() IWKBackForwardListItem
	SetBackList(value IWKBackForwardListItem)
	CurrentItem() IWKBackForwardListItem
	SetCurrentItem(value IWKBackForwardListItem)
	ForwardItem() IWKBackForwardListItem
	SetForwardItem(value IWKBackForwardListItem)
	BackForwardList() IWKBackForwardList
	SetBackForwardList(value IWKBackForwardList)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (bc _BackForwardListClass) Alloc() BackForwardList {
	rv := objc.Send[BackForwardList](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The array of items that follow the current item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKBackForwardList/forwardList
func (b_ BackForwardList) ForwardList() []IBackForwardListItem {
	rv := objc.Send[[]BackForwardListItem](b_.ID, objc.Sel("forwardList"))
	return rv
}


// The item immediately preceding the current item, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkbackforwardlist/backitem
func (b_ BackForwardList) BackItem() IWKBackForwardListItem {
	rv := objc.Send[BackForwardListItem](b_.ID, objc.Sel("backItem"))
	return rv
}


// The item immediately preceding the current item, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkbackforwardlist/backitem
func (b_ BackForwardList) SetBackItem(value IWKBackForwardListItem) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBackItem:"), value)
}


// The array of items that precede the current item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkbackforwardlist/backlist
func (b_ BackForwardList) BackList() IWKBackForwardListItem {
	rv := objc.Send[BackForwardListItem](b_.ID, objc.Sel("backList"))
	return rv
}


// The array of items that precede the current item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkbackforwardlist/backlist
func (b_ BackForwardList) SetBackList(value IWKBackForwardListItem) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBackList:"), value)
}


// The current item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkbackforwardlist/currentitem
func (b_ BackForwardList) CurrentItem() IWKBackForwardListItem {
	rv := objc.Send[BackForwardListItem](b_.ID, objc.Sel("currentItem"))
	return rv
}


// The current item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkbackforwardlist/currentitem
func (b_ BackForwardList) SetCurrentItem(value IWKBackForwardListItem) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setCurrentItem:"), value)
}


// The item immediately following the current item, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkbackforwardlist/forwarditem
func (b_ BackForwardList) ForwardItem() IWKBackForwardListItem {
	rv := objc.Send[BackForwardListItem](b_.ID, objc.Sel("forwardItem"))
	return rv
}


// The item immediately following the current item, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkbackforwardlist/forwarditem
func (b_ BackForwardList) SetForwardItem(value IWKBackForwardListItem) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setForwardItem:"), value)
}


// The web view’s back-forward list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/backforwardlist
func (b_ BackForwardList) BackForwardList() IWKBackForwardList {
	rv := objc.Send[BackForwardList](b_.ID, objc.Sel("backForwardList"))
	return rv
}


// The web view’s back-forward list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebview/backforwardlist
func (b_ BackForwardList) SetBackForwardList(value IWKBackForwardList) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBackForwardList:"), value)
}



