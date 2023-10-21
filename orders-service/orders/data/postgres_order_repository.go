package data

//
//import (
//	"context"
//	"database/sql"
//	"github.com/blazejwylegly/orders-service/orders/models"
//	"github.com/google/uuid"
//	"go.mongodb.org/mongo-driver/mongo"
//	"go.mongodb.org/mongo-driver/mongo/options"
//	"log"
//)
//
//type OrderPostgresRepository struct {
//	db  *sql.DB
//	ctx context.Context
//}
//
//const (
//	dbDriverName       = "postgres"
//	dbConnectionString = "postgresql://<username>:<password>@<database_ip>/todos?sslmode=disable"
//)
//
//func Create(ctx context.Context) *OrderPostgresRepository {
//	dbConnection, err := sql.Open(dbDriverName, dbConnectionString)
//	if err != nil {
//		log.Fatal(err)
//	}
//	return &OrderPostgresRepository{
//		db:  dbConnection,
//		ctx: ctx,
//	}
//}
//
//func EstablishDbConnection(ctx context.Context) *mongo.Database {
//	// TODO externalize configuration
//	opts := options.Client().ApplyURI("mongodb://localhost:27017")
//	client, err := mongo.Connect(ctx, opts)
//	if err != nil {
//		panic(err)
//	}
//	return client.Database("poc")
//}
//
//func (repository OrderPostgresRepository) GetAll() ([]models.Order, error) {
//	rows, err := repository.db.Query("SELECT * FROM TABLE ( ORDERS )")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	var fetchResult []models.Order
//	if err != nil {
//		return nil, err
//	}
//	return fetchResult, nil
//}
//
//func (repository OrderPostgresRepository) Get(uuid uuid.UUID) (*models.Order, error) {
//	//TODO implement me
//	log.Print("Get called")
//	return nil, nil
//}
//
//func (repository OrderPostgresRepository) Add(order models.Order) error {
//	_, err := repository.ordersCollection.InsertOne(repository.ctx, order)
//	if err != nil {
//		log.Fatal("Error trying to save db entry!")
//		return err
//	}
//	return nil
//}
