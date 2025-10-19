// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [BlockOperation] class.
var (
	blockOperationClass     _BlockOperationClass
	blockOperationClassOnce sync.Once
)

func getBlockOperationClass() _BlockOperationClass {
	blockOperationClassOnce.Do(func() {
		blockOperationClass = _BlockOperationClass{objc.GetClass("NSBlockOperation")}
	})
	return blockOperationClass
}

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

// Alloc allocates a new instance without initialization.
func (bc _BlockOperationClass) Alloc() BlockOperation {
	rv := objc.Send[BlockOperation](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BlockOperationClass) New() BlockOperation {
	rv := objc.Send[BlockOperation](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BlockOperation) Init() BlockOperation {
	rv := objc.Send[BlockOperation](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BlockOperation) Autorelease() BlockOperation {
	rv := objc.Send[BlockOperation](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBlockOperation creates a new BlockOperation instance.
func NewBlockOperation() BlockOperation {
	return getBlockOperationClass().New()
}




