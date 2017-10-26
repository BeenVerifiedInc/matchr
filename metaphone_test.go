package matchr

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestDoubleMetaphone(t *testing.T) {
	testDoubleMetaphoneHelper(t, "double_metaphone_corpus.txt.gz", func(tester string) (string, string) {
		return DoubleMetaphone(tester)
	})
}

func TestDoubleMetaphoneWithMaxLength(t *testing.T) {
	testDoubleMetaphoneHelper(t, "double_metaphone_max_length_32_corpus.txt.gz", func(tester string) (string, string) {
		return DoubleMetaphone(tester, 32)
	})
}

func testDoubleMetaphoneHelper(t *testing.T, fileName string, testFn func(string) (string, string)) {
	// load gzipped corpus
	f, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Sprintf("Error opening file %s! Exiting.", fileName))
	}
	defer f.Close()

	g, err := gzip.NewReader(f)
	if err != nil {
		panic(fmt.Sprintf("Error with supposedly gzipped file %s! Exiting.", fileName))
	}

	r := bufio.NewReader(g)

	line, err := r.ReadString('\n')
	for err == nil {
		line = strings.TrimRight(line, "\n")
		v := strings.Split(line, "|")

		metaphone, alternate := testFn(v[0])
		if metaphone != v[1] || alternate != v[2] {
			t.Errorf("DoubleMetaphone('%s') = (%v, %v), want (%v, %v)", v[0], metaphone, alternate, v[1], v[2])
			t.FailNow()
		}

		line, err = r.ReadString('\n')
	}
}
