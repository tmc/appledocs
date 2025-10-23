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
	NSPropertyListErrorMaximum() int /* primitive/slice/pointer. */
	SetNSPropertyListErrorMaximum(value int /* primitive/slice/pointer. */)
	NSPropertyListErrorMinimum() int /* primitive/slice/pointer. */
	SetNSPropertyListErrorMinimum(value int /* primitive/slice/pointer. */)
	NSPropertyListReadCorruptError() int /* primitive/slice/pointer. */
	SetNSPropertyListReadCorruptError(value int /* primitive/slice/pointer. */)
	NSPropertyListReadStreamError() int /* primitive/slice/pointer. */
	SetNSPropertyListReadStreamError(value int /* primitive/slice/pointer. */)
	NSPropertyListReadUnknownVersionError() int /* primitive/slice/pointer. */
	SetNSPropertyListReadUnknownVersionError(value int /* primitive/slice/pointer. */)
	NSPropertyListWriteInvalidError() int /* primitive/slice/pointer. */
	SetNSPropertyListWriteInvalidError(value int /* primitive/slice/pointer. */)
	NSPropertyListWriteStreamError() int /* primitive/slice/pointer. */
	SetNSPropertyListWriteStreamError(value int /* primitive/slice/pointer. */)
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



// Returns an object containing a given property list in a specified format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/PropertyListSerialization/data(fromPropertyList:format:options:)
func (pc _PropertyListSerializationClass) DataWithPropertyListFormatOptionsError(plist objectivec.IObject, format PropertyListFormat, opt objc.IObject /* cross-framework PropertyListWriteOptions */, error_ IError) IData {
	rv := objc.Send[Data](objc.ID(pc.class), objc.Sel("dataWithPropertyList:format:options:error:"), plist, format, opt, error_)
	return rv
}


// This method is obsolete and will be deprecated soon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/PropertyListSerialization/dataFromPropertyList(_:format:errorDescription:)
func (pc _PropertyListSerializationClass) DataFromPropertyListFormatErrorDescription(plist objectivec.IObject, format PropertyListFormat, errorString IString) IData {
	rv := objc.Send[Data](objc.ID(pc.class), objc.Sel("dataFromPropertyList:format:errorDescription:"), plist, format, errorString)
	return rv
}


// Returns a Boolean value that indicates whether a given property list is valid for a given format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/PropertyListSerialization/propertyList(_:isValidFor:)
func (pc _PropertyListSerializationClass) PropertyListIsValidForFormat(plist objectivec.IObject, format PropertyListFormat) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](objc.ID(pc.class), objc.Sel("propertyList:isValidForFormat:"), plist, format)
	return rv
}


// Creates and returns a property list from the specified data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/PropertyListSerialization/propertyList(from:options:format:)
func (pc _PropertyListSerializationClass) PropertyListWithDataOptionsFormatError(data IData, opt objc.IObject /* cross-framework PropertyListReadOptions */, format PropertyListFormat, error_ IError) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(pc.class), objc.Sel("propertyListWithData:options:format:error:"), data, opt, format, error_)
	return rv
}


// Creates and returns a property list by reading from the specified stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/PropertyListSerialization/propertyList(with:options:format:)
func (pc _PropertyListSerializationClass) PropertyListWithStreamOptionsFormatError(stream IInputStream, opt objc.IObject /* cross-framework PropertyListReadOptions */, format PropertyListFormat, error_ IError) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(pc.class), objc.Sel("propertyListWithStream:options:format:error:"), stream, opt, format, error_)
	return rv
}


// This method is deprecated. Use instead.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/PropertyListSerialization/propertyListFromData(_:mutabilityOption:format:errorDescription:)
func (pc _PropertyListSerializationClass) PropertyListFromDataMutabilityOptionFormatErrorDescription(data IData, opt PropertyListMutabilityOptions, format PropertyListFormat, errorString IString) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(pc.class), objc.Sel("propertyListFromData:mutabilityOption:format:errorDescription:"), data, opt, format, errorString)
	return rv
}


// Writes a property list to the specified stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/PropertyListSerialization/writePropertyList(_:to:format:options:error:)
func (pc _PropertyListSerializationClass) WritePropertyListToStreamFormatOptionsError(plist objectivec.IObject, stream IOutputStream, format PropertyListFormat, opt objc.IObject /* cross-framework PropertyListWriteOptions */, error_ IError) int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](objc.ID(pc.class), objc.Sel("writePropertyList:toStream:format:options:error:"), plist, stream, format, opt, error_)
	return rv
}


// The end of the range of error codes reserved for property list errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspropertylisterrormaximum-swift.var
func (p_ PropertyListSerialization) NSPropertyListErrorMaximum() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](p_.ID, objc.Sel("NSPropertyListErrorMaximum"))
	return rv
}


// The end of the range of error codes reserved for property list errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspropertylisterrormaximum-swift.var
func (p_ PropertyListSerialization) SetNSPropertyListErrorMaximum(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNSPropertyListErrorMaximum:"), value)
}


// The start of the range of error codes reserved for property list errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspropertylisterrorminimum-swift.var
func (p_ PropertyListSerialization) NSPropertyListErrorMinimum() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](p_.ID, objc.Sel("NSPropertyListErrorMinimum"))
	return rv
}


// The start of the range of error codes reserved for property list errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspropertylisterrorminimum-swift.var
func (p_ PropertyListSerialization) SetNSPropertyListErrorMinimum(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNSPropertyListErrorMinimum:"), value)
}


// Parsing of the property list failed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspropertylistreadcorrupterror-swift.var
func (p_ PropertyListSerialization) NSPropertyListReadCorruptError() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](p_.ID, objc.Sel("NSPropertyListReadCorruptError"))
	return rv
}


// Parsing of the property list failed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspropertylistreadcorrupterror-swift.var
func (p_ PropertyListSerialization) SetNSPropertyListReadCorruptError(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNSPropertyListReadCorruptError:"), value)
}


// Reading of the property list failed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspropertylistreadstreamerror-swift.var
func (p_ PropertyListSerialization) NSPropertyListReadStreamError() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](p_.ID, objc.Sel("NSPropertyListReadStreamError"))
	return rv
}


// Reading of the property list failed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspropertylistreadstreamerror-swift.var
func (p_ PropertyListSerialization) SetNSPropertyListReadStreamError(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNSPropertyListReadStreamError:"), value)
}


// The version number of the property list cannot be determined.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspropertylistreadunknownversionerror-swift.var
func (p_ PropertyListSerialization) NSPropertyListReadUnknownVersionError() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](p_.ID, objc.Sel("NSPropertyListReadUnknownVersionError"))
	return rv
}


// The version number of the property list cannot be determined.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspropertylistreadunknownversionerror-swift.var
func (p_ PropertyListSerialization) SetNSPropertyListReadUnknownVersionError(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNSPropertyListReadUnknownVersionError:"), value)
}


// Writing failed because of an invalid property list object, or an invalid property list type was specified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspropertylistwriteinvaliderror-swift.var
func (p_ PropertyListSerialization) NSPropertyListWriteInvalidError() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](p_.ID, objc.Sel("NSPropertyListWriteInvalidError"))
	return rv
}


// Writing failed because of an invalid property list object, or an invalid property list type was specified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspropertylistwriteinvaliderror-swift.var
func (p_ PropertyListSerialization) SetNSPropertyListWriteInvalidError(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNSPropertyListWriteInvalidError:"), value)
}


// Writing to the property list failed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspropertylistwritestreamerror-swift.var
func (p_ PropertyListSerialization) NSPropertyListWriteStreamError() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](p_.ID, objc.Sel("NSPropertyListWriteStreamError"))
	return rv
}


// Writing to the property list failed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspropertylistwritestreamerror-swift.var
func (p_ PropertyListSerialization) SetNSPropertyListWriteStreamError(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNSPropertyListWriteStreamError:"), value)
}



