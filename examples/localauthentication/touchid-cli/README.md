# LocalAuthentication Touch ID/Face ID CLI

This example demonstrates the LocalAuthentication framework for biometric authentication on macOS using the generated Go bindings.

## Overview

LocalAuthentication provides APIs for:
- Biometric authentication (Touch ID, Face ID, Optic ID)
- Device passcode authentication
- Apple Watch authentication
- Secure credential evaluation

## Building

```bash
go build .
```

## Running

```bash
# Show framework overview and capabilities
./touchid-cli

# Show authentication workflow notes
./touchid-cli -auth
```

## Features Demonstrated

1. **Framework Overview** - Core capabilities and authentication methods
2. **Authentication Policies** - Different policy types (biometrics, passcode, Watch)
3. **LAContext Creation** - Creating and managing authentication contexts
4. **Biometric Availability** - Checking if biometrics are available
5. **Context Properties** - Customizing authentication prompts
6. **Error Handling** - Common error codes and their meanings
7. **Authentication Workflow** - Step-by-step authentication process
8. **Use Cases** - Real-world applications
9. **Biometry Types** - Touch ID, Face ID, Optic ID support
10. **Best Practices** - Recommended implementation patterns
11. **Platform Requirements** - Hardware and OS requirements
12. **Security Considerations** - How biometric data is protected
13. **Troubleshooting** - Common issues and solutions

## Key Classes

- `LAContext` - Manages authentication state and user interaction
- `LAPolicy` - Defines authentication policies
- `LAError` - Authentication error codes

## Authentication Policies

| Policy | Description |
|--------|-------------|
| DeviceOwnerAuthenticationWithBiometrics | Biometrics only (Touch ID/Face ID) |
| DeviceOwnerAuthentication | Biometrics or Passcode |
| DeviceOwnerAuthenticationWithWatch | Apple Watch authentication |
| DeviceOwnerAuthenticationWithBiometricsOrWatch | Biometrics or Watch |
| DeviceOwnerAuthenticationWithWristDetection | Watch with wrist detection |

## Requirements

- macOS 10.10+ for basic LocalAuthentication
- Touch ID: MacBook Pro (2016+), MacBook Air (2018+)
- Face ID: Not available on macOS (iOS/iPadOS only)
- Optic ID: Apple Vision Pro
- Device passcode must be configured
- Biometric data must be enrolled

## Security

- Biometric data never leaves the Secure Enclave
- Apps only receive success/failure (not actual biometric data)
- Multiple failed attempts trigger automatic lockout
- System controls all UI and user interaction
- Cannot be bypassed or manipulated

## Limitations

This example demonstrates the framework APIs but does not perform actual authentication because:

1. **Block Callbacks** - Full authentication requires Objective-C block handling for async callbacks
2. **Type Marshaling** - Complex parameter types need proper Objective-C type wrapping
3. **Hardware Requirements** - Actual authentication requires Touch ID/Face ID hardware

For production use, consider:
- Using cgo with Objective-C wrapper for block callbacks
- Testing on actual hardware (not VMs or CI)
- Implementing comprehensive error handling

## References

- [LocalAuthentication Framework](https://developer.apple.com/documentation/localauthentication)
- [LAContext](https://developer.apple.com/documentation/localauthentication/lacontext)
- [LAPolicy](https://developer.apple.com/documentation/localauthentication/lapolicy)
- [Logging a User into Your App with Face ID or Touch ID](https://developer.apple.com/documentation/localauthentication/logging_a_user_into_your_app_with_face_id_or_touch_id)
