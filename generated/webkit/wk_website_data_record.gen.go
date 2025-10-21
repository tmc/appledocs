// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [WebsiteDataRecord] class.
type IWebsiteDataRecord interface {
	objectivec.IObject
}

// A record of the data that a particular website stores persistently.
//
// Use objects to discover the types of information that a website stores. Records identify the data types a website stores, but don’t identify the actual data. You might use this information to help the user manage website data. For example, Safari provides a way for users to view and remove website data. The domain name of each record contains the website’s domain name and suffix. You don’t create objects directly. WebKit creates these records and stores them in the web view’s data store. Use the of that data store to retrieve the current record objects. You also use that object to remove unwanted records.
//
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

// Alloc allocates a new instance without initialization.
func (wc _WebsiteDataRecordClass) Alloc() WebsiteDataRecord {
	rv := objc.Send[WebsiteDataRecord](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The display name for the data record.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebsitedatarecord/displayname
func (w_ WebsiteDataRecord) DisplayName() string {
	rv := objc.Send[string](w_.ID, objc.Sel("displayName"))
	return rv
}


// SetDisplayName sets the value of the displayName property.
// The display name for the data record.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebsitedatarecord/displayname
func (w_ WebsiteDataRecord) SetDisplayName(value string) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDisplayName:"), objc.String(value))
}

// The types of data associated with the record.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebsiteDataRecord/dataTypes
func (w_ WebsiteDataRecord) DataTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("dataTypes"))
	return rv
}



