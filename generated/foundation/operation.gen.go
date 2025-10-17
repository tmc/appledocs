// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Operation] class.
var OperationClass = _OperationClass{objc.GetClass("NSOperation")}

type _OperationClass struct {
	class objc.Class
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
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/addDependency(_:)
func (o_ Operation) AddDependency(op unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("addDependency:"), op)
}
// Advises the operation object that it should stop executing its task. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/cancel()
func (o_ Operation) Cancel() {
	objc.Send[objc.ID](o_.ID, objc.Sel("cancel"))
}
// Performs the receiver’s non-concurrent task. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/main()
func (o_ Operation) Main() {
	objc.Send[objc.ID](o_.ID, objc.Sel("main"))
}
// Removes the receiver’s dependence on the specified operation. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/removeDependency(_:)
func (o_ Operation) RemoveDependency(op unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("removeDependency:"), op)
}
// Begins the execution of the operation. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/start()
func (o_ Operation) Start() {
	objc.Send[objc.ID](o_.ID, objc.Sel("start"))
}
// Blocks execution of the current thread until the operation object finishes its task. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/waitUntilFinished()
func (o_ Operation) WaitUntilFinished() {
	objc.Send[objc.ID](o_.ID, objc.Sel("waitUntilFinished"))
}


