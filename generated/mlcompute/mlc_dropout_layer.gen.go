// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CDropoutLayer] class.
var (
	CDropoutLayerClass     _CDropoutLayerClass
	CDropoutLayerClassOnce sync.Once
)

func getCDropoutLayerClass() _CDropoutLayerClass {
	CDropoutLayerClassOnce.Do(func() {
		CDropoutLayerClass = _CDropoutLayerClass{objc.GetClass("MLCDropoutLayer")}
	})
	return CDropoutLayerClass
}

type _CDropoutLayerClass struct {
	class objc.Class
}

// An interface definition for the [CDropoutLayer] class.
type ICDropoutLayer interface {
	ICLayer
}

// A layer that deactivates neurons randomly to avoid overfitting.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDropoutLayer
type CDropoutLayer struct {
	CLayer
}

// CDropoutLayerFrom constructs a [CDropoutLayer] from an unsafe.Pointer.
//
// A layer that deactivates neurons randomly to avoid overfitting.
func CDropoutLayerFrom(ptr unsafe.Pointer) CDropoutLayer {
	return CDropoutLayer{
		CLayer: CLayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CDropoutLayerClass) Alloc() CDropoutLayer {
	rv := objc.Send[CDropoutLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CDropoutLayerClass) New() CDropoutLayer {
	rv := objc.Send[CDropoutLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CDropoutLayer) Init() CDropoutLayer {
	rv := objc.Send[CDropoutLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CDropoutLayer) Autorelease() CDropoutLayer {
	rv := objc.Send[CDropoutLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCDropoutLayer creates a new CDropoutLayer instance.
func NewCDropoutLayer() CDropoutLayer {
	return getCDropoutLayerClass().New()
}


// The seed you use to generate random numbers.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcdropoutlayer/seed
func (c_ CDropoutLayer) Seed() int {
	rv := objc.Send[int](c_.ID, objc.Sel("seed"))
	return rv
}


// SetSeed sets the value of the seed property.
// The seed you use to generate random numbers.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcdropoutlayer/seed
func (c_ CDropoutLayer) SetSeed(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSeed:"), value)
}

// The dropout rate you use for each element.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcdropoutlayer/rate
func (c_ CDropoutLayer) Rate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("rate"))
	return rv
}


// SetRate sets the value of the rate property.
// The dropout rate you use for each element.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcdropoutlayer/rate
func (c_ CDropoutLayer) SetRate(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRate:"), value)
}



