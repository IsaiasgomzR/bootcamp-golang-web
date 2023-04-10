package movies

import (
	"github.com/IsaiasgomzR/bootcamp-golang-web.git/method_patch/internal/domain"

)

//constructor
func NewRepositoryLocal(db []*domain.Movie, lastId int) Respository {
	return &respositoryLocal{db: db ,lastId:lastId}
}


//controller
type respositoryLocal struct{
	db [] *domain.Movie
	lastId int
}

func (rp *respositoryLocal) GetId(id int) (mv *domain.Movie, err error)  {
	for _, m := range rp.db {
		if m.Id == id {
			mv = m
			return
		}
	}
	err = ErrRepoNotFound
	return
}

func (rp *respositoryLocal) Create(mv *domain.Movie) (lastId int , err error)  {
	rp.lastId++
	mv.Id = rp.lastId

	rp.db = append(rp.db, mv)

	lastId = rp.lastId
	return
}

func (rp *respositoryLocal) Update(id int, mv *domain.Movie) (err error)  {
	mv.Id = id
	for i, mv := range rp.db {
		if mv.Id == id {
			rp.db[i] = mv
			return
		}
	}
	err = ErrRepoNotFound
	return
}

func (rp *respositoryLocal) UpdateGenre( id int, genre string) (mv *domain.Movie, err error)  {
	for i, m := range rp.db {
		if m.Id == id {
			rp.db[i].Genre = genre
			mv = rp.db[i]
			return
		}
	}
	err = ErrRepoNotFound
	return
}

func (rp *respositoryLocal) Delete(id int) (err error){
	for i, m := range rp.db {
		if m.Id == id {
			rp.db = append(rp.db[:i],rp.db[i +1 :]... )
			return
		}
	}
	err = ErrRepoNotFound
	return
}