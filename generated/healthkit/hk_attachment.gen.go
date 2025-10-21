// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

// A file that is attached to a sample in the HealthKit store.
//
// To access the attachment’s data, get a data reader from the attachment store for each attachment. You can then asynchronously access the whole data object. Alternatively, you can access the file’s contents as an asynchronous sequence of bytes.
//
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


// Additional data associated with the attachment in the HealthKit store.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkattachment/metadata
func (h_ HKAttachment) Metadata() string {
	rv := objc.Send[string](h_.ID, objc.Sel("metadata"))
	return rv
}


// SetMetadata sets the value of the metadata property.
// Additional data associated with the attachment in the HealthKit store.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkattachment/metadata
func (h_ HKAttachment) SetMetadata(value string) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setMetadata:"), objc.String(value))
}

// The attachment’s creation date.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkattachment/creationdate
func (h_ HKAttachment) CreationDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("creationDate"))
	return rv
}


// SetCreationDate sets the value of the creationDate property.
// The attachment’s creation date.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkattachment/creationdate
func (h_ HKAttachment) SetCreationDate(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setCreationDate:"), value)
}

// The type of data stored in the attached file.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkattachment/contenttype
func (h_ HKAttachment) ContentType() UTType {
	rv := objc.Send[UTType](h_.ID, objc.Sel("contentType"))
	return rv
}


// SetContentType sets the value of the contentType property.
// The type of data stored in the attached file.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkattachment/contenttype
func (h_ HKAttachment) SetContentType(value UTType) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setContentType:"), value)
}

// The attachment’s size (in bytes).
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkattachment/size
func (h_ HKAttachment) Size() int {
	rv := objc.Send[int](h_.ID, objc.Sel("size"))
	return rv
}


// SetSize sets the value of the size property.
// The attachment’s size (in bytes).

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkattachment/size
func (h_ HKAttachment) SetSize(value int) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setSize:"), value)
}

// The universally unique identifier for the attached file.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkattachment/identifier
func (h_ HKAttachment) Identifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("identifier"))
	return rv
}


// SetIdentifier sets the value of the identifier property.
// The universally unique identifier for the attached file.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkattachment/identifier
func (h_ HKAttachment) SetIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setIdentifier:"), value)
}

// The name of the attached file.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkattachment/name
func (h_ HKAttachment) Name() string {
	rv := objc.Send[string](h_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
// The name of the attached file.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkattachment/name
func (h_ HKAttachment) SetName(value string) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setName:"), objc.String(value))
}



