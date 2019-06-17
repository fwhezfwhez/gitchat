package project

import (
	"fmt"
	"github.com/fwhezfwhez/model_convert"
	"testing"
)

func TestSql2Model(t *testing.T) {
	dataSouce := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s", "localhost", "5432", "postgres", "test", "disable", "123")
	// tableName := "user_prop"
	// tableName := "prop"
	// tableName := "user_activity_process"
	tableName := "activity_config"
	fmt.Println(model_convert.TableToStructWithTag(dataSouce, tableName))
}
