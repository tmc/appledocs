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
func (t_ Task) ExecutableURL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("executableURL"))
	return rv
}


// SetExecutableURL sets the value of the executableURL property.
// The receiver’s executable.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/executableURL
func (t_ Task) SetExecutableURL(value unsafe.Pointer) {
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


