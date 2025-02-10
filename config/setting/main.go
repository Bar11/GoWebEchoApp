package setting

type settings struct {
	routeSettings
	sqlSettings
}

type routeSettings struct {
	LoginUrl string `json:"loginUrl"`
	IndexUrl string `json:"indexUrl"`
}

type sqlSettings struct {
	Password string `json:"password"`
	Username string `json:"username"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
}

var sqlSetting = sqlSettings{
	Password: "123qweasd",
	Username: "root",
	Host:     "localhost",
	Port:     "3306",
	Database: "EchoAppData",
}

var routeSetting = routeSettings{
	LoginUrl: "/login",
	IndexUrl: "/index",
}

var Settings = settings{
	routeSettings: routeSetting,
	sqlSettings:   sqlSetting,
}
