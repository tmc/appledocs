// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [PHASESpatialPipeline] class.
var (
	PHASESpatialPipelineClass     _PHASESpatialPipelineClass
	PHASESpatialPipelineClassOnce sync.Once
)

func getPHASESpatialPipelineClass() _PHASESpatialPipelineClass {
	PHASESpatialPipelineClassOnce.Do(func() {
		PHASESpatialPipelineClass = _PHASESpatialPipelineClass{objc.GetClass("PHASESpatialPipeline")}
	})
	return PHASESpatialPipelineClass
}

type _PHASESpatialPipelineClass struct {
	class objc.Class
}

// An interface definition for the [PHASESpatialPipeline] class.
type IPHASESpatialPipeline interface {
	objectivec.IObject
}

// An object that specifies the volume of optional environmental effects.
//
// The class contains an instance of this class, , to add optional sound layers to the output. On top of the original audio signal designated by , this class optionally includes audio layers for environmental effects, such as or , in the output. To control the amount of volume that either audio layer possesses in the mixer’s output, adjust the for the layer’s respective member in the dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESpatialPipeline
type PHASESpatialPipeline struct {
	objectivec.Object
}

// PHASESpatialPipelineFrom constructs a [PHASESpatialPipeline] from an unsafe.Pointer.
//
// An object that specifies the volume of optional environmental effects.
func PHASESpatialPipelineFrom(ptr unsafe.Pointer) PHASESpatialPipeline {
	return PHASESpatialPipeline{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHASESpatialPipelineClass) Alloc() PHASESpatialPipeline {
	rv := objc.Send[PHASESpatialPipeline](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHASESpatialPipelineClass) New() PHASESpatialPipeline {
	rv := objc.Send[PHASESpatialPipeline](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASESpatialPipeline) Init() PHASESpatialPipeline {
	rv := objc.Send[PHASESpatialPipeline](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASESpatialPipeline) Autorelease() PHASESpatialPipeline {
	rv := objc.Send[PHASESpatialPipeline](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASESpatialPipeline creates a new PHASESpatialPipeline instance.
func NewPHASESpatialPipeline() PHASESpatialPipeline {
	return getPHASESpatialPipelineClass().New()
}


// Creates a spatial pipeline with the specified flags.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESpatialPipeline/init(flags:)
func NewPHASESpatialPipelineWithFlags(flags unsafe.Pointer) PHASESpatialPipeline {
	instance := getPHASESpatialPipelineClass().Alloc()
	rv := objc.Send[PHASESpatialPipeline](instance.ID, objc.Sel("initWithFlags:"), flags)
	rv.Autorelease()
	return rv
}


// Audio layers for environmental effects to add to the output.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESpatialPipeline/entries
func (p_ PHASESpatialPipeline) Entries() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("entries"))
	return rv
}

// A collection of environmental effects to include in the output.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESpatialPipeline/flags-swift.property
func (p_ PHASESpatialPipeline) Flags() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("flags"))
	return rv
}


