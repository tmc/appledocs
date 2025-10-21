// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
	Option(optionName string) objc.ID
	SetOptionValue(optionName string, value objc.ID)
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
func NewODModuleEntryWithNameXpcServiceName(name string, xpcServiceName string) ODModuleEntry {
	rv := objc.Send[ODModuleEntry](objc.ID(getODModuleEntryClass().class), objc.Sel("moduleEntryWithName:xpcServiceName:"), objc.String(name), objc.String(xpcServiceName))
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/init(name:xpcServiceName:)
func (oc _ODModuleEntryClass) ModuleEntryWithNameXpcServiceName(name string, xpcServiceName string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("moduleEntryWithName:xpcServiceName:"), objc.String(name), objc.String(xpcServiceName))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/option(_:)
func (o_ ODModuleEntry) Option(optionName string) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("option:"), objc.String(optionName))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/setOption(_:value:)
func (o_ ODModuleEntry) SetOptionValue(optionName string, value objc.ID) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setOption:value:"), objc.String(optionName), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/mappings-swift.property
func (o_ ODModuleEntry) Mappings() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("mappings"))
	return rv
}


// SetMappings sets the value of the mappings property.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/mappings-swift.property
func (o_ ODModuleEntry) SetMappings(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setMappings:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/name-swift.property
func (o_ ODModuleEntry) Name() string {
	rv := objc.Send[string](o_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/name-swift.property
func (o_ ODModuleEntry) SetName(value string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setName:"), objc.String(value))
}
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/supportedOptions-swift.property
func (o_ ODModuleEntry) SupportedOptions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("supportedOptions"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/uuidString-swift.property
func (o_ ODModuleEntry) UuidString() string {
	rv := objc.Send[string](o_.ID, objc.Sel("uuidString"))
	return rv
}


// SetUuidString sets the value of the uuidString property.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/uuidString-swift.property
func (o_ ODModuleEntry) SetUuidString(value string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setUuidString:"), objc.String(value))
}
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/xpcServiceName-swift.property
func (o_ ODModuleEntry) XpcServiceName() string {
	rv := objc.Send[string](o_.ID, objc.Sel("xpcServiceName"))
	return rv
}


// SetXpcServiceName sets the value of the xpcServiceName property.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/xpcServiceName-swift.property
func (o_ ODModuleEntry) SetXpcServiceName(value string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setXpcServiceName:"), objc.String(value))
}

