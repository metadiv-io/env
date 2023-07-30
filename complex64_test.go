package env_test

import (
	"testing"

	"github.com/metadiv-io/env"
	"github.com/metadiv-io/env/test"
	"github.com/stretchr/testify/require"
)

func TestComplex64(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Complex64("TEST_COMPLEX64")
	require.Equal(t, complex64(123), v)
}

func TestComplex64Default(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Complex64("TEST_COMPLEX64_NOT_EXISTS", 123)
	require.Equal(t, complex64(123), v)
}

func TestComplex64Must(t *testing.T) {
	test.Util.BeforeTest()

	v := env.MustComplex64("TEST_COMPLEX64")
	require.Equal(t, complex64(123), v)
}

func TestComplex64MustPanic(t *testing.T) {
	test.Util.BeforeTest()

	require.Panics(t, func() {
		env.MustComplex64("TEST_COMPLEX64_NOT_EXISTS")
	})
}

func TestComplex64s(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Complex64s("TEST_COMPLEX64S")
	require.Equal(t, []complex64{complex64(123), complex64(456)}, v)
}

func TestComplex64sDefault(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Complex64s("TEST_COMPLEX64S_NOT_EXISTS", []complex64{complex64(123), complex64(456)})
	require.Equal(t, []complex64{complex64(123), complex64(456)}, v)
}

func TestComplex64sMust(t *testing.T) {
	test.Util.BeforeTest()

	v := env.MustComplex64s("TEST_COMPLEX64S")
	require.Equal(t, []complex64{complex64(123), complex64(456)}, v)
}

func TestComplex64sMustPanic(t *testing.T) {
	test.Util.BeforeTest()

	require.Panics(t, func() {
		env.MustComplex64s("TEST_COMPLEX64S_NOT_EXISTS")
	})
}
