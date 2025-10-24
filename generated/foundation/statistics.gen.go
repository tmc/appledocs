// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class statistics */


/* debug [class_header]: Header for statistics */
// The class instance for the [statistics] class.
var (
	StatisticsClass     _statisticsClass
	StatisticsClassOnce sync.Once
)

func getstatisticsClass() _statisticsClass {
	StatisticsClassOnce.Do(func() {
		StatisticsClass = _statisticsClass{objc.GetClass("statistics")}
	})
	return StatisticsClass
}

type _statisticsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for statistics */
// An interface definition for the [statistics] class.
type Istatistics interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for statistics */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for statistics */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for statistics */
// Alloc allocates a new instance without initialization.
func (sc _statisticsClass) Alloc() statistics {
	rv := objc.Send[statistics](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _statisticsClass) New() statistics {
	rv := objc.Send[statistics](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ statistics) Init() statistics {
	rv := objc.Send[statistics](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ statistics) Autorelease() statistics {
	rv := objc.Send[statistics](s_.ID, objc.Sel("autorelease"))
	return rv
}

// Newstatistics creates a new statistics instance.
func Newstatistics() statistics {
	return getstatisticsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for statistics */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/statistics-c.ivar
type statistics struct {
	objectivec.Object
}

// statisticsFrom constructs a [statistics] from an unsafe.Pointer.
func statisticsFrom(ptr unsafe.Pointer) statistics {
	return statistics{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for statistics *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for statistics */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for statistics */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for statistics */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for statistics */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class statistics */



