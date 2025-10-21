// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRDeviceController] class.
var (
	MTRDeviceControllerClass     _MTRDeviceControllerClass
	MTRDeviceControllerClassOnce sync.Once
)

func getMTRDeviceControllerClass() _MTRDeviceControllerClass {
	MTRDeviceControllerClassOnce.Do(func() {
		MTRDeviceControllerClass = _MTRDeviceControllerClass{objc.GetClass("MTRDeviceController")}
	})
	return MTRDeviceControllerClass
}

type _MTRDeviceControllerClass struct {
	class objc.Class
}

// An interface definition for the [MTRDeviceController] class.
type IMTRDeviceController interface {
	objectivec.IObject
	CommissionNodeWithIDCommissioningParamsError(nodeID unsafe.Pointer, commissioningParams unsafe.Pointer, error_ unsafe.Pointer) bool
	SetDeviceControllerDelegateQueue(delegate objc.ID, queue unsafe.Pointer)
	SetupCommissioningSessionWithPayloadNewNodeIDError(payload unsafe.Pointer, newNodeID unsafe.Pointer, error_ unsafe.Pointer) bool
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController
type MTRDeviceController struct {
	objectivec.Object
}

// MTRDeviceControllerFrom constructs a [MTRDeviceController] from an unsafe.Pointer.
func MTRDeviceControllerFrom(ptr unsafe.Pointer) MTRDeviceController {
	return MTRDeviceController{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceControllerClass) Alloc() MTRDeviceController {
	rv := objc.Send[MTRDeviceController](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDeviceControllerClass) New() MTRDeviceController {
	rv := objc.Send[MTRDeviceController](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceController) Init() MTRDeviceController {
	rv := objc.Send[MTRDeviceController](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceController) Autorelease() MTRDeviceController {
	rv := objc.Send[MTRDeviceController](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceController creates a new MTRDeviceController instance.
func NewMTRDeviceController() MTRDeviceController {
	return getMTRDeviceControllerClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/commissionNode(withID:commissioningParams:)
func (m_ MTRDeviceController) CommissionNodeWithIDCommissioningParamsError(nodeID unsafe.Pointer, commissioningParams unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("commissionNodeWithID:commissioningParams:error:"), nodeID, commissioningParams, error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/setDeviceControllerDelegate(_:queue:)
func (m_ MTRDeviceController) SetDeviceControllerDelegateQueue(delegate objc.ID, queue unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDeviceControllerDelegate:queue:"), delegate, queue)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/setupCommissioningSession(with:newNodeID:)
func (m_ MTRDeviceController) SetupCommissioningSessionWithPayloadNewNodeIDError(payload unsafe.Pointer, newNodeID unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("setupCommissioningSessionWithPayload:newNodeID:error:"), payload, newNodeID, error_)
	return rv
}



