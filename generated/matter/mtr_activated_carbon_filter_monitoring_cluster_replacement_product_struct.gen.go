// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct] class.
var (
	MTRActivatedCarbonFilterMonitoringClusterReplacementProductStructClass     _MTRActivatedCarbonFilterMonitoringClusterReplacementProductStructClass
	MTRActivatedCarbonFilterMonitoringClusterReplacementProductStructClassOnce sync.Once
)

func getMTRActivatedCarbonFilterMonitoringClusterReplacementProductStructClass() _MTRActivatedCarbonFilterMonitoringClusterReplacementProductStructClass {
	MTRActivatedCarbonFilterMonitoringClusterReplacementProductStructClassOnce.Do(func() {
		MTRActivatedCarbonFilterMonitoringClusterReplacementProductStructClass = _MTRActivatedCarbonFilterMonitoringClusterReplacementProductStructClass{objc.GetClass("MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct")}
	})
	return MTRActivatedCarbonFilterMonitoringClusterReplacementProductStructClass
}

type _MTRActivatedCarbonFilterMonitoringClusterReplacementProductStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct] class.
type IMTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct interface {
	objectivec.IObject
	// properties:
	ProductIdentifierType() objc.IObject /* cross-framework: NSNumber */
	SetProductIdentifierType(value objc.IObject /* cross-framework: NSNumber */)
	ProductIdentifierValue() objc.IObject /* cross-framework: NSString */
	SetProductIdentifierValue(value objc.IObject /* cross-framework: NSString */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct
type MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct struct {
	objectivec.Object
}

// MTRActivatedCarbonFilterMonitoringClusterReplacementProductStructFrom constructs a [MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct] from an unsafe.Pointer.
func MTRActivatedCarbonFilterMonitoringClusterReplacementProductStructFrom(ptr unsafe.Pointer) MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct {
	return MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRActivatedCarbonFilterMonitoringClusterReplacementProductStructClass) Alloc() MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct {
	rv := objc.Send[MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRActivatedCarbonFilterMonitoringClusterReplacementProductStructClass) New() MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct {
	rv := objc.Send[MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct) Init() MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct {
	rv := objc.Send[MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct) Autorelease() MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct {
	rv := objc.Send[MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct creates a new MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct instance.
func NewMTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct() MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct {
	return getMTRActivatedCarbonFilterMonitoringClusterReplacementProductStructClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractivatedcarbonfiltermonitoringclusterreplacementproductstruct/productidentifiertype
func (m_ MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct) ProductIdentifierType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("productIdentifierType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractivatedcarbonfiltermonitoringclusterreplacementproductstruct/productidentifiertype
func (m_ MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct) SetProductIdentifierType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProductIdentifierType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractivatedcarbonfiltermonitoringclusterreplacementproductstruct/productidentifiervalue
func (m_ MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct) ProductIdentifierValue() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("productIdentifierValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtractivatedcarbonfiltermonitoringclusterreplacementproductstruct/productidentifiervalue
func (m_ MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct) SetProductIdentifierValue(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProductIdentifierValue:"), value)
}



