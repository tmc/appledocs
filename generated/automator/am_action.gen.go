// Code generated from Apple documentation for Automator. DO NOT EDIT.

package automator

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [AMAction] class.
type IAMAction interface {
	objectivec.IObject
	Activated()
	Closed()
	DidFinishRunningWithError(errorInfo unsafe.Pointer)
	FinishRunningWithError(error_ foundation.IError)
	LogMessageWithLevelFormat(level AMLogLevel, format string)
	Opened()
	ParametersUpdated()
	Reset()
	RunWithInputError(input objectivec.IObject, error_ unsafe.Pointer) objc.ID
	RunAsynchronouslyWithInput(input objectivec.IObject)
	RunWithInputFromActionError(input objectivec.IObject, anAction IAMAction, errorInfo unsafe.Pointer) objc.ID
	Stop()
	UpdateParameters()
	WillFinishRunning()
	WriteToDictionary(dictionary unsafe.Pointer)
	IgnoresInput() bool
	Stopped() bool
	Name() string
	Output() objc.ID
	SetOutput(value objc.ID)
	ProgressValue() float64
	SetProgressValue(value float64)
	SelectedInputType() string
	SetSelectedInputType(value string)
	SelectedOutputType() string
	SetSelectedOutputType(value string)
	IsStopped() bool
	SetIsStopped(value bool)
}

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

// Alloc allocates a new instance without initialization.
func (ac _AMActionClass) Alloc() AMAction {
	rv := objc.Send[AMAction](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Loads an Automator action from a file URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/init(contentsOf:)

func NewAMActionWithContentsOfURLError(fileURL foundation.IURL, outError unsafe.Pointer) AMAction {
	instance := getAMActionClass().Alloc()
	rv := objc.Send[AMAction](instance.ID, objc.Sel("initWithContentsOfURL:error:"), fileURL, outError)
	rv.Autorelease()
	return rv
}



// Initializes the action with the specified definition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/init(definition:fromArchive:)

func NewAMActionWithDefinitionFromArchive(dict unsafe.Pointer, archived bool) AMAction {
	instance := getAMActionClass().Alloc()
	rv := objc.Send[AMAction](instance.ID, objc.Sel("initWithDefinition:fromArchive:"), dict, archived)
	rv.Autorelease()
	return rv
}




// Allows the action to synchronize its information with settings in another app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/activated()

func (a_ AMAction) Activated() {
	objc.Send[objc.ID](a_.ID, objc.Sel("activated"))
}



// Invoked by Automator when the receiving action is removed from a workflow, allowing it to perform cleanup operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/closed()

func (a_ AMAction) Closed() {
	objc.Send[objc.ID](a_.ID, objc.Sel("closed"))
}



// Sent by the action to itself when it has finished running asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/didFinishRunningWithError:

func (a_ AMAction) DidFinishRunningWithError(errorInfo unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("didFinishRunningWithError:"), errorInfo)
}



// Causes the action to stop running and return an error, which, in turn, causes the workflow to stop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/finishRunningWithError(_:)

func (a_ AMAction) FinishRunningWithError(error_ foundation.IError) {
	objc.Send[objc.ID](a_.ID, objc.Sel("finishRunningWithError:"), error_)
}



// Displays a message in Automator’s log area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/logMessageWithLevel:format:

func (a_ AMAction) LogMessageWithLevelFormat(level AMLogLevel, format string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("logMessageWithLevel:format:"), level, objc.String(format))
}



// Allows the action to initialize its user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/opened()

func (a_ AMAction) Opened() {
	objc.Send[objc.ID](a_.ID, objc.Sel("opened"))
}



// Requests the action to update its user interface from its stored parameters, which have changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/parametersUpdated()

func (a_ AMAction) ParametersUpdated() {
	objc.Send[objc.ID](a_.ID, objc.Sel("parametersUpdated"))
}



// Resets the action to its initial state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/reset()

func (a_ AMAction) Reset() {
	objc.Send[objc.ID](a_.ID, objc.Sel("reset"))
}



// Requests the action to perform its task using the specified input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/run(withInput:)

func (a_ AMAction) RunWithInputError(input objectivec.IObject, error_ unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("runWithInput:error:"), input, error_)
	return rv
}



// Causes Automator to wait for notification that the action has completed execution, which allows the action to perform an asynchronous operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/runAsynchronously(withInput:)

func (a_ AMAction) RunAsynchronouslyWithInput(input objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("runAsynchronouslyWithInput:"), input)
}



// Requests the action to perform its task using the specified input from the specified action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/runWithInput:fromAction:error:

func (a_ AMAction) RunWithInputFromActionError(input objectivec.IObject, anAction IAMAction, errorInfo unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("runWithInput:fromAction:error:"), input, anAction, errorInfo)
	return rv
}



// Stops the action from running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/stop()

func (a_ AMAction) Stop() {
	objc.Send[objc.ID](a_.ID, objc.Sel("stop"))
}



// Requests the action to update its stored set of parameters from the settings in the action’s user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/updateParameters()

func (a_ AMAction) UpdateParameters() {
	objc.Send[objc.ID](a_.ID, objc.Sel("updateParameters"))
}



// Provides an opportunity for an action to perform cleanup operations, such as closing windows and deallocating memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/willFinishRunning()

func (a_ AMAction) WillFinishRunning() {
	objc.Send[objc.ID](a_.ID, objc.Sel("willFinishRunning"))
}



// Examines the parameters and other configuration information specified in the passed dictionary and adds its own information to it if appropriate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/write(to:)

func (a_ AMAction) WriteToDictionary(dictionary unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("writeToDictionary:"), dictionary)
}


// A Boolean value that indicates whether the action acts upon its input or the input is ignored.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/ignoresInput

func (a_ AMAction) IgnoresInput() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("ignoresInput"))
	return rv
}


// A Boolean value that indicates whether the user clicked the stop button on the parent workflow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/isStopped

func (a_ AMAction) Stopped() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("stopped"))
	return rv
}


// The name of the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/name

func (a_ AMAction) Name() string {
	rv := objc.Send[string](a_.ID, objc.Sel("name"))
	return rv
}


// The action’s output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/output

func (a_ AMAction) Output() objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("output"))
	return rv
}


// The action’s output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/output

func (a_ AMAction) SetOutput(value objc.ID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOutput:"), value)
}


// A float value between 0 and 1, which indicates how far along the action is while processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/progressValue

func (a_ AMAction) ProgressValue() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("progressValue"))
	return rv
}


// A float value between 0 and 1, which indicates how far along the action is while processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/progressValue

func (a_ AMAction) SetProgressValue(value float64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setProgressValue:"), value)
}


// The type of input, in UTI format, of the input received by the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/selectedInputType

func (a_ AMAction) SelectedInputType() string {
	rv := objc.Send[string](a_.ID, objc.Sel("selectedInputType"))
	return rv
}


// The type of input, in UTI format, of the input received by the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/selectedInputType

func (a_ AMAction) SetSelectedInputType(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSelectedInputType:"), objc.String(value))
}


// The type of output, in UTI format, of the output to be produced by the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/selectedOutputType

func (a_ AMAction) SelectedOutputType() string {
	rv := objc.Send[string](a_.ID, objc.Sel("selectedOutputType"))
	return rv
}


// The type of output, in UTI format, of the output to be produced by the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAction/selectedOutputType

func (a_ AMAction) SetSelectedOutputType(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSelectedOutputType:"), objc.String(value))
}


// A Boolean value that indicates whether the user clicked the stop button on the parent workflow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automator/amaction/isstopped

func (a_ AMAction) IsStopped() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isStopped"))
	return rv
}


// A Boolean value that indicates whether the user clicked the stop button on the parent workflow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automator/amaction/isstopped

func (a_ AMAction) SetIsStopped(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsStopped:"), value)
}


