# Granting Screen Recording Permission

The ScreenCaptureKit example requires Screen Recording permission to function. Here's how to grant it:

## Method 1: Via System Settings (Recommended)

1. Open **System Settings** (System Preferences on older macOS)
2. Go to **Privacy & Security** → **Screen Recording**
3. Look for **ScreenCaptureKit-Example.app** in the list
4. Toggle it ON (enable the checkbox)
5. Rerun the example

## Method 2: Reset and Rerun

If the app doesn't appear in the list, reset the permission and try again:

```bash
# Reset the TCC permission
tccutil reset ScreenCapture com.github.tmc.appledocs.screencapturekit-example

# Run the app - it should now prompt
./screencapturekit-generated-bindings
```

## Method 3: Open System Settings Directly

```bash
# Open Screen Recording preferences directly
open "x-apple.systempreferences:com.apple.preference.security?Privacy_ScreenCapture"
```

## Verification

Once permission is granted, the example should output:

```
✓ Found X display(s)
   Display 0: ID=..., WxH
✓ Found X window(s)
```

## Troubleshooting

**Problem**: App shows "user declined TCCs" immediately without prompting  
**Solution**: The permission was previously denied. Use Method 1 above to manually grant it.

**Problem**: App not in the Screen Recording list  
**Solution**: Run the app once first, then check System Settings.

**Problem**: Permission granted but still fails  
**Solution**: Try killing all instances and relaunching:
```bash
pkill -f ScreenCaptureKit-Example
./screencapturekit-generated-bindings
```

## Why This Happens

macOS remembers TCC decisions. Once you deny permission (either explicitly or by timeout), the system caches that decision. The permission must be manually granted in System Settings after a denial.

## App Bundle Location

The macgo wrapper creates an app bundle at:
```
/Users/tmc/go/bin/ScreenCaptureKit-Example.app
```

This is the bundle that needs to be granted Screen Recording permission in System Settings.
