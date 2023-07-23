package config

type Database struct {
	DBName                string `yaml:"DBName"`
	Host                  string `yaml:"Host"`
	Port                  int    `yaml:"Port"`
	Username              string `yaml:"Username"`
	Password              string `yaml:"Password"`
	Charset               string `yaml:"Charset"`
	Timeout               int    `yaml:"Timeout"`
	MaximumConnections    int    `yaml:"MaximumConnections"`
	MinimumConnections    int    `yaml:"MinimumConnections"`
	IdleConnectionTimeout int    `yaml:"IdleConnectionTimeout"`
}
