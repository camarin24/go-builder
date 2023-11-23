package types

type DbPing struct {
	ConnStr  string `json:"connStr"`
	ConnType string `json:"connType"`
}
