// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ODModuleEntry */


/* debug [class_header]: Header for ODModuleEntry */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ODModuleEntry */
// An interface definition for the [ODModuleEntry] class.
type IODModuleEntry interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ODModuleEntry */
	// properties:
	Mappings() IODMappings
	SetMappings(value IODMappings)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	SupportedOptions() objc.IObject /* cross-framework: NSArray */
	UuidString() objc.IObject /* cross-framework: NSString */
	SetUuidString(value objc.IObject /* cross-framework: NSString */)
	XpcServiceName() objc.IObject /* cross-framework: NSString */
	SetXpcServiceName(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ODModuleEntry */
	// methods:
	Option(optionName objc.IObject /* cross-framework: NSString */) objc.ID
	SetOptionValue(optionName objc.IObject /* cross-framework: NSString */, value objc.IObject)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ODModuleEntry */
// Alloc allocates a new instance without initialization.
func (oc _ODModuleEntryClass) Alloc() ODModuleEntry {
	rv := objc.Send[ODModuleEntry](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ODModuleEntry */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry
type ODModuleEntry struct {
	objectivec.Object
}

// ODModuleEntryFrom constructs a [ODModuleEntry] from an unsafe.Pointer.
func ODModuleEntryFrom(ptr unsafe.Pointer) ODModuleEntry {
	return ODModuleEntry{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ODModuleEntry */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/init(name:xpcServiceName:)
func NewODModuleEntryWithNameXpcServiceName(name objc.IObject /* cross-framework: NSString */, xpcServiceName objc.IObject /* cross-framework: NSString */) ODModuleEntry {
	rv := objc.Send[ODModuleEntry](objc.ID(getODModuleEntryClass().class), objc.Sel("moduleEntryWithName:xpcServiceName:"), name, xpcServiceName)
	return rv
}/* debug [class_init_methods/constructor]: NewODModuleEntryWithNameXpcServiceName */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ODModuleEntry */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/init(name:xpcServiceName:)
func (oc _ODModuleEntryClass) ModuleEntryWithNameXpcServiceName(name objc.IObject /* cross-framework: NSString */, xpcServiceName objc.IObject /* cross-framework: NSString */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("moduleEntryWithName:xpcServiceName:"), name, xpcServiceName)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ModuleEntryWithNameXpcServiceName) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ODModuleEntry */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ODModuleEntry */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/option(_:)
func (o_ ODModuleEntry) Option(optionName objc.IObject /* cross-framework: NSString */) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("option:"), optionName)
	return rv
}/* debug [instance_methods/method]: Option */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/setOption(_:value:)
func (o_ ODModuleEntry) SetOptionValue(optionName objc.IObject /* cross-framework: NSString */, value objc.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setOption:value:"), optionName, value)
}/* debug [instance_methods/method]: SetOptionValue */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ODModuleEntry */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/mappings-swift.property
func (o_ ODModuleEntry) Mappings() IODMappings {
	rv := objc.Send[ODMappings](o_.ID, objc.Sel("mappings"))
	return rv
}/* debug [instance_properties/getter]: mappings */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/mappings-swift.property
func (o_ ODModuleEntry) SetMappings(value IODMappings) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setMappings:"), value)
}/* debug [instance_properties/setter]: mappings */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/name-swift.property
func (o_ ODModuleEntry) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/name-swift.property
func (o_ ODModuleEntry) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/supportedOptions-swift.property
func (o_ ODModuleEntry) SupportedOptions() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](o_.ID, objc.Sel("supportedOptions"))
	return rv
}/* debug [instance_properties/getter]: supportedOptions */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/uuidString-swift.property
func (o_ ODModuleEntry) UuidString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("uuidString"))
	return rv
}/* debug [instance_properties/getter]: uuidString */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/uuidString-swift.property
func (o_ ODModuleEntry) SetUuidString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setUuidString:"), value)
}/* debug [instance_properties/setter]: uuidString */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/xpcServiceName-swift.property
func (o_ ODModuleEntry) XpcServiceName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("xpcServiceName"))
	return rv
}/* debug [instance_properties/getter]: xpcServiceName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/xpcServiceName-swift.property
func (o_ ODModuleEntry) SetXpcServiceName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setXpcServiceName:"), value)
}/* debug [instance_properties/setter]: xpcServiceName */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ODModuleEntry */


