// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/quartzcore"
)

// The class instance for the [QCCompositionLayer] class.
var (
	QCCompositionLayerClass     _QCCompositionLayerClass
	QCCompositionLayerClassOnce sync.Once
)

func getQCCompositionLayerClass() _QCCompositionLayerClass {
	QCCompositionLayerClassOnce.Do(func() {
		QCCompositionLayerClass = _QCCompositionLayerClass{objc.GetClass("QCCompositionLayer")}
	})
	return QCCompositionLayerClass
}

type _QCCompositionLayerClass struct {
	class objc.Class
}

// An interface definition for the [QCCompositionLayer] class.
type IQCCompositionLayer interface {
	quartzcore.IOpenGLLayer
	Composition() QCComposition
}

// A layer that loads, plays, and controls Quartz Composer compositions in a Core Animation layer hierarchy.
//
// The composition tracks the Core Animation layer time and is rendered directly at the current dimensions of the object. An archived object saves the composition that’s loaded at the time the layer is archived. It detects layer usage and pauses or resumes the composition appropriately. A object starts rendering the composition automatically when the layer is placed in a visible layer hierarchy. The layer stops rendering when it is hidden or removed from the visible layer hierarchy. You can pass data to the input ports, or retrieve data from the output ports, of the root patch of a composition by accessing the attribute of the instance using methods provided by the protocol.
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCCompositionLayer
type QCCompositionLayer struct {
	quartzcore.OpenGLLayer
}

// QCCompositionLayerFrom constructs a [QCCompositionLayer] from an unsafe.Pointer.
//
// A layer that loads, plays, and controls Quartz Composer compositions in a Core Animation layer hierarchy.
func QCCompositionLayerFrom(ptr unsafe.Pointer) QCCompositionLayer {
	return QCCompositionLayer{
		OpenGLLayer: quartzcore.OpenGLLayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (qc _QCCompositionLayerClass) Alloc() QCCompositionLayer {
	rv := objc.Send[QCCompositionLayer](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (qc _QCCompositionLayerClass) New() QCCompositionLayer {
	rv := objc.Send[QCCompositionLayer](objc.ID(qc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (q_ QCCompositionLayer) Init() QCCompositionLayer {
	rv := objc.Send[QCCompositionLayer](q_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (q_ QCCompositionLayer) Autorelease() QCCompositionLayer {
	rv := objc.Send[QCCompositionLayer](q_.ID, objc.Sel("autorelease"))
	return rv
}

// NewQCCompositionLayer creates a new QCCompositionLayer instance.
func NewQCCompositionLayer() QCCompositionLayer {
	return getQCCompositionLayerClass().New()
}




// Initializes and returns a composition layer using the provided Quartz Composer composition.
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCCompositionLayer/init(composition:)
func NewQCCompositionLayerWithComposition(composition IQCComposition) QCCompositionLayer {
	instance := getQCCompositionLayerClass().Alloc()
	rv := objc.Send[QCCompositionLayer](instance.ID, objc.Sel("initWithComposition:"), composition)
	rv.Autorelease()
	return rv
}



// Initializes and returns a composition layer using the Quartz Composer composition in the specified file.
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCCompositionLayer/init(file:)
func NewQCCompositionLayerWithFile(path appkit.string) QCCompositionLayer {
	instance := getQCCompositionLayerClass().Alloc()
	rv := objc.Send[QCCompositionLayer](instance.ID, objc.Sel("initWithFile:"), path)
	rv.Autorelease()
	return rv
}


// Creates and returns an instance of a composition layer using the provided Quartz Composer composition.
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCCompositionLayer/compositionLayerWithComposition:
func (qc _QCCompositionLayerClass) CompositionLayerWithComposition(composition IQCComposition) QCCompositionLayer {
	rv := objc.Send[QCCompositionLayer](objc.ID(qc.class), objc.Sel("compositionLayerWithComposition:"), composition)
	return rv
}

// Creates and returns an instance of a composition layer using the Quartz Composer composition in the specified file.
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCCompositionLayer/compositionLayerWithFile:
func (qc _QCCompositionLayerClass) CompositionLayerWithFile(path appkit.string) QCCompositionLayer {
	rv := objc.Send[QCCompositionLayer](objc.ID(qc.class), objc.Sel("compositionLayerWithFile:"), path)
	return rv
}

// Returns the composition associated with the layer.
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCCompositionLayer/composition()
func (q_ QCCompositionLayer) Composition() QCComposition {
	rv := objc.Send[QCComposition](q_.ID, objc.Sel("composition"))
	return rv
}


