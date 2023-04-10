package movies
 
import(
	"errors"
	"github.com/IsaiasgomzR/bootcamp-golang-web.git/method_patch/internal/domain"
)

type Respository interface{
	GetId(id int) (mv *domain.Movie, err error)
	Create(mv *domain.Movie) (lastId int,err error)
	Update(id int, mv *domain.Movie) (err error)
	UpdateGenre(in int, genre string) (mv *domain.Movie, err error)
	Delete(id int) (err error)
}


var(
	ErrRepoInternal =errors.New("internal error")
	ErrRepoNotUnique = errors.New("movie already exists")
	ErrRepoNotFound = errors.New("movie not found")
)

type Service interface{
	GetId(id int) (mv *domain.Movie, err error)
	Create(mv *domain.Movie) (err error)
	Update(id int, mv *domain.Movie) (err error)
	UpdateGenre(id int, genre string) (mv *domain.Movie, err error)
	Delete(id int) (err error)
}

var(
	ErrServiceInternal = errors.New("internal error")
	ErrServiceInvalid = errors.New("invalid movie")
	ErrServiceNotUnique = errors.New("movie already exists")
	ErrServiceNotFound = errors.New("movie not found")
)