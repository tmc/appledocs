// Code generated from Apple documentation for Automator. DO NOT EDIT.

package automator

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [AMWorkflowController] class.
type IAMWorkflowController interface {
	appkit.IController
	Pause(sender objectivec.IObject)
	Reset(sender objectivec.IObject)
	Run(sender objectivec.IObject)
	Step(sender objectivec.IObject)
	Stop(sender objectivec.IObject)
	CanRun() bool
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	Paused() bool
	Running() bool
	Workflow() AMWorkflow
	SetWorkflow(value IAMWorkflow)
	WorkflowView() AMWorkflowView
	SetWorkflowView(value IAMWorkflowView)
	IsPaused() bool
	SetIsPaused(value bool)
	IsRunning() bool
	SetIsRunning(value bool)
}

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

// Alloc allocates a new instance without initialization.
func (ac _AMWorkflowControllerClass) Alloc() AMWorkflowController {
	rv := objc.Send[AMWorkflowController](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Pauses a workflow that’s running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflowController/pause(_:)
func (a_ AMWorkflowController) Pause(sender objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("pause:"), sender)
}


// Stops a workflow, clears any action results, and resets the workflow back to an un-run state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflowController/reset(_:)
func (a_ AMWorkflowController) Reset(sender objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("reset:"), sender)
}


// Runs the associated workflow, after first clearing any results stored by its actions during any previous run.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflowController/run(_:)
func (a_ AMWorkflowController) Run(sender objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("run:"), sender)
}


// In a paused workflow, runs the next action in the workflow and then pauses again.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflowController/step(_:)
func (a_ AMWorkflowController) Step(sender objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("step:"), sender)
}


// Stops the associated workflow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflowController/stop(_:)
func (a_ AMWorkflowController) Stop(sender objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("stop:"), sender)
}


// A Boolean value that indicates whether the controller’s workflow is able to run.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflowController/canRun
func (a_ AMWorkflowController) CanRun() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canRun"))
	return rv
}


// The controller’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflowController/delegate
func (a_ AMWorkflowController) Delegate() objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("delegate"))
	return rv
}


// The controller’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflowController/delegate
func (a_ AMWorkflowController) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDelegate:"), value)
}


// A Boolean value that indicates whether the controller’s workflow is currently paused.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflowController/isPaused
func (a_ AMWorkflowController) Paused() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("paused"))
	return rv
}


// A Boolean value that indicates whether the controller’s workflow is currently running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflowController/isRunning
func (a_ AMWorkflowController) Running() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("running"))
	return rv
}


// The controller’s workflow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflowController/workflow
func (a_ AMWorkflowController) Workflow() AMWorkflow {
	rv := objc.Send[AMWorkflow](a_.ID, objc.Sel("workflow"))
	return rv
}


// The controller’s workflow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflowController/workflow
func (a_ AMWorkflowController) SetWorkflow(value IAMWorkflow) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setWorkflow:"), value)
}


// The controller’s workflow view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflowController/workflowView-swift.property
func (a_ AMWorkflowController) WorkflowView() AMWorkflowView {
	rv := objc.Send[AMWorkflowView](a_.ID, objc.Sel("workflowView"))
	return rv
}


// The controller’s workflow view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflowController/workflowView-swift.property
func (a_ AMWorkflowController) SetWorkflowView(value IAMWorkflowView) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setWorkflowView:"), value)
}


// A Boolean value that indicates whether the controller’s workflow is currently paused.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automator/amworkflowcontroller/ispaused
func (a_ AMWorkflowController) IsPaused() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isPaused"))
	return rv
}


// A Boolean value that indicates whether the controller’s workflow is currently paused.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automator/amworkflowcontroller/ispaused
func (a_ AMWorkflowController) SetIsPaused(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsPaused:"), value)
}


// A Boolean value that indicates whether the controller’s workflow is currently running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automator/amworkflowcontroller/isrunning
func (a_ AMWorkflowController) IsRunning() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isRunning"))
	return rv
}


// A Boolean value that indicates whether the controller’s workflow is currently running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automator/amworkflowcontroller/isrunning
func (a_ AMWorkflowController) SetIsRunning(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsRunning:"), value)
}



