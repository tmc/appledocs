// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSExtensionItem */


/* debug [class_header]: Header for NSExtensionItem */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ExtensionItem */
// An interface definition for the [ExtensionItem] class.
type IExtensionItem interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ExtensionItem */
	// properties:
	Attachments() []ItemProvider
	SetAttachments(value []ItemProvider)
	AttributedContentText() IAttributedString
	SetAttributedContentText(value IAttributedString)
	AttributedTitle() IAttributedString
	SetAttributedTitle(value IAttributedString)
	UserInfo() objectivec.IObject
	SetUserInfo(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ExtensionItem */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ExtensionItem */
// Alloc allocates a new instance without initialization.
func (ec _ExtensionItemClass) Alloc() ExtensionItem {
	rv := objc.Send[ExtensionItem](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ExtensionItem */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ExtensionItem *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ExtensionItem */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ExtensionItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ExtensionItem */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ExtensionItem */

// An optional array of media data associated with the extension item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionItem/attachments
func (e_ ExtensionItem) Attachments() []ItemProvider {
	rv := objc.Send[[]ItemProvider](e_.ID, objc.Sel("attachments"))
	return rv
}/* debug [instance_properties/getter]: attachments */


// An optional array of media data associated with the extension item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionItem/attachments
func (e_ ExtensionItem) SetAttachments(value []ItemProvider) {
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
}/* debug [instance_properties/setter]: attachments */


// An optional string describing the extension item content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensionitem/attributedcontenttext
func (e_ ExtensionItem) AttributedContentText() IAttributedString {
	rv := objc.Send[AttributedString](e_.ID, objc.Sel("attributedContentText"))
	return rv
}/* debug [instance_properties/getter]: attributedContentText */


// An optional string describing the extension item content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensionitem/attributedcontenttext
func (e_ ExtensionItem) SetAttributedContentText(value IAttributedString) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setAttributedContentText:"), value)
}/* debug [instance_properties/setter]: attributedContentText */


// An optional title for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensionitem/attributedtitle
func (e_ ExtensionItem) AttributedTitle() IAttributedString {
	rv := objc.Send[AttributedString](e_.ID, objc.Sel("attributedTitle"))
	return rv
}/* debug [instance_properties/getter]: attributedTitle */


// An optional title for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensionitem/attributedtitle
func (e_ ExtensionItem) SetAttributedTitle(value IAttributedString) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setAttributedTitle:"), value)
}/* debug [instance_properties/setter]: attributedTitle */


// An optional dictionary of keys and values corresponding to the extension item’s properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensionitem/userinfo
func (e_ ExtensionItem) UserInfo() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](e_.ID, objc.Sel("userInfo"))
	return rv
}/* debug [instance_properties/getter]: userInfo */


// An optional dictionary of keys and values corresponding to the extension item’s properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensionitem/userinfo
func (e_ ExtensionItem) SetUserInfo(value objectivec.IObject) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setUserInfo:"), value)
}/* debug [instance_properties/setter]: userInfo */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSExtensionItem */



