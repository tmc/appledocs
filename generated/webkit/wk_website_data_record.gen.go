// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WKWebsiteDataRecord */


/* debug [class_header]: Header for WKWebsiteDataRecord */
// The class instance for the [WebsiteDataRecord] class.
var (
	WebsiteDataRecordClass     _WebsiteDataRecordClass
	WebsiteDataRecordClassOnce sync.Once
)

func getWebsiteDataRecordClass() _WebsiteDataRecordClass {
	WebsiteDataRecordClassOnce.Do(func() {
		WebsiteDataRecordClass = _WebsiteDataRecordClass{objc.GetClass("WKWebsiteDataRecord")}
	})
	return WebsiteDataRecordClass
}

type _WebsiteDataRecordClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for WebsiteDataRecord */
// An interface definition for the [WebsiteDataRecord] class.
type IWebsiteDataRecord interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for WebsiteDataRecord */
	// properties:
	DataTypes() unsafe.Pointer
	DisplayName() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for WebsiteDataRecord */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for WebsiteDataRecord */
// Alloc allocates a new instance without initialization.
func (wc _WebsiteDataRecordClass) Alloc() WebsiteDataRecord {
	rv := objc.Send[WebsiteDataRecord](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (wc _WebsiteDataRecordClass) New() WebsiteDataRecord {
	rv := objc.Send[WebsiteDataRecord](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WebsiteDataRecord) Init() WebsiteDataRecord {
	rv := objc.Send[WebsiteDataRecord](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WebsiteDataRecord) Autorelease() WebsiteDataRecord {
	rv := objc.Send[WebsiteDataRecord](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWebsiteDataRecord creates a new WebsiteDataRecord instance.
func NewWebsiteDataRecord() WebsiteDataRecord {
	return getWebsiteDataRecordClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for WebsiteDataRecord */
// A record of the data that a particular website stores persistently.
//
// Use objects to discover the types of information that a website stores. Records identify the data types a website stores, but don’t identify the actual data. You might use this information to help the user manage website data. For example, Safari provides a way for users to view and remove website data. The domain name of each record contains the website’s domain name and suffix. You don’t create objects directly. WebKit creates these records and stores them in the web view’s data store. Use the of that data store to retrieve the current record objects. You also use that object to remove unwanted records.


// A record of the data that a particular website stores persistently.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebsiteDataRecord
type WebsiteDataRecord struct {
	objectivec.Object
}

// WebsiteDataRecordFrom constructs a [WebsiteDataRecord] from an unsafe.Pointer.
//
// A record of the data that a particular website stores persistently.
func WebsiteDataRecordFrom(ptr unsafe.Pointer) WebsiteDataRecord {
	return WebsiteDataRecord{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for WebsiteDataRecord *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for WebsiteDataRecord */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for WebsiteDataRecord */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for WebsiteDataRecord */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for WebsiteDataRecord */

// The types of data associated with the record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebsiteDataRecord/dataTypes
func (w_ WebsiteDataRecord) DataTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("dataTypes"))
	return rv
}/* debug [instance_properties/getter]: dataTypes */


// The display name for the data record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebsiteDataRecord/displayName
func (w_ WebsiteDataRecord) DisplayName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("displayName"))
	return rv
}/* debug [instance_properties/getter]: displayName */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class WKWebsiteDataRecord */



