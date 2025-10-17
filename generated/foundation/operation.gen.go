// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Operation] class.
var OperationClass objc.Class

func init() {
	OperationClass = objc.GetClass("NSOperation")
}

type Operation struct {
	objc.ID
}

func OperationFrom(ptr unsafe.Pointer) Operation {
	return Operation{
		ID: objc.ID(ptr),
	}
}


// Makes the receiver dependent on the completion of the specified operation. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/Operation/addDependency(_:)
func (o_ Operation) AddDependency(op unsafe.Pointer) {
	sel := objc.RegisterName("addDependency:")
	o_.ID.Send(sel, op)
}
// Advises the operation object that it should stop executing its task. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/Operation/cancel()
func (o_ Operation) Cancel() {
	sel := objc.RegisterName("cancel")
	o_.ID.Send(sel)
}
// Performs the receiver’s non-concurrent task. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/Operation/main()
func (o_ Operation) Main() {
	sel := objc.RegisterName("main")
	o_.ID.Send(sel)
}
// Removes the receiver’s dependence on the specified operation. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/Operation/removeDependency(_:)
func (o_ Operation) RemoveDependency(op unsafe.Pointer) {
	sel := objc.RegisterName("removeDependency:")
	o_.ID.Send(sel, op)
}
// Begins the execution of the operation. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/Operation/start()
func (o_ Operation) Start() {
	sel := objc.RegisterName("start")
	o_.ID.Send(sel)
}
// Blocks execution of the current thread until the operation object finishes its task. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/Operation/waitUntilFinished()
func (o_ Operation) WaitUntilFinished() {
	sel := objc.RegisterName("waitUntilFinished")
	o_.ID.Send(sel)
}

