package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-postgres/models"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Response struct {
	Id      int64  `json:"id,omitempty"`
	Message string `json:"message"`
}

func CreateConnection() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to database!")
	return db
}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock

	err := json.NewDecoder(r.Body).Decode(&stock)

	if err != nil {
		log.Fatal("Error reading the request body", err)
		return
	}

	insertId := insertStock(stock)
	res := Response{
		Id:      insertId,
		Message: "Stock created successfully",
	}

	json.NewEncoder(w).Encode(res)
}

func GetStocks(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stockId := params["id"]

	id, err := strconv.Atoi(stockId)

	if err != nil {
		log.Fatal("Error converting to int", err)
		return
	}

	stock, err := getStock(int64(id))

	if err != nil {
		log.Fatal("Error getting the stock", err)
		return
	}

	json.NewEncoder(w).Encode(stock)
}

func GetAllStocks(w http.ResponseWriter, r *http.Request) {
	stocks, err := getAllStocks()

	if err != nil {
		log.Fatal("Error getting the stocks", err)
		return
	}

	json.NewEncoder(w).Encode(stocks)
}

func UpdateStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stockId := params["id"]

	id, err := strconv.Atoi(stockId)

	if err != nil {
		log.Fatal("Error converting to int", err)
		return
	}

	var stock models.Stock

	err = json.NewDecoder(r.Body).Decode(&stock)

	if err != nil {
		log.Fatal("Error reading the request body", err)
		return
	}

	updatedRows := updateStock(int64(id), stock)

	res := Response{
		Id:      int64(updatedRows),
		Message: "Stock updated successfully",
	}

	json.NewEncoder(w).Encode(res)
}

func DeleteStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stockId := params["id"]

	id, err := strconv.Atoi(stockId)

	if err != nil {
		log.Fatal("Error converting to int", err)
		return
	}

	deletedRows := deleteStock(int64(id))

	res := Response{
		Id:      int64(deletedRows),
		Message: "Stock deleted successfully",
	}

	json.NewEncoder(w).Encode(res)

}

func insertStock(model models.Stock) int64 {
	db := CreateConnection()
	defer db.Close()
	sqlStatement := `INSERT INTO stocks (name, price, company) VALUES ($1, $2, $3) RETURNING id`
	var id int64

	db.QueryRow(sqlStatement, model.Name, model.Price, model.Company).Scan(&id)

	fmt.Printf("Inserted a single record %v", id)
	return id
}

func getStock(id int64) (models.Stock, error) {
	db := CreateConnection()
	defer db.Close()
	var stock models.Stock
	sqlStatement := `SELECT * FROM stocks WHERE stockid=$1`
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&stock.StockId, &stock.Name, &stock.Price, &stock.Company)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned")
		return stock, nil
	case nil:
		return stock, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	return stock, err
}

func getAllStocks() ([]models.Stock, error) {
	db := CreateConnection()
	defer db.Close()
	var stocks []models.Stock
	sqlStatement := `SELECT * FROM stocks`
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var stock models.Stock
		err = rows.Scan(&stock.StockId, &stock.Name, &stock.Price, &stock.Company)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		stocks = append(stocks, stock)
	}

	return stocks, err
}

func updateStock(id int64, model models.Stock) int64 {
	db := CreateConnection()
	defer db.Close()
	sqlStatement := `UPDATE stocks SET name=$2, price=$3, company=$4 WHERE stockid=$1`
	res, err := db.Exec(sqlStatement, id, model.Name, model.Price, model.Company)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error getting rows affected. %v", err)
	}

	fmt.Printf("Rows affected %v", rowsAffected)
	return rowsAffected
}

func deleteStock(id int64) int64 {
	db := CreateConnection()
	defer db.Close()
	sqlStatement := `DELETE FROM stocks WHERE stockid=$1`
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error getting rows affected. %v", err)
	}

	fmt.Printf("Rows affected %v", rowsAffected)
	return rowsAffected
}
