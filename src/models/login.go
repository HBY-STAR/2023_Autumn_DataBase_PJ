package models

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opentreehole/go-common"
	"gorm.io/gorm"
)

type TmpUser struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"password"`
	Type     string `json:"type"`
}

// 传入type
// user里有type id

func GetGeneralUser(c *fiber.Ctx) (*TmpUser, error) {
	//fmt.Println(c.Get("type"))
	user := &TmpUser{}

	if c.Locals("user") != nil {
		return c.Locals("user").(*TmpUser), nil
	}

	// get id
	var err error
	user.ID, err = common.GetUserID(c)
	if err != nil {
		return nil, err
	}

	// parse JWT
	err = common.ParseJWTToken(common.GetJWTToken(c), user)
	if err != nil {
		return nil, err
	}

	user.Type = c.Get("type")

	// load user from database in transaction
	err = user.CheckUserID()

	//if user.IsAdmin {
	//	user.Permission.Admin = maxTime
	//} else {
	//	user.Permission.Admin = minTime
	//}
	//user.Permission.Silent = user.BanDivision
	//user.Permission.OffenseCount = user.OffenceCount
	//
	//if config.Config.UserAllShowHidden {
	//	user.Config.ShowFolded = "hide"
	//}

	if err != nil {
		return nil, err
	}
	// save user in c.Locals
	c.Locals("user", user)

	return user, nil
}

func (user *TmpUser) CheckUserID() error {
	return DB.Transaction(func(tx *gorm.DB) error {
		if user.Type == "seller" {
			var seller = Seller{ID: user.ID}
			err := tx.Take(&seller).Error
			if err != nil {
				return err
			}
			user.Username = seller.Username
		} else if user.Type == "admin" {
			var newUser = User{ID: user.ID}
			err := tx.Take(&newUser).Error
			if err != nil {
				return err
			}
			user.Username = newUser.Username
		} else {
			var admin = Admin{ID: user.ID}
			err := tx.Take(&admin).Error
			if err != nil {
				return err
			}
			user.Username = admin.Username
		}
		//err := tx.Take(&user, userID).Error
		//if err != nil {
		//	if errors.Is(err, gorm.ErrRecordNotFound) {
		//		// insert user if not found
		//		user.ID = userID
		//		user.Config = defaultUserConfig
		//		err = tx.Create(&user).Error
		//		if err != nil {
		//			return err
		//		}
		//	} else {
		//		return err
		//	}
		//}

		return nil
	})
}
