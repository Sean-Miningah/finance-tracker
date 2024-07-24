package user

import (
	"finance-crud-app/internal/testutils"
	"finance-crud-app/internal/types"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	code := m.Run()
	clearTestUserStore()
	os.Exit(code)
}

func TestUserStore(t *testing.T) {
	test_data := map[string]struct {
		user   types.User
		result error
	}{
		"valid user data": {
			user: types.User{
				FirstName: "testfirst-1",
				LastName:  "testlast-1",
				Email:     "email@test.com",
				Password:  "00000000",
			},
			result: nil,
		},
		"invalid user data": {
			user: types.User{
				FirstName: "testfirst-1",
				LastName:  "testlast-1",
				Email:     "email@test.com",
				Password:  "00000000",
			},
			result: CreateUserError,
		},
	}

	test_store := NewStore(testutils.DB)

	for name, tc := range test_data {
		t.Run(name, func(t *testing.T) {
			got := test_store.CreateUser(tc.user)
			if got != tc.result {
				t.Errorf("test fail expected %v got %v instead", tc.result, got)
			}
		})
	}

	// t.Run("successfully create user", func(t *testing.T) {
	// 	testUser := types.User{
	// 		FirstName: "testfirst-1",
	// 		LastName:  "testlast-1",
	// 		Email:     "email@test.com",
	// 		Password:  "00000000",
	// 	}

	// 	got := test_store.CreateUser(testUser)
	// 	if got != nil {
	// 		t.Errorf("test fail: got %v  but wanted nil", got)
	// 	}
	// })

	// t.Run("unsuccessfully create user", func(t *testing.T) {
	// 	testUser := types.User{

	// 	}
	// })
}

func clearTestUserStore() {
	_, err := testutils.DB.Exec("DELETE FROM users")
	if err != nil {
		log.Fatalf("Error clearing users test data %v", err)
	}
}
