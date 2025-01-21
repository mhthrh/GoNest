package address

const (
	path = "/file/countries"
	name = "Countries.csv"
)

type Country struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}
