package views

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func ExpenseReport(a fyne.App, w fyne.Window) {
	view := a.NewWindow("Expense Report")
	view.SetMaster()
	w.Hide()

	view.SetContent(container.NewVBox(
		widget.NewLabel("Food"),
		widget.NewLabel("Groceries"),
		widget.NewLabel("Travel"),
		widget.NewLabel("Hobbies"),
		widget.NewLabel("Others"),
		widget.NewButton("Home", func() {
			w.SetMaster()
			view.Hide()
			w.Show()
		}),
	))

	view.Resize(fyne.NewSize(350, 300))
	view.Show()
}

func AddExpense(a fyne.App, w fyne.Window) {
	view := a.NewWindow("Add an Expense")
	view.SetMaster()
	w.Hide()

	var category string

	title := widget.NewEntry()
	price := widget.NewEntry()

	title.SetPlaceHolder("Expense name:")
	price.SetPlaceHolder("Expense:")

	categories := widget.NewRadioGroup([]string{
		"Food",
		"Groceries",
		"Travel",
		"Hobbies",
		"Others",
	}, func(val string) {
		category = val
	})

	view.SetContent(container.NewVBox(
		title,
		price,
		categories,
		widget.NewButton("Submit", func() {
			fmt.Println(category)
		}),

		widget.NewButton("Home", func() {
			w.SetMaster()
			view.Hide()
			w.Show()
		}),
	))

	view.Resize(fyne.NewSize(350, 300))
	view.Show()
}

func SetBudget(a fyne.App, w fyne.Window) {
	view := a.NewWindow("Set Budget")
	view.SetMaster()
	w.Hide()

	budgetInp := widget.NewEntry()
	budgetInp.SetPlaceHolder("Budget:")

	view.SetContent(container.NewVBox(
		budgetInp,
		widget.NewButton("Set Budget", func() {

		}),

		widget.NewButton("Home", func() {
			w.SetMaster()
			view.Hide()
			w.Show()
		}),
	))

	view.Resize(fyne.NewSize(350, 300))
	view.Show()
}
