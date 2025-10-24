// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSMorphologyPronoun */


/* debug [class_header]: Header for NSMorphologyPronoun */
// The class instance for the [MorphologyPronoun] class.
var (
	MorphologyPronounClass     _MorphologyPronounClass
	MorphologyPronounClassOnce sync.Once
)

func getMorphologyPronounClass() _MorphologyPronounClass {
	MorphologyPronounClassOnce.Do(func() {
		MorphologyPronounClass = _MorphologyPronounClass{objc.GetClass("NSMorphologyPronoun")}
	})
	return MorphologyPronounClass
}

type _MorphologyPronounClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MorphologyPronoun */
// An interface definition for the [MorphologyPronoun] class.
type IMorphologyPronoun interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MorphologyPronoun */
	// properties:
	DependentMorphology() IMorphology
	Morphology() IMorphology
	Pronoun() IString
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MorphologyPronoun */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MorphologyPronoun */
// Alloc allocates a new instance without initialization.
func (mc _MorphologyPronounClass) Alloc() MorphologyPronoun {
	rv := objc.Send[MorphologyPronoun](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MorphologyPronounClass) New() MorphologyPronoun {
	rv := objc.Send[MorphologyPronoun](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MorphologyPronoun) Init() MorphologyPronoun {
	rv := objc.Send[MorphologyPronoun](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MorphologyPronoun) Autorelease() MorphologyPronoun {
	rv := objc.Send[MorphologyPronoun](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMorphologyPronoun creates a new MorphologyPronoun instance.
func NewMorphologyPronoun() MorphologyPronoun {
	return getMorphologyPronounClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MorphologyPronoun */
// A custom pronoun for referring to a third person.
//
// Create instances of  when you need to define custom pronouns for a localized term of address. For examples of how to create custom pronouns, see .


// A custom pronoun for referring to a third person.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphologyPronoun
type MorphologyPronoun struct {
	objectivec.Object
}

// MorphologyPronounFrom constructs a [MorphologyPronoun] from an unsafe.Pointer.
//
// A custom pronoun for referring to a third person.
func MorphologyPronounFrom(ptr unsafe.Pointer) MorphologyPronoun {
	return MorphologyPronoun{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MorphologyPronoun */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphologyPronoun/initWithPronoun:morphology:dependentMorphology:
func NewMorphologyPronounWithPronounMorphologyDependentMorphology(pronoun IString, morphology IMorphology, dependentMorphology IMorphology) MorphologyPronoun {
	instance := getMorphologyPronounClass().Alloc()
	rv := objc.Send[MorphologyPronoun](instance.ID, objc.Sel("initWithPronoun:morphology:dependentMorphology:"), pronoun, morphology, dependentMorphology)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMorphologyPronounWithPronounMorphologyDependentMorphology */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MorphologyPronoun */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MorphologyPronoun */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MorphologyPronoun */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MorphologyPronoun */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphologyPronoun/dependentMorphology
func (m_ MorphologyPronoun) DependentMorphology() IMorphology {
	rv := objc.Send[Morphology](m_.ID, objc.Sel("dependentMorphology"))
	return rv
}/* debug [instance_properties/getter]: dependentMorphology */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphologyPronoun/morphology
func (m_ MorphologyPronoun) Morphology() IMorphology {
	rv := objc.Send[Morphology](m_.ID, objc.Sel("morphology"))
	return rv
}/* debug [instance_properties/getter]: morphology */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphologyPronoun/pronoun
func (m_ MorphologyPronoun) Pronoun() IString {
	rv := objc.Send[String](m_.ID, objc.Sel("pronoun"))
	return rv
}/* debug [instance_properties/getter]: pronoun */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSMorphologyPronoun */


