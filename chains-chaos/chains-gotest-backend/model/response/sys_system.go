package response

import "chains-gotest-backend/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
