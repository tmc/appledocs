// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLModelStructureProgramValueType */


/* debug [class_header]: Header for MLModelStructureProgramValueType */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ModelStructureProgramValueType */
// An interface definition for the [ModelStructureProgramValueType] class.
type IModelStructureProgramValueType interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ModelStructureProgramValueType */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ModelStructureProgramValueType */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ModelStructureProgramValueType */
// Alloc allocates a new instance without initialization.
func (mc _ModelStructureProgramValueTypeClass) Alloc() ModelStructureProgramValueType {
	rv := objc.Send[ModelStructureProgramValueType](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ModelStructureProgramValueType */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ModelStructureProgramValueType *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ModelStructureProgramValueType */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ModelStructureProgramValueType */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ModelStructureProgramValueType */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ModelStructureProgramValueType */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLModelStructureProgramValueType */



