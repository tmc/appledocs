// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ExtensionItem] class.
var (
	ExtensionItemClass     _ExtensionItemClass
	ExtensionItemClassOnce sync.Once
)

func getExtensionItemClass() _ExtensionItemClass {
	ExtensionItemClassOnce.Do(func() {
		ExtensionItemClass = _ExtensionItemClass{objc.GetClass("NSExtensionItem")}
	})
	return ExtensionItemClass
}

type _ExtensionItemClass struct {
	class objc.Class
}

// An interface definition for the [ExtensionItem] class.
type IExtensionItem interface {
	objectivec.IObject
	// properties:
	Attachments() []ItemProvider
	SetAttachments(value []ItemProvider)
	AttributedContentText() IAttributedString
	SetAttributedContentText(value IAttributedString)
	AttributedTitle() IAttributedString
	SetAttributedTitle(value IAttributedString)
	UserInfo() unsafe.Pointer
	SetUserInfo(value unsafe.Pointer)
	// methods:
}

// An immutable collection of values representing different aspects of an item for an extension to act upon.


// An immutable collection of values representing different aspects of an item for an extension to act upon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionItem
type ExtensionItem struct {
	objectivec.Object
}

// ExtensionItemFrom constructs a [ExtensionItem] from an unsafe.Pointer.
//
// An immutable collection of values representing different aspects of an item for an extension to act upon.
func ExtensionItemFrom(ptr unsafe.Pointer) ExtensionItem {
	return ExtensionItem{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ec _ExtensionItemClass) Alloc() ExtensionItem {
	rv := objc.Send[ExtensionItem](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _ExtensionItemClass) New() ExtensionItem {
	rv := objc.Send[ExtensionItem](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ ExtensionItem) Init() ExtensionItem {
	rv := objc.Send[ExtensionItem](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ ExtensionItem) Autorelease() ExtensionItem {
	rv := objc.Send[ExtensionItem](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewExtensionItem creates a new ExtensionItem instance.
func NewExtensionItem() ExtensionItem {
	return getExtensionItemClass().New()
}



// An optional array of media data associated with the extension item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionItem/attachments
func (e_ ExtensionItem) Attachments() []ItemProvider {
	rv := objc.Send[[]ItemProvider](e_.ID, objc.Sel("attachments"))
	return rv
}


// An optional array of media data associated with the extension item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionItem/attachments
func (e_ ExtensionItem) SetAttachments(value []ItemProvider) {
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
	objc.Send[objc.ID](e_.ID, objc.Sel("setAttachments:"), nsArray)
}


// An optional string describing the extension item content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensionitem/attributedcontenttext
func (e_ ExtensionItem) AttributedContentText() IAttributedString {
	rv := objc.Send[AttributedString](e_.ID, objc.Sel("attributedContentText"))
	return rv
}


// An optional string describing the extension item content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensionitem/attributedcontenttext
func (e_ ExtensionItem) SetAttributedContentText(value IAttributedString) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setAttributedContentText:"), value)
}


// An optional title for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensionitem/attributedtitle
func (e_ ExtensionItem) AttributedTitle() IAttributedString {
	rv := objc.Send[AttributedString](e_.ID, objc.Sel("attributedTitle"))
	return rv
}


// An optional title for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensionitem/attributedtitle
func (e_ ExtensionItem) SetAttributedTitle(value IAttributedString) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setAttributedTitle:"), value)
}


// An optional dictionary of keys and values corresponding to the extension item’s properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensionitem/userinfo
func (e_ ExtensionItem) UserInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("userInfo"))
	return rv
}


// An optional dictionary of keys and values corresponding to the extension item’s properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensionitem/userinfo
func (e_ ExtensionItem) SetUserInfo(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setUserInfo:"), value)
}



