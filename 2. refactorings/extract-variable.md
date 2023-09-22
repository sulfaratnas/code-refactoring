# Rename Method

If you have an expression that’s hard to understand. Place the result of the expression or its parts in separate variables that are self-explanatory.

## Example

### Before

```
func renderBanner() {
    if strings.Contains(strings.ToUpper(platform), "MAC") &&
        strings.Contains(strings.ToUpper(browser), "IE") &&
        wasInitialized() && resize > 0 {
        // do something
    }
}

func wasInitialized() bool {
    // Replace this with your initialization logic
    return true
}

var platform string // Initialize platform
var browser string  // Initialize browser
var resize int     // Initialize resize
```

### After

```
func renderBanner() {
    isMacOs := strings.Contains(strings.ToUpper(platform), "MAC")
    isIE := strings.Contains(strings.ToUpper(browser), "IE")
    wasResized := resize > 0

    if isMacOs && isIE && wasInitialized() && wasResized {
        // do something
    }
}

func wasInitialized() bool {
    // Replace this with your initialization logic
    return true
}

var platform string // Initialize platform
var browser string  // Initialize browser
var resize int     // Initialize resize

```

