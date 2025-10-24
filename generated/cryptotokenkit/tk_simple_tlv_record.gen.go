// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class TKSimpleTLVRecord */


/* debug [class_header]: Header for TKSimpleTLVRecord */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TKSimpleTLVRecord */
// An interface definition for the [TKSimpleTLVRecord] class.
type ITKSimpleTLVRecord interface {
	ITKTLVRecord
	
/* debug [class_interface_properties]: Properties for TKSimpleTLVRecord */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TKSimpleTLVRecord */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TKSimpleTLVRecord */
// Alloc allocates a new instance without initialization.
func (tc _TKSimpleTLVRecordClass) Alloc() TKSimpleTLVRecord {
	rv := objc.Send[TKSimpleTLVRecord](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TKSimpleTLVRecord */
// An object that implements encoding using Simple-TLV encoding according to ISO 7816-4.


// An object that implements encoding using Simple-TLV encoding according to ISO 7816-4.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSimpleTLVRecord
type TKSimpleTLVRecord struct {
	TKTLVRecord
}

// TKSimpleTLVRecordFrom constructs a [TKSimpleTLVRecord] from an unsafe.Pointer.
//
// An object that implements encoding using Simple-TLV encoding according to ISO 7816-4.
func TKSimpleTLVRecordFrom(ptr unsafe.Pointer) TKSimpleTLVRecord {
	return TKSimpleTLVRecord{
		TKTLVRecord: TKTLVRecordFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TKSimpleTLVRecord */

// Initializes a TLV record with the specified tag and value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSimpleTLVRecord/init(tag:value:)
func NewTKSimpleTLVRecordWithTagValue(tag unsafe.Pointer, value objc.IObject /* cross-framework: NSData */) TKSimpleTLVRecord {
	instance := getTKSimpleTLVRecordClass().Alloc()
	rv := objc.Send[TKSimpleTLVRecord](instance.ID, objc.Sel("initWithTag:value:"), tag, value)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTKSimpleTLVRecordWithTagValue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TKSimpleTLVRecord */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TKSimpleTLVRecord */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TKSimpleTLVRecord */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TKSimpleTLVRecord */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class TKSimpleTLVRecord */


