// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	ProjectExtensionData() objc.IObject /* cross-framework: Data */
	SetProjectExtensionData(value objc.IObject /* cross-framework: Data */)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

// A request to change asset data in a Photos project extension.
//
// Make a project change request to alter a project’s title or metadata. Respond to project change requests by updating your user interface as assets are added, modified, or removed.

// A request to change asset data in a Photos project extension.
//
// [Full Topic]
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

// Compressed project-specific data to use in the change request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phprojectchangerequest/projectextensiondata
func (p_ PHProjectChangeRequest) ProjectExtensionData() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](p_.ID, objc.Sel("projectExtensionData"))
	return rv
}

// Compressed project-specific data to use in the change request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phprojectchangerequest/projectextensiondata
func (p_ PHProjectChangeRequest) SetProjectExtensionData(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setProjectExtensionData:"), value)
}

// The title of the change request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phprojectchangerequest/title
func (p_ PHProjectChangeRequest) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("title"))
	return rv
}

// The title of the change request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phprojectchangerequest/title
func (p_ PHProjectChangeRequest) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTitle:"), value)
}
