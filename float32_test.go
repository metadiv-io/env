package env_test

import (
	"testing"

	"github.com/metadiv-io/env"
	"github.com/metadiv-io/env/test"
	"github.com/stretchr/testify/require"
)

func TestFloat32(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Float32("TEST_FLOAT32")
	require.Equal(t, float32(3.2), v)
}

func TestFloat32Default(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Float32("TEST_FLOAT32_NOT_EXISTS", 3.2)
	require.Equal(t, float32(3.2), v)
}

func TestFloat32Must(t *testing.T) {
	test.Util.BeforeTest()

	v := env.MustFloat32("TEST_FLOAT32")
	require.Equal(t, float32(3.2), v)
}

func TestFloat32MustPanic(t *testing.T) {
	test.Util.BeforeTest()

	require.Panics(t, func() {
		env.MustFloat32("TEST_FLOAT32_NOT_EXISTS")
	})
}

func TestFloat32s(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Float32s("TEST_FLOAT32S")
	require.Equal(t, []float32{3.2, 2.3}, v)
}

func TestFloat32sDefault(t *testing.T) {
	test.Util.BeforeTest()

	v := env.Float32s("TEST_FLOAT32S_NOT_EXISTS", []float32{3.2, 2.3})
	require.Equal(t, []float32{3.2, 2.3}, v)
}

func TestFloat32sMust(t *testing.T) {
	test.Util.BeforeTest()

	v := env.MustFloat32s("TEST_FLOAT32S")
	require.Equal(t, []float32{3.2, 2.3}, v)
}

func TestFloat32sMustPanic(t *testing.T) {
	test.Util.BeforeTest()

	require.Panics(t, func() {
		env.MustFloat32s("TEST_FLOAT32S_NOT_EXISTS")
	})
}
