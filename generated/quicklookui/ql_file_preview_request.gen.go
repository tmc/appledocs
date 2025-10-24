// Code generated from Apple documentation for QuickLookUI. DO NOT EDIT.

package quicklookui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class QLFilePreviewRequest */


/* debug [class_header]: Header for QLFilePreviewRequest */
// The class instance for the [FilePreviewRequest] class.
var (
	FilePreviewRequestClass     _FilePreviewRequestClass
	FilePreviewRequestClassOnce sync.Once
)

func getFilePreviewRequestClass() _FilePreviewRequestClass {
	FilePreviewRequestClassOnce.Do(func() {
		FilePreviewRequestClass = _FilePreviewRequestClass{objc.GetClass("QLFilePreviewRequest")}
	})
	return FilePreviewRequestClass
}

type _FilePreviewRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FilePreviewRequest */
// An interface definition for the [FilePreviewRequest] class.
type IFilePreviewRequest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FilePreviewRequest */
	// properties:
	FileURL() objc.IObject /* cross-framework: NSURL */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FilePreviewRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FilePreviewRequest */
// Alloc allocates a new instance without initialization.
func (fc _FilePreviewRequestClass) Alloc() FilePreviewRequest {
	rv := objc.Send[FilePreviewRequest](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FilePreviewRequestClass) New() FilePreviewRequest {
	rv := objc.Send[FilePreviewRequest](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FilePreviewRequest) Init() FilePreviewRequest {
	rv := objc.Send[FilePreviewRequest](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FilePreviewRequest) Autorelease() FilePreviewRequest {
	rv := objc.Send[FilePreviewRequest](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFilePreviewRequest creates a new FilePreviewRequest instance.
func NewFilePreviewRequest() FilePreviewRequest {
	return getFilePreviewRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FilePreviewRequest */
// A Quick Look preview request that indicates the content to preview.
//
// The system provides a to the method of your data-based Quick Look extension.


// A Quick Look preview request that indicates the content to preview.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLFilePreviewRequest
type FilePreviewRequest struct {
	objectivec.Object
}

// FilePreviewRequestFrom constructs a [FilePreviewRequest] from an unsafe.Pointer.
//
// A Quick Look preview request that indicates the content to preview.
func FilePreviewRequestFrom(ptr unsafe.Pointer) FilePreviewRequest {
	return FilePreviewRequest{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FilePreviewRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FilePreviewRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FilePreviewRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FilePreviewRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FilePreviewRequest */

// The URL that indicates the content to preview.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookUI/QLFilePreviewRequest/fileURL
func (f_ FilePreviewRequest) FileURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](f_.ID, objc.Sel("fileURL"))
	return rv
}/* debug [instance_properties/getter]: fileURL */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class QLFilePreviewRequest */



