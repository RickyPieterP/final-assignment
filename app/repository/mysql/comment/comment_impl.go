package comment

import (
	"fmt"
	"mygram/app/adapter/database"
	"mygram/app/models/mysql"
)

type commentRepo struct {
	mysql *database.MySQL
}

func NewCommentRepo(mysql *database.MySQL) RepositoryComment {
	return &commentRepo{
		mysql: mysql,
	}
}

func (p *commentRepo) Find(in int) (out []mysql.Comment, err error) {
	err = p.mysql.Where("user_id = ?", in).Find(&out).Error
	return
}

func (p *commentRepo) FindOne(in int) (out *mysql.Comment, err error) {
	out = &mysql.Comment{}
	fmt.Println(in, "in findone")
	err = p.mysql.Where("id = ?", &in).First(out).Error
	fmt.Println(err, "error di findOne")
	return
}

func (p *commentRepo) Create(in *mysql.Comment) (out *mysql.Comment, err error) {
	err = p.mysql.Create(in).Error
	if err != nil {
		return nil, err
	}
	return in, nil
}

func (p *commentRepo) Update(in *mysql.Comment) (out *mysql.Comment, err error) {
	out = &mysql.Comment{}
	err = p.mysql.Where("id = ?", in.Id).Updates(in).Error
	if err != nil {
		return nil, err
	}
	p.mysql.Where("id = ?", in.Id).First(out)
	return out, nil
}

func (p *commentRepo) Delete(in int) (out bool, err error) {
	row := p.mysql.Model(&mysql.Comment{}).Where("id = ?", in).Delete(&mysql.Comment{})
	if row.RowsAffected != 0 {
		out = true
		return out, nil
	} else {
		err = row.Error
		if err != nil {
			return false, err
		}
		return false, nil
	}
}
