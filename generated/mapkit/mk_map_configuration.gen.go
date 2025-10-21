// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MKMapConfiguration] class.
var (
	MKMapConfigurationClass     _MKMapConfigurationClass
	MKMapConfigurationClassOnce sync.Once
)

func getMKMapConfigurationClass() _MKMapConfigurationClass {
	MKMapConfigurationClassOnce.Do(func() {
		MKMapConfigurationClass = _MKMapConfigurationClass{objc.GetClass("MKMapConfiguration")}
	})
	return MKMapConfigurationClass
}

type _MKMapConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [MKMapConfiguration] class.
type IMKMapConfiguration interface {
	objectivec.IObject
}

// An abstract class that represents the shared elements of map configurations.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapConfiguration
type MKMapConfiguration struct {
	objectivec.Object
}

// MKMapConfigurationFrom constructs a [MKMapConfiguration] from an unsafe.Pointer.
//
// An abstract class that represents the shared elements of map configurations.
func MKMapConfigurationFrom(ptr unsafe.Pointer) MKMapConfiguration {
	return MKMapConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MKMapConfigurationClass) Alloc() MKMapConfiguration {
	rv := objc.Send[MKMapConfiguration](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKMapConfigurationClass) New() MKMapConfiguration {
	rv := objc.Send[MKMapConfiguration](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKMapConfiguration) Init() MKMapConfiguration {
	rv := objc.Send[MKMapConfiguration](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKMapConfiguration) Autorelease() MKMapConfiguration {
	rv := objc.Send[MKMapConfiguration](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKMapConfiguration creates a new MKMapConfiguration instance.
func NewMKMapConfiguration() MKMapConfiguration {
	return getMKMapConfigurationClass().New()
}




