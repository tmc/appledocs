// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSPredicate */


/* debug [class_header]: Header for NSPredicate */
// The class instance for the [Predicate] class.
var (
	PredicateClass     _PredicateClass
	PredicateClassOnce sync.Once
)

func getPredicateClass() _PredicateClass {
	PredicateClassOnce.Do(func() {
		PredicateClass = _PredicateClass{objc.GetClass("NSPredicate")}
	})
	return PredicateClass
}

type _PredicateClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Predicate */
// An interface definition for the [Predicate] class.
type IPredicate interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Predicate */
	// properties:
	PredicateFormat() IString
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Predicate */
	// methods:
	AllowEvaluation()
	EvaluateWithObject(object objc.IObject) bool
	EvaluateWithObjectSubstitutionVariables(object objc.IObject, bindings IDictionary) bool
	PredicateWithSubstitutionVariables(variables IDictionary) objectivec.IObject
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Predicate */
// Alloc allocates a new instance without initialization.
func (pc _PredicateClass) Alloc() Predicate {
	rv := objc.Send[Predicate](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PredicateClass) New() Predicate {
	rv := objc.Send[Predicate](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ Predicate) Init() Predicate {
	rv := objc.Send[Predicate](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ Predicate) Autorelease() Predicate {
	rv := objc.Send[Predicate](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPredicate creates a new Predicate instance.
func NewPredicate() Predicate {
	return getPredicateClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Predicate */
// A definition of logical conditions for constraining a search for a fetch or for in-memory filtering.
//
// Predicates represent logical conditions, which you can use to filter collections of objects. Although it’s common to create predicates directly from instances of , , and , you often create predicates from a format string that the class methods parse on . Examples of predicate format strings include: Simple comparisons, such as or Case- and diacritic-insensitive lookups, such as Logical operations, such as Temporal range constraints, such as Relational conditions, such as Aggregate operations, such as For a complete syntax reference, refer to the . You can also create predicates that include variables using the method so that you can predefine the predicate before substituting concrete values at runtime.


// A definition of logical conditions for constraining a search for a fetch or for in-memory filtering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPredicate
type Predicate struct {
	objectivec.Object
}

// PredicateFrom constructs a [Predicate] from an unsafe.Pointer.
//
// A definition of logical conditions for constraining a search for a fetch or for in-memory filtering.
func PredicateFrom(ptr unsafe.Pointer) Predicate {
	return Predicate{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Predicate */

// Creates a predicate with a metadata query string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPredicate/init(fromMetadataQueryString:)
func NewPredicateFromMetadataQueryString(queryString IString) Predicate {
	rv := objc.Send[Predicate](objc.ID(getPredicateClass().class), objc.Sel("predicateFromMetadataQueryString:"), queryString)
	return rv
}/* debug [class_init_methods/constructor]: NewPredicateFromMetadataQueryString */


// Creates a predicate that evaluates using a specified block object and bindings dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPredicate/init(block:)
func NewPredicateWithBlock(block unsafe.Pointer) Predicate {
	rv := objc.Send[Predicate](objc.ID(getPredicateClass().class), objc.Sel("predicateWithBlock:"), block)
	return rv
}/* debug [class_init_methods/constructor]: NewPredicateWithBlock */


// Creates a predicate by substituting the values in a specified array into a format string and parsing the result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPredicate/init(format:argumentArray:)
func NewPredicateWithFormatArgumentArray(predicateFormat IString, arguments IArray) Predicate {
	rv := objc.Send[Predicate](objc.ID(getPredicateClass().class), objc.Sel("predicateWithFormat:argumentArray:"), predicateFormat, arguments)
	return rv
}/* debug [class_init_methods/constructor]: NewPredicateWithFormatArgumentArray */


// Creates a predicate by substituting the values in an argument list into a format string and parsing the result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPredicate/init(format:arguments:)
func NewPredicateWithFormatArguments(predicateFormat IString, argList objectivec.IObject) Predicate {
	rv := objc.Send[Predicate](objc.ID(getPredicateClass().class), objc.Sel("predicateWithFormat:arguments:"), predicateFormat, argList)
	return rv
}/* debug [class_init_methods/constructor]: NewPredicateWithFormatArguments */


// Creates and returns a predicate that always evaluates to a specified Boolean value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPredicate/init(value:)
func NewPredicateWithValue(value bool) Predicate {
	rv := objc.Send[Predicate](objc.ID(getPredicateClass().class), objc.Sel("predicateWithValue:"), value)
	return rv
}/* debug [class_init_methods/constructor]: NewPredicateWithValue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Predicate */

// Creates a predicate that evaluates using a specified block object and bindings dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPredicate/init(block:)
func (pc _PredicateClass) PredicateWithBlock(block unsafe.Pointer) IPredicate {
	rv := objc.Send[Predicate](objc.ID(pc.class), objc.Sel("predicateWithBlock:"), block)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateWithBlock) */


// Creates a predicate by substituting the values in a specified array into a format string and parsing the result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPredicate/init(format:argumentArray:)
func (pc _PredicateClass) PredicateWithFormatArgumentArray(predicateFormat IString, arguments IArray) IPredicate {
	rv := objc.Send[Predicate](objc.ID(pc.class), objc.Sel("predicateWithFormat:argumentArray:"), predicateFormat, arguments)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateWithFormatArgumentArray) */


// Creates a predicate by substituting the values in an argument list into a format string and parsing the result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPredicate/init(format:arguments:)
func (pc _PredicateClass) PredicateWithFormatArguments(predicateFormat IString, argList objectivec.IObject) IPredicate {
	rv := objc.Send[Predicate](objc.ID(pc.class), objc.Sel("predicateWithFormat:arguments:"), predicateFormat, argList)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateWithFormatArguments) */


// Creates a predicate with a metadata query string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPredicate/init(fromMetadataQueryString:)
func (pc _PredicateClass) PredicateFromMetadataQueryString(queryString IString) IPredicate {
	rv := objc.Send[Predicate](objc.ID(pc.class), objc.Sel("predicateFromMetadataQueryString:"), queryString)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateFromMetadataQueryString) */


// Creates and returns a predicate that always evaluates to a specified Boolean value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPredicate/init(value:)
func (pc _PredicateClass) PredicateWithValue(value bool) IPredicate {
	rv := objc.Send[Predicate](objc.ID(pc.class), objc.Sel("predicateWithValue:"), value)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateWithValue) */


// Creates and returns a new predicate formed by creating a new string with a specified format and parsing the result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPredicate/predicateWithFormat:
func (pc _PredicateClass) PredicateWithFormat(predicateFormat IString) IPredicate {
	rv := objc.Send[Predicate](objc.ID(pc.class), objc.Sel("predicateWithFormat:"), predicateFormat)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateWithFormat) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Predicate */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Predicate */

// Forces a securely decoded predicate to allow evaluation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPredicate/allowEvaluation()
func (p_ Predicate) AllowEvaluation() {
	objc.Send[objc.ID](p_.ID, objc.Sel("allowEvaluation"))
}/* debug [instance_methods/method]: AllowEvaluation */


// Returns a Boolean value that indicates whether the specified object matches the conditions that the predicate specifies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPredicate/evaluate(with:)
func (p_ Predicate) EvaluateWithObject(object objc.IObject) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("evaluateWithObject:"), object)
	return rv
}/* debug [instance_methods/method]: EvaluateWithObject */


// Returns a Boolean value that indicates whether the specified object matches the conditions that the predicate specifies after substituting in the values from a specified variables dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPredicate/evaluate(with:substitutionVariables:)
func (p_ Predicate) EvaluateWithObjectSubstitutionVariables(object objc.IObject, bindings IDictionary) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("evaluateWithObject:substitutionVariables:"), object, bindings)
	return rv
}/* debug [instance_methods/method]: EvaluateWithObjectSubstitutionVariables */


// Returns a copy of the predicate and substitutes the predicates variables with specified values from a specified substitution variables dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPredicate/withSubstitutionVariables(_:)
func (p_ Predicate) PredicateWithSubstitutionVariables(variables IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("predicateWithSubstitutionVariables:"), variables)
	return rv
}/* debug [instance_methods/method]: PredicateWithSubstitutionVariables */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Predicate */

// The predicate’s format string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPredicate/predicateFormat
func (p_ Predicate) PredicateFormat() IString {
	rv := objc.Send[String](p_.ID, objc.Sel("predicateFormat"))
	return rv
}/* debug [instance_properties/getter]: predicateFormat */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSPredicate */



