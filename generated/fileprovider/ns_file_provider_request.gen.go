// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSFileProviderRequest */


/* debug [class_header]: Header for NSFileProviderRequest */
// The class instance for the [FileProviderRequest] class.
var (
	FileProviderRequestClass     _FileProviderRequestClass
	FileProviderRequestClassOnce sync.Once
)

func getFileProviderRequestClass() _FileProviderRequestClass {
	FileProviderRequestClassOnce.Do(func() {
		FileProviderRequestClass = _FileProviderRequestClass{objc.GetClass("NSFileProviderRequest")}
	})
	return FileProviderRequestClass
}

type _FileProviderRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FileProviderRequest */
// An interface definition for the [FileProviderRequest] class.
type IFileProviderRequest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FileProviderRequest */
	// properties:
	DomainVersion() IFileProviderDomainVersion
	IsFileViewerRequest() bool
	IsSystemRequest() bool
	RequestingExecutable() objc.IObject /* cross-framework: NSURL */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FileProviderRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FileProviderRequest */
// Alloc allocates a new instance without initialization.
func (fc _FileProviderRequestClass) Alloc() FileProviderRequest {
	rv := objc.Send[FileProviderRequest](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FileProviderRequestClass) New() FileProviderRequest {
	rv := objc.Send[FileProviderRequest](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FileProviderRequest) Init() FileProviderRequest {
	rv := objc.Send[FileProviderRequest](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FileProviderRequest) Autorelease() FileProviderRequest {
	rv := objc.Send[FileProviderRequest](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFileProviderRequest creates a new FileProviderRequest instance.
func NewFileProviderRequest() FileProviderRequest {
	return getFileProviderRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FileProviderRequest */
// An object that provides information about the application requesting data from the File Provider extension.


// An object that provides information about the application requesting data from the File Provider extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderRequest
type FileProviderRequest struct {
	objectivec.Object
}

// FileProviderRequestFrom constructs a [FileProviderRequest] from an unsafe.Pointer.
//
// An object that provides information about the application requesting data from the File Provider extension.
func FileProviderRequestFrom(ptr unsafe.Pointer) FileProviderRequest {
	return FileProviderRequest{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FileProviderRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FileProviderRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FileProviderRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FileProviderRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FileProviderRequest */

// The version of the domain for the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderRequest/domainVersion
func (f_ FileProviderRequest) DomainVersion() IFileProviderDomainVersion {
	rv := objc.Send[FileProviderDomainVersion](f_.ID, objc.Sel("domainVersion"))
	return rv
}/* debug [instance_properties/getter]: domainVersion */


// A Boolean value that indicates whether the request came from Finder or related system file browsers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderRequest/isFileViewerRequest
func (f_ FileProviderRequest) IsFileViewerRequest() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isFileViewerRequest"))
	return rv
}/* debug [instance_properties/getter]: isFileViewerRequest */


// A Boolean value that indicates whether the request came from a system process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderRequest/isSystemRequest
func (f_ FileProviderRequest) IsSystemRequest() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isSystemRequest"))
	return rv
}/* debug [instance_properties/getter]: isSystemRequest */


// The URL of the requesting executable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderRequest/requestingExecutable
func (f_ FileProviderRequest) RequestingExecutable() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](f_.ID, objc.Sel("requestingExecutable"))
	return rv
}/* debug [instance_properties/getter]: requestingExecutable */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSFileProviderRequest */



