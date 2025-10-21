// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TKCompactTLVRecord] class.
var (
	TKCompactTLVRecordClass     _TKCompactTLVRecordClass
	TKCompactTLVRecordClassOnce sync.Once
)

func getTKCompactTLVRecordClass() _TKCompactTLVRecordClass {
	TKCompactTLVRecordClassOnce.Do(func() {
		TKCompactTLVRecordClass = _TKCompactTLVRecordClass{objc.GetClass("TKCompactTLVRecord")}
	})
	return TKCompactTLVRecordClass
}

type _TKCompactTLVRecordClass struct {
	class objc.Class
}

// An interface definition for the [TKCompactTLVRecord] class.
type ITKCompactTLVRecord interface {
	objectivec.IObject
}

// An object that implements encoding using Compact-TLV encoding according to ISO 7816-4.
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKCompactTLVRecord
type TKCompactTLVRecord struct {
	objectivec.Object
}

// TKCompactTLVRecordFrom constructs a [TKCompactTLVRecord] from an unsafe.Pointer.
//
// An object that implements encoding using Compact-TLV encoding according to ISO 7816-4.
func TKCompactTLVRecordFrom(ptr unsafe.Pointer) TKCompactTLVRecord {
	return TKCompactTLVRecord{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TKCompactTLVRecordClass) Alloc() TKCompactTLVRecord {
	rv := objc.Send[TKCompactTLVRecord](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TKCompactTLVRecordClass) New() TKCompactTLVRecord {
	rv := objc.Send[TKCompactTLVRecord](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKCompactTLVRecord) Init() TKCompactTLVRecord {
	rv := objc.Send[TKCompactTLVRecord](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKCompactTLVRecord) Autorelease() TKCompactTLVRecord {
	rv := objc.Send[TKCompactTLVRecord](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKCompactTLVRecord creates a new TKCompactTLVRecord instance.
func NewTKCompactTLVRecord() TKCompactTLVRecord {
	return getTKCompactTLVRecordClass().New()
}




// Initializes a TLV record with the specified tag and value.
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKCompactTLVRecord/init(tag:value:)
func NewTKCompactTLVRecordWithTagValue(tag unsafe.Pointer, value unsafe.Pointer) TKCompactTLVRecord {
	instance := getTKCompactTLVRecordClass().Alloc()
	rv := objc.Send[TKCompactTLVRecord](instance.ID, objc.Sel("initWithTag:value:"), tag, value)
	rv.Autorelease()
	return rv
}



