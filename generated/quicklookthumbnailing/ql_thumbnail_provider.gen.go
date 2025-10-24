// Code generated from Apple documentation for QuickLookThumbnailing. DO NOT EDIT.

package quicklookthumbnailing

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class QLThumbnailProvider */


/* debug [class_header]: Header for QLThumbnailProvider */
// The class instance for the [ThumbnailProvider] class.
var (
	ThumbnailProviderClass     _ThumbnailProviderClass
	ThumbnailProviderClassOnce sync.Once
)

func getThumbnailProviderClass() _ThumbnailProviderClass {
	ThumbnailProviderClassOnce.Do(func() {
		ThumbnailProviderClass = _ThumbnailProviderClass{objc.GetClass("QLThumbnailProvider")}
	})
	return ThumbnailProviderClass
}

type _ThumbnailProviderClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ThumbnailProvider */
// An interface definition for the [ThumbnailProvider] class.
type IThumbnailProvider interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ThumbnailProvider */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ThumbnailProvider */
	// methods:
	ProvideThumbnailForFileRequestCompletionHandler(request IQLFileThumbnailRequest, handler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ThumbnailProvider */
// Alloc allocates a new instance without initialization.
func (tc _ThumbnailProviderClass) Alloc() ThumbnailProvider {
	rv := objc.Send[ThumbnailProvider](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _ThumbnailProviderClass) New() ThumbnailProvider {
	rv := objc.Send[ThumbnailProvider](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ ThumbnailProvider) Init() ThumbnailProvider {
	rv := objc.Send[ThumbnailProvider](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ ThumbnailProvider) Autorelease() ThumbnailProvider {
	rv := objc.Send[ThumbnailProvider](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewThumbnailProvider creates a new ThumbnailProvider instance.
func NewThumbnailProvider() ThumbnailProvider {
	return getThumbnailProviderClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ThumbnailProvider */
// An abstract base class for creating thumbnails of custom file types.


// An abstract base class for creating thumbnails of custom file types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailProvider
type ThumbnailProvider struct {
	objectivec.Object
}

// ThumbnailProviderFrom constructs a [ThumbnailProvider] from an unsafe.Pointer.
//
// An abstract base class for creating thumbnails of custom file types.
func ThumbnailProviderFrom(ptr unsafe.Pointer) ThumbnailProvider {
	return ThumbnailProvider{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ThumbnailProvider *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ThumbnailProvider */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ThumbnailProvider */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ThumbnailProvider */

// Creates a thumbnail of a custom file type for a specific request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailProvider/provideThumbnail(for:_:)
func (t_ ThumbnailProvider) ProvideThumbnailForFileRequestCompletionHandler(request IQLFileThumbnailRequest, handler unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("provideThumbnailForFileRequest:completionHandler:"), request, handler)
}/* debug [instance_methods/method]: ProvideThumbnailForFileRequestCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ThumbnailProvider */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class QLThumbnailProvider */



