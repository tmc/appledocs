// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var operationQueueClass _OperationQueueClass

func init() {
	operationQueueClass = _OperationQueueClass{objc.GetClass("NSOperationQueue")}
}

type _OperationQueueClass struct {
	class objc.Class
}

type OperationQueue struct {
	objc.ID
}

func OperationQueueFrom(ptr unsafe.Pointer) OperationQueue {
	return OperationQueue{
		ID: objc.ID(ptr),
	}
}


// Cancels all queued and executing operations. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OperationQueue/cancelAllOperations()
func (o_ OperationQueue) CancelAllOperations() {
	objc.Send[objc.ID](o_.ID, objc.Sel("cancelAllOperations"))
}


