//go:build darwin && ios

// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// iOS-only methods for GLKViewController


// iOS-only properties

// The view controller’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKViewController/delegate
func (g_ GLKViewController) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("delegate"))
	return rv
}
func (g_ GLKViewController) SetDelegate(value unsafe.Pointer) {
	g_.ID.Send(objc.RegisterName("setDelegate:"), value)
}

// The number of frame updates that have been sent by the view controller since it was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKViewController/framesDisplayed
func (g_ GLKViewController) FramesDisplayed() int {
	rv := objc.Send[int](g_.ID, objc.Sel("framesDisplayed"))
	return rv
}

// The actual rate that the view controller attempts to call the view to update its contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKViewController/framesPerSecond
func (g_ GLKViewController) FramesPerSecond() int {
	rv := objc.Send[int](g_.ID, objc.Sel("framesPerSecond"))
	return rv
}

// A Boolean value that indicates whether the rendering loop is paused.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKViewController/isPaused
func (g_ GLKViewController) Paused() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("paused"))
	return rv
}
func (g_ GLKViewController) SetPaused(value bool) {
	g_.ID.Send(objc.RegisterName("setPaused:"), value)
}

// A Boolean value that indicates whether the view controller automatically pauses the rendering loop when the application resigns the active state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKViewController/pauseOnWillResignActive
func (g_ GLKViewController) PauseOnWillResignActive() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("pauseOnWillResignActive"))
	return rv
}
func (g_ GLKViewController) SetPauseOnWillResignActive(value bool) {
	g_.ID.Send(objc.RegisterName("setPauseOnWillResignActive:"), value)
}

// The rate you want the view controller to call the view to update the contents of the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKViewController/preferredFramesPerSecond
func (g_ GLKViewController) PreferredFramesPerSecond() int {
	rv := objc.Send[int](g_.ID, objc.Sel("preferredFramesPerSecond"))
	return rv
}
func (g_ GLKViewController) SetPreferredFramesPerSecond(value int) {
	g_.ID.Send(objc.RegisterName("setPreferredFramesPerSecond:"), value)
}

// A Boolean value that indicates whether the view controller automatically resumes the rendering loop when the application becomes active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKViewController/resumeOnDidBecomeActive
func (g_ GLKViewController) ResumeOnDidBecomeActive() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("resumeOnDidBecomeActive"))
	return rv
}
func (g_ GLKViewController) SetResumeOnDidBecomeActive(value bool) {
	g_.ID.Send(objc.RegisterName("setResumeOnDidBecomeActive:"), value)
}

// The amount of time that has passed since first time the view controller resumed sending update events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKViewController/timeSinceFirstResume
func (g_ GLKViewController) TimeSinceFirstResume() float64 {
	rv := objc.Send[float64](g_.ID, objc.Sel("timeSinceFirstResume"))
	return rv
}

// The amount of time that has passed since the last time the view controller called the view’s method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKViewController/timeSinceLastDraw
func (g_ GLKViewController) TimeSinceLastDraw() float64 {
	rv := objc.Send[float64](g_.ID, objc.Sel("timeSinceLastDraw"))
	return rv
}

// The amount of time that has passed since the last time the view controller resumed sending update events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKViewController/timeSinceLastResume
func (g_ GLKViewController) TimeSinceLastResume() float64 {
	rv := objc.Send[float64](g_.ID, objc.Sel("timeSinceLastResume"))
	return rv
}

// The amount of time that has passed since the last time the view controller called the delegate’s method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKViewController/timeSinceLastUpdate
func (g_ GLKViewController) TimeSinceLastUpdate() float64 {
	rv := objc.Send[float64](g_.ID, objc.Sel("timeSinceLastUpdate"))
	return rv
}







