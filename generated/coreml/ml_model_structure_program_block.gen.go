// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLModelStructureProgramBlock */


/* debug [class_header]: Header for MLModelStructureProgramBlock */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ModelStructureProgramBlock */
// An interface definition for the [ModelStructureProgramBlock] class.
type IModelStructureProgramBlock interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ModelStructureProgramBlock */
	// properties:
	Inputs() []ModelStructureProgramNamedValueType
	Operations() []ModelStructureProgramOperation
	OutputNames() []string
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ModelStructureProgramBlock */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ModelStructureProgramBlock */
// Alloc allocates a new instance without initialization.
func (mc _ModelStructureProgramBlockClass) Alloc() ModelStructureProgramBlock {
	rv := objc.Send[ModelStructureProgramBlock](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ModelStructureProgramBlock */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ModelStructureProgramBlock *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ModelStructureProgramBlock */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ModelStructureProgramBlock */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ModelStructureProgramBlock */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ModelStructureProgramBlock */

// The named inputs to the block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramBlock/inputs
func (m_ ModelStructureProgramBlock) Inputs() []ModelStructureProgramNamedValueType {
	rv := objc.Send[[]ModelStructureProgramNamedValueType](m_.ID, objc.Sel("inputs"))
	return rv
}/* debug [instance_properties/getter]: inputs */


// The list of topologically sorted operations in the block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramBlock/operations
func (m_ ModelStructureProgramBlock) Operations() []ModelStructureProgramOperation {
	rv := objc.Send[[]ModelStructureProgramOperation](m_.ID, objc.Sel("operations"))
	return rv
}/* debug [instance_properties/getter]: operations */


// The output names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramBlock/outputNames
func (m_ ModelStructureProgramBlock) OutputNames() []string {
	rv := objc.Send[[]string](m_.ID, objc.Sel("outputNames"))
	return rv
}/* debug [instance_properties/getter]: outputNames */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLModelStructureProgramBlock */



