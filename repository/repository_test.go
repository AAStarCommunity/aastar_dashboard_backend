package repository

import (
	"testing"
)

func TestRepository(t *testing.T) {

}

//func GetDbClient() *gorm.DB {
//	onceDb.Do(func() {
//		if "1" == os.Getenv("UnitTestEnv") {
//			db, _ = getInMemoryDbClient()
//			db = db.Debug()
//		} else {
//			dsn := getDbConnectionString(getConf())
//			_db, err := gorm.Open(postgres.Open(*dsn), &gorm.Config{
//				DisableForeignKeyConstraintWhenMigrating: true,
//			})
//			if err != nil {
//				panic(err)
//			}
//
//			if Environment.IsDevelopment() {
//				_db = _db.Debug()
//			}
//			db = _db
//		}
//	})
//	return db
//}

//if err := os.Setenv("UnitTestEnv", "1"); err != nil {
//t.Skip("skipped due to CI")
//} else {
//defer func() {
//os.Unsetenv("UnitTestEnv")
//}()
//}
