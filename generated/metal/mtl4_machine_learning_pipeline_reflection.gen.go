// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTL4MachineLearningPipelineReflection */


/* debug [class_header]: Header for MTL4MachineLearningPipelineReflection */
// The class instance for the [MTL4MachineLearningPipelineReflection] class.
var (
	MTL4MachineLearningPipelineReflectionClass     _MTL4MachineLearningPipelineReflectionClass
	MTL4MachineLearningPipelineReflectionClassOnce sync.Once
)

func getMTL4MachineLearningPipelineReflectionClass() _MTL4MachineLearningPipelineReflectionClass {
	MTL4MachineLearningPipelineReflectionClassOnce.Do(func() {
		MTL4MachineLearningPipelineReflectionClass = _MTL4MachineLearningPipelineReflectionClass{objc.GetClass("MTL4MachineLearningPipelineReflection")}
	})
	return MTL4MachineLearningPipelineReflectionClass
}

type _MTL4MachineLearningPipelineReflectionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTL4MachineLearningPipelineReflection */
// An interface definition for the [MTL4MachineLearningPipelineReflection] class.
type IMTL4MachineLearningPipelineReflection interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTL4MachineLearningPipelineReflection */
	// properties:
	Bindings() []objc.ID
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTL4MachineLearningPipelineReflection */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTL4MachineLearningPipelineReflection */
// Alloc allocates a new instance without initialization.
func (mc _MTL4MachineLearningPipelineReflectionClass) Alloc() MTL4MachineLearningPipelineReflection {
	rv := objc.Send[MTL4MachineLearningPipelineReflection](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTL4MachineLearningPipelineReflectionClass) New() MTL4MachineLearningPipelineReflection {
	rv := objc.Send[MTL4MachineLearningPipelineReflection](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4MachineLearningPipelineReflection) Init() MTL4MachineLearningPipelineReflection {
	rv := objc.Send[MTL4MachineLearningPipelineReflection](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4MachineLearningPipelineReflection) Autorelease() MTL4MachineLearningPipelineReflection {
	rv := objc.Send[MTL4MachineLearningPipelineReflection](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4MachineLearningPipelineReflection creates a new MTL4MachineLearningPipelineReflection instance.
func NewMTL4MachineLearningPipelineReflection() MTL4MachineLearningPipelineReflection {
	return getMTL4MachineLearningPipelineReflectionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTL4MachineLearningPipelineReflection */
// Represents reflection information for a machine learning pipeline state.


// Represents reflection information for a machine learning pipeline state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MachineLearningPipelineReflection
type MTL4MachineLearningPipelineReflection struct {
	objectivec.Object
}

// MTL4MachineLearningPipelineReflectionFrom constructs a [MTL4MachineLearningPipelineReflection] from an unsafe.Pointer.
//
// Represents reflection information for a machine learning pipeline state.
func MTL4MachineLearningPipelineReflectionFrom(ptr unsafe.Pointer) MTL4MachineLearningPipelineReflection {
	return MTL4MachineLearningPipelineReflection{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTL4MachineLearningPipelineReflection *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTL4MachineLearningPipelineReflection */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTL4MachineLearningPipelineReflection */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTL4MachineLearningPipelineReflection */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTL4MachineLearningPipelineReflection */

// Describes every input and output of the pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MachineLearningPipelineReflection/bindings
func (m_ MTL4MachineLearningPipelineReflection) Bindings() []objc.ID {
	rv := objc.Send[[]objc.ID](m_.ID, objc.Sel("bindings"))
	return rv
}/* debug [instance_properties/getter]: bindings */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTL4MachineLearningPipelineReflection */



