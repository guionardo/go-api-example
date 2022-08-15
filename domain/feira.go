package domain

//NOCOVER - REMOVES COVERAGE FOR THIS FILE

type (
	Feira struct {
		ID         int    `json:"id" gorm:"primary_key"`
		Registro   string `json:"registro" gorm:"index:idx_registro,unique;<-:create"`
		Long       int    `json:"long"`
		Lat        int    `json:"lat"`
		SetCens    int    `json:"set_cens"`
		AreaP      int    `json:"area_p"`
		CodDist    int    `json:"cod_dist"`
		Distrito   string `json:"distrito"`
		CodSubPref int    `json:"cod_sub_pref"`
		SubPrefe   string `json:"sub_prefe"`
		Regiao5    string `json:"regiao_5"`
		Regiao8    string `json:"regiao_8"`
		NomeFeira  string `json:"nome_feira"`
		Logradouro string `json:"logradouro"`
		Numero     string `json:"numero"`
		Bairro     string `json:"bairro"`
		Referencia string `json:"referencia"`
		//	ID,LONG,LAT,SETCENS,AREAP,CODDIST,DISTRITO,CODSUBPREF,SUBPREFE,REGIAO5,REGIAO8,NOME_FEIRA,REGISTRO,LOGRADOURO,NUMERO,BAIRRO,REFERENCIA
		//
		// 1,-46550164,-23558733,355030885000091,3550308005040,87,VILA FORMOSA,26,ARICANDUVA-FORMOSA-CARRAO,Leste,Leste 1,VILA FORMOSA,4041-0,RUA MARAGOJIPE,S/N,VL FORMOSA,TV RUA PRETORIA
	}

	FeiraService interface {
		Reset() error
		FindAll() ([]Feira, error)
		BulkSave(feiras []*Feira) error
		Save(feira *Feira) error
		Create(feira *Feira) error
		Update(feira *Feira) error
		FindByID(id int) (*Feira, error)
		FindByRegistro(registro string) (*Feira, error)
		DeleteByRegistro(registro string) error
		Query(distrito string, regiao5 string, nome_feira string, bairro string) ([]Feira, error)
	}
)
