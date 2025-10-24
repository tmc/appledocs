// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLBufferLayoutDescriptor */


/* debug [class_header]: Header for MTLBufferLayoutDescriptor */
// The class instance for the [BufferLayoutDescriptor] class.
var (
	BufferLayoutDescriptorClass     _BufferLayoutDescriptorClass
	BufferLayoutDescriptorClassOnce sync.Once
)

func getBufferLayoutDescriptorClass() _BufferLayoutDescriptorClass {
	BufferLayoutDescriptorClassOnce.Do(func() {
		BufferLayoutDescriptorClass = _BufferLayoutDescriptorClass{objc.GetClass("MTLBufferLayoutDescriptor")}
	})
	return BufferLayoutDescriptorClass
}

type _BufferLayoutDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BufferLayoutDescriptor */
// An interface definition for the [BufferLayoutDescriptor] class.
type IBufferLayoutDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for BufferLayoutDescriptor */
	// properties:
	StepFunction() StepFunction
	SetStepFunction(value StepFunction)
	StepRate() uint
	SetStepRate(value uint)
	Stride() uint
	SetStride(value uint)
	StageInputDescriptor() IMTLStageInputOutputDescriptor
	SetStageInputDescriptor(value IMTLStageInputOutputDescriptor)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BufferLayoutDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BufferLayoutDescriptor */
// Alloc allocates a new instance without initialization.
func (bc _BufferLayoutDescriptorClass) Alloc() BufferLayoutDescriptor {
	rv := objc.Send[BufferLayoutDescriptor](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _BufferLayoutDescriptorClass) New() BufferLayoutDescriptor {
	rv := objc.Send[BufferLayoutDescriptor](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BufferLayoutDescriptor) Init() BufferLayoutDescriptor {
	rv := objc.Send[BufferLayoutDescriptor](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BufferLayoutDescriptor) Autorelease() BufferLayoutDescriptor {
	rv := objc.Send[BufferLayoutDescriptor](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBufferLayoutDescriptor creates a new BufferLayoutDescriptor instance.
func NewBufferLayoutDescriptor() BufferLayoutDescriptor {
	return getBufferLayoutDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BufferLayoutDescriptor */
// A description of how a compute function fetches input data for an attribute.


// A description of how a compute function fetches input data for an attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBufferLayoutDescriptor
type BufferLayoutDescriptor struct {
	objectivec.Object
}

// BufferLayoutDescriptorFrom constructs a [BufferLayoutDescriptor] from an unsafe.Pointer.
//
// A description of how a compute function fetches input data for an attribute.
func BufferLayoutDescriptorFrom(ptr unsafe.Pointer) BufferLayoutDescriptor {
	return BufferLayoutDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BufferLayoutDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BufferLayoutDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BufferLayoutDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BufferLayoutDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BufferLayoutDescriptor */

// Determines how and when compute functions fetch data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBufferLayoutDescriptor/stepFunction
func (b_ BufferLayoutDescriptor) StepFunction() StepFunction {
	rv := objc.Send[StepFunction](b_.ID, objc.Sel("stepFunction"))
	return rv
}/* debug [instance_properties/getter]: stepFunction */


// Determines how and when compute functions fetch data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBufferLayoutDescriptor/stepFunction
func (b_ BufferLayoutDescriptor) SetStepFunction(value StepFunction) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setStepFunction:"), value)
}/* debug [instance_properties/setter]: stepFunction */


// How frequently the step function should load data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBufferLayoutDescriptor/stepRate
func (b_ BufferLayoutDescriptor) StepRate() uint {
	rv := objc.Send[uint](b_.ID, objc.Sel("stepRate"))
	return rv
}/* debug [instance_properties/getter]: stepRate */


// How frequently the step function should load data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBufferLayoutDescriptor/stepRate
func (b_ BufferLayoutDescriptor) SetStepRate(value uint) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setStepRate:"), value)
}/* debug [instance_properties/setter]: stepRate */


// The number of bytes from one buffer entry to the next.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBufferLayoutDescriptor/stride
func (b_ BufferLayoutDescriptor) Stride() uint {
	rv := objc.Send[uint](b_.ID, objc.Sel("stride"))
	return rv
}/* debug [instance_properties/getter]: stride */


// The number of bytes from one buffer entry to the next.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBufferLayoutDescriptor/stride
func (b_ BufferLayoutDescriptor) SetStride(value uint) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setStride:"), value)
}/* debug [instance_properties/setter]: stride */


// The organization of input and output data for the next kernel call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/stageinputdescriptor
func (b_ BufferLayoutDescriptor) StageInputDescriptor() IMTLStageInputOutputDescriptor {
	rv := objc.Send[StageInputOutputDescriptor](b_.ID, objc.Sel("stageInputDescriptor"))
	return rv
}/* debug [instance_properties/getter]: stageInputDescriptor */


// The organization of input and output data for the next kernel call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/stageinputdescriptor
func (b_ BufferLayoutDescriptor) SetStageInputDescriptor(value IMTLStageInputOutputDescriptor) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setStageInputDescriptor:"), value)
}/* debug [instance_properties/setter]: stageInputDescriptor */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLBufferLayoutDescriptor */



