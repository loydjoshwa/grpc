package repository

import (
	"grpc/internal/model"

	"gorm.io/gorm"
)

type UserRepository struct{
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository{
	return &UserRepository{
		db:db,
	}
}

func (r *UserRepository) Create(user *model.User) error{
	return r.db.Create(user).Error
}

func (r *UserRepository) GetAll() ([]model.User,error){
	var users []model.User

	err:=r.db.Find(&users).Error

	if err!=nil{
		return nil,err
	}
	return users,nil
}

func (r *UserRepository) GetById(id uint) (*model.User,error){
	var user model.User

	err:=r.db.First(&user,id).Error

	if err!=nil{
		return nil,err
	}

	return &user,nil
}

func (r *UserRepository) Update(user *model.User) error{
	return r.db.Save(user).Error
}

func (r *UserRepository) Delete(id uint) error{
	return r.db.Delete(&model.User{},id).Error
}