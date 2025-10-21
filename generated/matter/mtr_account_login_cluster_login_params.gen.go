// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRAccountLoginClusterLoginParams] class.
var (
	MTRAccountLoginClusterLoginParamsClass     _MTRAccountLoginClusterLoginParamsClass
	MTRAccountLoginClusterLoginParamsClassOnce sync.Once
)

func getMTRAccountLoginClusterLoginParamsClass() _MTRAccountLoginClusterLoginParamsClass {
	MTRAccountLoginClusterLoginParamsClassOnce.Do(func() {
		MTRAccountLoginClusterLoginParamsClass = _MTRAccountLoginClusterLoginParamsClass{objc.GetClass("MTRAccountLoginClusterLoginParams")}
	})
	return MTRAccountLoginClusterLoginParamsClass
}

type _MTRAccountLoginClusterLoginParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRAccountLoginClusterLoginParams] class.
type IMTRAccountLoginClusterLoginParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccountLoginClusterLoginParams
type MTRAccountLoginClusterLoginParams struct {
	objectivec.Object
}

// MTRAccountLoginClusterLoginParamsFrom constructs a [MTRAccountLoginClusterLoginParams] from an unsafe.Pointer.
func MTRAccountLoginClusterLoginParamsFrom(ptr unsafe.Pointer) MTRAccountLoginClusterLoginParams {
	return MTRAccountLoginClusterLoginParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRAccountLoginClusterLoginParamsClass) Alloc() MTRAccountLoginClusterLoginParams {
	rv := objc.Send[MTRAccountLoginClusterLoginParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRAccountLoginClusterLoginParamsClass) New() MTRAccountLoginClusterLoginParams {
	rv := objc.Send[MTRAccountLoginClusterLoginParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAccountLoginClusterLoginParams) Init() MTRAccountLoginClusterLoginParams {
	rv := objc.Send[MTRAccountLoginClusterLoginParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAccountLoginClusterLoginParams) Autorelease() MTRAccountLoginClusterLoginParams {
	rv := objc.Send[MTRAccountLoginClusterLoginParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAccountLoginClusterLoginParams creates a new MTRAccountLoginClusterLoginParams instance.
func NewMTRAccountLoginClusterLoginParams() MTRAccountLoginClusterLoginParams {
	return getMTRAccountLoginClusterLoginParamsClass().New()
}




