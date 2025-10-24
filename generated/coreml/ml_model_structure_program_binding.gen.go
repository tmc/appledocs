// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ModelStructureProgramBinding] class.
var (
	ModelStructureProgramBindingClass     _ModelStructureProgramBindingClass
	ModelStructureProgramBindingClassOnce sync.Once
)

func getModelStructureProgramBindingClass() _ModelStructureProgramBindingClass {
	ModelStructureProgramBindingClassOnce.Do(func() {
		ModelStructureProgramBindingClass = _ModelStructureProgramBindingClass{objc.GetClass("MLModelStructureProgramBinding")}
	})
	return ModelStructureProgramBindingClass
}

type _ModelStructureProgramBindingClass struct {
	class objc.Class
}

// An interface definition for the [ModelStructureProgramBinding] class.
type IModelStructureProgramBinding interface {
	objectivec.IObject
	// properties:
	Name() objc.IObject /* cross-framework: NSString */
	Value() IMLModelStructureProgramValue
	// methods:
}

// A class representing a binding in the Program
//
// A Binding is either a previously defined name of a variable or a constant value in the Program.


// A class representing a binding in the Program
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramBinding
type ModelStructureProgramBinding struct {
	objectivec.Object
}

// ModelStructureProgramBindingFrom constructs a [ModelStructureProgramBinding] from an unsafe.Pointer.
//
// A class representing a binding in the Program
func ModelStructureProgramBindingFrom(ptr unsafe.Pointer) ModelStructureProgramBinding {
	return ModelStructureProgramBinding{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _ModelStructureProgramBindingClass) Alloc() ModelStructureProgramBinding {
	rv := objc.Send[ModelStructureProgramBinding](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _ModelStructureProgramBindingClass) New() ModelStructureProgramBinding {
	rv := objc.Send[ModelStructureProgramBinding](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ ModelStructureProgramBinding) Init() ModelStructureProgramBinding {
	rv := objc.Send[ModelStructureProgramBinding](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ ModelStructureProgramBinding) Autorelease() ModelStructureProgramBinding {
	rv := objc.Send[ModelStructureProgramBinding](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewModelStructureProgramBinding creates a new ModelStructureProgramBinding instance.
func NewModelStructureProgramBinding() ModelStructureProgramBinding {
	return getModelStructureProgramBindingClass().New()
}



// The name of the variable in the Program.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramBinding/name
func (m_ ModelStructureProgramBinding) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}


// The compile time constant value in the Program.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramBinding/value
func (m_ ModelStructureProgramBinding) Value() IMLModelStructureProgramValue {
	rv := objc.Send[ModelStructureProgramValue](m_.ID, objc.Sel("value"))
	return rv
}



