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
	LaunchAndReturnError(error_ unsafe.Pointer) bool
	Terminate()
}

// An object that represents a subprocess of the current process.
//
// Using this class, your program can run another program as a subprocess and monitor that program’s execution. Unlike , it doesn’t share memory space with the process that creates it. A process operates within an environment defined by the current values for several items: the current directory, standard input, standard output, standard error, and the values of any environment variables, inheriting its environment from the process that launches it. If there are any environment variables that should be different for the subprocess (for example, if the current directory needs to change), change it in the instance after initialization, before your app launches it. Your app can’t change a process’s environment while it’s running. You can only run the subprocess once per instance. Subsequent attempts raise an error.
//
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



// Runs the process with the current environment.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/run()
func (t_ Task) LaunchAndReturnError(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("launchAndReturnError:"), error_)
	return rv
}

// Sends a terminate signal to the receiver and all of its subtasks.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/terminate()
func (t_ Task) Terminate() {
	objc.Send[objc.ID](t_.ID, objc.Sel("terminate"))
}

// A status that indicates whether the receiver is still running.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/isrunning
func (t_ Task) IsRunning() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isRunning"))
	return rv
}


// SetIsRunning sets the value of the isRunning property.
// A status that indicates whether the receiver is still running.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/isrunning
func (t_ Task) SetIsRunning(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsRunning:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/launchrequirement
func (t_ Task) LaunchRequirement() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("launchRequirement"))
	return rv
}


// SetLaunchRequirement sets the value of the launchRequirement property.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/launchrequirement
func (t_ Task) SetLaunchRequirement(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLaunchRequirement:"), value)
}

// A completion block the system invokes when the task completes.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/terminationhandler
func (t_ Task) TerminationHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("terminationHandler"))
	return rv
}


// SetTerminationHandler sets the value of the terminationHandler property.
// A completion block the system invokes when the task completes.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/terminationhandler
func (t_ Task) SetTerminationHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTerminationHandler:"), value)
}

// The standard error for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/standarderror
func (t_ Task) StandardError() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("standardError"))
	return rv
}


// SetStandardError sets the value of the standardError property.
// The standard error for the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/standarderror
func (t_ Task) SetStandardError(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStandardError:"), value)
}

// The reason the system terminated the task.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/terminationreason-swift.property
func (t_ Task) TerminationReason() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("terminationReason"))
	return rv
}


// SetTerminationReason sets the value of the terminationReason property.
// The reason the system terminated the task.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/terminationreason-swift.property
func (t_ Task) SetTerminationReason(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTerminationReason:"), value)
}

// The current directory for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/currentdirectoryurl
func (t_ Task) CurrentDirectoryURL() URL {
	rv := objc.Send[URL](t_.ID, objc.Sel("currentDirectoryURL"))
	return rv
}


// SetCurrentDirectoryURL sets the value of the currentDirectoryURL property.
// The current directory for the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/currentdirectoryurl
func (t_ Task) SetCurrentDirectoryURL(value URL) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCurrentDirectoryURL:"), value)
}

// The command arguments that the system uses to launch the executable.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/arguments
func (t_ Task) Arguments() string {
	rv := objc.Send[string](t_.ID, objc.Sel("arguments"))
	return rv
}


// SetArguments sets the value of the arguments property.
// The command arguments that the system uses to launch the executable.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/arguments
func (t_ Task) SetArguments(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setArguments:"), objc.String(value))
}

// The standard input for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/standardinput
func (t_ Task) StandardInput() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("standardInput"))
	return rv
}


// SetStandardInput sets the value of the standardInput property.
// The standard input for the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/standardinput
func (t_ Task) SetStandardInput(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStandardInput:"), value)
}

// The exit status the receiver’s executable returns.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/terminationstatus
func (t_ Task) TerminationStatus() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("terminationStatus"))
	return rv
}


// SetTerminationStatus sets the value of the terminationStatus property.
// The exit status the receiver’s executable returns.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/terminationstatus
func (t_ Task) SetTerminationStatus(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTerminationStatus:"), value)
}

// Sets the receiver’s executable.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/launchpath
func (t_ Task) LaunchPath() string {
	rv := objc.Send[string](t_.ID, objc.Sel("launchPath"))
	return rv
}


// SetLaunchPath sets the value of the launchPath property.
// Sets the receiver’s executable.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/launchpath
func (t_ Task) SetLaunchPath(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLaunchPath:"), objc.String(value))
}

// Sets the current directory for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/currentdirectorypath
func (t_ Task) CurrentDirectoryPath() string {
	rv := objc.Send[string](t_.ID, objc.Sel("currentDirectoryPath"))
	return rv
}


// SetCurrentDirectoryPath sets the value of the currentDirectoryPath property.
// Sets the current directory for the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/currentdirectorypath
func (t_ Task) SetCurrentDirectoryPath(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCurrentDirectoryPath:"), objc.String(value))
}

// The default quality of service level the system applies to operations the task executes.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/qualityofservice
func (t_ Task) QualityOfService() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("qualityOfService"))
	return rv
}


// SetQualityOfService sets the value of the qualityOfService property.
// The default quality of service level the system applies to operations the task executes.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/qualityofservice
func (t_ Task) SetQualityOfService(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setQualityOfService:"), value)
}

// The standard output for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/standardoutput
func (t_ Task) StandardOutput() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("standardOutput"))
	return rv
}


// SetStandardOutput sets the value of the standardOutput property.
// The standard output for the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/standardoutput
func (t_ Task) SetStandardOutput(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStandardOutput:"), value)
}

// The environment for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/environment
func (t_ Task) Environment() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("environment"))
	return rv
}


// SetEnvironment sets the value of the environment property.
// The environment for the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/environment
func (t_ Task) SetEnvironment(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setEnvironment:"), value)
}

// The receiver’s executable.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/executableURL
func (t_ Task) ExecutableURL() URL {
	rv := objc.Send[URL](t_.ID, objc.Sel("executableURL"))
	return rv
}


// SetExecutableURL sets the value of the executableURL property.
// The receiver’s executable.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/executableURL
func (t_ Task) SetExecutableURL(value URL) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setExecutableURL:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/launchRequirementData
func (t_ Task) LaunchRequirementData() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("launchRequirementData"))
	return rv
}


// SetLaunchRequirementData sets the value of the launchRequirementData property.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/launchRequirementData
func (t_ Task) SetLaunchRequirementData(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLaunchRequirementData:"), value)
}

// The receiver’s process identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/processIdentifier
func (t_ Task) ProcessIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("processIdentifier"))
	return rv
}


