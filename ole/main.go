package main

import (
	"fmt"
	"path/filepath"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

type Word struct {
    app       *ole.IDispatch
    documents *ole.VARIANT
    doc       *ole.VARIANT
}

func (wd *Word) open(inFile string) (err error) {
    ole.CoInitialize(0)

    var unknown *ole.IUnknown

    unknown, err = oleutil.CreateObject("Word.Application")
    if err != nil {
        return
    }

    wd.app, err = unknown.QueryInterface(ole.IID_IDispatch)
    if err != nil {
        return
    }

    _, err = oleutil.PutProperty(wd.app, "Visible", false)
    if err != nil {
        return
    }

    _, err = oleutil.PutProperty(wd.app, "DisplayAlerts", 0)
    if err != nil {
        return
	}
	
    wd.documents, err = oleutil.GetProperty(wd.app, "Documents")
    if err != nil {
        return
    }

    wd.doc, err = oleutil.CallMethod(wd.documents.ToIDispatch(), "Open", inFile)
    if err != nil {
        return
    }

    return
}

// Convert 转换
func (wd *Word) Convert(inFile, outDir string) (outFile string, err error) {

    outFile = filepath.Join(outDir, filepath.Base(inFile+".html"))

    defer func() {
        if err != nil {
            outFile = ""
        }
        wd.close()
    }()

    err = wd.open(inFile)
    if err != nil {
        return
    }

    _, err = oleutil.CallMethod(wd.doc.ToIDispatch(), "ExportAsFixedFormat", outFile, 17)
    if err != nil {
        return
    }

    return
}

func (wd *Word) close() {

    if wd.doc != nil {
        oleutil.MustPutProperty(wd.doc.ToIDispatch(), "Saved", true)
    }

    if wd.documents != nil {
        oleutil.MustCallMethod(wd.documents.ToIDispatch(), "Close")
    }

    if wd.app != nil {
        oleutil.MustCallMethod(wd.app, "Quit")
        wd.app.Release()
    }

    ole.CoUninitialize()
}


func main() {
	word := Word{}

	sample := "E:\\GO\\workspace\\examples\\ole\\test\\example.docx"
	export := "E:\\GO\\workspace\\examples\\ole\\test"

	_, err := word.Convert(sample, export)
	fmt.Println(err)
}