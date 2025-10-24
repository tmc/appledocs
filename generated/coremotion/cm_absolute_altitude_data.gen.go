// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class CMAbsoluteAltitudeData */


/* debug [class_header]: Header for CMAbsoluteAltitudeData */
// The class instance for the [AbsoluteAltitudeData] class.
var (
	AbsoluteAltitudeDataClass     _AbsoluteAltitudeDataClass
	AbsoluteAltitudeDataClassOnce sync.Once
)

func getAbsoluteAltitudeDataClass() _AbsoluteAltitudeDataClass {
	AbsoluteAltitudeDataClassOnce.Do(func() {
		AbsoluteAltitudeDataClass = _AbsoluteAltitudeDataClass{objc.GetClass("CMAbsoluteAltitudeData")}
	})
	return AbsoluteAltitudeDataClass
}

type _AbsoluteAltitudeDataClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AbsoluteAltitudeData */
// An interface definition for the [AbsoluteAltitudeData] class.
type IAbsoluteAltitudeData interface {
	ILogItem
	
/* debug [class_interface_properties]: Properties for AbsoluteAltitudeData */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AbsoluteAltitudeData */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AbsoluteAltitudeData */
// Alloc allocates a new instance without initialization.
func (ac _AbsoluteAltitudeDataClass) Alloc() AbsoluteAltitudeData {
	rv := objc.Send[AbsoluteAltitudeData](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AbsoluteAltitudeDataClass) New() AbsoluteAltitudeData {
	rv := objc.Send[AbsoluteAltitudeData](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AbsoluteAltitudeData) Init() AbsoluteAltitudeData {
	rv := objc.Send[AbsoluteAltitudeData](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AbsoluteAltitudeData) Autorelease() AbsoluteAltitudeData {
	rv := objc.Send[AbsoluteAltitudeData](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAbsoluteAltitudeData creates a new AbsoluteAltitudeData instance.
func NewAbsoluteAltitudeData() AbsoluteAltitudeData {
	return getAbsoluteAltitudeDataClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AbsoluteAltitudeData */
// Data that records a change in absolute altitude.
//
// Absolute altitude is only available on iPhone 12 and later and Apple Watch 6 or SE and later.


// Data that records a change in absolute altitude.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAbsoluteAltitudeData
type AbsoluteAltitudeData struct {
	LogItem
}

// AbsoluteAltitudeDataFrom constructs a [AbsoluteAltitudeData] from an unsafe.Pointer.
//
// Data that records a change in absolute altitude.
func AbsoluteAltitudeDataFrom(ptr unsafe.Pointer) AbsoluteAltitudeData {
	return AbsoluteAltitudeData{
		LogItem: LogItemFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AbsoluteAltitudeData *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AbsoluteAltitudeData */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AbsoluteAltitudeData */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AbsoluteAltitudeData */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AbsoluteAltitudeData */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CMAbsoluteAltitudeData */


