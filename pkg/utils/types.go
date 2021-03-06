package utils

// ContextKey defines a type for context keys shared in the app
type ContextKey string

// ServerConfig defines the configuration for the server
type ServerConfig struct {
	Host          string
	Port          string
	URISchema     string
	Version       string
	SessionSecret string
	JWT           JWTConfig
	GraphQL       GQLConfig
	Database      DBConfig
	AuthProviders []AuthProvider
}

//JWTConfig defines the options for JWT tokens
type JWTConfig struct {
	Secret    string
	Algorithm string
}

// GQLConfig defines the configuration for the GQL Server
type GQLConfig struct {
	Path                string
	PlaygroundPath      string
	IsPlaygroundEnabled bool
}

// DBConfig defines the configuration for the DB config
type DBConfig struct {
	Dialect     string
	DSN         string
	SeedDB      bool
	LogMode     bool
	AutoMigrate bool
}

// AuthProvider defines the configuration for the Goth config
type AuthProvider struct {
	Provider  string
	ClientKey string
	Secret    string
	Domain    string // If needed, like with auth0
	Scopes    []string
}

// ListenEndpoint builds the endpoint string (host + port)
func (s *ServerConfig) ListenEndpoint() string {
	if s.Port == "80" {
		return s.Host
	}
	return s.Host + ":" + s.Port
}

// VersionedEndpoint builds the endpoint string (host + port + version)
func (s *ServerConfig) VersionedEndpoint(path string) string {
	return "/" + s.Version + path
}

// SchemaVersionedEndpoint builds the schema endpoint string (schema + host + port + version)
func (s *ServerConfig) SchemaVersionedEndpoint(path string) string {
	if s.Port == "80" {
		return s.URISchema + s.Host + "/" + s.Version + path
	}
	return s.URISchema + s.Host + ":" + s.Port + "/" + s.Version + path
}
