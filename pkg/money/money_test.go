package money

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type moneyValTestCase struct {
	money    Money
	expected Money
}

func TestMoney_Val(t *testing.T) {
	for _, c := range casesMoneyVal() {
		assert.Equal(t, c.expected, c.money.Val())
	}
}

func casesMoneyVal() []moneyValTestCase {
	return []moneyValTestCase{
		{100, 100},
		{100.1, 100.1},
		{100.12, 100.12},
		{100.123, 100.123},
		{100.1234, 100.1234},
		{100.12345, 100.1235},
		{100.12344, 100.1234},
	}
}
