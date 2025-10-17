// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Task] class.
var TaskClass objc.Class

func init() {
	TaskClass = objc.GetClass("NSTask")
}

type Task struct {
	objc.ID
}

func TaskFrom(ptr unsafe.Pointer) Task {
	return Task{
		ID: objc.ID(ptr),
	}
}


// Runs the process with the current environment. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/Process/run()
func (t_ Task) LaunchAndReturnError(error unsafe.Pointer) bool {
	sel := objc.RegisterName("launchAndReturnError:")
	ret := t_.ID.Send(sel, error)
	return ret != 0
}
// Sends a terminate signal to the receiver and all of its subtasks. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/Process/terminate()
func (t_ Task) Terminate() {
	sel := objc.RegisterName("terminate")
	t_.ID.Send(sel)
}


