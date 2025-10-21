// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRBridgedDeviceBasicInformationClusterProductAppearanceStruct] class.
var (
	MTRBridgedDeviceBasicInformationClusterProductAppearanceStructClass     _MTRBridgedDeviceBasicInformationClusterProductAppearanceStructClass
	MTRBridgedDeviceBasicInformationClusterProductAppearanceStructClassOnce sync.Once
)

func getMTRBridgedDeviceBasicInformationClusterProductAppearanceStructClass() _MTRBridgedDeviceBasicInformationClusterProductAppearanceStructClass {
	MTRBridgedDeviceBasicInformationClusterProductAppearanceStructClassOnce.Do(func() {
		MTRBridgedDeviceBasicInformationClusterProductAppearanceStructClass = _MTRBridgedDeviceBasicInformationClusterProductAppearanceStructClass{objc.GetClass("MTRBridgedDeviceBasicInformationClusterProductAppearanceStruct")}
	})
	return MTRBridgedDeviceBasicInformationClusterProductAppearanceStructClass
}

type _MTRBridgedDeviceBasicInformationClusterProductAppearanceStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRBridgedDeviceBasicInformationClusterProductAppearanceStruct] class.
type IMTRBridgedDeviceBasicInformationClusterProductAppearanceStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBridgedDeviceBasicInformationClusterProductAppearanceStruct
type MTRBridgedDeviceBasicInformationClusterProductAppearanceStruct struct {
	objectivec.Object
}

// MTRBridgedDeviceBasicInformationClusterProductAppearanceStructFrom constructs a [MTRBridgedDeviceBasicInformationClusterProductAppearanceStruct] from an unsafe.Pointer.
func MTRBridgedDeviceBasicInformationClusterProductAppearanceStructFrom(ptr unsafe.Pointer) MTRBridgedDeviceBasicInformationClusterProductAppearanceStruct {
	return MTRBridgedDeviceBasicInformationClusterProductAppearanceStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBridgedDeviceBasicInformationClusterProductAppearanceStructClass) Alloc() MTRBridgedDeviceBasicInformationClusterProductAppearanceStruct {
	rv := objc.Send[MTRBridgedDeviceBasicInformationClusterProductAppearanceStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBridgedDeviceBasicInformationClusterProductAppearanceStructClass) New() MTRBridgedDeviceBasicInformationClusterProductAppearanceStruct {
	rv := objc.Send[MTRBridgedDeviceBasicInformationClusterProductAppearanceStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBridgedDeviceBasicInformationClusterProductAppearanceStruct) Init() MTRBridgedDeviceBasicInformationClusterProductAppearanceStruct {
	rv := objc.Send[MTRBridgedDeviceBasicInformationClusterProductAppearanceStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBridgedDeviceBasicInformationClusterProductAppearanceStruct) Autorelease() MTRBridgedDeviceBasicInformationClusterProductAppearanceStruct {
	rv := objc.Send[MTRBridgedDeviceBasicInformationClusterProductAppearanceStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBridgedDeviceBasicInformationClusterProductAppearanceStruct creates a new MTRBridgedDeviceBasicInformationClusterProductAppearanceStruct instance.
func NewMTRBridgedDeviceBasicInformationClusterProductAppearanceStruct() MTRBridgedDeviceBasicInformationClusterProductAppearanceStruct {
	return getMTRBridgedDeviceBasicInformationClusterProductAppearanceStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbridgeddevicebasicinformationclusterproductappearancestruct/primarycolor
func (m_ MTRBridgedDeviceBasicInformationClusterProductAppearanceStruct) PrimaryColor() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("primaryColor"))
	return rv
}


// SetPrimaryColor sets the value of the primaryColor property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbridgeddevicebasicinformationclusterproductappearancestruct/primarycolor
func (m_ MTRBridgedDeviceBasicInformationClusterProductAppearanceStruct) SetPrimaryColor(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPrimaryColor:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbridgeddevicebasicinformationclusterproductappearancestruct/finish
func (m_ MTRBridgedDeviceBasicInformationClusterProductAppearanceStruct) Finish() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("finish"))
	return rv
}


// SetFinish sets the value of the finish property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbridgeddevicebasicinformationclusterproductappearancestruct/finish
func (m_ MTRBridgedDeviceBasicInformationClusterProductAppearanceStruct) SetFinish(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFinish:"), value)
}



