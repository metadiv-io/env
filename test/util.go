package test

import "os"

var Util = new(util)

type util struct{}

func (t *util) BeforeTest() {
	os.Setenv("GIN_MODE", "release")

	os.Setenv("TEST_BOOL", "true")
	os.Setenv("TEST_BOOLS", "true,false,true")

	os.Setenv("TEST_COMPLEX64", "123")
	os.Setenv("TEST_COMPLEX64S", "123,456")

	os.Setenv("TEST_COMPLEX128", "456")
	os.Setenv("TEST_COMPLEX128S", "456,789")

	os.Setenv("TEST_FLOAT32", "3.2")
	os.Setenv("TEST_FLOAT32S", "3.2,2.3")

	os.Setenv("TEST_FLOAT64", "6.4")
	os.Setenv("TEST_FLOAT64S", "6.4,4.6")

	os.Setenv("TEST_INT", "10")
	os.Setenv("TEST_INTS", "10,100")

	os.Setenv("TEST_INT8", "1")
	os.Setenv("TEST_INT8S", "1,10")

	os.Setenv("TEST_INT16", "10")
	os.Setenv("TEST_INT16S", "10,100")

	os.Setenv("TEST_INT32", "100")
	os.Setenv("TEST_INT32S", "100,1000")

	os.Setenv("TEST_INT64", "1000")
	os.Setenv("TEST_INT64S", "1000,10000")

	os.Setenv("TEST_STRING", "test")
	os.Setenv("TEST_STRINGS", "test,test2")

	os.Setenv("TEST_UINT", "1")
	os.Setenv("TEST_UINTS", "1,2")

	os.Setenv("TEST_UINT8", "1")
	os.Setenv("TEST_UINT8S", "1,2")

	os.Setenv("TEST_UINT16", "1")
	os.Setenv("TEST_UINT16S", "1,2")

	os.Setenv("TEST_UINT32", "1")
	os.Setenv("TEST_UINT32S", "1,2")

	os.Setenv("TEST_UINT64", "1")
	os.Setenv("TEST_UINT64S", "1,2")
}
