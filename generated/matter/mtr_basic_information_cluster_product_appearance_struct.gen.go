// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRBasicInformationClusterProductAppearanceStruct] class.
var (
	MTRBasicInformationClusterProductAppearanceStructClass     _MTRBasicInformationClusterProductAppearanceStructClass
	MTRBasicInformationClusterProductAppearanceStructClassOnce sync.Once
)

func getMTRBasicInformationClusterProductAppearanceStructClass() _MTRBasicInformationClusterProductAppearanceStructClass {
	MTRBasicInformationClusterProductAppearanceStructClassOnce.Do(func() {
		MTRBasicInformationClusterProductAppearanceStructClass = _MTRBasicInformationClusterProductAppearanceStructClass{objc.GetClass("MTRBasicInformationClusterProductAppearanceStruct")}
	})
	return MTRBasicInformationClusterProductAppearanceStructClass
}

type _MTRBasicInformationClusterProductAppearanceStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRBasicInformationClusterProductAppearanceStruct] class.
type IMTRBasicInformationClusterProductAppearanceStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBasicInformationClusterProductAppearanceStruct
type MTRBasicInformationClusterProductAppearanceStruct struct {
	objectivec.Object
}

// MTRBasicInformationClusterProductAppearanceStructFrom constructs a [MTRBasicInformationClusterProductAppearanceStruct] from an unsafe.Pointer.
func MTRBasicInformationClusterProductAppearanceStructFrom(ptr unsafe.Pointer) MTRBasicInformationClusterProductAppearanceStruct {
	return MTRBasicInformationClusterProductAppearanceStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBasicInformationClusterProductAppearanceStructClass) Alloc() MTRBasicInformationClusterProductAppearanceStruct {
	rv := objc.Send[MTRBasicInformationClusterProductAppearanceStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBasicInformationClusterProductAppearanceStructClass) New() MTRBasicInformationClusterProductAppearanceStruct {
	rv := objc.Send[MTRBasicInformationClusterProductAppearanceStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBasicInformationClusterProductAppearanceStruct) Init() MTRBasicInformationClusterProductAppearanceStruct {
	rv := objc.Send[MTRBasicInformationClusterProductAppearanceStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBasicInformationClusterProductAppearanceStruct) Autorelease() MTRBasicInformationClusterProductAppearanceStruct {
	rv := objc.Send[MTRBasicInformationClusterProductAppearanceStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBasicInformationClusterProductAppearanceStruct creates a new MTRBasicInformationClusterProductAppearanceStruct instance.
func NewMTRBasicInformationClusterProductAppearanceStruct() MTRBasicInformationClusterProductAppearanceStruct {
	return getMTRBasicInformationClusterProductAppearanceStructClass().New()
}




