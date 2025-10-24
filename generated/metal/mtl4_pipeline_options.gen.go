// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTL4PipelineOptions */


/* debug [class_header]: Header for MTL4PipelineOptions */
// The class instance for the [MTL4PipelineOptions] class.
var (
	MTL4PipelineOptionsClass     _MTL4PipelineOptionsClass
	MTL4PipelineOptionsClassOnce sync.Once
)

func getMTL4PipelineOptionsClass() _MTL4PipelineOptionsClass {
	MTL4PipelineOptionsClassOnce.Do(func() {
		MTL4PipelineOptionsClass = _MTL4PipelineOptionsClass{objc.GetClass("MTL4PipelineOptions")}
	})
	return MTL4PipelineOptionsClass
}

type _MTL4PipelineOptionsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTL4PipelineOptions */
// An interface definition for the [MTL4PipelineOptions] class.
type IMTL4PipelineOptions interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTL4PipelineOptions */
	// properties:
	ShaderReflection() MTL4ShaderReflection
	SetShaderReflection(value MTL4ShaderReflection)
	ShaderValidation() ShaderValidation
	SetShaderValidation(value ShaderValidation)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTL4PipelineOptions */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTL4PipelineOptions */
// Alloc allocates a new instance without initialization.
func (mc _MTL4PipelineOptionsClass) Alloc() MTL4PipelineOptions {
	rv := objc.Send[MTL4PipelineOptions](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTL4PipelineOptionsClass) New() MTL4PipelineOptions {
	rv := objc.Send[MTL4PipelineOptions](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4PipelineOptions) Init() MTL4PipelineOptions {
	rv := objc.Send[MTL4PipelineOptions](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4PipelineOptions) Autorelease() MTL4PipelineOptions {
	rv := objc.Send[MTL4PipelineOptions](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4PipelineOptions creates a new MTL4PipelineOptions instance.
func NewMTL4PipelineOptions() MTL4PipelineOptions {
	return getMTL4PipelineOptionsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTL4PipelineOptions */
// Provides options controlling how to compile a pipeline state.
//
// You provide these options through the class at compilation time.


// Provides options controlling how to compile a pipeline state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PipelineOptions
type MTL4PipelineOptions struct {
	objectivec.Object
}

// MTL4PipelineOptionsFrom constructs a [MTL4PipelineOptions] from an unsafe.Pointer.
//
// Provides options controlling how to compile a pipeline state.
func MTL4PipelineOptionsFrom(ptr unsafe.Pointer) MTL4PipelineOptions {
	return MTL4PipelineOptions{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTL4PipelineOptions *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTL4PipelineOptions */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTL4PipelineOptions */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTL4PipelineOptions */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTL4PipelineOptions */

// Controls whether to include Metal shader reflection in this pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PipelineOptions/shaderReflection
func (m_ MTL4PipelineOptions) ShaderReflection() MTL4ShaderReflection {
	rv := objc.Send[MTL4ShaderReflection](m_.ID, objc.Sel("shaderReflection"))
	return rv
}/* debug [instance_properties/getter]: shaderReflection */


// Controls whether to include Metal shader reflection in this pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PipelineOptions/shaderReflection
func (m_ MTL4PipelineOptions) SetShaderReflection(value MTL4ShaderReflection) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShaderReflection:"), value)
}/* debug [instance_properties/setter]: shaderReflection */


// Controls whether to enable or disable Metal Shader Validation for the pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PipelineOptions/shaderValidation
func (m_ MTL4PipelineOptions) ShaderValidation() ShaderValidation {
	rv := objc.Send[ShaderValidation](m_.ID, objc.Sel("shaderValidation"))
	return rv
}/* debug [instance_properties/getter]: shaderValidation */


// Controls whether to enable or disable Metal Shader Validation for the pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PipelineOptions/shaderValidation
func (m_ MTL4PipelineOptions) SetShaderValidation(value ShaderValidation) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShaderValidation:"), value)
}/* debug [instance_properties/setter]: shaderValidation */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTL4PipelineOptions */



