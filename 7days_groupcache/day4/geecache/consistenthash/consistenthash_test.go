package consistenthash

import (
	"strconv"
	"testing"
)

func TestHashing(t *testing.T) {
	hash := New(3, func(key []byte) uint32 {
		i, _ := strconv.Atoi(string(key))
		return uint32(i)
	})

	// 06, 16, 26, 04, 14, 24, 02, 12, 22
	hash.Add("6", "4", "2")

	testCases := map[string]string{
		"16": "6",
		"22": "2",
		"4":  "4",
		"26": "6",
		"7":  "2",
	}

	for k, v := range testCases {
		if hash.Get(k) != v {
			t.Errorf("Hash %s expect %s got %s", k, v, hash.Get(k))
		}
	}

	// 08, 18, 28
	hash.Add("8")

	testCases["28"] = "8"
	testCases["7"] = "8"
	testCases["27"] = "8"

	for k, v := range testCases {
		if hash.Get(k) != v {
			t.Errorf("Hash %s expect %s got %s", k, v, hash.Get(k))
		}
	}
}
