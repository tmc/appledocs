// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Predicate] class.
var PredicateClass objc.Class

func init() {
	PredicateClass = objc.GetClass("NSPredicate")
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
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSPredicate/init(block:)
func (pc Predicate) PredicateWithBlock(block unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("predicateWithBlock:")
	ret := objc.ID(PredicateClass).Send(sel, block)
	return unsafe.Pointer(ret)
}
// Creates a predicate by substituting the values in a specified array into a format string and parsing the result. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSPredicate/init(format:argumentArray:)
func (pc Predicate) PredicateWithFormatArgumentArray(predicateFormat string, arguments unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("predicateWithFormat:argumentArray:")
	ret := objc.ID(PredicateClass).Send(sel, predicateFormat, arguments)
	return unsafe.Pointer(ret)
}
// Creates a predicate by substituting the values in an argument list into a format string and parsing the result. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSPredicate/init(format:arguments:)
func (pc Predicate) PredicateWithFormatArguments(predicateFormat string, argList unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("predicateWithFormat:arguments:")
	ret := objc.ID(PredicateClass).Send(sel, predicateFormat, argList)
	return unsafe.Pointer(ret)
}
// Creates a predicate with a metadata query string. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSPredicate/init(fromMetadataQueryString:)
func (pc Predicate) PredicateFromMetadataQueryString(queryString string) unsafe.Pointer {
	sel := objc.RegisterName("predicateFromMetadataQueryString:")
	ret := objc.ID(PredicateClass).Send(sel, queryString)
	return unsafe.Pointer(ret)
}
// Creates and returns a predicate that always evaluates to a specified Boolean value. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSPredicate/init(value:)
func (pc Predicate) PredicateWithValue(value bool) unsafe.Pointer {
	sel := objc.RegisterName("predicateWithValue:")
	ret := objc.ID(PredicateClass).Send(sel, value)
	return unsafe.Pointer(ret)
}
// Creates and returns a new predicate formed by creating a new string with a specified format and parsing the result. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSPredicate/predicateWithFormat:
func (pc Predicate) PredicateWithFormat(predicateFormat string) unsafe.Pointer {
	sel := objc.RegisterName("predicateWithFormat:")
	ret := objc.ID(PredicateClass).Send(sel, predicateFormat)
	return unsafe.Pointer(ret)
}
// Forces a securely decoded predicate to allow evaluation. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSPredicate/allowEvaluation()
func (p_ Predicate) AllowEvaluation() {
	sel := objc.RegisterName("allowEvaluation")
	p_.ID.Send(sel)
}
// Returns a Boolean value that indicates whether the specified object matches the conditions that the predicate specifies. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSPredicate/evaluate(with:)
func (p_ Predicate) EvaluateWithObject(object objc.ID) bool {
	sel := objc.RegisterName("evaluateWithObject:")
	ret := p_.ID.Send(sel, object)
	return ret != 0
}
// Returns a Boolean value that indicates whether the specified object matches the conditions that the predicate specifies after substituting in the values from a specified variables dictionary. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSPredicate/evaluate(with:substitutionVariables:)
func (p_ Predicate) EvaluateWithObjectSubstitutionVariables(object objc.ID, bindings unsafe.Pointer) bool {
	sel := objc.RegisterName("evaluateWithObject:substitutionVariables:")
	ret := p_.ID.Send(sel, object, bindings)
	return ret != 0
}
// Returns a copy of the predicate and substitutes the predicates variables with specified values from a specified substitution variables dictionary. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSPredicate/withSubstitutionVariables(_:)
func (p_ Predicate) PredicateWithSubstitutionVariables(variables unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("predicateWithSubstitutionVariables:")
	ret := p_.ID.Send(sel, variables)
	return unsafe.Pointer(ret)
}

