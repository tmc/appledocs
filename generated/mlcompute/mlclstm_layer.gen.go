// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CLSTMLayer] class.
var (
	CLSTMLayerClass     _CLSTMLayerClass
	CLSTMLayerClassOnce sync.Once
)

func getCLSTMLayerClass() _CLSTMLayerClass {
	CLSTMLayerClassOnce.Do(func() {
		CLSTMLayerClass = _CLSTMLayerClass{objc.GetClass("MLCLSTMLayer")}
	})
	return CLSTMLayerClass
}

type _CLSTMLayerClass struct {
	class objc.Class
}

// An interface definition for the [CLSTMLayer] class.
type ICLSTMLayer interface {
	ICLayer
}

// A layer that represents long short-term memory (LSTM) networks.
//
// Use this class to create an LSTM layer with one of the following configurations:
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCLSTMLayer
type CLSTMLayer struct {
	CLayer
}

// CLSTMLayerFrom constructs a [CLSTMLayer] from an unsafe.Pointer.
//
// A layer that represents long short-term memory (LSTM) networks.
func CLSTMLayerFrom(ptr unsafe.Pointer) CLSTMLayer {
	return CLSTMLayer{
		CLayer: CLayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _CLSTMLayerClass) Alloc() CLSTMLayer {
	rv := objc.Send[CLSTMLayer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _CLSTMLayerClass) New() CLSTMLayer {
	rv := objc.Send[CLSTMLayer](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ CLSTMLayer) Init() CLSTMLayer {
	rv := objc.Send[CLSTMLayer](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ CLSTMLayer) Autorelease() CLSTMLayer {
	rv := objc.Send[CLSTMLayer](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCLSTMLayer creates a new CLSTMLayer instance.
func NewCLSTMLayer() CLSTMLayer {
	return getCLSTMLayerClass().New()
}




