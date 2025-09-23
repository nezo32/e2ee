package security

import (
	"hash/fnv"
)

func HashSum(data []byte) uint32 {
	h := fnv.New32a()
	h.Write(data)
	return h.Sum32()
}
