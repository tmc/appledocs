// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CMDyskineticSymptomResult */


/* debug [class_header]: Header for CMDyskineticSymptomResult */
// The class instance for the [DyskineticSymptomResult] class.
var (
	DyskineticSymptomResultClass     _DyskineticSymptomResultClass
	DyskineticSymptomResultClassOnce sync.Once
)

func getDyskineticSymptomResultClass() _DyskineticSymptomResultClass {
	DyskineticSymptomResultClassOnce.Do(func() {
		DyskineticSymptomResultClass = _DyskineticSymptomResultClass{objc.GetClass("CMDyskineticSymptomResult")}
	})
	return DyskineticSymptomResultClass
}

type _DyskineticSymptomResultClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DyskineticSymptomResult */
// An interface definition for the [DyskineticSymptomResult] class.
type IDyskineticSymptomResult interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for DyskineticSymptomResult */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DyskineticSymptomResult */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DyskineticSymptomResult */
// Alloc allocates a new instance without initialization.
func (dc _DyskineticSymptomResultClass) Alloc() DyskineticSymptomResult {
	rv := objc.Send[DyskineticSymptomResult](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DyskineticSymptomResultClass) New() DyskineticSymptomResult {
	rv := objc.Send[DyskineticSymptomResult](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DyskineticSymptomResult) Init() DyskineticSymptomResult {
	rv := objc.Send[DyskineticSymptomResult](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DyskineticSymptomResult) Autorelease() DyskineticSymptomResult {
	rv := objc.Send[DyskineticSymptomResult](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDyskineticSymptomResult creates a new DyskineticSymptomResult instance.
func NewDyskineticSymptomResult() DyskineticSymptomResult {
	return getDyskineticSymptomResultClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DyskineticSymptomResult */
// A result object that contains data about the likely presence of dyskinetic symptoms during a one-minute interval.
//
// Dyskinesias are uncontrolled, involuntary movements that occur as a side effect of taking Levadopa to control Parkinson’s disease. Dyskinesias can manifest in a single body part, such as the arm, leg, or head, or they can affect the entire body. Particular dyskinesias resemble actions like fidgeting, writhing, wriggling, head bobbing, or body swaying. These symptoms tend to occur during the drug’s peak dosage. Dyskinesias typically occur in patients with advanced Parkinson’s disease, who may require higher dosages of Levadopa. The following equation is always true: .


// A result object that contains data about the likely presence of dyskinetic symptoms during a one-minute interval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMDyskineticSymptomResult
type DyskineticSymptomResult struct {
	objectivec.Object
}

// DyskineticSymptomResultFrom constructs a [DyskineticSymptomResult] from an unsafe.Pointer.
//
// A result object that contains data about the likely presence of dyskinetic symptoms during a one-minute interval.
func DyskineticSymptomResultFrom(ptr unsafe.Pointer) DyskineticSymptomResult {
	return DyskineticSymptomResult{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DyskineticSymptomResult *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DyskineticSymptomResult */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DyskineticSymptomResult */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DyskineticSymptomResult */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DyskineticSymptomResult */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CMDyskineticSymptomResult */


