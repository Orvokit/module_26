/*
Написать программу аналог cat.
Программа должна получать на вход имена двух файлов, необходимо  конкатенировать их содержимое, используя strings.Join.
*/

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	//------------------ работа если у нас 1 файл -------------------------

	if len(os.Args) > 1 { //prog и 1 файл

		//открытие файла 1
		file1, err1 := os.Open(os.Args[1])
		if err1 != nil {
			fmt.Println("Ошибка при открытии файла 1", err1)
		}
		defer file1.Close()

		data1 := make([]byte, 64)

		//чтение из файла 1
		for {
			_, err := file1.Read(data1)
			if err == io.EOF {
				break
			}
		}

		data1String := string(data1)
		if len(os.Args) == 2 {
			fmt.Println(data1String)
		}

		//------------------ работа если у нас 2 файла -------------------------

		if len(os.Args) > 2 { //prog и 2 файла
			file2, err2 := os.Open(os.Args[2])
			if err2 != nil {
				fmt.Println("Ошибка при открытии файла 2", err2)
			}
			defer file2.Close()

			data2 := make([]byte, 64)

			//чтение из файла 2
			_, err := file2.Read(data2)
			if err == io.EOF {
				file2.Close()
			}

			data2String := string(data2)
			if len(os.Args) == 3 {
				fmt.Println(data1String + " " + data2String)
			}

			//------------------ работа если у нас 3 файлами -------------------------

			if len(os.Args) > 3 {
				file3, err3 := os.Create(os.Args[3])
				if err3 != nil {
					fmt.Println("Ошибка при создании файла 3", err3)
				}
				defer file3.Close()

				file3.WriteString(data1String + "\n" + data2String)
				fmt.Println("Результат записан в файл", os.Args[3])
			}

		}
	}
}

//запуск: go run module_26/homework.go module_26/first.txt module_26/second.txt module_26/result.txt
