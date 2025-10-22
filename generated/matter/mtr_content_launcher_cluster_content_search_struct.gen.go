// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRContentLauncherClusterContentSearchStruct] class.
var (
	MTRContentLauncherClusterContentSearchStructClass     _MTRContentLauncherClusterContentSearchStructClass
	MTRContentLauncherClusterContentSearchStructClassOnce sync.Once
)

func getMTRContentLauncherClusterContentSearchStructClass() _MTRContentLauncherClusterContentSearchStructClass {
	MTRContentLauncherClusterContentSearchStructClassOnce.Do(func() {
		MTRContentLauncherClusterContentSearchStructClass = _MTRContentLauncherClusterContentSearchStructClass{objc.GetClass("MTRContentLauncherClusterContentSearchStruct")}
	})
	return MTRContentLauncherClusterContentSearchStructClass
}

type _MTRContentLauncherClusterContentSearchStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRContentLauncherClusterContentSearchStruct] class.
type IMTRContentLauncherClusterContentSearchStruct interface {
	objectivec.IObject
	ParameterList() unsafe.Pointer
	SetParameterList(value unsafe.Pointer)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterContentSearchStruct
type MTRContentLauncherClusterContentSearchStruct struct {
	objectivec.Object
}

// MTRContentLauncherClusterContentSearchStructFrom constructs a [MTRContentLauncherClusterContentSearchStruct] from an unsafe.Pointer.
func MTRContentLauncherClusterContentSearchStructFrom(ptr unsafe.Pointer) MTRContentLauncherClusterContentSearchStruct {
	return MTRContentLauncherClusterContentSearchStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRContentLauncherClusterContentSearchStructClass) Alloc() MTRContentLauncherClusterContentSearchStruct {
	rv := objc.Send[MTRContentLauncherClusterContentSearchStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRContentLauncherClusterContentSearchStructClass) New() MTRContentLauncherClusterContentSearchStruct {
	rv := objc.Send[MTRContentLauncherClusterContentSearchStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRContentLauncherClusterContentSearchStruct) Init() MTRContentLauncherClusterContentSearchStruct {
	rv := objc.Send[MTRContentLauncherClusterContentSearchStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRContentLauncherClusterContentSearchStruct) Autorelease() MTRContentLauncherClusterContentSearchStruct {
	rv := objc.Send[MTRContentLauncherClusterContentSearchStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRContentLauncherClusterContentSearchStruct creates a new MTRContentLauncherClusterContentSearchStruct instance.
func NewMTRContentLauncherClusterContentSearchStruct() MTRContentLauncherClusterContentSearchStruct {
	return getMTRContentLauncherClusterContentSearchStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclustercontentsearchstruct/parameterlist
func (m_ MTRContentLauncherClusterContentSearchStruct) ParameterList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("parameterList"))
	return rv
}


// SetParameterList sets the value of the parameterList property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclustercontentsearchstruct/parameterlist
func (m_ MTRContentLauncherClusterContentSearchStruct) SetParameterList(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setParameterList:"), value)
}



