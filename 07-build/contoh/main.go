package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Book struct{
	Id int
	Title string
	Author string
	RealeaseYear string
	Pages int
}

var books []Book

var fileName string = "data.csv"

func main() {
	createFile(fileName) 
	loadDataFromCSV(fileName) // init data set

	menu() // call view menu
}

func menu(){
	fmt.Println(strings.Repeat("=", 50))
	fmt.Printf("1.\tView All Books (%d)\n", len(books))
	fmt.Println("2.\tAdd New Book")
	fmt.Println("3.\tUpdate Book")
	fmt.Println("4.\tDelete Book")
	fmt.Println("5.\tExit")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Print("Enter your choice : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	selection, err := strconv.Atoi(scanner.Text())

	fmt.Println(strings.Repeat("=", 50))
	if err != nil {
		fmt.Println("Hanya berupa angka 1 Sampai 5")
		menu()
		return
	}
	
	switch selection {
		case 1:
			viewAllBook()
			cek := continueProgrez("Data berhasil ditampilkan, apakah ingin kembali ke menu")
			if !cek {
				return // exit program
			}
			menu()
			return
		case 2:
			addNew := addNewBook()
			if addNew != nil {
				cek := continueProgrez(addNew.Error()+ ", back to menu")
				if !cek {
					return // exit program
				}
			}else{
				fmt.Println(strings.Repeat("=", 50))
				cek := continueProgrez("Book added successfully, back to menu")
				if !cek {
					return // exit program
				}
			}
			menu()
			return
		case 3:
			updateBook := editBook()
			if updateBook != nil {
				cek := continueProgrez(updateBook.Error()+ ", back to menu")
				if !cek {
					return // exit program
				}
			}
			menu()
			return
		case 4:
			delBook := delateBook()
			if delBook != nil {
				cek := continueProgrez(delBook.Error()+ ", back to menu")
				if !cek {
					return // exit program
				}
			}
			menu()
			return
		case 5:
			fmt.Println("Menu exit di pilih")
		default:
			fmt.Println("Maaf pilihan tidak ada di daftar menu")
			menu()
			return
	}
}

func continueProgrez(message string) bool {
	fmt.Print(message, " Y/n: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input := scanner.Text()
	yes := [4]string{"YES", "Y", "yes", "y"}
	no := [4]string{"NO", "N", "no", "n"}

	for i := 0; i < 4; i++ {
		if input == yes[i]{
			return true
		}else if input == no[i]{
			return false
		} 
	}
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("Maaf keyword tidak sesuai!")
	menu()
	return false
}

func addNewBook() error {
	var newBook Book

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter Book Details:")
	fmt.Print("Book Id : ")
	scanner.Scan()
	newBook.Id, _ = strconv.Atoi(scanner.Text())

	fmt.Print("Book Title : ")
	scanner.Scan()
	newBook.Title = scanner.Text()

	fmt.Print("Book Author : ")
	scanner.Scan()
	newBook.Author = scanner.Text()

	fmt.Print("Book RealeaseYear : ")
	scanner.Scan()
	newBook.RealeaseYear = scanner.Text()

	fmt.Print("Book Pages : ")
	scanner.Scan()
	newBook.Pages, _ = strconv.Atoi(scanner.Text())

	_, err := findBookById(newBook.Id)
	if err != nil {
		fmt.Println(strings.Repeat("=", 50))
		continues := continueProgrez("Are you sure want to add this book")
		if continues{
			books = append(books, newBook)
		} else {
			menu()
		}
		
	}else{
		fmt.Println(strings.Repeat("=", 50))
		return fmt.Errorf("Book with id %d already exist", newBook.Id)
	}

	err = saveDataToCSV(fileName) // update file
	if err != nil {
		return err
	}

	return nil
}

func editBook() error {
	var newBook Book

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter Book Id to Update : ")
	scanner.Scan()
	newBook.Id, _ = strconv.Atoi(scanner.Text())
	fmt.Println(strings.Repeat("=", 50))

	fmt.Print("Book Title : ")
	scanner.Scan()
	newBook.Title = scanner.Text()

	fmt.Print("Book Author : ")
	scanner.Scan()
	newBook.Author = scanner.Text()

	fmt.Print("Book RealeaseYear : ")
	scanner.Scan()
	newBook.RealeaseYear = scanner.Text()

	fmt.Print("Book Pages : ")
	scanner.Scan()
	newBook.Pages, _ = strconv.Atoi(scanner.Text())

	buku, err := findBookById(newBook.Id)
	if err == nil {
		fmt.Println(strings.Repeat("=", 50))
		continues := continueProgrez("Are you sure want to update this book")
		if continues{
			// empty Handling
			if newBook.Title == "" {
				newBook.Title = buku.Title
			}else if newBook.Author == "" {
				newBook.Author = buku.Author
			}else if newBook.RealeaseYear == "" {
				newBook.RealeaseYear = buku.RealeaseYear
			}else if newBook.Pages == 0 {
				newBook.Pages = buku.Pages
			}

			// update Slice
			for i := 0; i < len(books); i++ {
				if books[i].Id == buku.Id{
					books[i] = newBook
				}
			}
		} else {
			menu()
		}
		
	}else{
		fmt.Println(strings.Repeat("=", 50))
		return fmt.Errorf("Book Id is nil or empty")
	}

	err = saveDataToCSV(fileName) // update file
	if err != nil {
		return err
	}

	return nil
}

func delateBook() error {
	var newBook Book

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Delate book by Id : ")
	scanner.Scan()
	newBook.Id, _ = strconv.Atoi(scanner.Text())

	buku, err := findBookById(newBook.Id)
	if err == nil {
		fmt.Println(strings.Repeat("=", 50))
		continues := continueProgrez("Are you sure want to delete "+buku.Title+"'s")
		if continues{
			for i := len(books) - 1; i >= 0; i-- {
				if books[i].Id == buku.Id {
					books = append(books[:i], books[i+1:]...) // Remove Slice By index
				}
			}
		} else {
			menu()
		}
		
	}else{
		fmt.Println(strings.Repeat("=", 50))
		return fmt.Errorf("Book Id is nil or empty")
	}

	err = saveDataToCSV(fileName) // update file
	if err != nil {
		return err
	}

	return nil
}

func saveDataToCSV(fileName string)error {
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("Error opening csv file: %w", err)
	}
	defer file.Close()

	write := bufio.NewWriter(file)

	for _, book := range books {
		row := strconv.Itoa(book.Id) + "," + book.Title + "," + book.Author+ "," +book.RealeaseYear + "," + strconv.Itoa(book.Pages) + "\n"
		write.WriteString(row)
		write.Flush()
	}

	return nil
}

func loadDataFromCSV(fileName string)error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("error opening csv file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		record := strings.Split(scanner.Text(), ",")
		//fmt.Printf("format id: %T \n",record[0]) // id berupa string
		
		id, _ := strconv.Atoi(record[0])
		pages, _ := strconv.Atoi(record[4])

		book := Book {
			Id: id,
			Pages: pages,
			Title: record[1],
			Author: record[2],
			RealeaseYear: record[3],
		}

		books = append(books, book)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error opening scv file: %w", err)
	}

	return nil
}

func viewAllBook()error {
	if len(books) == 0 {
		return fmt.Errorf("no books available")
	}

	for i, book := range books {
		fmt.Println("Book - ", i+1)
		fmt.Println("Book Id ", book.Id)
		fmt.Println("Book Title ", book.Title)
		fmt.Println("Book Author ", book.Author)
		fmt.Println("Book Release Year ", book.RealeaseYear)
		fmt.Println("Book Pages ", book.Pages)

		fmt.Println(strings.Repeat("=", 50))
	}

	return nil
}

func findBookById(id int) (Book, error){
	for _, book := range books {
		if book.Id == id {
			return book, nil
		}
	}

	return Book{}, fmt.Errorf("id: %d not found", id)
}

func createFile(fileName string){
	var fileInfo os.FileInfo
	fileInfo, err:= os.Stat(fileName)

	if os.IsNotExist(err) {
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Println("file", fileName, "berhasil dibuat!")
		
	}else{
		fmt.Println("file", fileInfo.Name(), fileInfo.ModTime())
	}
}