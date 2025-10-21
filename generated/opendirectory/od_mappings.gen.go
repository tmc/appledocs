// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [ODMappings] class.
var (
	ODMappingsClass     _ODMappingsClass
	ODMappingsClassOnce sync.Once
)

func getODMappingsClass() _ODMappingsClass {
	ODMappingsClassOnce.Do(func() {
		ODMappingsClass = _ODMappingsClass{objc.GetClass("ODMappings")}
	})
	return ODMappingsClass
}

type _ODMappingsClass struct {
	class objc.Class
}

// An interface definition for the [ODMappings] class.
type IODMappings interface {
	objectivec.IObject
	RecordMapForStandardRecordType(stdType string) unsafe.Pointer
	SetRecordMapForStandardRecordType(map_ unsafe.Pointer, stdType string)
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings
type ODMappings struct {
	objectivec.Object
}

// ODMappingsFrom constructs a [ODMappings] from an unsafe.Pointer.
func ODMappingsFrom(ptr unsafe.Pointer) ODMappings {
	return ODMappings{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (oc _ODMappingsClass) Alloc() ODMappings {
	rv := objc.Send[ODMappings](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _ODMappingsClass) New() ODMappings {
	rv := objc.Send[ODMappings](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ ODMappings) Init() ODMappings {
	rv := objc.Send[ODMappings](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ ODMappings) Autorelease() ODMappings {
	rv := objc.Send[ODMappings](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewODMappings creates a new ODMappings instance.
func NewODMappings() ODMappings {
	return getODMappingsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/mappings
func (oc _ODMappingsClass) Mappings() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("mappings"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/recordMap(forStandardRecordType:)
func (o_ ODMappings) RecordMapForStandardRecordType(stdType string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("recordMapForStandardRecordType:"), objc.String(stdType))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/setRecordMap(_:forStandardRecordType:)
func (o_ ODMappings) SetRecordMapForStandardRecordType(map_ unsafe.Pointer, stdType string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setRecordMap:forStandardRecordType:"), map_, objc.String(stdType))
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/comment-swift.property
func (o_ ODMappings) Comment() string {
	rv := objc.Send[string](o_.ID, objc.Sel("comment"))
	return rv
}


// SetComment sets the value of the comment property.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/comment-swift.property
func (o_ ODMappings) SetComment(value string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setComment:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/function-swift.property
func (o_ ODMappings) Function() string {
	rv := objc.Send[string](o_.ID, objc.Sel("function"))
	return rv
}


// SetFunction sets the value of the function property.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/function-swift.property
func (o_ ODMappings) SetFunction(value string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setFunction:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/functionAttributes-swift.property
func (o_ ODMappings) FunctionAttributes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("functionAttributes"))
	return rv
}


// SetFunctionAttributes sets the value of the functionAttributes property.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/functionAttributes-swift.property
func (o_ ODMappings) SetFunctionAttributes(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setFunctionAttributes:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/identifier-swift.property
func (o_ ODMappings) Identifier() string {
	rv := objc.Send[string](o_.ID, objc.Sel("identifier"))
	return rv
}


// SetIdentifier sets the value of the identifier property.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/identifier-swift.property
func (o_ ODMappings) SetIdentifier(value string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setIdentifier:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/recordTypes-swift.property
func (o_ ODMappings) RecordTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("recordTypes"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/templateName-swift.property
func (o_ ODMappings) TemplateName() string {
	rv := objc.Send[string](o_.ID, objc.Sel("templateName"))
	return rv
}


// SetTemplateName sets the value of the templateName property.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/templateName-swift.property
func (o_ ODMappings) SetTemplateName(value string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setTemplateName:"), objc.String(value))
}



