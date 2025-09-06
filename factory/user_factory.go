package factory

// Repository es una interfaz genérica para repositorios de cualquier entidad.
type Repository interface {
	Exists(id string) (bool, error)
	Save(id string, data interface{}) error
	Get(id string) (interface{}, error)
}

// DatabaseFactory es una interfaz genérica para crear repositorios.
// Cada microservicio implementa esta interfaz con su DB específica.
type DatabaseFactory interface {
	CreateRepository(entity string) Repository
}

// NewService crea un servicio genérico usando un DatabaseFactory.
// entity: Tipo de entidad (ej. "user", "product").
func NewService(dbFactory DatabaseFactory, entity string) interface{} {
	repo := dbFactory.CreateRepository(entity)
	// Aquí podrías devolver un servicio genérico o específico
	// Por simplicidad, devolver el repo; ajusta según necesidad
	return repo
}
