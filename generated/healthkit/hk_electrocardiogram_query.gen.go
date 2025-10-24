// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKElectrocardiogramQuery */


/* debug [class_header]: Header for HKElectrocardiogramQuery */
// The class instance for the [HKElectrocardiogramQuery] class.
var (
	HKElectrocardiogramQueryClass     _HKElectrocardiogramQueryClass
	HKElectrocardiogramQueryClassOnce sync.Once
)

func getHKElectrocardiogramQueryClass() _HKElectrocardiogramQueryClass {
	HKElectrocardiogramQueryClassOnce.Do(func() {
		HKElectrocardiogramQueryClass = _HKElectrocardiogramQueryClass{objc.GetClass("HKElectrocardiogramQuery")}
	})
	return HKElectrocardiogramQueryClass
}

type _HKElectrocardiogramQueryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKElectrocardiogramQuery */
// An interface definition for the [HKElectrocardiogramQuery] class.
type IHKElectrocardiogramQuery interface {
	IHKQuery
	
/* debug [class_interface_properties]: Properties for HKElectrocardiogramQuery */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKElectrocardiogramQuery */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKElectrocardiogramQuery */
// Alloc allocates a new instance without initialization.
func (hc _HKElectrocardiogramQueryClass) Alloc() HKElectrocardiogramQuery {
	rv := objc.Send[HKElectrocardiogramQuery](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKElectrocardiogramQueryClass) New() HKElectrocardiogramQuery {
	rv := objc.Send[HKElectrocardiogramQuery](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKElectrocardiogramQuery) Init() HKElectrocardiogramQuery {
	rv := objc.Send[HKElectrocardiogramQuery](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKElectrocardiogramQuery) Autorelease() HKElectrocardiogramQuery {
	rv := objc.Send[HKElectrocardiogramQuery](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKElectrocardiogramQuery creates a new HKElectrocardiogramQuery instance.
func NewHKElectrocardiogramQuery() HKElectrocardiogramQuery {
	return getHKElectrocardiogramQueryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKElectrocardiogramQuery */
// A query that returns the underlying voltage measurements for an electrocardiogram sample.
//
// Use the query to access the individual voltage measurements associated with an sample. The query calls the data handler once for each voltage measurement, passing a instance that contains the voltage data. After it has sent all the voltage measurements, the query calls the data handler one last time, passing . If an error occurs, it stops collecting voltage data and passes instead. Electrocardiogram queries are immutable: You set query’s properties when you create it, and they don’t change.


// A query that returns the underlying voltage measurements for an electrocardiogram sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKElectrocardiogramQuery
type HKElectrocardiogramQuery struct {
	HKQuery
}

// HKElectrocardiogramQueryFrom constructs a [HKElectrocardiogramQuery] from an unsafe.Pointer.
//
// A query that returns the underlying voltage measurements for an electrocardiogram sample.
func HKElectrocardiogramQueryFrom(ptr unsafe.Pointer) HKElectrocardiogramQuery {
	return HKElectrocardiogramQuery{
		HKQuery: HKQueryFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKElectrocardiogramQuery */

// Creates a new electrocardiogram query object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKElectrocardiogramQuery/init(electrocardiogram:dataHandler:)
func NewHKElectrocardiogramQueryWithElectrocardiogramDataHandler(electrocardiogram IHKElectrocardiogram, dataHandler unsafe.Pointer) HKElectrocardiogramQuery {
	instance := getHKElectrocardiogramQueryClass().Alloc()
	rv := objc.Send[HKElectrocardiogramQuery](instance.ID, objc.Sel("initWithElectrocardiogram:dataHandler:"), electrocardiogram, dataHandler)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKElectrocardiogramQueryWithElectrocardiogramDataHandler */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKElectrocardiogramQuery */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKElectrocardiogramQuery */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKElectrocardiogramQuery */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKElectrocardiogramQuery */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKElectrocardiogramQuery */


