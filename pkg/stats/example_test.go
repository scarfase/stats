package stats

import (
	"fmt"

	"github.com/scarfase/bank/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			Amount: 10_000_00,
		},
		{
			Amount: 10_000_00,
		},
		{
			Amount: 10_000_00,
		},
	}

	fmt.Println(Avg(payments))
	//Output: 1000000

}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			Amount:   10_000_00,
			Category: "bank",
		},
		{
			Amount:   10_000_00,
			Category: "bank",
		},
		{
			Amount:   10_000_00,
			Category: "gas",
		},
		{
			Amount:   10_000_00,
			Category: "playgroud",
		},
	}

	fmt.Println(TotalInCategory(payments, "bank"))
	//Output: 2000000

}
