package env_test

import (
	"testing"

	"github.com/metadiv-io/env"
	"github.com/metadiv-io/env/test"
	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	test.Util.BeforeTest()

	v := env.String("TEST_STRING")
	require.Equal(t, string("test"), v)
}

func TestStringDefault(t *testing.T) {
	test.Util.BeforeTest()

	v := env.String("TEST_STRING_NOT_EXISTS", "test")
	require.Equal(t, string("test"), v)
}

func TestStringMust(t *testing.T) {
	test.Util.BeforeTest()

	v := env.MustString("TEST_STRING")
	require.Equal(t, string("test"), v)
}

func TestStringMustPanic(t *testing.T) {
	test.Util.BeforeTest()

	require.Panics(t, func() {
		env.MustString("TEST_STRING_NOT_EXISTS")
	})
}

func TestStrings(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Strings("TEST_STRINGS")
	require.Equal(t, []string{"test", "test2"}, v)
}

func TestStringsDefault(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Strings("TEST_STRINGS_NOT_EXISTS", []string{"test", "test2"})
	require.Equal(t, []string{"test", "test2"}, v)
}

func TestStringsMust(t *testing.T) {
	test.Util.BeforeTest()

	v := env.MustStrings("TEST_STRINGS")
	require.Equal(t, []string{"test", "test2"}, v)
}

func TestStringsMustPanic(t *testing.T) {
	test.Util.BeforeTest()

	require.Panics(t, func() {
		env.MustStrings("TEST_STRINGS_NOT_EXISTS")
	})
}
