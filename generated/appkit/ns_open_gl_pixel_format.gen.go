// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSOpenGLPixelFormat */


/* debug [class_header]: Header for NSOpenGLPixelFormat */
// The class instance for the [OpenGLPixelFormat] class.
var (
	OpenGLPixelFormatClass     _OpenGLPixelFormatClass
	OpenGLPixelFormatClassOnce sync.Once
)

func getOpenGLPixelFormatClass() _OpenGLPixelFormatClass {
	OpenGLPixelFormatClassOnce.Do(func() {
		OpenGLPixelFormatClass = _OpenGLPixelFormatClass{objc.GetClass("NSOpenGLPixelFormat")}
	})
	return OpenGLPixelFormatClass
}

type _OpenGLPixelFormatClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for OpenGLPixelFormat */
// An interface definition for the [OpenGLPixelFormat] class.
type IOpenGLPixelFormat interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for OpenGLPixelFormat */
	// properties:
	CGLPixelFormatObj() LPixelFormatObj /* not a class type */
	NumberOfVirtualScreens() objectivec.IObject
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for OpenGLPixelFormat */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for OpenGLPixelFormat */
// Alloc allocates a new instance without initialization.
func (oc _OpenGLPixelFormatClass) Alloc() OpenGLPixelFormat {
	rv := objc.Send[OpenGLPixelFormat](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (oc _OpenGLPixelFormatClass) New() OpenGLPixelFormat {
	rv := objc.Send[OpenGLPixelFormat](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OpenGLPixelFormat) Init() OpenGLPixelFormat {
	rv := objc.Send[OpenGLPixelFormat](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OpenGLPixelFormat) Autorelease() OpenGLPixelFormat {
	rv := objc.Send[OpenGLPixelFormat](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOpenGLPixelFormat creates a new OpenGLPixelFormat instance.
func NewOpenGLPixelFormat() OpenGLPixelFormat {
	return getOpenGLPixelFormatClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for OpenGLPixelFormat */
// An object that specifies the types of buffers and other attributes of the OpenGL context.
//
// To render with OpenGL into an , you must specify the context’s pixel format. Every object wraps a low-level, platform-specific Core OpenGL (CGL) pixel format object. Your application can retrieve the CGL pixel format object by calling the method. For more information on the underling CGL pixel format object, see .


// An object that specifies the types of buffers and other attributes of the OpenGL context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLPixelFormat
type OpenGLPixelFormat struct {
	objectivec.Object
}

// OpenGLPixelFormatFrom constructs a [OpenGLPixelFormat] from an unsafe.Pointer.
//
// An object that specifies the types of buffers and other attributes of the OpenGL context.
func OpenGLPixelFormatFrom(ptr unsafe.Pointer) OpenGLPixelFormat {
	return OpenGLPixelFormat{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for OpenGLPixelFormat */

// Returns an OpenGL pixel format object initialized with specified pixel format attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLPixelFormat/init(attributes:)
func NewOpenGLPixelFormatWithAttributes(attribs OpenGLPixelFormatAttribute /* typedef */) OpenGLPixelFormat {
	instance := getOpenGLPixelFormatClass().Alloc()
	rv := objc.Send[OpenGLPixelFormat](instance.ID, objc.Sel("initWithAttributes:"), attribs)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewOpenGLPixelFormatWithAttributes */


// Returns an OpenGL pixel format object initialized with using an existing CGL pixel format object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLPixelFormat/init(cglPixelFormatObj:)
func NewOpenGLPixelFormatWithCGLPixelFormatObj(format LPixelFormatObj /* not a class type */) OpenGLPixelFormat {
	instance := getOpenGLPixelFormatClass().Alloc()
	rv := objc.Send[OpenGLPixelFormat](instance.ID, objc.Sel("initWithCGLPixelFormatObj:"), format)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewOpenGLPixelFormatWithCGLPixelFormatObj */


// Returns an OpenGL pixel format object initialized with specified pixel format attribute data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLPixelFormat/initWithData:
func NewOpenGLPixelFormatWithData(attribs objc.IObject /* cross-framework: NSData */) OpenGLPixelFormat {
	instance := getOpenGLPixelFormatClass().Alloc()
	rv := objc.Send[OpenGLPixelFormat](instance.ID, objc.Sel("initWithData:"), attribs)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewOpenGLPixelFormatWithData */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for OpenGLPixelFormat */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for OpenGLPixelFormat */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for OpenGLPixelFormat */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for OpenGLPixelFormat */

// The low-level, platform-specific Core OpenGL (CGL) pixel format object represented by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLPixelFormat/cglPixelFormatObj
func (o_ OpenGLPixelFormat) CGLPixelFormatObj() LPixelFormatObj /* not a class type */ {
	rv := objc.Send[LPixelFormatObj](o_.ID, objc.Sel("CGLPixelFormatObj"))
	return rv
}/* debug [instance_properties/getter]: CGLPixelFormatObj */


// The number of virtual screens associated with the OpenGL pixel format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLPixelFormat/numberOfVirtualScreens
func (o_ OpenGLPixelFormat) NumberOfVirtualScreens() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](o_.ID, objc.Sel("numberOfVirtualScreens"))
	return rv
}/* debug [instance_properties/getter]: numberOfVirtualScreens */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSOpenGLPixelFormat */


