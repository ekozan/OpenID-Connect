package OpenIDConnect

//Storage .
type Storage interface {
	getClients() []Client
	getClient(id int) Client
	createClient(c Client)
	deleteClient(id int)
	getUsers()
	getUser()
	createUser()
	deleteUser()
	getTokens()
	getToken(id string)
	createToken()
	deleteToken()
	purgeToken()
	createCode()
	getCode(id string)
	purgeCodes()
	getConfig() Config
}

//mysqlStorage .
type mysqlStorage struct {
	host     string
	port     int
	base     string
	user     string
	password string
}
