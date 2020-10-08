package config

// Config holds the entire application configuration.
type Config struct {
	CassandraConfig `yaml:"cassandra"`
	GDAXConfig      *GDAXConfig `yaml:"gdax"`
}

// GDAXConfig is the collection of configuration values
// to access the GDAX service.
type GDAXConfig struct {
	APIKey        string `yaml:"api_key"`
	APISecret     string `yaml:"api_secret"`
	APIPassphrase string `yaml:"api_passphrase"`
}

// CassandraConfig is the collection of configuration values
// needed to connect to a cassandra cluster.
type CassandraConfig struct {
	Cluster  []string `yaml:"cluster"`
	Keyspace string   `yaml:"keyspace"`
}
