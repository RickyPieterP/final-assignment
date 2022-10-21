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

func (c *commentRepo) Find(in int) (out []mysql.Comment, err error) {
	err = c.mysql.Where("user_id = ?", in).Find(&out).Error
	return
}

func (c *commentRepo) FindOne(in int) (out *mysql.Comment, err error) {
	out = &mysql.Comment{}
	err = c.mysql.Model(out).Where("id = ?", &in).First(out).Error
	fmt.Println(err, "error di findOne comment")
	return
}

func (c *commentRepo) Create(in *mysql.Comment) (out *mysql.Comment, err error) {
	err = c.mysql.Model(in).Create(in).Error
	if err != nil {

		return nil, err
	}
	return in, nil
}

func (c *commentRepo) Update(in *mysql.Comment) (out *mysql.Comment, err error) {
	out = &mysql.Comment{}
	err = c.mysql.Model(in).Where("id = ?", in.Id).Updates(in).Error
	if err != nil {
		return nil, err
	}
	c.mysql.Where("id = ?", in.Id).First(out)
	return out, nil
}

func (c *commentRepo) Delete(in int) (out bool, err error) {
	row := c.mysql.Model(&mysql.Comment{}).Where("id = ?", in).Delete(&mysql.Comment{})
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
