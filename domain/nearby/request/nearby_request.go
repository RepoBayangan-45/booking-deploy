package request

type NearbyCreateRequest struct {
	NameFacilities string `json:"namefacilities"`
	Jenis          string `json:"jenis"`
	Jarak          string `json:"jarak"`
	Latitude       string `json:"latitude"`
	Longtitude     string `json:"longtitude"`
}
