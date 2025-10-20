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
	extensionItemClass     _ExtensionItemClass
	extensionItemClassOnce sync.Once
)

func getExtensionItemClass() _ExtensionItemClass {
	extensionItemClassOnce.Do(func() {
		extensionItemClass = _ExtensionItemClass{objc.GetClass("NSExtensionItem")}
	})
	return extensionItemClass
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
	objc.Send[objc.ID](e_.ID, objc.Sel("setAttachments:"), value)
}


