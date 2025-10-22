// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [BlockOperation] class.
var (
	BlockOperationClass     _BlockOperationClass
	BlockOperationClassOnce sync.Once
)

func getBlockOperationClass() _BlockOperationClass {
	BlockOperationClassOnce.Do(func() {
		BlockOperationClass = _BlockOperationClass{objc.GetClass("NSBlockOperation")}
	})
	return BlockOperationClass
}

type _BlockOperationClass struct {
	class objc.Class
}

// An interface definition for the [BlockOperation] class.
type IBlockOperation interface {
	IOperation
	AddExecutionBlock(block unsafe.Pointer)
	ExecutionBlocks() []func()
}

// An operation that manages the concurrent execution of one or more blocks.
//
// The class is a concrete subclass of that manages the concurrent execution of one or more blocks. You can use this object to execute several blocks at once without having to create separate operation objects for each. When executing more than one block, the operation itself is considered finished only when all blocks have finished executing. Blocks added to a block operation are dispatched with default priority to an appropriate work queue. The blocks themselves should not make any assumptions about the configuration of their execution environment. For more information about blocks, see .


// An operation that manages the concurrent execution of one or more blocks.
//
// [Full Topic]
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





// Creates and returns an object and adds the specified block to it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/BlockOperation/init(block:)

func NewBlockOperationWithBlock(block unsafe.Pointer) BlockOperation {
	rv := objc.Send[BlockOperation](objc.ID(getBlockOperationClass().class), objc.Sel("blockOperationWithBlock:"), block)
	return rv
}



// Creates and returns an object and adds the specified block to it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/BlockOperation/init(block:)

func (bc _BlockOperationClass) BlockOperationWithBlock(block unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("blockOperationWithBlock:"), block)
	return rv
}


// Adds the specified block to the receiver’s list of blocks to perform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/BlockOperation/addExecutionBlock(_:)

func (b_ BlockOperation) AddExecutionBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("addExecutionBlock:"), block)
}


// The blocks associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/BlockOperation/executionBlocks

func (b_ BlockOperation) ExecutionBlocks() []func() {
	rv := objc.Send[[]func()](b_.ID, objc.Sel("executionBlocks"))
	return rv
}


