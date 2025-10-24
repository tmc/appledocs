// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MLCDropoutLayer */


/* debug [class_header]: Header for MLCDropoutLayer */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CDropoutLayer */
// An interface definition for the [CDropoutLayer] class.
type ICDropoutLayer interface {
	ICLayer
	
/* debug [class_interface_properties]: Properties for CDropoutLayer */
	// properties:
	Rate() float32
	Seed() uint
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CDropoutLayer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CDropoutLayer */
// Alloc allocates a new instance without initialization.
func (cc _CDropoutLayerClass) Alloc() CDropoutLayer {
	rv := objc.Send[CDropoutLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CDropoutLayer */
// A layer that deactivates neurons randomly to avoid overfitting.


// A layer that deactivates neurons randomly to avoid overfitting.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CDropoutLayer */

// Creates a dropout layer with the probability rate and random number generator seed you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDropoutLayer/init(rate:seed:)
func NewCDropoutLayerWithRateSeed(rate float32, seed uint) CDropoutLayer {
	rv := objc.Send[CDropoutLayer](objc.ID(getCDropoutLayerClass().class), objc.Sel("layerWithRate:seed:"), rate, seed)
	return rv
}/* debug [class_init_methods/constructor]: NewCDropoutLayerWithRateSeed */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CDropoutLayer */

// Creates a dropout layer with the probability rate and random number generator seed you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDropoutLayer/init(rate:seed:)
func (cc _CDropoutLayerClass) LayerWithRateSeed(rate float32, seed uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layerWithRate:seed:"), rate, seed)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayerWithRateSeed) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CDropoutLayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CDropoutLayer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CDropoutLayer */

// The dropout rate you use for each element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDropoutLayer/rate
func (c_ CDropoutLayer) Rate() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("rate"))
	return rv
}/* debug [instance_properties/getter]: rate */


// The seed you use to generate random numbers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCDropoutLayer/seed
func (c_ CDropoutLayer) Seed() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("seed"))
	return rv
}/* debug [instance_properties/getter]: seed */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCDropoutLayer */


