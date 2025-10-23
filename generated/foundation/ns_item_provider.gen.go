// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
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
	CanLoadObjectOfClass(aClass unsafe.Pointer) bool
	HasItemConformingToTypeIdentifier(typeIdentifier string) bool
	HasRepresentationConformingToTypeIdentifierFileOptions(typeIdentifier string, fileOptions NSItemProviderFileOptions) bool
	LoadDataRepresentationForTypeIdentifierCompletionHandler(typeIdentifier string, completionHandler unsafe.Pointer) Progress
	LoadDataRepresentationForContentTypeCompletionHandler(contentType objectivec.IObject, completionHandler unsafe.Pointer) Progress
	LoadFileRepresentationForTypeIdentifierCompletionHandler(typeIdentifier string, completionHandler unsafe.Pointer) Progress
	LoadFileRepresentationForContentTypeOpenInPlaceCompletionHandler(contentType objectivec.IObject, openInPlace bool, completionHandler unsafe.Pointer) Progress
	LoadInPlaceFileRepresentationForTypeIdentifierCompletionHandler(typeIdentifier string, completionHandler unsafe.Pointer) Progress
	LoadItemForTypeIdentifierOptionsCompletionHandler(typeIdentifier string, options objectivec.IObject, completionHandler unsafe.Pointer)
	LoadObjectOfClassCompletionHandler(aClass unsafe.Pointer, completionHandler unsafe.Pointer) Progress
	LoadPreviewImageWithOptionsCompletionHandler(options objectivec.IObject, completionHandler unsafe.Pointer)
	RegisterCKShareContainerAllowedSharingOptions(share objectivec.IObject, container objectivec.IObject, allowedOptions objectivec.IObject)
	RegisterCKShareWithContainerAllowedSharingOptionsPreparationHandler(container objectivec.IObject, allowedOptions objectivec.IObject, preparationHandler unsafe.Pointer)
	RegisterCloudKitShareContainer(share objectivec.IObject, container objectivec.IObject)
	RegisterCloudKitShareWithPreparationHandler(preparationHandler unsafe.Pointer)
	RegisterDataRepresentationForTypeIdentifierVisibilityLoadHandler(typeIdentifier string, visibility NSItemProviderRepresentationVisibility, loadHandler unsafe.Pointer)
	RegisterDataRepresentationForContentTypeVisibilityLoadHandler(contentType objectivec.IObject, visibility NSItemProviderRepresentationVisibility, loadHandler unsafe.Pointer)
	RegisterFileRepresentationForTypeIdentifierFileOptionsVisibilityLoadHandler(typeIdentifier string, fileOptions NSItemProviderFileOptions, visibility NSItemProviderRepresentationVisibility, loadHandler unsafe.Pointer)
	RegisterFileRepresentationForContentTypeVisibilityOpenInPlaceLoadHandler(contentType objectivec.IObject, visibility NSItemProviderRepresentationVisibility, openInPlace bool, loadHandler unsafe.Pointer)
	RegisterItemForTypeIdentifierLoadHandler(typeIdentifier string, loadHandler unsafe.Pointer)
	RegisterObjectVisibility(object objectivec.IObject, visibility NSItemProviderRepresentationVisibility)
	RegisterObjectOfClassVisibilityLoadHandler(aClass unsafe.Pointer, visibility NSItemProviderRepresentationVisibility, loadHandler unsafe.Pointer)
	RegisteredContentTypesConformingToContentType(contentType objectivec.IObject) []objectivec.IObject
	RegisteredTypeIdentifiersWithFileOptions(fileOptions NSItemProviderFileOptions) []string
	ContainerFrame() Rect
	PreferredPresentationSize() coregraphics.CGSize
	SetPreferredPresentationSize(value coregraphics.CGSize)
	PreferredPresentationStyle() UIPreferredPresentationStyle
	SetPreferredPresentationStyle(value UIPreferredPresentationStyle)
	PreviewImageHandler() unsafe.Pointer
	SetPreviewImageHandler(value unsafe.Pointer)
	RegisteredContentTypes() []objectivec.IObject
	RegisteredContentTypesForOpenInPlace() []objectivec.IObject
	RegisteredTypeIdentifiers() []string
	SourceFrame() Rect
	SuggestedName() string
	SetSuggestedName(value string)
	TeamData() IData
	SetTeamData(value IData)
	Attachments() IItemProvider
	SetAttachments(value IItemProvider)
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



// Provides data-backed content from an existing file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/init(contentsOf:)
func NewItemProviderWithContentsOfURL(fileURL IURL) ItemProvider {
	instance := getItemProviderClass().Alloc()
	rv := objc.Send[ItemProvider](instance.ID, objc.Sel("initWithContentsOfURL:"), fileURL)
	rv.Autorelease()
	return rv
}


// Provides data-backed content from an existing file with the specified parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/initWithContentsOfURL:contentType:openInPlace:coordinated:visibility:
func NewItemProviderWithContentsOfURLContentTypeOpenInPlaceCoordinatedVisibility(fileURL IURL, contentType objectivec.IObject, openInPlace bool, coordinated bool, visibility NSItemProviderRepresentationVisibility) ItemProvider {
	instance := getItemProviderClass().Alloc()
	rv := objc.Send[ItemProvider](instance.ID, objc.Sel("initWithContentsOfURL:contentType:openInPlace:coordinated:visibility:"), fileURL, contentType, openInPlace, coordinated, visibility)
	rv.Autorelease()
	return rv
}


// Creates an item provider with an object, according to the item provider type coercion policy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/init(item:typeIdentifier:)
func NewItemProviderWithItemTypeIdentifier(item objectivec.IObject, typeIdentifier string) ItemProvider {
	instance := getItemProviderClass().Alloc()
	rv := objc.Send[ItemProvider](instance.ID, objc.Sel("initWithItem:typeIdentifier:"), item, objc.String(typeIdentifier))
	rv.Autorelease()
	return rv
}


// Creates a new item provider, employing a specified object’s type identifiers to specify the data representations eligible for the provider to load.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/init(object:)
func NewItemProviderWithObject(object objectivec.IObject) ItemProvider {
	instance := getItemProviderClass().Alloc()
	rv := objc.Send[ItemProvider](instance.ID, objc.Sel("initWithObject:"), object)
	rv.Autorelease()
	return rv
}



// Returns a Boolean value indicating whether an item provider can load objects of a specified class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/canLoadObject(ofClass:)-3eig9
func (i_ ItemProvider) CanLoadObjectOfClass(aClass unsafe.Pointer) bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("canLoadObjectOfClass:"), aClass)
	return rv
}


// Returns a Boolean value indicating whether an item provider contains a data representation conforming to a specified universal type identifier file options parameter with a value of zero.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/hasItemConformingToTypeIdentifier(_:)
func (i_ ItemProvider) HasItemConformingToTypeIdentifier(typeIdentifier string) bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("hasItemConformingToTypeIdentifier:"), objc.String(typeIdentifier))
	return rv
}


// Returns a Boolean value indicating whether an item provider contains a data representation conforming to a specified universal type identifier and to specified open-in-place behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/hasRepresentationConforming(toTypeIdentifier:fileOptions:)
func (i_ ItemProvider) HasRepresentationConformingToTypeIdentifierFileOptions(typeIdentifier string, fileOptions NSItemProviderFileOptions) bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("hasRepresentationConformingToTypeIdentifier:fileOptions:"), objc.String(typeIdentifier), fileOptions)
	return rv
}


// Asynchronously copies the provided, typed data into a generic data object, returning a progress object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/loadDataRepresentation(forTypeIdentifier:completionHandler:)
func (i_ ItemProvider) LoadDataRepresentationForTypeIdentifierCompletionHandler(typeIdentifier string, completionHandler unsafe.Pointer) Progress {
	rv := objc.Send[Progress](i_.ID, objc.Sel("loadDataRepresentationForTypeIdentifier:completionHandler:"), objc.String(typeIdentifier), completionHandler)
	return rv
}


// Asynchronously copies the provided, typed data into a generic data object, returning a progress object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/loadDataRepresentationForContentType:completionHandler:
func (i_ ItemProvider) LoadDataRepresentationForContentTypeCompletionHandler(contentType objectivec.IObject, completionHandler unsafe.Pointer) Progress {
	rv := objc.Send[Progress](i_.ID, objc.Sel("loadDataRepresentationForContentType:completionHandler:"), contentType, completionHandler)
	return rv
}


// Asynchronously writes a copy of the provided, typed data to a temporary file, returning a progress object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/loadFileRepresentation(forTypeIdentifier:completionHandler:)
func (i_ ItemProvider) LoadFileRepresentationForTypeIdentifierCompletionHandler(typeIdentifier string, completionHandler unsafe.Pointer) Progress {
	rv := objc.Send[Progress](i_.ID, objc.Sel("loadFileRepresentationForTypeIdentifier:completionHandler:"), objc.String(typeIdentifier), completionHandler)
	return rv
}


// Asynchronously copies the content type data into a generic data object with the specified parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/loadFileRepresentationForContentType:openInPlace:completionHandler:
func (i_ ItemProvider) LoadFileRepresentationForContentTypeOpenInPlaceCompletionHandler(contentType objectivec.IObject, openInPlace bool, completionHandler unsafe.Pointer) Progress {
	rv := objc.Send[Progress](i_.ID, objc.Sel("loadFileRepresentationForContentType:openInPlace:completionHandler:"), contentType, openInPlace, completionHandler)
	return rv
}


// Asynchronously opens a file in place, if possible, returning a progress object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/loadInPlaceFileRepresentation(forTypeIdentifier:completionHandler:)
func (i_ ItemProvider) LoadInPlaceFileRepresentationForTypeIdentifierCompletionHandler(typeIdentifier string, completionHandler unsafe.Pointer) Progress {
	rv := objc.Send[Progress](i_.ID, objc.Sel("loadInPlaceFileRepresentationForTypeIdentifier:completionHandler:"), objc.String(typeIdentifier), completionHandler)
	return rv
}


// Loads the item’s data and coerces it to the specified type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/loadItem(forTypeIdentifier:options:completionHandler:)
func (i_ ItemProvider) LoadItemForTypeIdentifierOptionsCompletionHandler(typeIdentifier string, options objectivec.IObject, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("loadItemForTypeIdentifier:options:completionHandler:"), objc.String(typeIdentifier), options, completionHandler)
}


// Asynchronously loads an object of a specified class to an item provider, returning a progress object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/loadObject(ofClass:completionHandler:)-8ak5d
func (i_ ItemProvider) LoadObjectOfClassCompletionHandler(aClass unsafe.Pointer, completionHandler unsafe.Pointer) Progress {
	rv := objc.Send[Progress](i_.ID, objc.Sel("loadObjectOfClass:completionHandler:"), aClass, completionHandler)
	return rv
}


// Loads the preview image for the item that the item provider represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/loadPreviewImage(options:completionHandler:)
func (i_ ItemProvider) LoadPreviewImageWithOptionsCompletionHandler(options objectivec.IObject, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("loadPreviewImageWithOptions:completionHandler:"), options, completionHandler)
}


// Registers an existing collaboration object on a server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/registerCKShare:container:allowedSharingOptions:
func (i_ ItemProvider) RegisterCKShareContainerAllowedSharingOptions(share objectivec.IObject, container objectivec.IObject, allowedOptions objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("registerCKShare:container:allowedSharingOptions:"), share, container, allowedOptions)
}


// Creates and registers a new collaboration object using a collection of records to share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/registerCKShareWithContainer:allowedSharingOptions:preparationHandler:
func (i_ ItemProvider) RegisterCKShareWithContainerAllowedSharingOptionsPreparationHandler(container objectivec.IObject, allowedOptions objectivec.IObject, preparationHandler unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("registerCKShareWithContainer:allowedSharingOptions:preparationHandler:"), container, allowedOptions, preparationHandler)
}


// Registers a CloudKit share for the user to modify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/registerCloudKitShare(_:container:)
func (i_ ItemProvider) RegisterCloudKitShareContainer(share objectivec.IObject, container objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("registerCloudKitShare:container:"), share, container)
}


// Registers a handler that prepares a new CloudKit share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/registerCloudKitShare(preparationHandler:)
func (i_ ItemProvider) RegisterCloudKitShareWithPreparationHandler(preparationHandler unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("registerCloudKitShareWithPreparationHandler:"), preparationHandler)
}


// Registers a data-backed representation for an item, specifiying item visibility and a load handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/registerDataRepresentation(forTypeIdentifier:visibility:loadHandler:)
func (i_ ItemProvider) RegisterDataRepresentationForTypeIdentifierVisibilityLoadHandler(typeIdentifier string, visibility NSItemProviderRepresentationVisibility, loadHandler unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("registerDataRepresentationForTypeIdentifier:visibility:loadHandler:"), objc.String(typeIdentifier), visibility, loadHandler)
}


// Lazily registers an item, according to the item provider type coercion policy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/registerDataRepresentationForContentType:visibility:loadHandler:
func (i_ ItemProvider) RegisterDataRepresentationForContentTypeVisibilityLoadHandler(contentType objectivec.IObject, visibility NSItemProviderRepresentationVisibility, loadHandler unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("registerDataRepresentationForContentType:visibility:loadHandler:"), contentType, visibility, loadHandler)
}


// Registers a file-backed representation for an item, specifying file options, item visibility, and a load handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/registerFileRepresentation(forTypeIdentifier:fileOptions:visibility:loadHandler:)
func (i_ ItemProvider) RegisterFileRepresentationForTypeIdentifierFileOptionsVisibilityLoadHandler(typeIdentifier string, fileOptions NSItemProviderFileOptions, visibility NSItemProviderRepresentationVisibility, loadHandler unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("registerFileRepresentationForTypeIdentifier:fileOptions:visibility:loadHandler:"), objc.String(typeIdentifier), fileOptions, visibility, loadHandler)
}


// Registers a file-backed representation for an item with item visibility, an open-in-place option, and a load handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/registerFileRepresentationForContentType:visibility:openInPlace:loadHandler:
func (i_ ItemProvider) RegisterFileRepresentationForContentTypeVisibilityOpenInPlaceLoadHandler(contentType objectivec.IObject, visibility NSItemProviderRepresentationVisibility, openInPlace bool, loadHandler unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("registerFileRepresentationForContentType:visibility:openInPlace:loadHandler:"), contentType, visibility, openInPlace, loadHandler)
}


// Lazily registers an item, according to the item provider type coercion policy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/registerItem(forTypeIdentifier:loadHandler:)
func (i_ ItemProvider) RegisterItemForTypeIdentifierLoadHandler(typeIdentifier string, loadHandler unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("registerItemForTypeIdentifier:loadHandler:"), objc.String(typeIdentifier), loadHandler)
}


// Adds representations of a specified object to an item provider, based on the object’s implementation of the item provider writing protocol, and adhering to a visibility specification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/registerObject(_:visibility:)
func (i_ ItemProvider) RegisterObjectVisibility(object objectivec.IObject, visibility NSItemProviderRepresentationVisibility) {
	objc.Send[objc.ID](i_.ID, objc.Sel("registerObject:visibility:"), object, visibility)
}


// Lazily adds representations of a specified object class to an item provider, based on the object’s implementation of the item provider writing protocol, and adhering to a visibility specification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/registerObject(ofClass:visibility:loadHandler:)-9sndn
func (i_ ItemProvider) RegisterObjectOfClassVisibilityLoadHandler(aClass unsafe.Pointer, visibility NSItemProviderRepresentationVisibility, loadHandler unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("registerObjectOfClass:visibility:loadHandler:"), aClass, visibility, loadHandler)
}


// Returns an array of registered content types that conform to a specified content type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/registeredContentTypes(conformingTo:)
func (i_ ItemProvider) RegisteredContentTypesConformingToContentType(contentType objectivec.IObject) []objectivec.IObject {
	rv := objc.Send[[]objectivec.IObject](i_.ID, objc.Sel("registeredContentTypesConformingToContentType:"), contentType)
	return rv
}


// Returns an array with a subset of type identifiers for the item provider, according to the specified file options, in the same order they were registered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/registeredTypeIdentifiers(fileOptions:)
func (i_ ItemProvider) RegisteredTypeIdentifiersWithFileOptions(fileOptions NSItemProviderFileOptions) []string {
	rv := objc.Send[[]string](i_.ID, objc.Sel("registeredTypeIdentifiersWithFileOptions:"), fileOptions)
	return rv
}


// The rectangle of the item’s visible content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/containerFrame
func (i_ ItemProvider) ContainerFrame() Rect {
	rv := objc.Send[Rect](i_.ID, objc.Sel("containerFrame"))
	return rv
}


// The ideal presentation size of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/preferredPresentationSize
func (i_ ItemProvider) PreferredPresentationSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](i_.ID, objc.Sel("preferredPresentationSize"))
	return rv
}


// The ideal presentation size of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/preferredPresentationSize
func (i_ ItemProvider) SetPreferredPresentationSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPreferredPresentationSize:"), value)
}


// The preferred style for presenting the item provider’s data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/preferredPresentationStyle-swift.property
func (i_ ItemProvider) PreferredPresentationStyle() UIPreferredPresentationStyle {
	rv := objc.Send[PreferredPresentationStyle](i_.ID, objc.Sel("preferredPresentationStyle"))
	return rv
}


// The preferred style for presenting the item provider’s data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/preferredPresentationStyle-swift.property
func (i_ ItemProvider) SetPreferredPresentationStyle(value UIPreferredPresentationStyle) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPreferredPresentationStyle:"), value)
}


// The custom preview image handler block for the item provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/previewImageHandler
func (i_ ItemProvider) PreviewImageHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("previewImageHandler"))
	return rv
}


// The custom preview image handler block for the item provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/previewImageHandler
func (i_ ItemProvider) SetPreviewImageHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPreviewImageHandler:"), value)
}


// Registered content types in the order the app registers each type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/registeredContentTypes
func (i_ ItemProvider) RegisteredContentTypes() []objectivec.IObject {
	rv := objc.Send[[]objectivec.IObject](i_.ID, objc.Sel("registeredContentTypes"))
	return rv
}


// Registered content types that the system can load as open-in-place files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/registeredContentTypesForOpenInPlace
func (i_ ItemProvider) RegisteredContentTypesForOpenInPlace() []objectivec.IObject {
	rv := objc.Send[[]objectivec.IObject](i_.ID, objc.Sel("registeredContentTypesForOpenInPlace"))
	return rv
}


// Returns the array of type identifiers for the item provider, in the same order they were registered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/registeredTypeIdentifiers
func (i_ ItemProvider) RegisteredTypeIdentifiers() []string {
	rv := objc.Send[[]string](i_.ID, objc.Sel("registeredTypeIdentifiers"))
	return rv
}


// The rectangle that the item occupies in the host app’s source window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/sourceFrame
func (i_ ItemProvider) SourceFrame() Rect {
	rv := objc.Send[Rect](i_.ID, objc.Sel("sourceFrame"))
	return rv
}


// The filename to use when writing the provided data to a file on disk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/suggestedName
func (i_ ItemProvider) SuggestedName() string {
	rv := objc.Send[string](i_.ID, objc.Sel("suggestedName"))
	return rv
}


// The filename to use when writing the provided data to a file on disk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/suggestedName
func (i_ ItemProvider) SetSuggestedName(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSuggestedName:"), objc.String(value))
}


// The collection of data an app uses to hold private team information during drag and drop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/teamData
func (i_ ItemProvider) TeamData() IData {
	rv := objc.Send[NSData](i_.ID, objc.Sel("teamData"))
	return rv
}


// The collection of data an app uses to hold private team information during drag and drop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/teamData
func (i_ ItemProvider) SetTeamData(value IData) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTeamData:"), value)
}


// An optional array of media data associated with the extension item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensionitem/attachments
func (i_ ItemProvider) Attachments() IItemProvider {
	rv := objc.Send[NSItemProvider](i_.ID, objc.Sel("attachments"))
	return rv
}


// An optional array of media data associated with the extension item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensionitem/attachments
func (i_ ItemProvider) SetAttachments(value IItemProvider) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAttachments:"), value)
}


