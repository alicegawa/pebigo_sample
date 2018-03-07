package main
import (
	"fmt"
	"log"

	"github.com/alicegawa/pebigo"
)

func main() {
	opts := pebigo.NewOptions()
	opts.SetCache(pebigo.NewLRUCache(3<<30))
	opts.SetCreateIfMissing(true)
	db, err := pebigo.Open("./database/pebblesdb/sample", opts)
	if err != nil {
		log.Fatal(err)
	}

	wo := pebigo.NewWriteOptions()
	var key, value string
	tmpKey := "sampleKey-%d"
	tmpValue := "sampleValue-%d"
	for i := 0; i < 256; i++ {
		key = fmt.Sprintf(tmpKey, i)
		value = fmt.Sprintf(tmpValue, i)
		err = db.Put(wo, []byte(key), []byte(value))
		if err != nil {
			log.Fatal(err)
		}
	}
	wo.Close()

	ro := pebigo.NewReadOptions()
	var getValue []byte
	for i := 0; i < 256; i++ {
		key = fmt.Sprintf(tmpKey, i)
		getValue, err = db.Get(ro, []byte(key))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(getValue))
	}
	ro.Close()

	db.Close()
}
