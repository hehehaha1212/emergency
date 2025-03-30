
import (
    "crypto/sha256"
    "encoding/hex"
)

// HashPassword hashes the password using SHA-256
func HashPassword(password string) string {
    hash := sha256.Sum256([]byte(password))
    return hex.EncodeToString(hash[:])
}
