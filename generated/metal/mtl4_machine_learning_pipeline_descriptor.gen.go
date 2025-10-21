// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTL4MachineLearningPipelineDescriptor] class.
var (
	MTL4MachineLearningPipelineDescriptorClass     _MTL4MachineLearningPipelineDescriptorClass
	MTL4MachineLearningPipelineDescriptorClassOnce sync.Once
)

func getMTL4MachineLearningPipelineDescriptorClass() _MTL4MachineLearningPipelineDescriptorClass {
	MTL4MachineLearningPipelineDescriptorClassOnce.Do(func() {
		MTL4MachineLearningPipelineDescriptorClass = _MTL4MachineLearningPipelineDescriptorClass{objc.GetClass("MTL4MachineLearningPipelineDescriptor")}
	})
	return MTL4MachineLearningPipelineDescriptorClass
}

type _MTL4MachineLearningPipelineDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [MTL4MachineLearningPipelineDescriptor] class.
type IMTL4MachineLearningPipelineDescriptor interface {
	IMTL4PipelineDescriptor
}

// Description for a machine learning pipeline state.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4MachineLearningPipelineDescriptor
type MTL4MachineLearningPipelineDescriptor struct {
	MTL4PipelineDescriptor
}

// MTL4MachineLearningPipelineDescriptorFrom constructs a [MTL4MachineLearningPipelineDescriptor] from an unsafe.Pointer.
//
// Description for a machine learning pipeline state.
func MTL4MachineLearningPipelineDescriptorFrom(ptr unsafe.Pointer) MTL4MachineLearningPipelineDescriptor {
	return MTL4MachineLearningPipelineDescriptor{
		MTL4PipelineDescriptor: MTL4PipelineDescriptorFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTL4MachineLearningPipelineDescriptorClass) Alloc() MTL4MachineLearningPipelineDescriptor {
	rv := objc.Send[MTL4MachineLearningPipelineDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTL4MachineLearningPipelineDescriptorClass) New() MTL4MachineLearningPipelineDescriptor {
	rv := objc.Send[MTL4MachineLearningPipelineDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4MachineLearningPipelineDescriptor) Init() MTL4MachineLearningPipelineDescriptor {
	rv := objc.Send[MTL4MachineLearningPipelineDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4MachineLearningPipelineDescriptor) Autorelease() MTL4MachineLearningPipelineDescriptor {
	rv := objc.Send[MTL4MachineLearningPipelineDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4MachineLearningPipelineDescriptor creates a new MTL4MachineLearningPipelineDescriptor instance.
func NewMTL4MachineLearningPipelineDescriptor() MTL4MachineLearningPipelineDescriptor {
	return getMTL4MachineLearningPipelineDescriptorClass().New()
}


// Assigns an optional string that helps identify pipeline states you create from this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4machinelearningpipelinedescriptor/label
func (m_ MTL4MachineLearningPipelineDescriptor) Label() string {
	rv := objc.Send[string](m_.ID, objc.Sel("label"))
	return rv
}


// SetLabel sets the value of the label property.
// Assigns an optional string that helps identify pipeline states you create from this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4machinelearningpipelinedescriptor/label
func (m_ MTL4MachineLearningPipelineDescriptor) SetLabel(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), objc.String(value))
}

// Assigns the function that the machine learning pipeline you create from this descriptor executes.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4machinelearningpipelinedescriptor/machinelearningfunctiondescriptor
func (m_ MTL4MachineLearningPipelineDescriptor) MachineLearningFunctionDescriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("machineLearningFunctionDescriptor"))
	return rv
}


// SetMachineLearningFunctionDescriptor sets the value of the machineLearningFunctionDescriptor property.
// Assigns the function that the machine learning pipeline you create from this descriptor executes.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4machinelearningpipelinedescriptor/machinelearningfunctiondescriptor
func (m_ MTL4MachineLearningPipelineDescriptor) SetMachineLearningFunctionDescriptor(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMachineLearningFunctionDescriptor:"), value)
}



