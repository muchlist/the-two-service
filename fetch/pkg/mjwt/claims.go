package mjwt

// {
// 	"fresh": true,
// 	"iat": 1654303970,
// 	"jti": "95574c72-c4d6-462a-ba17-4010286744a6",
// 	"type": "access",
// 	"sub": "+81231741",
// 	"nbf": 1654303970,
// 	"exp": 1654304870,
// 	"name": "morala",
// 	"phone": "+81231741",
// 	"role": "admin",
// 	"timestamp": "2022-06-03 17:10:45"
//  }
type CustomClaim struct {
	UniqueID  string `json:"jti"`
	Sub       string `json:"sub"`
	Phone     string `json:"phone"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Timestamp string `json:"timestamp"`
	Exp       int64  `json:"exp"`
	Type      string `json:"type"`
	Fresh     bool   `json:"fresh"`
}
