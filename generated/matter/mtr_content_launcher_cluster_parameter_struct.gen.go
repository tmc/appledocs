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
	// properties:
	ExternalIDList() unsafe.Pointer
	SetExternalIDList(value unsafe.Pointer)
	Type() objc.IObject /* cross-framework: NSNumber */
	SetType(value objc.IObject /* cross-framework: NSNumber */)
	Value() objc.IObject /* cross-framework: NSString */
	SetValue(value objc.IObject /* cross-framework: NSString */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterparameterstruct/externalidlist
func (m_ MTRContentLauncherClusterParameterStruct) ExternalIDList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("externalIDList"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterparameterstruct/externalidlist
func (m_ MTRContentLauncherClusterParameterStruct) SetExternalIDList(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExternalIDList:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterparameterstruct/type
func (m_ MTRContentLauncherClusterParameterStruct) Type() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("type"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterparameterstruct/type
func (m_ MTRContentLauncherClusterParameterStruct) SetType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterparameterstruct/value
func (m_ MTRContentLauncherClusterParameterStruct) Value() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("value"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterparameterstruct/value
func (m_ MTRContentLauncherClusterParameterStruct) SetValue(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}



