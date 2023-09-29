package entity

const (
	ANALISE_TIPO_AUTOMATICA = AnaliseTipo("AU")
	ANALISE_TIPO_MANUAL     = AnaliseTipo("AM")
)

const (
	ANALISE_STATUS_ATIVA   = AnaliseStatus("A")
	ANALISE_STATUS_INATIVA = AnaliseStatus("I")
)

type Analise struct {
	Id             uint
	Uuid           string
	Wkt            string
	Usuario        *Usuario
	Ativo          *Ativo
	Status         AnaliseStatus
	Tipo           AnaliseTipo
	TotalArea      float32
	TotalDeteccoes uint
	ImagemAntes    *Imagem
	ImagemDepois   *Imagem
	Alertas        []*Alerta
	Camadas        []*AnaliseCamada
}

func NewAnalise(id uint) *Analise {
	return &Analise{
		Id:      id,
		Status:  ANALISE_STATUS_ATIVA,
		Tipo:    ANALISE_TIPO_AUTOMATICA,
		Alertas: []*Alerta{},
		Camadas: []*AnaliseCamada{},
	}
}

func (analise *Analise) PopuleCamadasAfetadas() {
	camadasAfetadas := make(map[uint]*AnaliseCamada)

	for _, alerta := range analise.Alertas {
		for _, camada := range alerta.Deteccao.Camadas {
			if _, ok := camadasAfetadas[camada.Id]; ok {
				camadasAfetadas[camada.Id].Total++
			} else {
				camadasAfetadas[camada.Id] = NewAnaliseCamada(0, camada, 1)
			}
		}
	}

	for _, camada := range camadasAfetadas {
		analise.Camadas = append(analise.Camadas, camada)
	}
}
