package infra

import (
	"fmt"
	"log"
	"path"
	"reflect"
	"testing"

	"github.com/guionardo/go-api-example/domain"
)

var (
	service *FeiraService
	feiras  = []domain.Feira{
		{
			ID:         1,
			Long:       100,
			Lat:        200,
			SetCens:    309102309123,
			AreaP:      300,
			CodDist:    40,
			Distrito:   "DISTRITO",
			CodSubPref: 19,
			SubPrefe:   "SUB_PREF",
			Regiao5:    "REGIAO_5",
			Regiao8:    "REGIAO_8",
			NomeFeira:  "NOME_FEIRA",
			Registro:   "REGISTRO",
			Logradouro: "LOGRADOURO",
			Numero:     "NUMERO",
			Bairro:     "BAIRRO",
			Referencia: "REFERENCIA",
		},
		{
			ID:         2,
			Long:       101,
			Lat:        202,
			SetCens:    309102309123321,
			AreaP:      303,
			CodDist:    60,
			Distrito:   "DISTRITO_2",
			CodSubPref: 20,
			SubPrefe:   "SUB_PREF_2",
			Regiao5:    "REGIAO_5_2",
			Regiao8:    "REGIAO_8_2",
			NomeFeira:  "NOME_FEIRA_2",
			Registro:   "REGISTRO_2",
			Logradouro: "LOGRADOURO_2",
			Numero:     "NUMERO_2",
			Bairro:     "BAIRRO_2",
			Referencia: "REFERENCIA_2",
		},
	}
)

func setupSuite(tb testing.TB, withData bool) func(tb testing.TB) {
	log.Printf("Setup repository test suite")
	connectionString := path.Join(tb.TempDir(), "feiras.db")
	var err error
	service, err = NewFeiraService(&Config{ConnectionString: connectionString})
	if err != nil {
		tb.Fatalf("Error creating service: %v", err)
	}
	if err = service.Reset(); err != nil {
		tb.Fatalf("Error reseting database service: %v", err)
	}

	if withData {
		for _, feira := range feiras {
			if err := service.Save(&feira); err != nil {
				tb.Fatalf("Error saving feira: %s", err)
			}
		}
	}
	return func(tb testing.TB) {
		log.Printf("Teardown repository test suite")
	}
}

func TestFindAllFeiras(t *testing.T) {
	teardownSuite := setupSuite(t, true)
	defer teardownSuite(t)
	t.Run("Default", func(t *testing.T) {
		feiras2, err := service.FindAll()
		if err != nil {
			t.Errorf("Error finding all feiras: %s", err)
			return
		}
		if len(feiras2) != len(feiras) {
			t.Errorf("Expected %d feiras: got %d", len(feiras), len(feiras2))
			return
		}
	})
}

func TestGetFeirasById(t *testing.T) {
	teardownSuite := setupSuite(t, true)
	defer teardownSuite(t)
	for _, feira := range feiras {
		t.Run(fmt.Sprintf("Feira #%d", feira.ID), func(t *testing.T) {
			findFeira, err := service.FindByID(feira.ID)
			if err != nil {
				t.Errorf("Error finding feira: %s", err)
				return
			}
			if findFeira == nil || findFeira.ID != feira.ID {
				t.Errorf("Expected feira %v: got %v", feira, findFeira)
				return
			}

		})
	}
}

func TestDeleteByRegistro(t *testing.T) {
	teardownSuite := setupSuite(t, true)
	defer teardownSuite(t)
	deleted := 0
	for _, feira := range feiras {
		t.Run(feira.Registro, func(t *testing.T) {
			if err := service.DeleteByRegistro(feira.Registro); err != nil {
				t.Errorf("Error deleting feira: %s", err)
				return
			}
			deleted++
			feiras2, err := service.FindAll()
			if err != nil {
				t.Errorf("Error finding all feiras: %s", err)
				return
			}
			if len(feiras2) != len(feiras)-deleted {
				t.Errorf("Expected %d feiras: got %d", len(feiras)-deleted, len(feiras2))
				return
			}
		})
	}
}

func TestFeiraService_Query(t *testing.T) {
	teardownSuite := setupSuite(t, true)
	defer teardownSuite(t)
	type args struct {
		distrito   string
		regiao5    string
		nome_feira string
		bairro     string
	}
	tests := []struct {
		name    string
		args    args
		want    []domain.Feira
		wantErr bool
	}{
		{
			name:    "No query fields",
			args:    args{},
			wantErr: true,
		}, {
			name: "Distrito",
			args: args{
				distrito: "DISTRITO",
			},
			want: []domain.Feira{
				feiras[0],
			},
		}, {
			name: "Distrito e região",
			args: args{
				distrito: "DISTRITO",
				regiao5:  "REGIAO_52",
			},
			want: []domain.Feira{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := service.Query(tt.args.distrito, tt.args.regiao5, tt.args.nome_feira, tt.args.bairro)
			if (err != nil) != tt.wantErr {
				t.Errorf("FeiraService.Query() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FeiraService.Query() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFeiraService_Save_BulkSave(t *testing.T) {
	teardownSuite := setupSuite(t, false)
	defer teardownSuite(t)

	t.Run("Default", func(t *testing.T) {
		feiras := []*domain.Feira{
			{
				Registro: "REGISTRO",
				Bairro:   "BAIRRO",
				Distrito: "DISTRITO",
				Regiao5:  "REGIAO_5",
			},
			{
				Registro: "REGISTRO2",
				Bairro:   "BAIRRO2",
				Distrito: "DISTRITO2",
				Regiao5:  "REGIAO_52",
			},
		}
		err := service.BulkSave(feiras)
		if err != nil {
			t.Errorf("FeiraService.BulkSave() error = %v", err)
			return
		}
		feiras2, err := service.FindAll()
		if err == nil && len(feiras2) != len(feiras) {
			err = fmt.Errorf("Expected %d feiras: got %d", len(feiras), len(feiras2))
		}
		if err != nil {
			t.Errorf("FeiraService.BulkSave() error = %v", err)
		}

		feira3 := &domain.Feira{
			Registro: "REGISTRO3",
			Bairro:   "BAIRRO3",
			Distrito: "DISTRITO3",
			Regiao5:  "REGIAO_53",
		}
		err = service.Save(feira3)
		if err != nil {
			t.Errorf("FeiraService.Save() error = %v", err)
			return
		}
	})

}

func TestValidateFeira(t *testing.T) {
	type args struct {
		feira *domain.Feira
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid",
			args: args{
				feira: &domain.Feira{
					NomeFeira: "NOME_FEIRA",
					Registro:  "REGISTRO",
					Bairro:    "BAIRRO",
					Distrito:  "DISTRITO",
					Regiao5:   "REGIAO_5",
				},
			},
		},
		{
			name: "Missing NomeFeira",
			args: args{
				feira: &domain.Feira{
					Registro: "REGISTRO",
					Bairro:   "BAIRRO",
					Distrito: "DISTRITO",
					Regiao5:  "REGIAO_5",
				},
			},
			wantErr: true,
		},
		{
			name: "Missing Registro",
			args: args{
				feira: &domain.Feira{
					NomeFeira: "NOME_FEIRA",
					Bairro:    "BAIRRO",
					Distrito:  "DISTRITO",
					Regiao5:   "REGIAO_5",
				},
			},
			wantErr: true,
		},

		{
			name: "Missing Bairro",
			args: args{
				feira: &domain.Feira{
					NomeFeira: "NOME_FEIRA",
					Registro:  "REGISTRO",
					Distrito:  "DISTRITO",
					Regiao5:   "REGIAO_5",
				},
			},
			wantErr: true,
		},
		{
			name: "Missing Distrito",
			args: args{
				feira: &domain.Feira{
					NomeFeira: "NOME_FEIRA",
					Registro:  "REGISTRO",
					Bairro:    "BAIRRO",
					Regiao5:   "REGIAO_5",
				},
			},
			wantErr: true,
		},
		{
			name: "Missing Regiao5",
			args: args{
				feira: &domain.Feira{
					NomeFeira: "NOME_FEIRA",
					Registro:  "REGISTRO",
					Bairro:    "BAIRRO",
					Distrito:  "DISTRITO",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateFeira(tt.args.feira); (err != nil) != tt.wantErr {
				t.Errorf("ValidateFeira() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
