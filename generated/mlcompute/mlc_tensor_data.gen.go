// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CTensorData] class.
var (
	CTensorDataClass     _CTensorDataClass
	CTensorDataClassOnce sync.Once
)

func getCTensorDataClass() _CTensorDataClass {
	CTensorDataClassOnce.Do(func() {
		CTensorDataClass = _CTensorDataClass{objc.GetClass("MLCTensorData")}
	})
	return CTensorDataClass
}

type _CTensorDataClass struct {
	class objc.Class
}

// An interface definition for the [CTensorData] class.
type ICTensorData interface {
	objectivec.IObject
}

// An encapsulation of the memory that tensor data uses.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorData
type CTensorData struct {
	objectivec.Object
}

// CTensorDataFrom constructs a [CTensorData] from an unsafe.Pointer.
//
// An encapsulation of the memory that tensor data uses.
func CTensorDataFrom(ptr unsafe.Pointer) CTensorData {
	return CTensorData{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CTensorDataClass) Alloc() CTensorData {
	rv := objc.Send[CTensorData](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CTensorDataClass) New() CTensorData {
	rv := objc.Send[CTensorData](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CTensorData) Init() CTensorData {
	rv := objc.Send[CTensorData](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CTensorData) Autorelease() CTensorData {
	rv := objc.Send[CTensorData](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCTensorData creates a new CTensorData instance.
func NewCTensorData() CTensorData {
	return getCTensorDataClass().New()
}


// A buffer that conains data.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctensordata/bytes
func (c_ CTensorData) Bytes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("bytes"))
	return rv
}


// SetBytes sets the value of the bytes property.
// A buffer that conains data.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctensordata/bytes
func (c_ CTensorData) SetBytes(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBytes:"), value)
}

// The number of bytes you choose to hold for this tensor data instance.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctensordata/length
func (c_ CTensorData) Length() int {
	rv := objc.Send[int](c_.ID, objc.Sel("length"))
	return rv
}


// SetLength sets the value of the length property.
// The number of bytes you choose to hold for this tensor data instance.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctensordata/length
func (c_ CTensorData) SetLength(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLength:"), value)
}



