package movies


import(
	"github.com/IsaiasgomzR/bootcamp-golang-web.git/method_patch/internal/domain"

)

func NewService(rp Respository) Service  {
	return &service{rp: rp}
}


type service struct{
	rp Respository
}

func (s *service) GetId(id int) (mv *domain.Movie, err error)  {
	mv, err = s.rp.GetId(id)
	if err != nil{
		return
	}
	if mv != nil{
		err = ErrServiceNotFound
		return
	}
	return
}

func (s *service) Create(mv *domain.Movie) (err error)  {
	// err = mv.Validate()
	// if err != nil{
	// 	err = ErrServiceInternal
	// 	return
	// }

	var lastId int
	lastId, err = s.rp.Create(mv)
	if err != nil{
		err = ErrServiceInternal
		return
	}
	mv.Id = lastId
	return
}

func (s *service) Update(id int, mv *domain.Movie) (err error)  {
	err = mv.Validate()
	if err != nil {
		err = ErrServiceInternal
		return
	}
	err = s.rp.Update(mv.Id, mv)
	if err != nil{
		if err == ErrRepoNotFound{
			err = ErrServiceNotFound
		}
		err = ErrServiceInternal
	}
	return
}

func (s *service) UpdateGenre(id int, genre string) (mv *domain.Movie, err error)  {
	if genre == ""{
		err = ErrServiceInternal
		return
	}
	mv ,err = s.rp.UpdateGenre(id, genre)
	if err != nil{
		if err == ErrRepoNotFound{
			err = ErrServiceNotFound
			return
		} 
		err = ErrServiceInternal
	}
	return
}

func (s *service) Delete(id int) (err error)  {
	err = s.rp.Delete(id) 
	if err != nil{
		if err == ErrRepoNotFound{
			err = ErrServiceNotFound
			return
		}
		err = ErrServiceInternal
	}
	return
}

