package mobile

import "C"

// Exported function (must start with uppercase to be exported)
func StartCore(config string) string {
    return "Core started with config: " + config
}
