// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ItemProvider] class.
var (
	ItemProviderClass     _ItemProviderClass
	ItemProviderClassOnce sync.Once
)

func getItemProviderClass() _ItemProviderClass {
	ItemProviderClassOnce.Do(func() {
		ItemProviderClass = _ItemProviderClass{objc.GetClass("NSItemProvider")}
	})
	return ItemProviderClass
}

type _ItemProviderClass struct {
	class objc.Class
}

// An interface definition for the [ItemProvider] class.
type IItemProvider interface {
	objectivec.IObject
	// properties:
	Attachments() IItemProvider
	SetAttachments(value IItemProvider)
	ContainerFrame() objc.IObject /* cross-framework: Rect */
	SetContainerFrame(value objc.IObject /* cross-framework: Rect */)
	PreferredPresentationSize() objc.IObject /* cross-framework: Size */
	SetPreferredPresentationSize(value objc.IObject /* cross-framework: Size */)
	PreferredPresentationStyle() unsafe.Pointer
	SetPreferredPresentationStyle(value unsafe.Pointer)
	PreviewImageHandler() unsafe.Pointer
	SetPreviewImageHandler(value unsafe.Pointer)
	RegisteredContentTypes() objc.IObject
	SetRegisteredContentTypes(value objc.IObject)
	RegisteredContentTypesForOpenInPlace() objc.IObject
	SetRegisteredContentTypesForOpenInPlace(value objc.IObject)
	RegisteredTypeIdentifiers() IString
	SetRegisteredTypeIdentifiers(value IString)
	SourceFrame() objc.IObject /* cross-framework: Rect */
	SetSourceFrame(value objc.IObject /* cross-framework: Rect */)
	SuggestedName() IString
	SetSuggestedName(value IString)
	TeamData() IData
	SetTeamData(value IData)
	// methods:
}

// An item provider for conveying data or a file between processes during drag-and-drop or copy-and-paste activities, or from a host app to an app extension.
//
// Starting in iOS 11, item providers play a central role in drag and drop, and in copy and paste. They continue to play a role with app extensions. The system uses an internal queue when calling the completion blocks for the class. When using an item provider with drag and drop, ensure that UI updates take place on the main queue as follows:


// An item provider for conveying data or a file between processes during drag-and-drop or copy-and-paste activities, or from a host app to an app extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider
type ItemProvider struct {
	objectivec.Object
}

// ItemProviderFrom constructs a [ItemProvider] from an unsafe.Pointer.
//
// An item provider for conveying data or a file between processes during drag-and-drop or copy-and-paste activities, or from a host app to an app extension.
func ItemProviderFrom(ptr unsafe.Pointer) ItemProvider {
	return ItemProvider{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _ItemProviderClass) Alloc() ItemProvider {
	rv := objc.Send[ItemProvider](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ItemProviderClass) New() ItemProvider {
	rv := objc.Send[ItemProvider](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ItemProvider) Init() ItemProvider {
	rv := objc.Send[ItemProvider](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ItemProvider) Autorelease() ItemProvider {
	rv := objc.Send[ItemProvider](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewItemProvider creates a new ItemProvider instance.
func NewItemProvider() ItemProvider {
	return getItemProviderClass().New()
}



// An optional array of media data associated with the extension item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensionitem/attachments
func (i_ ItemProvider) Attachments() IItemProvider {
	rv := objc.Send[ItemProvider](i_.ID, objc.Sel("attachments"))
	return rv
}


// An optional array of media data associated with the extension item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensionitem/attachments
func (i_ ItemProvider) SetAttachments(value IItemProvider) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAttachments:"), value)
}


// The rectangle of the item’s visible content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/containerframe
func (i_ ItemProvider) ContainerFrame() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("containerFrame"))
	return rv
}


// The rectangle of the item’s visible content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/containerframe
func (i_ ItemProvider) SetContainerFrame(value objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setContainerFrame:"), value)
}


// The ideal presentation size of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/preferredpresentationsize
func (i_ ItemProvider) PreferredPresentationSize() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("preferredPresentationSize"))
	return rv
}


// The ideal presentation size of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/preferredpresentationsize
func (i_ ItemProvider) SetPreferredPresentationSize(value objc.IObject /* cross-framework: Size */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPreferredPresentationSize:"), value)
}


// The preferred style for presenting the item provider’s data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/preferredpresentationstyle-swift.property
func (i_ ItemProvider) PreferredPresentationStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("preferredPresentationStyle"))
	return rv
}


// The preferred style for presenting the item provider’s data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/preferredpresentationstyle-swift.property
func (i_ ItemProvider) SetPreferredPresentationStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPreferredPresentationStyle:"), value)
}


// The custom preview image handler block for the item provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/previewimagehandler
func (i_ ItemProvider) PreviewImageHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("previewImageHandler"))
	return rv
}


// The custom preview image handler block for the item provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/previewimagehandler
func (i_ ItemProvider) SetPreviewImageHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPreviewImageHandler:"), value)
}


// Registered content types in the order the app registers each type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/registeredcontenttypes
func (i_ ItemProvider) RegisteredContentTypes() objc.IObject {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("registeredContentTypes"))
	return rv
}


// Registered content types in the order the app registers each type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/registeredcontenttypes
func (i_ ItemProvider) SetRegisteredContentTypes(value objc.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRegisteredContentTypes:"), value)
}


// Registered content types that the system can load as open-in-place files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/registeredcontenttypesforopeninplace
func (i_ ItemProvider) RegisteredContentTypesForOpenInPlace() objc.IObject {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("registeredContentTypesForOpenInPlace"))
	return rv
}


// Registered content types that the system can load as open-in-place files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/registeredcontenttypesforopeninplace
func (i_ ItemProvider) SetRegisteredContentTypesForOpenInPlace(value objc.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRegisteredContentTypesForOpenInPlace:"), value)
}


// Returns the array of type identifiers for the item provider, in the same order they were registered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/registeredtypeidentifiers
func (i_ ItemProvider) RegisteredTypeIdentifiers() IString {
	rv := objc.Send[String](i_.ID, objc.Sel("registeredTypeIdentifiers"))
	return rv
}


// Returns the array of type identifiers for the item provider, in the same order they were registered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/registeredtypeidentifiers
func (i_ ItemProvider) SetRegisteredTypeIdentifiers(value IString) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRegisteredTypeIdentifiers:"), value)
}


// The rectangle that the item occupies in the host app’s source window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/sourceframe
func (i_ ItemProvider) SourceFrame() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("sourceFrame"))
	return rv
}


// The rectangle that the item occupies in the host app’s source window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/sourceframe
func (i_ ItemProvider) SetSourceFrame(value objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSourceFrame:"), value)
}


// The filename to use when writing the provided data to a file on disk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/suggestedname
func (i_ ItemProvider) SuggestedName() IString {
	rv := objc.Send[String](i_.ID, objc.Sel("suggestedName"))
	return rv
}


// The filename to use when writing the provided data to a file on disk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/suggestedname
func (i_ ItemProvider) SetSuggestedName(value IString) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSuggestedName:"), value)
}


// The collection of data an app uses to hold private team information during drag and drop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/teamdata
func (i_ ItemProvider) TeamData() IData {
	rv := objc.Send[Data](i_.ID, objc.Sel("teamData"))
	return rv
}


// The collection of data an app uses to hold private team information during drag and drop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/teamdata
func (i_ ItemProvider) SetTeamData(value IData) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTeamData:"), value)
}



