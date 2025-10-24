// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRLaundryWasherModeClusterModeTagStruct] class.
var (
	MTRLaundryWasherModeClusterModeTagStructClass     _MTRLaundryWasherModeClusterModeTagStructClass
	MTRLaundryWasherModeClusterModeTagStructClassOnce sync.Once
)

func getMTRLaundryWasherModeClusterModeTagStructClass() _MTRLaundryWasherModeClusterModeTagStructClass {
	MTRLaundryWasherModeClusterModeTagStructClassOnce.Do(func() {
		MTRLaundryWasherModeClusterModeTagStructClass = _MTRLaundryWasherModeClusterModeTagStructClass{objc.GetClass("MTRLaundryWasherModeClusterModeTagStruct")}
	})
	return MTRLaundryWasherModeClusterModeTagStructClass
}

type _MTRLaundryWasherModeClusterModeTagStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRLaundryWasherModeClusterModeTagStruct] class.
type IMTRLaundryWasherModeClusterModeTagStruct interface {
	objectivec.IObject
	// properties:
	MfgCode() objc.IObject /* cross-framework: NSNumber */
	SetMfgCode(value objc.IObject /* cross-framework: NSNumber */)
	Value() objc.IObject /* cross-framework: NSNumber */
	SetValue(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLaundryWasherModeClusterModeTagStruct
type MTRLaundryWasherModeClusterModeTagStruct struct {
	objectivec.Object
}

// MTRLaundryWasherModeClusterModeTagStructFrom constructs a [MTRLaundryWasherModeClusterModeTagStruct] from an unsafe.Pointer.
func MTRLaundryWasherModeClusterModeTagStructFrom(ptr unsafe.Pointer) MTRLaundryWasherModeClusterModeTagStruct {
	return MTRLaundryWasherModeClusterModeTagStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRLaundryWasherModeClusterModeTagStructClass) Alloc() MTRLaundryWasherModeClusterModeTagStruct {
	rv := objc.Send[MTRLaundryWasherModeClusterModeTagStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRLaundryWasherModeClusterModeTagStructClass) New() MTRLaundryWasherModeClusterModeTagStruct {
	rv := objc.Send[MTRLaundryWasherModeClusterModeTagStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRLaundryWasherModeClusterModeTagStruct) Init() MTRLaundryWasherModeClusterModeTagStruct {
	rv := objc.Send[MTRLaundryWasherModeClusterModeTagStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRLaundryWasherModeClusterModeTagStruct) Autorelease() MTRLaundryWasherModeClusterModeTagStruct {
	rv := objc.Send[MTRLaundryWasherModeClusterModeTagStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRLaundryWasherModeClusterModeTagStruct creates a new MTRLaundryWasherModeClusterModeTagStruct instance.
func NewMTRLaundryWasherModeClusterModeTagStruct() MTRLaundryWasherModeClusterModeTagStruct {
	return getMTRLaundryWasherModeClusterModeTagStructClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLaundryWasherModeClusterModeTagStruct/mfgCode
func (m_ MTRLaundryWasherModeClusterModeTagStruct) MfgCode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("mfgCode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLaundryWasherModeClusterModeTagStruct/mfgCode
func (m_ MTRLaundryWasherModeClusterModeTagStruct) SetMfgCode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMfgCode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLaundryWasherModeClusterModeTagStruct/value
func (m_ MTRLaundryWasherModeClusterModeTagStruct) Value() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("value"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLaundryWasherModeClusterModeTagStruct/value
func (m_ MTRLaundryWasherModeClusterModeTagStruct) SetValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}



