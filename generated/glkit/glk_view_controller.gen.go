// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
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
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	FramesDisplayed() int
	FramesPerSecond() int
	Paused() bool
	SetPaused(value bool)
	PauseOnWillResignActive() bool
	SetPauseOnWillResignActive(value bool)
	PreferredFramesPerSecond() int
	SetPreferredFramesPerSecond(value int)
	ResumeOnDidBecomeActive() bool
	SetResumeOnDidBecomeActive(value bool)
	TimeSinceFirstResume() foundation.TimeInterval
	TimeSinceLastDraw() foundation.TimeInterval
	TimeSinceLastResume() foundation.TimeInterval
	TimeSinceLastUpdate() foundation.TimeInterval
	IsPaused() bool
	SetIsPaused(value bool)
}

// A view controller that manages an OpenGL ES rendering loop.
//
// A object works in conjunction with a object to display frames of animation in the view, and also provides standard view controller functionality. To use this class, allocate and initialize a new subclass and set its property to point to a object. Then, configure the view controller’s property to the desired frame rate your application requires. You can set a delegate or configure other properties on the view controller, such as whether the animation loop is automatically paused or resumed when the application moves into the background. When active, rendering loop automatically updates the view’s contents each time a new frame must be displayed. Each frame is rendered by the view controller using these steps: The view controller calls its delegate’s method. Your delegate should update frame data that does not involve rendering the results to the screen. The view controller calls its view’s method. Your view should redraw the frame.
//
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


// The view controller’s delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKViewController/delegate
func (g_ GLKViewController) Delegate() objc.ID {
	rv := objc.Send[objc.ID](g_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The view controller’s delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKViewController/delegate
func (g_ GLKViewController) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDelegate:"), value)
}

// The number of frame updates that have been sent by the view controller since it was created.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKViewController/framesDisplayed
func (g_ GLKViewController) FramesDisplayed() int {
	rv := objc.Send[int](g_.ID, objc.Sel("framesDisplayed"))
	return rv
}

// The actual rate that the view controller attempts to call the view to update its contents.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKViewController/framesPerSecond
func (g_ GLKViewController) FramesPerSecond() int {
	rv := objc.Send[int](g_.ID, objc.Sel("framesPerSecond"))
	return rv
}

// A Boolean value that indicates whether the rendering loop is paused.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKViewController/isPaused
func (g_ GLKViewController) Paused() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("paused"))
	return rv
}


// SetPaused sets the value of the paused property.
// A Boolean value that indicates whether the rendering loop is paused.

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKViewController/isPaused
func (g_ GLKViewController) SetPaused(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPaused:"), value)
}

// A Boolean value that indicates whether the view controller automatically pauses the rendering loop when the application resigns the active state.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKViewController/pauseOnWillResignActive
func (g_ GLKViewController) PauseOnWillResignActive() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("pauseOnWillResignActive"))
	return rv
}


// SetPauseOnWillResignActive sets the value of the pauseOnWillResignActive property.
// A Boolean value that indicates whether the view controller automatically pauses the rendering loop when the application resigns the active state.

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKViewController/pauseOnWillResignActive
func (g_ GLKViewController) SetPauseOnWillResignActive(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPauseOnWillResignActive:"), value)
}

// The rate you want the view controller to call the view to update the contents of the view.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKViewController/preferredFramesPerSecond
func (g_ GLKViewController) PreferredFramesPerSecond() int {
	rv := objc.Send[int](g_.ID, objc.Sel("preferredFramesPerSecond"))
	return rv
}


// SetPreferredFramesPerSecond sets the value of the preferredFramesPerSecond property.
// The rate you want the view controller to call the view to update the contents of the view.

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKViewController/preferredFramesPerSecond
func (g_ GLKViewController) SetPreferredFramesPerSecond(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPreferredFramesPerSecond:"), value)
}

// A Boolean value that indicates whether the view controller automatically resumes the rendering loop when the application becomes active.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKViewController/resumeOnDidBecomeActive
func (g_ GLKViewController) ResumeOnDidBecomeActive() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("resumeOnDidBecomeActive"))
	return rv
}


// SetResumeOnDidBecomeActive sets the value of the resumeOnDidBecomeActive property.
// A Boolean value that indicates whether the view controller automatically resumes the rendering loop when the application becomes active.

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKViewController/resumeOnDidBecomeActive
func (g_ GLKViewController) SetResumeOnDidBecomeActive(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setResumeOnDidBecomeActive:"), value)
}

// The amount of time that has passed since first time the view controller resumed sending update events.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKViewController/timeSinceFirstResume
func (g_ GLKViewController) TimeSinceFirstResume() foundation.TimeInterval {
	rv := objc.Send[foundation.TimeInterval](g_.ID, objc.Sel("timeSinceFirstResume"))
	return rv
}

// The amount of time that has passed since the last time the view controller called the view’s method.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKViewController/timeSinceLastDraw
func (g_ GLKViewController) TimeSinceLastDraw() foundation.TimeInterval {
	rv := objc.Send[foundation.TimeInterval](g_.ID, objc.Sel("timeSinceLastDraw"))
	return rv
}

// The amount of time that has passed since the last time the view controller resumed sending update events.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKViewController/timeSinceLastResume
func (g_ GLKViewController) TimeSinceLastResume() foundation.TimeInterval {
	rv := objc.Send[foundation.TimeInterval](g_.ID, objc.Sel("timeSinceLastResume"))
	return rv
}

// The amount of time that has passed since the last time the view controller called the delegate’s method.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKViewController/timeSinceLastUpdate
func (g_ GLKViewController) TimeSinceLastUpdate() foundation.TimeInterval {
	rv := objc.Send[foundation.TimeInterval](g_.ID, objc.Sel("timeSinceLastUpdate"))
	return rv
}

// A Boolean value that indicates whether the rendering loop is paused.
//
// [Full Topic]: https://developer.apple.com/documentation/glkit/glkviewcontroller/ispaused
func (g_ GLKViewController) IsPaused() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("isPaused"))
	return rv
}


// SetIsPaused sets the value of the isPaused property.
// A Boolean value that indicates whether the rendering loop is paused.

//
// [Full Topic]: https://developer.apple.com/documentation/glkit/glkviewcontroller/ispaused
func (g_ GLKViewController) SetIsPaused(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIsPaused:"), value)
}




