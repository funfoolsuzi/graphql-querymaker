package querymaker_test

import (
	"fmt"
	"testing"

	qm "github.com/nbcnews/graphql-querymaker"
)

type dummyQuery struct {
	Dummies []dummy `json:"dummy"`
}

type dummy struct {
	Name              string `json:"name"`
	Height            int    `graphqlvar:"hhh,abc"`
	NotImportantField string `json:"-"`
	Child             *subDummy
}

type subDummy struct {
	Location string `graphqlvar:"iii,qwe"`
}

func TestMakeQuery(t *testing.T) {
	d := &dummyQuery{}
	tmpl, _ := qm.MakeQuery(d)

	fmt.Println(tmpl)
}
