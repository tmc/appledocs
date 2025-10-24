// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class QCRenderer */


/* debug [class_header]: Header for QCRenderer */
// The class instance for the [QCRenderer] class.
var (
	QCRendererClass     _QCRendererClass
	QCRendererClassOnce sync.Once
)

func getQCRendererClass() _QCRendererClass {
	QCRendererClassOnce.Do(func() {
		QCRendererClass = _QCRendererClass{objc.GetClass("QCRenderer")}
	})
	return QCRendererClass
}

type _QCRendererClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for QCRenderer */
// An interface definition for the [QCRenderer] class.
type IQCRenderer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for QCRenderer */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for QCRenderer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for QCRenderer */
// Alloc allocates a new instance without initialization.
func (qc _QCRendererClass) Alloc() QCRenderer {
	rv := objc.Send[QCRenderer](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (qc _QCRendererClass) New() QCRenderer {
	rv := objc.Send[QCRenderer](objc.ID(qc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (q_ QCRenderer) Init() QCRenderer {
	rv := objc.Send[QCRenderer](q_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (q_ QCRenderer) Autorelease() QCRenderer {
	rv := objc.Send[QCRenderer](q_.ID, objc.Sel("autorelease"))
	return rv
}

// NewQCRenderer creates a new QCRenderer instance.
func NewQCRenderer() QCRenderer {
	return getQCRendererClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for QCRenderer */
// A base class for low-level rendering.
//
// A class is designed for low-level rendering of Quartz Composer compositions. This is the class to use if you want to be in charge of rendering a composition to a specific OpenGL context—either using the class or a object. also allows you to load, play, and control a composition. To render a composition to a specific OpenGL context: Create an instance of using one of the initialization methods, such as . Render frames by calling the method If you use double buffering in OpenGL, you must swap the OpenGL buffers. Release the renderer when you no longer need it. This code snippet shows how to implement these tasks:


// A base class for low-level rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCRenderer
type QCRenderer struct {
	objectivec.Object
}

// QCRendererFrom constructs a [QCRenderer] from an unsafe.Pointer.
//
// A base class for low-level rendering.
func QCRendererFrom(ptr unsafe.Pointer) QCRenderer {
	return QCRenderer{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for QCRenderer */

// Creates an offscreen renderer of a given size with the provided color space and composition object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCRenderer/init(offScreenWith:colorSpace:composition:)
func NewQCRendererOffScreenWithSizeColorSpaceComposition(size Size /* not a class type */, colorSpace ColorSpaceRef /* not a class type */, composition IQCComposition) QCRenderer {
	instance := getQCRendererClass().Alloc()
	rv := objc.Send[QCRenderer](instance.ID, objc.Sel("initOffScreenWithSize:colorSpace:composition:"), size, colorSpace, composition)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewQCRendererOffScreenWithSizeColorSpaceComposition */


// Creates a renderer object with a object, a pixel format, a color space, and a composition object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCRenderer/init(cglContext:pixelFormat:colorSpace:composition:)
func NewQCRendererWithCGLContextPixelFormatColorSpaceComposition(context LContextObj /* not a class type */, format LPixelFormatObj /* not a class type */, colorSpace ColorSpaceRef /* not a class type */, composition IQCComposition) QCRenderer {
	instance := getQCRendererClass().Alloc()
	rv := objc.Send[QCRenderer](instance.ID, objc.Sel("initWithCGLContext:pixelFormat:colorSpace:composition:"), context, format, colorSpace, composition)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewQCRendererWithCGLContextPixelFormatColorSpaceComposition */


// Creates a renderer object with a composition object and a color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCRenderer/init(composition:colorSpace:)
func NewQCRendererWithCompositionColorSpace(composition IQCComposition, colorSpace ColorSpaceRef /* not a class type */) QCRenderer {
	instance := getQCRendererClass().Alloc()
	rv := objc.Send[QCRenderer](instance.ID, objc.Sel("initWithComposition:colorSpace:"), composition, colorSpace)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewQCRendererWithCompositionColorSpace */


// Creates a renderer object with an object and a composition file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCRenderer/init(openGLContext:pixelFormat:file:)
func NewQCRendererWithOpenGLContextPixelFormatFile(context appkit.OpenGLContext, format appkit.OpenGLPixelFormat, path objc.IObject /* cross-framework: NSString */) QCRenderer {
	instance := getQCRendererClass().Alloc()
	rv := objc.Send[QCRenderer](instance.ID, objc.Sel("initWithOpenGLContext:pixelFormat:file:"), context, format, path)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewQCRendererWithOpenGLContextPixelFormatFile */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for QCRenderer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for QCRenderer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for QCRenderer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for QCRenderer */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class QCRenderer */


