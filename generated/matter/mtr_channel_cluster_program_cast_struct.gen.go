// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRChannelClusterProgramCastStruct] class.
var (
	MTRChannelClusterProgramCastStructClass     _MTRChannelClusterProgramCastStructClass
	MTRChannelClusterProgramCastStructClassOnce sync.Once
)

func getMTRChannelClusterProgramCastStructClass() _MTRChannelClusterProgramCastStructClass {
	MTRChannelClusterProgramCastStructClassOnce.Do(func() {
		MTRChannelClusterProgramCastStructClass = _MTRChannelClusterProgramCastStructClass{objc.GetClass("MTRChannelClusterProgramCastStruct")}
	})
	return MTRChannelClusterProgramCastStructClass
}

type _MTRChannelClusterProgramCastStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRChannelClusterProgramCastStruct] class.
type IMTRChannelClusterProgramCastStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramCastStruct
type MTRChannelClusterProgramCastStruct struct {
	objectivec.Object
}

// MTRChannelClusterProgramCastStructFrom constructs a [MTRChannelClusterProgramCastStruct] from an unsafe.Pointer.
func MTRChannelClusterProgramCastStructFrom(ptr unsafe.Pointer) MTRChannelClusterProgramCastStruct {
	return MTRChannelClusterProgramCastStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRChannelClusterProgramCastStructClass) Alloc() MTRChannelClusterProgramCastStruct {
	rv := objc.Send[MTRChannelClusterProgramCastStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRChannelClusterProgramCastStructClass) New() MTRChannelClusterProgramCastStruct {
	rv := objc.Send[MTRChannelClusterProgramCastStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRChannelClusterProgramCastStruct) Init() MTRChannelClusterProgramCastStruct {
	rv := objc.Send[MTRChannelClusterProgramCastStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRChannelClusterProgramCastStruct) Autorelease() MTRChannelClusterProgramCastStruct {
	rv := objc.Send[MTRChannelClusterProgramCastStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRChannelClusterProgramCastStruct creates a new MTRChannelClusterProgramCastStruct instance.
func NewMTRChannelClusterProgramCastStruct() MTRChannelClusterProgramCastStruct {
	return getMTRChannelClusterProgramCastStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramCastStruct/name
func (m_ MTRChannelClusterProgramCastStruct) Name() string {
	rv := objc.Send[string](m_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramCastStruct/name
func (m_ MTRChannelClusterProgramCastStruct) SetName(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramCastStruct/role
func (m_ MTRChannelClusterProgramCastStruct) Role() string {
	rv := objc.Send[string](m_.ID, objc.Sel("role"))
	return rv
}


// SetRole sets the value of the role property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterProgramCastStruct/role
func (m_ MTRChannelClusterProgramCastStruct) SetRole(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRole:"), objc.String(value))
}



