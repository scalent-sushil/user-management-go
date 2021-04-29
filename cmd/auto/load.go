package auto

import (
	"log"

	"github.com/scalent-sushil/user-management-go/database"
	"github.com/scalent-sushil/user-management-go/pkg/models"
	"github.com/scalent-sushil/user-management-go/utils/console"
)

func Load() {
	// db, err := database.Connect()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = db.Debug().DropTableIfExists(&models.Blog{}, &models.User{}).Error
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = db.Debug().AutoMigrate(&models.Blog{}, &models.User{}).Error
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = db.Debug().Model(&models.Blog{}).AddForeignKey("Author_id", "users(id)", "cascade", "cascade").Error
	// if err != nil {
	// 	log.Fatal(err)
	// }

	for i, _ := range users {
		err := database.DB.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatal(err)
		}
		// blogs[i].AuthorID = users[i].ID
		// err = db.Debug().Model(&models.Blog{}).Create(&blogs).Error
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// err = db.Debug().Model(&blogs[i]).Related(&blogs[i].AuthorName).Error
		// if err != nil {
		// 	log.Fatal(err)
		// }
		console.Pretty(users)
	}
}
