// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSFileProviderStringSearchRequest */


/* debug [class_header]: Header for NSFileProviderStringSearchRequest */
// The class instance for the [FileProviderStringSearchRequest] class.
var (
	FileProviderStringSearchRequestClass     _FileProviderStringSearchRequestClass
	FileProviderStringSearchRequestClassOnce sync.Once
)

func getFileProviderStringSearchRequestClass() _FileProviderStringSearchRequestClass {
	FileProviderStringSearchRequestClassOnce.Do(func() {
		FileProviderStringSearchRequestClass = _FileProviderStringSearchRequestClass{objc.GetClass("NSFileProviderStringSearchRequest")}
	})
	return FileProviderStringSearchRequestClass
}

type _FileProviderStringSearchRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FileProviderStringSearchRequest */
// An interface definition for the [FileProviderStringSearchRequest] class.
type IFileProviderStringSearchRequest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FileProviderStringSearchRequest */
	// properties:
	DesiredNumberOfResults() int
	Query() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FileProviderStringSearchRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FileProviderStringSearchRequest */
// Alloc allocates a new instance without initialization.
func (fc _FileProviderStringSearchRequestClass) Alloc() FileProviderStringSearchRequest {
	rv := objc.Send[FileProviderStringSearchRequest](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FileProviderStringSearchRequestClass) New() FileProviderStringSearchRequest {
	rv := objc.Send[FileProviderStringSearchRequest](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FileProviderStringSearchRequest) Init() FileProviderStringSearchRequest {
	rv := objc.Send[FileProviderStringSearchRequest](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FileProviderStringSearchRequest) Autorelease() FileProviderStringSearchRequest {
	rv := objc.Send[FileProviderStringSearchRequest](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFileProviderStringSearchRequest creates a new FileProviderStringSearchRequest instance.
func NewFileProviderStringSearchRequest() FileProviderStringSearchRequest {
	return getFileProviderStringSearchRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FileProviderStringSearchRequest */
// A type that contains details of a string-based search request.


// A type that contains details of a string-based search request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderStringSearchRequest
type FileProviderStringSearchRequest struct {
	objectivec.Object
}

// FileProviderStringSearchRequestFrom constructs a [FileProviderStringSearchRequest] from an unsafe.Pointer.
//
// A type that contains details of a string-based search request.
func FileProviderStringSearchRequestFrom(ptr unsafe.Pointer) FileProviderStringSearchRequest {
	return FileProviderStringSearchRequest{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FileProviderStringSearchRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FileProviderStringSearchRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FileProviderStringSearchRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FileProviderStringSearchRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FileProviderStringSearchRequest */

// How many results the system is requesting. This is a hint to the extension, to help avoid unnecessary work. The extension may return more results than this.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderStringSearchRequest/desiredNumberOfResults
func (f_ FileProviderStringSearchRequest) DesiredNumberOfResults() int {
	rv := objc.Send[int](f_.ID, objc.Sel("desiredNumberOfResults"))
	return rv
}/* debug [instance_properties/getter]: desiredNumberOfResults */


// A plaintext string, representing the query a person entered into the system search UI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderStringSearchRequest/query
func (f_ FileProviderStringSearchRequest) Query() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](f_.ID, objc.Sel("query"))
	return rv
}/* debug [instance_properties/getter]: query */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSFileProviderStringSearchRequest */





