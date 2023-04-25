package main

import (
	"fmt"
)

func main() {
	fmt.Println("Три числа. Ещё попытка.")

	//переменная создана для удобства смены условий программы.
	numRequired := 5
	fmt.Println("Эта программа определит, есть ли среди ваших трех чисел, количество чисел, которые больше или равны", numRequired)

	fmt.Println("Введите первое число:")
	var numOne int
	fmt.Scan(&numOne)

	fmt.Println("Введите второе число:")
	var numTwo int
	fmt.Scan(&numTwo)

	fmt.Println("Введите третье число:")
	var numThree int
	fmt.Scan(&numThree)

	//переменная, в которую будет "складываться" подходящее по условию количество чисел при проверке.
	numBasket := 0

	if numOne >= numRequired {
		numBasket++

		if numTwo >= numRequired {
			numBasket++

			if numThree >= numRequired {
				numBasket++
				fmt.Println("Среди введённых чисел", numBasket, "больше или равны", numRequired)
			} else {
				fmt.Println("Среди введённых чисел", numBasket, "больше или равны", numRequired)
			}
		} else if numThree >= numRequired {
			numBasket++
			fmt.Println("Среди введённых чисел", numBasket, "больше или равны", numRequired)
		} else {
			fmt.Println("Среди введённых чисел", numBasket, "больше или равны", numRequired)
		}
	} else if numTwo >= numRequired {
		numBasket++

		if numThree >= numRequired {
			numBasket++
			fmt.Println("Среди введённых чисел", numBasket, "больше или равны", numRequired)
		} else {
			fmt.Println("Среди введённых чисел", numBasket, "больше или равны", numRequired)
		}
	} else if numThree >= numRequired {
		numBasket++
		fmt.Println("Среди введённых чисел", numBasket, "больше или равны", numRequired)
	} else {
		fmt.Println("Среди введённых чисел нет чисел больше или равных", numRequired)
	}
}
