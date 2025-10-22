// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ModelStructureProgram] class.
var (
	ModelStructureProgramClass     _ModelStructureProgramClass
	ModelStructureProgramClassOnce sync.Once
)

func getModelStructureProgramClass() _ModelStructureProgramClass {
	ModelStructureProgramClassOnce.Do(func() {
		ModelStructureProgramClass = _ModelStructureProgramClass{objc.GetClass("MLModelStructureProgram")}
	})
	return ModelStructureProgramClass
}

type _ModelStructureProgramClass struct {
	class objc.Class
}

// An interface definition for the [ModelStructureProgram] class.
type IModelStructureProgram interface {
	objectivec.IObject
	Functions() unsafe.Pointer
}

// A class representing the structure of an ML Program model.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureProgram
type ModelStructureProgram struct {
	objectivec.Object
}

// ModelStructureProgramFrom constructs a [ModelStructureProgram] from an unsafe.Pointer.
//
// A class representing the structure of an ML Program model.
func ModelStructureProgramFrom(ptr unsafe.Pointer) ModelStructureProgram {
	return ModelStructureProgram{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _ModelStructureProgramClass) Alloc() ModelStructureProgram {
	rv := objc.Send[ModelStructureProgram](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _ModelStructureProgramClass) New() ModelStructureProgram {
	rv := objc.Send[ModelStructureProgram](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ ModelStructureProgram) Init() ModelStructureProgram {
	rv := objc.Send[ModelStructureProgram](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ ModelStructureProgram) Autorelease() ModelStructureProgram {
	rv := objc.Send[ModelStructureProgram](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewModelStructureProgram creates a new ModelStructureProgram instance.
func NewModelStructureProgram() ModelStructureProgram {
	return getModelStructureProgramClass().New()
}


// The functions in the program.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureProgram/functions
func (m_ ModelStructureProgram) Functions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("functions"))
	return rv
}



