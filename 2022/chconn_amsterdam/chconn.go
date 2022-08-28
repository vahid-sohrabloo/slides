package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/vahid-sohrabloo/chconn/v2/chpool"
	"github.com/vahid-sohrabloo/chconn/v2/column"
)

func Insert() {
	// CONNECT START OMIT
	conn, err := chpool.New(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	// to check if the connection is alive
	err = conn.Ping(context.Background())
	if err != nil {
		panic(err)
	}

	err = conn.Exec(context.Background(), `DROP TABLE IF EXISTS example_table`)
	if err != nil {
		panic(err)
	}

	err = conn.Exec(context.Background(), `CREATE TABLE  example_table (
		uint64 UInt64,
		uint64_nullable Nullable(UInt64)
	) Engine=Memory`)

	// CONNECT END OMIT
	if err != nil {
		panic(err)
	}

	// COLUMN START OMIT
	col1 := column.New[uint64]()
	col2 := column.New[uint64]().Nullable()
	rows := 1_000_0000 // One hundred million rows- insert in 10 times
	numInsert := 10
	col1.SetWriteBufferSize(rows)
	col2.SetWriteBufferSize(rows)
	// COLUMN END OMIT
	// INSERT START OMIT
	startInsert := time.Now()
	for i := 0; i < numInsert; i++ {
		for y := 0; y < rows; y++ {
			col1.Append(uint64(i))
			if i%2 == 0 {
				col2.Append(uint64(i))
			} else {
				col2.AppendNil()
			}
		}

		ctxInsert, cancelInsert := context.WithTimeout(context.Background(), time.Second*30)
		// insert data
		err = conn.Insert(
			ctxInsert, "INSERT INTO example_table (uint64,uint64_nullable) VALUES",
			col1, col2)
		if err != nil {
			cancelInsert()
			panic(err)
		}
		cancelInsert()
	}
	fmt.Println("inserted 10M rows in ", time.Since(startInsert))
	// INSERT END OMIT

	// select data
	// COLUMNSEL START OMIT
	col1Read := column.New[uint64]()
	col2Read := column.New[uint64]().Nullable()
	// COLUMNSEL END OMIT

	ctxSelect, cancelSelect := context.WithTimeout(context.Background(), time.Second*30)
	defer cancelSelect()

	startSelect := time.Now()
	// SELECT START OMIT
	selectStmt, err := conn.Select(ctxSelect,
		"SELECT uint64,uint64_nullable FROM  example_table",
		col1Read, col2Read)
	if err != nil {
		panic(err)
	}

	defer selectStmt.Close()

	var col1Data []uint64
	var col2DataNil []bool
	var col2Data []uint64
	for selectStmt.Next() {
		col1Data = col1Read.Read(col1Data[:0])
		col2DataNil = col2Read.ReadNil(col2DataNil[:0])
		col2Data = col2Read.Read(col2Data[:0])
	}

	if selectStmt.Err() != nil {
		panic(selectStmt.Err())
	}
	fmt.Println("selected 10M rows in ", time.Since(startSelect))
	// SELECT END OMIT

	// RESULT START OMIT
	// inserted 10M rows in  1.206666188s
	// selected 10M rows in  880.505004ms
	// RESULT END OMIT

}
