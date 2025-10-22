// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DataDetector] class.
var (
	DataDetectorClass     _DataDetectorClass
	DataDetectorClassOnce sync.Once
)

func getDataDetectorClass() _DataDetectorClass {
	DataDetectorClassOnce.Do(func() {
		DataDetectorClass = _DataDetectorClass{objc.GetClass("NSDataDetector")}
	})
	return DataDetectorClass
}

type _DataDetectorClass struct {
	class objc.Class
}

// An interface definition for the [DataDetector] class.
type IDataDetector interface {
	IRegularExpression
	CheckingTypes() unsafe.Pointer
	SetCheckingTypes(value unsafe.Pointer)
	NSNotFound() int
	Date() Date
	SetDate(value IDate)
	Duration() TimeInterval
	SetDuration(value ITimeInterval)
	TimeZone() TimeZone
	SetTimeZone(value ITimeZone)
	Url() URL
	SetUrl(value IURL)
}

// A specialized regular expression object that matches natural language text for predefined data patterns.
//
// Find dates, addresses, links, phone numbers, and transit information in natural language text with . returns the results of matching content in objects. The objects that returns are different from those that returns. The results are one of the data detector’s types and contain the corresponding properties. For example, results of type have a , , and ; and results of type have a .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDataDetector
type DataDetector struct {
	RegularExpression
}

// DataDetectorFrom constructs a [DataDetector] from an unsafe.Pointer.
//
// A specialized regular expression object that matches natural language text for predefined data patterns.
func DataDetectorFrom(ptr unsafe.Pointer) DataDetector {
	return DataDetector{
		RegularExpression: RegularExpressionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (dc _DataDetectorClass) Alloc() DataDetector {
	rv := objc.Send[DataDetector](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DataDetectorClass) New() DataDetector {
	rv := objc.Send[DataDetector](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DataDetector) Init() DataDetector {
	rv := objc.Send[DataDetector](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DataDetector) Autorelease() DataDetector {
	rv := objc.Send[DataDetector](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDataDetector creates a new DataDetector instance.
func NewDataDetector() DataDetector {
	return getDataDetectorClass().New()
}


// Returns the checking types for the data detector.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatadetector/checkingtypes
func (d_ DataDetector) CheckingTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("checkingTypes"))
	return rv
}


// SetCheckingTypes sets the value of the checkingTypes property.
// Returns the checking types for the data detector.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatadetector/checkingtypes
func (d_ DataDetector) SetCheckingTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCheckingTypes:"), value)
}

// A value indicating that a requested item couldn’t be found or doesn’t exist.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotfound-4qp9h
func (d_ DataDetector) NSNotFound() int {
	rv := objc.Send[int](d_.ID, objc.Sel("NSNotFound"))
	return rv
}

// The date component of a type checking result.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/date
func (d_ DataDetector) Date() Date {
	rv := objc.Send[Date](d_.ID, objc.Sel("date"))
	return rv
}


// SetDate sets the value of the date property.
// The date component of a type checking result.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/date
func (d_ DataDetector) SetDate(value IDate) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDate:"), value)
}

// The duration component of a type checking result.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/duration
func (d_ DataDetector) Duration() TimeInterval {
	rv := objc.Send[TimeInterval](d_.ID, objc.Sel("duration"))
	return rv
}


// SetDuration sets the value of the duration property.
// The duration component of a type checking result.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/duration
func (d_ DataDetector) SetDuration(value ITimeInterval) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDuration:"), value)
}

// The time zone component of a type checking result.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/timezone
func (d_ DataDetector) TimeZone() TimeZone {
	rv := objc.Send[TimeZone](d_.ID, objc.Sel("timeZone"))
	return rv
}


// SetTimeZone sets the value of the timeZone property.
// The time zone component of a type checking result.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/timezone
func (d_ DataDetector) SetTimeZone(value ITimeZone) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTimeZone:"), value)
}

// The URL of a type checking result.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/url
func (d_ DataDetector) Url() URL {
	rv := objc.Send[URL](d_.ID, objc.Sel("url"))
	return rv
}


// SetUrl sets the value of the url property.
// The URL of a type checking result.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/url
func (d_ DataDetector) SetUrl(value IURL) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setUrl:"), value)
}



