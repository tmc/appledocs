// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	AttributeSet() ICSSearchableItemAttributeSet
	SetAttributeSet(value ICSSearchableItemAttributeSet)
	DomainIdentifier() string
	SetDomainIdentifier(value string)
	ExpirationDate() foundation.NSDate
	SetExpirationDate(value foundation.NSDate)
	IsUpdate() bool
	SetIsUpdate(value bool)
	UniqueIdentifier() string
	SetUniqueIdentifier(value string)
	UpdateListenerOptions() CSSearchableItemUpdateListenerOptions
	SetUpdateListenerOptions(value CSSearchableItemUpdateListenerOptions)
	CSQueryContinuationActionType() string
	CSSearchQueryString() string
	CSSearchableItemActionType() string
	CSSearchableItemActivityIdentifier() string
	ContentType() string
	SetContentType(value string)
	ContentURL() foundation.URL
	SetContentURL(value foundation.URL)
	DisplayName() string
	SetDisplayName(value string)
	Title() string
	SetTitle(value string)
	CompareByRank(other ICSSearchableItem) unsafe.Pointer
}

// The details of your app-specific content that someone might search for on their devices.
//
// A uniquely identifies a part of your app’s content, and provides the metadata that Spotlight indexes and uses to find that content later. As part of indexing your app’s content, you create searchable items and fill them with details about your app’s content and where to find it. After indexing the content, you can then execute queries using the Core Spotlight APIs to find the items you indexed. People can also use the system’s Spotlight search interface to find your app’s content. When you create or update content in your app, create a for that content if you want it to be searchable. A searchable item contains identification strings you use to locate that item in your content and a object with details about the item. For the metadata, you typically want to provide values for the , , and attributes at a minimum. If you’re indexing a file on disk, provide a value for the attribute. Fill in as many other attributes as makes sense for the content you’re indexing. After creating a searchable item, index it using a object. As you update your app’s content, update your objects for that content and index them right away. If you delete content, similarly delete the searchable items from the index. Keeping your app’s indexes current ensures that searches return valid information. For more information on indexing your content, see .


// The details of your app-specific content that someone might search for on their devices.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/init(uniqueIdentifier:domainIdentifier:attributeSet:)
func NewCSSearchableItemWithUniqueIdentifierDomainIdentifierAttributeSet(uniqueIdentifier string, domainIdentifier string, attributeSet ICSSearchableItemAttributeSet) CSSearchableItem {
	instance := getCSSearchableItemClass().Alloc()
	rv := objc.Send[CSSearchableItem](instance.ID, objc.Sel("initWithUniqueIdentifier:domainIdentifier:attributeSet:"), objc.String(uniqueIdentifier), objc.String(domainIdentifier), attributeSet)
	rv.Autorelease()
	return rv
}



// Compares two items by rank and returns the result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/compare(byRank:)
func (c_ CSSearchableItem) CompareByRank(other ICSSearchableItem) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("compareByRank:"), other)
	return rv
}


// The set of attributes that contain metadata associated with the item in a object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/attributeSet
func (c_ CSSearchableItem) AttributeSet() ICSSearchableItemAttributeSet {
	rv := objc.Send[CSSearchableItemAttributeSet](c_.ID, objc.Sel("attributeSet"))
	return rv
}


// The set of attributes that contain metadata associated with the item in a object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/attributeSet
func (c_ CSSearchableItem) SetAttributeSet(value ICSSearchableItemAttributeSet) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAttributeSet:"), value)
}


// An optional identifier that represents the domain or owner of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/domainIdentifier
func (c_ CSSearchableItem) DomainIdentifier() string {
	rv := objc.Send[string](c_.ID, objc.Sel("domainIdentifier"))
	return rv
}


// An optional identifier that represents the domain or owner of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/domainIdentifier
func (c_ CSSearchableItem) SetDomainIdentifier(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDomainIdentifier:"), objc.String(value))
}


// The date after which the searchable item should no longer exist.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/expirationDate
func (c_ CSSearchableItem) ExpirationDate() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](c_.ID, objc.Sel("expirationDate"))
	return rv
}


// The date after which the searchable item should no longer exist.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/expirationDate
func (c_ CSSearchableItem) SetExpirationDate(value foundation.NSDate) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExpirationDate:"), value)
}


// A Boolean value that indicates whether to treat the item as an update instead of a new item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/isUpdate
func (c_ CSSearchableItem) IsUpdate() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isUpdate"))
	return rv
}


// A Boolean value that indicates whether to treat the item as an update instead of a new item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/isUpdate
func (c_ CSSearchableItem) SetIsUpdate(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsUpdate:"), value)
}


// The value that uniquely identifies the searchable item within your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/uniqueIdentifier
func (c_ CSSearchableItem) UniqueIdentifier() string {
	rv := objc.Send[string](c_.ID, objc.Sel("uniqueIdentifier"))
	return rv
}


// The value that uniquely identifies the searchable item within your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/uniqueIdentifier
func (c_ CSSearchableItem) SetUniqueIdentifier(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUniqueIdentifier:"), objc.String(value))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/updateListenerOptions-swift.property
func (c_ CSSearchableItem) UpdateListenerOptions() CSSearchableItemUpdateListenerOptions {
	rv := objc.Send[CSSearchableItemUpdateListenerOptions](c_.ID, objc.Sel("updateListenerOptions"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/updateListenerOptions-swift.property
func (c_ CSSearchableItem) SetUpdateListenerOptions(value CSSearchableItemUpdateListenerOptions) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUpdateListenerOptions:"), value)
}


// Indicates that the activity type to continue is a search or query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/csquerycontinuationactiontype
func (c_ CSSearchableItem) CSQueryContinuationActionType() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CSQueryContinuationActionType"))
	return rv
}


// Provides the key for the current query in the info dictionary of the user activity object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchquerystring
func (c_ CSSearchableItem) CSSearchQueryString() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CSSearchQueryString"))
	return rv
}


// Indicates that the activity type to continue is related to a searchable item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemactiontype
func (c_ CSSearchableItem) CSSearchableItemActionType() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CSSearchableItemActionType"))
	return rv
}


// The key you use to access a searchable item in a user activity object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemactivityidentifier
func (c_ CSSearchableItem) CSSearchableItemActivityIdentifier() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CSSearchableItemActivityIdentifier"))
	return rv
}


// The uniform type identifier (UTI) of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contenttype
func (c_ CSSearchableItem) ContentType() string {
	rv := objc.Send[string](c_.ID, objc.Sel("contentType"))
	return rv
}


// The uniform type identifier (UTI) of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contenttype
func (c_ CSSearchableItem) SetContentType(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentType:"), objc.String(value))
}


// The file URL of the content to index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contenturl
func (c_ CSSearchableItem) ContentURL() foundation.URL {
	rv := objc.Send[foundation.URL](c_.ID, objc.Sel("contentURL"))
	return rv
}


// The file URL of the content to index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contenturl
func (c_ CSSearchableItem) SetContentURL(value foundation.URL) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentURL:"), value)
}


// A localized string that contains the name of the item, suitable to display in the user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/displayname
func (c_ CSSearchableItem) DisplayName() string {
	rv := objc.Send[string](c_.ID, objc.Sel("displayName"))
	return rv
}


// A localized string that contains the name of the item, suitable to display in the user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/displayname
func (c_ CSSearchableItem) SetDisplayName(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDisplayName:"), objc.String(value))
}


// The title of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/title
func (c_ CSSearchableItem) Title() string {
	rv := objc.Send[string](c_.ID, objc.Sel("title"))
	return rv
}


// The title of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/title
func (c_ CSSearchableItem) SetTitle(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTitle:"), objc.String(value))
}


