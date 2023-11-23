package database

type FactoryConfig struct {
	ConnStr  string
	ConnType string
}

type DbFactory interface {
	Ping() bool
}

func (fc *FactoryConfig) NewDbFactory() DbFactory {
	switch fc.ConnType {
	case "pg":
		return NewPostgres(fc)
	case "sqlite":
		return NewSqlite(fc)
	case "mysql":
		return NewMySql(fc)
	default:
		return NewPostgres(fc)
	}
}
