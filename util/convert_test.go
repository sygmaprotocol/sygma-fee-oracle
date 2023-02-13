// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package util_test

import (
	"github.com/ChainSafe/sygma-fee-oracle/util"

	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStr2BigInt(t *testing.T) {
	type str2BigIntTest struct {
		name   string
		input  string
		output *big.Int
		errMsg string
	}

	testcases := make([]str2BigIntTest, 0)
	testcases = append(testcases, str2BigIntTest{
		name:   "valid string as input 1",
		input:  "100",
		output: big.NewInt(100),
		errMsg: "",
	})
	testcases = append(testcases, str2BigIntTest{
		name:   "valid string as input 2",
		input:  "123456789",
		output: big.NewInt(123456789),
		errMsg: "",
	})
	testcases = append(testcases, str2BigIntTest{
		name:   "invalid string as input",
		input:  "abc",
		output: nil,
		errMsg: "failed to convert from string to bigInt",
	})
	testcases = append(testcases, str2BigIntTest{
		name:   "invalid number as input(not int)",
		input:  "100.123",
		output: nil,
		errMsg: "failed to convert from string to bigInt",
	})
	testcases = append(testcases, str2BigIntTest{
		name:   "invalid number as input(not int)",
		input:  "0.123",
		output: nil,
		errMsg: "failed to convert from string to bigInt",
	})

	for _, testcase := range testcases {
		re, err := util.Str2BigInt(testcase.input)
		if err != nil {
			if testcase.errMsg == "" {
				assert.Fail(t, "error not equal", testcase.name)
			}
			assert.EqualError(t, err, testcase.errMsg, testcase.name)
		}
		assert.EqualValues(t, testcase.output, re, testcase.name)
	}
}

func TestLarge2SmallUnitConverter(t *testing.T) {
	type large2SmallUnitConverterTest struct {
		name      string
		input1    string
		input2    uint
		output    *big.Int
		outputErr error
	}

	testcases := make([]large2SmallUnitConverterTest, 0)
	testcases = append(testcases, large2SmallUnitConverterTest{
		name:      "valid string as input 1",
		input1:    "100",
		input2:    2,
		output:    big.NewInt(10000),
		outputErr: nil,
	})
	testcases = append(testcases, large2SmallUnitConverterTest{
		name:      "valid string as input 2",
		input1:    "1",
		input2:    18,
		output:    big.NewInt(1000000000000000000),
		outputErr: nil,
	})
	testcases = append(testcases, large2SmallUnitConverterTest{
		name:      "valid string as input 3",
		input1:    "1",
		input2:    9,
		output:    big.NewInt(1000000000),
		outputErr: nil,
	})
	testcases = append(testcases, large2SmallUnitConverterTest{
		name:      "valid string as input 4",
		input1:    "100",
		input2:    0,
		output:    big.NewInt(100),
		outputErr: nil,
	})
	testcases = append(testcases, large2SmallUnitConverterTest{
		name:      "short decimal number",
		input1:    "0.12",
		input2:    18,
		output:    big.NewInt(120000000000000000),
		outputErr: nil,
	})
	testcases = append(testcases, large2SmallUnitConverterTest{
		name:      "long decimal number",
		input1:    "0.00000012512521",
		input2:    18,
		output:    big.NewInt(125125210000),
		outputErr: nil,
	})
	testcases = append(testcases, large2SmallUnitConverterTest{
		name:      "large exponent",
		input1:    "0.000000000000000012512521",
		input2:    22,
		output:    big.NewInt(125125),
		outputErr: nil,
	})

	for _, testcase := range testcases {
		re, err := util.Large2SmallUnitConverter(testcase.input1, testcase.input2)
		assert.EqualValues(t, testcase.output, re, testcase.name)
		assert.EqualValues(t, testcase.outputErr, err, testcase.name)
	}
}
