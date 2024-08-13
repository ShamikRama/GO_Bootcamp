package main

import (
	c "Data_Api/ex00/client"
	"fmt"
)

func main() {
	indexName := "places"

	es := c.GetConnectionClient()

	err := c.CreateIndex(indexName, es)
	if err != nil {
		fmt.Print(err)
	}

	err1 := c.Mapping(indexName, es)
	if err1 != nil {
		fmt.Print(err)
	}

	err2 := c.FillESDATA(indexName, es, "Data_Api/ex00/cmd/data.csv")
	if err2 != nil {
		fmt.Print(err)
	}

}
