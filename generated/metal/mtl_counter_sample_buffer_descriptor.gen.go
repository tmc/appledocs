// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [CounterSampleBufferDescriptor] class.
type ICounterSampleBufferDescriptor interface {
	objectivec.IObject
	

	// properties:
	CounterSet() unsafe.Pointer
	SetCounterSet(value unsafe.Pointer)
	Label() foundation.foundation.INSString
	SetLabel(value foundation.foundation.INSString)
	SampleCount() uint
	SetSampleCount(value uint)
	StorageMode() StorageMode
	SetStorageMode(value StorageMode)
	MTLCounterDontSample() int
	SetMTLCounterDontSample(value int)


	

	// methods:


}





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

























// A GPU device’s counter set instance that you want to sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCounterSampleBufferDescriptor/counterSet
func (c_ CounterSampleBufferDescriptor) CounterSet() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("counterSet"))
	return rv
}


// A GPU device’s counter set instance that you want to sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCounterSampleBufferDescriptor/counterSet
func (c_ CounterSampleBufferDescriptor) SetCounterSet(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCounterSet:"), value)
}


// The name for the counter sample buffer you create with the descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCounterSampleBufferDescriptor/label
func (c_ CounterSampleBufferDescriptor) Label() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("label"))
	return rv
}


// The name for the counter sample buffer you create with the descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCounterSampleBufferDescriptor/label
func (c_ CounterSampleBufferDescriptor) SetLabel(value foundation.foundation.INSString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLabel:"), value)
}


// The number of instances of a counter set’s data that a counter sample buffer can store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCounterSampleBufferDescriptor/sampleCount
func (c_ CounterSampleBufferDescriptor) SampleCount() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("sampleCount"))
	return rv
}


// The number of instances of a counter set’s data that a counter sample buffer can store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCounterSampleBufferDescriptor/sampleCount
func (c_ CounterSampleBufferDescriptor) SetSampleCount(value uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSampleCount:"), value)
}


// The memory storage mode for the counter sample buffers you create with the descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCounterSampleBufferDescriptor/storageMode
func (c_ CounterSampleBufferDescriptor) StorageMode() StorageMode {
	rv := objc.Send[StorageMode](c_.ID, objc.Sel("storageMode"))
	return rv
}


// The memory storage mode for the counter sample buffers you create with the descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCounterSampleBufferDescriptor/storageMode
func (c_ CounterSampleBufferDescriptor) SetStorageMode(value StorageMode) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStorageMode:"), value)
}


// A sentinel value that instructs an encoder to skip sampling a counter as the GPU runs the encoder’s pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcounterdontsample
func (c_ CounterSampleBufferDescriptor) MTLCounterDontSample() int {
	rv := objc.Send[int](c_.ID, objc.Sel("MTLCounterDontSample"))
	return rv
}


// A sentinel value that instructs an encoder to skip sampling a counter as the GPU runs the encoder’s pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcounterdontsample
func (c_ CounterSampleBufferDescriptor) SetMTLCounterDontSample(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMTLCounterDontSample:"), value)
}








