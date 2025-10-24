// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLModelStructureProgramArgument */


/* debug [class_header]: Header for MLModelStructureProgramArgument */
// The class instance for the [ModelStructureProgramArgument] class.
var (
	ModelStructureProgramArgumentClass     _ModelStructureProgramArgumentClass
	ModelStructureProgramArgumentClassOnce sync.Once
)

func getModelStructureProgramArgumentClass() _ModelStructureProgramArgumentClass {
	ModelStructureProgramArgumentClassOnce.Do(func() {
		ModelStructureProgramArgumentClass = _ModelStructureProgramArgumentClass{objc.GetClass("MLModelStructureProgramArgument")}
	})
	return ModelStructureProgramArgumentClass
}

type _ModelStructureProgramArgumentClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ModelStructureProgramArgument */
// An interface definition for the [ModelStructureProgramArgument] class.
type IModelStructureProgramArgument interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ModelStructureProgramArgument */
	// properties:
	Bindings() []ModelStructureProgramBinding
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ModelStructureProgramArgument */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ModelStructureProgramArgument */
// Alloc allocates a new instance without initialization.
func (mc _ModelStructureProgramArgumentClass) Alloc() ModelStructureProgramArgument {
	rv := objc.Send[ModelStructureProgramArgument](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _ModelStructureProgramArgumentClass) New() ModelStructureProgramArgument {
	rv := objc.Send[ModelStructureProgramArgument](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ ModelStructureProgramArgument) Init() ModelStructureProgramArgument {
	rv := objc.Send[ModelStructureProgramArgument](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ ModelStructureProgramArgument) Autorelease() ModelStructureProgramArgument {
	rv := objc.Send[ModelStructureProgramArgument](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewModelStructureProgramArgument creates a new ModelStructureProgramArgument instance.
func NewModelStructureProgramArgument() ModelStructureProgramArgument {
	return getModelStructureProgramArgumentClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ModelStructureProgramArgument */
// A class representing an argument in the Program.


// A class representing an argument in the Program.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramArgument
type ModelStructureProgramArgument struct {
	objectivec.Object
}

// ModelStructureProgramArgumentFrom constructs a [ModelStructureProgramArgument] from an unsafe.Pointer.
//
// A class representing an argument in the Program.
func ModelStructureProgramArgumentFrom(ptr unsafe.Pointer) ModelStructureProgramArgument {
	return ModelStructureProgramArgument{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ModelStructureProgramArgument *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ModelStructureProgramArgument */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ModelStructureProgramArgument */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ModelStructureProgramArgument */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ModelStructureProgramArgument */

// The array of bindings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramArgument/bindings
func (m_ ModelStructureProgramArgument) Bindings() []ModelStructureProgramBinding {
	rv := objc.Send[[]ModelStructureProgramBinding](m_.ID, objc.Sel("bindings"))
	return rv
}/* debug [instance_properties/getter]: bindings */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLModelStructureProgramArgument */



