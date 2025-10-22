// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [ValueTransformer] class.
type IValueTransformer interface {
	objectivec.IObject
	ReverseTransformedValue(value objectivec.IObject) objc.ID
	TransformedValue(value objectivec.IObject) objc.ID
}

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

// Alloc allocates a new instance without initialization.
func (vc _ValueTransformerClass) Alloc() ValueTransformer {
	rv := objc.Send[ValueTransformer](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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





// Returns the value transformer identified by a given identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ValueTransformer/init(forName:)

func NewValueTransformerForName(name IValueTransformerName) ValueTransformer {
	rv := objc.Send[ValueTransformer](objc.ID(getValueTransformerClass().class), objc.Sel("valueTransformerForName:"), name)
	return rv
}



// Returns a Boolean value that indicates whether the receiver can reverse a transformation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ValueTransformer/allowsReverseTransformation()

func (vc _ValueTransformerClass) AllowsReverseTransformation() bool {
	rv := objc.Send[bool](objc.ID(vc.class), objc.Sel("allowsReverseTransformation"))
	return rv
}


// Returns the value transformer identified by a given identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ValueTransformer/init(forName:)

func (vc _ValueTransformerClass) ValueTransformerForName(name IValueTransformerName) ValueTransformer {
	rv := objc.Send[ValueTransformer](objc.ID(vc.class), objc.Sel("valueTransformerForName:"), name)
	return rv
}


// Registers the provided value transformer with a given identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ValueTransformer/setValueTransformer(_:forName:)

func (vc _ValueTransformerClass) SetValueTransformerForName(transformer IValueTransformer, name IValueTransformerName) {
	objc.Send[objc.ID](objc.ID(vc.class), objc.Sel("setValueTransformer:forName:"), transformer, name)
}


// Returns the class of the value returned by the receiver for a forward transformation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ValueTransformer/transformedValueClass()

func (vc _ValueTransformerClass) TransformedValueClass() objc.Class {
	rv := objc.Send[objc.Class](objc.ID(vc.class), objc.Sel("transformedValueClass"))
	return rv
}


// Returns an array of all the registered value transformers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ValueTransformer/valueTransformerNames()

func (vc _ValueTransformerClass) ValueTransformerNames() []string {
	rv := objc.Send[[]string](objc.ID(vc.class), objc.Sel("valueTransformerNames"))
	return rv
}


// Returns the result of the reverse transformation of a given value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ValueTransformer/reverseTransformedValue(_:)

func (v_ ValueTransformer) ReverseTransformedValue(value objectivec.IObject) objc.ID {
	rv := objc.Send[objc.ID](v_.ID, objc.Sel("reverseTransformedValue:"), value)
	return rv
}


// Returns the result of transforming a given value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ValueTransformer/transformedValue(_:)

func (v_ ValueTransformer) TransformedValue(value objectivec.IObject) objc.ID {
	rv := objc.Send[objc.ID](v_.ID, objc.Sel("transformedValue:"), value)
	return rv
}


