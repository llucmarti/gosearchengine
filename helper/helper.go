package helper

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
)

func GetNextPage(nPage int, perPage int, nProducts int) int {
	if (nPage+1)*perPage >= nProducts {
		return -1
	}
	return nPage + 1
}

func GenerateMaterialID(materialName string) string {
	name := NormalizeString(materialName)
	hash := sha256.Sum256([]byte(name))
	code := hex.EncodeToString(hash[:])[:4]
	return strings.ToUpper(code)
}

func NormalizeString(s string) string {
	return strings.ToLower(strings.ReplaceAll(s, " ", ""))
}

func CalculatePrice(s string) float64 {
	price, _ := strconv.ParseFloat(s, 64)
	return price / 100
}

func CalculateAmount(s string) int {
	amount, _ := strconv.Atoi(s)
	return amount
}
