package main

import (
	c "Data_Api/ex00/client"
	"fmt"
)

func main() {
	es := c.GetConnectionClient()

	err := c.CreateIndex("places", es)
	if err != nil {
		fmt.Print(err)
	}

	err1 := c.Mapping("places", es)
	if err1 != nil {
		fmt.Print(err)
	}

	err2 := c.FillData("places", es, "data.csv")
	if err2 != nil {
		fmt.Print(err)
	}

}
