// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRContentLauncherClusterParameter] class.
var (
	MTRContentLauncherClusterParameterClass     _MTRContentLauncherClusterParameterClass
	MTRContentLauncherClusterParameterClassOnce sync.Once
)

func getMTRContentLauncherClusterParameterClass() _MTRContentLauncherClusterParameterClass {
	MTRContentLauncherClusterParameterClassOnce.Do(func() {
		MTRContentLauncherClusterParameterClass = _MTRContentLauncherClusterParameterClass{objc.GetClass("MTRContentLauncherClusterParameter")}
	})
	return MTRContentLauncherClusterParameterClass
}

type _MTRContentLauncherClusterParameterClass struct {
	class objc.Class
}

// An interface definition for the [MTRContentLauncherClusterParameter] class.
type IMTRContentLauncherClusterParameter interface {
	IMTRContentLauncherClusterParameterStruct
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterParameter
type MTRContentLauncherClusterParameter struct {
	MTRContentLauncherClusterParameterStruct
}

// MTRContentLauncherClusterParameterFrom constructs a [MTRContentLauncherClusterParameter] from an unsafe.Pointer.
func MTRContentLauncherClusterParameterFrom(ptr unsafe.Pointer) MTRContentLauncherClusterParameter {
	return MTRContentLauncherClusterParameter{
		MTRContentLauncherClusterParameterStruct: MTRContentLauncherClusterParameterStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRContentLauncherClusterParameterClass) Alloc() MTRContentLauncherClusterParameter {
	rv := objc.Send[MTRContentLauncherClusterParameter](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRContentLauncherClusterParameterClass) New() MTRContentLauncherClusterParameter {
	rv := objc.Send[MTRContentLauncherClusterParameter](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRContentLauncherClusterParameter) Init() MTRContentLauncherClusterParameter {
	rv := objc.Send[MTRContentLauncherClusterParameter](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRContentLauncherClusterParameter) Autorelease() MTRContentLauncherClusterParameter {
	rv := objc.Send[MTRContentLauncherClusterParameter](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRContentLauncherClusterParameter creates a new MTRContentLauncherClusterParameter instance.
func NewMTRContentLauncherClusterParameter() MTRContentLauncherClusterParameter {
	return getMTRContentLauncherClusterParameterClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterparameter/externalidlist
func (m_ MTRContentLauncherClusterParameter) ExternalIDList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("externalIDList"))
	return rv
}


// SetExternalIDList sets the value of the externalIDList property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterparameter/externalidlist
func (m_ MTRContentLauncherClusterParameter) SetExternalIDList(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExternalIDList:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterparameter/type
func (m_ MTRContentLauncherClusterParameter) Type() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("type"))
	return rv
}


// SetType sets the value of the type property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterparameter/type
func (m_ MTRContentLauncherClusterParameter) SetType(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterparameter/value
func (m_ MTRContentLauncherClusterParameter) Value() string {
	rv := objc.Send[string](m_.ID, objc.Sel("value"))
	return rv
}


// SetValue sets the value of the value property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterparameter/value
func (m_ MTRContentLauncherClusterParameter) SetValue(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), objc.String(value))
}



