// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRSetupPayload] class.
var (
	MTRSetupPayloadClass     _MTRSetupPayloadClass
	MTRSetupPayloadClassOnce sync.Once
)

func getMTRSetupPayloadClass() _MTRSetupPayloadClass {
	MTRSetupPayloadClassOnce.Do(func() {
		MTRSetupPayloadClass = _MTRSetupPayloadClass{objc.GetClass("MTRSetupPayload")}
	})
	return MTRSetupPayloadClass
}

type _MTRSetupPayloadClass struct {
	class objc.Class
}

// An interface definition for the [MTRSetupPayload] class.
type IMTRSetupPayload interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSetupPayload
type MTRSetupPayload struct {
	objectivec.Object
}

// MTRSetupPayloadFrom constructs a [MTRSetupPayload] from an unsafe.Pointer.
func MTRSetupPayloadFrom(ptr unsafe.Pointer) MTRSetupPayload {
	return MTRSetupPayload{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRSetupPayloadClass) Alloc() MTRSetupPayload {
	rv := objc.Send[MTRSetupPayload](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRSetupPayloadClass) New() MTRSetupPayload {
	rv := objc.Send[MTRSetupPayload](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRSetupPayload) Init() MTRSetupPayload {
	rv := objc.Send[MTRSetupPayload](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRSetupPayload) Autorelease() MTRSetupPayload {
	rv := objc.Send[MTRSetupPayload](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRSetupPayload creates a new MTRSetupPayload instance.
func NewMTRSetupPayload() MTRSetupPayload {
	return getMTRSetupPayloadClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSetupPayload/init(onboardingPayload:)
func NewMTRSetupPayloadWithOnboardingPayloadError(onboardingPayload string, error_ unsafe.Pointer) MTRSetupPayload {
	rv := objc.Send[MTRSetupPayload](objc.ID(getMTRSetupPayloadClass().class), objc.Sel("setupPayloadWithOnboardingPayload:error:"), objc.String(onboardingPayload), error_)
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSetupPayload/init(onboardingPayload:)
func (mc _MTRSetupPayloadClass) SetupPayloadWithOnboardingPayloadError(onboardingPayload string, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("setupPayloadWithOnboardingPayload:error:"), objc.String(onboardingPayload), error_)
	return rv
}


