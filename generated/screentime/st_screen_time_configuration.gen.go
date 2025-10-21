// Code generated from Apple documentation for ScreenTime. DO NOT EDIT.

package screentime

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [STScreenTimeConfiguration] class.
var (
	STScreenTimeConfigurationClass     _STScreenTimeConfigurationClass
	STScreenTimeConfigurationClassOnce sync.Once
)

func getSTScreenTimeConfigurationClass() _STScreenTimeConfigurationClass {
	STScreenTimeConfigurationClassOnce.Do(func() {
		STScreenTimeConfigurationClass = _STScreenTimeConfigurationClass{objc.GetClass("STScreenTimeConfiguration")}
	})
	return STScreenTimeConfigurationClass
}

type _STScreenTimeConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [STScreenTimeConfiguration] class.
type ISTScreenTimeConfiguration interface {
	objectivec.IObject
}

// The configuration for this device.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STScreenTimeConfiguration
type STScreenTimeConfiguration struct {
	objectivec.Object
}

// STScreenTimeConfigurationFrom constructs a [STScreenTimeConfiguration] from an unsafe.Pointer.
//
// The configuration for this device.
func STScreenTimeConfigurationFrom(ptr unsafe.Pointer) STScreenTimeConfiguration {
	return STScreenTimeConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _STScreenTimeConfigurationClass) Alloc() STScreenTimeConfiguration {
	rv := objc.Send[STScreenTimeConfiguration](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _STScreenTimeConfigurationClass) New() STScreenTimeConfiguration {
	rv := objc.Send[STScreenTimeConfiguration](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ STScreenTimeConfiguration) Init() STScreenTimeConfiguration {
	rv := objc.Send[STScreenTimeConfiguration](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ STScreenTimeConfiguration) Autorelease() STScreenTimeConfiguration {
	rv := objc.Send[STScreenTimeConfiguration](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSTScreenTimeConfiguration creates a new STScreenTimeConfiguration instance.
func NewSTScreenTimeConfiguration() STScreenTimeConfiguration {
	return getSTScreenTimeConfigurationClass().New()
}


// A Boolean that indicates whether the device is currently enforcing child restrictions.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STScreenTimeConfiguration/enforcesChildRestrictions
func (s_ STScreenTimeConfiguration) EnforcesChildRestrictions() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("enforcesChildRestrictions"))
	return rv
}



