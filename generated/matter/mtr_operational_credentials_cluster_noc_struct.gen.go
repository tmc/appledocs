// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROperationalCredentialsClusterNOCStruct] class.
var (
	MTROperationalCredentialsClusterNOCStructClass     _MTROperationalCredentialsClusterNOCStructClass
	MTROperationalCredentialsClusterNOCStructClassOnce sync.Once
)

func getMTROperationalCredentialsClusterNOCStructClass() _MTROperationalCredentialsClusterNOCStructClass {
	MTROperationalCredentialsClusterNOCStructClassOnce.Do(func() {
		MTROperationalCredentialsClusterNOCStructClass = _MTROperationalCredentialsClusterNOCStructClass{objc.GetClass("MTROperationalCredentialsClusterNOCStruct")}
	})
	return MTROperationalCredentialsClusterNOCStructClass
}

type _MTROperationalCredentialsClusterNOCStructClass struct {
	class objc.Class
}

// An interface definition for the [MTROperationalCredentialsClusterNOCStruct] class.
type IMTROperationalCredentialsClusterNOCStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterNOCStruct
type MTROperationalCredentialsClusterNOCStruct struct {
	objectivec.Object
}

// MTROperationalCredentialsClusterNOCStructFrom constructs a [MTROperationalCredentialsClusterNOCStruct] from an unsafe.Pointer.
func MTROperationalCredentialsClusterNOCStructFrom(ptr unsafe.Pointer) MTROperationalCredentialsClusterNOCStruct {
	return MTROperationalCredentialsClusterNOCStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROperationalCredentialsClusterNOCStructClass) Alloc() MTROperationalCredentialsClusterNOCStruct {
	rv := objc.Send[MTROperationalCredentialsClusterNOCStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROperationalCredentialsClusterNOCStructClass) New() MTROperationalCredentialsClusterNOCStruct {
	rv := objc.Send[MTROperationalCredentialsClusterNOCStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalCredentialsClusterNOCStruct) Init() MTROperationalCredentialsClusterNOCStruct {
	rv := objc.Send[MTROperationalCredentialsClusterNOCStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalCredentialsClusterNOCStruct) Autorelease() MTROperationalCredentialsClusterNOCStruct {
	rv := objc.Send[MTROperationalCredentialsClusterNOCStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalCredentialsClusterNOCStruct creates a new MTROperationalCredentialsClusterNOCStruct instance.
func NewMTROperationalCredentialsClusterNOCStruct() MTROperationalCredentialsClusterNOCStruct {
	return getMTROperationalCredentialsClusterNOCStructClass().New()
}




