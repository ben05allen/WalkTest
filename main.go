package main

import (
	"fmt"
	"time"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	var mw *walk.MainWindow
	var openFileBtn, calcAgeBtn *walk.PushButton
	var filePathLabel *walk.Label
	var nameEdit, dobEdit *walk.LineEdit
    var nameLabel, dobLabel *walk.Label

	MainWindow{
		AssignTo: &mw,
		Title:   "Walk Demo",
		MaxSize: Size{Width: 400, Height: 300},
		Layout:  VBox{},
		Children: []Widget{
			Composite{
                Layout: HBox{},
                Children: []Widget{
                    Label{
                        AssignTo: &filePathLabel,
                        Text:     "File 1 path",
                    },
                    PushButton{
                        AssignTo: &openFileBtn,
                        Text:     "Open File 1",
                        OnClicked: func() {
                            dlg := new(walk.FileDialog)
                            if ok, err := dlg.ShowOpen(mw); err != nil {
                                walk.MsgBox(mw, "Error", err.Error(), walk.MsgBoxIconError)
                            } else if ok {
                                filePathLabel.SetText(dlg.FilePath)
                            }
                        },
                    },
                },
            },

			Composite{
                Layout: HBox{},
                Children: []Widget{
                    Label{AssignTo: &nameLabel, Text: "Name:"},
                    LineEdit{AssignTo: &nameEdit},
                },
            },

            Composite{
                Layout: HBox{},
                Children: []Widget{
                    Label{AssignTo: &dobLabel, Text: "Date of Birth (DD/MM/YYYY):"},
                    LineEdit{AssignTo: &dobEdit},
                },
			},

			Composite{
				Layout: HBox{},
                Children: []Widget{
                    PushButton{
                        AssignTo: &calcAgeBtn,
                        Text:     "Calculate Age",
                        OnClicked: func() {
                            name := nameEdit.Text()
                            dobText := dobEdit.Text()
                            years, err := calculateAge(dobText)
                            if err != nil {
                                walk.MsgBox(mw, "Error", "Please enter a valid date in DD/MM/YYYY format.", walk.MsgBoxIconError)
                                return
                            }
                            message := fmt.Sprintf("Hi %s, you are %d years old.", name, years)
                            walk.MsgBox(mw, "Age Calculation", message, walk.MsgBoxIconInformation)
                        },
                    },
                },
			},

			PushButton{
                Text: "Close",
                OnClicked: func() {
                    mw.Close()
                },
            },
		},
	}.Run()
}

func calculateAge(dobStr string) (int, error) {
    dob, err := time.Parse("2/1/2006", dobStr)
    if err != nil {
        return 0, err
    }

    years := time.Now().Year() - dob.Year()
    if time.Now().YearDay() < dob.YearDay() {
        years--
    }
    return years, nil
}