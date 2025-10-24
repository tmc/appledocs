// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CMTremorResult */


/* debug [class_header]: Header for CMTremorResult */
// The class instance for the [TremorResult] class.
var (
	TremorResultClass     _TremorResultClass
	TremorResultClassOnce sync.Once
)

func getTremorResultClass() _TremorResultClass {
	TremorResultClassOnce.Do(func() {
		TremorResultClass = _TremorResultClass{objc.GetClass("CMTremorResult")}
	})
	return TremorResultClass
}

type _TremorResultClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TremorResult */
// An interface definition for the [TremorResult] class.
type ITremorResult interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TremorResult */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TremorResult */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TremorResult */
// Alloc allocates a new instance without initialization.
func (tc _TremorResultClass) Alloc() TremorResult {
	rv := objc.Send[TremorResult](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TremorResultClass) New() TremorResult {
	rv := objc.Send[TremorResult](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TremorResult) Init() TremorResult {
	rv := objc.Send[TremorResult](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TremorResult) Autorelease() TremorResult {
	rv := objc.Send[TremorResult](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTremorResult creates a new TremorResult instance.
func NewTremorResult() TremorResult {
	return getTremorResultClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TremorResult */
// A result object that contains data about the presence and strength of tremors during a one-minute interval.
//
// The following equation is always true: .


// A result object that contains data about the presence and strength of tremors during a one-minute interval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMTremorResult
type TremorResult struct {
	objectivec.Object
}

// TremorResultFrom constructs a [TremorResult] from an unsafe.Pointer.
//
// A result object that contains data about the presence and strength of tremors during a one-minute interval.
func TremorResultFrom(ptr unsafe.Pointer) TremorResult {
	return TremorResult{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TremorResult *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TremorResult */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TremorResult */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TremorResult */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TremorResult */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CMTremorResult */


