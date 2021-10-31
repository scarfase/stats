package stats

import (
	"fmt"

	"github.com/scarfase/bank/v2/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			Amount: 10_000_00,
			Status: "OK",
		},
		{
			Amount: 10_000_00,
			Status: "OK",
		},
		{
			Amount: 10_000_00,
			Status: "INPROGRESS",
		}, {
			Amount: 10_000_00,
			Status: "FAIL",
		},
	}

	fmt.Println(Avg(payments))
	//Output: 750000

}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			Amount:   10_000_00,
			Category: "bank",
			Status:   "OK",
		},
		{
			Amount:   10_000_00,
			Category: "bank",
			Status:   "INPROGRESS",
		},
		{
			Amount:   10_000_00,
			Category: "bank",
			Status:   "FAIL",
		},
		{
			Amount:   10_000_00,
			Category: "gas",
			Status:   "OK",
		},
		{
			Amount:   10_000_00,
			Category: "playgroud",
			Status:   "OK",
		},
		{
			Amount:   10_000_00,
			Category: "playgroud",
			Status:   "INPROGRESS",
		},
		{
			Amount:   10_000_00,
			Category: "bank",
			Status:   "FAIL",
		},
	}

	fmt.Println(TotalInCategory(payments, "bank"))
	//Output: 2000000

}
