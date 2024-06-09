package ninjarmm

type Cursor struct {
	Name    string    `json:"name"`
	Offset  int32     `json:"offset"`
	Count   int32     `json:"count"`
	Expires Timestamp `json:"expires"`
}

type QueryResult[T any] struct {
	Cursor  *Cursor `json:"cursor"`
	Results []T     `json:"results"`
}

type SoftwareInventoryResult struct {
	InstallDate InstallDate `json:"installDate"`
	Location    string      `json:"location"`
	Name        string      `json:"name"`
	Publisher   string      `json:"publisher"`
	Size        int32       `json:"size"`
	Version     string      `json:"version"`
	ProductCode string      `json:"productCode"`
	DeviceID    int32       `json:"deviceId"`
	Timestamp   Timestamp   `json:"timestamp"`
}

type SoftwareInventory = QueryResult[SoftwareInventoryResult]
