// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRChannelClusterProgramGuideResponseParams] class.
var (
	MTRChannelClusterProgramGuideResponseParamsClass     _MTRChannelClusterProgramGuideResponseParamsClass
	MTRChannelClusterProgramGuideResponseParamsClassOnce sync.Once
)

func getMTRChannelClusterProgramGuideResponseParamsClass() _MTRChannelClusterProgramGuideResponseParamsClass {
	MTRChannelClusterProgramGuideResponseParamsClassOnce.Do(func() {
		MTRChannelClusterProgramGuideResponseParamsClass = _MTRChannelClusterProgramGuideResponseParamsClass{objc.GetClass("MTRChannelClusterProgramGuideResponseParams")}
	})
	return MTRChannelClusterProgramGuideResponseParamsClass
}

type _MTRChannelClusterProgramGuideResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRChannelClusterProgramGuideResponseParams] class.
type IMTRChannelClusterProgramGuideResponseParams interface {
	objectivec.IObject
	// properties:
	Paging() IMTRChannelClusterChannelPagingStruct
	SetPaging(value IMTRChannelClusterChannelPagingStruct)
	ProgramList() objc.IObject /* cross-framework: NSArray */
	SetProgramList(value objc.IObject /* cross-framework: NSArray */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramGuideResponseParams
type MTRChannelClusterProgramGuideResponseParams struct {
	objectivec.Object
}

// MTRChannelClusterProgramGuideResponseParamsFrom constructs a [MTRChannelClusterProgramGuideResponseParams] from an unsafe.Pointer.
func MTRChannelClusterProgramGuideResponseParamsFrom(ptr unsafe.Pointer) MTRChannelClusterProgramGuideResponseParams {
	return MTRChannelClusterProgramGuideResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRChannelClusterProgramGuideResponseParamsClass) Alloc() MTRChannelClusterProgramGuideResponseParams {
	rv := objc.Send[MTRChannelClusterProgramGuideResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRChannelClusterProgramGuideResponseParamsClass) New() MTRChannelClusterProgramGuideResponseParams {
	rv := objc.Send[MTRChannelClusterProgramGuideResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRChannelClusterProgramGuideResponseParams) Init() MTRChannelClusterProgramGuideResponseParams {
	rv := objc.Send[MTRChannelClusterProgramGuideResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRChannelClusterProgramGuideResponseParams) Autorelease() MTRChannelClusterProgramGuideResponseParams {
	rv := objc.Send[MTRChannelClusterProgramGuideResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRChannelClusterProgramGuideResponseParams creates a new MTRChannelClusterProgramGuideResponseParams instance.
func NewMTRChannelClusterProgramGuideResponseParams() MTRChannelClusterProgramGuideResponseParams {
	return getMTRChannelClusterProgramGuideResponseParamsClass().New()
}



// Initialize an MTRChannelClusterProgramGuideResponseParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramGuideResponseParams/init(responseValue:)
func NewMTRChannelClusterProgramGuideResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTRChannelClusterProgramGuideResponseParams {
	instance := getMTRChannelClusterProgramGuideResponseParamsClass().Alloc()
	rv := objc.Send[MTRChannelClusterProgramGuideResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramGuideResponseParams/paging
func (m_ MTRChannelClusterProgramGuideResponseParams) Paging() IMTRChannelClusterChannelPagingStruct {
	rv := objc.Send[MTRChannelClusterChannelPagingStruct](m_.ID, objc.Sel("paging"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramGuideResponseParams/paging
func (m_ MTRChannelClusterProgramGuideResponseParams) SetPaging(value IMTRChannelClusterChannelPagingStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPaging:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramGuideResponseParams/programList
func (m_ MTRChannelClusterProgramGuideResponseParams) ProgramList() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("programList"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramGuideResponseParams/programList
func (m_ MTRChannelClusterProgramGuideResponseParams) SetProgramList(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProgramList:"), value)
}


