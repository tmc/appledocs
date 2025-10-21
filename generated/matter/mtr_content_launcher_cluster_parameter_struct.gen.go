// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRContentLauncherClusterParameterStruct] class.
var (
	MTRContentLauncherClusterParameterStructClass     _MTRContentLauncherClusterParameterStructClass
	MTRContentLauncherClusterParameterStructClassOnce sync.Once
)

func getMTRContentLauncherClusterParameterStructClass() _MTRContentLauncherClusterParameterStructClass {
	MTRContentLauncherClusterParameterStructClassOnce.Do(func() {
		MTRContentLauncherClusterParameterStructClass = _MTRContentLauncherClusterParameterStructClass{objc.GetClass("MTRContentLauncherClusterParameterStruct")}
	})
	return MTRContentLauncherClusterParameterStructClass
}

type _MTRContentLauncherClusterParameterStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRContentLauncherClusterParameterStruct] class.
type IMTRContentLauncherClusterParameterStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterParameterStruct
type MTRContentLauncherClusterParameterStruct struct {
	objectivec.Object
}

// MTRContentLauncherClusterParameterStructFrom constructs a [MTRContentLauncherClusterParameterStruct] from an unsafe.Pointer.
func MTRContentLauncherClusterParameterStructFrom(ptr unsafe.Pointer) MTRContentLauncherClusterParameterStruct {
	return MTRContentLauncherClusterParameterStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRContentLauncherClusterParameterStructClass) Alloc() MTRContentLauncherClusterParameterStruct {
	rv := objc.Send[MTRContentLauncherClusterParameterStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRContentLauncherClusterParameterStructClass) New() MTRContentLauncherClusterParameterStruct {
	rv := objc.Send[MTRContentLauncherClusterParameterStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRContentLauncherClusterParameterStruct) Init() MTRContentLauncherClusterParameterStruct {
	rv := objc.Send[MTRContentLauncherClusterParameterStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRContentLauncherClusterParameterStruct) Autorelease() MTRContentLauncherClusterParameterStruct {
	rv := objc.Send[MTRContentLauncherClusterParameterStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRContentLauncherClusterParameterStruct creates a new MTRContentLauncherClusterParameterStruct instance.
func NewMTRContentLauncherClusterParameterStruct() MTRContentLauncherClusterParameterStruct {
	return getMTRContentLauncherClusterParameterStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterparameterstruct/value
func (m_ MTRContentLauncherClusterParameterStruct) Value() string {
	rv := objc.Send[string](m_.ID, objc.Sel("value"))
	return rv
}


// SetValue sets the value of the value property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterparameterstruct/value
func (m_ MTRContentLauncherClusterParameterStruct) SetValue(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterparameterstruct/externalidlist
func (m_ MTRContentLauncherClusterParameterStruct) ExternalIDList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("externalIDList"))
	return rv
}


// SetExternalIDList sets the value of the externalIDList property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterparameterstruct/externalidlist
func (m_ MTRContentLauncherClusterParameterStruct) SetExternalIDList(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExternalIDList:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterparameterstruct/type
func (m_ MTRContentLauncherClusterParameterStruct) Type() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("type"))
	return rv
}


// SetType sets the value of the type property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterparameterstruct/type
func (m_ MTRContentLauncherClusterParameterStruct) SetType(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setType:"), value)
}



