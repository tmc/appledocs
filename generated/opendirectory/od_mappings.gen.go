// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
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
	RecordMapForStandardRecordType(stdType appkit.string) ODRecordMap
	SetRecordMapForStandardRecordType(map_ IODRecordMap, stdType appkit.string)
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
func (o_ ODMappings) RecordMapForStandardRecordType(stdType appkit.string) ODRecordMap {
	rv := objc.Send[ODRecordMap](o_.ID, objc.Sel("recordMapForStandardRecordType:"), stdType)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/setRecordMap(_:forStandardRecordType:)
func (o_ ODMappings) SetRecordMapForStandardRecordType(map_ IODRecordMap, stdType appkit.string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setRecordMap:forStandardRecordType:"), map_, stdType)
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/comment-swift.property
func (o_ ODMappings) Comment() appkit.string {
	rv := objc.Send[appkit.string](o_.ID, objc.Sel("comment"))
	return rv
}


// SetComment sets the value of the comment property.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/comment-swift.property
func (o_ ODMappings) SetComment(value appkit.string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setComment:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/function-swift.property
func (o_ ODMappings) Function() appkit.string {
	rv := objc.Send[appkit.string](o_.ID, objc.Sel("function"))
	return rv
}


// SetFunction sets the value of the function property.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/function-swift.property
func (o_ ODMappings) SetFunction(value appkit.string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setFunction:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/functionAttributes-swift.property
func (o_ ODMappings) FunctionAttributes() objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("functionAttributes"))
	return rv
}


// SetFunctionAttributes sets the value of the functionAttributes property.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/functionAttributes-swift.property
func (o_ ODMappings) SetFunctionAttributes(value objc.ID) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setFunctionAttributes:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/identifier-swift.property
func (o_ ODMappings) Identifier() appkit.string {
	rv := objc.Send[appkit.string](o_.ID, objc.Sel("identifier"))
	return rv
}


// SetIdentifier sets the value of the identifier property.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/identifier-swift.property
func (o_ ODMappings) SetIdentifier(value appkit.string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setIdentifier:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/recordTypes-swift.property
func (o_ ODMappings) RecordTypes() objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("recordTypes"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/templateName-swift.property
func (o_ ODMappings) TemplateName() appkit.string {
	rv := objc.Send[appkit.string](o_.ID, objc.Sel("templateName"))
	return rv
}


// SetTemplateName sets the value of the templateName property.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/templateName-swift.property
func (o_ ODMappings) SetTemplateName(value appkit.string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setTemplateName:"), value)
}



