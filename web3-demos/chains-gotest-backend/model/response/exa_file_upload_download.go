package response

import "chains-gotest-backend/model"

type ExaFileResponse struct {
	File model.ExaFileUploadAndDownload `json:"file"`
}
