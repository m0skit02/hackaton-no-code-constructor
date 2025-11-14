package repository

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	models "hackaton-no-code-constructor/pkg/model"
)

type UserPostgres struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserPostgres {
	logrus.Info("[UserRepo] Initialized UserPostgres repository")
	return &UserPostgres{db: db}
}

func (r *UserPostgres) Create(user models.User) (*models.User, error) {
	logrus.Infof("[UserRepo] Creating user_context: Username=%s", user.Username)
	if err := r.db.Create(&user).Error; err != nil {
		logrus.Errorf("[UserRepo] Failed to create user_context '%s': %v", user.Username, err)
		return nil, err
	}
	logrus.Infof("[UserRepo] User created successfully: ID=%s", user.ID)
	return &user, nil
}

func (r *UserPostgres) GetAll() ([]models.User, error) {
	logrus.Info("[UserRepo] Fetching all users")
	var users []models.User
	if err := r.db.Order("created_at asc").Find(&users).Error; err != nil {
		logrus.Errorf("[UserRepo] Failed to fetch users: %v", err)
		return nil, err
	}
	logrus.Infof("[UserRepo] Retrieved %d users", len(users))
	return users, nil
}

func (r *UserPostgres) GetByID(id uuid.UUID) (*models.User, error) {
	logrus.Infof("[UserRepo] Fetching user_context by ID=%s", id)
	var user models.User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			logrus.Warnf("[UserRepo] User not found: ID=%s", id)
			return nil, nil
		}
		logrus.Errorf("[UserRepo] Error fetching user_context by ID=%s: %v", id, err)
		return nil, err
	}
	logrus.Infof("[UserRepo] User found: ID=%s, Username=%s", user.ID, user.Username)
	return &user, nil
}

func (r *UserPostgres) GetByUsername(username string) (*models.User, error) {
	logrus.Infof("[UserRepo] Fetching user_context by Username=%s", username)
	var user models.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			logrus.Warnf("[UserRepo] User not found: Username=%s", username)
			return nil, nil
		}
		logrus.Errorf("[UserRepo] Error fetching user_context by Username=%s: %v", username, err)
		return nil, err
	}
	logrus.Infof("[UserRepo] User found: ID=%s, Username=%s", user.ID, user.Username)
	return &user, nil
}

func (r *UserPostgres) Update(user *models.User) (*models.User, error) {
	logrus.Infof("[UserRepo] Updating user_context: ID=%s", user.ID)
	if err := r.db.Save(user).Error; err != nil {
		logrus.Errorf("[UserRepo] Failed to update user_context ID=%s: %v", user.ID, err)
		return nil, err
	}
	logrus.Infof("[UserRepo] User updated successfully: ID=%s, Username=%s", user.ID, user.Username)
	return user, nil
}

func (r *UserPostgres) Delete(id uuid.UUID) error {
	logrus.Infof("[UserRepo] Deleting user_context ID=%s", id)
	if err := r.db.Delete(&models.User{}, "id = ?", id).Error; err != nil {
		logrus.Errorf("[UserRepo] Failed to delete user_context ID=%s: %v", id, err)
		return err
	}
	logrus.Infof("[UserRepo] User deleted successfully: ID=%s", id)
	return nil
}
