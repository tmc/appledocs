// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [OverlayConfiguration] class.
var (
	OverlayConfigurationClass     _OverlayConfigurationClass
	OverlayConfigurationClassOnce sync.Once
)

func getOverlayConfigurationClass() _OverlayConfigurationClass {
	OverlayConfigurationClassOnce.Do(func() {
		OverlayConfigurationClass = _OverlayConfigurationClass{objc.GetClass("SKOverlayConfiguration")}
	})
	return OverlayConfigurationClass
}

type _OverlayConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [OverlayConfiguration] class.
type IOverlayConfiguration interface {
	objectivec.IObject
}

// The abstract superclass for all classes that represent an overlay’s attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/Configuration-swift.class
type OverlayConfiguration struct {
	objectivec.Object
}

// OverlayConfigurationFrom constructs a [OverlayConfiguration] from an unsafe.Pointer.
//
// The abstract superclass for all classes that represent an overlay’s attributes.
func OverlayConfigurationFrom(ptr unsafe.Pointer) OverlayConfiguration {
	return OverlayConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (oc _OverlayConfigurationClass) Alloc() OverlayConfiguration {
	rv := objc.Send[OverlayConfiguration](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _OverlayConfigurationClass) New() OverlayConfiguration {
	rv := objc.Send[OverlayConfiguration](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OverlayConfiguration) Init() OverlayConfiguration {
	rv := objc.Send[OverlayConfiguration](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OverlayConfiguration) Autorelease() OverlayConfiguration {
	rv := objc.Send[OverlayConfiguration](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOverlayConfiguration creates a new OverlayConfiguration instance.
func NewOverlayConfiguration() OverlayConfiguration {
	return getOverlayConfigurationClass().New()
}




