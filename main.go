package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	file := "/Users/zhangsheng/Downloads/pailiao/Pailiao.exe"
	fileNew := "/Users/zhangsheng/Downloads/pailiao/PL.exe"
	//	file := "/Users/zhangsheng/Downloads/pailiao/PL.exe"
	//	fileNew := "/Users/zhangsheng/Downloads/pailiao/PLnew.exe"

	fin, err := os.OpenFile(fileNew, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("create Error: ", err.Error())
		return
	} else {
		defer fin.Close()
	}

	f, ok := os.Open(file)
	if ok != nil {
		fmt.Println("Open Error: ", ok.Error())
	} else {
		defer f.Close()
		decoder := mahonia.NewDecoder("gb18030")
		if decoder == nil {
			fmt.Println("编码不存在!")
			return
		}
		encoder := mahonia.NewEncoder("gb18030")
		if encoder == nil {
			fmt.Println("编码不存在!")
			return
		}

		data := make([]byte, 4096)
		for {
			_, end := f.Read(data)
			if end == io.EOF {
				fmt.Println("end!")
				break
			} else if end == nil {
				oldstr := encoder.ConvertString("瑞丽超级自动排料系统")
				//newstr := encoder.ConvertString("                    ")
				newstr := encoder.ConvertString("深圳新群力-排版系统 ")

				var newdata []byte
				if bytes.Contains(data, []byte(oldstr)) {
					fmt.Println("Oh,Oh,Yeah!!")
					newdata = bytes.Replace(data, []byte(oldstr), []byte(newstr), -1)
					fmt.Println(len(oldstr), len(newstr))
					fmt.Println(len(data), len(newdata))
				} else {
					newdata = data
				}

				if _, err := fin.Write(newdata); err != nil {
					fmt.Println("Write Error:", err.Error())
					break
				}
			} else {
				fmt.Println("Read Error: ", ok.Error())
			}

		}
	}
}
