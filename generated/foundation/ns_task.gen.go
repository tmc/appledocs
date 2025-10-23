// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [Task] class.
type ITask interface {
	objectivec.IObject
	// properties:
	Arguments() []string /* primitive/slice/pointer */
	SetArguments(value []string /* primitive/slice/pointer */)
	CurrentDirectoryPath() string /* primitive/slice/pointer */
	SetCurrentDirectoryPath(value string /* primitive/slice/pointer */)
	CurrentDirectoryURL() IURL
	SetCurrentDirectoryURL(value IURL)
	Environment() IDictionary /* already interface */
	SetEnvironment(value IDictionary /* already interface */)
	ExecutableURL() IURL
	SetExecutableURL(value IURL)
	Running() bool /* primitive/slice/pointer */
	LaunchPath() string /* primitive/slice/pointer */
	SetLaunchPath(value string /* primitive/slice/pointer */)
	LaunchRequirementData() IData
	SetLaunchRequirementData(value IData)
	ProcessIdentifier() int /* primitive/slice/pointer */
	QualityOfService() QualityOfService
	SetQualityOfService(value QualityOfService)
	StandardError() objc.ID
	SetStandardError(value objc.ID)
	StandardInput() objc.ID
	SetStandardInput(value objc.ID)
	StandardOutput() objc.ID
	SetStandardOutput(value objc.ID)
	TerminationHandler() unsafe.Pointer
	SetTerminationHandler(value unsafe.Pointer)
	TerminationReason() TaskTerminationReason
	TerminationStatus() int /* primitive/slice/pointer */
	IsRunning() bool /* primitive/slice/pointer */
	SetIsRunning(value bool /* primitive/slice/pointer */)
	LaunchRequirement() unsafe.Pointer
	SetLaunchRequirement(value unsafe.Pointer)
	// methods:
	Interrupt()
	Resume() bool /* primitive/slice/pointer */
	LaunchAndReturnError(error_ IError) bool /* primitive/slice/pointer */
	Suspend() bool /* primitive/slice/pointer */
	Terminate()
	WaitUntilExit()
}

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

// Alloc allocates a new instance without initialization.
func (tc _TaskClass) Alloc() Task {
	rv := objc.Send[Task](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates and launches a task with a specified executable and arguments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/launchedProcess(launchPath:arguments:)
func (tc _TaskClass) LaunchedTaskWithLaunchPathArguments(path string /* primitive/slice/pointer */, arguments []string /* primitive/slice/pointer */) ITask {
	rv := objc.Send[Task](objc.ID(tc.class), objc.Sel("launchedTaskWithLaunchPath:arguments:"), objc.String(path), arguments)
	return rv
}


// Creates and runs a task with a specified executable and arguments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/run(_:arguments:terminationHandler:)
func (tc _TaskClass) LaunchedTaskWithExecutableURLArgumentsErrorTerminationHandler(url IURL, arguments []string /* primitive/slice/pointer */, error_ IError, terminationHandler unsafe.Pointer) ITask {
	rv := objc.Send[Task](objc.ID(tc.class), objc.Sel("launchedTaskWithExecutableURL:arguments:error:terminationHandler:"), url, arguments, error_, terminationHandler)
	return rv
}


// Sends an interrupt signal to the receiver and all of its subtasks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/interrupt()
func (t_ Task) Interrupt() {
	objc.Send[objc.ID](t_.ID, objc.Sel("interrupt"))
}


// Resumes execution of a suspended task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/resume()
func (t_ Task) Resume() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("resume"))
	return rv
}


// Runs the process with the current environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/run()
func (t_ Task) LaunchAndReturnError(error_ IError) bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("launchAndReturnError:"), error_)
	return rv
}


// Suspends execution of the receiver task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/suspend()
func (t_ Task) Suspend() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("suspend"))
	return rv
}


// Sends a terminate signal to the receiver and all of its subtasks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/terminate()
func (t_ Task) Terminate() {
	objc.Send[objc.ID](t_.ID, objc.Sel("terminate"))
}


// Blocks the process until the receiver is finished.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/waitUntilExit()
func (t_ Task) WaitUntilExit() {
	objc.Send[objc.ID](t_.ID, objc.Sel("waitUntilExit"))
}


// The command arguments that the system uses to launch the executable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/arguments
func (t_ Task) Arguments() []string /* primitive/slice/pointer */ {
	rv := objc.Send[[]string](t_.ID, objc.Sel("arguments"))
	return rv
}


// The command arguments that the system uses to launch the executable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/arguments
func (t_ Task) SetArguments(value []string /* primitive/slice/pointer */) {
	// Convert Go slice to NSArray
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
}


// Sets the current directory for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/currentDirectoryPath
func (t_ Task) CurrentDirectoryPath() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](t_.ID, objc.Sel("currentDirectoryPath"))
	return rv
}


// Sets the current directory for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/currentDirectoryPath
func (t_ Task) SetCurrentDirectoryPath(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCurrentDirectoryPath:"), objc.String(value))
}


// The current directory for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/currentDirectoryURL
func (t_ Task) CurrentDirectoryURL() IURL {
	rv := objc.Send[URL](t_.ID, objc.Sel("currentDirectoryURL"))
	return rv
}


// The current directory for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/currentDirectoryURL
func (t_ Task) SetCurrentDirectoryURL(value IURL) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCurrentDirectoryURL:"), value)
}


// The environment for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/environment
func (t_ Task) Environment() IDictionary /* already interface */ {
	rv := objc.Send[IDictionary](t_.ID, objc.Sel("environment"))
	return rv
}


// The environment for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/environment
func (t_ Task) SetEnvironment(value IDictionary /* already interface */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setEnvironment:"), value)
}


// The receiver’s executable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/executableURL
func (t_ Task) ExecutableURL() IURL {
	rv := objc.Send[URL](t_.ID, objc.Sel("executableURL"))
	return rv
}


// The receiver’s executable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/executableURL
func (t_ Task) SetExecutableURL(value IURL) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setExecutableURL:"), value)
}


// A status that indicates whether the receiver is still running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/isRunning
func (t_ Task) Running() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("running"))
	return rv
}


// Sets the receiver’s executable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/launchPath
func (t_ Task) LaunchPath() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](t_.ID, objc.Sel("launchPath"))
	return rv
}


// Sets the receiver’s executable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/launchPath
func (t_ Task) SetLaunchPath(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLaunchPath:"), objc.String(value))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/launchRequirementData
func (t_ Task) LaunchRequirementData() IData {
	rv := objc.Send[Data](t_.ID, objc.Sel("launchRequirementData"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/launchRequirementData
func (t_ Task) SetLaunchRequirementData(value IData) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLaunchRequirementData:"), value)
}


// The receiver’s process identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/processIdentifier
func (t_ Task) ProcessIdentifier() int /* primitive/slice/pointer */ {
	rv := objc.Send[int](t_.ID, objc.Sel("processIdentifier"))
	return rv
}


// The default quality of service level the system applies to operations the task executes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/qualityOfService
func (t_ Task) QualityOfService() QualityOfService {
	rv := objc.Send[QualityOfService](t_.ID, objc.Sel("qualityOfService"))
	return rv
}


// The default quality of service level the system applies to operations the task executes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/qualityOfService
func (t_ Task) SetQualityOfService(value QualityOfService) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setQualityOfService:"), value)
}


// The standard error for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/standardError
func (t_ Task) StandardError() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("standardError"))
	return rv
}


// The standard error for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/standardError
func (t_ Task) SetStandardError(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStandardError:"), value)
}


// The standard input for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/standardInput
func (t_ Task) StandardInput() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("standardInput"))
	return rv
}


// The standard input for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/standardInput
func (t_ Task) SetStandardInput(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStandardInput:"), value)
}


// The standard output for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/standardOutput
func (t_ Task) StandardOutput() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("standardOutput"))
	return rv
}


// The standard output for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/standardOutput
func (t_ Task) SetStandardOutput(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStandardOutput:"), value)
}


// A completion block the system invokes when the task completes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/terminationHandler
func (t_ Task) TerminationHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("terminationHandler"))
	return rv
}


// A completion block the system invokes when the task completes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/terminationHandler
func (t_ Task) SetTerminationHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTerminationHandler:"), value)
}


// The reason the system terminated the task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/terminationReason-swift.property
func (t_ Task) TerminationReason() TaskTerminationReason {
	rv := objc.Send[TaskTerminationReason](t_.ID, objc.Sel("terminationReason"))
	return rv
}


// The exit status the receiver’s executable returns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/terminationStatus
func (t_ Task) TerminationStatus() int /* primitive/slice/pointer */ {
	rv := objc.Send[int](t_.ID, objc.Sel("terminationStatus"))
	return rv
}


// A status that indicates whether the receiver is still running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/isrunning
func (t_ Task) IsRunning() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("isRunning"))
	return rv
}


// A status that indicates whether the receiver is still running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/isrunning
func (t_ Task) SetIsRunning(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsRunning:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/launchrequirement
func (t_ Task) LaunchRequirement() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("launchRequirement"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/launchrequirement
func (t_ Task) SetLaunchRequirement(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLaunchRequirement:"), value)
}


