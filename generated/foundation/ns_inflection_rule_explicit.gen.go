// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSInflectionRuleExplicit */


/* debug [class_header]: Header for NSInflectionRuleExplicit */
// The class instance for the [InflectionRuleExplicit] class.
var (
	InflectionRuleExplicitClass     _InflectionRuleExplicitClass
	InflectionRuleExplicitClassOnce sync.Once
)

func getInflectionRuleExplicitClass() _InflectionRuleExplicitClass {
	InflectionRuleExplicitClassOnce.Do(func() {
		InflectionRuleExplicitClass = _InflectionRuleExplicitClass{objc.GetClass("NSInflectionRuleExplicit")}
	})
	return InflectionRuleExplicitClass
}

type _InflectionRuleExplicitClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for InflectionRuleExplicit */
// An interface definition for the [InflectionRuleExplicit] class.
type IInflectionRuleExplicit interface {
	IInflectionRule
	
/* debug [class_interface_properties]: Properties for InflectionRuleExplicit */
	// properties:
	Morphology() IMorphology
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for InflectionRuleExplicit */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for InflectionRuleExplicit */
// Alloc allocates a new instance without initialization.
func (ic _InflectionRuleExplicitClass) Alloc() InflectionRuleExplicit {
	rv := objc.Send[InflectionRuleExplicit](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _InflectionRuleExplicitClass) New() InflectionRuleExplicit {
	rv := objc.Send[InflectionRuleExplicit](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ InflectionRuleExplicit) Init() InflectionRuleExplicit {
	rv := objc.Send[InflectionRuleExplicit](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ InflectionRuleExplicit) Autorelease() InflectionRuleExplicit {
	rv := objc.Send[InflectionRuleExplicit](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewInflectionRuleExplicit creates a new InflectionRuleExplicit instance.
func NewInflectionRuleExplicit() InflectionRuleExplicit {
	return getInflectionRuleExplicitClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for InflectionRuleExplicit */
// An inflection rule that uses a morphology instance to determine how to inflect attribued strings.


// An inflection rule that uses a morphology instance to determine how to inflect attribued strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInflectionRuleExplicit
type InflectionRuleExplicit struct {
	InflectionRule
}

// InflectionRuleExplicitFrom constructs a [InflectionRuleExplicit] from an unsafe.Pointer.
//
// An inflection rule that uses a morphology instance to determine how to inflect attribued strings.
func InflectionRuleExplicitFrom(ptr unsafe.Pointer) InflectionRuleExplicit {
	return InflectionRuleExplicit{
		InflectionRule: InflectionRuleFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for InflectionRuleExplicit */

// Creates an inflection rule with the given morphology.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInflectionRuleExplicit/initWithMorphology:
func NewInflectionRuleExplicitWithMorphology(morphology IMorphology) InflectionRuleExplicit {
	instance := getInflectionRuleExplicitClass().Alloc()
	rv := objc.Send[InflectionRuleExplicit](instance.ID, objc.Sel("initWithMorphology:"), morphology)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewInflectionRuleExplicitWithMorphology */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for InflectionRuleExplicit */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for InflectionRuleExplicit */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for InflectionRuleExplicit */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for InflectionRuleExplicit */

// The morphology used by this inflection rule.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInflectionRuleExplicit/morphology
func (i_ InflectionRuleExplicit) Morphology() IMorphology {
	rv := objc.Send[Morphology](i_.ID, objc.Sel("morphology"))
	return rv
}/* debug [instance_properties/getter]: morphology */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSInflectionRuleExplicit */


