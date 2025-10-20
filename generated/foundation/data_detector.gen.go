// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DataDetector] class.
var (
	dataDetectorClass     _DataDetectorClass
	dataDetectorClassOnce sync.Once
)

func getDataDetectorClass() _DataDetectorClass {
	dataDetectorClassOnce.Do(func() {
		dataDetectorClass = _DataDetectorClass{objc.GetClass("NSDataDetector")}
	})
	return dataDetectorClass
}

type _DataDetectorClass struct {
	class objc.Class
}

// An interface definition for the [DataDetector] class.
type IDataDetector interface {
	IRegularExpression
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




