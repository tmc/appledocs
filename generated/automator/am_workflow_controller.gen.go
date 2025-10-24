// Code generated from Apple documentation for Automator. DO NOT EDIT.

package automator

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AMWorkflowController */


/* debug [class_header]: Header for AMWorkflowController */
// The class instance for the [AMWorkflowController] class.
var (
	AMWorkflowControllerClass     _AMWorkflowControllerClass
	AMWorkflowControllerClassOnce sync.Once
)

func getAMWorkflowControllerClass() _AMWorkflowControllerClass {
	AMWorkflowControllerClassOnce.Do(func() {
		AMWorkflowControllerClass = _AMWorkflowControllerClass{objc.GetClass("AMWorkflowController")}
	})
	return AMWorkflowControllerClass
}

type _AMWorkflowControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AMWorkflowController */
// An interface definition for the [AMWorkflowController] class.
type IAMWorkflowController interface {
	appkit.IController
	
/* debug [class_interface_properties]: Properties for AMWorkflowController */
	// properties:
	CanRun() bool
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	Paused() bool
	Running() bool
	Workflow() IAMWorkflow
	SetWorkflow(value IAMWorkflow)
	WorkflowView() IAMWorkflowView
	SetWorkflowView(value IAMWorkflowView)
	IsPaused() bool
	SetIsPaused(value bool)
	IsRunning() bool
	SetIsRunning(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AMWorkflowController */
	// methods:
	Pause(sender objc.IObject)
	Reset(sender objc.IObject)
	Run(sender objc.IObject)
	Step(sender objc.IObject)
	Stop(sender objc.IObject)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AMWorkflowController */
// Alloc allocates a new instance without initialization.
func (ac _AMWorkflowControllerClass) Alloc() AMWorkflowController {
	rv := objc.Send[AMWorkflowController](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AMWorkflowControllerClass) New() AMWorkflowController {
	rv := objc.Send[AMWorkflowController](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AMWorkflowController) Init() AMWorkflowController {
	rv := objc.Send[AMWorkflowController](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AMWorkflowController) Autorelease() AMWorkflowController {
	rv := objc.Send[AMWorkflowController](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAMWorkflowController creates a new AMWorkflowController instance.
func NewAMWorkflowController() AMWorkflowController {
	return getAMWorkflowControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AMWorkflowController */
// An object that lets you manage an Automator workflow in your app.
//
// A controller can run and stop a workflow and obtain information about its state. The controller’s delegate ( ) receives messages as the workflow is executed and its actions are run. You can load and run a workflow with minimal overhead by using the class method . Use where you need greater control, such as the ability to start and stop the workflow. In that case, you must create and initialize both the workflow and the workflow controller objects.


// An object that lets you manage an Automator workflow in your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflowController
type AMWorkflowController struct {
	appkit.Controller
}

// AMWorkflowControllerFrom constructs a [AMWorkflowController] from an unsafe.Pointer.
//
// An object that lets you manage an Automator workflow in your app.
func AMWorkflowControllerFrom(ptr unsafe.Pointer) AMWorkflowController {
	return AMWorkflowController{
		Controller: appkit.ControllerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AMWorkflowController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AMWorkflowController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AMWorkflowController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AMWorkflowController */

// Pauses a workflow that’s running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflowController/pause(_:)
func (a_ AMWorkflowController) Pause(sender objc.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("pause:"), sender)
}/* debug [instance_methods/method]: Pause */


// Stops a workflow, clears any action results, and resets the workflow back to an un-run state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflowController/reset(_:)
func (a_ AMWorkflowController) Reset(sender objc.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("reset:"), sender)
}/* debug [instance_methods/method]: Reset */


// Runs the associated workflow, after first clearing any results stored by its actions during any previous run.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflowController/run(_:)
func (a_ AMWorkflowController) Run(sender objc.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("run:"), sender)
}/* debug [instance_methods/method]: Run */


// In a paused workflow, runs the next action in the workflow and then pauses again.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflowController/step(_:)
func (a_ AMWorkflowController) Step(sender objc.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("step:"), sender)
}/* debug [instance_methods/method]: Step */


// Stops the associated workflow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflowController/stop(_:)
func (a_ AMWorkflowController) Stop(sender objc.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("stop:"), sender)
}/* debug [instance_methods/method]: Stop */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AMWorkflowController */

// A Boolean value that indicates whether the controller’s workflow is able to run.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflowController/canRun
func (a_ AMWorkflowController) CanRun() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canRun"))
	return rv
}/* debug [instance_properties/getter]: canRun */


// The controller’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflowController/delegate
func (a_ AMWorkflowController) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The controller’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflowController/delegate
func (a_ AMWorkflowController) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// A Boolean value that indicates whether the controller’s workflow is currently paused.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflowController/isPaused
func (a_ AMWorkflowController) Paused() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("paused"))
	return rv
}/* debug [instance_properties/getter]: paused */


// A Boolean value that indicates whether the controller’s workflow is currently running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflowController/isRunning
func (a_ AMWorkflowController) Running() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("running"))
	return rv
}/* debug [instance_properties/getter]: running */


// The controller’s workflow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflowController/workflow
func (a_ AMWorkflowController) Workflow() IAMWorkflow {
	rv := objc.Send[AMWorkflow](a_.ID, objc.Sel("workflow"))
	return rv
}/* debug [instance_properties/getter]: workflow */


// The controller’s workflow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflowController/workflow
func (a_ AMWorkflowController) SetWorkflow(value IAMWorkflow) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setWorkflow:"), value)
}/* debug [instance_properties/setter]: workflow */


// The controller’s workflow view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflowController/workflowView-swift.property
func (a_ AMWorkflowController) WorkflowView() IAMWorkflowView {
	rv := objc.Send[AMWorkflowView](a_.ID, objc.Sel("workflowView"))
	return rv
}/* debug [instance_properties/getter]: workflowView */


// The controller’s workflow view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflowController/workflowView-swift.property
func (a_ AMWorkflowController) SetWorkflowView(value IAMWorkflowView) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setWorkflowView:"), value)
}/* debug [instance_properties/setter]: workflowView */


// A Boolean value that indicates whether the controller’s workflow is currently paused.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automator/amworkflowcontroller/ispaused
func (a_ AMWorkflowController) IsPaused() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isPaused"))
	return rv
}/* debug [instance_properties/getter]: isPaused */


// A Boolean value that indicates whether the controller’s workflow is currently paused.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automator/amworkflowcontroller/ispaused
func (a_ AMWorkflowController) SetIsPaused(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsPaused:"), value)
}/* debug [instance_properties/setter]: isPaused */


// A Boolean value that indicates whether the controller’s workflow is currently running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automator/amworkflowcontroller/isrunning
func (a_ AMWorkflowController) IsRunning() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isRunning"))
	return rv
}/* debug [instance_properties/getter]: isRunning */


// A Boolean value that indicates whether the controller’s workflow is currently running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automator/amworkflowcontroller/isrunning
func (a_ AMWorkflowController) SetIsRunning(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsRunning:"), value)
}/* debug [instance_properties/setter]: isRunning */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AMWorkflowController */



