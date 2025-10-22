// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CKSyncEngineConfiguration] class.
var (
	CKSyncEngineConfigurationClass     _CKSyncEngineConfigurationClass
	CKSyncEngineConfigurationClassOnce sync.Once
)

func getCKSyncEngineConfigurationClass() _CKSyncEngineConfigurationClass {
	CKSyncEngineConfigurationClassOnce.Do(func() {
		CKSyncEngineConfigurationClass = _CKSyncEngineConfigurationClass{objc.GetClass("CKSyncEngineConfiguration")}
	})
	return CKSyncEngineConfigurationClass
}

type _CKSyncEngineConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [CKSyncEngineConfiguration] class.
type ICKSyncEngineConfiguration interface {
	objectivec.IObject
}

// A type that configures the attributes and behavior of a sync engine.


// A type that configures the attributes and behavior of a sync engine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineConfiguration

type CKSyncEngineConfiguration struct {
	objectivec.Object
}

// CKSyncEngineConfigurationFrom constructs a [CKSyncEngineConfiguration] from an unsafe.Pointer.
//
// A type that configures the attributes and behavior of a sync engine.
func CKSyncEngineConfigurationFrom(ptr unsafe.Pointer) CKSyncEngineConfiguration {
	return CKSyncEngineConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineConfigurationClass) Alloc() CKSyncEngineConfiguration {
	rv := objc.Send[CKSyncEngineConfiguration](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKSyncEngineConfigurationClass) New() CKSyncEngineConfiguration {
	rv := objc.Send[CKSyncEngineConfiguration](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineConfiguration) Init() CKSyncEngineConfiguration {
	rv := objc.Send[CKSyncEngineConfiguration](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineConfiguration) Autorelease() CKSyncEngineConfiguration {
	rv := objc.Send[CKSyncEngineConfiguration](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineConfiguration creates a new CKSyncEngineConfiguration instance.
func NewCKSyncEngineConfiguration() CKSyncEngineConfiguration {
	return getCKSyncEngineConfigurationClass().New()
}




