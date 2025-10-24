// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLModelStructureProgramValue */


/* debug [class_header]: Header for MLModelStructureProgramValue */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ModelStructureProgramValue */
// An interface definition for the [ModelStructureProgramValue] class.
type IModelStructureProgramValue interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ModelStructureProgramValue */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ModelStructureProgramValue */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ModelStructureProgramValue */
// Alloc allocates a new instance without initialization.
func (mc _ModelStructureProgramValueClass) Alloc() ModelStructureProgramValue {
	rv := objc.Send[ModelStructureProgramValue](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ModelStructureProgramValue */
// A class representing a constant value in the Program.


// A class representing a constant value in the Program.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ModelStructureProgramValue *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ModelStructureProgramValue */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ModelStructureProgramValue */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ModelStructureProgramValue */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ModelStructureProgramValue */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLModelStructureProgramValue */



