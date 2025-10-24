// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [Compare] class.
var (
	CompareClass     _CompareClass
	CompareClassOnce sync.Once
)

func getCompareClass() _CompareClass {
	CompareClassOnce.Do(func() {
		CompareClass = _CompareClass{objc.GetClass("MPSNNCompare")}
	})
	return CompareClass
}

type _CompareClass struct {
	class objc.Class
}





// An interface definition for the [Compare] class.
type ICompare interface {
	ICNNArithmetic
	

	// properties:
	ComparisonType() ComparisonType get set /* not a class type */
	SetComparisonType(value ComparisonType get set /* not a class type */)
	Threshold() objectivec.IObject
	SetThreshold(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CompareClass) Alloc() Compare {
	rv := objc.Send[Compare](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CompareClass) New() Compare {
	rv := objc.Send[Compare](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ Compare) Init() Compare {
	rv := objc.Send[Compare](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ Compare) Autorelease() Compare {
	rv := objc.Send[Compare](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCompare creates a new Compare instance.
func NewCompare() Compare {
	return getCompareClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNCompare
type Compare struct {
	CNNArithmetic
}

// CompareFrom constructs a [Compare] from an unsafe.Pointer.
func CompareFrom(ptr unsafe.Pointer) Compare {
	return Compare{
		CNNArithmetic: CNNArithmeticFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncompare/3037375-initwithdevice
func NewCompareWithDevice(device unsafe.Pointer) Compare {
	instance := getCompareClass().Alloc()
	rv := objc.Send[Compare](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}






















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncompare/3037374-comparisontype
func (c_ Compare) ComparisonType() ComparisonType get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("comparisonType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncompare/3037374-comparisontype
func (c_ Compare) SetComparisonType(value ComparisonType get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setComparisonType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncompare/3037376-threshold
func (c_ Compare) Threshold() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("threshold"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncompare/3037376-threshold
func (c_ Compare) SetThreshold(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setThreshold:"), value)
}







