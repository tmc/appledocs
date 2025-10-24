// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSWritingToolsCoordinatorAnimationParameters */


/* debug [class_header]: Header for NSWritingToolsCoordinatorAnimationParameters */
// The class instance for the [WritingToolsCoordinatorAnimationParameters] class.
var (
	WritingToolsCoordinatorAnimationParametersClass     _WritingToolsCoordinatorAnimationParametersClass
	WritingToolsCoordinatorAnimationParametersClassOnce sync.Once
)

func getWritingToolsCoordinatorAnimationParametersClass() _WritingToolsCoordinatorAnimationParametersClass {
	WritingToolsCoordinatorAnimationParametersClassOnce.Do(func() {
		WritingToolsCoordinatorAnimationParametersClass = _WritingToolsCoordinatorAnimationParametersClass{objc.GetClass("NSWritingToolsCoordinatorAnimationParameters")}
	})
	return WritingToolsCoordinatorAnimationParametersClass
}

type _WritingToolsCoordinatorAnimationParametersClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for WritingToolsCoordinatorAnimationParameters */
// An interface definition for the [WritingToolsCoordinatorAnimationParameters] class.
type IWritingToolsCoordinatorAnimationParameters interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for WritingToolsCoordinatorAnimationParameters */
	// properties:
	CompletionHandler() unsafe.Pointer
	SetCompletionHandler(value unsafe.Pointer)
	Delay() float64
	Duration() float64
	ProgressHandler() unsafe.Pointer
	SetProgressHandler(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for WritingToolsCoordinatorAnimationParameters */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for WritingToolsCoordinatorAnimationParameters */
// Alloc allocates a new instance without initialization.
func (wc _WritingToolsCoordinatorAnimationParametersClass) Alloc() WritingToolsCoordinatorAnimationParameters {
	rv := objc.Send[WritingToolsCoordinatorAnimationParameters](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (wc _WritingToolsCoordinatorAnimationParametersClass) New() WritingToolsCoordinatorAnimationParameters {
	rv := objc.Send[WritingToolsCoordinatorAnimationParameters](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WritingToolsCoordinatorAnimationParameters) Init() WritingToolsCoordinatorAnimationParameters {
	rv := objc.Send[WritingToolsCoordinatorAnimationParameters](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WritingToolsCoordinatorAnimationParameters) Autorelease() WritingToolsCoordinatorAnimationParameters {
	rv := objc.Send[WritingToolsCoordinatorAnimationParameters](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWritingToolsCoordinatorAnimationParameters creates a new WritingToolsCoordinatorAnimationParameters instance.
func NewWritingToolsCoordinatorAnimationParameters() WritingToolsCoordinatorAnimationParameters {
	return getWritingToolsCoordinatorAnimationParametersClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for WritingToolsCoordinatorAnimationParameters */
// An object you use to configure additional tasks or animations to run alongside the Writing Tools animations.
//
// When Writing Tools replaces text in one of your context objects, it provides an object for you to use to configure any additional animations. During a Writing Tools session, you hide the text under evaluation and provide a targeted preview of your content. Writing Tools animations changes to that preview, but you might need to provide additional animations for other parts of your view’s content. For example, you might need to animate any layout changes caused by the insertion or removal of text in other parts of your view. Use this object to configure those animations. You don’t create an object directly. Instead, the system creates one and passes it to the method of your object. Use that object to specify the blocks to run during and after the system animations.


// An object you use to configure additional tasks or animations to run alongside the Writing Tools animations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/AnimationParameters
type WritingToolsCoordinatorAnimationParameters struct {
	objectivec.Object
}

// WritingToolsCoordinatorAnimationParametersFrom constructs a [WritingToolsCoordinatorAnimationParameters] from an unsafe.Pointer.
//
// An object you use to configure additional tasks or animations to run alongside the Writing Tools animations.
func WritingToolsCoordinatorAnimationParametersFrom(ptr unsafe.Pointer) WritingToolsCoordinatorAnimationParameters {
	return WritingToolsCoordinatorAnimationParameters{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for WritingToolsCoordinatorAnimationParameters *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for WritingToolsCoordinatorAnimationParameters */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for WritingToolsCoordinatorAnimationParameters */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for WritingToolsCoordinatorAnimationParameters */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for WritingToolsCoordinatorAnimationParameters */

// A custom block to run when the system animations finish.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/AnimationParameters/completionHandler
func (w_ WritingToolsCoordinatorAnimationParameters) CompletionHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("completionHandler"))
	return rv
}/* debug [instance_properties/getter]: completionHandler */


// A custom block to run when the system animations finish.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/AnimationParameters/completionHandler
func (w_ WritingToolsCoordinatorAnimationParameters) SetCompletionHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCompletionHandler:"), value)
}/* debug [instance_properties/setter]: completionHandler */


// The number of seconds the system waits before starting its animations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/AnimationParameters/delay
func (w_ WritingToolsCoordinatorAnimationParameters) Delay() float64 {
	rv := objc.Send[float64](w_.ID, objc.Sel("delay"))
	return rv
}/* debug [instance_properties/getter]: delay */


// The number of seconds it takes the system animations to run.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/AnimationParameters/duration
func (w_ WritingToolsCoordinatorAnimationParameters) Duration() float64 {
	rv := objc.Send[float64](w_.ID, objc.Sel("duration"))
	return rv
}/* debug [instance_properties/getter]: duration */


// A custom block that runs at the same time as the system animations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/AnimationParameters/progressHandler
func (w_ WritingToolsCoordinatorAnimationParameters) ProgressHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("progressHandler"))
	return rv
}/* debug [instance_properties/getter]: progressHandler */


// A custom block that runs at the same time as the system animations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/AnimationParameters/progressHandler
func (w_ WritingToolsCoordinatorAnimationParameters) SetProgressHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setProgressHandler:"), value)
}/* debug [instance_properties/setter]: progressHandler */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSWritingToolsCoordinatorAnimationParameters */



