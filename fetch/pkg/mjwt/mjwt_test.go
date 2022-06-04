package mjwt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMjwt_ValidateToken(t *testing.T) {
	token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJmcmVzaCI6dHJ1ZSwiaWF0IjoxNjU0MzAzOTcwLCJqdGkiOiI5NTU3NGM3Mi1jNGQ2LTQ2MmEtYmExNy00MDEwMjg2NzQ0YTYiLCJ0eXBlIjoiYWNjZXNzIiwic3ViIjoiKzgxMjMxNzQxIiwibmJmIjoxNjU0MzAzOTcwLCJleHAiOjI2NTQzMDQ4NzAsIm5hbWUiOiJtb3JhbGEiLCJwaG9uZSI6Iis4MTIzMTc0MSIsInJvbGUiOiJhZG1pbiIsInRpbWVzdGFtcCI6IjIwMjItMDYtMDMgMTc6MTA6NDUifQ.FAHrgv_BmvYmFeprBoyZZDfWtsnBifqDV1gIW9nEtG4"
	keyExample := "example"
	tokenValid, err := New(keyExample).ValidateToken(token)

	assert.Nil(t, err)
	assert.NotEmpty(t, tokenValid)
}

func TestMjwt_NotValidateToken(t *testing.T) {
	invalidToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9XX.eyJleHAiOjE2MDM4MDcyMzEsImlkZW50aXR5IjoibXVjaGxpc0BnbWFpbC5jb20iLCJpc19hZG1pbiI6dHJ1ZSwianRpIjoiIn0.dzKZdhPFtF-YC6uh5JZqBv7mhBjGTz1_rgIP-sRbYrU"
	keyExample := "example"
	tokenValid, err := New(keyExample).ValidateToken(invalidToken)

	assert.Empty(t, tokenValid)
	assert.NotNil(t, err)
	assert.Equal(t, "token not valid", err.Error())
}

func TestMjwt_ReadToken(t *testing.T) {
	token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJmcmVzaCI6dHJ1ZSwiaWF0IjoxNjU0MzAzOTcwLCJqdGkiOiI5NTU3NGM3Mi1jNGQ2LTQ2MmEtYmExNy00MDEwMjg2NzQ0YTYiLCJ0eXBlIjoiYWNjZXNzIiwic3ViIjoiKzgxMjMxNzQxIiwibmJmIjoxNjU0MzAzOTcwLCJleHAiOjI2NTQzMDQ4NzAsIm5hbWUiOiJtb3JhbGEiLCJwaG9uZSI6Iis4MTIzMTc0MSIsInJvbGUiOiJhZG1pbiIsInRpbWVzdGFtcCI6IjIwMjItMDYtMDMgMTc6MTA6NDUifQ.FAHrgv_BmvYmFeprBoyZZDfWtsnBifqDV1gIW9nEtG4"
	keyExample := "example"

	jwtMgr := New(keyExample)

	tokenValid, err := jwtMgr.ValidateToken(token)
	assert.Nil(t, err)
	if err != nil {
		assert.FailNow(t, err.Error())
		return
	}

	claims, err := jwtMgr.ReadToken(tokenValid)
	assert.Nil(t, err)

	assert.Equal(t, "+81231741", claims.Phone)
	assert.Equal(t, "admin", claims.Role)
}
