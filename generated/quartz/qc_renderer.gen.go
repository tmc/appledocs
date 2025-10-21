// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [QCRenderer] class.
type IQCRenderer interface {
	objectivec.IObject
	Composition() QCComposition
	CreateSnapshotImageOfType(type_ appkit.string) objc.ID
	RenderAtTimeArguments(time foundation.ITimeInterval, arguments objectivec.IObject) bool
	RenderingTimeForTimeArguments(time foundation.ITimeInterval, arguments objectivec.IObject) foundation.TimeInterval
	SnapshotImage() appkit.Image
}

// A base class for low-level rendering.
//
// A class is designed for low-level rendering of Quartz Composer compositions. This is the class to use if you want to be in charge of rendering a composition to a specific OpenGL context—either using the class or a object. also allows you to load, play, and control a composition. To render a composition to a specific OpenGL context: Create an instance of using one of the initialization methods, such as . Render frames by calling the method If you use double buffering in OpenGL, you must swap the OpenGL buffers. Release the renderer when you no longer need it. This code snippet shows how to implement these tasks:
//
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

// Alloc allocates a new instance without initialization.
func (qc _QCRendererClass) Alloc() QCRenderer {
	rv := objc.Send[QCRenderer](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates an offscreen renderer of a given size with the provided color space and composition object.
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCRenderer/init(offScreenWith:colorSpace:composition:)
func NewQCRendererOffScreenWithSizeColorSpaceComposition(size foundation.ISize, colorSpace coregraphics.CGColorSpaceRef, composition IQCComposition) QCRenderer {
	instance := getQCRendererClass().Alloc()
	rv := objc.Send[QCRenderer](instance.ID, objc.Sel("initOffScreenWithSize:colorSpace:composition:"), size, colorSpace, composition)
	rv.Autorelease()
	return rv
}



// Creates a renderer object with a object, a pixel format, a color space, and a composition object.
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCRenderer/init(cglContext:pixelFormat:colorSpace:composition:)
func NewQCRendererWithCGLContextPixelFormatColorSpaceComposition(context unsafe.Pointer, format unsafe.Pointer, colorSpace coregraphics.CGColorSpaceRef, composition IQCComposition) QCRenderer {
	instance := getQCRendererClass().Alloc()
	rv := objc.Send[QCRenderer](instance.ID, objc.Sel("initWithCGLContext:pixelFormat:colorSpace:composition:"), context, format, colorSpace, composition)
	rv.Autorelease()
	return rv
}



// Creates a renderer object with a composition object and a color space.
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCRenderer/init(composition:colorSpace:)
func NewQCRendererWithCompositionColorSpace(composition IQCComposition, colorSpace coregraphics.CGColorSpaceRef) QCRenderer {
	instance := getQCRendererClass().Alloc()
	rv := objc.Send[QCRenderer](instance.ID, objc.Sel("initWithComposition:colorSpace:"), composition, colorSpace)
	rv.Autorelease()
	return rv
}



// Creates a renderer object with an object and a composition file.
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCRenderer/init(openGLContext:pixelFormat:file:)
func NewQCRendererWithOpenGLContextPixelFormatFile(context appkit.IOpenGLContext, format appkit.OpenGLPixelFormat, path appkit.string) QCRenderer {
	instance := getQCRendererClass().Alloc()
	rv := objc.Send[QCRenderer](instance.ID, objc.Sel("initWithOpenGLContext:pixelFormat:file:"), context, format, path)
	rv.Autorelease()
	return rv
}


// Returns the composition object associated with the renderer.
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCRenderer/composition()
func (q_ QCRenderer) Composition() QCComposition {
	rv := objc.Send[QCComposition](q_.ID, objc.Sel("composition"))
	return rv
}

// Returns the current image in the OpenGL context associated with the renderer, as an image object of the provided image type.
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCRenderer/createSnapshotImage(ofType:)
func (q_ QCRenderer) CreateSnapshotImageOfType(type_ appkit.string) objc.ID {
	rv := objc.Send[objc.ID](q_.ID, objc.Sel("createSnapshotImageOfType:"), type_)
	return rv
}

// Renders a frame of a composition at the specified time.
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCRenderer/render(atTime:arguments:)
func (q_ QCRenderer) RenderAtTimeArguments(time foundation.ITimeInterval, arguments objectivec.IObject) bool {
	rv := objc.Send[bool](q_.ID, objc.Sel("renderAtTime:arguments:"), time, arguments)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCRenderer/renderingTime(forTime:arguments:)
func (q_ QCRenderer) RenderingTimeForTimeArguments(time foundation.ITimeInterval, arguments objectivec.IObject) foundation.TimeInterval {
	rv := objc.Send[foundation.TimeInterval](q_.ID, objc.Sel("renderingTimeForTime:arguments:"), time, arguments)
	return rv
}

// Returns an object of the current image in the OpenGL context associated with the renderer.
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCRenderer/snapshotImage()
func (q_ QCRenderer) SnapshotImage() appkit.Image {
	rv := objc.Send[appkit.Image](q_.ID, objc.Sel("snapshotImage"))
	return rv
}


