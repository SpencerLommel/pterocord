package pterodactyl

type Config struct {
	APIURL string `json:"api_url"`
	APIKey string `json:"api_key"`
}

var (
	config *Config
)

func sendCommand(command string) {
	print(command)
}

func whitelistUser(username string) {}

func unwhitelistUser(username string) {
}
