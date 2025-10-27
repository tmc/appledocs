// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
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
	ContainerFrame() corefoundation.CGRect
	SetContainerFrame(value corefoundation.CGRect)
	PreferredPresentationSize() corefoundation.CGSize
	SetPreferredPresentationSize(value corefoundation.CGSize)
	PreferredPresentationStyle() objectivec.IObject
	SetPreferredPresentationStyle(value objectivec.IObject)
	PreviewImageHandler() objectivec.IObject
	SetPreviewImageHandler(value objectivec.IObject)
	RegisteredContentTypes() objectivec.IObject
	SetRegisteredContentTypes(value objectivec.IObject)
	RegisteredContentTypesForOpenInPlace() objectivec.IObject
	SetRegisteredContentTypesForOpenInPlace(value objectivec.IObject)
	RegisteredTypeIdentifiers() IString
	SetRegisteredTypeIdentifiers(value IString)
	SourceFrame() corefoundation.CGRect
	SetSourceFrame(value corefoundation.CGRect)
	SuggestedName() IString
	SetSuggestedName(value IString)
	TeamData() IData
	SetTeamData(value IData)


	

	// methods:
	LoadDataRepresentationForTypeIdentifierCompletionHandler(typeIdentifier IString, completionHandler unsafe.Pointer) IProgress
	LoadDataRepresentationForContentTypeCompletionHandler(contentType objectivec.IObject, completionHandler unsafe.Pointer) IProgress
	LoadFileRepresentationForTypeIdentifierCompletionHandler(typeIdentifier IString, completionHandler unsafe.Pointer) IProgress
	LoadFileRepresentationForContentTypeOpenInPlaceCompletionHandler(contentType objectivec.IObject, openInPlace bool, completionHandler unsafe.Pointer) IProgress
	LoadInPlaceFileRepresentationForTypeIdentifierCompletionHandler(typeIdentifier IString, completionHandler unsafe.Pointer) IProgress
	LoadItemForTypeIdentifierOptionsCompletionHandler(typeIdentifier IString, options IDictionary, completionHandler ItemProviderCompletionHandler /* not a class type */)
	LoadObjectOfClassCompletionHandler(aClass unsafe.Pointer, completionHandler unsafe.Pointer) IProgress
	RegisterDataRepresentationForTypeIdentifierVisibilityLoadHandler(typeIdentifier IString, visibility ItemProviderRepresentationVisibility, loadHandler unsafe.Pointer)
	RegisterDataRepresentationForContentTypeVisibilityLoadHandler(contentType objectivec.IObject, visibility ItemProviderRepresentationVisibility, loadHandler unsafe.Pointer)
	RegisterItemForTypeIdentifierLoadHandler(typeIdentifier IString, loadHandler ItemProviderLoadHandler /* not a class type */)


}





// Alloc allocates a new instance without initialization.
func (ic _ItemProviderClass) Alloc() ItemProvider {
	rv := objc.Send[ItemProvider](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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




















// Asynchronously copies the provided, typed data into a generic data object, returning a progress object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/loadDataRepresentation(forTypeIdentifier:completionHandler:)
func (i_ ItemProvider) LoadDataRepresentationForTypeIdentifierCompletionHandler(typeIdentifier IString, completionHandler unsafe.Pointer) IProgress {
	rv := objc.Send[Progress](i_.ID, objc.Sel("loadDataRepresentationForTypeIdentifier:completionHandler:"), typeIdentifier, completionHandler)
	return rv
}


// Asynchronously copies the provided, typed data into a generic data object, returning a progress object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/loadDataRepresentationForContentType:completionHandler:
func (i_ ItemProvider) LoadDataRepresentationForContentTypeCompletionHandler(contentType objectivec.IObject, completionHandler unsafe.Pointer) IProgress {
	rv := objc.Send[Progress](i_.ID, objc.Sel("loadDataRepresentationForContentType:completionHandler:"), contentType, completionHandler)
	return rv
}


// Asynchronously writes a copy of the provided, typed data to a temporary file, returning a progress object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/loadFileRepresentation(forTypeIdentifier:completionHandler:)
func (i_ ItemProvider) LoadFileRepresentationForTypeIdentifierCompletionHandler(typeIdentifier IString, completionHandler unsafe.Pointer) IProgress {
	rv := objc.Send[Progress](i_.ID, objc.Sel("loadFileRepresentationForTypeIdentifier:completionHandler:"), typeIdentifier, completionHandler)
	return rv
}


// Asynchronously copies the content type data into a generic data object with the specified parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/loadFileRepresentationForContentType:openInPlace:completionHandler:
func (i_ ItemProvider) LoadFileRepresentationForContentTypeOpenInPlaceCompletionHandler(contentType objectivec.IObject, openInPlace bool, completionHandler unsafe.Pointer) IProgress {
	rv := objc.Send[Progress](i_.ID, objc.Sel("loadFileRepresentationForContentType:openInPlace:completionHandler:"), contentType, openInPlace, completionHandler)
	return rv
}


// Asynchronously opens a file in place, if possible, returning a progress object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/loadInPlaceFileRepresentation(forTypeIdentifier:completionHandler:)
func (i_ ItemProvider) LoadInPlaceFileRepresentationForTypeIdentifierCompletionHandler(typeIdentifier IString, completionHandler unsafe.Pointer) IProgress {
	rv := objc.Send[Progress](i_.ID, objc.Sel("loadInPlaceFileRepresentationForTypeIdentifier:completionHandler:"), typeIdentifier, completionHandler)
	return rv
}


// Loads the item’s data and coerces it to the specified type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/loadItem(forTypeIdentifier:options:completionHandler:)
func (i_ ItemProvider) LoadItemForTypeIdentifierOptionsCompletionHandler(typeIdentifier IString, options IDictionary, completionHandler ItemProviderCompletionHandler /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("loadItemForTypeIdentifier:options:completionHandler:"), typeIdentifier, options, completionHandler)
}


// Asynchronously loads an object of a specified class to an item provider, returning a progress object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/loadObject(ofClass:completionHandler:)-8ak5d
func (i_ ItemProvider) LoadObjectOfClassCompletionHandler(aClass unsafe.Pointer, completionHandler unsafe.Pointer) IProgress {
	rv := objc.Send[Progress](i_.ID, objc.Sel("loadObjectOfClass:completionHandler:"), aClass, completionHandler)
	return rv
}


// Registers a data-backed representation for an item, specifiying item visibility and a load handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/registerDataRepresentation(forTypeIdentifier:visibility:loadHandler:)
func (i_ ItemProvider) RegisterDataRepresentationForTypeIdentifierVisibilityLoadHandler(typeIdentifier IString, visibility ItemProviderRepresentationVisibility, loadHandler unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("registerDataRepresentationForTypeIdentifier:visibility:loadHandler:"), typeIdentifier, visibility, loadHandler)
}


// Lazily registers an item, according to the item provider type coercion policy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/registerDataRepresentationForContentType:visibility:loadHandler:
func (i_ ItemProvider) RegisterDataRepresentationForContentTypeVisibilityLoadHandler(contentType objectivec.IObject, visibility ItemProviderRepresentationVisibility, loadHandler unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("registerDataRepresentationForContentType:visibility:loadHandler:"), contentType, visibility, loadHandler)
}


// Lazily registers an item, according to the item provider type coercion policy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/registerItem(forTypeIdentifier:loadHandler:)
func (i_ ItemProvider) RegisterItemForTypeIdentifierLoadHandler(typeIdentifier IString, loadHandler ItemProviderLoadHandler /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("registerItemForTypeIdentifier:loadHandler:"), typeIdentifier, loadHandler)
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
func (i_ ItemProvider) ContainerFrame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](i_.ID, objc.Sel("containerFrame"))
	return rv
}


// The rectangle of the item’s visible content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/containerframe
func (i_ ItemProvider) SetContainerFrame(value corefoundation.CGRect) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setContainerFrame:"), value)
}


// The ideal presentation size of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/preferredpresentationsize
func (i_ ItemProvider) PreferredPresentationSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](i_.ID, objc.Sel("preferredPresentationSize"))
	return rv
}


// The ideal presentation size of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/preferredpresentationsize
func (i_ ItemProvider) SetPreferredPresentationSize(value corefoundation.CGSize) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPreferredPresentationSize:"), value)
}


// The preferred style for presenting the item provider’s data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/preferredpresentationstyle-swift.property
func (i_ ItemProvider) PreferredPresentationStyle() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("preferredPresentationStyle"))
	return rv
}


// The preferred style for presenting the item provider’s data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/preferredpresentationstyle-swift.property
func (i_ ItemProvider) SetPreferredPresentationStyle(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPreferredPresentationStyle:"), value)
}


// The custom preview image handler block for the item provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/previewimagehandler
func (i_ ItemProvider) PreviewImageHandler() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("previewImageHandler"))
	return rv
}


// The custom preview image handler block for the item provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/previewimagehandler
func (i_ ItemProvider) SetPreviewImageHandler(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPreviewImageHandler:"), value)
}


// Registered content types in the order the app registers each type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/registeredcontenttypes
func (i_ ItemProvider) RegisteredContentTypes() objectivec.IObject {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("registeredContentTypes"))
	return rv
}


// Registered content types in the order the app registers each type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/registeredcontenttypes
func (i_ ItemProvider) SetRegisteredContentTypes(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRegisteredContentTypes:"), value)
}


// Registered content types that the system can load as open-in-place files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/registeredcontenttypesforopeninplace
func (i_ ItemProvider) RegisteredContentTypesForOpenInPlace() objectivec.IObject {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("registeredContentTypesForOpenInPlace"))
	return rv
}


// Registered content types that the system can load as open-in-place files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/registeredcontenttypesforopeninplace
func (i_ ItemProvider) SetRegisteredContentTypesForOpenInPlace(value objectivec.IObject) {
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
func (i_ ItemProvider) SourceFrame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](i_.ID, objc.Sel("sourceFrame"))
	return rv
}


// The rectangle that the item occupies in the host app’s source window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovider/sourceframe
func (i_ ItemProvider) SetSourceFrame(value corefoundation.CGRect) {
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








