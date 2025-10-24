// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MXSignpostRecord */


/* debug [class_header]: Header for MXSignpostRecord */
// The class instance for the [MXSignpostRecord] class.
var (
	MXSignpostRecordClass     _MXSignpostRecordClass
	MXSignpostRecordClassOnce sync.Once
)

func getMXSignpostRecordClass() _MXSignpostRecordClass {
	MXSignpostRecordClassOnce.Do(func() {
		MXSignpostRecordClass = _MXSignpostRecordClass{objc.GetClass("MXSignpostRecord")}
	})
	return MXSignpostRecordClass
}

type _MXSignpostRecordClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MXSignpostRecord */
// An interface definition for the [MXSignpostRecord] class.
type IMXSignpostRecord interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MXSignpostRecord */
	// properties:
	BeginTimeStamp() objc.IObject /* cross-framework: NSDate */
	Category() objc.IObject /* cross-framework: NSString */
	Duration() unsafe.Pointer
	EndTimeStamp() objc.IObject /* cross-framework: NSDate */
	IsInterval() bool
	Name() objc.IObject /* cross-framework: NSString */
	Subsystem() objc.IObject /* cross-framework: NSString */
	MXErrorDomain() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MXSignpostRecord */
	// methods:
	DictionaryRepresentation() foundation.Dictionary
	JSONRepresentation() foundation.Data
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MXSignpostRecord */
// Alloc allocates a new instance without initialization.
func (mc _MXSignpostRecordClass) Alloc() MXSignpostRecord {
	rv := objc.Send[MXSignpostRecord](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MXSignpostRecordClass) New() MXSignpostRecord {
	rv := objc.Send[MXSignpostRecord](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXSignpostRecord) Init() MXSignpostRecord {
	rv := objc.Send[MXSignpostRecord](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXSignpostRecord) Autorelease() MXSignpostRecord {
	rv := objc.Send[MXSignpostRecord](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXSignpostRecord creates a new MXSignpostRecord instance.
func NewMXSignpostRecord() MXSignpostRecord {
	return getMXSignpostRecordClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MXSignpostRecord */
// An object representing the record for a signpost interval or event.


// An object representing the record for a signpost interval or event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXSignpostRecord
type MXSignpostRecord struct {
	objectivec.Object
}

// MXSignpostRecordFrom constructs a [MXSignpostRecord] from an unsafe.Pointer.
//
// An object representing the record for a signpost interval or event.
func MXSignpostRecordFrom(ptr unsafe.Pointer) MXSignpostRecord {
	return MXSignpostRecord{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MXSignpostRecord *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MXSignpostRecord */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MXSignpostRecord */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MXSignpostRecord */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXSignpostRecord/dictionaryRepresentation()
func (m_ MXSignpostRecord) DictionaryRepresentation() foundation.Dictionary {
	rv := objc.Send[foundation.Dictionary](m_.ID, objc.Sel("dictionaryRepresentation"))
	return rv
}/* debug [instance_methods/method]: DictionaryRepresentation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXSignpostRecord/jsonRepresentation()
func (m_ MXSignpostRecord) JSONRepresentation() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("JSONRepresentation"))
	return rv
}/* debug [instance_methods/method]: JSONRepresentation */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MXSignpostRecord */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXSignpostRecord/beginTimeStamp
func (m_ MXSignpostRecord) BeginTimeStamp() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("beginTimeStamp"))
	return rv
}/* debug [instance_properties/getter]: beginTimeStamp */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXSignpostRecord/category
func (m_ MXSignpostRecord) Category() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("category"))
	return rv
}/* debug [instance_properties/getter]: category */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXSignpostRecord/duration
func (m_ MXSignpostRecord) Duration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("duration"))
	return rv
}/* debug [instance_properties/getter]: duration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXSignpostRecord/endTimeStamp
func (m_ MXSignpostRecord) EndTimeStamp() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("endTimeStamp"))
	return rv
}/* debug [instance_properties/getter]: endTimeStamp */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXSignpostRecord/isInterval
func (m_ MXSignpostRecord) IsInterval() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isInterval"))
	return rv
}/* debug [instance_properties/getter]: isInterval */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXSignpostRecord/name
func (m_ MXSignpostRecord) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXSignpostRecord/subsystem
func (m_ MXSignpostRecord) Subsystem() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("subsystem"))
	return rv
}/* debug [instance_properties/getter]: subsystem */


// Error domain for error values from app metrics.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxerrordomain
func (m_ MXSignpostRecord) MXErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("MXErrorDomain"))
	return rv
}/* debug [instance_properties/getter]: MXErrorDomain */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MXSignpostRecord */



