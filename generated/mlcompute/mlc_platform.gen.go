// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLCPlatform */


/* debug [class_header]: Header for MLCPlatform */
// The class instance for the [CPlatform] class.
var (
	CPlatformClass     _CPlatformClass
	CPlatformClassOnce sync.Once
)

func getCPlatformClass() _CPlatformClass {
	CPlatformClassOnce.Do(func() {
		CPlatformClass = _CPlatformClass{objc.GetClass("MLCPlatform")}
	})
	return CPlatformClass
}

type _CPlatformClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CPlatform */
// An interface definition for the [CPlatform] class.
type ICPlatform interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CPlatform */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CPlatform */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CPlatform */
// Alloc allocates a new instance without initialization.
func (cc _CPlatformClass) Alloc() CPlatform {
	rv := objc.Send[CPlatform](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CPlatformClass) New() CPlatform {
	rv := objc.Send[CPlatform](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CPlatform) Init() CPlatform {
	rv := objc.Send[CPlatform](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CPlatform) Autorelease() CPlatform {
	rv := objc.Send[CPlatform](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPlatform creates a new CPlatform instance.
func NewCPlatform() CPlatform {
	return getCPlatformClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CPlatform */
// A utility class for setting global properties in the framework.


// A utility class for setting global properties in the framework.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPlatform
type CPlatform struct {
	objectivec.Object
}

// CPlatformFrom constructs a [CPlatform] from an unsafe.Pointer.
//
// A utility class for setting global properties in the framework.
func CPlatformFrom(ptr unsafe.Pointer) CPlatform {
	return CPlatform{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CPlatform *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CPlatform */

// Returns the global random number generator seed value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPlatform/getRNGseed()
func (cc _CPlatformClass) GetRNGseed() foundation.Number {
	rv := objc.Send[foundation.Number](objc.ID(cc.class), objc.Sel("getRNGseed"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GetRNGseed) */


// Sets the global random number generator seed value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPlatform/setRNGSeedTo(_:)
func (cc _CPlatformClass) SetRNGSeedTo(seed objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("setRNGSeedTo:"), seed)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SetRNGSeedTo) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CPlatform */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CPlatform */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CPlatform */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCPlatform */



