package blob

type UploadBlobResponseDTO struct {
	Location string `json:"location"`
}

type ListBlobsResponseDTO struct {
	Blobs []string `json:"blobs"`
	Count int      `json:"count"`
}
