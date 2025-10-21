// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ModelStructureProgramFunction] class.
var (
	ModelStructureProgramFunctionClass     _ModelStructureProgramFunctionClass
	ModelStructureProgramFunctionClassOnce sync.Once
)

func getModelStructureProgramFunctionClass() _ModelStructureProgramFunctionClass {
	ModelStructureProgramFunctionClassOnce.Do(func() {
		ModelStructureProgramFunctionClass = _ModelStructureProgramFunctionClass{objc.GetClass("MLModelStructureProgramFunction")}
	})
	return ModelStructureProgramFunctionClass
}

type _ModelStructureProgramFunctionClass struct {
	class objc.Class
}

// An interface definition for the [ModelStructureProgramFunction] class.
type IModelStructureProgramFunction interface {
	objectivec.IObject
}

// A class representing a function in the Program.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramFunction
type ModelStructureProgramFunction struct {
	objectivec.Object
}

// ModelStructureProgramFunctionFrom constructs a [ModelStructureProgramFunction] from an unsafe.Pointer.
//
// A class representing a function in the Program.
func ModelStructureProgramFunctionFrom(ptr unsafe.Pointer) ModelStructureProgramFunction {
	return ModelStructureProgramFunction{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _ModelStructureProgramFunctionClass) Alloc() ModelStructureProgramFunction {
	rv := objc.Send[ModelStructureProgramFunction](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _ModelStructureProgramFunctionClass) New() ModelStructureProgramFunction {
	rv := objc.Send[ModelStructureProgramFunction](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ ModelStructureProgramFunction) Init() ModelStructureProgramFunction {
	rv := objc.Send[ModelStructureProgramFunction](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ ModelStructureProgramFunction) Autorelease() ModelStructureProgramFunction {
	rv := objc.Send[ModelStructureProgramFunction](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewModelStructureProgramFunction creates a new ModelStructureProgramFunction instance.
func NewModelStructureProgramFunction() ModelStructureProgramFunction {
	return getModelStructureProgramFunctionClass().New()
}


// The active block in the function.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramFunction/block
func (m_ ModelStructureProgramFunction) Block() MLModelStructureProgramBlock {
	rv := objc.Send[MLModelStructureProgramBlock](m_.ID, objc.Sel("block"))
	return rv
}

// The named inputs to the function.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramFunction/inputs
func (m_ ModelStructureProgramFunction) Inputs() []ModelStructureProgramNamedValueType {
	rv := objc.Send[[]ModelStructureProgramNamedValueType](m_.ID, objc.Sel("inputs"))
	return rv
}



