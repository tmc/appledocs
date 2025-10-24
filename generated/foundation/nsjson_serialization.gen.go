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
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (jc _JSONSerializationClass) Alloc() JSONSerialization {
	rv := objc.Send[JSONSerialization](objc.ID(jc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// An object that converts between JSON and the equivalent Foundation objects.
//
// You use the class to convert JSON to Foundation objects and convert Foundation objects to JSON. To convert a Foundation object to JSON, the object must have the following properties: The top level object is an or , unless you set the option. All objects are instances of , , , , or . All dictionary keys are instances of . Numbers are neither nor infinity. Other rules may apply. Calling or attempting a conversion are the definitive ways to tell if the class can convert given object to JSON data.


// An object that converts between JSON and the equivalent Foundation objects.
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










// Returns JSON data from a Foundation object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/JSONSerialization/data(withJSONObject:options:)
func (jc _JSONSerializationClass) DataWithJSONObjectOptionsError(obj objc.IObject, opt JSONWritingOptions, error_ IError) IData {
	rv := objc.Send[Data](objc.ID(jc.class), objc.Sel("dataWithJSONObject:options:error:"), obj, opt, error_)
	return rv
}


// Returns a Boolean value that indicates whether the serializer can convert a given object to JSON data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/JSONSerialization/isValidJSONObject(_:)
func (jc _JSONSerializationClass) IsValidJSONObject(obj objc.IObject) bool {
	rv := objc.Send[bool](objc.ID(jc.class), objc.Sel("isValidJSONObject:"), obj)
	return rv
}


// Returns a Foundation object from JSON data in a given stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/JSONSerialization/jsonObject(with:options:)-3afap
func (jc _JSONSerializationClass) JSONObjectWithStreamOptionsError(stream IInputStream, opt JSONReadingOptions, error_ IError) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(jc.class), objc.Sel("JSONObjectWithStream:options:error:"), stream, opt, error_)
	return rv
}


// Returns a Foundation object from given JSON data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/JSONSerialization/jsonObject(with:options:)-8demi
func (jc _JSONSerializationClass) JSONObjectWithDataOptionsError(data IData, opt JSONReadingOptions, error_ IError) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(jc.class), objc.Sel("JSONObjectWithData:options:error:"), data, opt, error_)
	return rv
}


// Writes a given JSON object to a stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/JSONSerialization/writeJSONObject(_:to:options:error:)
func (jc _JSONSerializationClass) WriteJSONObjectToStreamOptionsError(obj objc.IObject, stream IOutputStream, opt JSONWritingOptions, error_ IError) int {
	rv := objc.Send[int](objc.ID(jc.class), objc.Sel("writeJSONObject:toStream:options:error:"), obj, stream, opt, error_)
	return rv
}























