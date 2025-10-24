// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class GLKViewController */


/* debug [class_header]: Header for GLKViewController */
// The class instance for the [GLKViewController] class.
var (
	GLKViewControllerClass     _GLKViewControllerClass
	GLKViewControllerClassOnce sync.Once
)

func getGLKViewControllerClass() _GLKViewControllerClass {
	GLKViewControllerClassOnce.Do(func() {
		GLKViewControllerClass = _GLKViewControllerClass{objc.GetClass("GLKViewController")}
	})
	return GLKViewControllerClass
}

type _GLKViewControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GLKViewController */
// An interface definition for the [GLKViewController] class.
type IGLKViewController interface {
	IViewController
	
/* debug [class_interface_properties]: Properties for GLKViewController */
	// properties:
	IsPaused() bool
	SetIsPaused(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GLKViewController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GLKViewController */
// Alloc allocates a new instance without initialization.
func (gc _GLKViewControllerClass) Alloc() GLKViewController {
	rv := objc.Send[GLKViewController](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GLKViewControllerClass) New() GLKViewController {
	rv := objc.Send[GLKViewController](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GLKViewController) Init() GLKViewController {
	rv := objc.Send[GLKViewController](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GLKViewController) Autorelease() GLKViewController {
	rv := objc.Send[GLKViewController](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGLKViewController creates a new GLKViewController instance.
func NewGLKViewController() GLKViewController {
	return getGLKViewControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GLKViewController */
// A view controller that manages an OpenGL ES rendering loop.
//
// A object works in conjunction with a object to display frames of animation in the view, and also provides standard view controller functionality. To use this class, allocate and initialize a new subclass and set its property to point to a object. Then, configure the view controller’s property to the desired frame rate your application requires. You can set a delegate or configure other properties on the view controller, such as whether the animation loop is automatically paused or resumed when the application moves into the background. When active, rendering loop automatically updates the view’s contents each time a new frame must be displayed. Each frame is rendered by the view controller using these steps: The view controller calls its delegate’s method. Your delegate should update frame data that does not involve rendering the results to the screen. The view controller calls its view’s method. Your view should redraw the frame.


// A view controller that manages an OpenGL ES rendering loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKViewController
type GLKViewController struct {
	ViewController
}

// GLKViewControllerFrom constructs a [GLKViewController] from an unsafe.Pointer.
//
// A view controller that manages an OpenGL ES rendering loop.
func GLKViewControllerFrom(ptr unsafe.Pointer) GLKViewController {
	return GLKViewController{
		ViewController: ViewControllerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GLKViewController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GLKViewController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GLKViewController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GLKViewController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GLKViewController */

// A Boolean value that indicates whether the rendering loop is paused.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/glkit/glkviewcontroller/ispaused
func (g_ GLKViewController) IsPaused() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("isPaused"))
	return rv
}/* debug [instance_properties/getter]: isPaused */


// A Boolean value that indicates whether the rendering loop is paused.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/glkit/glkviewcontroller/ispaused
func (g_ GLKViewController) SetIsPaused(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIsPaused:"), value)
}/* debug [instance_properties/setter]: isPaused */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GLKViewController */


