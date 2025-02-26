package db

import (
	"fmt"
	"os"

	"github.com/medyagh/gopogh/pkg/models"
)

// FlagValues are the config values from flags
type FlagValues struct {
	Backend     string
	Host        string
	Path        string
	UseCloudSQL bool
	UseIAMAuth  bool
}

// config is database configuration
type config struct {
	dbType     string
	path       string
	host       string
	useIAMAuth bool
}

// Datab is the database interface we support
type Datab interface {
	Set(models.DBEnvironmentTest, []models.DBTestCase) error

	Initialize() error

	GetEnvironmentTestsAndTestCases() (map[string]interface{}, error)

	GetEnvCharts(string, int) (map[string]interface{}, error)

	GetOverview() (map[string]interface{}, error)

	GetTestCharts(string, string) (map[string]interface{}, error)
}

// newDB handles which database driver to use and initializes the db
func newDB(cfg config) (Datab, error) {
	switch cfg.dbType {
	case "sqlite":
		return newSQLite(cfg)
	case "postgres":
		return newPostgres(cfg)
	default:
		return nil, fmt.Errorf("unknown backend: %q", cfg.dbType)
	}
}

// FromEnv configures and returns a database instance.
// backend and path parameters are default config, otherwise gets config from the environment variables DB_BACKEND and DB_PATH
func FromEnv(fv FlagValues) (c Datab, err error) {
	backend, err := getFlagOrEnv(fv.Backend, "DB_BACKEND")
	if err != nil {
		return nil, err
	}
	path, err := getFlagOrEnv(fv.Path, "DB_PATH")
	if err != nil {
		return nil, err
	}
	host, err := getFlagOrEnv(fv.Host, "DB_HOST")
	if err != nil {
		return nil, err
	}
	cfg := config{
		dbType:     backend,
		path:       path,
		host:       host,
		useIAMAuth: fv.UseIAMAuth,
	}
	if fv.UseCloudSQL {
		c, err = NewCloudSQL(cfg)
	} else {
		c, err = newDB(cfg)
	}
	if err != nil {
		return nil, fmt.Errorf("new from %s: %s: %v", backend, path, err)
	}

	return c, nil
}

func getFlagOrEnv(flagValue, envName string) (string, error) {
	if flagValue != "" {
		return flagValue, nil
	}
	env := os.Getenv(envName)
	if env != "" {
		return env, nil
	}
	return "", fmt.Errorf("missing %s", envName)
}
