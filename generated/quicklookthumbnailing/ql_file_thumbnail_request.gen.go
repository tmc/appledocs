// Code generated from Apple documentation for QuickLookThumbnailing. DO NOT EDIT.

package quicklookthumbnailing

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class QLFileThumbnailRequest */


/* debug [class_header]: Header for QLFileThumbnailRequest */
// The class instance for the [FileThumbnailRequest] class.
var (
	FileThumbnailRequestClass     _FileThumbnailRequestClass
	FileThumbnailRequestClassOnce sync.Once
)

func getFileThumbnailRequestClass() _FileThumbnailRequestClass {
	FileThumbnailRequestClassOnce.Do(func() {
		FileThumbnailRequestClass = _FileThumbnailRequestClass{objc.GetClass("QLFileThumbnailRequest")}
	})
	return FileThumbnailRequestClass
}

type _FileThumbnailRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FileThumbnailRequest */
// An interface definition for the [FileThumbnailRequest] class.
type IFileThumbnailRequest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FileThumbnailRequest */
	// properties:
	FileURL() objc.IObject /* cross-framework: NSURL */
	MaximumSize() corefoundation.CGSize
	MinimumSize() corefoundation.CGSize
	Scale() float64
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FileThumbnailRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FileThumbnailRequest */
// Alloc allocates a new instance without initialization.
func (fc _FileThumbnailRequestClass) Alloc() FileThumbnailRequest {
	rv := objc.Send[FileThumbnailRequest](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FileThumbnailRequestClass) New() FileThumbnailRequest {
	rv := objc.Send[FileThumbnailRequest](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FileThumbnailRequest) Init() FileThumbnailRequest {
	rv := objc.Send[FileThumbnailRequest](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FileThumbnailRequest) Autorelease() FileThumbnailRequest {
	rv := objc.Send[FileThumbnailRequest](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFileThumbnailRequest creates a new FileThumbnailRequest instance.
func NewFileThumbnailRequest() FileThumbnailRequest {
	return getFileThumbnailRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FileThumbnailRequest */
// A request to generate a thumbnail for a custom file type.


// A request to generate a thumbnail for a custom file type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLFileThumbnailRequest
type FileThumbnailRequest struct {
	objectivec.Object
}

// FileThumbnailRequestFrom constructs a [FileThumbnailRequest] from an unsafe.Pointer.
//
// A request to generate a thumbnail for a custom file type.
func FileThumbnailRequestFrom(ptr unsafe.Pointer) FileThumbnailRequest {
	return FileThumbnailRequest{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FileThumbnailRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FileThumbnailRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FileThumbnailRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FileThumbnailRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FileThumbnailRequest */

// The URL of the image file to use for the thumbnail.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLFileThumbnailRequest/fileURL
func (f_ FileThumbnailRequest) FileURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](f_.ID, objc.Sel("fileURL"))
	return rv
}/* debug [instance_properties/getter]: fileURL */


// The maximum accepted size of a thumbnail.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLFileThumbnailRequest/maximumSize
func (f_ FileThumbnailRequest) MaximumSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](f_.ID, objc.Sel("maximumSize"))
	return rv
}/* debug [instance_properties/getter]: maximumSize */


// The minimum accepted size of a thumbnail.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLFileThumbnailRequest/minimumSize
func (f_ FileThumbnailRequest) MinimumSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](f_.ID, objc.Sel("minimumSize"))
	return rv
}/* debug [instance_properties/getter]: minimumSize */


// The scale of the requested thumbnail.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLFileThumbnailRequest/scale
func (f_ FileThumbnailRequest) Scale() float64 {
	rv := objc.Send[float64](f_.ID, objc.Sel("scale"))
	return rv
}/* debug [instance_properties/getter]: scale */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class QLFileThumbnailRequest */



