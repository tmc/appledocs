// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROnboardingPayloadParser] class.
var (
	MTROnboardingPayloadParserClass     _MTROnboardingPayloadParserClass
	MTROnboardingPayloadParserClassOnce sync.Once
)

func getMTROnboardingPayloadParserClass() _MTROnboardingPayloadParserClass {
	MTROnboardingPayloadParserClassOnce.Do(func() {
		MTROnboardingPayloadParserClass = _MTROnboardingPayloadParserClass{objc.GetClass("MTROnboardingPayloadParser")}
	})
	return MTROnboardingPayloadParserClass
}

type _MTROnboardingPayloadParserClass struct {
	class objc.Class
}

// An interface definition for the [MTROnboardingPayloadParser] class.
type IMTROnboardingPayloadParser interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnboardingPayloadParser
type MTROnboardingPayloadParser struct {
	objectivec.Object
}

// MTROnboardingPayloadParserFrom constructs a [MTROnboardingPayloadParser] from an unsafe.Pointer.
func MTROnboardingPayloadParserFrom(ptr unsafe.Pointer) MTROnboardingPayloadParser {
	return MTROnboardingPayloadParser{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROnboardingPayloadParserClass) Alloc() MTROnboardingPayloadParser {
	rv := objc.Send[MTROnboardingPayloadParser](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROnboardingPayloadParserClass) New() MTROnboardingPayloadParser {
	rv := objc.Send[MTROnboardingPayloadParser](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROnboardingPayloadParser) Init() MTROnboardingPayloadParser {
	rv := objc.Send[MTROnboardingPayloadParser](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROnboardingPayloadParser) Autorelease() MTROnboardingPayloadParser {
	rv := objc.Send[MTROnboardingPayloadParser](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROnboardingPayloadParser creates a new MTROnboardingPayloadParser instance.
func NewMTROnboardingPayloadParser() MTROnboardingPayloadParser {
	return getMTROnboardingPayloadParserClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnboardingPayloadParser/setupPayload(forOnboardingPayload:)
func (mc _MTROnboardingPayloadParserClass) SetupPayloadForOnboardingPayloadError(onboardingPayload string, error_ unsafe.Pointer) MTRSetupPayload {
	rv := objc.Send[MTRSetupPayload](objc.ID(mc.class), objc.Sel("setupPayloadForOnboardingPayload:error:"), objc.String(onboardingPayload), error_)
	return rv
}



