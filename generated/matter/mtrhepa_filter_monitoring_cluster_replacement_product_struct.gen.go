// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRHEPAFilterMonitoringClusterReplacementProductStruct] class.
var (
	MTRHEPAFilterMonitoringClusterReplacementProductStructClass     _MTRHEPAFilterMonitoringClusterReplacementProductStructClass
	MTRHEPAFilterMonitoringClusterReplacementProductStructClassOnce sync.Once
)

func getMTRHEPAFilterMonitoringClusterReplacementProductStructClass() _MTRHEPAFilterMonitoringClusterReplacementProductStructClass {
	MTRHEPAFilterMonitoringClusterReplacementProductStructClassOnce.Do(func() {
		MTRHEPAFilterMonitoringClusterReplacementProductStructClass = _MTRHEPAFilterMonitoringClusterReplacementProductStructClass{objc.GetClass("MTRHEPAFilterMonitoringClusterReplacementProductStruct")}
	})
	return MTRHEPAFilterMonitoringClusterReplacementProductStructClass
}

type _MTRHEPAFilterMonitoringClusterReplacementProductStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRHEPAFilterMonitoringClusterReplacementProductStruct] class.
type IMTRHEPAFilterMonitoringClusterReplacementProductStruct interface {
	objectivec.IObject
	ProductIdentifierType() foundation.Number
	SetProductIdentifierType(value foundation.INumber)
	ProductIdentifierValue() string
	SetProductIdentifierValue(value string)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRHEPAFilterMonitoringClusterReplacementProductStruct
type MTRHEPAFilterMonitoringClusterReplacementProductStruct struct {
	objectivec.Object
}

// MTRHEPAFilterMonitoringClusterReplacementProductStructFrom constructs a [MTRHEPAFilterMonitoringClusterReplacementProductStruct] from an unsafe.Pointer.
func MTRHEPAFilterMonitoringClusterReplacementProductStructFrom(ptr unsafe.Pointer) MTRHEPAFilterMonitoringClusterReplacementProductStruct {
	return MTRHEPAFilterMonitoringClusterReplacementProductStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRHEPAFilterMonitoringClusterReplacementProductStructClass) Alloc() MTRHEPAFilterMonitoringClusterReplacementProductStruct {
	rv := objc.Send[MTRHEPAFilterMonitoringClusterReplacementProductStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRHEPAFilterMonitoringClusterReplacementProductStructClass) New() MTRHEPAFilterMonitoringClusterReplacementProductStruct {
	rv := objc.Send[MTRHEPAFilterMonitoringClusterReplacementProductStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRHEPAFilterMonitoringClusterReplacementProductStruct) Init() MTRHEPAFilterMonitoringClusterReplacementProductStruct {
	rv := objc.Send[MTRHEPAFilterMonitoringClusterReplacementProductStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRHEPAFilterMonitoringClusterReplacementProductStruct) Autorelease() MTRHEPAFilterMonitoringClusterReplacementProductStruct {
	rv := objc.Send[MTRHEPAFilterMonitoringClusterReplacementProductStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRHEPAFilterMonitoringClusterReplacementProductStruct creates a new MTRHEPAFilterMonitoringClusterReplacementProductStruct instance.
func NewMTRHEPAFilterMonitoringClusterReplacementProductStruct() MTRHEPAFilterMonitoringClusterReplacementProductStruct {
	return getMTRHEPAFilterMonitoringClusterReplacementProductStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrhepafiltermonitoringclusterreplacementproductstruct/productidentifiertype
func (m_ MTRHEPAFilterMonitoringClusterReplacementProductStruct) ProductIdentifierType() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("productIdentifierType"))
	return rv
}


// SetProductIdentifierType sets the value of the productIdentifierType property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrhepafiltermonitoringclusterreplacementproductstruct/productidentifiertype
func (m_ MTRHEPAFilterMonitoringClusterReplacementProductStruct) SetProductIdentifierType(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProductIdentifierType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrhepafiltermonitoringclusterreplacementproductstruct/productidentifiervalue
func (m_ MTRHEPAFilterMonitoringClusterReplacementProductStruct) ProductIdentifierValue() string {
	rv := objc.Send[string](m_.ID, objc.Sel("productIdentifierValue"))
	return rv
}


// SetProductIdentifierValue sets the value of the productIdentifierValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrhepafiltermonitoringclusterreplacementproductstruct/productidentifiervalue
func (m_ MTRHEPAFilterMonitoringClusterReplacementProductStruct) SetProductIdentifierValue(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProductIdentifierValue:"), objc.String(value))
}



