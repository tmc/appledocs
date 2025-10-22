// Code generated from Apple documentation for ExternalAccessory. DO NOT EDIT.

package externalaccessory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [EASession] class.
var (
	EASessionClass     _EASessionClass
	EASessionClassOnce sync.Once
)

func getEASessionClass() _EASessionClass {
	EASessionClassOnce.Do(func() {
		EASessionClass = _EASessionClass{objc.GetClass("EASession")}
	})
	return EASessionClass
}

type _EASessionClass struct {
	class objc.Class
}

// An interface definition for the [EASession] class.
type IEASession interface {
	objectivec.IObject
	Accessory() EAAccessory
	InputStream() foundation.InputStream
	OutputStream() foundation.OutputStream
	ProtocolString() string
}

// The object you use to manage communications between your app and a connected hardware accessory.
//
// An object creates a communications channel between your app and a connected hardware accessory. The manufacturer of the device must share the accessory’s supported protocols with you. When you create your session, specify one of these protocols to initiate communication with the accessory. After initializing an object, use the provided output and input streams to transfer data to and from the accessory using that protocol. After creating a session object, immediately retrieve and configure the stream objects provided by the session. Streams send events to their associated delegate to notify it of changes in the stream status. For example, streams notify the delegate when data is waiting to be read or when more space is available for writing data. For more information about how to use stream objects, see . When sending and receiving data using the provided streams, it is your responsibility to ensure the data is formatted according to the specified protocol. The class has no knowledge of specific accessory protocols and doesn’t attempt to format the data in any way before or after transferring it.
//
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EASession
type EASession struct {
	objectivec.Object
}

// EASessionFrom constructs a [EASession] from an unsafe.Pointer.
//
// The object you use to manage communications between your app and a connected hardware accessory.
func EASessionFrom(ptr unsafe.Pointer) EASession {
	return EASession{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ec _EASessionClass) Alloc() EASession {
	rv := objc.Send[EASession](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _EASessionClass) New() EASession {
	rv := objc.Send[EASession](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EASession) Init() EASession {
	rv := objc.Send[EASession](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EASession) Autorelease() EASession {
	rv := objc.Send[EASession](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEASession creates a new EASession instance.
func NewEASession() EASession {
	return getEASessionClass().New()
}




// Initializes the session for the specified accessory and protocol.
//
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EASession/init(accessory:forProtocol:)
func NewEASessionWithAccessoryForProtocol(accessory IEAAccessory, protocolString string) EASession {
	instance := getEASessionClass().Alloc()
	rv := objc.Send[EASession](instance.ID, objc.Sel("initWithAccessory:forProtocol:"), accessory, objc.String(protocolString))
	rv.Autorelease()
	return rv
}


// The accessory attached to the session.
//
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EASession/accessory
func (e_ EASession) Accessory() EAAccessory {
	rv := objc.Send[EAAccessory](e_.ID, objc.Sel("accessory"))
	return rv
}

// The stream to use for receiving data from the accessory.
//
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EASession/inputStream
func (e_ EASession) InputStream() foundation.InputStream {
	rv := objc.Send[foundation.InputStream](e_.ID, objc.Sel("inputStream"))
	return rv
}

// The stream to use for sending data to the accessory.
//
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EASession/outputStream
func (e_ EASession) OutputStream() foundation.OutputStream {
	rv := objc.Send[foundation.OutputStream](e_.ID, objc.Sel("outputStream"))
	return rv
}

// The protocol being used for communication with the accessory.
//
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EASession/protocolString
func (e_ EASession) ProtocolString() string {
	rv := objc.Send[string](e_.ID, objc.Sel("protocolString"))
	return rv
}


