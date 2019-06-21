package staticenv

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Env struct {
	Prefix string
}

// NewEnv returns a pointer to a new Env.
func NewEnv() *Env {
	return &Env{}
}

// NewEnvWithPrefix returns a pointer to a new Env with the prefix set to `prefix`.
func NewEnvWithPrefix(prefix string) *Env {
	env := Env{}
	env.Prefix = prefix
	return &env
}

// SetPrefix sets the prefix for the environment.
func (env *Env) SetPrefix(prefix string) {
	env.Prefix = prefix
}

// Load a .env configuration file if it exists, and add its configuration to the environment.
func (env *Env) Load() error {
	curdir, err := os.Getwd()
	if err != nil {
		return err
	}
	p := path.Join(curdir, ".env")
	if f, err := os.Stat(p); err == nil && !f.IsDir() {
		err = godotenv.Load(p)
		return err
	}
	return errors.New(fmt.Sprintf("Could not open .env file in current directory: %s", curdir))
}

// Read a .env configuration file if it exists, and return the configuration as a map[string]string.
func (env *Env) Read() (map[string]string, error) {
	var m map[string]string
	curdir, err := os.Getwd()
	if err != nil {
		return m, err
	}
	p := path.Join(curdir, ".env")
	if f, err := os.Stat(p); err == nil && !f.IsDir() {
		m, err = godotenv.Read(p)
		return m, err
	}
	return m, errors.New(fmt.Sprintf("Could not read .env file in current directory: %s", curdir))
}

// Getenv returns the value of first valid environment variable, or the default provided value.
func (env *Env) Getenv(dfault string, v ...string) string {
	var value string
	for _, name := range v {
		if env.Prefix != "" {
			name = env.Prefix + "_" + name
		}
		value = os.Getenv(name)
		if value != "" {
			return value
		}
	}
	return dfault
}

// GetInt returns the value of the first valid environment variable as an integer, or the default provided value.
func (env *Env) GetInt(dfault int, v ...string) int {
	var value string
	for _, name := range v {
		if env.Prefix != "" {
			name = env.Prefix + "_" + name
		}
		value = os.Getenv(name)
		if value != "" {
			break
		}
	}
	if value == "" {
		return dfault
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return dfault
	}
	return intValue
}

// GetFloat returns the value of the first valid environment variable as a 64 bit float, or the default provided
// value.
//
// If the value is well-formed and near a valid floating point number, it will return the
// nearest floating point number rounded using IEEE754 unbiased rounding.
//
// If the value is not syntactically well-formed, it will return the provided default value.
//
// If the value is well-formed but is more than 1/2 ULP away from the largest 64 bit floating
// point number, it will return the provided default value.
func (env *Env) GetFloat(dfault float64, v ...string) float64 {
	var value string
	for _, name := range v {
		if env.Prefix != "" {
			name = env.Prefix + "_" + name
		}
		value = os.Getenv(name)
		if value != "" {
			break
		}
	}
	if value == "" {
		return dfault
	}
	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return dfault
	}
	return floatValue
}

// GetBool returns the value of the first valid environment variable as a bool, or it will return the
// provided default value. It accepts 1, t, T, TRUE, true, True, 0, f, F, False, false, False.
// If any other value is found, it will return the provided default value.
func (env *Env) GetBool(dfault bool, v ...string) bool {
	var value string
	for _, name := range v {
		if env.Prefix != "" {
			name = env.Prefix + "_" + name
		}
		value = os.Getenv(name)
		if value != "" {
			break
		}
	}
	if value == "" {
		return dfault
	}
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return dfault
	}
	return boolValue
}

// GetDuration returns the value of the first valid environment variable as a time.Duration, or it will return
// the provided default value.
//
// A duration is a possibly signed sequence of decimal numbers, each with optional fraction and a unit
// suffix, such as "300ms", "-1.5h" or "2h45m". Valid time units are "ns", "us" (or "µs"), "ms", "s",
// "m", "h".
func (env *Env) GetDuration(dfault time.Duration, v ...string) time.Duration {
	var value string
	for _, name := range v {
		if env.Prefix != "" {
			name = env.Prefix + "_" + name
		}
		value = os.Getenv(name)
		if value != "" {
			break
		}
	}
	if value == "" {
		return dfault
	}
	durationValue, err := time.ParseDuration(value)
	if err != nil {
		return dfault
	}
	return durationValue
}

// GetTime returns the value of the first environment variable as a time.Time, or it will return the
// provided default value. It also takes a layout string to parse the time.Time. See
// https://golang.org/pkg/time/#Parse for more information about the layout's format.
func (env *Env) GetTime(layout string, dfault time.Time, v ...string) time.Time {
	var value string
	for _, name := range v {
		if env.Prefix != "" {
			name = env.Prefix + "_" + name
		}
		value = os.Getenv(name)
		if value != "" {
			break
		}
	}
	if value == "" {
		return dfault
	}
	timeValue, err := time.Parse(layout, value)
	if err != nil {
		return dfault
	}
	return timeValue
}
