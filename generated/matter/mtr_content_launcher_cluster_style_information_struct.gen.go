// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRContentLauncherClusterStyleInformationStruct] class.
var (
	MTRContentLauncherClusterStyleInformationStructClass     _MTRContentLauncherClusterStyleInformationStructClass
	MTRContentLauncherClusterStyleInformationStructClassOnce sync.Once
)

func getMTRContentLauncherClusterStyleInformationStructClass() _MTRContentLauncherClusterStyleInformationStructClass {
	MTRContentLauncherClusterStyleInformationStructClassOnce.Do(func() {
		MTRContentLauncherClusterStyleInformationStructClass = _MTRContentLauncherClusterStyleInformationStructClass{objc.GetClass("MTRContentLauncherClusterStyleInformationStruct")}
	})
	return MTRContentLauncherClusterStyleInformationStructClass
}

type _MTRContentLauncherClusterStyleInformationStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRContentLauncherClusterStyleInformationStruct] class.
type IMTRContentLauncherClusterStyleInformationStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterStyleInformationStruct
type MTRContentLauncherClusterStyleInformationStruct struct {
	objectivec.Object
}

// MTRContentLauncherClusterStyleInformationStructFrom constructs a [MTRContentLauncherClusterStyleInformationStruct] from an unsafe.Pointer.
func MTRContentLauncherClusterStyleInformationStructFrom(ptr unsafe.Pointer) MTRContentLauncherClusterStyleInformationStruct {
	return MTRContentLauncherClusterStyleInformationStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRContentLauncherClusterStyleInformationStructClass) Alloc() MTRContentLauncherClusterStyleInformationStruct {
	rv := objc.Send[MTRContentLauncherClusterStyleInformationStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRContentLauncherClusterStyleInformationStructClass) New() MTRContentLauncherClusterStyleInformationStruct {
	rv := objc.Send[MTRContentLauncherClusterStyleInformationStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRContentLauncherClusterStyleInformationStruct) Init() MTRContentLauncherClusterStyleInformationStruct {
	rv := objc.Send[MTRContentLauncherClusterStyleInformationStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRContentLauncherClusterStyleInformationStruct) Autorelease() MTRContentLauncherClusterStyleInformationStruct {
	rv := objc.Send[MTRContentLauncherClusterStyleInformationStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRContentLauncherClusterStyleInformationStruct creates a new MTRContentLauncherClusterStyleInformationStruct instance.
func NewMTRContentLauncherClusterStyleInformationStruct() MTRContentLauncherClusterStyleInformationStruct {
	return getMTRContentLauncherClusterStyleInformationStructClass().New()
}




