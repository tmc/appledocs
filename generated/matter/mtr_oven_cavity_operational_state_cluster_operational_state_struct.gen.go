// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROvenCavityOperationalStateClusterOperationalStateStruct] class.
var (
	MTROvenCavityOperationalStateClusterOperationalStateStructClass     _MTROvenCavityOperationalStateClusterOperationalStateStructClass
	MTROvenCavityOperationalStateClusterOperationalStateStructClassOnce sync.Once
)

func getMTROvenCavityOperationalStateClusterOperationalStateStructClass() _MTROvenCavityOperationalStateClusterOperationalStateStructClass {
	MTROvenCavityOperationalStateClusterOperationalStateStructClassOnce.Do(func() {
		MTROvenCavityOperationalStateClusterOperationalStateStructClass = _MTROvenCavityOperationalStateClusterOperationalStateStructClass{objc.GetClass("MTROvenCavityOperationalStateClusterOperationalStateStruct")}
	})
	return MTROvenCavityOperationalStateClusterOperationalStateStructClass
}

type _MTROvenCavityOperationalStateClusterOperationalStateStructClass struct {
	class objc.Class
}

// An interface definition for the [MTROvenCavityOperationalStateClusterOperationalStateStruct] class.
type IMTROvenCavityOperationalStateClusterOperationalStateStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterOperationalStateStruct
type MTROvenCavityOperationalStateClusterOperationalStateStruct struct {
	objectivec.Object
}

// MTROvenCavityOperationalStateClusterOperationalStateStructFrom constructs a [MTROvenCavityOperationalStateClusterOperationalStateStruct] from an unsafe.Pointer.
func MTROvenCavityOperationalStateClusterOperationalStateStructFrom(ptr unsafe.Pointer) MTROvenCavityOperationalStateClusterOperationalStateStruct {
	return MTROvenCavityOperationalStateClusterOperationalStateStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROvenCavityOperationalStateClusterOperationalStateStructClass) Alloc() MTROvenCavityOperationalStateClusterOperationalStateStruct {
	rv := objc.Send[MTROvenCavityOperationalStateClusterOperationalStateStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROvenCavityOperationalStateClusterOperationalStateStructClass) New() MTROvenCavityOperationalStateClusterOperationalStateStruct {
	rv := objc.Send[MTROvenCavityOperationalStateClusterOperationalStateStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROvenCavityOperationalStateClusterOperationalStateStruct) Init() MTROvenCavityOperationalStateClusterOperationalStateStruct {
	rv := objc.Send[MTROvenCavityOperationalStateClusterOperationalStateStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROvenCavityOperationalStateClusterOperationalStateStruct) Autorelease() MTROvenCavityOperationalStateClusterOperationalStateStruct {
	rv := objc.Send[MTROvenCavityOperationalStateClusterOperationalStateStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROvenCavityOperationalStateClusterOperationalStateStruct creates a new MTROvenCavityOperationalStateClusterOperationalStateStruct instance.
func NewMTROvenCavityOperationalStateClusterOperationalStateStruct() MTROvenCavityOperationalStateClusterOperationalStateStruct {
	return getMTROvenCavityOperationalStateClusterOperationalStateStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterOperationalStateStruct/operationalStateID
func (m_ MTROvenCavityOperationalStateClusterOperationalStateStruct) OperationalStateID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("operationalStateID"))
	return rv
}


// SetOperationalStateID sets the value of the operationalStateID property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterOperationalStateStruct/operationalStateID
func (m_ MTROvenCavityOperationalStateClusterOperationalStateStruct) SetOperationalStateID(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationalStateID:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterOperationalStateStruct/operationalStateLabel
func (m_ MTROvenCavityOperationalStateClusterOperationalStateStruct) OperationalStateLabel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("operationalStateLabel"))
	return rv
}


// SetOperationalStateLabel sets the value of the operationalStateLabel property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterOperationalStateStruct/operationalStateLabel
func (m_ MTROvenCavityOperationalStateClusterOperationalStateStruct) SetOperationalStateLabel(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationalStateLabel:"), value)
}


