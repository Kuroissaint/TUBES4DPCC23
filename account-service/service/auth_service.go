package service

import (
	"account-service/model"
	"account-service/repository"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("rahasia-kelompok-23-c2")

type AuthService interface {
	Login(req model.LoginRequest) (*model.LoginResponse, error)
	Register(req model.RegisterRequest) (*model.RegisterResponse, error)
	//FUNGSI REGISTER DRIVER
	RegisterDriver(req model.DriverRegisterRequest) (*model.RegisterResponse, error)
    RegisterCustomer(req model.CustomerRegisterRequest) (*model.RegisterResponse, error)
}

type authService struct {
	repo repository.UserRepository
}


func NewAuthService(repo repository.UserRepository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) Login(req model.LoginRequest) (*model.LoginResponse, error) {
	user, err := s.repo.GetByEmail(req.Email)
	if err != nil || user == nil {
		return nil, errors.New("email atau password salah")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(req.Password))
	if err != nil {
		return nil, errors.New("email atau password salah")
	}

	token, err := generateJWT(user.ID, user.Email, user.Nama, user.NoHp, user.Role)
	if err != nil {
		return nil, errors.New("gagal membuat token")
	}

	return &model.LoginResponse{
		UserID: user.ID,
		Token:  token,
	}, nil
}

//REGISTER KHUSUS DRIVER
func (s *authService) RegisterDriver(req model.DriverRegisterRequest) (*model.RegisterResponse, error) {
	// 1. Cek email sudah ada
	existing, _ := s.repo.GetByEmail(req.Email)
	if existing != nil {
		return nil, errors.New("email sudah terdaftar")
	}

	// 2. Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("gagal memproses password")
	}

	// 3. Simpan ke users + driver_profiles sekaligus
	user, err := s.repo.RegisterDriver(
		model.User{
			Nama:           req.Nama,
			Email:          req.Email,
			NoHp:           req.NoHp,
			HashedPassword: string(hashedPassword),
		},
		model.DriverProfile{
			NoSim:          req.NoSim,
			NoPlat:         req.NoPlat,
			JenisKendaraan: req.JenisKendaraan,
		},
	)
	if err != nil {
		return nil, err
	}

	return &model.RegisterResponse{
		UserID:  user.ID,
		Message: "Registrasi driver berhasil!",
	}, nil
}

//REGISTER KHUSUS CUSTOMER//
func (s *authService) RegisterCustomer(req model.CustomerRegisterRequest) (*model.RegisterResponse, error) {
	existing, _ := s.repo.GetByEmail(req.Email)
	if existing != nil {
		return nil, errors.New("email sudah terdaftar")
	}

	if len(req.Password) < 8 {
		return nil, errors.New("password minimal 8 karakter")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("gagal memproses password")
	}

	user, err := s.repo.RegisterCustomer(
		model.User{
			Nama:           req.Nama,
			Email:          req.Email,
			NoHp:           req.NoHp,
			HashedPassword: string(hashedPassword),
		},
		req.Alamat,
	)
	if err != nil {
		return nil, err
	}

	return &model.RegisterResponse{
		UserID:  user.ID,
		Message: "Registrasi customer berhasil!",
	}, nil
}

func (s *authService) Register(req model.RegisterRequest) (*model.RegisterResponse, error) {
	existing, _ := s.repo.GetByEmail(req.Email)
	if existing != nil {
		return nil, errors.New("email sudah terdaftar")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("gagal memproses password")
	}

	role := req.Role
	if role == "" {
		role = "user"
	}

	user, err := s.repo.Register(model.User{
		Nama:           req.Nama,
		Email:          req.Email,
		NoHp:           req.NoHp,
		HashedPassword: string(hashedPassword),
		Role:           role,
	})
	if err != nil {
		return nil, err
	}

	return &model.RegisterResponse{
		UserID:  user.ID,
		Message: "Registrasi berhasil!",
	}, nil
}

func generateJWT(userID int, email string, nama string, noHp string, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"nama":    nama,
		"no_hp":   noHp,
		"role":    role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
