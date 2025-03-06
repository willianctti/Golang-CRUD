package database

import (
	"fmt"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"go-crud/models"
)

// DB eh uma var global que guarda a conexão com o banco
// outros arquivos podem usar database.DB para acessar o banco
var DB *gorm.DB

func Connect() {
	fmt.Println("Tentando conectar ao postgres...")
	
	// Primeiro conecta ao postgres
	dsn := "host=127.0.0.1 user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar no postgres: ", err)
	}

	// Verifica se o banco go_crud existe
	var count int64
	db.Raw("SELECT COUNT(*) FROM pg_database WHERE datname = 'go_crud'").Scan(&count)
	
	// Se não existe, cria o banco
	if count == 0 {
		fmt.Println("Criando banco go_crud...")
		err = db.Exec("CREATE DATABASE go_crud").Error
		if err != nil {
			// erro doidao do go != nil
			log.Fatal("Erro ao criar banco: ", err)
		}
		fmt.Println("Banco go_crud criado com sucesso!")
	} else {
		fmt.Println("Banco go_crud já existe!")
	}

	// Conecta ao banco go_crud
	dsn = "host=127.0.0.1 user=postgres password=postgres dbname=go_crud port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar no go_crud: ", err)
	}

	fmt.Println("Conexão bem-sucedida com o banco de dados!")

	// Cria/atualiza as tabelas
	fmt.Println("Criando/atualizando tabelas...")
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Erro ao criar tabelas: ", err)
	}
	fmt.Println("Tabelas criadas/atualizadas com sucesso!")
}
