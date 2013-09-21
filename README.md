# gomoney

This library makes handling money in Go easy. Heavily inspired by
https://github.com/RubyMoney/money

## Features

* Simplify interaction with money in Go.
* Currency agnostic.

## Downloading

Install with:

    go install github.com/zachlatta/gomoney

Use in project with:

    import "github.com/zachlatta/gomoney"

## Usage

``` go
package main

import (
  "fmt"
  "github.com/zachlatta/gomoney"
)

func main() {
  money := gomoney.NewMoney(btc())
  fmt.Println(money) // Prints 0Ƀ

  moneyWithInitialValue := gomoney.NewMoneyWithValue(15, btc())
  fmt.Println(moneyWithInitialValue) // Prints 15Ƀ

  money.Value += 42.4238
  fmt.Println(money) // Prints 42.4238Ƀ

  // ...and more! Check the source.
}

func btc() *gomoney.Currency {
  return gomoney.NewCurrency(
    "BTC",
    "Bitcoin",
    "Ƀ",
    false,
    "Satoshi",
    1000000,
    ",",
    ".")
}
```
