// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MTL4PipelineDataSetSerializerDescriptor] class.
var (
	MTL4PipelineDataSetSerializerDescriptorClass     _MTL4PipelineDataSetSerializerDescriptorClass
	MTL4PipelineDataSetSerializerDescriptorClassOnce sync.Once
)

func getMTL4PipelineDataSetSerializerDescriptorClass() _MTL4PipelineDataSetSerializerDescriptorClass {
	MTL4PipelineDataSetSerializerDescriptorClassOnce.Do(func() {
		MTL4PipelineDataSetSerializerDescriptorClass = _MTL4PipelineDataSetSerializerDescriptorClass{objc.GetClass("MTL4PipelineDataSetSerializerDescriptor")}
	})
	return MTL4PipelineDataSetSerializerDescriptorClass
}

type _MTL4PipelineDataSetSerializerDescriptorClass struct {
	class objc.Class
}





// An interface definition for the [MTL4PipelineDataSetSerializerDescriptor] class.
type IMTL4PipelineDataSetSerializerDescriptor interface {
	objectivec.IObject
	

	// properties:
	Configuration() MTL4PipelineDataSetSerializerConfiguration
	SetConfiguration(value MTL4PipelineDataSetSerializerConfiguration)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MTL4PipelineDataSetSerializerDescriptorClass) Alloc() MTL4PipelineDataSetSerializerDescriptor {
	rv := objc.Send[MTL4PipelineDataSetSerializerDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTL4PipelineDataSetSerializerDescriptorClass) New() MTL4PipelineDataSetSerializerDescriptor {
	rv := objc.Send[MTL4PipelineDataSetSerializerDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4PipelineDataSetSerializerDescriptor) Init() MTL4PipelineDataSetSerializerDescriptor {
	rv := objc.Send[MTL4PipelineDataSetSerializerDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4PipelineDataSetSerializerDescriptor) Autorelease() MTL4PipelineDataSetSerializerDescriptor {
	rv := objc.Send[MTL4PipelineDataSetSerializerDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4PipelineDataSetSerializerDescriptor creates a new MTL4PipelineDataSetSerializerDescriptor instance.
func NewMTL4PipelineDataSetSerializerDescriptor() MTL4PipelineDataSetSerializerDescriptor {
	return getMTL4PipelineDataSetSerializerDescriptorClass().New()
}





// Groups together properties to create a pipeline data set serializer.


// Groups together properties to create a pipeline data set serializer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PipelineDataSetSerializerDescriptor
type MTL4PipelineDataSetSerializerDescriptor struct {
	objectivec.Object
}

// MTL4PipelineDataSetSerializerDescriptorFrom constructs a [MTL4PipelineDataSetSerializerDescriptor] from an unsafe.Pointer.
//
// Groups together properties to create a pipeline data set serializer.
func MTL4PipelineDataSetSerializerDescriptorFrom(ptr unsafe.Pointer) MTL4PipelineDataSetSerializerDescriptor {
	return MTL4PipelineDataSetSerializerDescriptor{objectivec.Object{objc.ID(ptr)}}
}

























// Specifies the configuration of the serialization process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PipelineDataSetSerializerDescriptor/configuration
func (m_ MTL4PipelineDataSetSerializerDescriptor) Configuration() MTL4PipelineDataSetSerializerConfiguration {
	rv := objc.Send[MTL4PipelineDataSetSerializerConfiguration](m_.ID, objc.Sel("configuration"))
	return rv
}


// Specifies the configuration of the serialization process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4PipelineDataSetSerializerDescriptor/configuration
func (m_ MTL4PipelineDataSetSerializerDescriptor) SetConfiguration(value MTL4PipelineDataSetSerializerConfiguration) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setConfiguration:"), value)
}








