// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CPlatform] class.
var (
	CPlatformClass     _CPlatformClass
	CPlatformClassOnce sync.Once
)

func getCPlatformClass() _CPlatformClass {
	CPlatformClassOnce.Do(func() {
		CPlatformClass = _CPlatformClass{objc.GetClass("MLCPlatform")}
	})
	return CPlatformClass
}

type _CPlatformClass struct {
	class objc.Class
}

// An interface definition for the [CPlatform] class.
type ICPlatform interface {
	objectivec.IObject
}

// A utility class for setting global properties in the framework.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPlatform
type CPlatform struct {
	objectivec.Object
}

// CPlatformFrom constructs a [CPlatform] from an unsafe.Pointer.
//
// A utility class for setting global properties in the framework.
func CPlatformFrom(ptr unsafe.Pointer) CPlatform {
	return CPlatform{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CPlatformClass) Alloc() CPlatform {
	rv := objc.Send[CPlatform](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CPlatformClass) New() CPlatform {
	rv := objc.Send[CPlatform](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CPlatform) Init() CPlatform {
	rv := objc.Send[CPlatform](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CPlatform) Autorelease() CPlatform {
	rv := objc.Send[CPlatform](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPlatform creates a new CPlatform instance.
func NewCPlatform() CPlatform {
	return getCPlatformClass().New()
}


// Returns the global random number generator seed value.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPlatform/getRNGseed()
func (cc _CPlatformClass) GetRNGseed() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("getRNGseed"))
	return rv
}

// Sets the global random number generator seed value.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPlatform/setRNGSeedTo(_:)
func (cc _CPlatformClass) SetRNGSeedTo(seed unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("setRNGSeedTo:"), seed)
}



