package usecase

import (
	"myapp/internal/domain/entity"
	"myapp/internal/domain/repository"
	"myapp/pkg/auth"

	"go.uber.org/zap"
)

type (
	// UserUsecase はユーザに対するユースケースです。
	UserUsecase interface {
		// Create はユーザを作成します。
		Create(username, password, email, firstName, lastName, profileImage string) error
		// GetByID はIDに一致するユーザを取得します。
		GetByID(id uint) (*entity.User, error)
		// Search は条件に一致するユーザを取得します。
		Search(condition *entity.User) ([]*entity.User, error)
		// GetAll はユーザを全件取得します。
		GetAll() ([]*entity.User, error)
		// Update はユーザを更新します。
		Update(id, version uint, username, email, firstName, lastName, profileImage string) error
		// UpdatePassword はユーザのパスワードを更新します。
		UpdatePassword(username, password string) error
		// Delete はユーザを削除します。
		Delete(id, version uint) error
	}

	// UserUsecaseImpl はUserUsecaseの実装です。
	UserUsecaseImpl struct {
		userRepository repository.UserRepository
	}
)

// NewUserUsecase はUserUsecaseを生成します。
func NewUserUsecase(userRepository repository.UserRepository) *UserUsecaseImpl {
	return &UserUsecaseImpl{userRepository}
}

// Create はユーザを作成します。
func (u *UserUsecaseImpl) Create(username, password, email, firstName, lastName, profileImage string) error {
	hashedPassword, err := auth.HashPassword(password)
	if err != nil {
		zap.L().Error("failed to hash password", zap.Error(err), zap.String("username", username))
		return err
	}

	user := &entity.User{
		Username:     username,
		Password:     hashedPassword,
		Email:        email,
		FirstName:    firstName,
		LastName:     lastName,
		ProfileImage: profileImage,
	}

	if err := u.userRepository.Create(user); err != nil {
		return err
	}

	return nil
}

// GetByID はIDに一致するユーザを取得します。
func (u *UserUsecaseImpl) GetByID(id uint) (*entity.User, error) {
	return u.userRepository.FindByID(id)
}

// Search は条件に一致するユーザを取得します。
func (u *UserUsecaseImpl) Search(condition *entity.User) ([]*entity.User, error) {
	return u.userRepository.FindByCondition(condition)
}

// GetAll はユーザを全件取得します。
func (u *UserUsecaseImpl) GetAll() ([]*entity.User, error) {
	return u.userRepository.FindAll()
}

// Update はユーザを更新します。
func (u *UserUsecaseImpl) Update(id, version uint, username, email, firstName, lastName, profileImage string) error {
	user, err := u.userRepository.FindByID(id)
	if err != nil {
		return err
	}

	if username != "" {
		user.Username = username
	}

	if email != "" {
		user.Email = email
	}

	if firstName != "" {
		user.FirstName = firstName
	}

	if lastName != "" {
		user.LastName = lastName
	}

	if profileImage != "" {
		user.ProfileImage = profileImage
	}

	user.Version = version

	if err := u.userRepository.Update(user); err != nil {
		return err
	}

	return nil
}

// UpdatePassword はユーザのパスワードを更新します。
func (u *UserUsecaseImpl) UpdatePassword(username, password string) error {
	user, err := u.userRepository.FindByUsername(username)
	if err != nil {
		return err
	}

	hashedPassword, err := auth.HashPassword(password)
	if err != nil {
		zap.L().Error("failed to hash password", zap.Error(err), zap.String("username", username))
	}

	user.Password = hashedPassword

	if err := u.userRepository.Update(user); err != nil {
		return err
	}

	return nil
}

// Delete はユーザを削除します。
func (u *UserUsecaseImpl) Delete(id, version uint) error {
	return u.userRepository.Delete(id, version)
}
