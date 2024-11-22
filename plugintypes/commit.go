package plugintypes

import (
	"time"

	cciptypes "github.com/goplugin/plugin-common/pkg/types/ccipocr3"
)

// NOTE: The following type should be moved to internal plugin types after it's not required anymore in pluginv3.0 repo.
// Right now it's only used in a pluginv3.0 repo test: TestCCIPReader_CommitReportsGTETimestamp

type CommitPluginReportWithMeta struct {
	Report    cciptypes.CommitPluginReport `json:"report"`
	Timestamp time.Time                    `json:"timestamp"`
	BlockNum  uint64                       `json:"blockNum"`
}
