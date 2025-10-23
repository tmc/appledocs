// Code generated from Apple documentation for Automator. DO NOT EDIT.

package automator

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AMWorkflow] class.
var (
	AMWorkflowClass     _AMWorkflowClass
	AMWorkflowClassOnce sync.Once
)

func getAMWorkflowClass() _AMWorkflowClass {
	AMWorkflowClassOnce.Do(func() {
		AMWorkflowClass = _AMWorkflowClass{objc.GetClass("AMWorkflow")}
	})
	return AMWorkflowClass
}

type _AMWorkflowClass struct {
	class objc.Class
}

// An interface definition for the [AMWorkflow] class.
type IAMWorkflow interface {
	objectivec.IObject
	AddAction(action IAMAction)
	InsertActionAtIndex(action IAMAction, index uint)
	MoveActionAtIndexToIndex(startIndex uint, endIndex uint)
	RemoveAction(action IAMAction)
	SetValueForVariableWithName(value objectivec.IObject, variableName string) bool
	ValueForVariableWithName(variableName string) objc.ID
	WriteToURLError(fileURL foundation.IURL, outError unsafe.Pointer) bool
	Actions() []AMAction
	FileURL() foundation.URL
	Input() objc.ID
	SetInput(value objc.ID)
	Output() objc.ID
}

// An object that lets you use an Automator workflow in your app.
//
// A consists of one or more actions, discrete tasks which together can perform complex automation tasks. Your app can use workflows to package its own features and to take advantage of features provided by other apps. You create actions with Xcode, while you create workflows with the Automator app. You can load and run a workflow with minimal overhead by using the class method . However, in situations where you need greater control, such as the ability to start and stop the workflow, you can use an instance of the class instead. In that case, you must create and initialize both the workflow and the workflow controller objects. In either case, the workflow runs in a separate process so that any actions it contains are executed in a separate memory space. That helps to insulate your app from crashes, memory leaks, or exceptions that might occur from running the actions in the workflow. You can display a workflow with an instance of .


// An object that lets you use an Automator workflow in your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflow
type AMWorkflow struct {
	objectivec.Object
}

// AMWorkflowFrom constructs a [AMWorkflow] from an unsafe.Pointer.
//
// An object that lets you use an Automator workflow in your app.
func AMWorkflowFrom(ptr unsafe.Pointer) AMWorkflow {
	return AMWorkflow{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AMWorkflowClass) Alloc() AMWorkflow {
	rv := objc.Send[AMWorkflow](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AMWorkflowClass) New() AMWorkflow {
	rv := objc.Send[AMWorkflow](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AMWorkflow) Init() AMWorkflow {
	rv := objc.Send[AMWorkflow](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AMWorkflow) Autorelease() AMWorkflow {
	rv := objc.Send[AMWorkflow](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAMWorkflow creates a new AMWorkflow instance.
func NewAMWorkflow() AMWorkflow {
	return getAMWorkflowClass().New()
}



// Creates and initializes a workflow based on the contents of the specified file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflow/init(contentsOf:)
func NewAMWorkflowWithContentsOfURLError(fileURL foundation.IURL, outError unsafe.Pointer) AMWorkflow {
	instance := getAMWorkflowClass().Alloc()
	rv := objc.Send[AMWorkflow](instance.ID, objc.Sel("initWithContentsOfURL:error:"), fileURL, outError)
	rv.Autorelease()
	return rv
}



// Loads and runs the specified workflow file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflow/run(at:withInput:)
func (ac _AMWorkflowClass) RunWorkflowAtURLWithInputError(fileURL foundation.IURL, input objectivec.IObject, error_ unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("runWorkflowAtURL:withInput:error:"), fileURL, input, error_)
	return rv
}


// Adds the specified action at the end of the receiving workflow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflow/addAction(_:)
func (a_ AMWorkflow) AddAction(action IAMAction) {
	objc.Send[objc.ID](a_.ID, objc.Sel("addAction:"), action)
}


// Inserts the specified action at the specified position of the receiving workflow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflow/insertAction(_:at:)
func (a_ AMWorkflow) InsertActionAtIndex(action IAMAction, index uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("insertAction:atIndex:"), action, index)
}


// Moves the action from the specified start position to the specified end position in the receiving workflow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflow/moveAction(at:to:)
func (a_ AMWorkflow) MoveActionAtIndexToIndex(startIndex uint, endIndex uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("moveActionAtIndex:toIndex:"), startIndex, endIndex)
}


// Removes the specified action from the workflow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflow/removeAction(_:)
func (a_ AMWorkflow) RemoveAction(action IAMAction) {
	objc.Send[objc.ID](a_.ID, objc.Sel("removeAction:"), action)
}


// Sets the value of the workflow variable with the specified name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflow/setValue(_:forVariableWithName:)
func (a_ AMWorkflow) SetValueForVariableWithName(value objectivec.IObject, variableName string) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setValue:forVariableWithName:"), value, objc.String(variableName))
	return rv
}


// Returns the value of the workflow variable with the specified name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflow/valueForVariable(withName:)
func (a_ AMWorkflow) ValueForVariableWithName(variableName string) objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("valueForVariableWithName:"), objc.String(variableName))
	return rv
}


// Writes the workflow to the specified file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflow/write(to:)
func (a_ AMWorkflow) WriteToURLError(fileURL foundation.IURL, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("writeToURL:error:"), fileURL, outError)
	return rv
}


// An array of the workflow’s actions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflow/actions
func (a_ AMWorkflow) Actions() []AMAction {
	rv := objc.Send[[]AMAction](a_.ID, objc.Sel("actions"))
	return rv
}


// A URL that specifies the location of the workflow file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflow/fileURL
func (a_ AMWorkflow) FileURL() foundation.URL {
	rv := objc.Send[foundation.URL](a_.ID, objc.Sel("fileURL"))
	return rv
}


// The input data that is passed to the first action in the workflow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflow/input
func (a_ AMWorkflow) Input() objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("input"))
	return rv
}


// The input data that is passed to the first action in the workflow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflow/input
func (a_ AMWorkflow) SetInput(value objc.ID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setInput:"), value)
}


// The output data that is provided by the last action in the workflow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflow/output
func (a_ AMWorkflow) Output() objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("output"))
	return rv
}


