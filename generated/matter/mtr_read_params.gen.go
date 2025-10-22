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
	FabricFiltered() foundation.Number
	SetFabricFiltered(value foundation.INumber)
	MinEventNumber() foundation.Number
	SetMinEventNumber(value foundation.INumber)
	AssumeUnknownAttributesReportable() bool
	SetAssumeUnknownAttributesReportable(value bool)
	FilterByFabric() bool
	SetFilterByFabric(value bool)
	ShouldAssumeUnknownAttributesReportable() bool
	SetShouldAssumeUnknownAttributesReportable(value bool)
	ShouldFilterByFabric() bool
	SetShouldFilterByFabric(value bool)
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRReadParams/fabricFiltered
func (m_ MTRReadParams) FabricFiltered() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("fabricFiltered"))
	return rv
}


// SetFabricFiltered sets the value of the fabricFiltered property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRReadParams/fabricFiltered
func (m_ MTRReadParams) SetFabricFiltered(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricFiltered:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRReadParams/minEventNumber
func (m_ MTRReadParams) MinEventNumber() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("minEventNumber"))
	return rv
}


// SetMinEventNumber sets the value of the minEventNumber property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRReadParams/minEventNumber
func (m_ MTRReadParams) SetMinEventNumber(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinEventNumber:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRReadParams/shouldAssumeUnknownAttributesReportable
func (m_ MTRReadParams) AssumeUnknownAttributesReportable() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("assumeUnknownAttributesReportable"))
	return rv
}


// SetAssumeUnknownAttributesReportable sets the value of the assumeUnknownAttributesReportable property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRReadParams/shouldAssumeUnknownAttributesReportable
func (m_ MTRReadParams) SetAssumeUnknownAttributesReportable(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAssumeUnknownAttributesReportable:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRReadParams/shouldFilterByFabric
func (m_ MTRReadParams) FilterByFabric() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("filterByFabric"))
	return rv
}


// SetFilterByFabric sets the value of the filterByFabric property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRReadParams/shouldFilterByFabric
func (m_ MTRReadParams) SetFilterByFabric(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFilterByFabric:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrreadparams/shouldassumeunknownattributesreportable
func (m_ MTRReadParams) ShouldAssumeUnknownAttributesReportable() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("shouldAssumeUnknownAttributesReportable"))
	return rv
}


// SetShouldAssumeUnknownAttributesReportable sets the value of the shouldAssumeUnknownAttributesReportable property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrreadparams/shouldassumeunknownattributesreportable
func (m_ MTRReadParams) SetShouldAssumeUnknownAttributesReportable(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShouldAssumeUnknownAttributesReportable:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrreadparams/shouldfilterbyfabric
func (m_ MTRReadParams) ShouldFilterByFabric() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("shouldFilterByFabric"))
	return rv
}


// SetShouldFilterByFabric sets the value of the shouldFilterByFabric property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrreadparams/shouldfilterbyfabric
func (m_ MTRReadParams) SetShouldFilterByFabric(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShouldFilterByFabric:"), value)
}



