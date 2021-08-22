package file

import "time"

type File struct {
	tableName struct{} `pg:"lk_dev.file_storage_item"`

	ID               uint
	OriginalFilename string
	Extension        string
	MimeType         string
	Size             uint
	Checksum         string
	CreatedAt        time.Time
}
