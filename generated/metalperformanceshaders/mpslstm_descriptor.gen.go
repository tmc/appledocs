// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [LSTMDescriptor] class.
var (
	LSTMDescriptorClass     _LSTMDescriptorClass
	LSTMDescriptorClassOnce sync.Once
)

func getLSTMDescriptorClass() _LSTMDescriptorClass {
	LSTMDescriptorClassOnce.Do(func() {
		LSTMDescriptorClass = _LSTMDescriptorClass{objc.GetClass("MPSLSTMDescriptor")}
	})
	return LSTMDescriptorClass
}

type _LSTMDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [LSTMDescriptor] class.
type ILSTMDescriptor interface {
	objectivec.IObject
}

// A description of a long short-term memory block or layer.
//
// The recurrent neural network (RNN) layer initialized with transforms the input data (image or matrix), the memory cell data, and previous output with a set of filters. Each produces one feature map in the output data and memory cell according to the long short-term memory (LSTM) formula detailed below. You may provide the LSTM unit with a single input or a sequence of inputs.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSLSTMDescriptor
type LSTMDescriptor struct {
	objectivec.Object
}

// LSTMDescriptorFrom constructs a [LSTMDescriptor] from an unsafe.Pointer.
//
// A description of a long short-term memory block or layer.
func LSTMDescriptorFrom(ptr unsafe.Pointer) LSTMDescriptor {
	return LSTMDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (lc _LSTMDescriptorClass) Alloc() LSTMDescriptor {
	rv := objc.Send[LSTMDescriptor](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (lc _LSTMDescriptorClass) New() LSTMDescriptor {
	rv := objc.Send[LSTMDescriptor](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LSTMDescriptor) Init() LSTMDescriptor {
	rv := objc.Send[LSTMDescriptor](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LSTMDescriptor) Autorelease() LSTMDescriptor {
	rv := objc.Send[LSTMDescriptor](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLSTMDescriptor creates a new LSTMDescriptor instance.
func NewLSTMDescriptor() LSTMDescriptor {
	return getLSTMDescriptorClass().New()
}




