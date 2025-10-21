// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [ModelStructureProgramValue] class.
var (
	ModelStructureProgramValueClass     _ModelStructureProgramValueClass
	ModelStructureProgramValueClassOnce sync.Once
)

func getModelStructureProgramValueClass() _ModelStructureProgramValueClass {
	ModelStructureProgramValueClassOnce.Do(func() {
		ModelStructureProgramValueClass = _ModelStructureProgramValueClass{objc.GetClass("MLModelStructureProgramValue")}
	})
	return ModelStructureProgramValueClass
}

type _ModelStructureProgramValueClass struct {
	class objc.Class
}

// An interface definition for the [ModelStructureProgramValue] class.
type IModelStructureProgramValue interface {
	objectivec.IObject
}

// A class representing a constant value in the Program.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramValue
type ModelStructureProgramValue struct {
	objectivec.Object
}

// ModelStructureProgramValueFrom constructs a [ModelStructureProgramValue] from an unsafe.Pointer.
//
// A class representing a constant value in the Program.
func ModelStructureProgramValueFrom(ptr unsafe.Pointer) ModelStructureProgramValue {
	return ModelStructureProgramValue{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _ModelStructureProgramValueClass) Alloc() ModelStructureProgramValue {
	rv := objc.Send[ModelStructureProgramValue](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _ModelStructureProgramValueClass) New() ModelStructureProgramValue {
	rv := objc.Send[ModelStructureProgramValue](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ ModelStructureProgramValue) Init() ModelStructureProgramValue {
	rv := objc.Send[ModelStructureProgramValue](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ ModelStructureProgramValue) Autorelease() ModelStructureProgramValue {
	rv := objc.Send[ModelStructureProgramValue](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewModelStructureProgramValue creates a new ModelStructureProgramValue instance.
func NewModelStructureProgramValue() ModelStructureProgramValue {
	return getModelStructureProgramValueClass().New()
}




