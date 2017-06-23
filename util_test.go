package buffstreams

import (
	"testing"
)

func TestMessageSizeToBitLength(t *testing.T) {
	cases := []struct {
		input, output int
	}{
		{1, 2},
		{2, 2},
		{4, 2},
		{8, 2},
		{32, 2},
		{64, 2},
		{255, 2},
		{256, 3},
		{257, 3},
		{512, 3},
		{2048, 3},
		{4096, 3},
		{8192, 3},
	}

	for _, c := range cases {
		length := messageSizeToBitLength(c.input)
		if length != c.output {
			t.Errorf("Bit Length incorrect. For message size %d, got %d, expected %d", c.input, length, c.output)
		}
	}
}

func TestMessageBytesToInt(t *testing.T) {
	cases := []struct {
		input, output int64
	}{
		{1, 2},
		{2, 3},
		{4, 5},
		{16, 17},
		{32, 33},
		{64, 65},
		{128, 129},
		{256, 257},
		{1024, 1025},
		{2048, 2049},
		{4096, 4097},
		{8192, 8193},
		{17, 18},
		{456, 457},
		{24569045, 24569046},
	}

	for _, c := range cases {
		byteSize := messageSizeToBitLength(int(c.input))
		bytes := intToByteArray(c.input, byteSize, 1)
		result, _ := byteArrayToUInt32(bytes)
		if int64(result) != c.output {
			t.Errorf("Conversion between bytes incorrect. Original value %d, got %d", c.input, result)
		}
	}
}
