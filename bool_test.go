package env_test

import (
	"testing"

	"github.com/metadiv-io/env"
	"github.com/metadiv-io/env/test"
	"github.com/stretchr/testify/require"
)

func TestBool(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Bool("TEST_BOOL")
	require.Equal(t, true, v)
}

func TestBoolDefault(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Bool("TEST_BOOL_NOT_EXISTS", true)
	require.Equal(t, true, v)
}

func TestBoolMust(t *testing.T) {
	test.Util.BeforeTest()

	v := env.MustBool("TEST_BOOL")
	require.Equal(t, true, v)
}

func TestBoolMustPanic(t *testing.T) {
	test.Util.BeforeTest()

	require.Panics(t, func() {
		env.MustBool("TEST_BOOL_NOT_EXISTS")
	})
}

func TestBools(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Bools("TEST_BOOLS")
	require.Equal(t, []bool{true, false, true}, v)
}

func TestBoolsDefault(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Bools("TEST_BOOLS_NOT_EXISTS", []bool{false, true, false})
	require.Equal(t, []bool{false, true, false}, v)
}

func TestBoolsMust(t *testing.T) {
	test.Util.BeforeTest()

	v := env.MustBools("TEST_BOOLS")
	require.Equal(t, []bool{true, false, true}, v)
}

func TestBoolsMustPanic(t *testing.T) {
	test.Util.BeforeTest()

	require.Panics(t, func() {
		env.MustBools("TEST_BOOLS_NOT_EXISTS")
	})
}
