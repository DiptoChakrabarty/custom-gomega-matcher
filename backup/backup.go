package random

import (
	"fmt"
	"strings"

	"github.com/olekukonko/tablewriter"
)

func main() {
	fmt.Println("This is New Table")

	tableString := &strings.Builder{}
	data := [][]string{
		[]string{"1/1/2014", "Domain name", "2233", "$10.98"},
		[]string{"1/1/2014", "January Hosting", "2233", "$54.95"},
		[]string{"1/4/2014", "February Hosting", "2233", "$51.00"},
		[]string{"1/4/2014", "February Extra Bandwidth", "2233", "$30.00"},
	}

	table := tablewriter.NewWriter(tableString)
	table.SetHeader([]string{"Date", "Description", "CV2", "Amount"})
	table.SetFooter([]string{"", "", "Total", "$146.93"}) // Add Footer
	table.SetBorder(false)                                // Set Border to false
	table.AppendBulk(data)

	table.Render()
	fmt.Println(tableString.String())
	fmt.Println(len(tableString.String()))

	random := []string{
		"Object/root",
		"├─Object/child1",
		"│ │           ├─C1.1", // first condition child gets pipes, children pipe and ├─
		"│ │           └─C1.2", // last condition child gets pipes, children pipe and └─
		"│ └─Object/child1.1",
		"└─Object/child2",
		"  │           ├─C2.1", // first condition child gets spaces, children pipe and ├─
		"  │           └─C2.2", // last condition child gets spaces, children pipe and └─
		"  └─Object/child2.1",
	}

	for i := range random {
		for j := range random[i] {
			fmt.Println(string(random[i][j]))
		}
		//fmt.Println("\n")
	}
}
