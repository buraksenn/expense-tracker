package receipt

import "time"

type UploadReceiptCommand struct {
	Link string
}

type GetReceiptsCommand struct {
	StartTime time.Time
}
