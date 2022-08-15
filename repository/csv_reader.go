package repository

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"encoding/csv"

	"github.com/guionardo/go-api-example/domain"
)

func ReadCsvFile(csvFile string) (feiras []*domain.Feira, err error) {
	file, err := os.Open(csvFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.FieldsPerRecord = -1
	r.TrimLeadingSpace = true
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	feiras = make([]*domain.Feira, len(records)-1)
	for i, record := range records {
		if i == 0 {
			continue //skip header
		}
		feira, err := NewFeira(record)
		if err != nil {
			log.Printf("Error on line %d: %s", i, err)
		}
		feiras[i-1] = feira
	}
	return feiras, nil
}

func NewFeira(fields []string) (feira *domain.Feira, err error) {
	if len(fields) < 16 {
		return nil, fmt.Errorf("invalid number of fields: %d", len(fields))
	}
	feira = &domain.Feira{}
	if err = strToInt(fields[0], "id", &feira.ID); err != nil {
		return
	}
	if err = strToInt(fields[1], "long", &feira.Long); err != nil {
		return
	}
	if err = strToInt(fields[2], "lat", &feira.Lat); err != nil {
		return
	}
	if err = strToInt(fields[3], "setcens", &feira.SetCens); err != nil {
		return
	}
	if err = strToInt(fields[4], "areap", &feira.AreaP); err != nil {
		return
	}
	if err = strToInt(fields[5], "coddist", &feira.CodDist); err != nil {
		return
	}
	feira.Distrito = fields[6]
	if err = strToInt(fields[7], "codsubpref", &feira.CodSubPref); err != nil {
		return
	}
	feira.SubPrefe = fields[8]
	feira.Regiao5 = fields[9]
	feira.Regiao8 = fields[10]
	feira.NomeFeira = fields[11]
	feira.Registro = fields[12]
	feira.Logradouro = fields[13]
	feira.Numero = fields[14]
	feira.Bairro = fields[15]
	if len(fields) > 16 {
		feira.Referencia = fields[16]
	}

	return
}

func strToInt(fieldValue string, fieldName string, field *int) error {
	value, err := strconv.Atoi(fieldValue)
	if err != nil {
		return fmt.Errorf("invalid %s: %s", fieldName, fieldValue)
	}
	*field = value
	return nil
}