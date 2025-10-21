// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ClassDescription] class.
var (
	ClassDescriptionClass     _ClassDescriptionClass
	ClassDescriptionClassOnce sync.Once
)

func getClassDescriptionClass() _ClassDescriptionClass {
	ClassDescriptionClassOnce.Do(func() {
		ClassDescriptionClass = _ClassDescriptionClass{objc.GetClass("NSClassDescription")}
	})
	return ClassDescriptionClass
}

type _ClassDescriptionClass struct {
	class objc.Class
}

// An interface definition for the [ClassDescription] class.
type IClassDescription interface {
	objectivec.IObject
}

// An abstract class that provides the interface for querying the relationships and properties of a class.
//
// Concrete subclasses of provide the available attributes of objects of a particular class and the relationships between that class and other classes. Defining these relationships between classes allows for more intelligent and flexible manipulation of objects with key-value coding. It is important to note that there are no class descriptions by default. To use objects in your code you have to implement them for your model classes. For all concrete subclasses, you must provide implementations for all instance methods of . ( provides only the implementation for the class methods that maintain the cache of registered class descriptions.) Once created, you must register a class description with the method . You can use the objects in the arrays returned by methods such as and to access—using key-value coding—the properties of an instance of the class to which a class description object corresponds. For more about attributes and relationships, see Cocoa Fundamentals Guide. For more about key-value coding, see . , which is used to map the relationships between scriptable classes, is the only concrete subclass of provided as part of the Cocoa framework.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSClassDescription
type ClassDescription struct {
	objectivec.Object
}

// ClassDescriptionFrom constructs a [ClassDescription] from an unsafe.Pointer.
//
// An abstract class that provides the interface for querying the relationships and properties of a class.
func ClassDescriptionFrom(ptr unsafe.Pointer) ClassDescription {
	return ClassDescription{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _ClassDescriptionClass) Alloc() ClassDescription {
	rv := objc.Send[ClassDescription](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ClassDescriptionClass) New() ClassDescription {
	rv := objc.Send[ClassDescription](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ClassDescription) Init() ClassDescription {
	rv := objc.Send[ClassDescription](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ClassDescription) Autorelease() ClassDescription {
	rv := objc.Send[ClassDescription](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewClassDescription creates a new ClassDescription instance.
func NewClassDescription() ClassDescription {
	return getClassDescriptionClass().New()
}

// Returns the class description for a given class.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSClassDescription/init(for:)
func NewClassDescriptionForClass(aClass objc.Class) ClassDescription {
	rv := objc.Send[ClassDescription](objc.ID(getClassDescriptionClass().class), objc.Sel("classDescriptionForClass:"), aClass)
	return rv
}

// Returns the class description for a given class.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSClassDescription/init(for:)
func (cc _ClassDescriptionClass) ClassDescriptionForClass(aClass objc.Class) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("classDescriptionForClass:"), aClass)
	return rv
}
