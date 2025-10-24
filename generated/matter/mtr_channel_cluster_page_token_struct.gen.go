// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	After() objc.IObject /* cross-framework: NSString */
	SetAfter(value objc.IObject /* cross-framework: NSString */)
	Before() objc.IObject /* cross-framework: NSString */
	SetBefore(value objc.IObject /* cross-framework: NSString */)
	Limit() objc.IObject /* cross-framework: NSNumber */
	SetLimit(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterPageTokenStruct/after
func (m_ MTRChannelClusterPageTokenStruct) After() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("after"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterPageTokenStruct/after
func (m_ MTRChannelClusterPageTokenStruct) SetAfter(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAfter:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterPageTokenStruct/before
func (m_ MTRChannelClusterPageTokenStruct) Before() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("before"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterPageTokenStruct/before
func (m_ MTRChannelClusterPageTokenStruct) SetBefore(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBefore:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterPageTokenStruct/limit
func (m_ MTRChannelClusterPageTokenStruct) Limit() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("limit"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterPageTokenStruct/limit
func (m_ MTRChannelClusterPageTokenStruct) SetLimit(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLimit:"), value)
}



