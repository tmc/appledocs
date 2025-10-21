// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CTensorOptimizerDeviceData] class.
var (
	CTensorOptimizerDeviceDataClass     _CTensorOptimizerDeviceDataClass
	CTensorOptimizerDeviceDataClassOnce sync.Once
)

func getCTensorOptimizerDeviceDataClass() _CTensorOptimizerDeviceDataClass {
	CTensorOptimizerDeviceDataClassOnce.Do(func() {
		CTensorOptimizerDeviceDataClass = _CTensorOptimizerDeviceDataClass{objc.GetClass("MLCTensorOptimizerDeviceData")}
	})
	return CTensorOptimizerDeviceDataClass
}

type _CTensorOptimizerDeviceDataClass struct {
	class objc.Class
}

// An interface definition for the [CTensorOptimizerDeviceData] class.
type ICTensorOptimizerDeviceData interface {
	objectivec.IObject
}

// An encapsulation of the device memory associated with a tensor that an optimizer uses.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorOptimizerDeviceData
type CTensorOptimizerDeviceData struct {
	objectivec.Object
}

// CTensorOptimizerDeviceDataFrom constructs a [CTensorOptimizerDeviceData] from an unsafe.Pointer.
//
// An encapsulation of the device memory associated with a tensor that an optimizer uses.
func CTensorOptimizerDeviceDataFrom(ptr unsafe.Pointer) CTensorOptimizerDeviceData {
	return CTensorOptimizerDeviceData{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CTensorOptimizerDeviceDataClass) Alloc() CTensorOptimizerDeviceData {
	rv := objc.Send[CTensorOptimizerDeviceData](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CTensorOptimizerDeviceDataClass) New() CTensorOptimizerDeviceData {
	rv := objc.Send[CTensorOptimizerDeviceData](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CTensorOptimizerDeviceData) Init() CTensorOptimizerDeviceData {
	rv := objc.Send[CTensorOptimizerDeviceData](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CTensorOptimizerDeviceData) Autorelease() CTensorOptimizerDeviceData {
	rv := objc.Send[CTensorOptimizerDeviceData](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCTensorOptimizerDeviceData creates a new CTensorOptimizerDeviceData instance.
func NewCTensorOptimizerDeviceData() CTensorOptimizerDeviceData {
	return getCTensorOptimizerDeviceDataClass().New()
}




