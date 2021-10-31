package stats

import (
	"github.com/scarfase/bank/v2/pkg/types"
)

// Avg рассчитывает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money {
	var sum types.Money
	var successPayments int64

	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			sum += payment.Amount
			successPayments++
		}
	}

	return types.Money(int64(sum) / successPayments)
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

// FilterByCategory возвращает платежи в указаной катергории.
func FilterByCategory(payments []types.Payment, category types.Category) []types.Payment {
	var filtered []types.Payment
	for _, payment := range payments {
		if payment.Category == category {
			filtered = append(filtered, payment)
		}
	}
	return filtered
}

// CategoriesTotal возвращает сумму платежей по каждой категории.
func CategoriesTotal(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}

	for _, payment := range payments {
		categories[payment.Category] += payment.Amount
	}
	return categories
}

// CategoriesAvg возвращает среднюю сумму платежа по каждой категории
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	mapAvg := map[types.Category]types.Money{}
	mapCount := map[types.Category]int64{}

	consolideted := CategoriesTotal(payments)

	for _, payment := range payments {
		mapCount[payment.Category]++
	}

	for key, value := range consolideted {
		count, ok := mapCount[key]
		if ok {
			mapAvg[key] = types.Money(int64(value) / count)
		}
	}

	return mapAvg
}
