// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLCounterSampleBufferDescriptor */


/* debug [class_header]: Header for MTLCounterSampleBufferDescriptor */
// The class instance for the [CounterSampleBufferDescriptor] class.
var (
	CounterSampleBufferDescriptorClass     _CounterSampleBufferDescriptorClass
	CounterSampleBufferDescriptorClassOnce sync.Once
)

func getCounterSampleBufferDescriptorClass() _CounterSampleBufferDescriptorClass {
	CounterSampleBufferDescriptorClassOnce.Do(func() {
		CounterSampleBufferDescriptorClass = _CounterSampleBufferDescriptorClass{objc.GetClass("MTLCounterSampleBufferDescriptor")}
	})
	return CounterSampleBufferDescriptorClass
}

type _CounterSampleBufferDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CounterSampleBufferDescriptor */
// An interface definition for the [CounterSampleBufferDescriptor] class.
type ICounterSampleBufferDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CounterSampleBufferDescriptor */
	// properties:
	CounterSet() unsafe.Pointer
	SetCounterSet(value unsafe.Pointer)
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	SampleCount() uint
	SetSampleCount(value uint)
	StorageMode() StorageMode
	SetStorageMode(value StorageMode)
	MTLCounterDontSample() int
	SetMTLCounterDontSample(value int)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CounterSampleBufferDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CounterSampleBufferDescriptor */
// Alloc allocates a new instance without initialization.
func (cc _CounterSampleBufferDescriptorClass) Alloc() CounterSampleBufferDescriptor {
	rv := objc.Send[CounterSampleBufferDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CounterSampleBufferDescriptorClass) New() CounterSampleBufferDescriptor {
	rv := objc.Send[CounterSampleBufferDescriptor](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CounterSampleBufferDescriptor) Init() CounterSampleBufferDescriptor {
	rv := objc.Send[CounterSampleBufferDescriptor](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CounterSampleBufferDescriptor) Autorelease() CounterSampleBufferDescriptor {
	rv := objc.Send[CounterSampleBufferDescriptor](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCounterSampleBufferDescriptor creates a new CounterSampleBufferDescriptor instance.
func NewCounterSampleBufferDescriptor() CounterSampleBufferDescriptor {
	return getCounterSampleBufferDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CounterSampleBufferDescriptor */
// A group of properties that configures the counter sample buffers you create with it.
//
// To create a new counter sample buffer, create and configure an instance, and then call an instance’s method. See . Each new sample counter buffer inherits the values of the descriptor’s properties when you create it. You can modify a descriptor and reuse it to create other counter sample buffers, which has no effect on existing counter sample buffers.


// A group of properties that configures the counter sample buffers you create with it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCounterSampleBufferDescriptor
type CounterSampleBufferDescriptor struct {
	objectivec.Object
}

// CounterSampleBufferDescriptorFrom constructs a [CounterSampleBufferDescriptor] from an unsafe.Pointer.
//
// A group of properties that configures the counter sample buffers you create with it.
func CounterSampleBufferDescriptorFrom(ptr unsafe.Pointer) CounterSampleBufferDescriptor {
	return CounterSampleBufferDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CounterSampleBufferDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CounterSampleBufferDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CounterSampleBufferDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CounterSampleBufferDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CounterSampleBufferDescriptor */

// A GPU device’s counter set instance that you want to sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCounterSampleBufferDescriptor/counterSet
func (c_ CounterSampleBufferDescriptor) CounterSet() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("counterSet"))
	return rv
}/* debug [instance_properties/getter]: counterSet */


// A GPU device’s counter set instance that you want to sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCounterSampleBufferDescriptor/counterSet
func (c_ CounterSampleBufferDescriptor) SetCounterSet(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCounterSet:"), value)
}/* debug [instance_properties/setter]: counterSet */


// The name for the counter sample buffer you create with the descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCounterSampleBufferDescriptor/label
func (c_ CounterSampleBufferDescriptor) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// The name for the counter sample buffer you create with the descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCounterSampleBufferDescriptor/label
func (c_ CounterSampleBufferDescriptor) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */


// The number of instances of a counter set’s data that a counter sample buffer can store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCounterSampleBufferDescriptor/sampleCount
func (c_ CounterSampleBufferDescriptor) SampleCount() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("sampleCount"))
	return rv
}/* debug [instance_properties/getter]: sampleCount */


// The number of instances of a counter set’s data that a counter sample buffer can store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCounterSampleBufferDescriptor/sampleCount
func (c_ CounterSampleBufferDescriptor) SetSampleCount(value uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSampleCount:"), value)
}/* debug [instance_properties/setter]: sampleCount */


// The memory storage mode for the counter sample buffers you create with the descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCounterSampleBufferDescriptor/storageMode
func (c_ CounterSampleBufferDescriptor) StorageMode() StorageMode {
	rv := objc.Send[StorageMode](c_.ID, objc.Sel("storageMode"))
	return rv
}/* debug [instance_properties/getter]: storageMode */


// The memory storage mode for the counter sample buffers you create with the descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCounterSampleBufferDescriptor/storageMode
func (c_ CounterSampleBufferDescriptor) SetStorageMode(value StorageMode) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStorageMode:"), value)
}/* debug [instance_properties/setter]: storageMode */


// A sentinel value that instructs an encoder to skip sampling a counter as the GPU runs the encoder’s pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcounterdontsample
func (c_ CounterSampleBufferDescriptor) MTLCounterDontSample() int {
	rv := objc.Send[int](c_.ID, objc.Sel("MTLCounterDontSample"))
	return rv
}/* debug [instance_properties/getter]: MTLCounterDontSample */


// A sentinel value that instructs an encoder to skip sampling a counter as the GPU runs the encoder’s pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcounterdontsample
func (c_ CounterSampleBufferDescriptor) SetMTLCounterDontSample(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMTLCounterDontSample:"), value)
}/* debug [instance_properties/setter]: MTLCounterDontSample */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLCounterSampleBufferDescriptor */



