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
}

// An immutable collection of values representing different aspects of an item for an extension to act upon.
//
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
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionItem/attachments
func (e_ ExtensionItem) Attachments() []ItemProvider {
	rv := objc.Send[[]ItemProvider](e_.ID, objc.Sel("attachments"))
	return rv
}


// SetAttachments sets the value of the attachments property.
// An optional array of media data associated with the extension item.

//
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
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionItem/attributedContentText
func (e_ ExtensionItem) AttributedContentText() NSAttributedString {
	rv := objc.Send[NSAttributedString](e_.ID, objc.Sel("attributedContentText"))
	return rv
}


// SetAttributedContentText sets the value of the attributedContentText property.
// An optional string describing the extension item content.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionItem/attributedContentText
func (e_ ExtensionItem) SetAttributedContentText(value IAttributedString) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setAttributedContentText:"), value)
}

// An optional title for the item.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionItem/attributedTitle
func (e_ ExtensionItem) AttributedTitle() NSAttributedString {
	rv := objc.Send[NSAttributedString](e_.ID, objc.Sel("attributedTitle"))
	return rv
}


// SetAttributedTitle sets the value of the attributedTitle property.
// An optional title for the item.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionItem/attributedTitle
func (e_ ExtensionItem) SetAttributedTitle(value IAttributedString) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setAttributedTitle:"), value)
}

// An optional dictionary of keys and values corresponding to the extension item’s properties.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionItem/userInfo
func (e_ ExtensionItem) UserInfo() objc.ID {
	rv := objc.Send[objc.ID](e_.ID, objc.Sel("userInfo"))
	return rv
}


// SetUserInfo sets the value of the userInfo property.
// An optional dictionary of keys and values corresponding to the extension item’s properties.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionItem/userInfo
func (e_ ExtensionItem) SetUserInfo(value objc.ID) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setUserInfo:"), value)
}



