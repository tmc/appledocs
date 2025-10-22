// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRMediaInputClusterRenameInputParams] class.
var (
	MTRMediaInputClusterRenameInputParamsClass     _MTRMediaInputClusterRenameInputParamsClass
	MTRMediaInputClusterRenameInputParamsClassOnce sync.Once
)

func getMTRMediaInputClusterRenameInputParamsClass() _MTRMediaInputClusterRenameInputParamsClass {
	MTRMediaInputClusterRenameInputParamsClassOnce.Do(func() {
		MTRMediaInputClusterRenameInputParamsClass = _MTRMediaInputClusterRenameInputParamsClass{objc.GetClass("MTRMediaInputClusterRenameInputParams")}
	})
	return MTRMediaInputClusterRenameInputParamsClass
}

type _MTRMediaInputClusterRenameInputParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRMediaInputClusterRenameInputParams] class.
type IMTRMediaInputClusterRenameInputParams interface {
	objectivec.IObject
	Index() foundation.Number
	SetIndex(value foundation.INumber)
	Name() string
	SetName(value string)
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaInputClusterRenameInputParams
type MTRMediaInputClusterRenameInputParams struct {
	objectivec.Object
}

// MTRMediaInputClusterRenameInputParamsFrom constructs a [MTRMediaInputClusterRenameInputParams] from an unsafe.Pointer.
func MTRMediaInputClusterRenameInputParamsFrom(ptr unsafe.Pointer) MTRMediaInputClusterRenameInputParams {
	return MTRMediaInputClusterRenameInputParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRMediaInputClusterRenameInputParamsClass) Alloc() MTRMediaInputClusterRenameInputParams {
	rv := objc.Send[MTRMediaInputClusterRenameInputParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRMediaInputClusterRenameInputParamsClass) New() MTRMediaInputClusterRenameInputParams {
	rv := objc.Send[MTRMediaInputClusterRenameInputParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMediaInputClusterRenameInputParams) Init() MTRMediaInputClusterRenameInputParams {
	rv := objc.Send[MTRMediaInputClusterRenameInputParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMediaInputClusterRenameInputParams) Autorelease() MTRMediaInputClusterRenameInputParams {
	rv := objc.Send[MTRMediaInputClusterRenameInputParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMediaInputClusterRenameInputParams creates a new MTRMediaInputClusterRenameInputParams instance.
func NewMTRMediaInputClusterRenameInputParams() MTRMediaInputClusterRenameInputParams {
	return getMTRMediaInputClusterRenameInputParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediainputclusterrenameinputparams/index
func (m_ MTRMediaInputClusterRenameInputParams) Index() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("index"))
	return rv
}


// SetIndex sets the value of the index property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediainputclusterrenameinputparams/index
func (m_ MTRMediaInputClusterRenameInputParams) SetIndex(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIndex:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediainputclusterrenameinputparams/name
func (m_ MTRMediaInputClusterRenameInputParams) Name() string {
	rv := objc.Send[string](m_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediainputclusterrenameinputparams/name
func (m_ MTRMediaInputClusterRenameInputParams) SetName(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediainputclusterrenameinputparams/serversideprocessingtimeout
func (m_ MTRMediaInputClusterRenameInputParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediainputclusterrenameinputparams/serversideprocessingtimeout
func (m_ MTRMediaInputClusterRenameInputParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediainputclusterrenameinputparams/timedinvoketimeoutms
func (m_ MTRMediaInputClusterRenameInputParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediainputclusterrenameinputparams/timedinvoketimeoutms
func (m_ MTRMediaInputClusterRenameInputParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



