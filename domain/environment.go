package domain

type Environment struct {
	HOST string
	PORT string
	DBUSERNAME string
	DBPASSWORD string
	DBHOST string
	DBPORT string
	DBNAME string
	LOGPATH string
}

func NewEnvironment(host, port, dbUsername, dbPassword, dbHost, dbPort, dbName, logPath string) Environment {
	return Environment{
		HOST:       host,
		PORT:       port,
		DBUSERNAME: dbUsername,
		DBPASSWORD: dbPassword,
		DBHOST:     dbHost,
		DBPORT:     dbPort,
		DBNAME:     dbName,
		LOGPATH:    logPath,
	}
}