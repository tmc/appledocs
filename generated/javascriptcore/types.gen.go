// Code generated from Apple documentation for JavaScriptCore. DO NOT EDIT.

package javascriptcore
import (
	"unsafe"
)


// C struct types
// JSClassDefinition - A structure that contains properties and callbacks that define a type of object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSClassDefinition
type JSClassDefinition struct {
	Attributes JSClassAttributes // A set of class attributes to give to the class.
	CallAsConstructor JSObjectCallAsConstructorCallback // The callback for using an object as a constructor.
	CallAsFunction JSObjectCallAsFunctionCallback // The callback for calling an object as a function.
	ClassName unsafe.Pointer // A null-terminated UTF-8 string that contains the class’s name.
	ConvertToType JSObjectConvertToTypeCallback // The callback for converting an object to a particular JavaScript type.
	DeleteProperty JSObjectDeletePropertyCallback // The callback for deleting a property.
	Finalize JSObjectFinalizeCallback // The callback for preparing the object for garbage collection.
	GetProperty JSObjectGetPropertyCallback // The callback for getting a property’s value.
	GetPropertyNames JSObjectGetPropertyNamesCallback // The callback for collecting the names of an object’s properties.
	HasInstance JSObjectHasInstanceCallback // The callback for checking whether an object is an instance of a particular type.
	HasProperty JSObjectHasPropertyCallback // The callback for determining whether an object has a property.
	Initialize JSObjectInitializeCallback // The callback for creating the object.
	ParentClass JSClassRef // A JavaScript class to set as the class’s parent class.
	SetProperty JSObjectSetPropertyCallback // The callback for setting a property’s value.
	StaticFunctions JSStaticFunction // An array that contains the class’s statically declared function properties.
	StaticValues JSStaticValue // An array that contains the class’s statically declared value properties.
	Version int // The version of the class definition structure.
}/* debug [types.gen.go/struct]: JSClassDefinition */

// JSStaticFunction - A statically declared function property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSStaticFunction
type JSStaticFunction struct {
	Attributes JSPropertyAttributes // A set of property attributes to give to the property.
	CallAsFunction JSObjectCallAsFunctionCallback // A callback to invoke when calling the property as a function.
	Name unsafe.Pointer // A null-terminated UTF-8 string that contains the property’s name.
}/* debug [types.gen.go/struct]: JSStaticFunction */

// JSStaticValue - A statically declared value property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSStaticValue
type JSStaticValue struct {
	Attributes JSPropertyAttributes // A set of property attributes to give to the property.
	GetProperty JSObjectGetPropertyCallback // A callback to invoke when getting the property’s value.
	Name unsafe.Pointer // A null-terminated UTF-8 string that contains the property’s name.
	SetProperty JSObjectSetPropertyCallback // A callback to invoke when setting the property’s value.
}/* debug [types.gen.go/struct]: JSStaticValue */





