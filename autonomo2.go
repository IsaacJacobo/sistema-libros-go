package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Interfaz para representar un comportamiento imprimible
type Printable interface {
	PrintInfo()
}

// Estructura para representar un libro (con encapsulación)
type Book struct {
	title    string
	author   string
	category string
}

// Métodos Setter (Encapsulación)
func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b *Book) SetCategory(category string) {
	b.category = category
}

// Métodos Getter
func (b Book) GetTitle() string {
	return b.title
}

func (b Book) GetAuthor() string {
	return b.author
}

func (b Book) GetCategory() string {
	return b.category
}

// Implementación de la interfaz Printable
func (b Book) PrintInfo() {
	fmt.Printf("Título: %s | Autor: %s | Categoría: %s\n", b.title, b.author, b.category)
}

// Lista global de libros
var books []Book

// Función para agregar un libro con validación
func addBook(title, author, category string) {
	// Validación básica (manejo de errores)
	if title == "" || author == "" || category == "" {
		fmt.Println("Todos los campos son obligatorios. Intenta de nuevo.")
		return
	}

	// Crear libro usando setters (encapsulación)
	var book Book
	book.SetTitle(title)
	book.SetAuthor(author)
	book.SetCategory(category)

	books = append(books, book)
	fmt.Println("Libro agregado correctamente.")
}

// Función para listar libros
func listBooks() {
	if len(books) == 0 {
		fmt.Println("No hay libros registrados.")
		return
	}
	fmt.Println("\nLista de libros:")
	for i, book := range books {
		fmt.Printf("%d. ", i+1)
		book.PrintInfo() // uso de interfaz
	}
}

// Buscar libros por autor con comparación flexible
func searchByAuthor(author string) {
	found := false
	for i, book := range books {
		if strings.EqualFold(strings.TrimSpace(book.GetAuthor()), strings.TrimSpace(author)) {
			if !found {
				fmt.Println("\nLibros encontrados:")
			}
			fmt.Printf("%d. ", i+1)
			book.PrintInfo()
			found = true
		}
	}
	if !found {
		fmt.Println(" No se encontraron libros de ese autor.")
	}
}

// Función principal con manejo de errores
func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nSistema de Gestión de Libros Electrónicos")
		fmt.Println("1. Agregar libro")
		fmt.Println("2. Listar libros")
		fmt.Println("3. Buscar libros por autor")
		fmt.Println("4. Salir")
		fmt.Print("Selecciona una opción: ")

		input, _ := reader.ReadString('\n')
		opcion, err := strconv.Atoi(strings.TrimSpace(input))

		// Manejo de error al convertir la opción
		if err != nil {
			fmt.Println("Entrada no válida. Debe ser un número. Intenta de nuevo.")
			continue
		}

		switch opcion {
		case 1:
			fmt.Print("Título: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)

			fmt.Print("Autor: ")
			author, _ := reader.ReadString('\n')
			author = strings.TrimSpace(author)

			fmt.Print("Categoría: ")
			category, _ := reader.ReadString('\n')
			category = strings.TrimSpace(category)

			addBook(title, author, category)

		case 2:
			listBooks()

		case 3:
			fmt.Print("Autor a buscar: ")
			searchAuthor, _ := reader.ReadString('\n')
			searchAuthor = strings.TrimSpace(searchAuthor)

			searchByAuthor(searchAuthor)

		case 4:
			fmt.Println("Saliendo del sistema. ¡Hasta luego!")
			return

		default:
			fmt.Println("Opción no válida. Intenta de nuevo.")
		}
	}
}
