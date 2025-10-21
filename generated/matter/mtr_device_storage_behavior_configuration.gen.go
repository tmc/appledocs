// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRDeviceStorageBehaviorConfiguration] class.
var (
	MTRDeviceStorageBehaviorConfigurationClass     _MTRDeviceStorageBehaviorConfigurationClass
	MTRDeviceStorageBehaviorConfigurationClassOnce sync.Once
)

func getMTRDeviceStorageBehaviorConfigurationClass() _MTRDeviceStorageBehaviorConfigurationClass {
	MTRDeviceStorageBehaviorConfigurationClassOnce.Do(func() {
		MTRDeviceStorageBehaviorConfigurationClass = _MTRDeviceStorageBehaviorConfigurationClass{objc.GetClass("MTRDeviceStorageBehaviorConfiguration")}
	})
	return MTRDeviceStorageBehaviorConfigurationClass
}

type _MTRDeviceStorageBehaviorConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [MTRDeviceStorageBehaviorConfiguration] class.
type IMTRDeviceStorageBehaviorConfiguration interface {
	objectivec.IObject
}

// Class that configures how MTRDevice objects persist their attributes to storage, so as to not overwhelm the underlying storage system.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceStorageBehaviorConfiguration
type MTRDeviceStorageBehaviorConfiguration struct {
	objectivec.Object
}

// MTRDeviceStorageBehaviorConfigurationFrom constructs a [MTRDeviceStorageBehaviorConfiguration] from an unsafe.Pointer.
//
// Class that configures how MTRDevice objects persist their attributes to storage, so as to not overwhelm the underlying storage system.
func MTRDeviceStorageBehaviorConfigurationFrom(ptr unsafe.Pointer) MTRDeviceStorageBehaviorConfiguration {
	return MTRDeviceStorageBehaviorConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceStorageBehaviorConfigurationClass) Alloc() MTRDeviceStorageBehaviorConfiguration {
	rv := objc.Send[MTRDeviceStorageBehaviorConfiguration](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDeviceStorageBehaviorConfigurationClass) New() MTRDeviceStorageBehaviorConfiguration {
	rv := objc.Send[MTRDeviceStorageBehaviorConfiguration](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceStorageBehaviorConfiguration) Init() MTRDeviceStorageBehaviorConfiguration {
	rv := objc.Send[MTRDeviceStorageBehaviorConfiguration](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceStorageBehaviorConfiguration) Autorelease() MTRDeviceStorageBehaviorConfiguration {
	rv := objc.Send[MTRDeviceStorageBehaviorConfiguration](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceStorageBehaviorConfiguration creates a new MTRDeviceStorageBehaviorConfiguration instance.
func NewMTRDeviceStorageBehaviorConfiguration() MTRDeviceStorageBehaviorConfiguration {
	return getMTRDeviceStorageBehaviorConfigurationClass().New()
}




