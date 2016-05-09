package translate

import (
	"os"
	"fmt"
	"bufio"
	"io"
)

func Trans(frame Frame) {


	for info := frame.nextFile(); info != nil; info = frame.nextFile() {

		file, err := os.Open(info.name())
		if err != nil {
			fmt.Printf("when Trans %v: %v\n", info.name(), err.Error())
			return
		}

//		defer file.Close()

		r := bufio.NewReader(file)

		loop:	for {

			lineByte, err := r.ReadBytes('\n')
			switch err {
			case nil:
				frame.pushLine(lineByte)
			case io.EOF:
				//			fmt.Println("end of file")
//				frame.pushEnd()
				break loop
			default:
				fmt.Println("there's an error when read line : " + err.Error())
				break loop
			}

		}

		if frame.GetFlushInfo() != nil {

			frame.flush()
		}

		file.Close()
	}

}

