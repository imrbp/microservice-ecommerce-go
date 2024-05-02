package dependency

import ("gorm.io/gorm"
"gorm.io/driver/postgres")

type Database struct {
  DB *gorm.DB
}

func NewDatabase() *gorm.DB {
  stringConnection := "host=localhost user=imr.bp password=/ dbname=order port=5432"
	db, err := gorm.Open(postgres.Open(stringConnection), &gorm.Config{})

	if err != nil {
		panic(db)
	}

  return db 
}
