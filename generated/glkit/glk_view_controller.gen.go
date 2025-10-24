// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

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

// An interface definition for the [GLKViewController] class.
type IGLKViewController interface {
	appkit.IViewController
	// properties:
	IsPaused() bool
	SetIsPaused(value bool)
	// methods:
}

// A view controller that manages an OpenGL ES rendering loop.
//
// A object works in conjunction with a object to display frames of animation in the view, and also provides standard view controller functionality. To use this class, allocate and initialize a new subclass and set its property to point to a object. Then, configure the view controller’s property to the desired frame rate your application requires. You can set a delegate or configure other properties on the view controller, such as whether the animation loop is automatically paused or resumed when the application moves into the background. When active, rendering loop automatically updates the view’s contents each time a new frame must be displayed. Each frame is rendered by the view controller using these steps: The view controller calls its delegate’s method. Your delegate should update frame data that does not involve rendering the results to the screen. The view controller calls its view’s method. Your view should redraw the frame.


// A view controller that manages an OpenGL ES rendering loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKViewController
type GLKViewController struct {
	appkit.ViewController
}

// GLKViewControllerFrom constructs a [GLKViewController] from an unsafe.Pointer.
//
// A view controller that manages an OpenGL ES rendering loop.
func GLKViewControllerFrom(ptr unsafe.Pointer) GLKViewController {
	return GLKViewController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GLKViewControllerClass) Alloc() GLKViewController {
	rv := objc.Send[GLKViewController](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// A Boolean value that indicates whether the rendering loop is paused.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/glkit/glkviewcontroller/ispaused
func (g_ GLKViewController) IsPaused() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("isPaused"))
	return rv
}


// A Boolean value that indicates whether the rendering loop is paused.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/glkit/glkviewcontroller/ispaused
func (g_ GLKViewController) SetIsPaused(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIsPaused:"), value)
}


