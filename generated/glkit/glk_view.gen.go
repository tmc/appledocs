// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/corefoundation"
)

/* debug [class.gen.go]: Generating class GLKView */


/* debug [class_header]: Header for GLKView */
// The class instance for the [GLKView] class.
var (
	GLKViewClass     _GLKViewClass
	GLKViewClassOnce sync.Once
)

func getGLKViewClass() _GLKViewClass {
	GLKViewClassOnce.Do(func() {
		GLKViewClass = _GLKViewClass{objc.GetClass("GLKView")}
	})
	return GLKViewClass
}

type _GLKViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GLKView */
// An interface definition for the [GLKView] class.
type IGLKView interface {
	IView
	
/* debug [class_interface_properties]: Properties for GLKView */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GLKView */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GLKView */
// Alloc allocates a new instance without initialization.
func (gc _GLKViewClass) Alloc() GLKView {
	rv := objc.Send[GLKView](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GLKViewClass) New() GLKView {
	rv := objc.Send[GLKView](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GLKView) Init() GLKView {
	rv := objc.Send[GLKView](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GLKView) Autorelease() GLKView {
	rv := objc.Send[GLKView](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGLKView creates a new GLKView instance.
func NewGLKView() GLKView {
	return getGLKViewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GLKView */
// A default implementation for views that draw their content using OpenGL ES.
//
// The class simplifies the effort required to create an OpenGL ES application by directly managing a framebuffer object on your behalf; your application simply needs to draw into the framebuffer when the contents need to be updated. To use this class in your application, create a new object and provide it an OpenGL ES context. Then, modify the view’s , , , and properties to configure the format of the drawable’s framebuffer object. After this, the view automatically creates or updates the framebuffer object whenever the view must be redrawn. A object uses the regular view drawing cycle for a object, calling its method whenever the contents of the view need to be updated. Before calling its method, the view makes its object the current OpenGL ES context and binds its framebuffer object to the OpenGL ES context as the target for rendering commands. Your application’s implementation of the method should call one or more OpenGL ES functions to render an image into the framebuffer object. Then, the view resolves any multisampling that you may have enabled and delivers the finished results. The class can be used in conjunction with a object to create an animation rendering loop that redraws the contents of the view at a specified frame rate.


// A default implementation for views that draw their content using OpenGL ES.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKView
type GLKView struct {
	View
}

// GLKViewFrom constructs a [GLKView] from an unsafe.Pointer.
//
// A default implementation for views that draw their content using OpenGL ES.
func GLKViewFrom(ptr unsafe.Pointer) GLKView {
	return GLKView{
		View: ViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GLKView */

// Initializes a new view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKView/init(frame:context:)
func NewGLKViewWithFrameContext(frame corefoundation.CGRect, context unsafe.Pointer) GLKView {
	instance := getGLKViewClass().Alloc()
	rv := objc.Send[GLKView](instance.ID, objc.Sel("initWithFrame:context:"), frame, context)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewGLKViewWithFrameContext */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GLKView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GLKView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GLKView */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GLKView */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GLKView */


