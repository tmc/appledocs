// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ModelStructureProgramValueType] class.
var (
	ModelStructureProgramValueTypeClass     _ModelStructureProgramValueTypeClass
	ModelStructureProgramValueTypeClassOnce sync.Once
)

func getModelStructureProgramValueTypeClass() _ModelStructureProgramValueTypeClass {
	ModelStructureProgramValueTypeClassOnce.Do(func() {
		ModelStructureProgramValueTypeClass = _ModelStructureProgramValueTypeClass{objc.GetClass("MLModelStructureProgramValueType")}
	})
	return ModelStructureProgramValueTypeClass
}

type _ModelStructureProgramValueTypeClass struct {
	class objc.Class
}

// An interface definition for the [ModelStructureProgramValueType] class.
type IModelStructureProgramValueType interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A class representing the type of a value or a variable in the Program.


// A class representing the type of a value or a variable in the Program.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramValueType
type ModelStructureProgramValueType struct {
	objectivec.Object
}

// ModelStructureProgramValueTypeFrom constructs a [ModelStructureProgramValueType] from an unsafe.Pointer.
//
// A class representing the type of a value or a variable in the Program.
func ModelStructureProgramValueTypeFrom(ptr unsafe.Pointer) ModelStructureProgramValueType {
	return ModelStructureProgramValueType{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _ModelStructureProgramValueTypeClass) Alloc() ModelStructureProgramValueType {
	rv := objc.Send[ModelStructureProgramValueType](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _ModelStructureProgramValueTypeClass) New() ModelStructureProgramValueType {
	rv := objc.Send[ModelStructureProgramValueType](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ ModelStructureProgramValueType) Init() ModelStructureProgramValueType {
	rv := objc.Send[ModelStructureProgramValueType](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ ModelStructureProgramValueType) Autorelease() ModelStructureProgramValueType {
	rv := objc.Send[ModelStructureProgramValueType](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewModelStructureProgramValueType creates a new ModelStructureProgramValueType instance.
func NewModelStructureProgramValueType() ModelStructureProgramValueType {
	return getModelStructureProgramValueTypeClass().New()
}




