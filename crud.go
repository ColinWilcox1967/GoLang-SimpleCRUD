package crud

func CreateUser(name string, email string) (*User, error) {
	user := User{Name: name, Email: email}
	result := DB.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func GetAllUsers() ([]User, error) {
	var users []User

	result := DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func GetUserByID(id uint) (*User, error) {
	var user User

	result := DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
func UpdateUser(id uint, name string, email string) (*User, error) {
	var user User

	result := DB.First(&user, id)

	if result.Error != nil {
		return nil, result.Error
	}

	user.Name = name
	user.Email = email
	DB.Save(&user)

	return &user, nil
}

func DeleteUser(id uint) error {
	result := DB.Delete(&User{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}