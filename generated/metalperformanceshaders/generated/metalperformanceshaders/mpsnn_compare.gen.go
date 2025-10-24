// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNCompare */


/* debug [class_header]: Header for MPSNNCompare */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Compare */
// An interface definition for the [Compare] class.
type ICompare interface {
	ICNNArithmetic
	
/* debug [class_interface_properties]: Properties for Compare */
	// properties:
	ComparisonType() ComparisonType get set /* not a class type */
	SetComparisonType(value ComparisonType get set /* not a class type */)
	Threshold() objectivec.IObject
	SetThreshold(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Compare */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Compare */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Compare */


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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Compare */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncompare/3037375-initwithdevice
func NewCompareWithDevice(device unsafe.Pointer) Compare {
	instance := getCompareClass().Alloc()
	rv := objc.Send[Compare](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCompareWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Compare */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Compare */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Compare */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Compare */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncompare/3037374-comparisontype
func (c_ Compare) ComparisonType() ComparisonType get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("comparisonType"))
	return rv
}/* debug [instance_properties/getter]: comparisonType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncompare/3037374-comparisontype
func (c_ Compare) SetComparisonType(value ComparisonType get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setComparisonType:"), value)
}/* debug [instance_properties/setter]: comparisonType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncompare/3037376-threshold
func (c_ Compare) Threshold() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("threshold"))
	return rv
}/* debug [instance_properties/getter]: threshold */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncompare/3037376-threshold
func (c_ Compare) SetThreshold(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setThreshold:"), value)
}/* debug [instance_properties/setter]: threshold */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNCompare */


