// Code generated from Apple documentation for ObjectiveC. DO NOT EDIT.

package objectivec
import (
	"unsafe"
)



// C struct types
// NXHashState
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXHashState
type NXHashState struct {
	I int
	J int
}// NXHashTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXHashTable
type NXHashTable struct {
	Buckets unsafe.Pointer
	Count unsafe.Pointer
	Info unsafe.Pointer
	NbBuckets unsafe.Pointer
	Prototype unsafe.Pointer
}// NXHashTablePrototype
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXHashTablePrototype
type NXHashTablePrototype struct {
	Free unsafe.Pointer
	Hash unsafe.Pointer
	IsEqual unsafe.Pointer
	Style int
}// objc_method_description - Defines an Objective-C method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_method_description
type objc_method_description struct {
	Name unsafe.Pointer // The name of the method at runtime.
	Types unsafe.Pointer // The types of the method arguments.
}// objc_object - Represents an instance of a class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_object
type objc_object struct {
	Isa unsafe.Pointer // A pointer to the class definition of which this object is an instance.
}// objc_property_attribute_t - Defines a property attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_property_attribute_t
type objc_property_attribute_t struct {
	Name unsafe.Pointer // The name of the attribute.
	Value unsafe.Pointer // The value of the attribute (usually empty).
}// objc_super - Specifies the superclass of an instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_super-swift.struct
type objc_super struct {
	Receiver unsafe.Pointer // A pointer of type  . Specifies an instance of a class.
	Super_class unsafe.Pointer // A pointer to a   data structure. Specifies the particular superclass of the instance to message.
}



