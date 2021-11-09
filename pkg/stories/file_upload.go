package stories

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/keygaen/pkg/components"
)

type FileUploadStory struct {
	Story

	fileContents []byte
}

func (c *FileUploadStory) Render() app.UI {
	c.EnableShallowReflection()

	return c.WithRoot(
		&components.FileUpload{
			ID:                    "file-upload-story",
			FileSelectionLabel:    "Drag and drop a file or select one",
			ClearLabel:            "Clear",
			TextEntryLabel:        "Or enter text here",
			TextEntryBlockedLabel: "File has been selected.",
			FileContents:          c.fileContents,

			OnChange: func(fileContents []byte) {
				c.fileContents = fileContents
			},
			OnClear: func() {
				c.fileContents = []byte{}
			},
		},
	)
}
