// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRChannelClusterPageTokenStruct] class.
var (
	MTRChannelClusterPageTokenStructClass     _MTRChannelClusterPageTokenStructClass
	MTRChannelClusterPageTokenStructClassOnce sync.Once
)

func getMTRChannelClusterPageTokenStructClass() _MTRChannelClusterPageTokenStructClass {
	MTRChannelClusterPageTokenStructClassOnce.Do(func() {
		MTRChannelClusterPageTokenStructClass = _MTRChannelClusterPageTokenStructClass{objc.GetClass("MTRChannelClusterPageTokenStruct")}
	})
	return MTRChannelClusterPageTokenStructClass
}

type _MTRChannelClusterPageTokenStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRChannelClusterPageTokenStruct] class.
type IMTRChannelClusterPageTokenStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterPageTokenStruct
type MTRChannelClusterPageTokenStruct struct {
	objectivec.Object
}

// MTRChannelClusterPageTokenStructFrom constructs a [MTRChannelClusterPageTokenStruct] from an unsafe.Pointer.
func MTRChannelClusterPageTokenStructFrom(ptr unsafe.Pointer) MTRChannelClusterPageTokenStruct {
	return MTRChannelClusterPageTokenStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRChannelClusterPageTokenStructClass) Alloc() MTRChannelClusterPageTokenStruct {
	rv := objc.Send[MTRChannelClusterPageTokenStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRChannelClusterPageTokenStructClass) New() MTRChannelClusterPageTokenStruct {
	rv := objc.Send[MTRChannelClusterPageTokenStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRChannelClusterPageTokenStruct) Init() MTRChannelClusterPageTokenStruct {
	rv := objc.Send[MTRChannelClusterPageTokenStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRChannelClusterPageTokenStruct) Autorelease() MTRChannelClusterPageTokenStruct {
	rv := objc.Send[MTRChannelClusterPageTokenStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRChannelClusterPageTokenStruct creates a new MTRChannelClusterPageTokenStruct instance.
func NewMTRChannelClusterPageTokenStruct() MTRChannelClusterPageTokenStruct {
	return getMTRChannelClusterPageTokenStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterPageTokenStruct/after
func (m_ MTRChannelClusterPageTokenStruct) After() string {
	rv := objc.Send[string](m_.ID, objc.Sel("after"))
	return rv
}


// SetAfter sets the value of the after property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterPageTokenStruct/after
func (m_ MTRChannelClusterPageTokenStruct) SetAfter(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAfter:"), objc.String(value))
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterPageTokenStruct/before
func (m_ MTRChannelClusterPageTokenStruct) Before() string {
	rv := objc.Send[string](m_.ID, objc.Sel("before"))
	return rv
}


// SetBefore sets the value of the before property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterPageTokenStruct/before
func (m_ MTRChannelClusterPageTokenStruct) SetBefore(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBefore:"), objc.String(value))
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterPageTokenStruct/limit
func (m_ MTRChannelClusterPageTokenStruct) Limit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("limit"))
	return rv
}


// SetLimit sets the value of the limit property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterPageTokenStruct/limit
func (m_ MTRChannelClusterPageTokenStruct) SetLimit(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLimit:"), value)
}


