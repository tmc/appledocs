// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WKWebExtensionDataRecord */


/* debug [class_header]: Header for WKWebExtensionDataRecord */
// The class instance for the [WebExtensionDataRecord] class.
var (
	WebExtensionDataRecordClass     _WebExtensionDataRecordClass
	WebExtensionDataRecordClassOnce sync.Once
)

func getWebExtensionDataRecordClass() _WebExtensionDataRecordClass {
	WebExtensionDataRecordClassOnce.Do(func() {
		WebExtensionDataRecordClass = _WebExtensionDataRecordClass{objc.GetClass("WKWebExtensionDataRecord")}
	})
	return WebExtensionDataRecordClass
}

type _WebExtensionDataRecordClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for WebExtensionDataRecord */
// An interface definition for the [WebExtensionDataRecord] class.
type IWebExtensionDataRecord interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for WebExtensionDataRecord */
	// properties:
	ContainedDataTypes() unsafe.Pointer
	DisplayName() objc.IObject /* cross-framework: NSString */
	Errors() []objc.IObject /* cross-framework: Error */
	TotalSizeInBytes() uint
	UniqueIdentifier() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for WebExtensionDataRecord */
	// methods:
	SizeInBytesOfTypes(dataTypes unsafe.Pointer) uint
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for WebExtensionDataRecord */
// Alloc allocates a new instance without initialization.
func (wc _WebExtensionDataRecordClass) Alloc() WebExtensionDataRecord {
	rv := objc.Send[WebExtensionDataRecord](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (wc _WebExtensionDataRecordClass) New() WebExtensionDataRecord {
	rv := objc.Send[WebExtensionDataRecord](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WebExtensionDataRecord) Init() WebExtensionDataRecord {
	rv := objc.Send[WebExtensionDataRecord](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WebExtensionDataRecord) Autorelease() WebExtensionDataRecord {
	rv := objc.Send[WebExtensionDataRecord](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWebExtensionDataRecord creates a new WebExtensionDataRecord instance.
func NewWebExtensionDataRecord() WebExtensionDataRecord {
	return getWebExtensionDataRecordClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for WebExtensionDataRecord */
// An object that represents a record of stored data for a specific web extension context.
//
// Contains properties and methods to query the data types and sizes.


// An object that represents a record of stored data for a specific web extension context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/DataRecord
type WebExtensionDataRecord struct {
	objectivec.Object
}

// WebExtensionDataRecordFrom constructs a [WebExtensionDataRecord] from an unsafe.Pointer.
//
// An object that represents a record of stored data for a specific web extension context.
func WebExtensionDataRecordFrom(ptr unsafe.Pointer) WebExtensionDataRecord {
	return WebExtensionDataRecord{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for WebExtensionDataRecord *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for WebExtensionDataRecord */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for WebExtensionDataRecord */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for WebExtensionDataRecord */

// Retrieves the size in bytes of the specific data types in this data record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/DataRecord/sizeInBytes(ofTypes:)
func (w_ WebExtensionDataRecord) SizeInBytesOfTypes(dataTypes unsafe.Pointer) uint {
	rv := objc.Send[uint](w_.ID, objc.Sel("sizeInBytesOfTypes:"), dataTypes)
	return rv
}/* debug [instance_methods/method]: SizeInBytesOfTypes */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for WebExtensionDataRecord */

// The set of data types contained in this data record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/DataRecord/containedDataTypes
func (w_ WebExtensionDataRecord) ContainedDataTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("containedDataTypes"))
	return rv
}/* debug [instance_properties/getter]: containedDataTypes */


// The display name for the web extension to which this data record belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/DataRecord/displayName
func (w_ WebExtensionDataRecord) DisplayName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("displayName"))
	return rv
}/* debug [instance_properties/getter]: displayName */


// An array of errors that may have occurred when either calculating or deleting storage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/DataRecord/errors
func (w_ WebExtensionDataRecord) Errors() []objc.IObject /* cross-framework: Error */ {
	rv := objc.Send[[]coretelephony.Error](w_.ID, objc.Sel("errors"))
	return rv
}/* debug [instance_properties/getter]: errors */


// The total size in bytes of all data types contained in this data record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/DataRecord/totalSizeInBytes
func (w_ WebExtensionDataRecord) TotalSizeInBytes() uint {
	rv := objc.Send[uint](w_.ID, objc.Sel("totalSizeInBytes"))
	return rv
}/* debug [instance_properties/getter]: totalSizeInBytes */


// Unique identifier for the web extension context to which this data record belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/DataRecord/uniqueIdentifier
func (w_ WebExtensionDataRecord) UniqueIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("uniqueIdentifier"))
	return rv
}/* debug [instance_properties/getter]: uniqueIdentifier */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class WKWebExtensionDataRecord */



