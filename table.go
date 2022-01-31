package main

import (
	//"bufio"
	"bytes"
	"fmt"

	"github.com/olekukonko/tablewriter"
)

func main() {
	fmt.Println("This is New Table")
	/*f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}*/
	//stdout := bufio.NewWriter(os.Stdout)
	var output bytes.Buffer

	//tableString := &strings.Builder{}
	data := [][]string{
		[]string{"1/1/2014", "Domain name", "2233", "$10.98"},
	}

	table := tablewriter.NewWriter(&output)
	table.SetHeader([]string{"Date", "Description", "CV2", "Amount"})
	table.SetFooter([]string{"", "", "Total", "$146.93"}) // Add Footer
	table.SetBorder(true)                                 // Set Border to false
	table.AppendBulk(data)

	table.Render()
	tbl := output.String()
	/*fmt.Fprintln(f, tbl)
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}*/
	//fmt.Println(tbl)
	/*
		fmt.Println(len(tbl))
		for i, _ := range tbl {
			fmt.Println(string(tbl[i]), i)
		}*/
	//fmt.Println(stdout.Size())
	// https://stackoverflow.com/questions/26804642/how-to-test-a-functions-output-stdout-stderr-in-unit-tests
	// https://stackoverflow.com/questions/10473800/in-go-how-do-i-capture-stdout-of-a-function-into-a-string

	//fmt.Println(tableString.String())
	fmt.Println(tbl)

	random := []string{
		"+----------+-------------+-------+---------+\n",
		"|   DATE   | DESCRIPTION |  CV2  | AMOUNT  |\n",
		"+----------+-------------+-------+---------+\n",
		"| 1/1/2014 | Domain name |  2233 | $10.98  |\n",
		"+----------+-------------+-------+---------+\n",
		"|                          TOTAL | $146.93 |\n",
		"+----------+-------------+-------+---------+\n",
	}

	flag := 0
	tablelength := 0

	/*for i := range random {
		fmt.Println(i)
		fmt.Println("new")
		for j := range random[i] {
			fmt.Println(i, j, string(random[i][j]))
		}
	}*/

	fmt.Println(len(random), len(random[0]))
	for i := range random {
		for j := range random[i] {
			if string(random[i][j]) != string(tbl[tablelength]) {
				flag += 1
				fmt.Println("Invalid Characters")
				fmt.Println(i, j, string(random[i][j]), string(tbl[tablelength]))
			}
			tablelength += 1
		}
		//fmt.Println("\n")
	}
	fmt.Println(tablelength)
	fmt.Println("Flag status is", flag)

}
