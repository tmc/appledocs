// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSRunLoop */


/* debug [class_header]: Header for NSRunLoop */
// The class instance for the [RunLoop] class.
var (
	RunLoopClass     _RunLoopClass
	RunLoopClassOnce sync.Once
)

func getRunLoopClass() _RunLoopClass {
	RunLoopClassOnce.Do(func() {
		RunLoopClass = _RunLoopClass{objc.GetClass("NSRunLoop")}
	})
	return RunLoopClass
}

type _RunLoopClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RunLoop */
// An interface definition for the [RunLoop] class.
type IRunLoop interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for RunLoop */
	// properties:
	CurrentMode() RunLoopMode
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RunLoop */
	// methods:
	AcceptInputForModeBeforeDate(mode RunLoopMode, limitDate IDate)
	AddTimerForMode(timer ITimer, mode RunLoopMode)
	AddPortForMode(aPort IPort, mode RunLoopMode)
	CancelPerformSelectorTargetArgument(aSelector objc.SEL, target objc.IObject, arg objc.IObject)
	CancelPerformSelectorsWithTarget(target objc.IObject)
	GetCFRunLoop() RunLoopRef /* not a class type */
	LimitDateForMode(mode RunLoopMode) IDate
	PerformBlock(block unsafe.Pointer)
	PerformSelectorTargetArgumentOrderModes(aSelector objc.SEL, target objc.IObject, arg objc.IObject, order uint, modes []string)
	PerformInModesBlock(modes []string, block unsafe.Pointer)
	RemovePortForMode(aPort IPort, mode RunLoopMode)
	Run()
	RunModeBeforeDate(mode RunLoopMode, limitDate IDate) bool
	RunUntilDate(limitDate IDate)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RunLoop */
// Alloc allocates a new instance without initialization.
func (rc _RunLoopClass) Alloc() RunLoop {
	rv := objc.Send[RunLoop](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RunLoopClass) New() RunLoop {
	rv := objc.Send[RunLoop](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RunLoop) Init() RunLoop {
	rv := objc.Send[RunLoop](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RunLoop) Autorelease() RunLoop {
	rv := objc.Send[RunLoop](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRunLoop creates a new RunLoop instance.
func NewRunLoop() RunLoop {
	return getRunLoopClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RunLoop */
// The programmatic interface to objects that manage input sources.
//
// A object processes input for sources, such as mouse and keyboard events from the window system and objects. A object also processes events. Your application neither creates nor explicitly manages objects. The system creates a object as needed for each object, including the application’s main thread. If you need to access the current thread’s run loop, use the class method . Note that from the perspective of , objects aren’t “input”—they’re a special type, and they don’t cause the run loop to return when they fire.


// The programmatic interface to objects that manage input sources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RunLoop
type RunLoop struct {
	objectivec.Object
}

// RunLoopFrom constructs a [RunLoop] from an unsafe.Pointer.
//
// The programmatic interface to objects that manage input sources.
func RunLoopFrom(ptr unsafe.Pointer) RunLoop {
	return RunLoop{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RunLoop *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RunLoop */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RunLoop */

// Returns the run loop for the current thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RunLoop/current
func (rc _RunLoopClass) CurrentRunLoop() RunLoop {
	rv := objc.Send[RunLoop](objc.ID(rc.class), objc.Sel("currentRunLoop"))
	return rv
}/* debug [class_properties_class/property]: currentRunLoop */

// Returns the run loop of the main thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RunLoop/main
func (rc _RunLoopClass) MainRunLoop() RunLoop {
	rv := objc.Send[RunLoop](objc.ID(rc.class), objc.Sel("mainRunLoop"))
	return rv
}/* debug [class_properties_class/property]: mainRunLoop */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RunLoop */

// Runs the loop once or until the specified date, accepting input only for the specified mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RunLoop/acceptInput(forMode:before:)
func (r_ RunLoop) AcceptInputForModeBeforeDate(mode RunLoopMode, limitDate IDate) {
	objc.Send[objc.ID](r_.ID, objc.Sel("acceptInputForMode:beforeDate:"), mode, limitDate)
}/* debug [instance_methods/method]: AcceptInputForModeBeforeDate */


// Registers a given timer with a given input mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RunLoop/add(_:forMode:)-392ag
func (r_ RunLoop) AddTimerForMode(timer ITimer, mode RunLoopMode) {
	objc.Send[objc.ID](r_.ID, objc.Sel("addTimer:forMode:"), timer, mode)
}/* debug [instance_methods/method]: AddTimerForMode */


// Adds a port as an input source to the specified mode of the run loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RunLoop/add(_:forMode:)-6z982
func (r_ RunLoop) AddPortForMode(aPort IPort, mode RunLoopMode) {
	objc.Send[objc.ID](r_.ID, objc.Sel("addPort:forMode:"), aPort, mode)
}/* debug [instance_methods/method]: AddPortForMode */


// Cancels the sending of a previously scheduled message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RunLoop/cancelPerform(_:target:argument:)
func (r_ RunLoop) CancelPerformSelectorTargetArgument(aSelector objc.SEL, target objc.IObject, arg objc.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("cancelPerformSelector:target:argument:"), aSelector, target, arg)
}/* debug [instance_methods/method]: CancelPerformSelectorTargetArgument */


// Cancels all outstanding ordered performs scheduled with a given target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RunLoop/cancelPerformSelectors(withTarget:)
func (r_ RunLoop) CancelPerformSelectorsWithTarget(target objc.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("cancelPerformSelectorsWithTarget:"), target)
}/* debug [instance_methods/method]: CancelPerformSelectorsWithTarget */


// Returns the receiver’s underlying run loop object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RunLoop/getCFRunLoop()
func (r_ RunLoop) GetCFRunLoop() RunLoopRef /* not a class type */ {
	rv := objc.Send[RunLoopRef](r_.ID, objc.Sel("getCFRunLoop"))
	return rv
}/* debug [instance_methods/method]: GetCFRunLoop */


// Performs one pass through the run loop in the specified mode and returns the date at which the next timer is scheduled to fire.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RunLoop/limitDate(forMode:)
func (r_ RunLoop) LimitDateForMode(mode RunLoopMode) IDate {
	rv := objc.Send[Date](r_.ID, objc.Sel("limitDateForMode:"), mode)
	return rv
}/* debug [instance_methods/method]: LimitDateForMode */


// Schedules a block that the run loop invokes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RunLoop/perform(_:)
func (r_ RunLoop) PerformBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("performBlock:"), block)
}/* debug [instance_methods/method]: PerformBlock */


// Schedules the sending of a message on the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RunLoop/perform(_:target:argument:order:modes:)
func (r_ RunLoop) PerformSelectorTargetArgumentOrderModes(aSelector objc.SEL, target objc.IObject, arg objc.IObject, order uint, modes []string) {
	objc.Send[objc.ID](r_.ID, objc.Sel("performSelector:target:argument:order:modes:"), aSelector, target, arg, order, modes)
}/* debug [instance_methods/method]: PerformSelectorTargetArgumentOrderModes */


// Schedules a block that the run loop invokes when it’s running in any of the specified modes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RunLoop/perform(inModes:block:)
func (r_ RunLoop) PerformInModesBlock(modes []string, block unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("performInModes:block:"), modes, block)
}/* debug [instance_methods/method]: PerformInModesBlock */


// Removes a port from the specified input mode of the run loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RunLoop/remove(_:forMode:)
func (r_ RunLoop) RemovePortForMode(aPort IPort, mode RunLoopMode) {
	objc.Send[objc.ID](r_.ID, objc.Sel("removePort:forMode:"), aPort, mode)
}/* debug [instance_methods/method]: RemovePortForMode */


// Puts the receiver into a permanent loop, during which time it processes data from all attached input sources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RunLoop/run()
func (r_ RunLoop) Run() {
	objc.Send[objc.ID](r_.ID, objc.Sel("run"))
}/* debug [instance_methods/method]: Run */


// Runs the loop once, blocking for input in the specified mode until a given date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RunLoop/run(mode:before:)
func (r_ RunLoop) RunModeBeforeDate(mode RunLoopMode, limitDate IDate) bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("runMode:beforeDate:"), mode, limitDate)
	return rv
}/* debug [instance_methods/method]: RunModeBeforeDate */


// Runs the loop until the specified date, during which time it processes data from all attached input sources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RunLoop/run(until:)
func (r_ RunLoop) RunUntilDate(limitDate IDate) {
	objc.Send[objc.ID](r_.ID, objc.Sel("runUntilDate:"), limitDate)
}/* debug [instance_methods/method]: RunUntilDate */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RunLoop */

// Returns the run loop for the current thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RunLoop/current
func (r_ RunLoop) CurrentRunLoop() IRunLoop {
	rv := objc.Send[RunLoop](r_.ID, objc.Sel("currentRunLoop"))
	return rv
}/* debug [instance_properties/getter]: currentRunLoop */


// The receiver’s current input mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RunLoop/currentMode
func (r_ RunLoop) CurrentMode() RunLoopMode {
	rv := objc.Send[RunLoopMode](r_.ID, objc.Sel("currentMode"))
	return rv
}/* debug [instance_properties/getter]: currentMode */


// Returns the run loop of the main thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RunLoop/main
func (r_ RunLoop) MainRunLoop() IRunLoop {
	rv := objc.Send[RunLoop](r_.ID, objc.Sel("mainRunLoop"))
	return rv
}/* debug [instance_properties/getter]: mainRunLoop */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSRunLoop */



