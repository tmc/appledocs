// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRChannelClusterChannelPagingStruct] class.
var (
	MTRChannelClusterChannelPagingStructClass     _MTRChannelClusterChannelPagingStructClass
	MTRChannelClusterChannelPagingStructClassOnce sync.Once
)

func getMTRChannelClusterChannelPagingStructClass() _MTRChannelClusterChannelPagingStructClass {
	MTRChannelClusterChannelPagingStructClassOnce.Do(func() {
		MTRChannelClusterChannelPagingStructClass = _MTRChannelClusterChannelPagingStructClass{objc.GetClass("MTRChannelClusterChannelPagingStruct")}
	})
	return MTRChannelClusterChannelPagingStructClass
}

type _MTRChannelClusterChannelPagingStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRChannelClusterChannelPagingStruct] class.
type IMTRChannelClusterChannelPagingStruct interface {
	objectivec.IObject
	NextToken() MTRChannelClusterPageTokenStruct
	SetNextToken(value IMTRChannelClusterPageTokenStruct)
	PreviousToken() MTRChannelClusterPageTokenStruct
	SetPreviousToken(value IMTRChannelClusterPageTokenStruct)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterChannelPagingStruct
type MTRChannelClusterChannelPagingStruct struct {
	objectivec.Object
}

// MTRChannelClusterChannelPagingStructFrom constructs a [MTRChannelClusterChannelPagingStruct] from an unsafe.Pointer.
func MTRChannelClusterChannelPagingStructFrom(ptr unsafe.Pointer) MTRChannelClusterChannelPagingStruct {
	return MTRChannelClusterChannelPagingStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRChannelClusterChannelPagingStructClass) Alloc() MTRChannelClusterChannelPagingStruct {
	rv := objc.Send[MTRChannelClusterChannelPagingStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRChannelClusterChannelPagingStructClass) New() MTRChannelClusterChannelPagingStruct {
	rv := objc.Send[MTRChannelClusterChannelPagingStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRChannelClusterChannelPagingStruct) Init() MTRChannelClusterChannelPagingStruct {
	rv := objc.Send[MTRChannelClusterChannelPagingStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRChannelClusterChannelPagingStruct) Autorelease() MTRChannelClusterChannelPagingStruct {
	rv := objc.Send[MTRChannelClusterChannelPagingStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRChannelClusterChannelPagingStruct creates a new MTRChannelClusterChannelPagingStruct instance.
func NewMTRChannelClusterChannelPagingStruct() MTRChannelClusterChannelPagingStruct {
	return getMTRChannelClusterChannelPagingStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterChannelPagingStruct/nextToken
func (m_ MTRChannelClusterChannelPagingStruct) NextToken() MTRChannelClusterPageTokenStruct {
	rv := objc.Send[MTRChannelClusterPageTokenStruct](m_.ID, objc.Sel("nextToken"))
	return rv
}


// SetNextToken sets the value of the nextToken property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterChannelPagingStruct/nextToken
func (m_ MTRChannelClusterChannelPagingStruct) SetNextToken(value IMTRChannelClusterPageTokenStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNextToken:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterChannelPagingStruct/previousToken
func (m_ MTRChannelClusterChannelPagingStruct) PreviousToken() MTRChannelClusterPageTokenStruct {
	rv := objc.Send[MTRChannelClusterPageTokenStruct](m_.ID, objc.Sel("previousToken"))
	return rv
}


// SetPreviousToken sets the value of the previousToken property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterChannelPagingStruct/previousToken
func (m_ MTRChannelClusterChannelPagingStruct) SetPreviousToken(value IMTRChannelClusterPageTokenStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreviousToken:"), value)
}



