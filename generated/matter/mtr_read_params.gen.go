// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRReadParams */


/* debug [class_header]: Header for MTRReadParams */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRReadParams */
// An interface definition for the [MTRReadParams] class.
type IMTRReadParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRReadParams */
	// properties:
	FabricFiltered() objc.IObject /* cross-framework: NSNumber */
	SetFabricFiltered(value objc.IObject /* cross-framework: NSNumber */)
	MinEventNumber() objc.IObject /* cross-framework: NSNumber */
	SetMinEventNumber(value objc.IObject /* cross-framework: NSNumber */)
	AssumeUnknownAttributesReportable() bool
	SetAssumeUnknownAttributesReportable(value bool)
	FilterByFabric() bool
	SetFilterByFabric(value bool)
	ShouldAssumeUnknownAttributesReportable() bool
	SetShouldAssumeUnknownAttributesReportable(value bool)
	ShouldFilterByFabric() bool
	SetShouldFilterByFabric(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRReadParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRReadParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRReadParamsClass) Alloc() MTRReadParams {
	rv := objc.Send[MTRReadParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRReadParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRReadParams
type MTRReadParams struct {
	objectivec.Object
}

// MTRReadParamsFrom constructs a [MTRReadParams] from an unsafe.Pointer.
func MTRReadParamsFrom(ptr unsafe.Pointer) MTRReadParams {
	return MTRReadParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRReadParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRReadParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRReadParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRReadParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRReadParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRReadParams/fabricFiltered
func (m_ MTRReadParams) FabricFiltered() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricFiltered"))
	return rv
}/* debug [instance_properties/getter]: fabricFiltered */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRReadParams/fabricFiltered
func (m_ MTRReadParams) SetFabricFiltered(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricFiltered:"), value)
}/* debug [instance_properties/setter]: fabricFiltered */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRReadParams/minEventNumber
func (m_ MTRReadParams) MinEventNumber() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("minEventNumber"))
	return rv
}/* debug [instance_properties/getter]: minEventNumber */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRReadParams/minEventNumber
func (m_ MTRReadParams) SetMinEventNumber(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinEventNumber:"), value)
}/* debug [instance_properties/setter]: minEventNumber */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRReadParams/shouldAssumeUnknownAttributesReportable
func (m_ MTRReadParams) AssumeUnknownAttributesReportable() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("assumeUnknownAttributesReportable"))
	return rv
}/* debug [instance_properties/getter]: assumeUnknownAttributesReportable */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRReadParams/shouldAssumeUnknownAttributesReportable
func (m_ MTRReadParams) SetAssumeUnknownAttributesReportable(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAssumeUnknownAttributesReportable:"), value)
}/* debug [instance_properties/setter]: assumeUnknownAttributesReportable */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRReadParams/shouldFilterByFabric
func (m_ MTRReadParams) FilterByFabric() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("filterByFabric"))
	return rv
}/* debug [instance_properties/getter]: filterByFabric */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRReadParams/shouldFilterByFabric
func (m_ MTRReadParams) SetFilterByFabric(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFilterByFabric:"), value)
}/* debug [instance_properties/setter]: filterByFabric */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrreadparams/shouldassumeunknownattributesreportable
func (m_ MTRReadParams) ShouldAssumeUnknownAttributesReportable() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("shouldAssumeUnknownAttributesReportable"))
	return rv
}/* debug [instance_properties/getter]: shouldAssumeUnknownAttributesReportable */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrreadparams/shouldassumeunknownattributesreportable
func (m_ MTRReadParams) SetShouldAssumeUnknownAttributesReportable(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShouldAssumeUnknownAttributesReportable:"), value)
}/* debug [instance_properties/setter]: shouldAssumeUnknownAttributesReportable */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrreadparams/shouldfilterbyfabric
func (m_ MTRReadParams) ShouldFilterByFabric() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("shouldFilterByFabric"))
	return rv
}/* debug [instance_properties/getter]: shouldFilterByFabric */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrreadparams/shouldfilterbyfabric
func (m_ MTRReadParams) SetShouldFilterByFabric(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShouldFilterByFabric:"), value)
}/* debug [instance_properties/setter]: shouldFilterByFabric */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRReadParams */



