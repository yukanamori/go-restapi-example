package repository

import (
	"myapp/internal/domain/entity"
	"myapp/pkg/erreurs"
	"strings"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type (
	// UserRepository はユーザに関する永続化のためのインターフェースです。
	UserRepository interface {
		// Create はユーザを作成します。ユニーク制約に違反した場合はエラーを返します。
		Create(user *entity.User) error
		// FindByID はIDに一致するユーザを取得します。
		FindByID(id uint) (*entity.User, error)
		// FindByUsername はユーザ名に一致するユーザを取得します。
		FindByUsername(username string) (*entity.User, error)
		// FindByCondition は条件に一致するユーザを取得します。
		FindByCondition(condition *entity.User) ([]*entity.User, error)
		// FindAll はユーザを全件取得します。
		FindAll() ([]*entity.User, error)
		// Update はユーザを更新します。ユーザが存在しない、またはバージョンが一致しない場合はエラーを返します。
		Update(user *entity.User) error
		// Delete はユーザを削除します。ユーザが存在しない、またはバージョンが一致しない場合はエラーを返します。
		Delete(id uint, version uint) error
		// Exists はユーザが存在するかどうかを返します。
		Exists(id uint) bool
	}

	// UserRepositoryImpl はUserRepositoryの実装です。
	UserRepositoryImpl struct {
		db *gorm.DB
	}
)

// NewUserRepository はUserRepositoryを生成します。
func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db}
}

// Create はユーザを作成します。ユニーク制約に違反した場合はエラーを返します。
func (r *UserRepositoryImpl) Create(user *entity.User) error {
	if err := r.db.Create(user).Error; err != nil {
		if err := r.handleDuplicateKeyError(err); err != nil {
			if err == erreurs.ErrUsernameAlreadyExists {
				zap.L().Info("username already exists", zap.Error(err), zap.String("username", user.Username))
				return err
			} else if err == erreurs.ErrEmailAlreadyExists {
				zap.L().Info("email already exists", zap.Error(err), zap.String("email", user.Email))
				return err
			}
		}

		zap.L().Error("failed to create user", zap.Error(err))
		return err
	}

	return nil
}

func (r *UserRepositoryImpl) handleDuplicateKeyError(err error) error {
	if err == gorm.ErrDuplicatedKey {
		if strings.Contains(err.Error(), "username") {
			return erreurs.ErrUsernameAlreadyExists
		} else if strings.Contains(err.Error(), "email") {
			return erreurs.ErrEmailAlreadyExists
		}
	}

	return nil
}

// FindByID はIDに一致するユーザを取得します。
func (r *UserRepositoryImpl) FindByID(id uint) (*entity.User, error) {
	user := &entity.User{}
	if err := r.db.First(user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			zap.L().Info("user not found", zap.Error(err), zap.Uint("id", id))
			return nil, erreurs.ErrUserNotFound
		}

		zap.L().Error("failed to find user", zap.Error(err), zap.Uint("id", id))
		return nil, err
	}

	return user, nil
}

// FindByUsername はユーザ名に一致するユーザを取得します。
func (r *UserRepositoryImpl) FindByUsername(username string) (*entity.User, error) {
	user := &entity.User{}
	if err := r.db.Where("username = ?", username).First(user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			zap.L().Info("user not found", zap.Error(err), zap.String("username", username))
			return nil, erreurs.ErrUserNotFound
		}

		zap.L().Error("failed to find user", zap.Error(err), zap.String("username", username))
		return nil, err
	}

	return user, nil
}

// FindByCondition は条件に一致するユーザを取得します。
func (r *UserRepositoryImpl) FindByCondition(condition *entity.User) ([]*entity.User, error) {
	var users []*entity.User
	if err := r.db.Where(condition).Find(&users).Error; err != nil {
		zap.L().Error("failed to find users", zap.Error(err))
		return nil, err
	}

	return users, nil
}

// FindAll はユーザを全件取得します。
func (r *UserRepositoryImpl) FindAll() ([]*entity.User, error) {
	var users []*entity.User
	if err := r.db.Find(&users).Error; err != nil {
		zap.L().Error("failed to find users", zap.Error(err))
		return nil, err
	}

	return users, nil
}

// Update はユーザを更新します。ユーザーが存在しない、またはバージョンが一致しない、またはユニーク制約に違反した場合はエラーを返します。
func (r *UserRepositoryImpl) Update(user *entity.User) error {
	r.db.Transaction(func(tx *gorm.DB) error {
		u := &entity.User{}
		if err := tx.First(u, user.ID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				zap.L().Info("user not found", zap.Error(err), zap.Uint("id", user.ID))
				return erreurs.ErrUserNotFound
			}

			zap.L().Error("failed to find user", zap.Error(err), zap.Uint("id", user.ID))
			return err
		}

		if u.Version != user.Version {
			zap.L().Info("user version mismatch", zap.Uint("id", user.ID), zap.Uint("version", user.Version))
			return erreurs.ErrVersionMismatch
		}

		if err := tx.Save(user).Error; err != nil {
			if err := r.handleDuplicateKeyError(err); err != nil {
				if err == erreurs.ErrUsernameAlreadyExists {
					zap.L().Info("username already exists", zap.Error(err), zap.String("username", user.Username))
					return err
				} else if err == erreurs.ErrEmailAlreadyExists {
					zap.L().Info("email already exists", zap.Error(err), zap.String("email", user.Email))
					return err
				}
			}

			zap.L().Error("failed to update user", zap.Error(err), zap.Uint("id", user.ID))
			return err
		}

		return nil
	})

	return nil
}

// Delete はユーザを削除します。ユーザが存在しない、またはバージョンが一致しない場合はエラーを返します。
func (r *UserRepositoryImpl) Delete(id, version uint) error {
	r.db.Transaction(func(tx *gorm.DB) error {
		user := &entity.User{}
		if err := tx.First(user, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				zap.L().Info("user not found", zap.Error(err), zap.Uint("id", id))
				return erreurs.ErrUserNotFound
			}

			zap.L().Error("failed to find user", zap.Error(err), zap.Uint("id", id))
			return err
		}

		if user.Version != version {
			zap.L().Info("user version mismatch", zap.Uint("id", id), zap.Uint("version", version))
			return erreurs.ErrVersionMismatch
		}

		if err := tx.Delete(user).Error; err != nil {
			zap.L().Error("failed to delete user", zap.Error(err), zap.Uint("id", id))
			return err
		}

		return nil
	})

	return nil
}

// Exists はユーザが存在するかどうかを返します。
func (r *UserRepositoryImpl) Exists(id uint) bool {
	user := &entity.User{}
	if err := r.db.First(user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false
		}

		zap.L().Error("failed to find user", zap.Error(err), zap.Uint("id", id))
		return false
	}

	return true
}
