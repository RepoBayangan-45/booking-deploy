package response

type NearbyResponse struct {
	ID             int    `json:"id"`
	NameFacilities string `json:"namefacilities"`
	Jenis          string `json:"jenis"`
	Jarak          string `json:"jarak"`
	Latitude       string `json:"latitude"`
	Longtitude     string `json:"longtitude"`
}
