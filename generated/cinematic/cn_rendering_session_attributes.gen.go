// Code generated from Apple documentation for Cinematic. DO NOT EDIT.

package cinematic

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/avfoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CNRenderingSessionAttributes */


/* debug [class_header]: Header for CNRenderingSessionAttributes */
// The class instance for the [CNRenderingSessionAttributes] class.
var (
	CNRenderingSessionAttributesClass     _CNRenderingSessionAttributesClass
	CNRenderingSessionAttributesClassOnce sync.Once
)

func getCNRenderingSessionAttributesClass() _CNRenderingSessionAttributesClass {
	CNRenderingSessionAttributesClassOnce.Do(func() {
		CNRenderingSessionAttributesClass = _CNRenderingSessionAttributesClass{objc.GetClass("CNRenderingSessionAttributes")}
	})
	return CNRenderingSessionAttributesClass
}

type _CNRenderingSessionAttributesClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNRenderingSessionAttributes */
// An interface definition for the [CNRenderingSessionAttributes] class.
type ICNRenderingSessionAttributes interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CNRenderingSessionAttributes */
	// properties:
	RenderingVersion() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNRenderingSessionAttributes */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNRenderingSessionAttributes */
// Alloc allocates a new instance without initialization.
func (cc _CNRenderingSessionAttributesClass) Alloc() CNRenderingSessionAttributes {
	rv := objc.Send[CNRenderingSessionAttributes](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNRenderingSessionAttributesClass) New() CNRenderingSessionAttributes {
	rv := objc.Send[CNRenderingSessionAttributes](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNRenderingSessionAttributes) Init() CNRenderingSessionAttributes {
	rv := objc.Send[CNRenderingSessionAttributes](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNRenderingSessionAttributes) Autorelease() CNRenderingSessionAttributes {
	rv := objc.Send[CNRenderingSessionAttributes](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNRenderingSessionAttributes creates a new CNRenderingSessionAttributes instance.
func NewCNRenderingSessionAttributes() CNRenderingSessionAttributes {
	return getCNRenderingSessionAttributesClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNRenderingSessionAttributes */
// A structure for movie-wide attributes required for proper rendering.
//
// The attributes include camera intrinsics from the camera on which the video was originally recorded.


// A structure for movie-wide attributes required for proper rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNRenderingSessionAttributes
type CNRenderingSessionAttributes struct {
	objectivec.Object
}

// CNRenderingSessionAttributesFrom constructs a [CNRenderingSessionAttributes] from an unsafe.Pointer.
//
// A structure for movie-wide attributes required for proper rendering.
func CNRenderingSessionAttributesFrom(ptr unsafe.Pointer) CNRenderingSessionAttributes {
	return CNRenderingSessionAttributes{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNRenderingSessionAttributes *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNRenderingSessionAttributes */

// Loads the rendering session attributes from an asset asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNRenderingSessionAttributes/loadFromAsset:completionHandler:
func (cc _CNRenderingSessionAttributesClass) LoadFromAssetCompletionHandler(asset avfoundation.Asset, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("loadFromAsset:completionHandler:"), asset, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LoadFromAssetCompletionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNRenderingSessionAttributes */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNRenderingSessionAttributes */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNRenderingSessionAttributes */

// The primary version number used to render the original Cinematic move that determines compatibility.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNRenderingSessionAttributes/renderingVersion
func (c_ CNRenderingSessionAttributes) RenderingVersion() int {
	rv := objc.Send[int](c_.ID, objc.Sel("renderingVersion"))
	return rv
}/* debug [instance_properties/getter]: renderingVersion */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNRenderingSessionAttributes */



