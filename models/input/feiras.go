package input

// Input validation for FEIRAS data creation
type CreateFeirasInput struct {
	LONG       float64 `json:"long" binding:"required"`
	LAT        float64 `json:"lat" binding:"required"`
	SETCENS    int64   `json:"setcens" binding:"required"`
	AREAP      int64   `json:"areap" binding:"required"`
	CODDIST    int     `json:"coddist" binding:"required"`
	DISTRITO   string  `json:"distrito" binding:"required"`
	CODSUBPREF int     `json:"codsubpref" binding:"required"`
	SUBPREF    string  `json:"subpref" binding:"required"`
	REGIAO5    string  `json:"regiao5" binding:"required"`
	REGIAO8    string  `json:"regiao8" binding:"required"`
	NOME_FEIRA string  `json:"nome_feira" binding:"required"`
	REGISTRO   string  `json:"registro" binding:"required"`
	LOGRADOURO string  `json:"logradouro" binding:"required"`
	NUMERO     string  `json:"numero"`
	BAIRRO     string  `json:"bairro" binding:"required"`
	REFERENCIA string  `json:"referencia"`
}

// Input validation for FEIRAS data update
type UpdateFeirasInput struct {
	LONG       float64 `json:"long"`
	LAT        float64 `json:"lat"`
	SETCENS    int64   `json:"setcens"`
	AREAP      int64   `json:"areap"`
	CODDIST    int     `json:"coddist"`
	DISTRITO   string  `json:"distrito"`
	CODSUBPREF int     `json:"codsubpref"`
	SUBPREF    string  `json:"subpref"`
	REGIAO5    string  `json:"regiao5"`
	REGIAO8    string  `json:"regiao8"`
	NOME_FEIRA string  `json:"nome_feira"`
	REGISTRO   string  `json:"registro"`
	LOGRADOURO string  `json:"logradouro"`
	NUMERO     string  `json:"numero"`
	BAIRRO     string  `json:"bairro"`
	REFERENCIA string  `json:"referencia"`
}
