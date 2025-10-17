// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [OperationQueue] class.
var operationQueueClass = _OperationQueueClass{objc.GetClass("NSOperationQueue")}

type _OperationQueueClass struct {
	class objc.Class
}

// An interface definition for the [OperationQueue] class.
type IOperationQueue interface {
	objectivec.IObject
	CancelAllOperations()
}

// A queue that regulates the execution of operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue

type OperationQueue struct {
	objectivec.Object
}

// OperationQueueFrom constructs a [OperationQueue] from an unsafe.Pointer.
//
// A queue that regulates the execution of operations.
func OperationQueueFrom(ptr unsafe.Pointer) OperationQueue {
	return OperationQueue{objectivec.Object{objc.ID(ptr)}}
}

// Cancels all queued and executing operations. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/cancelAllOperations()
func (o_ OperationQueue) CancelAllOperations() {
	objc.Send[objc.ID](o_.ID, objc.Sel("cancelAllOperations"))
}


