package models

type User struct {
	ID       uint   `gorm:"primaryKey,column:id"`
	Username string `gorm:"unique,column:username"`
	Password string `gorm:"column:phone"`
}

//func GetUser(c *fiber.Ctx) (*User, error) {
//	user := &User{
//		BanDivision: make(map[int]*time.Time),
//	}
//	if config.Config.Mode == "dev" || config.Config.Mode == "test" {
//		user.ID = 1
//		user.IsAdmin = true
//		user.HasAnsweredQuestions = true
//		return user, nil
//	}
//
//	if c.Locals("user") != nil {
//		return c.Locals("user").(*User), nil
//	}
//
//	// get id
//	userID, err := common.GetUserID(c)
//	if err != nil {
//		return nil, err
//	}
//
//	// parse JWT
//	err = common.ParseJWTToken(common.GetJWTToken(c), user)
//	if err != nil {
//		return nil, err
//	}
//
//	// load user from database in transaction
//	err = user.LoadUserByID(userID)
//
//	if user.IsAdmin {
//		user.Permission.Admin = maxTime
//	} else {
//		user.Permission.Admin = minTime
//	}
//	user.Permission.Silent = user.BanDivision
//	user.Permission.OffenseCount = user.OffenceCount
//
//	if config.Config.UserAllShowHidden {
//		user.Config.ShowFolded = "hide"
//	}
//
//	// save user in c.Locals
//	c.Locals("user", user)
//
//	return user, err
//}
