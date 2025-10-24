// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ModelStructurePipeline] class.
var (
	ModelStructurePipelineClass     _ModelStructurePipelineClass
	ModelStructurePipelineClassOnce sync.Once
)

func getModelStructurePipelineClass() _ModelStructurePipelineClass {
	ModelStructurePipelineClassOnce.Do(func() {
		ModelStructurePipelineClass = _ModelStructurePipelineClass{objc.GetClass("MLModelStructurePipeline")}
	})
	return ModelStructurePipelineClass
}

type _ModelStructurePipelineClass struct {
	class objc.Class
}

// An interface definition for the [ModelStructurePipeline] class.
type IModelStructurePipeline interface {
	objectivec.IObject
	// properties:
	SubModelNames() []string
	SubModels() []IModelStructure
	// methods:
}

// A class representing the structure of a Pipeline model.


// A class representing the structure of a Pipeline model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructurePipeline
type ModelStructurePipeline struct {
	objectivec.Object
}

// ModelStructurePipelineFrom constructs a [ModelStructurePipeline] from an unsafe.Pointer.
//
// A class representing the structure of a Pipeline model.
func ModelStructurePipelineFrom(ptr unsafe.Pointer) ModelStructurePipeline {
	return ModelStructurePipeline{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _ModelStructurePipelineClass) Alloc() ModelStructurePipeline {
	rv := objc.Send[ModelStructurePipeline](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _ModelStructurePipelineClass) New() ModelStructurePipeline {
	rv := objc.Send[ModelStructurePipeline](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ ModelStructurePipeline) Init() ModelStructurePipeline {
	rv := objc.Send[ModelStructurePipeline](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ ModelStructurePipeline) Autorelease() ModelStructurePipeline {
	rv := objc.Send[ModelStructurePipeline](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewModelStructurePipeline creates a new ModelStructurePipeline instance.
func NewModelStructurePipeline() ModelStructurePipeline {
	return getModelStructurePipelineClass().New()
}



// The names of the sub models in the pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructurePipeline/subModelNames
func (m_ ModelStructurePipeline) SubModelNames() []string {
	rv := objc.Send[[]string](m_.ID, objc.Sel("subModelNames"))
	return rv
}


// The structure of the sub models in the pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructurePipeline/subModels
func (m_ ModelStructurePipeline) SubModels() []IModelStructure {
	rv := objc.Send[[]ModelStructure](m_.ID, objc.Sel("subModels"))
	return rv
}



