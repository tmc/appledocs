// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ModelStructureProgramOperation] class.
var (
	ModelStructureProgramOperationClass     _ModelStructureProgramOperationClass
	ModelStructureProgramOperationClassOnce sync.Once
)

func getModelStructureProgramOperationClass() _ModelStructureProgramOperationClass {
	ModelStructureProgramOperationClassOnce.Do(func() {
		ModelStructureProgramOperationClass = _ModelStructureProgramOperationClass{objc.GetClass("MLModelStructureProgramOperation")}
	})
	return ModelStructureProgramOperationClass
}

type _ModelStructureProgramOperationClass struct {
	class objc.Class
}





// An interface definition for the [ModelStructureProgramOperation] class.
type IModelStructureProgramOperation interface {
	objectivec.IObject
	

	// properties:
	Blocks() []ModelStructureProgramBlock
	Inputs() foundation.IDictionary
	OperatorName() foundation.foundation.INSString
	Outputs() []ModelStructureProgramNamedValueType


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _ModelStructureProgramOperationClass) Alloc() ModelStructureProgramOperation {
	rv := objc.Send[ModelStructureProgramOperation](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _ModelStructureProgramOperationClass) New() ModelStructureProgramOperation {
	rv := objc.Send[ModelStructureProgramOperation](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ ModelStructureProgramOperation) Init() ModelStructureProgramOperation {
	rv := objc.Send[ModelStructureProgramOperation](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ ModelStructureProgramOperation) Autorelease() ModelStructureProgramOperation {
	rv := objc.Send[ModelStructureProgramOperation](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewModelStructureProgramOperation creates a new ModelStructureProgramOperation instance.
func NewModelStructureProgramOperation() ModelStructureProgramOperation {
	return getModelStructureProgramOperationClass().New()
}





// A class representing an Operation in a Program.


// A class representing an Operation in a Program.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramOperation
type ModelStructureProgramOperation struct {
	objectivec.Object
}

// ModelStructureProgramOperationFrom constructs a [ModelStructureProgramOperation] from an unsafe.Pointer.
//
// A class representing an Operation in a Program.
func ModelStructureProgramOperationFrom(ptr unsafe.Pointer) ModelStructureProgramOperation {
	return ModelStructureProgramOperation{objectivec.Object{objc.ID(ptr)}}
}

























// Nested blocks for loops and conditionals, e.g., a conditional block will have two entries here.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramOperation/blocks
func (m_ ModelStructureProgramOperation) Blocks() []ModelStructureProgramBlock {
	rv := objc.Send[[]ModelStructureProgramBlock](m_.ID, objc.Sel("blocks"))
	return rv
}


// The arguments to the Operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramOperation/inputs
func (m_ ModelStructureProgramOperation) Inputs() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("inputs"))
	return rv
}


// The name of the operator, e.g., “conv”, “pool”, “softmax”, etc.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramOperation/operatorName
func (m_ ModelStructureProgramOperation) OperatorName() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("operatorName"))
	return rv
}


// The outputs of the Operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramOperation/outputs
func (m_ ModelStructureProgramOperation) Outputs() []ModelStructureProgramNamedValueType {
	rv := objc.Send[[]ModelStructureProgramNamedValueType](m_.ID, objc.Sel("outputs"))
	return rv
}








