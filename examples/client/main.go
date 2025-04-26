package main

import (
	"fmt"
	"log"
	"victorclient/client"
)

func main() {

	vClient := client.NewClient(&client.ClientOptions{
		Host: "localhost",
		Port: "8080",
	})

	indexParams := client.CreateIndexCommandInput{
		IndexType: 0,
		Method:    0,
		Dims:      5,
		IndexName: "indice1",
	}

	destroyIndexParams := client.DestroyIndexCommandInput{
		IndexType: 0,
		Method:    0,
		Dims:      5,
		IndexName: "indice1",
	}

	v1 := client.InsertVectorCommandInput{
		IndexName: "indice1",
		ID:        1,
		Vector:    []float32{0.1, 0.2, 0.3, 0.4, 0.5},
	}

	v2 := client.InsertVectorCommandInput{
		IndexName: "indice1",
		ID:        2,
		Vector:    []float32{0.6, 0.7, 0.8, 0.9, 1.0},
	}
	v3 := client.InsertVectorCommandInput{
		IndexName: "indice1",
		ID:        3,
		Vector:    []float32{1.1, 1.2, 1.3, 1.4, 1.5},
	}

	searchVector := client.SearchVectorCommandInput{
		IndexName: "indice1",
		TopK:      3,
		Vector:    []float32{1.5, 0.2, 2, -0.4, 0.9},
	}

	deleteVector := client.DeleteVectorCommandInput{
		IndexName: "indice1",
		VectorID:  1,
	}
	//---------------------------------------------------------------------
	indexResult, err := vClient.CreateIndex(&indexParams)
	if err != nil {
		log.Fatalf("Index creation error: %+v", err)
	}

	fmt.Printf("%s\n%s\n%+v\n", indexResult.Status, indexResult.Message, indexResult.Results)

	destroyResult, err := vClient.DestroyIndex(&destroyIndexParams);
	if err != nil {
		log.Fatalf("Index destroy error: %+v", err);
	}
	fmt.Printf("%s\n%s\n%+v\n", destroyResult.Status, destroyResult.Message, destroyResult.Results)


	insertResult1, err := vClient.InsertVector(&v1)
	if err != nil {
		log.Fatalf("Vector insertion error: %+v", err)
	}
	fmt.Printf("%s\n%s\n%+v\n", insertResult1.Status, insertResult1.Message, insertResult1.Results)

	insertResult2, err := vClient.InsertVector(&v2)
	if err != nil {
		log.Fatalf("Insertion vector error: %+v", err)
	}
	fmt.Printf("%s\n%s\n%+v\n", insertResult2.Status, insertResult2.Message, insertResult2.Results)

	insertResult3, err := vClient.InsertVector(&v3)
	if err != nil {
		log.Fatalf("Insertion vector Error: %+v", err)
	}
	fmt.Printf("%s\n%s\n%+v\n", insertResult3.Status, insertResult3.Message, insertResult3.Results)

	searchResult, err := vClient.SearchVector(&searchVector)
	if err != nil {
		log.Fatalf("Search error: %+v", err)
	}
	fmt.Printf("%s\n%s\n%+v\n", searchResult.Status, searchResult.Message, searchResult.Results)

	deleteResult, err := vClient.DeleteVector(&deleteVector)
	if err != nil {
		log.Fatalf("Delete vector error: %+v", err)
	}
	fmt.Printf("%s\n%s\n%+v\n", deleteResult.Status, deleteResult.Message, deleteResult.Results)

}
