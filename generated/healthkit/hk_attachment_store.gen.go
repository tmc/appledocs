// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/uniformtypeidentifiers"
)

/* debug [class.gen.go]: Generating class HKAttachmentStore */


/* debug [class_header]: Header for HKAttachmentStore */
// The class instance for the [HKAttachmentStore] class.
var (
	HKAttachmentStoreClass     _HKAttachmentStoreClass
	HKAttachmentStoreClassOnce sync.Once
)

func getHKAttachmentStoreClass() _HKAttachmentStoreClass {
	HKAttachmentStoreClassOnce.Do(func() {
		HKAttachmentStoreClass = _HKAttachmentStoreClass{objc.GetClass("HKAttachmentStore")}
	})
	return HKAttachmentStoreClass
}

type _HKAttachmentStoreClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKAttachmentStore */
// An interface definition for the [HKAttachmentStore] class.
type IHKAttachmentStore interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKAttachmentStore */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKAttachmentStore */
	// methods:
	AddAttachmentToObjectNameContentTypeURLMetadataCompletion(object IHKObject, name objc.IObject /* cross-framework: NSString */, contentType uniformtypeidentifiers.UTType, URL objc.IObject /* cross-framework: NSURL */, metadata foundation.IDictionary, completion unsafe.Pointer)
	GetAttachmentsForObjectCompletion(object IHKObject, completion unsafe.Pointer)
	GetDataForAttachmentCompletion(attachment IHKAttachment, completion unsafe.Pointer) foundation.Progress
	RemoveAttachmentFromObjectCompletion(attachment IHKAttachment, object IHKObject, completion unsafe.Pointer)
	StreamDataForAttachmentDataHandler(attachment IHKAttachment, dataHandler unsafe.Pointer) foundation.Progress
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKAttachmentStore */
// Alloc allocates a new instance without initialization.
func (hc _HKAttachmentStoreClass) Alloc() HKAttachmentStore {
	rv := objc.Send[HKAttachmentStore](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKAttachmentStoreClass) New() HKAttachmentStore {
	rv := objc.Send[HKAttachmentStore](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKAttachmentStore) Init() HKAttachmentStore {
	rv := objc.Send[HKAttachmentStore](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKAttachmentStore) Autorelease() HKAttachmentStore {
	rv := objc.Send[HKAttachmentStore](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKAttachmentStore creates a new HKAttachmentStore instance.
func NewHKAttachmentStore() HKAttachmentStore {
	return getHKAttachmentStoreClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKAttachmentStore */
// The access point for attachments associated with samples in the HealthKit store.
//
// Use an object to manage attachments for samples that your app has saved to the HealthKit store.


// The access point for attachments associated with samples in the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAttachmentStore
type HKAttachmentStore struct {
	objectivec.Object
}

// HKAttachmentStoreFrom constructs a [HKAttachmentStore] from an unsafe.Pointer.
//
// The access point for attachments associated with samples in the HealthKit store.
func HKAttachmentStoreFrom(ptr unsafe.Pointer) HKAttachmentStore {
	return HKAttachmentStore{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKAttachmentStore */

// Creates an attachment store for the provided HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAttachmentStore/init(healthStore:)
func NewHKAttachmentStoreWithHealthStore(healthStore IHKHealthStore) HKAttachmentStore {
	instance := getHKAttachmentStoreClass().Alloc()
	rv := objc.Send[HKAttachmentStore](instance.ID, objc.Sel("initWithHealthStore:"), healthStore)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKAttachmentStoreWithHealthStore */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKAttachmentStore */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKAttachmentStore */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKAttachmentStore */

// Adds an attachment to the specified object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAttachmentStore/addAttachmentToObject:name:contentType:URL:metadata:completion:
func (h_ HKAttachmentStore) AddAttachmentToObjectNameContentTypeURLMetadataCompletion(object IHKObject, name objc.IObject /* cross-framework: NSString */, contentType uniformtypeidentifiers.UTType, URL objc.IObject /* cross-framework: NSURL */, metadata foundation.IDictionary, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("addAttachmentToObject:name:contentType:URL:metadata:completion:"), object, name, contentType, URL, metadata, completion)
}/* debug [instance_methods/method]: AddAttachmentToObjectNameContentTypeURLMetadataCompletion */


// Returns all the attachments for the specified object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAttachmentStore/getAttachments(for:completion:)
func (h_ HKAttachmentStore) GetAttachmentsForObjectCompletion(object IHKObject, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("getAttachmentsForObject:completion:"), object, completion)
}/* debug [instance_methods/method]: GetAttachmentsForObjectCompletion */


// Returns an attachment’s data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAttachmentStore/getData(for:completion:)
func (h_ HKAttachmentStore) GetDataForAttachmentCompletion(attachment IHKAttachment, completion unsafe.Pointer) foundation.Progress {
	rv := objc.Send[foundation.Progress](h_.ID, objc.Sel("getDataForAttachment:completion:"), attachment, completion)
	return rv
}/* debug [instance_methods/method]: GetDataForAttachmentCompletion */


// Removes the specified attachment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAttachmentStore/removeAttachment(_:from:completion:)
func (h_ HKAttachmentStore) RemoveAttachmentFromObjectCompletion(attachment IHKAttachment, object IHKObject, completion unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("removeAttachment:fromObject:completion:"), attachment, object, completion)
}/* debug [instance_methods/method]: RemoveAttachmentFromObjectCompletion */


// Asynchronously returns the attachment’s data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAttachmentStore/streamData(for:dataHandler:)
func (h_ HKAttachmentStore) StreamDataForAttachmentDataHandler(attachment IHKAttachment, dataHandler unsafe.Pointer) foundation.Progress {
	rv := objc.Send[foundation.Progress](h_.ID, objc.Sel("streamDataForAttachment:dataHandler:"), attachment, dataHandler)
	return rv
}/* debug [instance_methods/method]: StreamDataForAttachmentDataHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKAttachmentStore */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKAttachmentStore */


