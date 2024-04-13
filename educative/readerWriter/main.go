package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var nrchar, nrwords, nrline int

func Counters(s string) {
	nrchar += len(s) - 2
	nrwords += strings.Count(s, " ") + 1
	nrline++
}

func openAndRead(pathToFile string) {
	inputFile, inputError := os.Open(pathToFile)
	if inputError != nil {
		fmt.Print("Error to opern file")
	}
	defer inputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			fmt.Print("End of file")
		}
		fmt.Printf("The input was :%s", inputString)
	}

}

func readColumns(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var col1, col2, col3 []string
	for {
		var v1, v2, v3 string
		_, err := fmt.Fscanln(file, &v1, &v2, &v3)
		if err != nil {
			break
		}
		col1 = append(col1, v1)
		col2 = append(col2, v2)
		col3 = append(col3, v3)
	}
	fmt.Println(col1)
	fmt.Println(col2)
	fmt.Println(col3)
}

func readZipFile(path string) {
	var r *bufio.Reader
	fi, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v, Cant open %s: error: %s\n", os.Args[0], path, err)
		os.Exit(1)
	}
	fz, err := gzip.NewReader(fi)
	if err != nil {
		r = bufio.NewReader(fi)
	} else {
		r = bufio.NewReader(fz)
	}
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Done reading file")
			os.Exit(0)
		}
		fmt.Println(line)
	}
}

func writeToFile() {
	outputFile, outputError := os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Print("error")
		os.Exit(1)
	}
	defer outputFile.Close()
	outputWriter := bufio.NewWriter(outputFile)
	outputString := "Hello, this is Emi\n"
	for i := 0; i < 5; i++ {
		outputWriter.WriteString(outputString)
	}
	outputWriter.Flush()
}

type Book struct {
	title    string
	price    float64
	quantity int
}

func readCVSFile() {
	bks := make([]Book, 1)
	file, err := os.Open("products.txt")
	if err != nil {
		log.Fatalf("Error %s opening file products.txt: ", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		// read one line from the file:
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		// remove \r and \n so 2 in Windows, in Linux only \n, so 1:
		line = string(line[:len(line)-2])
		//fmt.Printf("The input was: -%s-", line)

		strSl := strings.Split(line, ";")
		book := new(Book)
		book.title = strSl[0]
		book.price, err = strconv.ParseFloat(strSl[1], 32)
		if err != nil {
			fmt.Printf("Error in file: %v", err)
		}
		//fmt.Printf("The quan was:-%s-", strSl[2])
		book.quantity, err = strconv.Atoi(strSl[2])
		if err != nil {
			fmt.Printf("Error in file: %v", err)
		}
		if bks[0].title == "" {
			bks[0] = *book
		} else {
			bks = append(bks, *book)
		}
	}
	fmt.Println("We have read the following books from the file: ")
	for _, bk := range bks {
		fmt.Println(bk)
	}
}

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func load(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

func cat(r *bufio.Reader) {
	for {
		r, err := r.ReadBytes('\n')
		if err != io.EOF {
			break
		}
		fmt.Fprintf(os.Stdout, "%s", r)
	}
	return
}

func catSlice(f *os.File) {
	const NBUF = 512
	var buf [NBUF]byte
	for {
		switch nr, err := f.Read(buf[:]); true {
		case nr < 0:
			fmt.Fprintf(os.Stderr, "cat: error reading: %s\n", err.Error())
			os.Exit(1)
		case nr == 0:
			return
		case nr > 0:
			if nw, ew := os.Stdout.Write(buf[0:nr]); nw != nr {
				fmt.Fprintf(os.Stderr, "cat: error writing: %s\n", ew)
			}

		}
	}
}

func useScan() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please put some input: ")
	scanner.Split(bufio.ScanWords)
	for {
		scanner.Scan()
		fmt.Printf("Token ->%s<-\n", scanner.Text())
	}
}
func main() {
	// readCVSFile()

	// p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	// p1.save()
	// p2, _ := load("TestPage")
	// fmt.Println(string(p2.Body))

	// CopyFile("TestPage2.txt", "TestPage.txt")
	// fmt.Println("Copy Done")

	// flag.Parse()
	// if flag.NArg() == 0 {
	// 	catSlice(os.Stdin)
	// }
	// for i := 0; i < flag.NArg(); i++ {
	// 	f, err := os.Open(flag.Arg(i))
	// 	if f == nil {
	// 		fmt.Fprintf(os.Stderr, "cat: can't open %s: error %s\n", flag.Arg(i), err)
	// 		os.Exit(1)
	// 	}
	// 	catSlice(f)
	// 	f.Close()
	// }

	useScan()
}
