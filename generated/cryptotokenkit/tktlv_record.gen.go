// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class TKTLVRecord */


/* debug [class_header]: Header for TKTLVRecord */
// The class instance for the [TKTLVRecord] class.
var (
	TKTLVRecordClass     _TKTLVRecordClass
	TKTLVRecordClassOnce sync.Once
)

func getTKTLVRecordClass() _TKTLVRecordClass {
	TKTLVRecordClassOnce.Do(func() {
		TKTLVRecordClass = _TKTLVRecordClass{objc.GetClass("TKTLVRecord")}
	})
	return TKTLVRecordClass
}

type _TKTLVRecordClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TKTLVRecord */
// An interface definition for the [TKTLVRecord] class.
type ITKTLVRecord interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TKTLVRecord */
	// properties:
	Data() objc.IObject /* cross-framework: NSData */
	Tag() TKTLVTag /* typedef */
	Value() objc.IObject /* cross-framework: NSData */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TKTLVRecord */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TKTLVRecord */
// Alloc allocates a new instance without initialization.
func (tc _TKTLVRecordClass) Alloc() TKTLVRecord {
	rv := objc.Send[TKTLVRecord](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TKTLVRecordClass) New() TKTLVRecord {
	rv := objc.Send[TKTLVRecord](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKTLVRecord) Init() TKTLVRecord {
	rv := objc.Send[TKTLVRecord](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKTLVRecord) Autorelease() TKTLVRecord {
	rv := objc.Send[TKTLVRecord](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKTLVRecord creates a new TKTLVRecord instance.
func NewTKTLVRecord() TKTLVRecord {
	return getTKTLVRecordClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TKTLVRecord */
// The base class encapsulating a Tag-Length-Value record.
//
// The CryptoTokenKit framework provides the following concrete subclasses for various TLV record encodings: for BER-TLV encoding rules for Simple-TLV encoding according to ISO 7816-4 for Compact-TLV encoding according to ISO 7816-4


// The base class encapsulating a Tag-Length-Value record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTLVRecord
type TKTLVRecord struct {
	objectivec.Object
}

// TKTLVRecordFrom constructs a [TKTLVRecord] from an unsafe.Pointer.
//
// The base class encapsulating a Tag-Length-Value record.
func TKTLVRecordFrom(ptr unsafe.Pointer) TKTLVRecord {
	return TKTLVRecord{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TKTLVRecord */

// Creates and returns a TLV record from by parsing the specified data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTLVRecord/init(from:)
func NewTKTLVRecordFromData(data objc.IObject /* cross-framework: NSData */) TKTLVRecord {
	rv := objc.Send[TKTLVRecord](objc.ID(getTKTLVRecordClass().class), objc.Sel("recordFromData:"), data)
	return rv
}/* debug [class_init_methods/constructor]: NewTKTLVRecordFromData */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TKTLVRecord */

// Creates and returns a TLV record from by parsing the specified data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTLVRecord/init(from:)
func (tc _TKTLVRecordClass) RecordFromData(data objc.IObject /* cross-framework: NSData */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("recordFromData:"), data)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RecordFromData) */


// Creates and returns an array of TLV records from the specified data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTLVRecord/sequenceOfRecords(from:)
func (tc _TKTLVRecordClass) SequenceOfRecordsFromData(data objc.IObject /* cross-framework: NSData */) []TKTLVRecord {
	rv := objc.Send[[]TKTLVRecord](objc.ID(tc.class), objc.Sel("sequenceOfRecordsFromData:"), data)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SequenceOfRecordsFromData) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TKTLVRecord */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TKTLVRecord */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TKTLVRecord */

// The record data, including the tag, length, and value fields.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTLVRecord/data
func (t_ TKTLVRecord) Data() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](t_.ID, objc.Sel("data"))
	return rv
}/* debug [instance_properties/getter]: data */


// The tag field of the record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTLVRecord/tag
func (t_ TKTLVRecord) Tag() TKTLVTag /* typedef */ {
	rv := objc.Send[uint64](t_.ID, objc.Sel("tag"))
	return rv
}/* debug [instance_properties/getter]: tag */


// The value field of the record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTLVRecord/value
func (t_ TKTLVRecord) Value() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](t_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class TKTLVRecord */


