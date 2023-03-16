package passwords

import "golang.org/x/crypto/bcrypt"

type Generator interface {
	// GeneratePasswordHash generates a password hash for storage in the database
	GeneratePasswordHash(password string) (string, error)
	// CheckPasswordHash checks if the received password matches the one stored in the database
	CheckPasswordHash(password, hash string) bool
}

type generatePasswordHash struct {
	cost int
}

func NewGeneratePasswordHash(cost int) *generatePasswordHash {
	return &generatePasswordHash{
		cost: cost,
	}
}

func (g generatePasswordHash) GeneratePasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), g.cost)
	return string(bytes), err
}

func (g generatePasswordHash) CheckPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
