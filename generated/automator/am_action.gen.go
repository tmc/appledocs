// Code generated from Apple documentation for Automator. DO NOT EDIT.

package automator

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coretelephony"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AMAction */


/* debug [class_header]: Header for AMAction */
// The class instance for the [AMAction] class.
var (
	AMActionClass     _AMActionClass
	AMActionClassOnce sync.Once
)

func getAMActionClass() _AMActionClass {
	AMActionClassOnce.Do(func() {
		AMActionClass = _AMActionClass{objc.GetClass("AMAction")}
	})
	return AMActionClass
}

type _AMActionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AMAction */
// An interface definition for the [AMAction] class.
type IAMAction interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AMAction */
	// properties:
	IgnoresInput() bool
	Stopped() bool
	Name() objc.IObject /* cross-framework: NSString */
	Output() objc.ID
	SetOutput(value objc.ID)
	ProgressValue() float64
	SetProgressValue(value float64)
	SelectedInputType() objc.IObject /* cross-framework: NSString */
	SetSelectedInputType(value objc.IObject /* cross-framework: NSString */)
	SelectedOutputType() objc.IObject /* cross-framework: NSString */
	SetSelectedOutputType(value objc.IObject /* cross-framework: NSString */)
	IsStopped() bool
	SetIsStopped(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AMAction */
	// methods:
	Activated()
	Closed()
	FinishRunningWithError(error_ objc.IObject /* cross-framework: Error */)
	Opened()
	ParametersUpdated()
	Reset()
	RunWithInputError(input objc.IObject, error_ unsafe.Pointer) objc.ID
	RunAsynchronouslyWithInput(input objc.IObject)
	Stop()
	UpdateParameters()
	WillFinishRunning()
	WriteToDictionary(dictionary unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AMAction */
// Alloc allocates a new instance without initialization.
func (ac _AMActionClass) Alloc() AMAction {
	rv := objc.Send[AMAction](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AMActionClass) New() AMAction {
	rv := objc.Send[AMAction](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AMAction) Init() AMAction {
	rv := objc.Send[AMAction](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AMAction) Autorelease() AMAction {
	rv := objc.Send[AMAction](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAMAction creates a new AMAction instance.
func NewAMAction() AMAction {
	return getAMActionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AMAction */
// An abstract class that defines the interface and general characteristics of Automator actions.
//
// Automator is an Apple app that allows users to construct and execute workflows consisting of a sequence of discrete modules called actions. An action performs a specific task, such as copying a file or cropping an image, and passes its output to Automator to give to the next action in the workflow. Actions are currently implemented as loadable bundles owned by objects of the class, a subclass of . The critically important method declared by is . When Automator executes a workflow, it sends this message to each action object in the workflow (in workflow sequence), in most cases passing in the output of the previous action as input. The action object performs its task in this method and ends by returning an output object for the next action in the workflow. Subclassing is not recommended. For most situations requiring an enhancement to the Automator framework, it is sufficient to subclass .


// An abstract class that defines the interface and general characteristics of Automator actions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction
type AMAction struct {
	objectivec.Object
}

// AMActionFrom constructs a [AMAction] from an unsafe.Pointer.
//
// An abstract class that defines the interface and general characteristics of Automator actions.
func AMActionFrom(ptr unsafe.Pointer) AMAction {
	return AMAction{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AMAction */

// Loads an Automator action from a file URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/init(contentsOf:)
func NewAMActionWithContentsOfURLError(fileURL objc.IObject /* cross-framework: NSURL */, outError unsafe.Pointer) AMAction {
	instance := getAMActionClass().Alloc()
	rv := objc.Send[AMAction](instance.ID, objc.Sel("initWithContentsOfURL:error:"), fileURL, outError)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAMActionWithContentsOfURLError */


// Initializes the action with the specified definition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/init(definition:fromArchive:)
func NewAMActionWithDefinitionFromArchive(dict foundation.IDictionary, archived bool) AMAction {
	instance := getAMActionClass().Alloc()
	rv := objc.Send[AMAction](instance.ID, objc.Sel("initWithDefinition:fromArchive:"), dict, archived)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAMActionWithDefinitionFromArchive */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AMAction */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AMAction */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AMAction */

// Allows the action to synchronize its information with settings in another app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/activated()
func (a_ AMAction) Activated() {
	objc.Send[objc.ID](a_.ID, objc.Sel("activated"))
}/* debug [instance_methods/method]: Activated */


// Invoked by Automator when the receiving action is removed from a workflow, allowing it to perform cleanup operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/closed()
func (a_ AMAction) Closed() {
	objc.Send[objc.ID](a_.ID, objc.Sel("closed"))
}/* debug [instance_methods/method]: Closed */


// Causes the action to stop running and return an error, which, in turn, causes the workflow to stop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/finishRunningWithError(_:)
func (a_ AMAction) FinishRunningWithError(error_ objc.IObject /* cross-framework: Error */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("finishRunningWithError:"), error_)
}/* debug [instance_methods/method]: FinishRunningWithError */


// Allows the action to initialize its user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/opened()
func (a_ AMAction) Opened() {
	objc.Send[objc.ID](a_.ID, objc.Sel("opened"))
}/* debug [instance_methods/method]: Opened */


// Requests the action to update its user interface from its stored parameters, which have changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/parametersUpdated()
func (a_ AMAction) ParametersUpdated() {
	objc.Send[objc.ID](a_.ID, objc.Sel("parametersUpdated"))
}/* debug [instance_methods/method]: ParametersUpdated */


// Resets the action to its initial state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/reset()
func (a_ AMAction) Reset() {
	objc.Send[objc.ID](a_.ID, objc.Sel("reset"))
}/* debug [instance_methods/method]: Reset */


// Requests the action to perform its task using the specified input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/run(withInput:)
func (a_ AMAction) RunWithInputError(input objc.IObject, error_ unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("runWithInput:error:"), input, error_)
	return rv
}/* debug [instance_methods/method]: RunWithInputError */


// Causes Automator to wait for notification that the action has completed execution, which allows the action to perform an asynchronous operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/runAsynchronously(withInput:)
func (a_ AMAction) RunAsynchronouslyWithInput(input objc.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("runAsynchronouslyWithInput:"), input)
}/* debug [instance_methods/method]: RunAsynchronouslyWithInput */


// Stops the action from running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/stop()
func (a_ AMAction) Stop() {
	objc.Send[objc.ID](a_.ID, objc.Sel("stop"))
}/* debug [instance_methods/method]: Stop */


// Requests the action to update its stored set of parameters from the settings in the action’s user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/updateParameters()
func (a_ AMAction) UpdateParameters() {
	objc.Send[objc.ID](a_.ID, objc.Sel("updateParameters"))
}/* debug [instance_methods/method]: UpdateParameters */


// Provides an opportunity for an action to perform cleanup operations, such as closing windows and deallocating memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/willFinishRunning()
func (a_ AMAction) WillFinishRunning() {
	objc.Send[objc.ID](a_.ID, objc.Sel("willFinishRunning"))
}/* debug [instance_methods/method]: WillFinishRunning */


// Examines the parameters and other configuration information specified in the passed dictionary and adds its own information to it if appropriate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/write(to:)
func (a_ AMAction) WriteToDictionary(dictionary unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("writeToDictionary:"), dictionary)
}/* debug [instance_methods/method]: WriteToDictionary */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AMAction */

// A Boolean value that indicates whether the action acts upon its input or the input is ignored.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/ignoresInput
func (a_ AMAction) IgnoresInput() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("ignoresInput"))
	return rv
}/* debug [instance_properties/getter]: ignoresInput */


// A Boolean value that indicates whether the user clicked the stop button on the parent workflow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/isStopped
func (a_ AMAction) Stopped() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("stopped"))
	return rv
}/* debug [instance_properties/getter]: stopped */


// The name of the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/name
func (a_ AMAction) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The action’s output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/output
func (a_ AMAction) Output() objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("output"))
	return rv
}/* debug [instance_properties/getter]: output */


// The action’s output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/output
func (a_ AMAction) SetOutput(value objc.ID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOutput:"), value)
}/* debug [instance_properties/setter]: output */


// A float value between 0 and 1, which indicates how far along the action is while processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/progressValue
func (a_ AMAction) ProgressValue() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("progressValue"))
	return rv
}/* debug [instance_properties/getter]: progressValue */


// A float value between 0 and 1, which indicates how far along the action is while processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/progressValue
func (a_ AMAction) SetProgressValue(value float64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setProgressValue:"), value)
}/* debug [instance_properties/setter]: progressValue */


// The type of input, in UTI format, of the input received by the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/selectedInputType
func (a_ AMAction) SelectedInputType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("selectedInputType"))
	return rv
}/* debug [instance_properties/getter]: selectedInputType */


// The type of input, in UTI format, of the input received by the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/selectedInputType
func (a_ AMAction) SetSelectedInputType(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSelectedInputType:"), value)
}/* debug [instance_properties/setter]: selectedInputType */


// The type of output, in UTI format, of the output to be produced by the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/selectedOutputType
func (a_ AMAction) SelectedOutputType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("selectedOutputType"))
	return rv
}/* debug [instance_properties/getter]: selectedOutputType */


// The type of output, in UTI format, of the output to be produced by the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/selectedOutputType
func (a_ AMAction) SetSelectedOutputType(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSelectedOutputType:"), value)
}/* debug [instance_properties/setter]: selectedOutputType */


// A Boolean value that indicates whether the user clicked the stop button on the parent workflow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automator/amaction/isstopped
func (a_ AMAction) IsStopped() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isStopped"))
	return rv
}/* debug [instance_properties/getter]: isStopped */


// A Boolean value that indicates whether the user clicked the stop button on the parent workflow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automator/amaction/isstopped
func (a_ AMAction) SetIsStopped(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsStopped:"), value)
}/* debug [instance_properties/setter]: isStopped */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AMAction */


