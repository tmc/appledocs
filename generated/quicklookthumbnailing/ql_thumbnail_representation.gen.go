// Code generated from Apple documentation for QuickLookThumbnailing. DO NOT EDIT.

package quicklookthumbnailing

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class QLThumbnailRepresentation */


/* debug [class_header]: Header for QLThumbnailRepresentation */
// The class instance for the [ThumbnailRepresentation] class.
var (
	ThumbnailRepresentationClass     _ThumbnailRepresentationClass
	ThumbnailRepresentationClassOnce sync.Once
)

func getThumbnailRepresentationClass() _ThumbnailRepresentationClass {
	ThumbnailRepresentationClassOnce.Do(func() {
		ThumbnailRepresentationClass = _ThumbnailRepresentationClass{objc.GetClass("QLThumbnailRepresentation")}
	})
	return ThumbnailRepresentationClass
}

type _ThumbnailRepresentationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ThumbnailRepresentation */
// An interface definition for the [ThumbnailRepresentation] class.
type IThumbnailRepresentation interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ThumbnailRepresentation */
	// properties:
	CGImage() ImageRef /* not a class type */
	ContentRect() corefoundation.CGRect
	NSImage() appkit.Image
	Type() ThumbnailRepresentationType
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ThumbnailRepresentation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ThumbnailRepresentation */
// Alloc allocates a new instance without initialization.
func (tc _ThumbnailRepresentationClass) Alloc() ThumbnailRepresentation {
	rv := objc.Send[ThumbnailRepresentation](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _ThumbnailRepresentationClass) New() ThumbnailRepresentation {
	rv := objc.Send[ThumbnailRepresentation](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ ThumbnailRepresentation) Init() ThumbnailRepresentation {
	rv := objc.Send[ThumbnailRepresentation](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ ThumbnailRepresentation) Autorelease() ThumbnailRepresentation {
	rv := objc.Send[ThumbnailRepresentation](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewThumbnailRepresentation creates a new ThumbnailRepresentation instance.
func NewThumbnailRepresentation() ThumbnailRepresentation {
	return getThumbnailRepresentationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ThumbnailRepresentation */
// Information about the thumbnail that the thumbnail generator returns.
//
// QuickLook Thumbnailing is a non-UI framework, so your app doesn’t have to link to either or . Quicklook Thumbnailing generates a thumbnail as a Core Graphics image object and makes the thumbnail available as the property. If an app links to AppKit or UIKit, the thumbnail is available through the or properties. For more information on the different types of thumbnails that can create, see .


// Information about the thumbnail that the thumbnail generator returns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailRepresentation
type ThumbnailRepresentation struct {
	objectivec.Object
}

// ThumbnailRepresentationFrom constructs a [ThumbnailRepresentation] from an unsafe.Pointer.
//
// Information about the thumbnail that the thumbnail generator returns.
func ThumbnailRepresentationFrom(ptr unsafe.Pointer) ThumbnailRepresentation {
	return ThumbnailRepresentation{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ThumbnailRepresentation *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ThumbnailRepresentation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ThumbnailRepresentation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ThumbnailRepresentation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ThumbnailRepresentation */

// A thumbnail in the form of a Core Graphics image object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailRepresentation/cgImage
func (t_ ThumbnailRepresentation) CGImage() ImageRef /* not a class type */ {
	rv := objc.Send[ImageRef](t_.ID, objc.Sel("CGImage"))
	return rv
}/* debug [instance_properties/getter]: CGImage */


// The rectangle within the thumbnail image of the document that represents its contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailRepresentation/contentRect
func (t_ ThumbnailRepresentation) ContentRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t_.ID, objc.Sel("contentRect"))
	return rv
}/* debug [instance_properties/getter]: contentRect */


// A thumbnail in the form of an AppKit image object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailRepresentation/nsImage
func (t_ ThumbnailRepresentation) NSImage() appkit.Image {
	rv := objc.Send[appkit.Image](t_.ID, objc.Sel("NSImage"))
	return rv
}/* debug [instance_properties/getter]: NSImage */


// The type of thumbnail.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailRepresentation/type
func (t_ ThumbnailRepresentation) Type() ThumbnailRepresentationType {
	rv := objc.Send[ThumbnailRepresentationType](t_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class QLThumbnailRepresentation */


