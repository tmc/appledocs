// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MutableData] class.
var (
	mutableDataClass     _MutableDataClass
	mutableDataClassOnce sync.Once
)

func getMutableDataClass() _MutableDataClass {
	mutableDataClassOnce.Do(func() {
		mutableDataClass = _MutableDataClass{objc.GetClass("NSMutableData")}
	})
	return mutableDataClass
}

type _MutableDataClass struct {
	class objc.Class
}

// An interface definition for the [MutableData] class.
type IMutableData interface {
	IData
}

// An object representing a dynamic byte buffer in memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableData
type MutableData struct {
	Data
}

// MutableDataFrom constructs a [MutableData] from an unsafe.Pointer.
//
// An object representing a dynamic byte buffer in memory.
func MutableDataFrom(ptr unsafe.Pointer) MutableData {
	return MutableData{
		Data: DataFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MutableDataClass) Alloc() MutableData {
	rv := objc.Send[MutableData](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MutableDataClass) New() MutableData {
	rv := objc.Send[MutableData](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MutableData) Init() MutableData {
	rv := objc.Send[MutableData](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MutableData) Autorelease() MutableData {
	rv := objc.Send[MutableData](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMutableData creates a new MutableData instance.
func NewMutableData() MutableData {
	return getMutableDataClass().New()
}




