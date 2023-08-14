package Calculate

import (
	"fmt"
	"strconv"
	"workAPP/types"
)

func Orderinventory(item, amount string, a *[]string) {
	amountfloat, _ := strconv.ParseFloat(amount, 64)
	fmt.Println("str", types.Centralmap[item])
	if amountfloat <= types.Centralmap[item] {
		sentance := fmt.Sprintf("%s一箱,", item)
		*a = append(*a, sentance)
	}
}
