// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TKSimpleTLVRecord] class.
var (
	TKSimpleTLVRecordClass     _TKSimpleTLVRecordClass
	TKSimpleTLVRecordClassOnce sync.Once
)

func getTKSimpleTLVRecordClass() _TKSimpleTLVRecordClass {
	TKSimpleTLVRecordClassOnce.Do(func() {
		TKSimpleTLVRecordClass = _TKSimpleTLVRecordClass{objc.GetClass("TKSimpleTLVRecord")}
	})
	return TKSimpleTLVRecordClass
}

type _TKSimpleTLVRecordClass struct {
	class objc.Class
}

// An interface definition for the [TKSimpleTLVRecord] class.
type ITKSimpleTLVRecord interface {
	objectivec.IObject
}

// An object that implements encoding using Simple-TLV encoding according to ISO 7816-4.


// An object that implements encoding using Simple-TLV encoding according to ISO 7816-4.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSimpleTLVRecord
type TKSimpleTLVRecord struct {
	objectivec.Object
}

// TKSimpleTLVRecordFrom constructs a [TKSimpleTLVRecord] from an unsafe.Pointer.
//
// An object that implements encoding using Simple-TLV encoding according to ISO 7816-4.
func TKSimpleTLVRecordFrom(ptr unsafe.Pointer) TKSimpleTLVRecord {
	return TKSimpleTLVRecord{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TKSimpleTLVRecordClass) Alloc() TKSimpleTLVRecord {
	rv := objc.Send[TKSimpleTLVRecord](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TKSimpleTLVRecordClass) New() TKSimpleTLVRecord {
	rv := objc.Send[TKSimpleTLVRecord](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKSimpleTLVRecord) Init() TKSimpleTLVRecord {
	rv := objc.Send[TKSimpleTLVRecord](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKSimpleTLVRecord) Autorelease() TKSimpleTLVRecord {
	rv := objc.Send[TKSimpleTLVRecord](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKSimpleTLVRecord creates a new TKSimpleTLVRecord instance.
func NewTKSimpleTLVRecord() TKSimpleTLVRecord {
	return getTKSimpleTLVRecordClass().New()
}




