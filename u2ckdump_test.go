package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	logInit(ioutil.Discard, ioutil.Discard, ioutil.Discard, ioutil.Discard)

	err := Parse("res/dump-2018-04-16T23:46:00+0300/dump.xml")
	require.NoError(t, err)
}

func BenchmarkParse(b *testing.B) {
	logInit(ioutil.Discard, ioutil.Discard, ioutil.Discard, ioutil.Discard)

	const file = "res/dump-2018-04-16T23:46:00+0300/dump.xml"
	fi, err := os.Stat(file)
	require.NoError(b, err)
	b.SetBytes(fi.Size())

	b.ReportAllocs()
	b.ResetTimer()

	err = Parse(file)

	b.StopTimer()

	require.NoError(b, err)
}
