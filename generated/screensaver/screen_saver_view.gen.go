// Code generated from Apple documentation for ScreenSaver. DO NOT EDIT.

package screensaver

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [ScreenSaverView] class.
var (
	ScreenSaverViewClass     _ScreenSaverViewClass
	ScreenSaverViewClassOnce sync.Once
)

func getScreenSaverViewClass() _ScreenSaverViewClass {
	ScreenSaverViewClassOnce.Do(func() {
		ScreenSaverViewClass = _ScreenSaverViewClass{objc.GetClass("ScreenSaverView")}
	})
	return ScreenSaverViewClass
}

type _ScreenSaverViewClass struct {
	class objc.Class
}

// An interface definition for the [ScreenSaverView] class.
type IScreenSaverView interface {
	appkit.IView
	AnimateOneFrame()
	DrawRect(rect coregraphics.CGRect)
	StartAnimation()
	StopAnimation()
	AnimationTimeInterval() foundation.TimeInterval
	SetAnimationTimeInterval(value foundation.ITimeInterval)
	ConfigureSheet() appkit.Window
	HasConfigureSheet() bool
	Animating() bool
	Preview() bool
	IsAnimating() bool
	SetIsAnimating(value bool)
	IsPreview() bool
	SetIsPreview(value bool)
}

// An abstract class that defines the interface for subclassers to interact with the screen saver infrastructure.
//
// provides the interface for your screen saver, including the content you animate onscreen and an optional configuration sheet. Create your own custom subclass and add it to your screen saver bundle. Use your subclass to create the animations that you want to appear onscreen, and to specify additional animation details. You can draw from your view’s method, or you can draw directly from the method. If you prefer to use the method, use the method to call the method and specify the portions of your view that require updates.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenSaver/ScreenSaverView
type ScreenSaverView struct {
	appkit.View
}

// ScreenSaverViewFrom constructs a [ScreenSaverView] from an unsafe.Pointer.
//
// An abstract class that defines the interface for subclassers to interact with the screen saver infrastructure.
func ScreenSaverViewFrom(ptr unsafe.Pointer) ScreenSaverView {
	return ScreenSaverView{
		View: appkit.ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _ScreenSaverViewClass) Alloc() ScreenSaverView {
	rv := objc.Send[ScreenSaverView](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _ScreenSaverViewClass) New() ScreenSaverView {
	rv := objc.Send[ScreenSaverView](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScreenSaverView) Init() ScreenSaverView {
	rv := objc.Send[ScreenSaverView](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScreenSaverView) Autorelease() ScreenSaverView {
	rv := objc.Send[ScreenSaverView](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScreenSaverView creates a new ScreenSaverView instance.
func NewScreenSaverView() ScreenSaverView {
	return getScreenSaverViewClass().New()
}




// Creates a newly allocated screen saver view with the specified frame rectangle and preview information.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenSaver/ScreenSaverView/init(frame:isPreview:)
func NewScreenSaverViewWithFrameIsPreview(frame coregraphics.CGRect, isPreview bool) ScreenSaverView {
	instance := getScreenSaverViewClass().Alloc()
	rv := objc.Send[ScreenSaverView](instance.ID, objc.Sel("initWithFrame:isPreview:"), frame, isPreview)
	rv.Autorelease()
	return rv
}


// Returns the type of backing store you want for your screen saver’s window.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenSaver/ScreenSaverView/backingStoreType()
func (sc _ScreenSaverViewClass) BackingStoreType() appkit.BackingStoreType {
	rv := objc.Send[appkit.BackingStoreType](objc.ID(sc.class), objc.Sel("backingStoreType"))
	return rv
}

// Indicates whether to perform a gradual screen fade when the system starts and stops your screen saver’s animation.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenSaver/ScreenSaverView/performGammaFade()
func (sc _ScreenSaverViewClass) PerformGammaFade() bool {
	rv := objc.Send[bool](objc.ID(sc.class), objc.Sel("performGammaFade"))
	return rv
}

// Advances the screen saver’s animation by a single frame.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenSaver/ScreenSaverView/animateOneFrame()
func (s_ ScreenSaverView) AnimateOneFrame() {
	objc.Send[objc.ID](s_.ID, objc.Sel("animateOneFrame"))
}

// Draws the screen saver view.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenSaver/ScreenSaverView/draw(_:)
func (s_ ScreenSaverView) DrawRect(rect coregraphics.CGRect) {
	objc.Send[objc.ID](s_.ID, objc.Sel("drawRect:"), rect)
}

// Activates the periodic timer that animates the screen saver.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenSaver/ScreenSaverView/startAnimation()
func (s_ ScreenSaverView) StartAnimation() {
	objc.Send[objc.ID](s_.ID, objc.Sel("startAnimation"))
}

// Deactivates the timer that advances the animation.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenSaver/ScreenSaverView/stopAnimation()
func (s_ ScreenSaverView) StopAnimation() {
	objc.Send[objc.ID](s_.ID, objc.Sel("stopAnimation"))
}

// The time interval between animation frames.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenSaver/ScreenSaverView/animationTimeInterval
func (s_ ScreenSaverView) AnimationTimeInterval() foundation.TimeInterval {
	rv := objc.Send[foundation.TimeInterval](s_.ID, objc.Sel("animationTimeInterval"))
	return rv
}


// SetAnimationTimeInterval sets the value of the animationTimeInterval property.
// The time interval between animation frames.

//
// [Full Topic]: https://developer.apple.com/documentation/ScreenSaver/ScreenSaverView/animationTimeInterval
func (s_ ScreenSaverView) SetAnimationTimeInterval(value foundation.ITimeInterval) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAnimationTimeInterval:"), value)
}

// The window that contains the controls to configure the screen saver.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenSaver/ScreenSaverView/configureSheet
func (s_ ScreenSaverView) ConfigureSheet() appkit.Window {
	rv := objc.Send[appkit.Window](s_.ID, objc.Sel("configureSheet"))
	return rv
}

// A Boolean value that indicates whether the screen saver has an associated configuration sheet.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenSaver/ScreenSaverView/hasConfigureSheet
func (s_ ScreenSaverView) HasConfigureSheet() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("hasConfigureSheet"))
	return rv
}

// A Boolean value that indicates whether the screen saver is animating.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenSaver/ScreenSaverView/isAnimating
func (s_ ScreenSaverView) Animating() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("animating"))
	return rv
}

// A Boolean value that indicates whether the screen saver view is set to a size suitable for previewing its content.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenSaver/ScreenSaverView/isPreview
func (s_ ScreenSaverView) Preview() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("preview"))
	return rv
}

// A Boolean value that indicates whether the screen saver is animating.
//
// [Full Topic]: https://developer.apple.com/documentation/screensaver/screensaverview/isanimating
func (s_ ScreenSaverView) IsAnimating() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isAnimating"))
	return rv
}


// SetIsAnimating sets the value of the isAnimating property.
// A Boolean value that indicates whether the screen saver is animating.

//
// [Full Topic]: https://developer.apple.com/documentation/screensaver/screensaverview/isanimating
func (s_ ScreenSaverView) SetIsAnimating(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsAnimating:"), value)
}

// A Boolean value that indicates whether the screen saver view is set to a size suitable for previewing its content.
//
// [Full Topic]: https://developer.apple.com/documentation/screensaver/screensaverview/ispreview
func (s_ ScreenSaverView) IsPreview() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isPreview"))
	return rv
}


// SetIsPreview sets the value of the isPreview property.
// A Boolean value that indicates whether the screen saver view is set to a size suitable for previewing its content.

//
// [Full Topic]: https://developer.apple.com/documentation/screensaver/screensaverview/ispreview
func (s_ ScreenSaverView) SetIsPreview(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsPreview:"), value)
}


