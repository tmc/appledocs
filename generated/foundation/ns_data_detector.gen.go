// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSDataDetector */


/* debug [class_header]: Header for NSDataDetector */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DataDetector */
// An interface definition for the [DataDetector] class.
type IDataDetector interface {
	IRegularExpression
	
/* debug [class_interface_properties]: Properties for DataDetector */
	// properties:
	CheckingTypes() TextCheckingTypes /* typedef */
	NSNotFound() int
	Date() IDate
	SetDate(value IDate)
	Duration() float64
	SetDuration(value float64)
	TimeZone() ITimeZone
	SetTimeZone(value ITimeZone)
	Url() IURL
	SetUrl(value IURL)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DataDetector */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DataDetector */
// Alloc allocates a new instance without initialization.
func (dc _DataDetectorClass) Alloc() DataDetector {
	rv := objc.Send[DataDetector](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DataDetector */
// A specialized regular expression object that matches natural language text for predefined data patterns.
//
// Find dates, addresses, links, phone numbers, and transit information in natural language text with . returns the results of matching content in objects. The objects that returns are different from those that returns. The results are one of the data detector’s types and contain the corresponding properties. For example, results of type have a , , and ; and results of type have a .


// A specialized regular expression object that matches natural language text for predefined data patterns.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DataDetector */

// Initializes and returns a data detector instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDataDetector/init(types:)
func NewDataDetectorWithTypesError(checkingTypes TextCheckingTypes /* typedef */, error_ IError) DataDetector {
	instance := getDataDetectorClass().Alloc()
	rv := objc.Send[DataDetector](instance.ID, objc.Sel("initWithTypes:error:"), checkingTypes, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDataDetectorWithTypesError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DataDetector */

// Creates and returns a new data detector instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDataDetector/dataDetectorWithTypes:error:
func (dc _DataDetectorClass) DataDetectorWithTypesError(checkingTypes TextCheckingTypes /* typedef */, error_ IError) IDataDetector {
	rv := objc.Send[DataDetector](objc.ID(dc.class), objc.Sel("dataDetectorWithTypes:error:"), checkingTypes, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DataDetectorWithTypesError) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DataDetector */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DataDetector */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DataDetector */

// Returns the checking types for the data detector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDataDetector/checkingTypes
func (d_ DataDetector) CheckingTypes() TextCheckingTypes /* typedef */ {
	rv := objc.Send[uint64](d_.ID, objc.Sel("checkingTypes"))
	return rv
}/* debug [instance_properties/getter]: checkingTypes */


// A value indicating that a requested item couldn’t be found or doesn’t exist.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotfound-4qp9h
func (d_ DataDetector) NSNotFound() int {
	rv := objc.Send[int](d_.ID, objc.Sel("NSNotFound"))
	return rv
}/* debug [instance_properties/getter]: NSNotFound */


// The date component of a type checking result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/date
func (d_ DataDetector) Date() IDate {
	rv := objc.Send[Date](d_.ID, objc.Sel("date"))
	return rv
}/* debug [instance_properties/getter]: date */


// The date component of a type checking result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/date
func (d_ DataDetector) SetDate(value IDate) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDate:"), value)
}/* debug [instance_properties/setter]: date */


// The duration component of a type checking result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/duration
func (d_ DataDetector) Duration() float64 {
	rv := objc.Send[float64](d_.ID, objc.Sel("duration"))
	return rv
}/* debug [instance_properties/getter]: duration */


// The duration component of a type checking result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/duration
func (d_ DataDetector) SetDuration(value float64) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDuration:"), value)
}/* debug [instance_properties/setter]: duration */


// The time zone component of a type checking result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/timezone
func (d_ DataDetector) TimeZone() ITimeZone {
	rv := objc.Send[TimeZone](d_.ID, objc.Sel("timeZone"))
	return rv
}/* debug [instance_properties/getter]: timeZone */


// The time zone component of a type checking result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/timezone
func (d_ DataDetector) SetTimeZone(value ITimeZone) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTimeZone:"), value)
}/* debug [instance_properties/setter]: timeZone */


// The URL of a type checking result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/url
func (d_ DataDetector) Url() IURL {
	rv := objc.Send[URL](d_.ID, objc.Sel("url"))
	return rv
}/* debug [instance_properties/getter]: url */


// The URL of a type checking result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nstextcheckingresult/url
func (d_ DataDetector) SetUrl(value IURL) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setUrl:"), value)
}/* debug [instance_properties/setter]: url */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSDataDetector */


