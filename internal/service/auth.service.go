package service

import (
	"errors"

	"github.com/rlanakomara7/koda-b9-gin.git/internal/dto"
)

type AuthService struct {
	users []dto.User
}

func NewAuthService() *AuthService {
	return &AuthService{
		users: []dto.User{},
	}
}

// Validasi dan Logika Register
func (s *AuthService) Register(user dto.User) error {
	if len(user.Username) <= 6 || len(user.Email) <= 6 || len(user.Password) <= 6 {
		return errors.New("username, email, dan password harus lebih dari 6 karakter")
	}

	// Mengecek apakah email sudah pernah didaftarkan
	for _, u := range s.users {
		if u.Email == user.Email {
			return errors.New("email sudah terdaftar")
		}
	}

	// Simpan user baru
	s.users = append(s.users, user)
	return nil
}

// Validasi dan Logika Login
func (s *AuthService) Login(req dto.User) (dto.User, error) {
	if len(req.Email) <= 6 || len(req.Password) <= 6 {
		return dto.User{}, errors.New("email atau password tidak valid")
	}

	for _, u := range s.users {
		if u.Email == req.Email {
			if u.Password == req.Password {
				return u, nil // Berhasil login
			}
			return dto.User{}, errors.New("password salah")
		}
	}

	return dto.User{}, errors.New("email tidak ditemukan")
}
