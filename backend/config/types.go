package config

type Config struct {
	App      appConfig
	Security securityConfig
	Logger   loggerConfig
	Database databaseConfig
}

type appConfig struct {
	Name         string
	Version      string
	Domain       string
	Host         string
	Port         int
	Env          string
	Debug        bool
	ReadTimeout  int
	WriteTimeout int
	SSL          bool
	Prefork      bool
}

type securityConfig struct {
	Cors      corsConfig
	Csrf      csrfConfig
	RateLimit rateLimitConfig
	Cookie    cookieConfig
}

type corsConfig struct {
	AllowedOrigins   string
	AllowedMethods   string
	AllowCredentials bool
}

type csrfConfig struct {
	Enabled    bool
	CookieName string
	HeaderName string
}

type rateLimitConfig struct {
	Duration    int
	MaxRequests int
}

type loggerConfig struct {
	Level  int
	Pretty bool
}

type databaseConfig struct {
	ConnMaxIdleTime int
	ConnMaxLifetime int
	MaxIdleCons     int
	MaxOpenCons     int
	File            string
	DryRun          bool
}

type cookieConfig struct {
	Name     string
	Secure   bool
	HttpOnly bool
	SameSite string
	Domain   string
	MaxAge   int
	Key      string
}
