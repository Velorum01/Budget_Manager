package views

import (
	"fmt"
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
	"github.com/Velorum01/Budget_Manager/validations"
)

func popUp(a fyne.App, w fyne.Window, msg string) {
	view := a.NewWindow("Message")
	view.SetContent(widget.NewLabel(msg))
	view.Show()
}

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
	priceInp := widget.NewEntry()

	title.SetPlaceHolder("Expense name:")
	priceInp.SetPlaceHolder("Expense:")

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
		priceInp,
		categories,
		widget.NewButton("Submit", func() {
			if !validations.IsNewExpenseValid(title.Text, priceInp.Text, category) {
				popUp(a, view, "Please fill all fields")
				return
			}

			price, err := strconv.ParseFloat(priceInp.Text, 64)

			if err != nil {
				popUp(a, view, "Please only enter numbers")
			}

			fmt.Println(price)
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
			if len(budgetInp.Text) == 0 {
				popUp(a, view, "Please only enter numbers")
				return
			}

			budget, err := strconv.ParseFloat(budgetInp.Text, 64)

			if err != nil {
				popUp(a, view, "Please only enter numbers")
			}

			fmt.Println(budget)
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
