package string_utils

import (
	"fmt"
	"github.com/stader-labs/stader-node/stader-lib/types"
	"math/big"
	"strings"
)

func StringifyArray(arr []*big.Int) string {
	var strArr []string
	for _, v := range arr {
		strArr = append(strArr, v.String())
	}
	return strings.Join(strArr, ",")
}

// write a method given a comma seperated value of string numbers, get the big Ints back
func DestringifyArray(arr string) ([]*big.Int, error) {
	var intArr []*big.Int
	for _, v := range strings.Split(arr, ",") {
		i, ok := new(big.Int).SetString(v, 10)
		if !ok {
			return nil, fmt.Errorf("could not parse string to big int: %s", v)
		}
		intArr = append(intArr, i)
	}
	return intArr, nil
}

// Stringify array of types.ValidatorPubKeys
func StringifyValidatorPubKeys(arr []types.ValidatorPubkey) string {
	var strArr []string
	for _, v := range arr {
		strArr = append(strArr, v.String())
	}
	return strings.Join(strArr, ",")
}

// Destringify string of validator pubkeys to array of types.ValidatorPubKeys
func DestringifyValidatorPubKeys(arr string) []types.ValidatorPubkey {
	var keyArr []types.ValidatorPubkey
	for _, v := range strings.Split(arr, ",") {
		i := types.BytesToValidatorPubkey([]byte(v))
		keyArr = append(keyArr, i)
	}
	return keyArr
}

func DestringifyStringArray(arr string) []string {
	return strings.Split(arr, ",")
}
