// Code generated from Apple documentation for Automator. DO NOT EDIT.

package automator

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	// properties:
	IgnoresInput() bool /* primitive/slice/pointer. */
	SetIgnoresInput(value bool /* primitive/slice/pointer. */)
	IsStopped() bool /* primitive/slice/pointer. */
	SetIsStopped(value bool /* primitive/slice/pointer. */)
	Name() string /* primitive/slice/pointer. */
	SetName(value string /* primitive/slice/pointer. */)
	Output() unsafe.Pointer
	SetOutput(value unsafe.Pointer)
	ProgressValue() float64 /* primitive/slice/pointer. */
	SetProgressValue(value float64 /* primitive/slice/pointer. */)
	SelectedInputType() string /* primitive/slice/pointer. */
	SetSelectedInputType(value string /* primitive/slice/pointer. */)
	SelectedOutputType() string /* primitive/slice/pointer. */
	SetSelectedOutputType(value string /* primitive/slice/pointer. */)
	// methods:
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



// A Boolean value that indicates whether the action acts upon its input or the input is ignored.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automator/amaction/ignoresinput
func (a_ AMAction) IgnoresInput() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("ignoresInput"))
	return rv
}


// A Boolean value that indicates whether the action acts upon its input or the input is ignored.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automator/amaction/ignoresinput
func (a_ AMAction) SetIgnoresInput(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIgnoresInput:"), value)
}


// A Boolean value that indicates whether the user clicked the stop button on the parent workflow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automator/amaction/isstopped
func (a_ AMAction) IsStopped() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("isStopped"))
	return rv
}


// A Boolean value that indicates whether the user clicked the stop button on the parent workflow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automator/amaction/isstopped
func (a_ AMAction) SetIsStopped(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsStopped:"), value)
}


// The name of the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automator/amaction/name
func (a_ AMAction) Name() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](a_.ID, objc.Sel("name"))
	return rv
}


// The name of the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automator/amaction/name
func (a_ AMAction) SetName(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setName:"), objc.String(value))
}


// The action’s output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automator/amaction/output
func (a_ AMAction) Output() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("output"))
	return rv
}


// The action’s output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automator/amaction/output
func (a_ AMAction) SetOutput(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOutput:"), value)
}


// A float value between 0 and 1, which indicates how far along the action is while processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automator/amaction/progressvalue
func (a_ AMAction) ProgressValue() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](a_.ID, objc.Sel("progressValue"))
	return rv
}


// A float value between 0 and 1, which indicates how far along the action is while processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automator/amaction/progressvalue
func (a_ AMAction) SetProgressValue(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setProgressValue:"), value)
}


// The type of input, in UTI format, of the input received by the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automator/amaction/selectedinputtype
func (a_ AMAction) SelectedInputType() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](a_.ID, objc.Sel("selectedInputType"))
	return rv
}


// The type of input, in UTI format, of the input received by the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automator/amaction/selectedinputtype
func (a_ AMAction) SetSelectedInputType(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSelectedInputType:"), objc.String(value))
}


// The type of output, in UTI format, of the output to be produced by the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automator/amaction/selectedoutputtype
func (a_ AMAction) SelectedOutputType() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](a_.ID, objc.Sel("selectedOutputType"))
	return rv
}


// The type of output, in UTI format, of the output to be produced by the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automator/amaction/selectedoutputtype
func (a_ AMAction) SetSelectedOutputType(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSelectedOutputType:"), objc.String(value))
}



