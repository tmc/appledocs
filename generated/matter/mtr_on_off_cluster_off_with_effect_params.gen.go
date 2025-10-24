// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROnOffClusterOffWithEffectParams] class.
var (
	MTROnOffClusterOffWithEffectParamsClass     _MTROnOffClusterOffWithEffectParamsClass
	MTROnOffClusterOffWithEffectParamsClassOnce sync.Once
)

func getMTROnOffClusterOffWithEffectParamsClass() _MTROnOffClusterOffWithEffectParamsClass {
	MTROnOffClusterOffWithEffectParamsClassOnce.Do(func() {
		MTROnOffClusterOffWithEffectParamsClass = _MTROnOffClusterOffWithEffectParamsClass{objc.GetClass("MTROnOffClusterOffWithEffectParams")}
	})
	return MTROnOffClusterOffWithEffectParamsClass
}

type _MTROnOffClusterOffWithEffectParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROnOffClusterOffWithEffectParams] class.
type IMTROnOffClusterOffWithEffectParams interface {
	objectivec.IObject
	// properties:
	EffectId() objc.IObject /* cross-framework: NSNumber */
	SetEffectId(value objc.IObject /* cross-framework: NSNumber */)
	EffectIdentifier() objc.IObject /* cross-framework: NSNumber */
	SetEffectIdentifier(value objc.IObject /* cross-framework: NSNumber */)
	EffectVariant() objc.IObject /* cross-framework: NSNumber */
	SetEffectVariant(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterOffWithEffectParams
type MTROnOffClusterOffWithEffectParams struct {
	objectivec.Object
}

// MTROnOffClusterOffWithEffectParamsFrom constructs a [MTROnOffClusterOffWithEffectParams] from an unsafe.Pointer.
func MTROnOffClusterOffWithEffectParamsFrom(ptr unsafe.Pointer) MTROnOffClusterOffWithEffectParams {
	return MTROnOffClusterOffWithEffectParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROnOffClusterOffWithEffectParamsClass) Alloc() MTROnOffClusterOffWithEffectParams {
	rv := objc.Send[MTROnOffClusterOffWithEffectParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROnOffClusterOffWithEffectParamsClass) New() MTROnOffClusterOffWithEffectParams {
	rv := objc.Send[MTROnOffClusterOffWithEffectParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROnOffClusterOffWithEffectParams) Init() MTROnOffClusterOffWithEffectParams {
	rv := objc.Send[MTROnOffClusterOffWithEffectParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROnOffClusterOffWithEffectParams) Autorelease() MTROnOffClusterOffWithEffectParams {
	rv := objc.Send[MTROnOffClusterOffWithEffectParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROnOffClusterOffWithEffectParams creates a new MTROnOffClusterOffWithEffectParams instance.
func NewMTROnOffClusterOffWithEffectParams() MTROnOffClusterOffWithEffectParams {
	return getMTROnOffClusterOffWithEffectParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtronoffclusteroffwitheffectparams/effectid
func (m_ MTROnOffClusterOffWithEffectParams) EffectId() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("effectId"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtronoffclusteroffwitheffectparams/effectid
func (m_ MTROnOffClusterOffWithEffectParams) SetEffectId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEffectId:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtronoffclusteroffwitheffectparams/effectidentifier
func (m_ MTROnOffClusterOffWithEffectParams) EffectIdentifier() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("effectIdentifier"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtronoffclusteroffwitheffectparams/effectidentifier
func (m_ MTROnOffClusterOffWithEffectParams) SetEffectIdentifier(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEffectIdentifier:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtronoffclusteroffwitheffectparams/effectvariant
func (m_ MTROnOffClusterOffWithEffectParams) EffectVariant() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("effectVariant"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtronoffclusteroffwitheffectparams/effectvariant
func (m_ MTROnOffClusterOffWithEffectParams) SetEffectVariant(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEffectVariant:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtronoffclusteroffwitheffectparams/serversideprocessingtimeout
func (m_ MTROnOffClusterOffWithEffectParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtronoffclusteroffwitheffectparams/serversideprocessingtimeout
func (m_ MTROnOffClusterOffWithEffectParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtronoffclusteroffwitheffectparams/timedinvoketimeoutms
func (m_ MTROnOffClusterOffWithEffectParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtronoffclusteroffwitheffectparams/timedinvoketimeoutms
func (m_ MTROnOffClusterOffWithEffectParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



