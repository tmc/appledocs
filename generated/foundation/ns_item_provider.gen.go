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
	CanLoadObjectOfClass(aClass unsafe.Pointer) bool
	LoadDataRepresentationForTypeIdentifierCompletionHandler(typeIdentifier string, completionHandler unsafe.Pointer) unsafe.Pointer
	LoadFileRepresentationForTypeIdentifierCompletionHandler(typeIdentifier string, completionHandler unsafe.Pointer) unsafe.Pointer
	LoadInPlaceFileRepresentationForTypeIdentifierCompletionHandler(typeIdentifier string, completionHandler unsafe.Pointer) unsafe.Pointer
	LoadObjectOfClassCompletionHandler(aClass unsafe.Pointer, completionHandler unsafe.Pointer) unsafe.Pointer
	RegisterCKShareWithContainerAllowedSharingOptionsPreparationHandler(container unsafe.Pointer, allowedOptions unsafe.Pointer, preparationHandler unsafe.Pointer)
	RegisterFileRepresentationForTypeIdentifierFileOptionsVisibilityLoadHandler(typeIdentifier string, fileOptions unsafe.Pointer, visibility unsafe.Pointer, loadHandler unsafe.Pointer)
}

// An item provider for conveying data or a file between processes during drag-and-drop or copy-and-paste activities, or from a host app to an app extension.
//
// Starting in iOS 11, item providers play a central role in drag and drop, and in copy and paste. They continue to play a role with app extensions. The system uses an internal queue when calling the completion blocks for the class. When using an item provider with drag and drop, ensure that UI updates take place on the main queue as follows:
//
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




// Creates a new item provider, employing a specified object’s type identifiers to specify the data representations eligible for the provider to load.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/init(object:)
func NewItemProviderWithObject(object objc.ID) ItemProvider {
	instance := getItemProviderClass().Alloc()
	rv := objc.Send[ItemProvider](instance.ID, objc.Sel("initWithObject:"), object)
	rv.Autorelease()
	return rv
}


// Returns a Boolean value indicating whether an item provider can load objects of a specified class.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/canLoadObject(ofClass:)-3eig9
func (i_ ItemProvider) CanLoadObjectOfClass(aClass unsafe.Pointer) bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("canLoadObjectOfClass:"), aClass)
	return rv
}

// Asynchronously copies the provided, typed data into a generic data object, returning a progress object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/loadDataRepresentation(forTypeIdentifier:completionHandler:)
func (i_ ItemProvider) LoadDataRepresentationForTypeIdentifierCompletionHandler(typeIdentifier string, completionHandler unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("loadDataRepresentationForTypeIdentifier:completionHandler:"), objc.String(typeIdentifier), completionHandler)
	return rv
}

// Asynchronously writes a copy of the provided, typed data to a temporary file, returning a progress object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/loadFileRepresentation(forTypeIdentifier:completionHandler:)
func (i_ ItemProvider) LoadFileRepresentationForTypeIdentifierCompletionHandler(typeIdentifier string, completionHandler unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("loadFileRepresentationForTypeIdentifier:completionHandler:"), objc.String(typeIdentifier), completionHandler)
	return rv
}

// Asynchronously opens a file in place, if possible, returning a progress object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/loadInPlaceFileRepresentation(forTypeIdentifier:completionHandler:)
func (i_ ItemProvider) LoadInPlaceFileRepresentationForTypeIdentifierCompletionHandler(typeIdentifier string, completionHandler unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("loadInPlaceFileRepresentationForTypeIdentifier:completionHandler:"), objc.String(typeIdentifier), completionHandler)
	return rv
}

// Asynchronously loads an object of a specified class to an item provider, returning a progress object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/loadObject(ofClass:completionHandler:)-8ak5d
func (i_ ItemProvider) LoadObjectOfClassCompletionHandler(aClass unsafe.Pointer, completionHandler unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("loadObjectOfClass:completionHandler:"), aClass, completionHandler)
	return rv
}

// Creates and registers a new collaboration object using a collection of records to share.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/registerCKShareWithContainer:allowedSharingOptions:preparationHandler:
func (i_ ItemProvider) RegisterCKShareWithContainerAllowedSharingOptionsPreparationHandler(container unsafe.Pointer, allowedOptions unsafe.Pointer, preparationHandler unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("registerCKShareWithContainer:allowedSharingOptions:preparationHandler:"), container, allowedOptions, preparationHandler)
}

// Registers a file-backed representation for an item, specifying file options, item visibility, and a load handler.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/registerFileRepresentation(forTypeIdentifier:fileOptions:visibility:loadHandler:)
func (i_ ItemProvider) RegisterFileRepresentationForTypeIdentifierFileOptionsVisibilityLoadHandler(typeIdentifier string, fileOptions unsafe.Pointer, visibility unsafe.Pointer, loadHandler unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("registerFileRepresentationForTypeIdentifier:fileOptions:visibility:loadHandler:"), objc.String(typeIdentifier), fileOptions, visibility, loadHandler)
}

// The preferred style for presenting the item provider’s data.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/preferredPresentationStyle-swift.property
func (i_ ItemProvider) PreferredPresentationStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("preferredPresentationStyle"))
	return rv
}


// SetPreferredPresentationStyle sets the value of the preferredPresentationStyle property.
// The preferred style for presenting the item provider’s data.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/preferredPresentationStyle-swift.property
func (i_ ItemProvider) SetPreferredPresentationStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPreferredPresentationStyle:"), value)
}

// The rectangle that the item occupies in the host app’s source window.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/sourceFrame
func (i_ ItemProvider) SourceFrame() Rect {
	rv := objc.Send[Rect](i_.ID, objc.Sel("sourceFrame"))
	return rv
}


