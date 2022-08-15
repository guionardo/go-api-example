package infra

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/guionardo/go-api-example/domain"
	"gorm.io/gorm"
)

type FeiraService struct {
	lock sync.Mutex
	db   *gorm.DB
}

func NewFeiraService(config *Config) (service *FeiraService, err error) {
	service = &FeiraService{}
	service.db, err = GetDatabase(config)
	return
}

func (service *FeiraService) Reset() error {
	service.lock.Lock()
	defer service.lock.Unlock()
	return ResetDatabase(service.db)
}

func (service *FeiraService) BulkSave(feiras []*domain.Feira) error {
	service.lock.Lock()
	defer service.lock.Unlock()
	return service.db.CreateInBatches(feiras, 100).Error
}

func (service *FeiraService) Save(feira *domain.Feira) error {
	service.lock.Lock()
	defer service.lock.Unlock()
	return service.db.Save(feira).Error
}
func (service *FeiraService) Create(feira *domain.Feira) error {
	if err := ValidateFeira(feira); err != nil {
		return err
	}
	feira.ID = 0
	service.lock.Lock()
	defer service.lock.Unlock()
	return service.db.Create(feira).Error
}

func (service *FeiraService) Update(feira *domain.Feira) error {
	if err := ValidateFeira(feira); err != nil {
		return err
	}

	existingFeira, err := service.FindByRegistro(feira.Registro)
	if err != nil {
		return err
	}
	service.lock.Lock()
	defer service.lock.Unlock()
	feira.ID = existingFeira.ID
	return service.db.Save(feira).Error
}

func ValidateFeira(feira *domain.Feira) error {
	if feira == nil {
		return fmt.Errorf("feira is nil")
	}
	if len(feira.Registro) == 0 {
		return fmt.Errorf("registro is required")
	}
	if len(feira.NomeFeira) == 0 {
		return fmt.Errorf("nome_feira is required")
	}
	if len(feira.Distrito) == 0 {
		return fmt.Errorf("distrito is required")
	}
	if len(feira.Regiao5) == 0 {
		return fmt.Errorf("regiao5 is required")
	}
	if len(feira.Bairro) == 0 {
		return fmt.Errorf("bairro is required")
	}
	return nil
}

func (service *FeiraService) FindAll() ([]domain.Feira, error) {
	service.lock.Lock()
	defer service.lock.Unlock()
	var feiras []domain.Feira
	if err := service.db.Find(&feiras).Error; err != nil {
		return nil, err
	}
	return feiras, nil
}

func (service *FeiraService) FindByID(id int) (*domain.Feira, error) {
	service.lock.Lock()
	defer service.lock.Unlock()
	var feira domain.Feira
	if err := service.db.First(&feira, id).Error; err != nil {
		return nil, err
	}
	return &feira, nil
}

func (service *FeiraService) FindByRegistro(registro string) (*domain.Feira, error) {
	service.lock.Lock()
	defer service.lock.Unlock()
	var feira domain.Feira
	if err := service.db.First(&feira, "registro = ?", registro).Error; err != nil {
		return nil, err
	}
	return &feira, nil
}

func (service *FeiraService) Delete(feira *domain.Feira) error {
	service.lock.Lock()
	defer service.lock.Unlock()
	return service.db.Delete(feira).Error
}

func (repository *FeiraService) DeleteByRegistro(registro string) error {
	repository.lock.Lock()
	defer repository.lock.Unlock()
	var feira domain.Feira
	if err := repository.db.First(&feira, "registro = ?", registro).Error; err != nil {
		return err
	}
	return repository.db.Delete(feira).Error
}
func (repository *FeiraService) Query(distrito string, regiao5 string, nome_feira string, bairro string) ([]domain.Feira, error) {

	var feiras []domain.Feira
	queryFields := make([]string, 0)
	queryArgs := make([]interface{}, 0)
	if distrito != "" {
		queryFields = append(queryFields, "distrito like ?")
		queryArgs = append(queryArgs, distrito)
	}
	if regiao5 != "" {
		queryFields = append(queryFields, "regiao5 like ?")
		queryArgs = append(queryArgs, regiao5)
	}
	if nome_feira != "" {
		queryFields = append(queryFields, "nome_feira like ?")
		queryArgs = append(queryArgs, nome_feira)
	}
	if bairro != "" {
		queryFields = append(queryFields, "bairro like ?")
		queryArgs = append(queryArgs, bairro)
	}
	if len(queryFields) == 0 {
		//TODO: Trazer todas as feiras se não houver filtro?
		return nil, fmt.Errorf("no query fields provided (expected 'distrito', 'regiao5', 'nome_feira', or 'bairro')")
	}
	repository.lock.Lock()
	defer repository.lock.Unlock()
	queryString := strings.Join(queryFields, " and ")
	if err := repository.db.Where(queryString, queryArgs...).Find(&feiras).Error; err != nil {
		return nil, err
	}
	log.Printf("found %d feiras", len(feiras))
	return feiras, nil
}
