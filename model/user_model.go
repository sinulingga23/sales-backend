package model

import (
	"log"
	"errors"
	"fmt"
	"sales-backend/utility"
)
type User struct {
	UserId		string	`json:"userId"`
	RoleId		int 	`json:"roleId"`
	AddressId	int	`json:"addressId"`
	FirstName	string	`json:"firstName"`
	LastName	string	`json:"lastName"`
	Gender		string	`json:"gender"`
	Address		string	`json:"address"`
	Email		string	`json:"address"`
	PhoneNumber	string	`json:"phoneNumber"`
	Audit		Audit	`json:"audit"`
}

type UserRegister struct {
	UserId		string	`json:"userId"`
	RoleId		int 	`json:"roleId"`
	ProvinceId	int 	`json:"provinceId"`
	CityId		int 	`json:"cityId"`
	SubDistrictId	int 	`json:"subDistrictId"`
	FirstName	string	`json:"firstName"`
	LastName	string	`json:"lastName"`
	Gender		string	`json:"gender"`
	Address		string	`json:"address"`
	Email		string	`json:"email"`
	Password	string	`json:"password"`
	PhoneNumber	string	`json:"phoneNumber"`
	Audit		Audit	`json:"audit"`
}


func (uR *UserRegister) IsPhoneNumberExists(phoneNumber string) (bool, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return false, err
	}
	defer db.Close()

	check := 0
	err = db.QueryRow("SELECT COUNT(phone_number) FROM users WHERE phone_number = ?", phoneNumber).Scan(&check)
	if err != nil {
		return false, err
	}

	if check != 1 {
		return false, err
	}

	return true, nil
}

// This function is just temporary implemenation, soon will be moved
// into associated struct
func (uR *UserRegister) IsRoleExistsById(roleId int) (bool, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return false, errors.New("Somethings wrong!")
	}
	defer db.Close()

	check := 0
	err = db.QueryRow("SELECT COUNT(role_id) FROM roles WHERE role_id = ?", roleId).Scan(&check)
	if err != nil {
		log.Printf("%s", err)
		return false, err
	}

	if check != 1 {
		return false, nil
	}

	return true, nil
}

func (uR *UserRegister) IsUserExistsByEmail(email string) (bool, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return false, errors.New("Somethings wrong!")
	}
	defer db.Close()

	check := 0
	err  = db.QueryRow("SELECT COUNT(user_id) FROM users WHERE email = ?", email).Scan(&check)
	if err != nil {
		return false, err
	}

	if check != 1 {
		return false, nil
	}

	return true, nil
}

func (uR *UserRegister) IsUserExistsById(userId string) (bool, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		log.Printf("line 85: %s", err)
		return false, errors.New("Somethings wrong!")
	}
	defer db.Close()

	check := 0
	err = db.QueryRow("SELECT COUNT(user_id) FROM users WHERE user_id = ?", userId).Scan(&check)
	if err != nil {
		log.Printf("line 92: %s", err)
		return false, err
	}

	if check != 1 {
		log.Printf("line 97: %s, %d", err, check)
		return false, nil
	}

	return true, nil
}

func (uR *UserRegister) SaveUser() (*User, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return &User{}, err
	}
	defer db.Close()

	number := 0
	err = db.QueryRow("SELECT COUNT(user_id) FROM users").Scan(&number)
	if err != nil {
		log.Printf("%s", err)
		return &User{}, err
	}

	count := utility.DigitsCount(number)
	formatUserId := "USR00000000"
	userId := "USR"
	for i := 0; i < len(formatUserId)-count-5; i++ {
		userId += "0"
	}
	number += 1
	userId += fmt.Sprintf("%d", number)
	uR.UserId = userId

	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		log.Printf("%s", err)
	}

	loginModel := Login{}
	hashedPassword, err := loginModel.EncryptPassword(uR.Password)
	if err != nil {
		return &User{}, err
	}

	result, err := tx.Exec("INSERT INTO address (province_id, city_id, sub_district_id, address, created_at) VALUES (?, ?, ?, ?, ?)",
		uR.ProvinceId,
		uR.CityId,
		uR.SubDistrictId,
		uR.Address,
		uR.Audit.CreatedAt)
	if err != nil {
		tx.Rollback()
		return &User{}, err
	}

	addressId, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return &User{}, err
	}

	_, err = tx.Exec("INSERT INTO users (user_id, address_id, first_name, last_name, gender, address, email, password, phone_number, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		uR.UserId,
		addressId,
		uR.FirstName,
		uR.LastName,
		uR.Gender,
		uR.Address,
		uR.Email,
		hashedPassword,
		uR.PhoneNumber,
		uR.Audit.CreatedAt)
	if err != nil {
		tx.Rollback()
		return &User{}, err
	}

	_, err = tx.Exec("INSERT INTO users_roles (user_id, role_id) VALUES (?, ?)", uR.UserId, uR.RoleId)
	if err != nil {
		tx.Rollback()
		return &User{}, err
	}

	err = tx.Commit()
	if err != nil {
		return &User{}, err
	}

	createdUser := &User{
		UserId: 	uR.UserId,
		RoleId: 	uR.RoleId,
		AddressId:	int(addressId),
		FirstName:	uR.FirstName,
		LastName:	uR.LastName,
		Gender:		uR.Gender,
		Address:	uR.Address,
		Email:		uR.Email,
		PhoneNumber:	uR.PhoneNumber,
		Audit:		Audit{CreatedAt: uR.Audit.CreatedAt},
	}

	return createdUser, nil
}

func (uR *UserRegister) FindUserById(userId string) (*User, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return &User{}, err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		return &User{}, err
	}

	currentUser := &User{}
	err = tx.QueryRow("SELECT user_id, address_id, first_name, last_name, gender, address, email, phone_number, created_at, updated_at FROM users WHERE user_id = ?", userId).
		Scan(
			&currentUser.UserId,
			&currentUser.AddressId,
			&currentUser.FirstName,
			&currentUser.LastName,
			&currentUser.Gender,
			&currentUser.Address,
			&currentUser.Email,
			&currentUser.PhoneNumber,
			&currentUser.Audit.CreatedAt,
			&currentUser.Audit.UpdatedAt)
	if err != nil {
		tx.Rollback()
		return &User{}, err
	}

	err = tx.QueryRow("SELECT role_id FROM users_roles WHERE user_id = ?", userId).Scan(&currentUser.RoleId)
	if err != nil {
		tx.Rollback()
		return &User{}, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return &User{}, err
	}

	if currentUser == (&User{}) {
		return &User{}, errors.New("Somethings wrong!")
	}

	return currentUser, nil
}

func (uR *UserRegister) GetRoleIdByEmail(email string) (int, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	tx, err := db.Begin()

	userId := ""
	err = tx.QueryRow("SELECT user_id FROM users WHERE email = ?", email).Scan(&userId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	roleId := 0
	err = tx.QueryRow("SELECT role_id FROM users_roles WHERE user_id =?", userId).Scan(&roleId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return roleId, nil
}
