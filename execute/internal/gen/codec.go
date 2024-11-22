// Package gen wraps an external type to generate a mock object.
package gen

import cciptypes "github.com/goplugin/plugin-common/pkg/types/ccipocr3"

// ExecutePluginCodec is defined in plugin-common.
type ExecutePluginCodec interface {
	cciptypes.ExecutePluginCodec
}
