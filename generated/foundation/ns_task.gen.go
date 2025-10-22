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
	LaunchAndReturnError(error_ IError) bool
	Terminate()
	Environment() unsafe.Pointer
	SetEnvironment(value unsafe.Pointer)
	ExecutableURL() URL
	SetExecutableURL(value IURL)
	LaunchRequirementData() NSData
	SetLaunchRequirementData(value IData)
	ProcessIdentifier() int
	Arguments() string
	SetArguments(value string)
	CurrentDirectoryPath() string
	SetCurrentDirectoryPath(value string)
	CurrentDirectoryURL() URL
	SetCurrentDirectoryURL(value IURL)
	IsRunning() bool
	SetIsRunning(value bool)
	LaunchPath() string
	SetLaunchPath(value string)
	LaunchRequirement() unsafe.Pointer
	SetLaunchRequirement(value unsafe.Pointer)
	QualityOfService() QualityOfService
	SetQualityOfService(value IQualityOfService)
	StandardError() unsafe.Pointer
	SetStandardError(value unsafe.Pointer)
	StandardInput() unsafe.Pointer
	SetStandardInput(value unsafe.Pointer)
	StandardOutput() unsafe.Pointer
	SetStandardOutput(value unsafe.Pointer)
	TerminationHandler() unsafe.Pointer
	SetTerminationHandler(value unsafe.Pointer)
	TerminationReason() unsafe.Pointer
	SetTerminationReason(value unsafe.Pointer)
	TerminationStatus() unsafe.Pointer
	SetTerminationStatus(value unsafe.Pointer)
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




// Runs the process with the current environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/run()

func (t_ Task) LaunchAndReturnError(error_ IError) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("launchAndReturnError:"), error_)
	return rv
}


// Sends a terminate signal to the receiver and all of its subtasks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/terminate()

func (t_ Task) Terminate() {
	objc.Send[objc.ID](t_.ID, objc.Sel("terminate"))
}


// The environment for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/environment

func (t_ Task) Environment() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("environment"))
	return rv
}


// The environment for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/environment

func (t_ Task) SetEnvironment(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setEnvironment:"), value)
}


// The receiver’s executable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/executableURL

func (t_ Task) ExecutableURL() URL {
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


//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/launchRequirementData

func (t_ Task) LaunchRequirementData() NSData {
	rv := objc.Send[NSData](t_.ID, objc.Sel("launchRequirementData"))
	return rv
}


//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/launchRequirementData

func (t_ Task) SetLaunchRequirementData(value IData) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLaunchRequirementData:"), value)
}


// The receiver’s process identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/processIdentifier

func (t_ Task) ProcessIdentifier() int {
	rv := objc.Send[int](t_.ID, objc.Sel("processIdentifier"))
	return rv
}


// The command arguments that the system uses to launch the executable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/arguments

func (t_ Task) Arguments() string {
	rv := objc.Send[string](t_.ID, objc.Sel("arguments"))
	return rv
}


// The command arguments that the system uses to launch the executable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/arguments

func (t_ Task) SetArguments(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setArguments:"), objc.String(value))
}


// Sets the current directory for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/currentdirectorypath

func (t_ Task) CurrentDirectoryPath() string {
	rv := objc.Send[string](t_.ID, objc.Sel("currentDirectoryPath"))
	return rv
}


// Sets the current directory for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/currentdirectorypath

func (t_ Task) SetCurrentDirectoryPath(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCurrentDirectoryPath:"), objc.String(value))
}


// The current directory for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/currentdirectoryurl

func (t_ Task) CurrentDirectoryURL() URL {
	rv := objc.Send[URL](t_.ID, objc.Sel("currentDirectoryURL"))
	return rv
}


// The current directory for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/currentdirectoryurl

func (t_ Task) SetCurrentDirectoryURL(value IURL) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCurrentDirectoryURL:"), value)
}


// A status that indicates whether the receiver is still running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/isrunning

func (t_ Task) IsRunning() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isRunning"))
	return rv
}


// A status that indicates whether the receiver is still running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/isrunning

func (t_ Task) SetIsRunning(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsRunning:"), value)
}


// Sets the receiver’s executable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/launchpath

func (t_ Task) LaunchPath() string {
	rv := objc.Send[string](t_.ID, objc.Sel("launchPath"))
	return rv
}


// Sets the receiver’s executable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/launchpath

func (t_ Task) SetLaunchPath(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLaunchPath:"), objc.String(value))
}


//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/launchrequirement

func (t_ Task) LaunchRequirement() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("launchRequirement"))
	return rv
}


//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/launchrequirement

func (t_ Task) SetLaunchRequirement(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLaunchRequirement:"), value)
}


// The default quality of service level the system applies to operations the task executes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/qualityofservice

func (t_ Task) QualityOfService() QualityOfService {
	rv := objc.Send[QualityOfService](t_.ID, objc.Sel("qualityOfService"))
	return rv
}


// The default quality of service level the system applies to operations the task executes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/qualityofservice

func (t_ Task) SetQualityOfService(value IQualityOfService) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setQualityOfService:"), value)
}


// The standard error for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/standarderror

func (t_ Task) StandardError() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("standardError"))
	return rv
}


// The standard error for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/standarderror

func (t_ Task) SetStandardError(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStandardError:"), value)
}


// The standard input for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/standardinput

func (t_ Task) StandardInput() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("standardInput"))
	return rv
}


// The standard input for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/standardinput

func (t_ Task) SetStandardInput(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStandardInput:"), value)
}


// The standard output for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/standardoutput

func (t_ Task) StandardOutput() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("standardOutput"))
	return rv
}


// The standard output for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/standardoutput

func (t_ Task) SetStandardOutput(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStandardOutput:"), value)
}


// A completion block the system invokes when the task completes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/terminationhandler

func (t_ Task) TerminationHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("terminationHandler"))
	return rv
}


// A completion block the system invokes when the task completes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/terminationhandler

func (t_ Task) SetTerminationHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTerminationHandler:"), value)
}


// The reason the system terminated the task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/terminationreason-swift.property

func (t_ Task) TerminationReason() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("terminationReason"))
	return rv
}


// The reason the system terminated the task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/terminationreason-swift.property

func (t_ Task) SetTerminationReason(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTerminationReason:"), value)
}


// The exit status the receiver’s executable returns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/terminationstatus

func (t_ Task) TerminationStatus() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("terminationStatus"))
	return rv
}


// The exit status the receiver’s executable returns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/process/terminationstatus

func (t_ Task) SetTerminationStatus(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTerminationStatus:"), value)
}


