package stats

import "github.com/scarfase/bank/v2/pkg/types"

// Avg рассчитывает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money {
	var sum types.Money

	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			sum += payment.Amount
		}
	}

	return types.Money(int(sum) / len(payments))
}

// TotalInCategory находит сумму покупок в определенной категории
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var sum types.Money

	for _, payment := range payments {
		if payment.Category == category && payment.Status != types.StatusFail {
			sum += payment.Amount
		}
	}
	return sum
}
