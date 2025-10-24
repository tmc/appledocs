// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ModelStructureProgramBlock] class.
var (
	ModelStructureProgramBlockClass     _ModelStructureProgramBlockClass
	ModelStructureProgramBlockClassOnce sync.Once
)

func getModelStructureProgramBlockClass() _ModelStructureProgramBlockClass {
	ModelStructureProgramBlockClassOnce.Do(func() {
		ModelStructureProgramBlockClass = _ModelStructureProgramBlockClass{objc.GetClass("MLModelStructureProgramBlock")}
	})
	return ModelStructureProgramBlockClass
}

type _ModelStructureProgramBlockClass struct {
	class objc.Class
}

// An interface definition for the [ModelStructureProgramBlock] class.
type IModelStructureProgramBlock interface {
	objectivec.IObject
	// properties:
	Inputs() []IModelStructureProgramNamedValueType
	Operations() []IModelStructureProgramOperation
	OutputNames() []string
	// methods:
}

// A class representing a block in the Program.


// A class representing a block in the Program.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramBlock
type ModelStructureProgramBlock struct {
	objectivec.Object
}

// ModelStructureProgramBlockFrom constructs a [ModelStructureProgramBlock] from an unsafe.Pointer.
//
// A class representing a block in the Program.
func ModelStructureProgramBlockFrom(ptr unsafe.Pointer) ModelStructureProgramBlock {
	return ModelStructureProgramBlock{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _ModelStructureProgramBlockClass) Alloc() ModelStructureProgramBlock {
	rv := objc.Send[ModelStructureProgramBlock](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _ModelStructureProgramBlockClass) New() ModelStructureProgramBlock {
	rv := objc.Send[ModelStructureProgramBlock](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ ModelStructureProgramBlock) Init() ModelStructureProgramBlock {
	rv := objc.Send[ModelStructureProgramBlock](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ ModelStructureProgramBlock) Autorelease() ModelStructureProgramBlock {
	rv := objc.Send[ModelStructureProgramBlock](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewModelStructureProgramBlock creates a new ModelStructureProgramBlock instance.
func NewModelStructureProgramBlock() ModelStructureProgramBlock {
	return getModelStructureProgramBlockClass().New()
}



// The named inputs to the block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramBlock/inputs
func (m_ ModelStructureProgramBlock) Inputs() []IModelStructureProgramNamedValueType {
	rv := objc.Send[[]ModelStructureProgramNamedValueType](m_.ID, objc.Sel("inputs"))
	return rv
}


// The list of topologically sorted operations in the block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramBlock/operations
func (m_ ModelStructureProgramBlock) Operations() []IModelStructureProgramOperation {
	rv := objc.Send[[]ModelStructureProgramOperation](m_.ID, objc.Sel("operations"))
	return rv
}


// The output names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramBlock/outputNames
func (m_ ModelStructureProgramBlock) OutputNames() []string {
	rv := objc.Send[[]string](m_.ID, objc.Sel("outputNames"))
	return rv
}



