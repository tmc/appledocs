// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class QCCompositionLayer */


/* debug [class_header]: Header for QCCompositionLayer */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for QCCompositionLayer */
// An interface definition for the [QCCompositionLayer] class.
type IQCCompositionLayer interface {
	IOpenGLLayer
	
/* debug [class_interface_properties]: Properties for QCCompositionLayer */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for QCCompositionLayer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for QCCompositionLayer */
// Alloc allocates a new instance without initialization.
func (qc _QCCompositionLayerClass) Alloc() QCCompositionLayer {
	rv := objc.Send[QCCompositionLayer](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for QCCompositionLayer */
// A layer that loads, plays, and controls Quartz Composer compositions in a Core Animation layer hierarchy.
//
// The composition tracks the Core Animation layer time and is rendered directly at the current dimensions of the object. An archived object saves the composition that’s loaded at the time the layer is archived. It detects layer usage and pauses or resumes the composition appropriately. A object starts rendering the composition automatically when the layer is placed in a visible layer hierarchy. The layer stops rendering when it is hidden or removed from the visible layer hierarchy. You can pass data to the input ports, or retrieve data from the output ports, of the root patch of a composition by accessing the attribute of the instance using methods provided by the protocol.


// A layer that loads, plays, and controls Quartz Composer compositions in a Core Animation layer hierarchy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCCompositionLayer
type QCCompositionLayer struct {
	OpenGLLayer
}

// QCCompositionLayerFrom constructs a [QCCompositionLayer] from an unsafe.Pointer.
//
// A layer that loads, plays, and controls Quartz Composer compositions in a Core Animation layer hierarchy.
func QCCompositionLayerFrom(ptr unsafe.Pointer) QCCompositionLayer {
	return QCCompositionLayer{
		OpenGLLayer: OpenGLLayerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for QCCompositionLayer */

// Initializes and returns a composition layer using the provided Quartz Composer composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCCompositionLayer/init(composition:)
func NewQCCompositionLayerWithComposition(composition IQCComposition) QCCompositionLayer {
	instance := getQCCompositionLayerClass().Alloc()
	rv := objc.Send[QCCompositionLayer](instance.ID, objc.Sel("initWithComposition:"), composition)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewQCCompositionLayerWithComposition */


// Initializes and returns a composition layer using the Quartz Composer composition in the specified file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCCompositionLayer/init(file:)
func NewQCCompositionLayerWithFile(path objc.IObject /* cross-framework: NSString */) QCCompositionLayer {
	instance := getQCCompositionLayerClass().Alloc()
	rv := objc.Send[QCCompositionLayer](instance.ID, objc.Sel("initWithFile:"), path)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewQCCompositionLayerWithFile */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for QCCompositionLayer */

// Creates and returns an instance of a composition layer using the provided Quartz Composer composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCCompositionLayer/compositionLayerWithComposition:
func (qc _QCCompositionLayerClass) CompositionLayerWithComposition(composition IQCComposition) QCCompositionLayer {
	rv := objc.Send[QCCompositionLayer](objc.ID(qc.class), objc.Sel("compositionLayerWithComposition:"), composition)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CompositionLayerWithComposition) */


// Creates and returns an instance of a composition layer using the Quartz Composer composition in the specified file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCCompositionLayer/compositionLayerWithFile:
func (qc _QCCompositionLayerClass) CompositionLayerWithFile(path objc.IObject /* cross-framework: NSString */) QCCompositionLayer {
	rv := objc.Send[QCCompositionLayer](objc.ID(qc.class), objc.Sel("compositionLayerWithFile:"), path)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CompositionLayerWithFile) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for QCCompositionLayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for QCCompositionLayer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for QCCompositionLayer */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class QCCompositionLayer */


