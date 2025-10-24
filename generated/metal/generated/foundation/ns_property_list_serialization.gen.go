// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PropertyListSerialization] class.
var (
	PropertyListSerializationClass     _PropertyListSerializationClass
	PropertyListSerializationClassOnce sync.Once
)

func getPropertyListSerializationClass() _PropertyListSerializationClass {
	PropertyListSerializationClassOnce.Do(func() {
		PropertyListSerializationClass = _PropertyListSerializationClass{objc.GetClass("NSPropertyListSerialization")}
	})
	return PropertyListSerializationClass
}

type _PropertyListSerializationClass struct {
	class objc.Class
}

// An interface definition for the [PropertyListSerialization] class.
type IPropertyListSerialization interface {
	objectivec.IObject
	// properties:
	NSPropertyListErrorMaximum() int
	SetNSPropertyListErrorMaximum(value int)
	NSPropertyListErrorMinimum() int
	SetNSPropertyListErrorMinimum(value int)
	NSPropertyListReadCorruptError() int
	SetNSPropertyListReadCorruptError(value int)
	NSPropertyListReadStreamError() int
	SetNSPropertyListReadStreamError(value int)
	NSPropertyListReadUnknownVersionError() int
	SetNSPropertyListReadUnknownVersionError(value int)
	NSPropertyListWriteInvalidError() int
	SetNSPropertyListWriteInvalidError(value int)
	NSPropertyListWriteStreamError() int
	SetNSPropertyListWriteStreamError(value int)
	// methods:
}

// An object that converts between a property list and one of several serialized representations.
//
// The class provides methods that convert a property list to and from several serialized formats. A property list is itself an array or dictionary that contains only , , , , , and objects. Property list objects are toll-free bridged with their respective Core Foundation types ( , , and so on). See for more information on toll-free bridging.


// An object that converts between a property list and one of several serialized representations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/PropertyListSerialization
type PropertyListSerialization struct {
	objectivec.Object
}

// PropertyListSerializationFrom constructs a [PropertyListSerialization] from an unsafe.Pointer.
//
// An object that converts between a property list and one of several serialized representations.
func PropertyListSerializationFrom(ptr unsafe.Pointer) PropertyListSerialization {
	return PropertyListSerialization{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PropertyListSerializationClass) Alloc() PropertyListSerialization {
	rv := objc.Send[PropertyListSerialization](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PropertyListSerializationClass) New() PropertyListSerialization {
	rv := objc.Send[PropertyListSerialization](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PropertyListSerialization) Init() PropertyListSerialization {
	rv := objc.Send[PropertyListSerialization](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PropertyListSerialization) Autorelease() PropertyListSerialization {
	rv := objc.Send[PropertyListSerialization](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPropertyListSerialization creates a new PropertyListSerialization instance.
func NewPropertyListSerialization() PropertyListSerialization {
	return getPropertyListSerializationClass().New()
}



// The end of the range of error codes reserved for property list errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspropertylisterrormaximum-swift.var
func (p_ PropertyListSerialization) NSPropertyListErrorMaximum() int {
	rv := objc.Send[int](p_.ID, objc.Sel("NSPropertyListErrorMaximum"))
	return rv
}


// The end of the range of error codes reserved for property list errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspropertylisterrormaximum-swift.var
func (p_ PropertyListSerialization) SetNSPropertyListErrorMaximum(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNSPropertyListErrorMaximum:"), value)
}


// The start of the range of error codes reserved for property list errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspropertylisterrorminimum-swift.var
func (p_ PropertyListSerialization) NSPropertyListErrorMinimum() int {
	rv := objc.Send[int](p_.ID, objc.Sel("NSPropertyListErrorMinimum"))
	return rv
}


// The start of the range of error codes reserved for property list errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspropertylisterrorminimum-swift.var
func (p_ PropertyListSerialization) SetNSPropertyListErrorMinimum(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNSPropertyListErrorMinimum:"), value)
}


// Parsing of the property list failed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspropertylistreadcorrupterror-swift.var
func (p_ PropertyListSerialization) NSPropertyListReadCorruptError() int {
	rv := objc.Send[int](p_.ID, objc.Sel("NSPropertyListReadCorruptError"))
	return rv
}


// Parsing of the property list failed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspropertylistreadcorrupterror-swift.var
func (p_ PropertyListSerialization) SetNSPropertyListReadCorruptError(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNSPropertyListReadCorruptError:"), value)
}


// Reading of the property list failed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspropertylistreadstreamerror-swift.var
func (p_ PropertyListSerialization) NSPropertyListReadStreamError() int {
	rv := objc.Send[int](p_.ID, objc.Sel("NSPropertyListReadStreamError"))
	return rv
}


// Reading of the property list failed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspropertylistreadstreamerror-swift.var
func (p_ PropertyListSerialization) SetNSPropertyListReadStreamError(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNSPropertyListReadStreamError:"), value)
}


// The version number of the property list cannot be determined.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspropertylistreadunknownversionerror-swift.var
func (p_ PropertyListSerialization) NSPropertyListReadUnknownVersionError() int {
	rv := objc.Send[int](p_.ID, objc.Sel("NSPropertyListReadUnknownVersionError"))
	return rv
}


// The version number of the property list cannot be determined.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspropertylistreadunknownversionerror-swift.var
func (p_ PropertyListSerialization) SetNSPropertyListReadUnknownVersionError(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNSPropertyListReadUnknownVersionError:"), value)
}


// Writing failed because of an invalid property list object, or an invalid property list type was specified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspropertylistwriteinvaliderror-swift.var
func (p_ PropertyListSerialization) NSPropertyListWriteInvalidError() int {
	rv := objc.Send[int](p_.ID, objc.Sel("NSPropertyListWriteInvalidError"))
	return rv
}


// Writing failed because of an invalid property list object, or an invalid property list type was specified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspropertylistwriteinvaliderror-swift.var
func (p_ PropertyListSerialization) SetNSPropertyListWriteInvalidError(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNSPropertyListWriteInvalidError:"), value)
}


// Writing to the property list failed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspropertylistwritestreamerror-swift.var
func (p_ PropertyListSerialization) NSPropertyListWriteStreamError() int {
	rv := objc.Send[int](p_.ID, objc.Sel("NSPropertyListWriteStreamError"))
	return rv
}


// Writing to the property list failed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspropertylistwritestreamerror-swift.var
func (p_ PropertyListSerialization) SetNSPropertyListWriteStreamError(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNSPropertyListWriteStreamError:"), value)
}



