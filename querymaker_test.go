package querymaker_test

import (
	"crypto/sha1"
	"fmt"
	"io"
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

	h := sha1.New()
	io.WriteString(h, tmpl)
	sha1Hash := fmt.Sprintf("%x", h.Sum(nil))
	if sha1Hash != "b4f2ba89cfed17295f5f613680677344c4e02ad2" { // original text for this hash can be found in `test_hash` directory
		t.Fail()
	}
}
