// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [BlockOperation] class.
var blockOperationClass = _BlockOperationClass{objc.GetClass("NSBlockOperation")}

type _BlockOperationClass struct {
	class objc.Class
}

// An interface definition for the [BlockOperation] class.
type IBlockOperation interface {
	IOperation
}

// An operation that manages the concurrent execution of one or more blocks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/BlockOperation

type BlockOperation struct {
	Operation
}

// BlockOperationFrom constructs a [BlockOperation] from an unsafe.Pointer.
//
// An operation that manages the concurrent execution of one or more blocks.
func BlockOperationFrom(ptr unsafe.Pointer) BlockOperation {
	return BlockOperation{
		Operation: OperationFrom(ptr),
	}
}



