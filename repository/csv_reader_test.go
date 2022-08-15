package repository

import (
	"reflect"
	"testing"

	"github.com/guionardo/go-api-example/domain"
)

func TestReadCsvFile(t *testing.T) {

		t.Run("Read", func(t *testing.T) {
			gotFeiras, err := ReadCsvFile("DEINFO_AB_FEIRASLIVRES_2014.csv")
			if err != nil {
				t.Errorf("ReadCsvFile() error = %v", err)
				return
			}
			if len(gotFeiras) != 880 {
				t.Errorf("ReadCsvFile() len(gotFeiras) = %d, want %d", len(gotFeiras), 880)
			}
		})

}

func TestNewFeira(t *testing.T) {
	type args struct {
		fields []string
	}
	tests := []struct {
		name      string
		args      args
		wantFeira *domain.Feira
		wantErr   bool
	}{
		{
			name: "invalid number of fields",
			args: args{
				fields: []string{"1"},
			},
			wantErr: true,
		},
		{
			name: "valid fields",
			args: args{
				fields: []string{"1","-46550164","-23558733","355030885000091","3550308005040","87","VILA FORMOSA","26","ARICANDUVA-FORMOSA-CARRAO","Leste","Leste 1","VILA FORMOSA","4041-0","RUA MARAGOJIPE","S/N","VL FORMOSA","TV RUA PRETORIA"},				
			},
			wantErr: false,
			wantFeira: &domain.Feira{
				ID:         1,
				Long:       -46550164,
				Lat:        -23558733,
				SetCens:    355030885000091,
				AreaP:      3550308005040,
				CodDist:    87,
				Distrito:   "VILA FORMOSA",
				CodSubPref: 26,
				SubPrefe:   "ARICANDUVA-FORMOSA-CARRAO",
				Regiao5:    "Leste",
				Regiao8:    "Leste 1",
				NomeFeira:  "VILA FORMOSA",
				Registro:   "4041-0",
				Logradouro: "RUA MARAGOJIPE",
				Numero:     "S/N",
				Bairro:     "VL FORMOSA",
				Referencia: "TV RUA PRETORIA",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFeira, err := NewFeira(tt.args.fields)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFeira() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFeira, tt.wantFeira) {
				t.Errorf("NewFeira() = %v, want %v", gotFeira, tt.wantFeira)
			}
		})
	}
}