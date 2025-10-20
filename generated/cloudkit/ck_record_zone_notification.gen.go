// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CKRecordZoneNotification] class.
var (
	CKRecordZoneNotificationClass     _CKRecordZoneNotificationClass
	CKRecordZoneNotificationClassOnce sync.Once
)

func getCKRecordZoneNotificationClass() _CKRecordZoneNotificationClass {
	CKRecordZoneNotificationClassOnce.Do(func() {
		CKRecordZoneNotificationClass = _CKRecordZoneNotificationClass{objc.GetClass("CKRecordZoneNotification")}
	})
	return CKRecordZoneNotificationClass
}

type _CKRecordZoneNotificationClass struct {
	class objc.Class
}

// An interface definition for the [CKRecordZoneNotification] class.
type ICKRecordZoneNotification interface {
	ICKNotification
}

// A notification that triggers when the contents of a record zone change.
//
// A record zone subscription executes when a user, or in certain scenarios, CloudKit, modifies a record in that zone, for example, when a field’s value changes in a record. When CloudKit registers the change, it sends push notifications to the user’s devices to inform your app about the change. You can then fetch the changes and cache them on-device. When appropriate, CloudKit excludes the device where the change originates. You configure a subscription’s notifications by setting it’s property. Do this before you save it to the server. A subscription generates either high-priority or medium-priority push notifications. CloudKit delivers medium-priority notifications to your app in the background. High-priority notifications are visual and the system displays them to the user. Visual notifications need the user’s permission. For more information, see . A subscription uses to configure its notifications. For background delivery, set only its property to . If you set any other property, CloudKit treats the notification as high-priority. Don’t rely on push notifications for specific changes to records because the system can coalesce them. CloudKit can omit data to keep the notification’s payload size under the APNs size limit. Consider notifications an indication of remote changes. Use to determine which database contains the changed record zone, and to determine which zone contains changed records. You can then fetch just those changes using . A notification’s property is if CloudKit omits data. You don’t instantiate this class. Instead, implement in your app delegate. Initialize with the dictionary that CloudKit passes to the method. This returns an instance of the appropriate subclass. Use the property to determine the type. Then cast to that type to access type-specific properties and methods.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecordZoneNotification
type CKRecordZoneNotification struct {
	CKNotification
}

// CKRecordZoneNotificationFrom constructs a [CKRecordZoneNotification] from an unsafe.Pointer.
//
// A notification that triggers when the contents of a record zone change.
func CKRecordZoneNotificationFrom(ptr unsafe.Pointer) CKRecordZoneNotification {
	return CKRecordZoneNotification{
		CKNotification: CKNotificationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CKRecordZoneNotificationClass) Alloc() CKRecordZoneNotification {
	rv := objc.Send[CKRecordZoneNotification](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKRecordZoneNotificationClass) New() CKRecordZoneNotification {
	rv := objc.Send[CKRecordZoneNotification](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKRecordZoneNotification) Init() CKRecordZoneNotification {
	rv := objc.Send[CKRecordZoneNotification](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKRecordZoneNotification) Autorelease() CKRecordZoneNotification {
	rv := objc.Send[CKRecordZoneNotification](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKRecordZoneNotification creates a new CKRecordZoneNotification instance.
func NewCKRecordZoneNotification() CKRecordZoneNotification {
	return getCKRecordZoneNotificationClass().New()
}


// The ID of the record zone that has changes.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKRecordZoneNotification/recordZoneID
func (c_ CKRecordZoneNotification) RecordZoneID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("recordZoneID"))
	return rv
}



