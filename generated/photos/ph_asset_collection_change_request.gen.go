// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [PHAssetCollectionChangeRequest] class.
var (
	PHAssetCollectionChangeRequestClass     _PHAssetCollectionChangeRequestClass
	PHAssetCollectionChangeRequestClassOnce sync.Once
)

func getPHAssetCollectionChangeRequestClass() _PHAssetCollectionChangeRequestClass {
	PHAssetCollectionChangeRequestClassOnce.Do(func() {
		PHAssetCollectionChangeRequestClass = _PHAssetCollectionChangeRequestClass{objc.GetClass("PHAssetCollectionChangeRequest")}
	})
	return PHAssetCollectionChangeRequestClass
}

type _PHAssetCollectionChangeRequestClass struct {
	class objc.Class
}

// An interface definition for the [PHAssetCollectionChangeRequest] class.
type IPHAssetCollectionChangeRequest interface {
	IPHChangeRequest
}

// A request to create, delete, or modify a Photos asset collection, for use in a photo library change block.
//
// You use the class to request changes for objects. To make changes to asset collections (such as user-created albums) in the Photos library, create a change request using the appropriate class method for the change you want to perform. Call the method to create a new asset collection. Call the method to delete existing asset collections. Call the or method to modify a collection’s metadata or list of member assets. Before creating a change request, use the method to verify that the collection allows the edit operation you’re requesting. If you attempt to perform an unsupported edit operation, Photos throws an exception. A change request for creating or modifying an asset collection works like a mutable version of the asset collection object. Use the change request’s properties and instance methods to request changes to the asset collection itself. For example, the following code removes an asset from an album. After Photos runs the change block and calls your completion handler, the asset collection’s state reflects the changes you requested in the block. If you create or use a change request object outside a photo library change block, Photos raises an Objective-C exception. For details on change blocks, see .
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetCollectionChangeRequest
type PHAssetCollectionChangeRequest struct {
	PHChangeRequest
}

// PHAssetCollectionChangeRequestFrom constructs a [PHAssetCollectionChangeRequest] from an unsafe.Pointer.
//
// A request to create, delete, or modify a Photos asset collection, for use in a photo library change block.
func PHAssetCollectionChangeRequestFrom(ptr unsafe.Pointer) PHAssetCollectionChangeRequest {
	return PHAssetCollectionChangeRequest{
		PHChangeRequest: PHChangeRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PHAssetCollectionChangeRequestClass) Alloc() PHAssetCollectionChangeRequest {
	rv := objc.Send[PHAssetCollectionChangeRequest](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHAssetCollectionChangeRequestClass) New() PHAssetCollectionChangeRequest {
	rv := objc.Send[PHAssetCollectionChangeRequest](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHAssetCollectionChangeRequest) Init() PHAssetCollectionChangeRequest {
	rv := objc.Send[PHAssetCollectionChangeRequest](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHAssetCollectionChangeRequest) Autorelease() PHAssetCollectionChangeRequest {
	rv := objc.Send[PHAssetCollectionChangeRequest](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHAssetCollectionChangeRequest creates a new PHAssetCollectionChangeRequest instance.
func NewPHAssetCollectionChangeRequest() PHAssetCollectionChangeRequest {
	return getPHAssetCollectionChangeRequestClass().New()
}


// Creates a request for adding a new asset collection to the Photos library.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetCollectionChangeRequest/creationRequestForAssetCollection(withTitle:)
func (pc _PHAssetCollectionChangeRequestClass) CreationRequestForAssetCollectionWithTitle(title appkit.string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("creationRequestForAssetCollectionWithTitle:"), title)
	return rv
}

// A placeholder object for the asset collection that the change request creates.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetcollectionchangerequest/placeholderforcreatedassetcollection
func (p_ PHAssetCollectionChangeRequest) PlaceholderForCreatedAssetCollection() PHObjectPlaceholder {
	rv := objc.Send[PHObjectPlaceholder](p_.ID, objc.Sel("placeholderForCreatedAssetCollection"))
	return rv
}


// SetPlaceholderForCreatedAssetCollection sets the value of the placeholderForCreatedAssetCollection property.
// A placeholder object for the asset collection that the change request creates.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetcollectionchangerequest/placeholderforcreatedassetcollection
func (p_ PHAssetCollectionChangeRequest) SetPlaceholderForCreatedAssetCollection(value IPHObjectPlaceholder) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPlaceholderForCreatedAssetCollection:"), value)
}

// The displayed name of the asset collection.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetcollectionchangerequest/title
func (p_ PHAssetCollectionChangeRequest) Title() appkit.string {
	rv := objc.Send[appkit.string](p_.ID, objc.Sel("title"))
	return rv
}


// SetTitle sets the value of the title property.
// The displayed name of the asset collection.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetcollectionchangerequest/title
func (p_ PHAssetCollectionChangeRequest) SetTitle(value appkit.string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTitle:"), value)
}



