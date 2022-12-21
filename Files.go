package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

/**
В базе данных планово-экономического отдела хранятся сведения о выпуске изделий и затратах (по статьям калькуляции) на каждое изделие.
Шифр статьи затрат Изделие Затраты на единицу (условных единиц) План(шт.)
Структура входного файла in.txt
2 Принтер 200 5
1 Монитор 100 5
2 Картридж 10 20
...
Определить суммарную себестоимость по каждой статье затрат, упорядочив по номеру статьи
Структура выходного файла out.txt
Шифр статьи затрат Себестоимость (условных единиц)
1 500
2 1200
*/
func main() {
	resMap := map[int]int{}
	in, err := os.Open("in")

	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()

	wr, err := os.Create("write")

	if err != nil {
		log.Fatal(err)
	}
	defer wr.Close()

	scanner := bufio.NewScanner(in)

	for scanner.Scan() {
		arrOfStr := strings.Split(scanner.Text(), " ")

		numberString, err := strconv.Atoi(arrOfStr[0])
		if err != nil {
			log.Fatal(err)
		}

		costs, err := strconv.Atoi(arrOfStr[2])
		if err != nil {
			log.Fatal(err)
		}

		units, err := strconv.Atoi(arrOfStr[3])
		if err != nil {
			log.Fatal(err)
		}

		resMap[numberString] += costs * units

	}
	for i := 1; i < len(resMap)+i; i++ {
		sum, ok := resMap[i]
		if !ok {
			continue
		}

		_, err := wr.WriteString(strconv.Itoa(i) + " " + strconv.Itoa(sum) + "\n")

		if err != nil {
			panic(err)
		}
	}
}
