// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [JSONSerialization] class.
var (
	JSONSerializationClass     _JSONSerializationClass
	JSONSerializationClassOnce sync.Once
)

func getJSONSerializationClass() _JSONSerializationClass {
	JSONSerializationClassOnce.Do(func() {
		JSONSerializationClass = _JSONSerializationClass{objc.GetClass("NSJSONSerialization")}
	})
	return JSONSerializationClass
}

type _JSONSerializationClass struct {
	class objc.Class
}

// An interface definition for the [JSONSerialization] class.
type IJSONSerialization interface {
	objectivec.IObject
}

// An object that converts between JSON and the equivalent Foundation objects.
//
// You use the class to convert JSON to Foundation objects and convert Foundation objects to JSON. To convert a Foundation object to JSON, the object must have the following properties: The top level object is an or , unless you set the option. All objects are instances of , , , , or . All dictionary keys are instances of . Numbers are neither nor infinity. Other rules may apply. Calling or attempting a conversion are the definitive ways to tell if the class can convert given object to JSON data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/JSONSerialization
type JSONSerialization struct {
	objectivec.Object
}

// JSONSerializationFrom constructs a [JSONSerialization] from an unsafe.Pointer.
//
// An object that converts between JSON and the equivalent Foundation objects.
func JSONSerializationFrom(ptr unsafe.Pointer) JSONSerialization {
	return JSONSerialization{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (jc _JSONSerializationClass) Alloc() JSONSerialization {
	rv := objc.Send[JSONSerialization](objc.ID(jc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (jc _JSONSerializationClass) New() JSONSerialization {
	rv := objc.Send[JSONSerialization](objc.ID(jc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (j_ JSONSerialization) Init() JSONSerialization {
	rv := objc.Send[JSONSerialization](j_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (j_ JSONSerialization) Autorelease() JSONSerialization {
	rv := objc.Send[JSONSerialization](j_.ID, objc.Sel("autorelease"))
	return rv
}

// NewJSONSerialization creates a new JSONSerialization instance.
func NewJSONSerialization() JSONSerialization {
	return getJSONSerializationClass().New()
}



// Returns a Boolean value that indicates whether the serializer can convert a given object to JSON data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/JSONSerialization/isValidJSONObject(_:)

func (jc _JSONSerializationClass) IsValidJSONObject(obj objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(jc.class), objc.Sel("isValidJSONObject:"), obj)
	return rv
}



