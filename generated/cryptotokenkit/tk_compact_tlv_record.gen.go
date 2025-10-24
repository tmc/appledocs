// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class TKCompactTLVRecord */


/* debug [class_header]: Header for TKCompactTLVRecord */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TKCompactTLVRecord */
// An interface definition for the [TKCompactTLVRecord] class.
type ITKCompactTLVRecord interface {
	ITKTLVRecord
	
/* debug [class_interface_properties]: Properties for TKCompactTLVRecord */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TKCompactTLVRecord */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TKCompactTLVRecord */
// Alloc allocates a new instance without initialization.
func (tc _TKCompactTLVRecordClass) Alloc() TKCompactTLVRecord {
	rv := objc.Send[TKCompactTLVRecord](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TKCompactTLVRecord */
// An object that implements encoding using Compact-TLV encoding according to ISO 7816-4.


// An object that implements encoding using Compact-TLV encoding according to ISO 7816-4.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKCompactTLVRecord
type TKCompactTLVRecord struct {
	TKTLVRecord
}

// TKCompactTLVRecordFrom constructs a [TKCompactTLVRecord] from an unsafe.Pointer.
//
// An object that implements encoding using Compact-TLV encoding according to ISO 7816-4.
func TKCompactTLVRecordFrom(ptr unsafe.Pointer) TKCompactTLVRecord {
	return TKCompactTLVRecord{
		TKTLVRecord: TKTLVRecordFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TKCompactTLVRecord */

// Initializes a TLV record with the specified tag and value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKCompactTLVRecord/init(tag:value:)
func NewTKCompactTLVRecordWithTagValue(tag unsafe.Pointer, value objc.IObject /* cross-framework: NSData */) TKCompactTLVRecord {
	instance := getTKCompactTLVRecordClass().Alloc()
	rv := objc.Send[TKCompactTLVRecord](instance.ID, objc.Sel("initWithTag:value:"), tag, value)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTKCompactTLVRecordWithTagValue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TKCompactTLVRecord */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TKCompactTLVRecord */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TKCompactTLVRecord */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TKCompactTLVRecord */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class TKCompactTLVRecord */


