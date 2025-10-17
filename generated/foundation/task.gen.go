// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Task] class.
var taskClass = _TaskClass{objc.GetClass("NSTask")}

type _TaskClass struct {
	class objc.Class
}

// An interface definition for the [Task] class.
type ITask interface {
	objectivec.IObject
	LaunchAndReturnError(error unsafe.Pointer) bool
	Terminate()
}

// An object that represents a subprocess of the current process. [Full Topic]
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

// Runs the process with the current environment. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/run()
func (t_ Task) LaunchAndReturnError(error unsafe.Pointer) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("launchAndReturnError:"), error)
	return rv
}
// Sends a terminate signal to the receiver and all of its subtasks. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Process/terminate()
func (t_ Task) Terminate() {
	objc.Send[objc.ID](t_.ID, objc.Sel("terminate"))
}


