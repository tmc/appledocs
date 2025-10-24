// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [PHCollectionListChangeRequest] class.
var (
	PHCollectionListChangeRequestClass     _PHCollectionListChangeRequestClass
	PHCollectionListChangeRequestClassOnce sync.Once
)

func getPHCollectionListChangeRequestClass() _PHCollectionListChangeRequestClass {
	PHCollectionListChangeRequestClassOnce.Do(func() {
		PHCollectionListChangeRequestClass = _PHCollectionListChangeRequestClass{objc.GetClass("PHCollectionListChangeRequest")}
	})
	return PHCollectionListChangeRequestClass
}

type _PHCollectionListChangeRequestClass struct {
	class objc.Class
}

// An interface definition for the [PHCollectionListChangeRequest] class.
type IPHCollectionListChangeRequest interface {
	IPHChangeRequest
	// properties:
	PlaceholderForCreatedCollectionList() IPHObjectPlaceholder
	SetPlaceholderForCreatedCollectionList(value IPHObjectPlaceholder)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

// A request to create, delete, or modify a Photos collection list, for use in a photo library change block.
//
// You use the class to request changes for objects. To make changes to collection lists (such as folders containing user-created albums) in the Photos library, create a change request using the appropriate class method for the change you want to perform. Call the method to create a new asset collection. Call the method to delete existing asset collections. Call the or method to modify a collection’s metadata or its list of child collections. Before creating a change request, use the method to verify that the collection allows the edit operation you’re requesting. If you attempt to perform an unsupported edit operation, Photos throws an exception. A change request for creating or modifying a collection list works like a mutable version of the collection list object. Use the change request’s properties and instance methods to request changes to the collection list itself. For example, the following code removes an album from a folder. After Photos runs the change block and calls your completion handler, the collection list’s state reflects the changes you requested in the block. If you create or use a change request object outside a photo library change block, Photos raises an Objective-C exception. For details on change blocks, see .


// A request to create, delete, or modify a Photos collection list, for use in a photo library change block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionListChangeRequest
type PHCollectionListChangeRequest struct {
	PHChangeRequest
}

// PHCollectionListChangeRequestFrom constructs a [PHCollectionListChangeRequest] from an unsafe.Pointer.
//
// A request to create, delete, or modify a Photos collection list, for use in a photo library change block.
func PHCollectionListChangeRequestFrom(ptr unsafe.Pointer) PHCollectionListChangeRequest {
	return PHCollectionListChangeRequest{
		PHChangeRequest: PHChangeRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PHCollectionListChangeRequestClass) Alloc() PHCollectionListChangeRequest {
	rv := objc.Send[PHCollectionListChangeRequest](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHCollectionListChangeRequestClass) New() PHCollectionListChangeRequest {
	rv := objc.Send[PHCollectionListChangeRequest](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHCollectionListChangeRequest) Init() PHCollectionListChangeRequest {
	rv := objc.Send[PHCollectionListChangeRequest](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHCollectionListChangeRequest) Autorelease() PHCollectionListChangeRequest {
	rv := objc.Send[PHCollectionListChangeRequest](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHCollectionListChangeRequest creates a new PHCollectionListChangeRequest instance.
func NewPHCollectionListChangeRequest() PHCollectionListChangeRequest {
	return getPHCollectionListChangeRequestClass().New()
}



// A placeholder object for the collection list that the change request creates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcollectionlistchangerequest/placeholderforcreatedcollectionlist
func (p_ PHCollectionListChangeRequest) PlaceholderForCreatedCollectionList() IPHObjectPlaceholder {
	rv := objc.Send[PHObjectPlaceholder](p_.ID, objc.Sel("placeholderForCreatedCollectionList"))
	return rv
}


// A placeholder object for the collection list that the change request creates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcollectionlistchangerequest/placeholderforcreatedcollectionlist
func (p_ PHCollectionListChangeRequest) SetPlaceholderForCreatedCollectionList(value IPHObjectPlaceholder) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPlaceholderForCreatedCollectionList:"), value)
}


// The displayed name of the collection list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcollectionlistchangerequest/title
func (p_ PHCollectionListChangeRequest) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("title"))
	return rv
}


// The displayed name of the collection list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcollectionlistchangerequest/title
func (p_ PHCollectionListChangeRequest) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTitle:"), value)
}



