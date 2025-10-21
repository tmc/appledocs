// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PHProjectChangeRequest] class.
var (
	PHProjectChangeRequestClass     _PHProjectChangeRequestClass
	PHProjectChangeRequestClassOnce sync.Once
)

func getPHProjectChangeRequestClass() _PHProjectChangeRequestClass {
	PHProjectChangeRequestClassOnce.Do(func() {
		PHProjectChangeRequestClass = _PHProjectChangeRequestClass{objc.GetClass("PHProjectChangeRequest")}
	})
	return PHProjectChangeRequestClass
}

type _PHProjectChangeRequestClass struct {
	class objc.Class
}

// An interface definition for the [PHProjectChangeRequest] class.
type IPHProjectChangeRequest interface {
	IPHChangeRequest
	RemoveAssets(assets objc.ID)
	SetKeyAsset(keyAsset unsafe.Pointer)
	SetProjectPreviewImage(previewImage unsafe.Pointer)
}

// A request to change asset data in a Photos project extension.
//
// Make a project change request to alter a project’s title or metadata. Respond to project change requests by updating your user interface as assets are added, modified, or removed.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHProjectChangeRequest
type PHProjectChangeRequest struct {
	PHChangeRequest
}

// PHProjectChangeRequestFrom constructs a [PHProjectChangeRequest] from an unsafe.Pointer.
//
// A request to change asset data in a Photos project extension.
func PHProjectChangeRequestFrom(ptr unsafe.Pointer) PHProjectChangeRequest {
	return PHProjectChangeRequest{
		PHChangeRequest: PHChangeRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PHProjectChangeRequestClass) Alloc() PHProjectChangeRequest {
	rv := objc.Send[PHProjectChangeRequest](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHProjectChangeRequestClass) New() PHProjectChangeRequest {
	rv := objc.Send[PHProjectChangeRequest](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHProjectChangeRequest) Init() PHProjectChangeRequest {
	rv := objc.Send[PHProjectChangeRequest](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHProjectChangeRequest) Autorelease() PHProjectChangeRequest {
	rv := objc.Send[PHProjectChangeRequest](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHProjectChangeRequest creates a new PHProjectChangeRequest instance.
func NewPHProjectChangeRequest() PHProjectChangeRequest {
	return getPHProjectChangeRequestClass().New()
}


// Creates a change request around the specified project.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHProjectChangeRequest/init(project:)
func NewPHProjectChangeRequestWithProject(project unsafe.Pointer) PHProjectChangeRequest {
	instance := getPHProjectChangeRequestClass().Alloc()
	rv := objc.Send[PHProjectChangeRequest](instance.ID, objc.Sel("initWithProject:"), project)
	rv.Autorelease()
	return rv
}


// Removes the specified assets from the project.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHProjectChangeRequest/removeAssets:
func (p_ PHProjectChangeRequest) RemoveAssets(assets objc.ID) {
	objc.Send[objc.ID](p_.ID, objc.Sel("removeAssets:"), assets)
}

// Sets the key asset representing the project.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHProjectChangeRequest/setKeyAsset(_:)
func (p_ PHProjectChangeRequest) SetKeyAsset(keyAsset unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setKeyAsset:"), keyAsset)
}

// Updates the project preview in Photos.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHProjectChangeRequest/setProjectPreviewImage(_:)
func (p_ PHProjectChangeRequest) SetProjectPreviewImage(previewImage unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setProjectPreviewImage:"), previewImage)
}

// Compressed project-specific data to use in the change request.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHProjectChangeRequest/projectExtensionData
func (p_ PHProjectChangeRequest) ProjectExtensionData() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("projectExtensionData"))
	return rv
}


// SetProjectExtensionData sets the value of the projectExtensionData property.
// Compressed project-specific data to use in the change request.

//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHProjectChangeRequest/projectExtensionData
func (p_ PHProjectChangeRequest) SetProjectExtensionData(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setProjectExtensionData:"), value)
}
// The title of the change request.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHProjectChangeRequest/title
func (p_ PHProjectChangeRequest) Title() string {
	rv := objc.Send[string](p_.ID, objc.Sel("title"))
	return rv
}


// SetTitle sets the value of the title property.
// The title of the change request.

//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHProjectChangeRequest/title
func (p_ PHProjectChangeRequest) SetTitle(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTitle:"), objc.String(value))
}

