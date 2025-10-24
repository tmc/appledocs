// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class TKBERTLVRecord */


/* debug [class_header]: Header for TKBERTLVRecord */
// The class instance for the [TKBERTLVRecord] class.
var (
	TKBERTLVRecordClass     _TKBERTLVRecordClass
	TKBERTLVRecordClassOnce sync.Once
)

func getTKBERTLVRecordClass() _TKBERTLVRecordClass {
	TKBERTLVRecordClassOnce.Do(func() {
		TKBERTLVRecordClass = _TKBERTLVRecordClass{objc.GetClass("TKBERTLVRecord")}
	})
	return TKBERTLVRecordClass
}

type _TKBERTLVRecordClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TKBERTLVRecord */
// An interface definition for the [TKBERTLVRecord] class.
type ITKBERTLVRecord interface {
	ITKTLVRecord
	
/* debug [class_interface_properties]: Properties for TKBERTLVRecord */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TKBERTLVRecord */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TKBERTLVRecord */
// Alloc allocates a new instance without initialization.
func (tc _TKBERTLVRecordClass) Alloc() TKBERTLVRecord {
	rv := objc.Send[TKBERTLVRecord](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TKBERTLVRecordClass) New() TKBERTLVRecord {
	rv := objc.Send[TKBERTLVRecord](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKBERTLVRecord) Init() TKBERTLVRecord {
	rv := objc.Send[TKBERTLVRecord](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKBERTLVRecord) Autorelease() TKBERTLVRecord {
	rv := objc.Send[TKBERTLVRecord](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKBERTLVRecord creates a new TKBERTLVRecord instance.
func NewTKBERTLVRecord() TKBERTLVRecord {
	return getTKBERTLVRecordClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TKBERTLVRecord */
// An object that parses BER-encoded data and produces DER-encoded data for TLV records.


// An object that parses BER-encoded data and produces DER-encoded data for TLV records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKBERTLVRecord
type TKBERTLVRecord struct {
	TKTLVRecord
}

// TKBERTLVRecordFrom constructs a [TKBERTLVRecord] from an unsafe.Pointer.
//
// An object that parses BER-encoded data and produces DER-encoded data for TLV records.
func TKBERTLVRecordFrom(ptr unsafe.Pointer) TKBERTLVRecord {
	return TKBERTLVRecord{
		TKTLVRecord: TKTLVRecordFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TKBERTLVRecord */

// Initializes a BER-TLV record with the specified tag and an array of TLV subrecords.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKBERTLVRecord/init(tag:records:)
func NewTKBERTLVRecordWithTagRecords(tag TKTLVTag /* typedef */, records []TKTLVRecord) TKBERTLVRecord {
	instance := getTKBERTLVRecordClass().Alloc()
	rv := objc.Send[TKBERTLVRecord](instance.ID, objc.Sel("initWithTag:records:"), tag, records)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTKBERTLVRecordWithTagRecords */


// Initializes a BER-TLV record with the specified tag and value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKBERTLVRecord/init(tag:value:)
func NewTKBERTLVRecordWithTagValue(tag TKTLVTag /* typedef */, value objc.IObject /* cross-framework: NSData */) TKBERTLVRecord {
	instance := getTKBERTLVRecordClass().Alloc()
	rv := objc.Send[TKBERTLVRecord](instance.ID, objc.Sel("initWithTag:value:"), tag, value)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTKBERTLVRecordWithTagValue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TKBERTLVRecord */

// Encodes a specified tag using BER-TLV tag encoding rules.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKBERTLVRecord/data(forTag:)
func (tc _TKBERTLVRecordClass) DataForTag(tag TKTLVTag /* typedef */) foundation.Data {
	rv := objc.Send[foundation.Data](objc.ID(tc.class), objc.Sel("dataForTag:"), tag)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DataForTag) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TKBERTLVRecord */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TKBERTLVRecord */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TKBERTLVRecord */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class TKBERTLVRecord */


