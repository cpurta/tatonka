package config

type Config struct {
	CassandraConfig `yaml:"cassandra"`
	GDAXConfig      *GDAXConfig `yaml:"gdax"`
	Strategies      []Strategy  `yaml:"strategies"`
}

type GDAXConfig struct {
	APIKey        string `yaml:"api_key"`
	APISecret     string `yaml:"api_secret"`
	APIPassphrase string `yaml:"api_passphrase"`
}

type CassandraConfig struct {
	Cluster  []string `yaml:"cluster"`
	Keyspace string   `yaml:"keyspace"`
}

type Strategy struct {
	Name    string                 `yaml:"name"`
	Options map[string]interface{} `yaml:"options"`
}
