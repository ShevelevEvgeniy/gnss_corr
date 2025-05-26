package modules

type Service struct {
	Transport string `envconfig:"SERVER_TRANSPORT"    default:"tcp" logKey:"transport"`
	Port      int    `envconfig:"SERVER_PORT"         default:"80"  logKey:"port"`
}
