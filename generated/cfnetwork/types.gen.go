// Code generated from Apple documentation for CFNetwork. DO NOT EDIT.

package cfnetwork
import (
	"unsafe"
)


// C struct types
// CFHostClientContext - A structure containing user-defined data and callbacks for CFHost objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHostClientContext
type CFHostClientContext struct {
	CopyDescription AllocatorCopyDescriptionCallBack // The callback used to create a descriptive string representation of the info pointer (or the data pointed to by the info pointer) for debugging purposes. This callback is called by the   function.
	Info unsafe.Pointer // An arbitrary pointer to allocated memory containing user-defined data that can be associated with the host and that is passed to the callbacks.
	Release AllocatorReleaseCallBack // The callback used to remove a retain previously added for the host on the info pointer.
	Retain AllocatorRetainCallBack // The callback used to add a retain for the host on the info pointer for the life of the host, and may be used for temporary references the host needs to take. This callback returns the actual info pointer to store in the host, almost always just the pointer passed as the parameter.
	Version Index // The version number of the structure type passed as a parameter to the host client function. The only valid version number is  .
}

// CFNetServiceClientContext - A structure provided when a CFNetService is associated with a callback function or when a CFNetServiceBrowser is created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceClientContext
type CFNetServiceClientContext struct {
	CopyDescription AllocatorCopyDescriptionCallBack // Callback used to create a descriptive string representation of the data pointed to by  . In implementing this function, return a reference to a CFString object that describes your allocator and some characteristics of your user-defined data, which is used by  . You can set this field to  , in which case Core Foundation will provide a rudimentary description.
	Info unsafe.Pointer // Arbitrary pointer to user-allocated memory containing user-defined data that is associated with the service, browser, or monitor and is passed to their respective callback functions. The data must be valid for as long as the CFNetService, CFNetServiceBrowser, or CFNetServiceMonitor is valid. Set this field to   if your callback function doesn’t want to receive user-defined data.
	Release AllocatorReleaseCallBack // Callback that removes a retain previously added for the service or browser on the   pointer. This field can be  , but setting this field to   may result in memory leaks.
	Retain AllocatorRetainCallBack // The callback used to add a retain for the service or browser using   for the life of the service or browser. This callback may be used for temporary references the service or browser needs to take. This callback returns the actual   pointer so it can be stored in the service or browser. This field can be  .
	Version Index // Version number for this structure. Currently the only valid value is zero.
}





