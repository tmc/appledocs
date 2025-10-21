// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MXSignpostRecord] class.
type IMXSignpostRecord interface {
	objectivec.IObject
	DictionaryRepresentation() unsafe.Pointer
	JSONRepresentation() unsafe.Pointer
}

// An object representing the record for a signpost interval or event.
//
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

// Alloc allocates a new instance without initialization.
func (mc _MXSignpostRecordClass) Alloc() MXSignpostRecord {
	rv := objc.Send[MXSignpostRecord](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXSignpostRecord/dictionaryRepresentation()
func (m_ MXSignpostRecord) DictionaryRepresentation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("dictionaryRepresentation"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXSignpostRecord/jsonRepresentation()
func (m_ MXSignpostRecord) JSONRepresentation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("JSONRepresentation"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXSignpostRecord/beginTimeStamp
func (m_ MXSignpostRecord) BeginTimeStamp() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("beginTimeStamp"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXSignpostRecord/category
func (m_ MXSignpostRecord) Category() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("category"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXSignpostRecord/duration
func (m_ MXSignpostRecord) Duration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("duration"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXSignpostRecord/endTimeStamp
func (m_ MXSignpostRecord) EndTimeStamp() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("endTimeStamp"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXSignpostRecord/isInterval
func (m_ MXSignpostRecord) IsInterval() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isInterval"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXSignpostRecord/name
func (m_ MXSignpostRecord) Name() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("name"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXSignpostRecord/subsystem
func (m_ MXSignpostRecord) Subsystem() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("subsystem"))
	return rv
}



