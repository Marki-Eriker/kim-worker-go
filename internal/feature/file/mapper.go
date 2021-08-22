package file

import (
	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
)

func MapOneToGqlModel(f *File) *model.File {
	return &model.File{
		ID:               f.ID,
		OriginalFileName: f.OriginalFilename,
		Extension:        f.Extension,
		MimeType:         f.MimeType,
		Size:             f.Size,
		Checksum:         f.Checksum,
		CreatedAt:        f.CreatedAt,
	}
}
