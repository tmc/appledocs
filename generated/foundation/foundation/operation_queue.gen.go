// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [OperationQueue] class.
var OperationQueueClass objc.Class

func init() {
	OperationQueueClass = objc.GetClass("NSOperationQueue")
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
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/OperationQueue/cancelAllOperations()
func (o_ OperationQueue) CancelAllOperations() {
	sel := objc.RegisterName("cancelAllOperations")
	o_.ID.Send(sel)
}


