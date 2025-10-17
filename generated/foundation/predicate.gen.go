// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Predicate] class.
var PredicateClass = _PredicateClass{objc.GetClass("NSPredicate")}

type _PredicateClass struct {
	class objc.Class
}

type Predicate struct {
	objc.ID
}

func PredicateFrom(ptr unsafe.Pointer) Predicate {
	return Predicate{
		ID: objc.ID(ptr),
	}
}


// Creates a predicate that evaluates using a specified block object and bindings dictionary. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPredicate/init(block:)
func (pc _PredicateClass) PredicateWithBlock(block unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("predicateWithBlock:"), block)
	return rv
}
// Creates a predicate by substituting the values in a specified array into a format string and parsing the result. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPredicate/init(format:argumentArray:)
func (pc _PredicateClass) PredicateWithFormatArgumentArray(predicateFormat string, arguments unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("predicateWithFormat:argumentArray:"), predicateFormat, arguments)
	return rv
}
// Creates a predicate by substituting the values in an argument list into a format string and parsing the result. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPredicate/init(format:arguments:)
func (pc _PredicateClass) PredicateWithFormatArguments(predicateFormat string, argList unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("predicateWithFormat:arguments:"), predicateFormat, argList)
	return rv
}
// Creates a predicate with a metadata query string. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPredicate/init(fromMetadataQueryString:)
func (pc _PredicateClass) PredicateFromMetadataQueryString(queryString string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("predicateFromMetadataQueryString:"), queryString)
	return rv
}
// Creates and returns a predicate that always evaluates to a specified Boolean value. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPredicate/init(value:)
func (pc _PredicateClass) PredicateWithValue(value bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("predicateWithValue:"), value)
	return rv
}
// Creates and returns a new predicate formed by creating a new string with a specified format and parsing the result. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPredicate/predicateWithFormat:
func (pc _PredicateClass) PredicateWithFormat(predicateFormat string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("predicateWithFormat:"), predicateFormat)
	return rv
}
// Forces a securely decoded predicate to allow evaluation. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPredicate/allowEvaluation()
func (p_ Predicate) AllowEvaluation() {
	objc.Send[objc.ID](p_.ID, objc.Sel("allowEvaluation"))
}
// Returns a Boolean value that indicates whether the specified object matches the conditions that the predicate specifies. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPredicate/evaluate(with:)
func (p_ Predicate) EvaluateWithObject(object objc.ID) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("evaluateWithObject:"), object)
	return rv
}
// Returns a Boolean value that indicates whether the specified object matches the conditions that the predicate specifies after substituting in the values from a specified variables dictionary. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPredicate/evaluate(with:substitutionVariables:)
func (p_ Predicate) EvaluateWithObjectSubstitutionVariables(object objc.ID, bindings unsafe.Pointer) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("evaluateWithObject:substitutionVariables:"), object, bindings)
	return rv
}
// Returns a copy of the predicate and substitutes the predicates variables with specified values from a specified substitution variables dictionary. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPredicate/withSubstitutionVariables(_:)
func (p_ Predicate) PredicateWithSubstitutionVariables(variables unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("predicateWithSubstitutionVariables:"), variables)
	return rv
}


