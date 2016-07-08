package file

import (
    "hiank.net/translate/tool"
    "github.com/tealeg/xlsx"
    "github.com/extrame/xls"
    "fmt"
)

// XLSXFile used to operate xlsx file
type XLSXFile struct {



}


// LoadXLSX used to read file in xlsx type
func LoadXLSX(path string) tool.File {


    xl, err := xlsx.OpenFile(path)
    if err != nil {

        fmt.Println("xlsx read error : " + err.Error())
        return nil
    }

    for _, sheet := range xl.Sheets {

        for _, row := range sheet.Rows {

            for _, cell := range row.Cells {

                s, _ := cell.String()
                fmt.Printf("%s_", s)
                // fmt.Printf("#%s", cell.String())
            }
            fmt.Printf("\n")
        }
    }

    return nil
}

// LoadXLS used to read file in xls type
func LoadXLS(path string) tool.File {

    xl, err := xls.Open(path, "")
    if err != nil {
        fmt.Println("xls read error : " + err.Error())
        return nil
    }

    for i := 0; i < xl.NumSheets(); i++ {
        
        sheet := xl.GetSheet(i)

        for _, row := range sheet.Rows {

           if len(row.Cols) > 0 {
               
               row.Cols.Keys
                for _, col := range row.Cols {

                    // if uint16(len(data)) <= col.LastCol() {
                    //     data = append(data, make([]string], col.LastCol()-uint16(len(data))+1)...)
                    // }
                    strArr := col.String(xl)
                    for idx, str := range strArr {

                        fmt.Printf("%v_", idx)
                    }
                    
                }
            }
            fmt.Printf("\n")
        }

    }
    return nil
    // c := xl.ReadAllCells()
    // for _, sheet := range xl.Sh
}