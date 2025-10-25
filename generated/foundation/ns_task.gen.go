// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSTask */


/* debug [class_header]: Header for NSTask */
// The class instance for the [Task] class.
var (
	TaskClass     _TaskClass
	TaskClassOnce sync.Once
)

func getTaskClass() _TaskClass {
	TaskClassOnce.Do(func() {
		TaskClass = _TaskClass{objc.GetClass("NSTask")}
	})
	return TaskClass
}

type _TaskClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Task */
// An interface definition for the [Task] class.
type ITask interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Task */
	// properties:
	Arguments() []string
	SetArguments(value []string)
	CurrentDirectoryPath() IString
	SetCurrentDirectoryPath(value IString)
	CurrentDirectoryURL() IURL
	SetCurrentDirectoryURL(value IURL)
	Environment() IDictionary
	SetEnvironment(value IDictionary)
	ExecutableURL() IURL
	SetExecutableURL(value IURL)
	Running() bool
	LaunchPath() IString
	SetLaunchPath(value IString)
	LaunchRequirementData() IData
	SetLaunchRequirementData(value IData)
	ProcessIdentifier() int
	QualityOfService() QualityOfService
	SetQualityOfService(value QualityOfService)
	StandardError() objc.ID
	SetStandardError(value objc.ID)
	StandardInput() objc.ID
	SetStandardInput(value objc.ID)
	StandardOutput() objc.ID
	SetStandardOutput(value objc.ID)
	TerminationHandler() func(unsafe.Pointer)
	SetTerminationHandler(value func(unsafe.Pointer))
	TerminationReason() TaskTerminationReason
	TerminationStatus() int
	IsRunning() bool
	SetIsRunning(value bool)
	LaunchRequirement() objectivec.IObject
	SetLaunchRequirement(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Task */
	// methods:
	Interrupt()
	Resume() bool
	LaunchAndReturnError(error_ IError) bool
	Suspend() bool
	Terminate()
	WaitUntilExit()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Task */
// Alloc allocates a new instance without initialization.
func (tc _TaskClass) Alloc() Task {
	rv := objc.Send[Task](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TaskClass) New() Task {
	rv := objc.Send[Task](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ Task) Init() Task {
	rv := objc.Send[Task](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ Task) Autorelease() Task {
	rv := objc.Send[Task](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTask creates a new Task instance.
func NewTask() Task {
	return getTaskClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Task */
// An object that represents a subprocess of the current process.
//
// Using this class, your program can run another program as a subprocess and monitor that program’s execution. Unlike , it doesn’t share memory space with the process that creates it. A process operates within an environment defined by the current values for several items: the current directory, standard input, standard output, standard error, and the values of any environment variables, inheriting its environment from the process that launches it. If there are any environment variables that should be different for the subprocess (for example, if the current directory needs to change), change it in the instance after initialization, before your app launches it. Your app can’t change a process’s environment while it’s running. You can only run the subprocess once per instance. Subsequent attempts raise an error.


// An object that represents a subprocess of the current process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process
type Task struct {
	objectivec.Object
}

// TaskFrom constructs a [Task] from an unsafe.Pointer.
//
// An object that represents a subprocess of the current process.
func TaskFrom(ptr unsafe.Pointer) Task {
	return Task{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Task */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Task */

// Creates and launches a task with a specified executable and arguments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/launchedProcess(launchPath:arguments:)
func (tc _TaskClass) LaunchedTaskWithLaunchPathArguments(path IString, arguments []string) ITask {
	rv := objc.Send[Task](objc.ID(tc.class), objc.Sel("launchedTaskWithLaunchPath:arguments:"), path, arguments)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LaunchedTaskWithLaunchPathArguments) */


// Creates and runs a task with a specified executable and arguments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/run(_:arguments:terminationHandler:)
func (tc _TaskClass) LaunchedTaskWithExecutableURLArgumentsErrorTerminationHandler(url IURL, arguments []string, error_ IError, terminationHandler unsafe.Pointer) ITask {
	rv := objc.Send[Task](objc.ID(tc.class), objc.Sel("launchedTaskWithExecutableURL:arguments:error:terminationHandler:"), url, arguments, error_, terminationHandler)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LaunchedTaskWithExecutableURLArgumentsErrorTerminationHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Task */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Task */

// Sends an interrupt signal to the receiver and all of its subtasks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/interrupt()
func (t_ Task) Interrupt() {
	objc.Send[objc.ID](t_.ID, objc.Sel("interrupt"))
}/* debug [instance_methods/method]: Interrupt */


// Resumes execution of a suspended task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/resume()
func (t_ Task) Resume() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("resume"))
	return rv
}/* debug [instance_methods/method]: Resume */


// Runs the process with the current environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/run()
func (t_ Task) LaunchAndReturnError(error_ IError) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("launchAndReturnError:"), error_)
	return rv
}/* debug [instance_methods/method]: LaunchAndReturnError */


// Suspends execution of the receiver task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/suspend()
func (t_ Task) Suspend() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("suspend"))
	return rv
}/* debug [instance_methods/method]: Suspend */


// Sends a terminate signal to the receiver and all of its subtasks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/terminate()
func (t_ Task) Terminate() {
	objc.Send[objc.ID](t_.ID, objc.Sel("terminate"))
}/* debug [instance_methods/method]: Terminate */


// Blocks the process until the receiver is finished.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/waitUntilExit()
func (t_ Task) WaitUntilExit() {
	objc.Send[objc.ID](t_.ID, objc.Sel("waitUntilExit"))
}/* debug [instance_methods/method]: WaitUntilExit */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Task */

// The command arguments that the system uses to launch the executable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/arguments
func (t_ Task) Arguments() []string {
	rv := objc.Send[[]string](t_.ID, objc.Sel("arguments"))
	return rv
}/* debug [instance_properties/getter]: arguments */


// The command arguments that the system uses to launch the executable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/arguments
func (t_ Task) SetArguments(value []string) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](t_.ID, objc.Sel("setArguments:"), nsArray)
}/* debug [instance_properties/setter]: arguments */


// Sets the current directory for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/currentDirectoryPath
func (t_ Task) CurrentDirectoryPath() IString {
	rv := objc.Send[String](t_.ID, objc.Sel("currentDirectoryPath"))
	return rv
}/* debug [instance_properties/getter]: currentDirectoryPath */


// Sets the current directory for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/currentDirectoryPath
func (t_ Task) SetCurrentDirectoryPath(value IString) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCurrentDirectoryPath:"), value)
}/* debug [instance_properties/setter]: currentDirectoryPath */


// The current directory for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/currentDirectoryURL
func (t_ Task) CurrentDirectoryURL() IURL {
	rv := objc.Send[URL](t_.ID, objc.Sel("currentDirectoryURL"))
	return rv
}/* debug [instance_properties/getter]: currentDirectoryURL */


// The current directory for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/currentDirectoryURL
func (t_ Task) SetCurrentDirectoryURL(value IURL) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCurrentDirectoryURL:"), value)
}/* debug [instance_properties/setter]: currentDirectoryURL */


// The environment for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/environment
func (t_ Task) Environment() IDictionary {
	rv := objc.Send[Dictionary](t_.ID, objc.Sel("environment"))
	return rv
}/* debug [instance_properties/getter]: environment */


// The environment for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/environment
func (t_ Task) SetEnvironment(value IDictionary) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setEnvironment:"), value)
}/* debug [instance_properties/setter]: environment */


// The receiver’s executable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/executableURL
func (t_ Task) ExecutableURL() IURL {
	rv := objc.Send[URL](t_.ID, objc.Sel("executableURL"))
	return rv
}/* debug [instance_properties/getter]: executableURL */


// The receiver’s executable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/executableURL
func (t_ Task) SetExecutableURL(value IURL) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setExecutableURL:"), value)
}/* debug [instance_properties/setter]: executableURL */


// A status that indicates whether the receiver is still running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/isRunning
func (t_ Task) Running() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("running"))
	return rv
}/* debug [instance_properties/getter]: running */


// Sets the receiver’s executable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/launchPath
func (t_ Task) LaunchPath() IString {
	rv := objc.Send[String](t_.ID, objc.Sel("launchPath"))
	return rv
}/* debug [instance_properties/getter]: launchPath */


// Sets the receiver’s executable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/launchPath
func (t_ Task) SetLaunchPath(value IString) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLaunchPath:"), value)
}/* debug [instance_properties/setter]: launchPath */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/launchRequirementData
func (t_ Task) LaunchRequirementData() IData {
	rv := objc.Send[Data](t_.ID, objc.Sel("launchRequirementData"))
	return rv
}/* debug [instance_properties/getter]: launchRequirementData */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/launchRequirementData
func (t_ Task) SetLaunchRequirementData(value IData) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLaunchRequirementData:"), value)
}/* debug [instance_properties/setter]: launchRequirementData */


// The receiver’s process identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/processIdentifier
func (t_ Task) ProcessIdentifier() int {
	rv := objc.Send[int](t_.ID, objc.Sel("processIdentifier"))
	return rv
}/* debug [instance_properties/getter]: processIdentifier */


// The default quality of service level the system applies to operations the task executes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/qualityOfService
func (t_ Task) QualityOfService() QualityOfService {
	rv := objc.Send[QualityOfService](t_.ID, objc.Sel("qualityOfService"))
	return rv
}/* debug [instance_properties/getter]: qualityOfService */


// The default quality of service level the system applies to operations the task executes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/qualityOfService
func (t_ Task) SetQualityOfService(value QualityOfService) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setQualityOfService:"), value)
}/* debug [instance_properties/setter]: qualityOfService */


// The standard error for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/standardError
func (t_ Task) StandardError() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("standardError"))
	return rv
}/* debug [instance_properties/getter]: standardError */


// The standard error for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/standardError
func (t_ Task) SetStandardError(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStandardError:"), value)
}/* debug [instance_properties/setter]: standardError */


// The standard input for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/standardInput
func (t_ Task) StandardInput() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("standardInput"))
	return rv
}/* debug [instance_properties/getter]: standardInput */


// The standard input for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/standardInput
func (t_ Task) SetStandardInput(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStandardInput:"), value)
}/* debug [instance_properties/setter]: standardInput */


// The standard output for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/standardOutput
func (t_ Task) StandardOutput() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("standardOutput"))
	return rv
}/* debug [instance_properties/getter]: standardOutput */


// The standard output for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/standardOutput
func (t_ Task) SetStandardOutput(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStandardOutput:"), value)
}/* debug [instance_properties/setter]: standardOutput */


// A completion block the system invokes when the task completes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/terminationHandler
func (t_ Task) TerminationHandler() func(unsafe.Pointer) {
	rv := objc.Send[func(unsafe.Pointer)](t_.ID, objc.Sel("terminationHandler"))
	return rv
}/* debug [instance_properties/getter]: terminationHandler */


// A completion block the system invokes when the task completes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/terminationHandler
func (t_ Task) SetTerminationHandler(value func(unsafe.Pointer)) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTerminationHandler:"), value)
}/* debug [instance_properties/setter]: terminationHandler */


// The reason the system terminated the task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/terminationReason-swift.property
func (t_ Task) TerminationReason() TaskTerminationReason {
	rv := objc.Send[TaskTerminationReason](t_.ID, objc.Sel("terminationReason"))
	return rv
}/* debug [instance_properties/getter]: terminationReason */


// The exit status the receiver’s executable returns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/terminationStatus
func (t_ Task) TerminationStatus() int {
	rv := objc.Send[int](t_.ID, objc.Sel("terminationStatus"))
	return rv
}/* debug [instance_properties/getter]: terminationStatus */


// A status that indicates whether the receiver is still running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/isrunning
func (t_ Task) IsRunning() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isRunning"))
	return rv
}/* debug [instance_properties/getter]: isRunning */


// A status that indicates whether the receiver is still running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/isrunning
func (t_ Task) SetIsRunning(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsRunning:"), value)
}/* debug [instance_properties/setter]: isRunning */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/launchrequirement
func (t_ Task) LaunchRequirement() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](t_.ID, objc.Sel("launchRequirement"))
	return rv
}/* debug [instance_properties/getter]: launchRequirement */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/launchrequirement
func (t_ Task) SetLaunchRequirement(value objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLaunchRequirement:"), value)
}/* debug [instance_properties/setter]: launchRequirement */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSTask */


