// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	// Tenta criar o arquivo dentro do diretório ./temp que ainda não existe
// 	_, err := os.Create("./temp/file.txt")
// 	if err != nil {
// 		panic("Erro ao criar o arquivo ./temp/file.txt: " + err.Error())
// 	}

// 	fmt.Println("Arquivo criado com sucesso.")
// }

// ######################################################################
// package main

// import "fmt"

// func dividir(num, div float64) float64 {
// 	if div == 0 {
// 		panic("divisão por zero não permitida")
// 	}
// 	return num / div
// }

// func main() {
// 	fmt.Println("Resultado:", dividir(10, 2)) // Saída: Resultado: 5
// 	fmt.Println("Resultado:", dividir(10, 3)) // Saída: Resultado: 3.3333333333333335
// 	fmt.Println("Resultado:", dividir(10, 0)) // Isto irá disparar o panic
// }

//#######################################################################

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func conectarBancoDeDados() *sql.DB {
	// Conexão com o MySQL
	connStr := "root:123456@/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		panic("falha ao conectar no banco de dados: " + err.Error())
	}

	// Verifica se a conexão está ativa
	err = db.Ping()
	if err != nil {
		panic("não foi possível estabelecer uma conexão ativa: " + err.Error())
	}

	return db
}

func main() {
	db := conectarBancoDeDados()
	defer db.Close()

	fmt.Println("Conexão com o banco de dados estabelecida com sucesso.")
}
