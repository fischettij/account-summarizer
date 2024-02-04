package email

type Config struct {
	Port        string
	ServerURL   string
	From        string
	Username    string
	Password    string
	Identity    string
	TLSHostName string
}
