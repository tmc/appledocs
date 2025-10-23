// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ModelStructureProgramNamedValueType] class.
var (
	ModelStructureProgramNamedValueTypeClass     _ModelStructureProgramNamedValueTypeClass
	ModelStructureProgramNamedValueTypeClassOnce sync.Once
)

func getModelStructureProgramNamedValueTypeClass() _ModelStructureProgramNamedValueTypeClass {
	ModelStructureProgramNamedValueTypeClassOnce.Do(func() {
		ModelStructureProgramNamedValueTypeClass = _ModelStructureProgramNamedValueTypeClass{objc.GetClass("MLModelStructureProgramNamedValueType")}
	})
	return ModelStructureProgramNamedValueTypeClass
}

type _ModelStructureProgramNamedValueTypeClass struct {
	class objc.Class
}

// An interface definition for the [ModelStructureProgramNamedValueType] class.
type IModelStructureProgramNamedValueType interface {
	objectivec.IObject
	// properties:
	Name() string /* primitive/slice/pointer. */
	Type() IMLModelStructureProgramValueType
	// methods:
}

// A class representing a named value type in a Program.


// A class representing a named value type in a Program.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramNamedValueType
type ModelStructureProgramNamedValueType struct {
	objectivec.Object
}

// ModelStructureProgramNamedValueTypeFrom constructs a [ModelStructureProgramNamedValueType] from an unsafe.Pointer.
//
// A class representing a named value type in a Program.
func ModelStructureProgramNamedValueTypeFrom(ptr unsafe.Pointer) ModelStructureProgramNamedValueType {
	return ModelStructureProgramNamedValueType{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _ModelStructureProgramNamedValueTypeClass) Alloc() ModelStructureProgramNamedValueType {
	rv := objc.Send[ModelStructureProgramNamedValueType](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _ModelStructureProgramNamedValueTypeClass) New() ModelStructureProgramNamedValueType {
	rv := objc.Send[ModelStructureProgramNamedValueType](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ ModelStructureProgramNamedValueType) Init() ModelStructureProgramNamedValueType {
	rv := objc.Send[ModelStructureProgramNamedValueType](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ ModelStructureProgramNamedValueType) Autorelease() ModelStructureProgramNamedValueType {
	rv := objc.Send[ModelStructureProgramNamedValueType](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewModelStructureProgramNamedValueType creates a new ModelStructureProgramNamedValueType instance.
func NewModelStructureProgramNamedValueType() ModelStructureProgramNamedValueType {
	return getModelStructureProgramNamedValueTypeClass().New()
}



// The name of the parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramNamedValueType/name
func (m_ ModelStructureProgramNamedValueType) Name() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](m_.ID, objc.Sel("name"))
	return rv
}


// The type of the parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramNamedValueType/type
func (m_ ModelStructureProgramNamedValueType) Type() IMLModelStructureProgramValueType {
	rv := objc.Send[ModelStructureProgramValueType](m_.ID, objc.Sel("type"))
	return rv
}



