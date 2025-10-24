// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSBlockOperation */


/* debug [class_header]: Header for NSBlockOperation */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BlockOperation */
// An interface definition for the [BlockOperation] class.
type IBlockOperation interface {
	IOperation
	
/* debug [class_interface_properties]: Properties for BlockOperation */
	// properties:
	ExecutionBlocks() []func()
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BlockOperation */
	// methods:
	AddExecutionBlock(block unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BlockOperation */
// Alloc allocates a new instance without initialization.
func (bc _BlockOperationClass) Alloc() BlockOperation {
	rv := objc.Send[BlockOperation](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BlockOperation */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BlockOperation */

// Creates and returns an object and adds the specified block to it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/BlockOperation/init(block:)
func NewBlockOperationWithBlock(block unsafe.Pointer) BlockOperation {
	rv := objc.Send[BlockOperation](objc.ID(getBlockOperationClass().class), objc.Sel("blockOperationWithBlock:"), block)
	return rv
}/* debug [class_init_methods/constructor]: NewBlockOperationWithBlock */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BlockOperation */

// Creates and returns an object and adds the specified block to it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/BlockOperation/init(block:)
func (bc _BlockOperationClass) BlockOperationWithBlock(block unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(bc.class), objc.Sel("blockOperationWithBlock:"), block)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=BlockOperationWithBlock) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BlockOperation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BlockOperation */

// Adds the specified block to the receiver’s list of blocks to perform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/BlockOperation/addExecutionBlock(_:)
func (b_ BlockOperation) AddExecutionBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("addExecutionBlock:"), block)
}/* debug [instance_methods/method]: AddExecutionBlock */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BlockOperation */

// The blocks associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/BlockOperation/executionBlocks
func (b_ BlockOperation) ExecutionBlocks() []func() {
	rv := objc.Send[[]func()](b_.ID, objc.Sel("executionBlocks"))
	return rv
}/* debug [instance_properties/getter]: executionBlocks */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSBlockOperation */


