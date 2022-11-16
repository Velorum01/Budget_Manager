package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
	"github.com/Velorum01/Budget_Manager/views"
)

func main() {
	a := app.New()
	w := a.NewWindow("Budget Manager")

	w.SetContent(container.NewVBox(
		widget.NewLabel("Budgegt manager"),
		widget.NewButton("This month's spending", func() {
			views.ExpenseReport(a, w)
		}),
		widget.NewButton("Add an expense", func() {
			views.AddExpense(a, w)
		}),
		widget.NewButton("Set Budget", func() {
			views.SetBudget(a, w)
		}),
	))

	w.Resize(fyne.NewSize(350, 300))
	w.ShowAndRun()
}
