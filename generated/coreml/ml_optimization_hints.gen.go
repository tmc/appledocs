// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLOptimizationHints */


/* debug [class_header]: Header for MLOptimizationHints */
// The class instance for the [OptimizationHints] class.
var (
	OptimizationHintsClass     _OptimizationHintsClass
	OptimizationHintsClassOnce sync.Once
)

func getOptimizationHintsClass() _OptimizationHintsClass {
	OptimizationHintsClassOnce.Do(func() {
		OptimizationHintsClass = _OptimizationHintsClass{objc.GetClass("MLOptimizationHints")}
	})
	return OptimizationHintsClass
}

type _OptimizationHintsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for OptimizationHints */
// An interface definition for the [OptimizationHints] class.
type IOptimizationHints interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for OptimizationHints */
	// properties:
	ReshapeFrequency() ReshapeFrequencyHint
	SetReshapeFrequency(value ReshapeFrequencyHint)
	SpecializationStrategy() SpecializationStrategy
	SetSpecializationStrategy(value SpecializationStrategy)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for OptimizationHints */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for OptimizationHints */
// Alloc allocates a new instance without initialization.
func (oc _OptimizationHintsClass) Alloc() OptimizationHints {
	rv := objc.Send[OptimizationHints](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (oc _OptimizationHintsClass) New() OptimizationHints {
	rv := objc.Send[OptimizationHints](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OptimizationHints) Init() OptimizationHints {
	rv := objc.Send[OptimizationHints](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OptimizationHints) Autorelease() OptimizationHints {
	rv := objc.Send[OptimizationHints](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOptimizationHints creates a new OptimizationHints instance.
func NewOptimizationHints() OptimizationHints {
	return getOptimizationHintsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for OptimizationHints */
// MLOptimizationHints
//
// An object to hold hints that CoreML could use for further optimization


// MLOptimizationHints
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLOptimizationHints-c.class
type OptimizationHints struct {
	objectivec.Object
}

// OptimizationHintsFrom constructs a [OptimizationHints] from an unsafe.Pointer.
//
// MLOptimizationHints
func OptimizationHintsFrom(ptr unsafe.Pointer) OptimizationHints {
	return OptimizationHints{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for OptimizationHints *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for OptimizationHints */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for OptimizationHints */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for OptimizationHints */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for OptimizationHints */

// The anticipated reshape frequency
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLOptimizationHints-c.class/reshapeFrequency
func (o_ OptimizationHints) ReshapeFrequency() ReshapeFrequencyHint {
	rv := objc.Send[ReshapeFrequencyHint](o_.ID, objc.Sel("reshapeFrequency"))
	return rv
}/* debug [instance_properties/getter]: reshapeFrequency */


// The anticipated reshape frequency
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLOptimizationHints-c.class/reshapeFrequency
func (o_ OptimizationHints) SetReshapeFrequency(value ReshapeFrequencyHint) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setReshapeFrequency:"), value)
}/* debug [instance_properties/setter]: reshapeFrequency */


// Optimization strategy for the model specialization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLOptimizationHints-c.class/specializationStrategy
func (o_ OptimizationHints) SpecializationStrategy() SpecializationStrategy {
	rv := objc.Send[SpecializationStrategy](o_.ID, objc.Sel("specializationStrategy"))
	return rv
}/* debug [instance_properties/getter]: specializationStrategy */


// Optimization strategy for the model specialization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLOptimizationHints-c.class/specializationStrategy
func (o_ OptimizationHints) SetSpecializationStrategy(value SpecializationStrategy) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setSpecializationStrategy:"), value)
}/* debug [instance_properties/setter]: specializationStrategy */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLOptimizationHints */



