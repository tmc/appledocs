// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTL4MachineLearningPipelineDescriptor] class.
var mTL4MachineLearningPipelineDescriptorClass = _MTL4MachineLearningPipelineDescriptorClass{objc.GetClass("MTL4MachineLearningPipelineDescriptor")}

type _MTL4MachineLearningPipelineDescriptorClass struct {
	class objc.Class
}

// Description for a machine learning pipeline state. [Full Topic]
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



