package cmd

import (
	"github.com/jinzhu/gorm"
	"github.com/sanderdsz/codepix/application/grpc"
	"github.com/sanderdsz/codepix/infrastructure/db"
	"os"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}
