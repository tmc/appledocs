// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRValveConfigurationAndControlClusterOpenParams] class.
var (
	MTRValveConfigurationAndControlClusterOpenParamsClass     _MTRValveConfigurationAndControlClusterOpenParamsClass
	MTRValveConfigurationAndControlClusterOpenParamsClassOnce sync.Once
)

func getMTRValveConfigurationAndControlClusterOpenParamsClass() _MTRValveConfigurationAndControlClusterOpenParamsClass {
	MTRValveConfigurationAndControlClusterOpenParamsClassOnce.Do(func() {
		MTRValveConfigurationAndControlClusterOpenParamsClass = _MTRValveConfigurationAndControlClusterOpenParamsClass{objc.GetClass("MTRValveConfigurationAndControlClusterOpenParams")}
	})
	return MTRValveConfigurationAndControlClusterOpenParamsClass
}

type _MTRValveConfigurationAndControlClusterOpenParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRValveConfigurationAndControlClusterOpenParams] class.
type IMTRValveConfigurationAndControlClusterOpenParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRValveConfigurationAndControlClusterOpenParams
type MTRValveConfigurationAndControlClusterOpenParams struct {
	objectivec.Object
}

// MTRValveConfigurationAndControlClusterOpenParamsFrom constructs a [MTRValveConfigurationAndControlClusterOpenParams] from an unsafe.Pointer.
func MTRValveConfigurationAndControlClusterOpenParamsFrom(ptr unsafe.Pointer) MTRValveConfigurationAndControlClusterOpenParams {
	return MTRValveConfigurationAndControlClusterOpenParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRValveConfigurationAndControlClusterOpenParamsClass) Alloc() MTRValveConfigurationAndControlClusterOpenParams {
	rv := objc.Send[MTRValveConfigurationAndControlClusterOpenParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRValveConfigurationAndControlClusterOpenParamsClass) New() MTRValveConfigurationAndControlClusterOpenParams {
	rv := objc.Send[MTRValveConfigurationAndControlClusterOpenParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRValveConfigurationAndControlClusterOpenParams) Init() MTRValveConfigurationAndControlClusterOpenParams {
	rv := objc.Send[MTRValveConfigurationAndControlClusterOpenParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRValveConfigurationAndControlClusterOpenParams) Autorelease() MTRValveConfigurationAndControlClusterOpenParams {
	rv := objc.Send[MTRValveConfigurationAndControlClusterOpenParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRValveConfigurationAndControlClusterOpenParams creates a new MTRValveConfigurationAndControlClusterOpenParams instance.
func NewMTRValveConfigurationAndControlClusterOpenParams() MTRValveConfigurationAndControlClusterOpenParams {
	return getMTRValveConfigurationAndControlClusterOpenParamsClass().New()
}




