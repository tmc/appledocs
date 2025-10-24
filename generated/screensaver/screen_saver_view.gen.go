// Code generated from Apple documentation for ScreenSaver. DO NOT EDIT.

package screensaver

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

/* debug [class.gen.go]: Generating class ScreenSaverView */


/* debug [class_header]: Header for ScreenSaverView */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ScreenSaverView */
// An interface definition for the [ScreenSaverView] class.
type IScreenSaverView interface {
	appkit.IView
	
/* debug [class_interface_properties]: Properties for ScreenSaverView */
	// properties:
	AnimationTimeInterval() float64
	SetAnimationTimeInterval(value float64)
	ConfigureSheet() appkit.Window
	HasConfigureSheet() bool
	Animating() bool
	Preview() bool
	IsAnimating() bool
	SetIsAnimating(value bool)
	IsPreview() bool
	SetIsPreview(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ScreenSaverView */
	// methods:
	AnimateOneFrame()
	DrawRect(rect Rect /* not a class type */)
	StartAnimation()
	StopAnimation()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ScreenSaverView */
// Alloc allocates a new instance without initialization.
func (sc _ScreenSaverViewClass) Alloc() ScreenSaverView {
	rv := objc.Send[ScreenSaverView](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ScreenSaverView */
// An abstract class that defines the interface for subclassers to interact with the screen saver infrastructure.
//
// provides the interface for your screen saver, including the content you animate onscreen and an optional configuration sheet. Create your own custom subclass and add it to your screen saver bundle. Use your subclass to create the animations that you want to appear onscreen, and to specify additional animation details. You can draw from your view’s method, or you can draw directly from the method. If you prefer to use the method, use the method to call the method and specify the portions of your view that require updates.


// An abstract class that defines the interface for subclassers to interact with the screen saver infrastructure.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ScreenSaverView */

// Creates a newly allocated screen saver view with the specified frame rectangle and preview information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenSaver/ScreenSaverView/init(frame:isPreview:)
func NewScreenSaverViewWithFrameIsPreview(frame Rect /* not a class type */, isPreview bool) ScreenSaverView {
	instance := getScreenSaverViewClass().Alloc()
	rv := objc.Send[ScreenSaverView](instance.ID, objc.Sel("initWithFrame:isPreview:"), frame, isPreview)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewScreenSaverViewWithFrameIsPreview */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ScreenSaverView */

// Returns the type of backing store you want for your screen saver’s window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenSaver/ScreenSaverView/backingStoreType()
func (sc _ScreenSaverViewClass) BackingStoreType() BackingStoreType /* not a class type */ {
	rv := objc.Send[BackingStoreType](objc.ID(sc.class), objc.Sel("backingStoreType"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=BackingStoreType) */


// Indicates whether to perform a gradual screen fade when the system starts and stops your screen saver’s animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenSaver/ScreenSaverView/performGammaFade()
func (sc _ScreenSaverViewClass) PerformGammaFade() bool {
	rv := objc.Send[bool](objc.ID(sc.class), objc.Sel("performGammaFade"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PerformGammaFade) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ScreenSaverView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ScreenSaverView */

// Advances the screen saver’s animation by a single frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenSaver/ScreenSaverView/animateOneFrame()
func (s_ ScreenSaverView) AnimateOneFrame() {
	objc.Send[objc.ID](s_.ID, objc.Sel("animateOneFrame"))
}/* debug [instance_methods/method]: AnimateOneFrame */


// Draws the screen saver view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenSaver/ScreenSaverView/draw(_:)
func (s_ ScreenSaverView) DrawRect(rect Rect /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("drawRect:"), rect)
}/* debug [instance_methods/method]: DrawRect */


// Activates the periodic timer that animates the screen saver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenSaver/ScreenSaverView/startAnimation()
func (s_ ScreenSaverView) StartAnimation() {
	objc.Send[objc.ID](s_.ID, objc.Sel("startAnimation"))
}/* debug [instance_methods/method]: StartAnimation */


// Deactivates the timer that advances the animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenSaver/ScreenSaverView/stopAnimation()
func (s_ ScreenSaverView) StopAnimation() {
	objc.Send[objc.ID](s_.ID, objc.Sel("stopAnimation"))
}/* debug [instance_methods/method]: StopAnimation */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ScreenSaverView */

// The time interval between animation frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenSaver/ScreenSaverView/animationTimeInterval
func (s_ ScreenSaverView) AnimationTimeInterval() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("animationTimeInterval"))
	return rv
}/* debug [instance_properties/getter]: animationTimeInterval */


// The time interval between animation frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenSaver/ScreenSaverView/animationTimeInterval
func (s_ ScreenSaverView) SetAnimationTimeInterval(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAnimationTimeInterval:"), value)
}/* debug [instance_properties/setter]: animationTimeInterval */


// The window that contains the controls to configure the screen saver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenSaver/ScreenSaverView/configureSheet
func (s_ ScreenSaverView) ConfigureSheet() appkit.Window {
	rv := objc.Send[appkit.Window](s_.ID, objc.Sel("configureSheet"))
	return rv
}/* debug [instance_properties/getter]: configureSheet */


// A Boolean value that indicates whether the screen saver has an associated configuration sheet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenSaver/ScreenSaverView/hasConfigureSheet
func (s_ ScreenSaverView) HasConfigureSheet() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("hasConfigureSheet"))
	return rv
}/* debug [instance_properties/getter]: hasConfigureSheet */


// A Boolean value that indicates whether the screen saver is animating.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenSaver/ScreenSaverView/isAnimating
func (s_ ScreenSaverView) Animating() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("animating"))
	return rv
}/* debug [instance_properties/getter]: animating */


// A Boolean value that indicates whether the screen saver view is set to a size suitable for previewing its content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenSaver/ScreenSaverView/isPreview
func (s_ ScreenSaverView) Preview() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("preview"))
	return rv
}/* debug [instance_properties/getter]: preview */


// A Boolean value that indicates whether the screen saver is animating.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screensaver/screensaverview/isanimating
func (s_ ScreenSaverView) IsAnimating() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isAnimating"))
	return rv
}/* debug [instance_properties/getter]: isAnimating */


// A Boolean value that indicates whether the screen saver is animating.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screensaver/screensaverview/isanimating
func (s_ ScreenSaverView) SetIsAnimating(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsAnimating:"), value)
}/* debug [instance_properties/setter]: isAnimating */


// A Boolean value that indicates whether the screen saver view is set to a size suitable for previewing its content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screensaver/screensaverview/ispreview
func (s_ ScreenSaverView) IsPreview() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isPreview"))
	return rv
}/* debug [instance_properties/getter]: isPreview */


// A Boolean value that indicates whether the screen saver view is set to a size suitable for previewing its content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screensaver/screensaverview/ispreview
func (s_ ScreenSaverView) SetIsPreview(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsPreview:"), value)
}/* debug [instance_properties/setter]: isPreview */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ScreenSaverView */


