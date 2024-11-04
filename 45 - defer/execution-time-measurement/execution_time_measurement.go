package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func realizarConsulta() error {
	// Abre conexão com o banco
	db, err := sql.Open("mysql", "root:sua_senha@/defer_example")
	if err != nil {
		return fmt.Errorf("erro ao conectar ao banco: %v", err)
	}
	defer db.Close()

	// Inicia uma transação
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("erro ao iniciar transação: %v", err)
	}

	// Garante que a transação será finalizada
	defer func() {
		if err != nil {
			fmt.Println("Revertendo transação...")
			tx.Rollback()
			return
		}
		fmt.Println("Commitando transação...")
		tx.Commit()
	}()

	// Insere um usuário
	_, err = tx.Exec(`
        INSERT INTO usuarios (nome, email) 
        VALUES (?, ?)
    `, "João", "joao@exemplo.com")
	if err != nil {
		return fmt.Errorf("erro ao inserir usuário: %v", err)
	}

	// Tenta uma operação que pode falhar
	_, err = tx.Exec(`
        INSERT INTO usuarios (nome, email) 
        VALUES (?, ?)
    `, "Maria", "maria@exemplo.com")
	if err != nil {
		return fmt.Errorf("erro ao inserir segundo usuário: %v", err)
	}

	// Lista os usuários inseridos
	rows, err := tx.Query("SELECT id, nome, email FROM usuarios")
	if err != nil {
		return fmt.Errorf("erro ao consultar usuários: %v", err)
	}
	defer rows.Close()

	fmt.Println("\nUsuários inseridos:")
	for rows.Next() {
		var id int
		var nome, email string
		err := rows.Scan(&id, &nome, &email)
		if err != nil {
			return fmt.Errorf("erro ao ler usuário: %v", err)
		}
		fmt.Printf("ID: %d, Nome: %s, Email: %s\n", id, nome, email)
	}

	return nil
}

func main() {
	if err := realizarConsulta(); err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}
	fmt.Println("Operação concluída com sucesso!")
}
