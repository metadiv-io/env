package env_test

import (
	"testing"

	"github.com/metadiv-io/env"
	"github.com/metadiv-io/env/test"
	"github.com/stretchr/testify/require"
)

func TestComplex128(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Complex128("TEST_COMPLEX128")
	require.Equal(t, complex128(456), v)
}

func TestComplex128Default(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Complex128("TEST_COMPLEX128_NOT_EXISTS", 123)
	require.Equal(t, complex128(123), v)
}

func TestComplex128Must(t *testing.T) {
	test.Util.BeforeTest()

	v := env.MustComplex128("TEST_COMPLEX128")
	require.Equal(t, complex128(456), v)
}

func TestComplex128MustPanic(t *testing.T) {
	test.Util.BeforeTest()

	require.Panics(t, func() {
		env.MustComplex128("TEST_COMPLEX128_NOT_EXISTS")
	})
}

func TestComplex128s(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Complex128s("TEST_COMPLEX128S")
	require.Equal(t, []complex128{complex128(456), complex128(789)}, v)
}

func TestComplex128sDefault(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Complex128s("TEST_COMPLEX128S_NOT_EXISTS", []complex128{complex128(456), complex128(789)})
	require.Equal(t, []complex128{complex128(456), complex128(789)}, v)
}

func TestComplex128sMust(t *testing.T) {
	test.Util.BeforeTest()

	v := env.MustComplex128s("TEST_COMPLEX128S")
	require.Equal(t, []complex128{complex128(456), complex128(789)}, v)
}

func TestComplex128sMustPanic(t *testing.T) {
	test.Util.BeforeTest()

	require.Panics(t, func() {
		env.MustComplex128s("TEST_COMPLEX128S_NOT_EXISTS")
	})
}
