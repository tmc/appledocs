// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct */


/* debug [class_header]: Header for MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct */
// The class instance for the [MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct] class.
var (
	MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStructClass     _MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStructClass
	MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStructClassOnce sync.Once
)

func getMTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStructClass() _MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStructClass {
	MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStructClassOnce.Do(func() {
		MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStructClass = _MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStructClass{objc.GetClass("MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct")}
	})
	return MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStructClass
}

type _MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct */
// An interface definition for the [MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct] class.
type IMTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct */
	// properties:
	FixedMax() objc.IObject /* cross-framework: NSNumber */
	SetFixedMax(value objc.IObject /* cross-framework: NSNumber */)
	FixedMin() objc.IObject /* cross-framework: NSNumber */
	SetFixedMin(value objc.IObject /* cross-framework: NSNumber */)
	FixedTypical() objc.IObject /* cross-framework: NSNumber */
	SetFixedTypical(value objc.IObject /* cross-framework: NSNumber */)
	PercentMax() objc.IObject /* cross-framework: NSNumber */
	SetPercentMax(value objc.IObject /* cross-framework: NSNumber */)
	PercentMin() objc.IObject /* cross-framework: NSNumber */
	SetPercentMin(value objc.IObject /* cross-framework: NSNumber */)
	PercentTypical() objc.IObject /* cross-framework: NSNumber */
	SetPercentTypical(value objc.IObject /* cross-framework: NSNumber */)
	RangeMax() objc.IObject /* cross-framework: NSNumber */
	SetRangeMax(value objc.IObject /* cross-framework: NSNumber */)
	RangeMin() objc.IObject /* cross-framework: NSNumber */
	SetRangeMin(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStructClass) Alloc() MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct {
	rv := objc.Send[MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStructClass) New() MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct {
	rv := objc.Send[MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) Init() MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct {
	rv := objc.Send[MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) Autorelease() MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct {
	rv := objc.Send[MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct creates a new MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct instance.
func NewMTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct() MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct {
	return getMTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct
type MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct struct {
	objectivec.Object
}

// MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStructFrom constructs a [MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct] from an unsafe.Pointer.
func MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStructFrom(ptr unsafe.Pointer) MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct {
	return MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct/fixedMax
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) FixedMax() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fixedMax"))
	return rv
}/* debug [instance_properties/getter]: fixedMax */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct/fixedMax
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) SetFixedMax(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFixedMax:"), value)
}/* debug [instance_properties/setter]: fixedMax */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclustermeasurementaccuracyrangestruct/fixedmin
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) FixedMin() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fixedMin"))
	return rv
}/* debug [instance_properties/getter]: fixedMin */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclustermeasurementaccuracyrangestruct/fixedmin
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) SetFixedMin(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFixedMin:"), value)
}/* debug [instance_properties/setter]: fixedMin */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclustermeasurementaccuracyrangestruct/fixedtypical
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) FixedTypical() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fixedTypical"))
	return rv
}/* debug [instance_properties/getter]: fixedTypical */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclustermeasurementaccuracyrangestruct/fixedtypical
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) SetFixedTypical(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFixedTypical:"), value)
}/* debug [instance_properties/setter]: fixedTypical */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclustermeasurementaccuracyrangestruct/percentmax
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) PercentMax() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("percentMax"))
	return rv
}/* debug [instance_properties/getter]: percentMax */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclustermeasurementaccuracyrangestruct/percentmax
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) SetPercentMax(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPercentMax:"), value)
}/* debug [instance_properties/setter]: percentMax */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclustermeasurementaccuracyrangestruct/percentmin
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) PercentMin() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("percentMin"))
	return rv
}/* debug [instance_properties/getter]: percentMin */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclustermeasurementaccuracyrangestruct/percentmin
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) SetPercentMin(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPercentMin:"), value)
}/* debug [instance_properties/setter]: percentMin */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclustermeasurementaccuracyrangestruct/percenttypical
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) PercentTypical() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("percentTypical"))
	return rv
}/* debug [instance_properties/getter]: percentTypical */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclustermeasurementaccuracyrangestruct/percenttypical
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) SetPercentTypical(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPercentTypical:"), value)
}/* debug [instance_properties/setter]: percentTypical */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclustermeasurementaccuracyrangestruct/rangemax
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) RangeMax() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("rangeMax"))
	return rv
}/* debug [instance_properties/getter]: rangeMax */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclustermeasurementaccuracyrangestruct/rangemax
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) SetRangeMax(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRangeMax:"), value)
}/* debug [instance_properties/setter]: rangeMax */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclustermeasurementaccuracyrangestruct/rangemin
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) RangeMin() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("rangeMin"))
	return rv
}/* debug [instance_properties/getter]: rangeMin */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclustermeasurementaccuracyrangestruct/rangemin
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) SetRangeMin(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRangeMin:"), value)
}/* debug [instance_properties/setter]: rangeMin */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct */



