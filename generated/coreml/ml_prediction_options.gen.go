// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [PredictionOptions] class.
type IPredictionOptions interface {
	objectivec.IObject
}

// The options available when making a prediction.
//
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

// Alloc allocates a new instance without initialization.
func (pc _PredictionOptionsClass) Alloc() PredictionOptions {
	rv := objc.Send[PredictionOptions](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




