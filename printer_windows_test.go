package printer_test

import (
	"encoding/json"
	printer "github.com/gorpher/winspool-go"
	"testing"
)

func testGetDefaultPrinter(t *testing.T) printer.HANDLE {
	name, err := printer.GetDefaultPrinter()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(name)
	hPrinter, err := printer.OpenPrinter(name)
	if err != nil {
		t.Fatal(err)
	}
	return hPrinter

}

func TestGetDefaultPrinterName(t *testing.T) {
	name, err := printer.GetDefaultPrinter()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(name)
}

func TestOpenPrinter(t *testing.T) {
	hPrinter, err := printer.OpenPrinter("Canon LBP2900")
	if err != nil {
		t.Fatal(err)
	}
	err = hPrinter.ClosePrinter()
	if err != nil {
		t.Fatal(err)
	}
}

func TestEnumPrinters2(t *testing.T) {
	_, err := printer.EnumPrinters2()
	if err != nil {
		t.Fatal(err)
	}
}

func TestEnumPrinters5(t *testing.T) {
	_, err := printer.EnumPrinters5()
	if err != nil {
		t.Fatal(err)
	}
}

func TestEnumJobs(t *testing.T) {
	hPrinter, err := printer.OpenPrinter("Canon LBP2900")
	if err != nil {
		t.Fatal(err)
	}
	defer hPrinter.ClosePrinter()
	jobs1, err := hPrinter.EnumJobs1()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(len(jobs1))
	for _, job := range jobs1 {
		t.Log(job.JobID, "|", job.Document, "|", job.GetStatus())
	}
	printJson(t, jobs1)
}

func TestGetDriverInfo(t *testing.T) {
	handle := testGetDefaultPrinter(t)
	defer handle.ClosePrinter()
	driver, err := handle.Driver()
	if err != nil {
		t.Fatal(err)
	}
	printJson(t, driver)
}

func TestPrint(t *testing.T) {
	handle := testGetDefaultPrinter(t)
	defer handle.ClosePrinter()
	err := handle.Print("test", "TEXT", []byte("test\n"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestReleaseJob(t *testing.T) {
	handle := testGetDefaultPrinter(t)
	defer handle.ClosePrinter()
	jobs1, err := handle.EnumJobs1()
	if err != nil {
		t.Fatal(err)
	}
	for _, job := range jobs1 {
		err := handle.ReleaseJob(job.JobID)
		if err != nil {
			t.Fatal(err)
		}
		return
	}
}

func printJson(t *testing.T, o interface{}) {
	body, err := json.MarshalIndent(o, "", "\t")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(body))
}
