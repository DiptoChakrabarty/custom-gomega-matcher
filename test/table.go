package test

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

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
	// This writes the table to a file which can be viewed and matched to string
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
	//fmt.Println(tableString.String())
	fmt.Println(tbl)

	random := []string{
		"+----------+-------------+-------+---------+",
		"|   DATE   | DESCRIPTION |  CV2  | AMOUNT  |",
		"+----------+-------------+-------+---------+",
		"| 1/1/2014 | Domain name |  2233 | $10.98  |",
		"+----------+-------------+-------+---------+",
		"|                          TOTAL | $146.93 |",
		"+----------+-------------+-------+---------+",
	}

	flag := 0
	tablelength := 0

	s := strings.Split(tbl, "\n")

	fmt.Println(len(random), len(random[0]))
	for i := range random {
		fmt.Println(random[i], s[i], len(random[i]), len(s[i]))
		if !reflect.DeepEqual(random[i], s[i]) {
			fmt.Println("THIS IS WRONG")
			flag += 1
		}
		//fmt.Println("\n")
	}
	fmt.Println(tablelength)
	fmt.Println("Flag status is", flag)

}
