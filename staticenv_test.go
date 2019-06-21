package staticenv

import (
	"os"
	"testing"
	"time"
)

func TestEnv(t *testing.T) {
	os.Setenv("TEST", "TESTVAL")

	env := NewEnv()
	tval := env.Getenv("DEFAULT", "TEST")
	if tval != "TESTVAL" {
		t.Errorf("Expected TESTVAL, but got %s", tval)
	}
}

func TestEnvDefault(t *testing.T) {
	env := NewEnv()
	tval := env.Getenv("DEFAULT", "UNLIKELYTOEXIST")
	if tval != "DEFAULT" {
		t.Errorf("Expected DEFAULT, but got %s", tval)
	}
}

func TestWithPrefix(t *testing.T) {
	os.Setenv("PREFIX_TEST", "PREFIXTESTVAL")

	env := NewEnvWithPrefix("PREFIX")
	tval := env.Getenv("DEFAULT", "TEST")
	if env.Prefix != "PREFIX" {
		t.Errorf("Expected a prefix of PREFIX, but got %s", env.Prefix)
	}
	if tval != "PREFIXTESTVAL" {
		t.Errorf("Expected PREFIXTESTVAL, but got %s", tval)
	}
}

func TestSetPrefix(t *testing.T) {
	os.Setenv("PREFIX_TEST", "PREFIXTESTVAL")

	env := NewEnv()
	env.SetPrefix("PREFIX")
	tval := env.Getenv("DEFAULT", "TEST")
	if env.Prefix != "PREFIX" {
		t.Errorf("Expected a prefix of PREFIX, but got %s", env.Prefix)
	}
	if tval != "PREFIXTESTVAL" {
		t.Errorf("Expected PREFIXTESTVAL, but got %s", tval)
	}
}

func TestIntEnv(t *testing.T) {
	os.Setenv("INTTEST", "3")

	env := NewEnv()
	tval := env.GetInt(1, "INTTEST")
	if tval != 3 {
		t.Errorf("Expected 3, but got %d", tval)
	}
}

func TestIntEnvDefault(t *testing.T) {
	env := NewEnv()
	tval := env.GetInt(1, "UNLIKELYTOEXIST")
	if tval != 1 {
		t.Errorf("Expected 1, but got %d", tval)
	}
}

func TestFloatEnv(t *testing.T) {
	os.Setenv("FLOATTEST", "3.14159")
	env := NewEnv()
	tval := env.GetFloat(1.2, "FLOATTEST")
	if tval != 3.14159 {
		t.Errorf("Expected 3.14159, but got %f", tval)
	}
}

func TestFloatEnvDefault(t *testing.T) {
	env := NewEnv()
	tval := env.GetFloat(1.2, "UNLIKELYTOEXIST")
	if tval != 1.2 {
		t.Errorf("Expected 1.2, but got %f", tval)
	}
}

func TestBoolEnv(t *testing.T) {
	os.Setenv("BOOLTEST", "T")
	env := NewEnv()
	tval := env.GetBool(false, "BOOLTEST")
	if !tval {
		t.Errorf("Expected true, but got %t", tval)
	}
}

func TestBoolEnvDefault(t *testing.T) {
	env := NewEnv()
	tval := env.GetBool(true, "UNLIKELYTOEXIST")
	if !tval {
		t.Errorf("Expected true, but got %t", tval)
	}
}

func TestDurationEnv(t *testing.T) {
	os.Setenv("DURATIONTEST", "3m5s")
	env := NewEnv()
	tval := env.GetDuration(0, "DURATIONTEST")
	if float64(tval) != 185000000000.0 {
		t.Errorf("Expected 185000000000.0, but got %f", float64(tval))
	}
}

func TestDurationEnvDefault(t *testing.T) {
	env := NewEnv()
	d, err := time.ParseDuration("1m5s")
	if err != nil {
		t.Error(err)
	}
	tval := env.GetDuration(d, "UNLIKELYTOEXIST")
	if float64(tval) != 65000000000.0 {
		t.Errorf("Expected 65000000000.0, but got %f", float64(tval))
	}
}

func TestTimeEnv(t *testing.T) {
	os.Setenv("TIMETEST", "2018-06-07 13:10:20")
	env := NewEnv()
	tval := env.GetTime("2006-01-02 15:04:05", time.Now(), "TIMETEST")
	if tval.String() != "2018-06-07 13:10:20 +0000 UTC" {
		t.Errorf("Expected 2018-06-07 13:10:20 +0000 UTC, but got %v", tval)
	}
}

func TestTimeEnvDefault(t *testing.T) {
	env := NewEnv()
	now := time.Now()
	tval := env.GetTime("2006-01-02 15:04:05", now, "UNLIKELYTOEXIST")
	if tval != now {
		t.Errorf("Expected %v, but got %v", now, tval)
	}
}
