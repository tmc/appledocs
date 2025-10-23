// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HKAttachment] class.
var (
	HKAttachmentClass     _HKAttachmentClass
	HKAttachmentClassOnce sync.Once
)

func getHKAttachmentClass() _HKAttachmentClass {
	HKAttachmentClassOnce.Do(func() {
		HKAttachmentClass = _HKAttachmentClass{objc.GetClass("HKAttachment")}
	})
	return HKAttachmentClass
}

type _HKAttachmentClass struct {
	class objc.Class
}

// An interface definition for the [HKAttachment] class.
type IHKAttachment interface {
	objectivec.IObject
	// properties:
	ContentType() objectivec.IObject
	SetContentType(value objectivec.IObject)
	CreationDate() foundation.objc.IObject /* cross-framework: Date */
	SetCreationDate(value foundation.objc.IObject /* cross-framework: Date */)
	Identifier() foundation.objc.IObject /* cross-framework: UUID */
	SetIdentifier(value foundation.objc.IObject /* cross-framework: UUID */)
	Metadata() string /* primitive/slice/pointer. */
	SetMetadata(value string /* primitive/slice/pointer. */)
	Name() string /* primitive/slice/pointer. */
	SetName(value string /* primitive/slice/pointer. */)
	Size() int /* primitive/slice/pointer. */
	SetSize(value int /* primitive/slice/pointer. */)
	// methods:
}

// A file that is attached to a sample in the HealthKit store.
//
// To access the attachment’s data, get a data reader from the attachment store for each attachment. You can then asynchronously access the whole data object. Alternatively, you can access the file’s contents as an asynchronous sequence of bytes.


// A file that is attached to a sample in the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAttachment
type HKAttachment struct {
	objectivec.Object
}

// HKAttachmentFrom constructs a [HKAttachment] from an unsafe.Pointer.
//
// A file that is attached to a sample in the HealthKit store.
func HKAttachmentFrom(ptr unsafe.Pointer) HKAttachment {
	return HKAttachment{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKAttachmentClass) Alloc() HKAttachment {
	rv := objc.Send[HKAttachment](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKAttachmentClass) New() HKAttachment {
	rv := objc.Send[HKAttachment](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKAttachment) Init() HKAttachment {
	rv := objc.Send[HKAttachment](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKAttachment) Autorelease() HKAttachment {
	rv := objc.Send[HKAttachment](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKAttachment creates a new HKAttachment instance.
func NewHKAttachment() HKAttachment {
	return getHKAttachmentClass().New()
}



// The type of data stored in the attached file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkattachment/contenttype
func (h_ HKAttachment) ContentType() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](h_.ID, objc.Sel("contentType"))
	return rv
}


// The type of data stored in the attached file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkattachment/contenttype
func (h_ HKAttachment) SetContentType(value objectivec.IObject) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setContentType:"), value)
}


// The attachment’s creation date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkattachment/creationdate
func (h_ HKAttachment) CreationDate() foundation.objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](h_.ID, objc.Sel("creationDate"))
	return rv
}


// The attachment’s creation date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkattachment/creationdate
func (h_ HKAttachment) SetCreationDate(value foundation.objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setCreationDate:"), value)
}


// The universally unique identifier for the attached file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkattachment/identifier
func (h_ HKAttachment) Identifier() foundation.objc.IObject /* cross-framework: UUID */ {
	rv := objc.Send[foundation.UUID](h_.ID, objc.Sel("identifier"))
	return rv
}


// The universally unique identifier for the attached file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkattachment/identifier
func (h_ HKAttachment) SetIdentifier(value foundation.objc.IObject /* cross-framework: UUID */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setIdentifier:"), value)
}


// Additional data associated with the attachment in the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkattachment/metadata
func (h_ HKAttachment) Metadata() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](h_.ID, objc.Sel("metadata"))
	return rv
}


// Additional data associated with the attachment in the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkattachment/metadata
func (h_ HKAttachment) SetMetadata(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setMetadata:"), objc.String(value))
}


// The name of the attached file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkattachment/name
func (h_ HKAttachment) Name() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](h_.ID, objc.Sel("name"))
	return rv
}


// The name of the attached file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkattachment/name
func (h_ HKAttachment) SetName(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setName:"), objc.String(value))
}


// The attachment’s size (in bytes).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkattachment/size
func (h_ HKAttachment) Size() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](h_.ID, objc.Sel("size"))
	return rv
}


// The attachment’s size (in bytes).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkattachment/size
func (h_ HKAttachment) SetSize(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setSize:"), value)
}



