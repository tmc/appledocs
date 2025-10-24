// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRReadParams] class.
var (
	MTRReadParamsClass     _MTRReadParamsClass
	MTRReadParamsClassOnce sync.Once
)

func getMTRReadParamsClass() _MTRReadParamsClass {
	MTRReadParamsClassOnce.Do(func() {
		MTRReadParamsClass = _MTRReadParamsClass{objc.GetClass("MTRReadParams")}
	})
	return MTRReadParamsClass
}

type _MTRReadParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRReadParams] class.
type IMTRReadParams interface {
	objectivec.IObject
	// properties:
	FabricFiltered() objc.IObject /* cross-framework: NSNumber */
	SetFabricFiltered(value objc.IObject /* cross-framework: NSNumber */)
	MinEventNumber() objc.IObject /* cross-framework: NSNumber */
	SetMinEventNumber(value objc.IObject /* cross-framework: NSNumber */)
	ShouldAssumeUnknownAttributesReportable() bool
	SetShouldAssumeUnknownAttributesReportable(value bool)
	ShouldFilterByFabric() bool
	SetShouldFilterByFabric(value bool)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRReadParams
type MTRReadParams struct {
	objectivec.Object
}

// MTRReadParamsFrom constructs a [MTRReadParams] from an unsafe.Pointer.
func MTRReadParamsFrom(ptr unsafe.Pointer) MTRReadParams {
	return MTRReadParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRReadParamsClass) Alloc() MTRReadParams {
	rv := objc.Send[MTRReadParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRReadParamsClass) New() MTRReadParams {
	rv := objc.Send[MTRReadParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRReadParams) Init() MTRReadParams {
	rv := objc.Send[MTRReadParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRReadParams) Autorelease() MTRReadParams {
	rv := objc.Send[MTRReadParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRReadParams creates a new MTRReadParams instance.
func NewMTRReadParams() MTRReadParams {
	return getMTRReadParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrreadparams/fabricfiltered
func (m_ MTRReadParams) FabricFiltered() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricFiltered"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrreadparams/fabricfiltered
func (m_ MTRReadParams) SetFabricFiltered(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricFiltered:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrreadparams/mineventnumber
func (m_ MTRReadParams) MinEventNumber() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("minEventNumber"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrreadparams/mineventnumber
func (m_ MTRReadParams) SetMinEventNumber(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinEventNumber:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrreadparams/shouldassumeunknownattributesreportable
func (m_ MTRReadParams) ShouldAssumeUnknownAttributesReportable() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("shouldAssumeUnknownAttributesReportable"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrreadparams/shouldassumeunknownattributesreportable
func (m_ MTRReadParams) SetShouldAssumeUnknownAttributesReportable(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShouldAssumeUnknownAttributesReportable:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrreadparams/shouldfilterbyfabric
func (m_ MTRReadParams) ShouldFilterByFabric() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("shouldFilterByFabric"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrreadparams/shouldfilterbyfabric
func (m_ MTRReadParams) SetShouldFilterByFabric(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShouldFilterByFabric:"), value)
}



