package infra

import (
	"path"
	"testing"
)

func TestDatabaseFunctions(t *testing.T) {

	t.Run("DBFunctions", func(t *testing.T) {
		config := &Config{
			ConnectionString: path.Join(t.TempDir(), "feiras_test.db"),
		}
		db, err := GetDatabase(config)
		if err != nil {
			t.Errorf("GetDatabase() error = %v", err)
			return
		}
		if err = ResetDatabase(db); err != nil {
			t.Errorf("ResetDatabase() error = %v", err)			
		}
	})

}
