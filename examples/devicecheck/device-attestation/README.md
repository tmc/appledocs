# DeviceCheck Framework Overview

This example provides a comprehensive educational overview of Apple's DeviceCheck framework for device attestation and fraud prevention.

## Overview

DeviceCheck is Apple's framework for verifying device authenticity and preventing fraudulent activity. It enables:
- Device token generation and validation
- App integrity attestation (App Attest)
- Per-device data storage (2 bits)
- Fraud detection and prevention
- Bot and abuse detection
- Server-side device verification

## Building

```bash
go build .
```

## Running

```bash
# Show comprehensive framework overview
./device-attestation
```

## Features Demonstrated

1. **Framework Overview** - Core capabilities for device attestation and fraud prevention
2. **Core Classes** - DCDevice and DCAppAttestService
3. **Use Cases** - Fraud prevention, bot detection, rate limiting, promo abuse prevention
4. **DeviceCheck Workflow** - 8-step process from token generation to server validation
5. **App Attest Workflow** - 7-step cryptographic attestation process
6. **Two Bits Storage** - Per-device flag storage (00, 01, 10, 11)
7. **DeviceCheck vs App Attest** - Comparison of security approaches
8. **Server Integration** - API endpoints and authentication requirements
9. **Security Best Practices** - Validation, rate limiting, monitoring
10. **Framework Limitations** - Internet requirements, rate limits, OS versions
11. **Privacy Considerations** - GDPR/CCPA compliance, user privacy
12. **Implementation Steps** - From Xcode setup to production monitoring
13. **Error Handling** - Common errors and solutions
14. **Testing Strategy** - Development and production testing approaches
15. **Complementary Technologies** - reCAPTCHA, CAPTCHA, behavioral analysis

## DeviceCheck vs App Attest

### DCDevice (DeviceCheck)
- Simple device token generation
- Ephemeral tokens (change over time)
- Basic fraud detection
- 2 bits of per-device data
- Server-side validation with Apple
- Good for: Simple fraud prevention, rate limiting

### DCAppAttestService (App Attest)
- Cryptographic attestation
- Persistent key pairs
- Stronger security guarantees
- Verifies app integrity
- Replay protection with counters
- Good for: High-value transactions, security-critical apps

## Two Bits Per-Device Storage

Each device gets exactly 2 bits (4 possible values):

```
00 - New device, no flags
01 - Device has claimed signup bonus
10 - Device flagged for suspicious activity
11 - Device is permanently blocked
```

**Important**: This is NOT a database. Use only for simple flags. Store complex data server-side.

## Server-Side Integration

### Requirements
- Apple Developer account
- Team ID from developer portal
- Private key for server authentication
- HTTPS endpoints for Apple API communication

### DeviceCheck API Endpoints
```
POST https://api.devicecheck.apple.com/v1/validate_device_token
POST https://api.devicecheck.apple.com/v1/update_two_bits
```

### App Attest API Endpoints
```
POST https://api.devicecheck.apple.com/v1/attest
POST https://api.devicecheck.apple.com/v1/validate_assertion
```

## DeviceCheck Workflow

```
┌─────────────────────────────────────────────────────────────┐
│  1. App requests device token from DCDevice                 │
│  2. iOS generates ephemeral device token                    │
│  3. App sends token to your server                          │
│  4. Server queries Apple's DeviceCheck API                  │
│  5. Apple verifies device authenticity                      │
│  6. Server receives device validity + 2 bits of data        │
│  7. Server makes decision (allow/deny)                      │
│  8. Optionally update 2 bits for this device                │
└─────────────────────────────────────────────────────────────┘
```

## App Attest Workflow

```
┌─────────────────────────────────────────────────────────────┐
│  1. App generates attestation key pair                      │
│  2. App creates key identifier                              │
│  3. Server sends challenge                                  │
│  4. App attests key with challenge                          │
│  5. Server validates attestation with Apple                 │
│  6. Future requests include assertions signed with key      │
│  7. Server validates assertions                             │
└─────────────────────────────────────────────────────────────┘
```

## Use Cases

### Fraud Prevention
- Detect and block fraudulent devices
- Prevent account creation abuse
- Identify suspicious patterns

### Bot Detection
- Identify automated/scripted access
- Block non-human interactions
- Protect API endpoints

### Rate Limiting
- Track requests per device
- Enforce fair usage policies
- Prevent API abuse

### Promo Abuse Prevention
- Prevent multiple signups from same device
- Block referral code abuse
- Track bonus claims

### In-App Purchase Verification
- Verify legitimate purchase requests
- Prevent IAP fraud
- Validate transaction integrity

### API Security
- Authenticate device making API calls
- Verify app authenticity
- Prevent API key theft

## Security Best Practices

1. **Never trust client-side data alone** - Always validate with Apple's servers
2. **Use HTTPS** - All communication must be encrypted
3. **Implement rate limiting** - Prevent token request abuse
4. **Monitor patterns** - Look for suspicious behavior
5. **Combine with other methods** - Use alongside behavioral analysis, CAPTCHA, etc.
6. **Rotate attestation keys** - Periodically refresh App Attest keys
7. **Log verification failures** - Track and analyze failed attempts
8. **Handle errors gracefully** - Don't expose security details to clients

## Privacy & Compliance

DeviceCheck is privacy-preserving by design:

- ✅ Device tokens don't identify users
- ✅ Cannot track users across apps
- ✅ No personal information exposed
- ✅ GDPR/CCPA compliant
- ✅ Apple doesn't share user data
- ✅ Tokens reset on factory reset
- ✅ Cannot correlate with Apple ID

## Requirements

- **iOS**: 11.0+ for DeviceCheck, 14.0+ for App Attest
- **macOS**: 10.15+ for DeviceCheck, 11.0+ for App Attest
- **Internet**: Required for server validation
- **Apple Developer Account**: Required for server-side integration
- **Server Infrastructure**: Backend to communicate with Apple APIs

## Limitations

This example is educational only. DeviceCheck requires:

### Client-Side Limitations
- Device tokens are ephemeral (not permanent identifiers)
- Only 2 bits of data per device
- Cannot identify individual users
- Requires internet connection

### Server-Side Requirements
- Must implement Apple API authentication
- Rate limits on API calls
- Requires proper credential management
- Need monitoring infrastructure

### Platform Constraints
- iOS 11+ / macOS 10.15+ required
- App Attest requires iOS 14+ / macOS 11+
- Simulator support limited
- Real devices required for testing

## Implementation Steps

1. **Add DeviceCheck framework** to Xcode project
2. **Import framework** in Swift/Objective-C
3. **Request device token** in app
   ```swift
   DCDevice.current.generateToken { data, error in
       // Send data to your server
   }
   ```
4. **Set up Apple Developer credentials** (Team ID, private key)
5. **Implement server-side validation**
   - Parse JWT token
   - Query Apple's API
   - Validate response
   - Update 2 bits if needed
6. **Test in development** with real devices
7. **Monitor production usage** for fraud patterns
8. **Handle errors gracefully** with fallback logic

## Error Handling

Common errors and solutions:

| Error | Solution |
|-------|----------|
| Token generation fails | Device may not support DeviceCheck (iOS < 11) |
| Server validation fails | Check credentials, network, and API endpoint |
| Rate limit exceeded | Implement exponential backoff |
| Invalid token | Token expired, request new one |
| Attestation not supported | Requires iOS 14+ / macOS 11+ |

## Testing Strategy

### Development
- Use Xcode development environment
- Test on real devices (simulator has limited support)
- Verify server integration with test credentials
- Monitor API response times
- Test error handling paths

### Production
- Monitor API success rates
- Track fraud prevention effectiveness
- A/B test different strategies
- Log and analyze failures
- Set up alerts for anomalies

## Example Server Implementation (Pseudocode)

```python
# Validate device token
def validate_device(device_token, user_id):
    # 1. Generate JWT with your private key
    jwt_token = generate_jwt(team_id, key_id, private_key)

    # 2. Call Apple's validation API
    response = requests.post(
        'https://api.devicecheck.apple.com/v1/validate_device_token',
        headers={'Authorization': f'Bearer {jwt_token}'},
        json={'device_token': device_token.hex()}
    )

    # 3. Check response
    if response.status_code == 200:
        data = response.json()
        # data contains: bit0, bit1, last_update_time

        # 4. Make decision
        if data['bit1'] == 1:  # Device is blocked
            return False

        # 5. Optionally update bits
        update_bits(device_token, bit0=0, bit1=0)
        return True
    else:
        # Handle error
        return None  # Treat as unknown device
```

## Complementary Technologies

Combine DeviceCheck with:

- **reCAPTCHA** - Bot detection for web interfaces
- **CAPTCHA** - Human verification challenges
- **SMS Verification** - Phone number validation
- **Email Verification** - Email address validation
- **Behavioral Analysis** - User interaction patterns
- **IP Reputation** - Network-based fraud detection
- **Device Fingerprinting** - Browser/device characteristics
- **Machine Learning** - Pattern-based anomaly detection

## References

- [DeviceCheck Framework](https://developer.apple.com/documentation/devicecheck)
- [DCDevice](https://developer.apple.com/documentation/devicecheck/dcdevice)
- [DCAppAttestService](https://developer.apple.com/documentation/devicecheck/dcappattestservice)
- [Establishing Your App's Integrity](https://developer.apple.com/documentation/devicecheck/establishing-your-app-s-integrity)
- [DeviceCheck API Documentation](https://developer.apple.com/documentation/devicecheck/accessing_and_modifying_per-device_data)
- [App Attest Guide](https://developer.apple.com/documentation/devicecheck/validating_apps_that_connect_to_your_server)

## See Also

- `../../generated/devicecheck/` - Generated DeviceCheck bindings
- Security framework for related security APIs
- LocalAuthentication framework for biometric verification

## Production Considerations

When implementing DeviceCheck in production:

1. **Monitoring** - Track success rates, latency, error patterns
2. **Fallback Strategy** - Handle cases where DeviceCheck unavailable
3. **User Experience** - Don't block legitimate users on false positives
4. **Server Performance** - Cache decisions, implement rate limiting
5. **Cost Management** - Monitor API usage and costs
6. **Fraud Analysis** - Regularly review fraud patterns and adjust strategies
7. **Compliance** - Ensure GDPR/CCPA compliance in data handling
8. **Documentation** - Document your bit encoding scheme for team

## Notes

This example demonstrates framework capabilities only. Real implementation requires:
- iOS/macOS app with DeviceCheck framework integration
- Server-side implementation with Apple API authentication
- Proper credential management (Team ID, private key)
- Production monitoring and alerting
- Fraud pattern analysis

The generated bindings in `../../generated/devicecheck/` provide the Go interface to DeviceCheck classes, but actual usage requires running within an iOS/macOS app context with proper entitlements.
