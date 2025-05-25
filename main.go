package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Estructura para representar un libro
type Book struct {
	Title    string
	Author   string
	Category string
}

// Lista global para almacenar los libros
var books []Book

// Agregar libro
func addBook(title, author, category string) {
	book := Book{
		Title:    title,
		Author:   author,
		Category: category,
	}
	books = append(books, book)
	fmt.Println("✅ Libro agregado correctamente.")
}

// Listar todos los libros
func listBooks() {
	if len(books) == 0 {
		fmt.Println("⚠️  No hay libros registrados.")
		return
	}
	fmt.Println("\n📚 Lista de libros:")
	for i, book := range books {
		fmt.Printf("%d. Título: %s | Autor: %s | Categoría: %s\n", i+1, book.Title, book.Author, book.Category)
	}
}

// Buscar libros por autor (ignora mayúsculas/minúsculas y espacios)
func searchByAuthor(author string) {
	found := false
	for i, book := range books {
		if strings.EqualFold(strings.TrimSpace(book.Author), strings.TrimSpace(author)) {
			if !found {
				fmt.Println("\n🔍 Libros encontrados:")
			}
			fmt.Printf("%d. Título: %s | Autor: %s | Categoría: %s\n", i+1, book.Title, book.Author, book.Category)
			found = true
		}
	}
	if !found {
		fmt.Println("❌ No se encontraron libros de ese autor.")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n📖 Sistema de Gestión de Libros Electrónicos")
		fmt.Println("1. Agregar libro")
		fmt.Println("2. Listar libros")
		fmt.Println("3. Buscar libros por autor")
		fmt.Println("4. Salir")
		fmt.Print("Selecciona una opción: ")

		var opcion int
		fmt.Scanln(&opcion)

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
			fmt.Println("👋 Saliendo del sistema. ¡Hasta luego!")
			return

		default:
			fmt.Println("⚠️ Opción no válida. Intenta de nuevo.")
		}
	}
}
