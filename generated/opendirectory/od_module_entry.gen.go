// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ODModuleEntry] class.
var (
	ODModuleEntryClass     _ODModuleEntryClass
	ODModuleEntryClassOnce sync.Once
)

func getODModuleEntryClass() _ODModuleEntryClass {
	ODModuleEntryClassOnce.Do(func() {
		ODModuleEntryClass = _ODModuleEntryClass{objc.GetClass("ODModuleEntry")}
	})
	return ODModuleEntryClass
}

type _ODModuleEntryClass struct {
	class objc.Class
}

// An interface definition for the [ODModuleEntry] class.
type IODModuleEntry interface {
	objectivec.IObject
	Option(optionName appkit.string) objc.ID
	SetOptionValue(optionName appkit.string, value objectivec.IObject)
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry
type ODModuleEntry struct {
	objectivec.Object
}

// ODModuleEntryFrom constructs a [ODModuleEntry] from an unsafe.Pointer.
func ODModuleEntryFrom(ptr unsafe.Pointer) ODModuleEntry {
	return ODModuleEntry{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (oc _ODModuleEntryClass) Alloc() ODModuleEntry {
	rv := objc.Send[ODModuleEntry](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _ODModuleEntryClass) New() ODModuleEntry {
	rv := objc.Send[ODModuleEntry](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ ODModuleEntry) Init() ODModuleEntry {
	rv := objc.Send[ODModuleEntry](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ ODModuleEntry) Autorelease() ODModuleEntry {
	rv := objc.Send[ODModuleEntry](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewODModuleEntry creates a new ODModuleEntry instance.
func NewODModuleEntry() ODModuleEntry {
	return getODModuleEntryClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/init(name:xpcServiceName:)
func NewODModuleEntryWithNameXpcServiceName(name appkit.string, xpcServiceName appkit.string) ODModuleEntry {
	rv := objc.Send[ODModuleEntry](objc.ID(getODModuleEntryClass().class), objc.Sel("moduleEntryWithName:xpcServiceName:"), name, xpcServiceName)
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/init(name:xpcServiceName:)
func (oc _ODModuleEntryClass) ModuleEntryWithNameXpcServiceName(name appkit.string, xpcServiceName appkit.string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("moduleEntryWithName:xpcServiceName:"), name, xpcServiceName)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/option(_:)
func (o_ ODModuleEntry) Option(optionName appkit.string) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("option:"), optionName)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/setOption(_:value:)
func (o_ ODModuleEntry) SetOptionValue(optionName appkit.string, value objectivec.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setOption:value:"), optionName, value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/mappings-swift.property
func (o_ ODModuleEntry) Mappings() ODMappings {
	rv := objc.Send[ODMappings](o_.ID, objc.Sel("mappings"))
	return rv
}


// SetMappings sets the value of the mappings property.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/mappings-swift.property
func (o_ ODModuleEntry) SetMappings(value IODMappings) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setMappings:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/name-swift.property
func (o_ ODModuleEntry) Name() appkit.string {
	rv := objc.Send[appkit.string](o_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/name-swift.property
func (o_ ODModuleEntry) SetName(value appkit.string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setName:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/supportedOptions-swift.property
func (o_ ODModuleEntry) SupportedOptions() objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("supportedOptions"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/uuidString-swift.property
func (o_ ODModuleEntry) UuidString() appkit.string {
	rv := objc.Send[appkit.string](o_.ID, objc.Sel("uuidString"))
	return rv
}


// SetUuidString sets the value of the uuidString property.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/uuidString-swift.property
func (o_ ODModuleEntry) SetUuidString(value appkit.string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setUuidString:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/xpcServiceName-swift.property
func (o_ ODModuleEntry) XpcServiceName() appkit.string {
	rv := objc.Send[appkit.string](o_.ID, objc.Sel("xpcServiceName"))
	return rv
}


// SetXpcServiceName sets the value of the xpcServiceName property.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/xpcServiceName-swift.property
func (o_ ODModuleEntry) SetXpcServiceName(value appkit.string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setXpcServiceName:"), value)
}


