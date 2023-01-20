package blob

type Blob struct {
	Url       string `json:"url"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
	Title     string `json:"title"`
}
type BlobPointers struct {
	Blobs []Blob `json:"blobs"`
	Count int    `json:"count"`
}

type UploadBlobResponseDTO struct {
	Blob
}

type BlobPointersResponseDTO struct {
	BlobPointers
}
