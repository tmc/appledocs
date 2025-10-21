// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CSSearchableItem] class.
var (
	CSSearchableItemClass     _CSSearchableItemClass
	CSSearchableItemClassOnce sync.Once
)

func getCSSearchableItemClass() _CSSearchableItemClass {
	CSSearchableItemClassOnce.Do(func() {
		CSSearchableItemClass = _CSSearchableItemClass{objc.GetClass("CSSearchableItem")}
	})
	return CSSearchableItemClass
}

type _CSSearchableItemClass struct {
	class objc.Class
}

// An interface definition for the [CSSearchableItem] class.
type ICSSearchableItem interface {
	objectivec.IObject
	CompareByRank(other unsafe.Pointer) unsafe.Pointer
}

// The details of your app-specific content that someone might search for on their devices.
//
// A uniquely identifies a part of your app’s content, and provides the metadata that Spotlight indexes and uses to find that content later. As part of indexing your app’s content, you create searchable items and fill them with details about your app’s content and where to find it. After indexing the content, you can then execute queries using the Core Spotlight APIs to find the items you indexed. People can also use the system’s Spotlight search interface to find your app’s content. When you create or update content in your app, create a for that content if you want it to be searchable. A searchable item contains identification strings you use to locate that item in your content and a object with details about the item. For the metadata, you typically want to provide values for the , , and attributes at a minimum. If you’re indexing a file on disk, provide a value for the attribute. Fill in as many other attributes as makes sense for the content you’re indexing. After creating a searchable item, index it using a object. As you update your app’s content, update your objects for that content and index them right away. If you delete content, similarly delete the searchable items from the index. Keeping your app’s indexes current ensures that searches return valid information. For more information on indexing your content, see .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem
type CSSearchableItem struct {
	objectivec.Object
}

// CSSearchableItemFrom constructs a [CSSearchableItem] from an unsafe.Pointer.
//
// The details of your app-specific content that someone might search for on their devices.
func CSSearchableItemFrom(ptr unsafe.Pointer) CSSearchableItem {
	return CSSearchableItem{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CSSearchableItemClass) Alloc() CSSearchableItem {
	rv := objc.Send[CSSearchableItem](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CSSearchableItemClass) New() CSSearchableItem {
	rv := objc.Send[CSSearchableItem](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CSSearchableItem) Init() CSSearchableItem {
	rv := objc.Send[CSSearchableItem](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CSSearchableItem) Autorelease() CSSearchableItem {
	rv := objc.Send[CSSearchableItem](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCSSearchableItem creates a new CSSearchableItem instance.
func NewCSSearchableItem() CSSearchableItem {
	return getCSSearchableItemClass().New()
}


// Returns a searchable item associated with the specified identifier, domain identifier, and attribute set.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/init(uniqueIdentifier:domainIdentifier:attributeSet:)
func NewCSSearchableItemWithUniqueIdentifierDomainIdentifierAttributeSet(uniqueIdentifier string, domainIdentifier string, attributeSet unsafe.Pointer) CSSearchableItem {
	instance := getCSSearchableItemClass().Alloc()
	rv := objc.Send[CSSearchableItem](instance.ID, objc.Sel("initWithUniqueIdentifier:domainIdentifier:attributeSet:"), objc.String(uniqueIdentifier), objc.String(domainIdentifier), attributeSet)
	rv.Autorelease()
	return rv
}


// Compares two items by rank and returns the result.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/compare(byRank:)
func (c_ CSSearchableItem) CompareByRank(other unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("compareByRank:"), other)
	return rv
}

// The set of attributes that contain metadata associated with the item in a object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/attributeSet
func (c_ CSSearchableItem) AttributeSet() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("attributeSet"))
	return rv
}


// SetAttributeSet sets the value of the attributeSet property.
// The set of attributes that contain metadata associated with the item in a object.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/attributeSet
func (c_ CSSearchableItem) SetAttributeSet(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAttributeSet:"), value)
}
// An optional identifier that represents the domain or owner of the item.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/domainIdentifier
func (c_ CSSearchableItem) DomainIdentifier() string {
	rv := objc.Send[string](c_.ID, objc.Sel("domainIdentifier"))
	return rv
}


// SetDomainIdentifier sets the value of the domainIdentifier property.
// An optional identifier that represents the domain or owner of the item.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/domainIdentifier
func (c_ CSSearchableItem) SetDomainIdentifier(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDomainIdentifier:"), objc.String(value))
}
// The date after which the searchable item should no longer exist.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/expirationDate
func (c_ CSSearchableItem) ExpirationDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("expirationDate"))
	return rv
}


// SetExpirationDate sets the value of the expirationDate property.
// The date after which the searchable item should no longer exist.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/expirationDate
func (c_ CSSearchableItem) SetExpirationDate(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExpirationDate:"), value)
}
// A Boolean value that indicates whether to treat the item as an update instead of a new item.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/isUpdate
func (c_ CSSearchableItem) IsUpdate() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isUpdate"))
	return rv
}


// SetIsUpdate sets the value of the isUpdate property.
// A Boolean value that indicates whether to treat the item as an update instead of a new item.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/isUpdate
func (c_ CSSearchableItem) SetIsUpdate(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsUpdate:"), value)
}
// The value that uniquely identifies the searchable item within your app.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/uniqueIdentifier
func (c_ CSSearchableItem) UniqueIdentifier() string {
	rv := objc.Send[string](c_.ID, objc.Sel("uniqueIdentifier"))
	return rv
}


// SetUniqueIdentifier sets the value of the uniqueIdentifier property.
// The value that uniquely identifies the searchable item within your app.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/uniqueIdentifier
func (c_ CSSearchableItem) SetUniqueIdentifier(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUniqueIdentifier:"), objc.String(value))
}
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/updateListenerOptions-swift.property
func (c_ CSSearchableItem) UpdateListenerOptions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("updateListenerOptions"))
	return rv
}


// SetUpdateListenerOptions sets the value of the updateListenerOptions property.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/updateListenerOptions-swift.property
func (c_ CSSearchableItem) SetUpdateListenerOptions(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUpdateListenerOptions:"), value)
}

