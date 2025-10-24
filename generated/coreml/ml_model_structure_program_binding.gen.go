// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLModelStructureProgramBinding */


/* debug [class_header]: Header for MLModelStructureProgramBinding */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ModelStructureProgramBinding */
// An interface definition for the [ModelStructureProgramBinding] class.
type IModelStructureProgramBinding interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ModelStructureProgramBinding */
	// properties:
	Name() objc.IObject /* cross-framework: NSString */
	Value() IMLModelStructureProgramValue
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ModelStructureProgramBinding */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ModelStructureProgramBinding */
// Alloc allocates a new instance without initialization.
func (mc _ModelStructureProgramBindingClass) Alloc() ModelStructureProgramBinding {
	rv := objc.Send[ModelStructureProgramBinding](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ModelStructureProgramBinding */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ModelStructureProgramBinding *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ModelStructureProgramBinding */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ModelStructureProgramBinding */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ModelStructureProgramBinding */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ModelStructureProgramBinding */

// The name of the variable in the Program.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramBinding/name
func (m_ ModelStructureProgramBinding) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The compile time constant value in the Program.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramBinding/value
func (m_ ModelStructureProgramBinding) Value() IMLModelStructureProgramValue {
	rv := objc.Send[ModelStructureProgramValue](m_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLModelStructureProgramBinding */



