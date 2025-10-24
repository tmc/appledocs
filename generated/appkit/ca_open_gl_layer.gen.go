// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CAOpenGLLayer */


/* debug [class_header]: Header for CAOpenGLLayer */
// The class instance for the [OpenGLLayer] class.
var (
	OpenGLLayerClass     _OpenGLLayerClass
	OpenGLLayerClassOnce sync.Once
)

func getOpenGLLayerClass() _OpenGLLayerClass {
	OpenGLLayerClassOnce.Do(func() {
		OpenGLLayerClass = _OpenGLLayerClass{objc.GetClass("CAOpenGLLayer")}
	})
	return OpenGLLayerClass
}

type _OpenGLLayerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for OpenGLLayer */
// An interface definition for the [OpenGLLayer] class.
type IOpenGLLayer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for OpenGLLayer */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for OpenGLLayer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for OpenGLLayer */
// Alloc allocates a new instance without initialization.
func (oc _OpenGLLayerClass) Alloc() OpenGLLayer {
	rv := objc.Send[OpenGLLayer](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (oc _OpenGLLayerClass) New() OpenGLLayer {
	rv := objc.Send[OpenGLLayer](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OpenGLLayer) Init() OpenGLLayer {
	rv := objc.Send[OpenGLLayer](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OpenGLLayer) Autorelease() OpenGLLayer {
	rv := objc.Send[OpenGLLayer](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOpenGLLayer creates a new OpenGLLayer instance.
func NewOpenGLLayer() OpenGLLayer {
	return getOpenGLLayerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for OpenGLLayer */
// A parent class referenced by other AppKit classes.


// A parent class referenced by other AppKit classes. [Full Topic]
type OpenGLLayer struct {
	objectivec.Object
}

// OpenGLLayerFrom constructs a [OpenGLLayer] from an unsafe.Pointer.
//
// A parent class referenced by other AppKit classes.
func OpenGLLayerFrom(ptr unsafe.Pointer) OpenGLLayer {
	return OpenGLLayer{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for OpenGLLayer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for OpenGLLayer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for OpenGLLayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for OpenGLLayer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for OpenGLLayer */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CAOpenGLLayer */



