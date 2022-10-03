package ognoc

const (
	Alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Number   = "0123456789"
	Spechar  = "!@#$%&*()_+-=?.,:;<>"
	Mixed    = Alphabet + Number + Spechar
)

// Use Option Design Pattern
type Option func() string

// WithAlphabet returns alphabet in cipher
func WithAlphabet() Option {
	return func() string {
		return Alphabet
	}
}

// WithNumber returns number in cipher
func WithNumber() Option {
	return func() string {
		return Number
	}
}

// WithSpechar returns spechar in cipher
func WithSpechar() Option {
	return func() string {
		return Spechar
	}
}

// NewPassword generates a complicated cipher
func NewPassword(options ...Option) string {
	if len(options) != 0 {
		cipher := ""
		for _, opt := range options {
			cipher += opt()
		}
		return GenerateRandom(cipher)
	}

	return GenerateRandom(Mixed)
}
