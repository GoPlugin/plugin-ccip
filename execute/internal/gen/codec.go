// Package gen wraps an external type to generate a mock object.
package gen

import cciptypes "github.com/goplugin/plugin-ccip/pkg/types/ccipocr3"

// TODO: get rid of this.
// ExecutePluginCodec is defined in plugin-common.
type ExecutePluginCodec interface {
	cciptypes.ExecutePluginCodec
}
