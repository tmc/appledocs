// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [OptimizationHints] class.
type IOptimizationHints interface {
	objectivec.IObject
	ReshapeFrequency() ReshapeFrequencyHint
	SetReshapeFrequency(value IReshapeFrequencyHint)
	SpecializationStrategy() SpecializationStrategy
	SetSpecializationStrategy(value SpecializationStrategy)
}

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

// Alloc allocates a new instance without initialization.
func (oc _OptimizationHintsClass) Alloc() OptimizationHints {
	rv := objc.Send[OptimizationHints](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The anticipated reshape frequency
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLOptimizationHints-c.class/reshapeFrequency
func (o_ OptimizationHints) ReshapeFrequency() ReshapeFrequencyHint {
	rv := objc.Send[ReshapeFrequencyHint](o_.ID, objc.Sel("reshapeFrequency"))
	return rv
}


// The anticipated reshape frequency
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLOptimizationHints-c.class/reshapeFrequency
func (o_ OptimizationHints) SetReshapeFrequency(value IReshapeFrequencyHint) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setReshapeFrequency:"), value)
}


// Optimization strategy for the model specialization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLOptimizationHints-c.class/specializationStrategy
func (o_ OptimizationHints) SpecializationStrategy() SpecializationStrategy {
	rv := objc.Send[SpecializationStrategy](o_.ID, objc.Sel("specializationStrategy"))
	return rv
}


// Optimization strategy for the model specialization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLOptimizationHints-c.class/specializationStrategy
func (o_ OptimizationHints) SetSpecializationStrategy(value SpecializationStrategy) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setSpecializationStrategy:"), value)
}



