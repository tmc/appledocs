// Code generated from Apple documentation for PushKit. DO NOT EDIT.

package pushkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// PPushRegistryDelegate is the PKPushRegistryDelegate protocol interface.
//
// The methods that you use to handle incoming PushKit notifications and registration events.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.15+
//   - visionOS 1.0+
//   - watchOS 6.0+
//
// See: doc://com.apple.pushkit/documentation/PushKit/PKPushRegistryDelegate
type PPushRegistryDelegate interface {
	// Required methods
	PushRegistryDidUpdatePushCredentialsForType(registry IPKPushRegistry, pushCredentials IPKPushCredentials, type_ PushType /* typedef */)/* debug [protocol_interface/required_method]: PushRegistryDidUpdatePushCredentialsForType */
	// Optional methods
	PushRegistryDidInvalidatePushTokenForType(registry IPKPushRegistry, type_ PushType /* typedef */)
	HasPushRegistryDidInvalidatePushTokenForType() bool
	PushRegistryDidReceiveIncomingPushWithPayloadForType(registry IPKPushRegistry, payload IPKPushPayload, type_ PushType /* typedef */)
	HasPushRegistryDidReceiveIncomingPushWithPayloadForType() bool
	PushRegistryDidReceiveIncomingPushWithPayloadForTypeWithCompletionHandler(registry IPKPushRegistry, payload IPKPushPayload, type_ PushType /* typedef */, completion unsafe.Pointer)
	HasPushRegistryDidReceiveIncomingPushWithPayloadForTypeWithCompletionHandler() bool
}

// PushRegistryDelegate is a delegate implementation builder for the PPushRegistryDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type PushRegistryDelegate struct {
	_PushRegistryDidInvalidatePushTokenForType func(registry IPKPushRegistry, type_ PushType /* typedef */)
	_PushRegistryDidReceiveIncomingPushWithPayloadForType func(registry IPKPushRegistry, payload IPKPushPayload, type_ PushType /* typedef */)
	_PushRegistryDidReceiveIncomingPushWithPayloadForTypeWithCompletionHandler func(registry IPKPushRegistry, payload IPKPushPayload, type_ PushType /* typedef */, completion unsafe.Pointer)
	_PushRegistryDidUpdatePushCredentialsForType func(registry IPKPushRegistry, pushCredentials IPKPushCredentials, type_ PushType /* typedef */)
}

// SetPushRegistryDidInvalidatePushTokenForType sets the handler for the PushRegistryDidInvalidatePushTokenForType delegate method.
//
// Tells the delegate that the system invalidated the push token for the specified type.
func (d *PushRegistryDelegate) SetPushRegistryDidInvalidatePushTokenForType(f func(registry IPKPushRegistry, type_ PushType /* typedef */)) {
	d._PushRegistryDidInvalidatePushTokenForType = f
}

// SetPushRegistryDidReceiveIncomingPushWithPayloadForType sets the handler for the PushRegistryDidReceiveIncomingPushWithPayloadForType delegate method.
//
// Notifies the delegate that a remote push has been received.
func (d *PushRegistryDelegate) SetPushRegistryDidReceiveIncomingPushWithPayloadForType(f func(registry IPKPushRegistry, payload IPKPushPayload, type_ PushType /* typedef */)) {
	d._PushRegistryDidReceiveIncomingPushWithPayloadForType = f
}

// SetPushRegistryDidReceiveIncomingPushWithPayloadForTypeWithCompletionHandler sets the handler for the PushRegistryDidReceiveIncomingPushWithPayloadForTypeWithCompletionHandler delegate method.
//
// Tells the delegate that a remote push notification arrived.
func (d *PushRegistryDelegate) SetPushRegistryDidReceiveIncomingPushWithPayloadForTypeWithCompletionHandler(f func(registry IPKPushRegistry, payload IPKPushPayload, type_ PushType /* typedef */, completion unsafe.Pointer)) {
	d._PushRegistryDidReceiveIncomingPushWithPayloadForTypeWithCompletionHandler = f
}

// SetPushRegistryDidUpdatePushCredentialsForType sets the handler for the PushRegistryDidUpdatePushCredentialsForType delegate method.
//
// Tells the delegate that the system updated the credentials for the specified type of push notification.
func (d *PushRegistryDelegate) SetPushRegistryDidUpdatePushCredentialsForType(f func(registry IPKPushRegistry, pushCredentials IPKPushCredentials, type_ PushType /* typedef */)) {
	d._PushRegistryDidUpdatePushCredentialsForType = f
}

// PushRegistryDidInvalidatePushTokenForType implements the PPushRegistryDelegate interface.
func (d *PushRegistryDelegate) PushRegistryDidInvalidatePushTokenForType(registry IPKPushRegistry, type_ PushType /* typedef */) {
	if d._PushRegistryDidInvalidatePushTokenForType != nil {
		d._PushRegistryDidInvalidatePushTokenForType(registry, type_)
	}
}

// HasPushRegistryDidInvalidatePushTokenForType returns true if a handler for PushRegistryDidInvalidatePushTokenForType has been set.
func (d *PushRegistryDelegate) HasPushRegistryDidInvalidatePushTokenForType() bool {
	return d._PushRegistryDidInvalidatePushTokenForType != nil
}

// PushRegistryDidReceiveIncomingPushWithPayloadForType implements the PPushRegistryDelegate interface.
func (d *PushRegistryDelegate) PushRegistryDidReceiveIncomingPushWithPayloadForType(registry IPKPushRegistry, payload IPKPushPayload, type_ PushType /* typedef */) {
	if d._PushRegistryDidReceiveIncomingPushWithPayloadForType != nil {
		d._PushRegistryDidReceiveIncomingPushWithPayloadForType(registry, payload, type_)
	}
}

// HasPushRegistryDidReceiveIncomingPushWithPayloadForType returns true if a handler for PushRegistryDidReceiveIncomingPushWithPayloadForType has been set.
func (d *PushRegistryDelegate) HasPushRegistryDidReceiveIncomingPushWithPayloadForType() bool {
	return d._PushRegistryDidReceiveIncomingPushWithPayloadForType != nil
}

// PushRegistryDidReceiveIncomingPushWithPayloadForTypeWithCompletionHandler implements the PPushRegistryDelegate interface.
func (d *PushRegistryDelegate) PushRegistryDidReceiveIncomingPushWithPayloadForTypeWithCompletionHandler(registry IPKPushRegistry, payload IPKPushPayload, type_ PushType /* typedef */, completion unsafe.Pointer) {
	if d._PushRegistryDidReceiveIncomingPushWithPayloadForTypeWithCompletionHandler != nil {
		d._PushRegistryDidReceiveIncomingPushWithPayloadForTypeWithCompletionHandler(registry, payload, type_, completion)
	}
}

// HasPushRegistryDidReceiveIncomingPushWithPayloadForTypeWithCompletionHandler returns true if a handler for PushRegistryDidReceiveIncomingPushWithPayloadForTypeWithCompletionHandler has been set.
func (d *PushRegistryDelegate) HasPushRegistryDidReceiveIncomingPushWithPayloadForTypeWithCompletionHandler() bool {
	return d._PushRegistryDidReceiveIncomingPushWithPayloadForTypeWithCompletionHandler != nil
}

// PushRegistryDidUpdatePushCredentialsForType implements the PPushRegistryDelegate interface.
func (d *PushRegistryDelegate) PushRegistryDidUpdatePushCredentialsForType(registry IPKPushRegistry, pushCredentials IPKPushCredentials, type_ PushType /* typedef */) {
	if d._PushRegistryDidUpdatePushCredentialsForType != nil {
		d._PushRegistryDidUpdatePushCredentialsForType(registry, pushCredentials, type_)
	}
}

// HasPushRegistryDidUpdatePushCredentialsForType returns true if a handler for PushRegistryDidUpdatePushCredentialsForType has been set.
func (d *PushRegistryDelegate) HasPushRegistryDidUpdatePushCredentialsForType() bool {
	return d._PushRegistryDidUpdatePushCredentialsForType != nil
}
