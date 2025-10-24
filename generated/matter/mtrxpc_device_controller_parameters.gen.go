// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRXPCDeviceControllerParameters] class.
var (
	MTRXPCDeviceControllerParametersClass     _MTRXPCDeviceControllerParametersClass
	MTRXPCDeviceControllerParametersClassOnce sync.Once
)

func getMTRXPCDeviceControllerParametersClass() _MTRXPCDeviceControllerParametersClass {
	MTRXPCDeviceControllerParametersClassOnce.Do(func() {
		MTRXPCDeviceControllerParametersClass = _MTRXPCDeviceControllerParametersClass{objc.GetClass("MTRXPCDeviceControllerParameters")}
	})
	return MTRXPCDeviceControllerParametersClass
}

type _MTRXPCDeviceControllerParametersClass struct {
	class objc.Class
}

// An interface definition for the [MTRXPCDeviceControllerParameters] class.
type IMTRXPCDeviceControllerParameters interface {
	IMTRDeviceControllerAbstractParameters
	// properties:
	UniqueIdentifier() objc.IObject /* cross-framework: UUID */
	XpcConnectionBlock() NSXPCConnection * (^ /* not a class type */
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRXPCDeviceControllerParameters
type MTRXPCDeviceControllerParameters struct {
	MTRDeviceControllerAbstractParameters
}

// MTRXPCDeviceControllerParametersFrom constructs a [MTRXPCDeviceControllerParameters] from an unsafe.Pointer.
func MTRXPCDeviceControllerParametersFrom(ptr unsafe.Pointer) MTRXPCDeviceControllerParameters {
	return MTRXPCDeviceControllerParameters{
		MTRDeviceControllerAbstractParameters: MTRDeviceControllerAbstractParametersFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRXPCDeviceControllerParametersClass) Alloc() MTRXPCDeviceControllerParameters {
	rv := objc.Send[MTRXPCDeviceControllerParameters](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRXPCDeviceControllerParametersClass) New() MTRXPCDeviceControllerParameters {
	rv := objc.Send[MTRXPCDeviceControllerParameters](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRXPCDeviceControllerParameters) Init() MTRXPCDeviceControllerParameters {
	rv := objc.Send[MTRXPCDeviceControllerParameters](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRXPCDeviceControllerParameters) Autorelease() MTRXPCDeviceControllerParameters {
	rv := objc.Send[MTRXPCDeviceControllerParameters](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRXPCDeviceControllerParameters creates a new MTRXPCDeviceControllerParameters instance.
func NewMTRXPCDeviceControllerParameters() MTRXPCDeviceControllerParameters {
	return getMTRXPCDeviceControllerParametersClass().New()
}



// A controller created from this way will connect to a remote instance of an MTRDeviceController loaded in an XPC Service
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRXPCDeviceControllerParameters/init(xpcConnectionBlock:uniqueIdentifier:)
func NewMTRXPCDeviceControllerParametersWithXPCConnectionBlockUniqueIdentifier(xpcConnectionBlock XPCConnection * (^ /* not a class type */, uniqueIdentifier objc.IObject /* cross-framework: UUID */) MTRXPCDeviceControllerParameters {
	instance := getMTRXPCDeviceControllerParametersClass().Alloc()
	rv := objc.Send[MTRXPCDeviceControllerParameters](instance.ID, objc.Sel("initWithXPCConnectionBlock:uniqueIdentifier:"), xpcConnectionBlock, uniqueIdentifier)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRXPCDeviceControllerParameters/init(xpConnectionBlock:uniqueIdentifier:)
func NewMTRXPCDeviceControllerParametersWithXPConnectionBlockUniqueIdentifier(xpcConnectionBlock XPCConnection * (^ /* not a class type */, uniqueIdentifier objc.IObject /* cross-framework: UUID */) MTRXPCDeviceControllerParameters {
	instance := getMTRXPCDeviceControllerParametersClass().Alloc()
	rv := objc.Send[MTRXPCDeviceControllerParameters](instance.ID, objc.Sel("initWithXPConnectionBlock:uniqueIdentifier:"), xpcConnectionBlock, uniqueIdentifier)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRXPCDeviceControllerParameters/uniqueIdentifier
func (m_ MTRXPCDeviceControllerParameters) UniqueIdentifier() objc.IObject /* cross-framework: UUID */ {
	rv := objc.Send[foundation.UUID](m_.ID, objc.Sel("uniqueIdentifier"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRXPCDeviceControllerParameters/xpcConnectionBlock
func (m_ MTRXPCDeviceControllerParameters) XpcConnectionBlock() NSXPCConnection * (^ /* not a class type */ {
	rv := objc.Send[XPCConnection * (^](m_.ID, objc.Sel("xpcConnectionBlock"))
	return rv
}


