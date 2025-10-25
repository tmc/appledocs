// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSValueTransformer */


/* debug [class_header]: Header for NSValueTransformer */
// The class instance for the [ValueTransformer] class.
var (
	ValueTransformerClass     _ValueTransformerClass
	ValueTransformerClassOnce sync.Once
)

func getValueTransformerClass() _ValueTransformerClass {
	ValueTransformerClassOnce.Do(func() {
		ValueTransformerClass = _ValueTransformerClass{objc.GetClass("NSValueTransformer")}
	})
	return ValueTransformerClass
}

type _ValueTransformerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ValueTransformer */
// An interface definition for the [ValueTransformer] class.
type IValueTransformer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ValueTransformer */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ValueTransformer */
	// methods:
	ReverseTransformedValue(value objc.IObject) objc.ID
	TransformedValue(value objc.IObject) objc.ID
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ValueTransformer */
// Alloc allocates a new instance without initialization.
func (vc _ValueTransformerClass) Alloc() ValueTransformer {
	rv := objc.Send[ValueTransformer](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _ValueTransformerClass) New() ValueTransformer {
	rv := objc.Send[ValueTransformer](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ ValueTransformer) Init() ValueTransformer {
	rv := objc.Send[ValueTransformer](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ ValueTransformer) Autorelease() ValueTransformer {
	rv := objc.Send[ValueTransformer](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewValueTransformer creates a new ValueTransformer instance.
func NewValueTransformer() ValueTransformer {
	return getValueTransformerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ValueTransformer */
// An abstract class used to transform values from one representation to another.
//
// You create a value transformer by subclassing and overriding the necessary methods to provide the required custom transformation. You then register the value transformer using the method, so that other parts of your app can access it by name with . Use the method to transform a value from one representation into another. If a value transformer designates that its transformation is reversible by returning for , you can also use the to perform the transformation in reverse. For example, reversing the characters in a string is a reversible operation, whereas changing the characters in a string to be uppercase is a nonreversible operation. A value transformer can take inputs of one type and return a value of a different type. For example, a value transformer could take an or object and return an object containing the PNG representation of that image.


// An abstract class used to transform values from one representation to another.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ValueTransformer
type ValueTransformer struct {
	objectivec.Object
}

// ValueTransformerFrom constructs a [ValueTransformer] from an unsafe.Pointer.
//
// An abstract class used to transform values from one representation to another.
func ValueTransformerFrom(ptr unsafe.Pointer) ValueTransformer {
	return ValueTransformer{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ValueTransformer */

// Returns the value transformer identified by a given identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ValueTransformer/init(forName:)
func NewValueTransformerForName(name ValueTransformerName) ValueTransformer {
	rv := objc.Send[ValueTransformer](objc.ID(getValueTransformerClass().class), objc.Sel("valueTransformerForName:"), name)
	return rv
}/* debug [class_init_methods/constructor]: NewValueTransformerForName */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ValueTransformer */

// Returns a Boolean value that indicates whether the receiver can reverse a transformation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ValueTransformer/allowsReverseTransformation()
func (vc _ValueTransformerClass) AllowsReverseTransformation() bool {
	rv := objc.Send[bool](objc.ID(vc.class), objc.Sel("allowsReverseTransformation"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AllowsReverseTransformation) */


// Returns the value transformer identified by a given identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ValueTransformer/init(forName:)
func (vc _ValueTransformerClass) ValueTransformerForName(name ValueTransformerName) IValueTransformer {
	rv := objc.Send[ValueTransformer](objc.ID(vc.class), objc.Sel("valueTransformerForName:"), name)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ValueTransformerForName) */


// Registers the provided value transformer with a given identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ValueTransformer/setValueTransformer(_:forName:)
func (vc _ValueTransformerClass) SetValueTransformerForName(transformer IValueTransformer, name ValueTransformerName) {
	objc.Send[objc.ID](objc.ID(vc.class), objc.Sel("setValueTransformer:forName:"), transformer, name)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SetValueTransformerForName) */


// Returns the class of the value returned by the receiver for a forward transformation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ValueTransformer/transformedValueClass()
func (vc _ValueTransformerClass) TransformedValueClass() objc.Class {
	rv := objc.Send[objc.Class](objc.ID(vc.class), objc.Sel("transformedValueClass"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TransformedValueClass) */


// Returns an array of all the registered value transformers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ValueTransformer/valueTransformerNames()
func (vc _ValueTransformerClass) ValueTransformerNames() []string {
	rv := objc.Send[[]string](objc.ID(vc.class), objc.Sel("valueTransformerNames"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ValueTransformerNames) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ValueTransformer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ValueTransformer */

// Returns the result of the reverse transformation of a given value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ValueTransformer/reverseTransformedValue(_:)
func (v_ ValueTransformer) ReverseTransformedValue(value objc.IObject) objc.ID {
	rv := objc.Send[objc.ID](v_.ID, objc.Sel("reverseTransformedValue:"), value)
	return rv
}/* debug [instance_methods/method]: ReverseTransformedValue */


// Returns the result of transforming a given value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ValueTransformer/transformedValue(_:)
func (v_ ValueTransformer) TransformedValue(value objc.IObject) objc.ID {
	rv := objc.Send[objc.ID](v_.ID, objc.Sel("transformedValue:"), value)
	return rv
}/* debug [instance_methods/method]: TransformedValue */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ValueTransformer */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSValueTransformer */


