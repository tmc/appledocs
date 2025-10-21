// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRContentLauncherClusterAdditionalInfoStruct] class.
var (
	MTRContentLauncherClusterAdditionalInfoStructClass     _MTRContentLauncherClusterAdditionalInfoStructClass
	MTRContentLauncherClusterAdditionalInfoStructClassOnce sync.Once
)

func getMTRContentLauncherClusterAdditionalInfoStructClass() _MTRContentLauncherClusterAdditionalInfoStructClass {
	MTRContentLauncherClusterAdditionalInfoStructClassOnce.Do(func() {
		MTRContentLauncherClusterAdditionalInfoStructClass = _MTRContentLauncherClusterAdditionalInfoStructClass{objc.GetClass("MTRContentLauncherClusterAdditionalInfoStruct")}
	})
	return MTRContentLauncherClusterAdditionalInfoStructClass
}

type _MTRContentLauncherClusterAdditionalInfoStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRContentLauncherClusterAdditionalInfoStruct] class.
type IMTRContentLauncherClusterAdditionalInfoStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterAdditionalInfoStruct
type MTRContentLauncherClusterAdditionalInfoStruct struct {
	objectivec.Object
}

// MTRContentLauncherClusterAdditionalInfoStructFrom constructs a [MTRContentLauncherClusterAdditionalInfoStruct] from an unsafe.Pointer.
func MTRContentLauncherClusterAdditionalInfoStructFrom(ptr unsafe.Pointer) MTRContentLauncherClusterAdditionalInfoStruct {
	return MTRContentLauncherClusterAdditionalInfoStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRContentLauncherClusterAdditionalInfoStructClass) Alloc() MTRContentLauncherClusterAdditionalInfoStruct {
	rv := objc.Send[MTRContentLauncherClusterAdditionalInfoStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRContentLauncherClusterAdditionalInfoStructClass) New() MTRContentLauncherClusterAdditionalInfoStruct {
	rv := objc.Send[MTRContentLauncherClusterAdditionalInfoStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRContentLauncherClusterAdditionalInfoStruct) Init() MTRContentLauncherClusterAdditionalInfoStruct {
	rv := objc.Send[MTRContentLauncherClusterAdditionalInfoStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRContentLauncherClusterAdditionalInfoStruct) Autorelease() MTRContentLauncherClusterAdditionalInfoStruct {
	rv := objc.Send[MTRContentLauncherClusterAdditionalInfoStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRContentLauncherClusterAdditionalInfoStruct creates a new MTRContentLauncherClusterAdditionalInfoStruct instance.
func NewMTRContentLauncherClusterAdditionalInfoStruct() MTRContentLauncherClusterAdditionalInfoStruct {
	return getMTRContentLauncherClusterAdditionalInfoStructClass().New()
}




