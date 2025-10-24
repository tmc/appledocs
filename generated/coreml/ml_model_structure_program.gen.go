// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLModelStructureProgram */


/* debug [class_header]: Header for MLModelStructureProgram */
// The class instance for the [ModelStructureProgram] class.
var (
	ModelStructureProgramClass     _ModelStructureProgramClass
	ModelStructureProgramClassOnce sync.Once
)

func getModelStructureProgramClass() _ModelStructureProgramClass {
	ModelStructureProgramClassOnce.Do(func() {
		ModelStructureProgramClass = _ModelStructureProgramClass{objc.GetClass("MLModelStructureProgram")}
	})
	return ModelStructureProgramClass
}

type _ModelStructureProgramClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ModelStructureProgram */
// An interface definition for the [ModelStructureProgram] class.
type IModelStructureProgram interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ModelStructureProgram */
	// properties:
	Functions() foundation.IDictionary
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ModelStructureProgram */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ModelStructureProgram */
// Alloc allocates a new instance without initialization.
func (mc _ModelStructureProgramClass) Alloc() ModelStructureProgram {
	rv := objc.Send[ModelStructureProgram](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _ModelStructureProgramClass) New() ModelStructureProgram {
	rv := objc.Send[ModelStructureProgram](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ ModelStructureProgram) Init() ModelStructureProgram {
	rv := objc.Send[ModelStructureProgram](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ ModelStructureProgram) Autorelease() ModelStructureProgram {
	rv := objc.Send[ModelStructureProgram](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewModelStructureProgram creates a new ModelStructureProgram instance.
func NewModelStructureProgram() ModelStructureProgram {
	return getModelStructureProgramClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ModelStructureProgram */
// A class representing the structure of an ML Program model.


// A class representing the structure of an ML Program model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureProgram
type ModelStructureProgram struct {
	objectivec.Object
}

// ModelStructureProgramFrom constructs a [ModelStructureProgram] from an unsafe.Pointer.
//
// A class representing the structure of an ML Program model.
func ModelStructureProgramFrom(ptr unsafe.Pointer) ModelStructureProgram {
	return ModelStructureProgram{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ModelStructureProgram *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ModelStructureProgram */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ModelStructureProgram */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ModelStructureProgram */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ModelStructureProgram */

// The functions in the program.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureProgram/functions
func (m_ ModelStructureProgram) Functions() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("functions"))
	return rv
}/* debug [instance_properties/getter]: functions */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLModelStructureProgram */



