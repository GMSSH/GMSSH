package core

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/DemonZack/simplejrpc-go/core/config"
	"github.com/DemonZack/simplejrpc-go/core/gvalid"
	"github.com/DemonZack/simplejrpc-go/os/gpath"
)

// Verification rules: You can implement them yourself according to the basic template provided
type ExampleUser struct {
	// Username string `validate:"required#Required parameters are missing name|min_length:6#The length is too small"`

	Username string `validate:"min_length:6#The length is too small"`
	Age      any    `validate:"required#Required parameters are missing Age|range:18,100|int#Test verification error return"`

	Email string `validate:"required#Email address is required"`
}

func TestExampleValidate(t *testing.T) {
	env := "test"
	gpath.GmCfgPath = filepath.Join(filepath.Dir(""), "..")
	InitContainer(config.WithConfigEnvFormatterOptionFunc(env))

	user := ExampleUser{
		Username: "abcggg", // Does not meet min_length: 6
		// Age:      16,    // Does not meet range:18,100
		Email: "", // Does not meet required
	}

	err := Container.Valid().Walk(&user)
	if err != nil {
		t.Fatalf("valid failed : %v ", err)
	}
	fmt.Println("valid  successfully")
}

type CustomValidator struct {
	gvalid.ValidatorErrorMessage
}

func (c CustomValidator) Validate(field *gvalid.FieldInfo, value any) error {
	if _, ok := value.(int); !ok {
		return c.NewValidationError(field.Field.Name, "Please enter a numeric type").WithMessage()
	}
	return nil
}

func TestCustomValidate(t *testing.T) {
	env := "test"
	gpath.GmCfgPath = filepath.Join(filepath.Dir(""), "..")
	container := InitContainer(config.WithConfigEnvFormatterOptionFunc(env))

	// Adding a custom validation structure
	container.Valid().RegisterValidator("int", &CustomValidator{})

	user := ExampleUser{
		Username: "abcggg", //  Does not meet min_length:6
		// Age:      16,    //  Does not meet range:18,100
		Age:   "123",        //  Does not meet int
		Email: "xxx@qq.com", //  Does not meet required
	}

	err := Container.Valid().Walk(&user)
	if err != nil {
		t.Fatalf("valid failed : %v ", err)
	}
	fmt.Println("valid  successfully")
}

type Example2User struct {
	Age any `myvalidate:"int#Test verification error return"`
}

func TestCustomValidatorVisitor(t *testing.T) {
	env := "test"
	gpath.GmCfgPath = filepath.Join(filepath.Dir(""), "..")
	container := InitContainer(config.WithConfigEnvFormatterOptionFunc(env))

	// 自定义校验器
	visitor := gvalid.NewValidatorVisitor()
	visitor.RegisterValidator("int", &CustomValidator{})
	walker := gvalid.NewStructWalker(visitor, "myvalidate")
	container.Clone(WithContainerValidOption(walker))

	user := Example2User{
		Age: 12, // 测试满足int结构
	}

	err := Container.Valid().Walk(&user)
	if err != nil {
		t.Fatalf("valid failed : %v ", err)
	}
	fmt.Println("valid  successfully")
}
