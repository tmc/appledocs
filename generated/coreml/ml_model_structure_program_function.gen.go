// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLModelStructureProgramFunction */


/* debug [class_header]: Header for MLModelStructureProgramFunction */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ModelStructureProgramFunction */
// An interface definition for the [ModelStructureProgramFunction] class.
type IModelStructureProgramFunction interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ModelStructureProgramFunction */
	// properties:
	Block() IMLModelStructureProgramBlock
	Inputs() []ModelStructureProgramNamedValueType
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ModelStructureProgramFunction */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ModelStructureProgramFunction */
// Alloc allocates a new instance without initialization.
func (mc _ModelStructureProgramFunctionClass) Alloc() ModelStructureProgramFunction {
	rv := objc.Send[ModelStructureProgramFunction](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ModelStructureProgramFunction */
// A class representing a function in the Program.


// A class representing a function in the Program.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ModelStructureProgramFunction *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ModelStructureProgramFunction */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ModelStructureProgramFunction */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ModelStructureProgramFunction */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ModelStructureProgramFunction */

// The active block in the function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramFunction/block
func (m_ ModelStructureProgramFunction) Block() IMLModelStructureProgramBlock {
	rv := objc.Send[ModelStructureProgramBlock](m_.ID, objc.Sel("block"))
	return rv
}/* debug [instance_properties/getter]: block */


// The named inputs to the function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramFunction/inputs
func (m_ ModelStructureProgramFunction) Inputs() []ModelStructureProgramNamedValueType {
	rv := objc.Send[[]ModelStructureProgramNamedValueType](m_.ID, objc.Sel("inputs"))
	return rv
}/* debug [instance_properties/getter]: inputs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLModelStructureProgramFunction */



