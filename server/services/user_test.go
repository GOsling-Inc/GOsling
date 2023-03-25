package services_test

// import (
// 	"testing"

// 	"github.com/GOsling-Inc/GOsling/services"

// 	"github.com/stretchr/testify/assert"
// )

// func Test_User_BeforeCreate(t *testing.T) {
// 	u := services.TestUser(t)
// 	assert.NoError(t, HashPass(*u))
// }

// func Test_User_Validate(t *testing.T) {
// 	testCases := []struct {
// 		name    string
// 		u       func() *models.User
// 		isValid bool
// 	}{
// 		{
// 			name: "valid",
// 			u: func() *models.User {
// 				return services.TestUser(t)
// 			},
// 			isValid: true,
// 		},
// 		{
// 			name: "empty email",
// 			u: func() *models.User {
// 				u := services.TestUser(t)
// 				u.Email = ""

// 				return u
// 			},
// 			isValid: false,
// 		},
// 		{
// 			name: "invalid email",
// 			u: func() *models.User {
// 				u := services.TestUser(t)
// 				u.Email = "invalid"

// 				return u
// 			},
// 			isValid: false,
// 		},
// 		{
// 			name: "empty password",
// 			u: func() *models.User {
// 				u := services.TestUser(t)
// 				u.Password = ""

// 				return u
// 			},
// 			isValid: false,
// 		},
// 		{
// 			name: "short password",
// 			u: func() *models.User {
// 				u := services.TestUser(t)
// 				u.Password = "short"

// 				return u
// 			},
// 			isValid: false,
// 		},
// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			if tc.isValid {
// 				assert.NoError(t, tc.s.Validate(u()))
// 			} else {
// 				assert.Error(t, tc.s.Validate())
// 			}
// 		})
// 	}
// };
