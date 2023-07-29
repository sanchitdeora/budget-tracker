package exceptions

import "errors"

var ErrValidationError = errors.New("invalid request")

//transactions
var ErrNoTransactionsFound error = errors.New("no transactions found")
var ErrTransactionNotFound error = errors.New("transaction not found")

// budgets
var ErrNoBudgetsFound error = errors.New("no budgets found")