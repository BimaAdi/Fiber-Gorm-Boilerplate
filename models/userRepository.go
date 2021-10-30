package models

func (user User) GetAllUser(page int, size int) (*[]User, *int, error) {
	// get all users from database
	limit := size
	offset := (page - 1) * size
	users := []User{}
	if err := DBConn.Order("id DESC").Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, nil, err
	}

	// count all users in database
	var count int64
	if err := DBConn.Model(&users).Count(&count).Error; err != nil {
		return nil, nil, err
	}

	var intCount = int(count) // convert from int64 to int

	return &users, &intCount, nil
}

func (user User) GetDetailUser(id int) (*User, error) {
	if err := DBConn.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (user User) CreateUser(newUser User) (*User, error) {
	if err := DBConn.Create(&newUser).Error; err != nil {
		return nil, err
	}
	return &newUser, nil
}

func (user User) UpdateUser(updated User, id int) (*User, error) {
	if err := DBConn.First(&user, id).Error; err != nil {
		return nil, err
	}

	user.Username = updated.Username
	user.Password = updated.Password
	user.IsAdmin = updated.IsAdmin

	if err := DBConn.Save(user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (user User) DeleteUser(id int) error {
	if err := DBConn.Unscoped().Delete(&user, id).Error; err != nil {
		return err
	}
	return nil
}
