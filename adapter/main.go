package main

import "fmt"

type LegacyPrinter interface {
	Print(s string) string
}

type MyLegacyStruct struct{}

func (l *MyLegacyStruct) Print(s string) string {
	newMessage := "Legacy printer: " + s
	fmt.Println(newMessage)
	return newMessage
}

type ModernPrinter interface {
	PrintStored() string
}

type PrinterAdapter struct {
	OldPrinter LegacyPrinter
	Msg        string
}

func (p *PrinterAdapter) PrinterStored() string {
	if p.OldPrinter != nil {
		return p.OldPrinter.Print(p.Msg)
	}
	return p.Msg

}

func main() {
	legacy := &MyLegacyStruct{}
	legacy.Print("Hello Legacy!")

	adapter := PrinterAdapter{
		OldPrinter: legacy,
		Msg:        "Don't worry, be happy with adapted printer!",
	}
	adapter.PrinterStored()

}
