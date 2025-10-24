// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLPredictionOptions */


/* debug [class_header]: Header for MLPredictionOptions */
// The class instance for the [PredictionOptions] class.
var (
	PredictionOptionsClass     _PredictionOptionsClass
	PredictionOptionsClassOnce sync.Once
)

func getPredictionOptionsClass() _PredictionOptionsClass {
	PredictionOptionsClassOnce.Do(func() {
		PredictionOptionsClass = _PredictionOptionsClass{objc.GetClass("MLPredictionOptions")}
	})
	return PredictionOptionsClass
}

type _PredictionOptionsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PredictionOptions */
// An interface definition for the [PredictionOptions] class.
type IPredictionOptions interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PredictionOptions */
	// properties:
	OutputBackings() foundation.IDictionary
	SetOutputBackings(value foundation.IDictionary)
	UsesCPUOnly() bool
	SetUsesCPUOnly(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PredictionOptions */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PredictionOptions */
// Alloc allocates a new instance without initialization.
func (pc _PredictionOptionsClass) Alloc() PredictionOptions {
	rv := objc.Send[PredictionOptions](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PredictionOptionsClass) New() PredictionOptions {
	rv := objc.Send[PredictionOptions](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PredictionOptions) Init() PredictionOptions {
	rv := objc.Send[PredictionOptions](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PredictionOptions) Autorelease() PredictionOptions {
	rv := objc.Send[PredictionOptions](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPredictionOptions creates a new PredictionOptions instance.
func NewPredictionOptions() PredictionOptions {
	return getPredictionOptionsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PredictionOptions */
// The options available when making a prediction.


// The options available when making a prediction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLPredictionOptions
type PredictionOptions struct {
	objectivec.Object
}

// PredictionOptionsFrom constructs a [PredictionOptions] from an unsafe.Pointer.
//
// The options available when making a prediction.
func PredictionOptionsFrom(ptr unsafe.Pointer) PredictionOptions {
	return PredictionOptions{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PredictionOptions *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PredictionOptions */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PredictionOptions */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PredictionOptions */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PredictionOptions */

// A dictionary of feature names and client-allocated buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLPredictionOptions/outputBackings
func (p_ PredictionOptions) OutputBackings() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](p_.ID, objc.Sel("outputBackings"))
	return rv
}/* debug [instance_properties/getter]: outputBackings */


// A dictionary of feature names and client-allocated buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLPredictionOptions/outputBackings
func (p_ PredictionOptions) SetOutputBackings(value foundation.IDictionary) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setOutputBackings:"), value)
}/* debug [instance_properties/setter]: outputBackings */


// A Boolean value that indicates whether a prediction is computed using only the CPU.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLPredictionOptions/usesCPUOnly
func (p_ PredictionOptions) UsesCPUOnly() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("usesCPUOnly"))
	return rv
}/* debug [instance_properties/getter]: usesCPUOnly */


// A Boolean value that indicates whether a prediction is computed using only the CPU.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLPredictionOptions/usesCPUOnly
func (p_ PredictionOptions) SetUsesCPUOnly(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUsesCPUOnly:"), value)
}/* debug [instance_properties/setter]: usesCPUOnly */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLPredictionOptions */



