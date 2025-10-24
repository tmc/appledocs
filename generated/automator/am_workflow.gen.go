// Code generated from Apple documentation for Automator. DO NOT EDIT.

package automator

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AMWorkflow */


/* debug [class_header]: Header for AMWorkflow */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AMWorkflow */
// An interface definition for the [AMWorkflow] class.
type IAMWorkflow interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AMWorkflow */
	// properties:
	Actions() []AMAction
	FileURL() objc.IObject /* cross-framework: NSURL */
	Input() objc.ID
	SetInput(value objc.ID)
	Output() objc.ID
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AMWorkflow */
	// methods:
	AddAction(action IAMAction)
	InsertActionAtIndex(action IAMAction, index uint)
	MoveActionAtIndexToIndex(startIndex uint, endIndex uint)
	RemoveAction(action IAMAction)
	SetValueForVariableWithName(value objc.IObject, variableName objc.IObject /* cross-framework: NSString */) bool
	ValueForVariableWithName(variableName objc.IObject /* cross-framework: NSString */) objc.ID
	WriteToURLError(fileURL objc.IObject /* cross-framework: NSURL */, outError unsafe.Pointer) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AMWorkflow */
// Alloc allocates a new instance without initialization.
func (ac _AMWorkflowClass) Alloc() AMWorkflow {
	rv := objc.Send[AMWorkflow](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AMWorkflow */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AMWorkflow */

// Creates and initializes a workflow based on the contents of the specified file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflow/init(contentsOf:)
func NewAMWorkflowWithContentsOfURLError(fileURL objc.IObject /* cross-framework: NSURL */, outError unsafe.Pointer) AMWorkflow {
	instance := getAMWorkflowClass().Alloc()
	rv := objc.Send[AMWorkflow](instance.ID, objc.Sel("initWithContentsOfURL:error:"), fileURL, outError)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAMWorkflowWithContentsOfURLError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AMWorkflow */

// Loads and runs the specified workflow file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflow/run(at:withInput:)
func (ac _AMWorkflowClass) RunWorkflowAtURLWithInputError(fileURL objc.IObject /* cross-framework: NSURL */, input objc.IObject, error_ unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("runWorkflowAtURL:withInput:error:"), fileURL, input, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RunWorkflowAtURLWithInputError) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AMWorkflow */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AMWorkflow */

// Adds the specified action at the end of the receiving workflow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflow/addAction(_:)
func (a_ AMWorkflow) AddAction(action IAMAction) {
	objc.Send[objc.ID](a_.ID, objc.Sel("addAction:"), action)
}/* debug [instance_methods/method]: AddAction */


// Inserts the specified action at the specified position of the receiving workflow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflow/insertAction(_:at:)
func (a_ AMWorkflow) InsertActionAtIndex(action IAMAction, index uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("insertAction:atIndex:"), action, index)
}/* debug [instance_methods/method]: InsertActionAtIndex */


// Moves the action from the specified start position to the specified end position in the receiving workflow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflow/moveAction(at:to:)
func (a_ AMWorkflow) MoveActionAtIndexToIndex(startIndex uint, endIndex uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("moveActionAtIndex:toIndex:"), startIndex, endIndex)
}/* debug [instance_methods/method]: MoveActionAtIndexToIndex */


// Removes the specified action from the workflow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflow/removeAction(_:)
func (a_ AMWorkflow) RemoveAction(action IAMAction) {
	objc.Send[objc.ID](a_.ID, objc.Sel("removeAction:"), action)
}/* debug [instance_methods/method]: RemoveAction */


// Sets the value of the workflow variable with the specified name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflow/setValue(_:forVariableWithName:)
func (a_ AMWorkflow) SetValueForVariableWithName(value objc.IObject, variableName objc.IObject /* cross-framework: NSString */) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setValue:forVariableWithName:"), value, variableName)
	return rv
}/* debug [instance_methods/method]: SetValueForVariableWithName */


// Returns the value of the workflow variable with the specified name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflow/valueForVariable(withName:)
func (a_ AMWorkflow) ValueForVariableWithName(variableName objc.IObject /* cross-framework: NSString */) objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("valueForVariableWithName:"), variableName)
	return rv
}/* debug [instance_methods/method]: ValueForVariableWithName */


// Writes the workflow to the specified file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflow/write(to:)
func (a_ AMWorkflow) WriteToURLError(fileURL objc.IObject /* cross-framework: NSURL */, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("writeToURL:error:"), fileURL, outError)
	return rv
}/* debug [instance_methods/method]: WriteToURLError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AMWorkflow */

// An array of the workflow’s actions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflow/actions
func (a_ AMWorkflow) Actions() []AMAction {
	rv := objc.Send[[]AMAction](a_.ID, objc.Sel("actions"))
	return rv
}/* debug [instance_properties/getter]: actions */


// A URL that specifies the location of the workflow file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflow/fileURL
func (a_ AMWorkflow) FileURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](a_.ID, objc.Sel("fileURL"))
	return rv
}/* debug [instance_properties/getter]: fileURL */


// The input data that is passed to the first action in the workflow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflow/input
func (a_ AMWorkflow) Input() objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("input"))
	return rv
}/* debug [instance_properties/getter]: input */


// The input data that is passed to the first action in the workflow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflow/input
func (a_ AMWorkflow) SetInput(value objc.ID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setInput:"), value)
}/* debug [instance_properties/setter]: input */


// The output data that is provided by the last action in the workflow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflow/output
func (a_ AMWorkflow) Output() objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("output"))
	return rv
}/* debug [instance_properties/getter]: output */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AMWorkflow */


