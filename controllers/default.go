package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	m "github.com/xulei8/lifeapp/models"
	"strings"
	"time"
)

var mUid int
var mUidSer int

type MainController struct {
	beego.Controller
}

type AppSer struct {
	beego.Controller
}

func (c *AppSer) Get() {

	act := c.GetString("act")

	if act == "delete" {
		Deleteid, _ := c.GetInt("id", 0)
		mmd := m.AppMain{Id: Deleteid}
		mmd.Deleted = 1
		o := orm.NewOrm()
		o.Update(&mmd)
		c.Ctx.WriteString(`{"success":true }`)
		return
	}
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "app.html"
}

func (c *AppSer) Post() {
	o := orm.NewOrm()
	act := c.GetString("act")
	newid := ""
	if act == "save" {
		id, _ := c.GetInt("id", 0)
		mmd := m.AppMain{}
		if id > 0 {
			mmd.Id = id
			o.Read(&mmd)
		}
		mmd.Title = c.GetString("title")
		mmd.Modname = c.GetString("modname")
		mmd.Tags = c.GetString("tags")
		mmd.Content = c.GetString("content")
		mmd.Descwords = c.GetString("descwords")

		cid1, _ := c.GetInt("cateid", 0)
		cid2, _ := c.GetInt("cateid2", 0)
		oid, _ := c.GetInt("oid", 100)
		rid, _ := c.GetInt("relationid", 100)
		mmd.Relationid = rid
		mmd.Cateid = cid1
		mmd.Cateid2 = cid2
		mmd.Orderid = oid

		if id < 1 {
			mmd.Addtime = time.Now()
			o.Insert(&mmd)
			newid = fmt.Sprintf("%d", mmd.Id)
			fmt.Print("insert ", mmd)
			c.Ctx.WriteString(`{"success":true ,"id","` + newid + `"}`)
			return

		} else {

			mmd.Edittime = time.Now()
			o.Update(&mmd)
			c.Ctx.WriteString(`{"success":true  }`)
			return
		}

	}

	if act == "load" {
		mod := c.GetString("modname")
		if len(mod) < 2 {
			c.Ctx.WriteString(`{"error":true  }`)
			return
		}

		var datas []*m.AppMain
		dataset := o.QueryTable("app_main").Filter("modname", mod)
		cateid := c.GetString("cateid")
		cateid2 := c.GetString("cateid2")
		if len(cateid) > 0 {
			dataset = dataset.Filter("cateid", cateid)
		}
		if len(cateid2) > 0 {
			dataset = dataset.Filter("cateid2", cateid2)
		}

		dataset.All(&datas)
		fmt.Print(datas)
		i := 0
		var rows []string
		rown := len(datas)
		for i = 0; i < rown; i++ {
			strrow := fmt.Sprint(`{"id":"`, datas[i].Id, `","tile":"`+datas[i].Title+`" `)
			strrow += fmt.Sprint(`, "cateid":"`, datas[i].Cateid, `"`)
			strrow += fmt.Sprint(`, "cateid2":"`, datas[i].Cateid2, `"`)
			strrow += fmt.Sprint(`, "tags":"`, datas[i].Tags, `"`)
			strrow += fmt.Sprint(`, "modname":"`, datas[i].Modname, `"`)
			strrow += fmt.Sprint(`, "addtime":"`, datas[i].Addtime, `"`)
			strrow += fmt.Sprint(`, "edittime":"`, datas[i].Edittime, `"`)
			strrow += fmt.Sprint(`, "filename":"`, datas[i].Filename, `"`)
			strrow += "}"
			rows = append(rows, strrow)
		}
		datastr := "[" + strings.Join(rows, ",") + "]"

		c.Ctx.WriteString(`{"rows":` + fmt.Sprintf("%d", rown) + ` ,"data:"` + datastr + `}`)
	}

	if act == "loadone" {

		id, _ := c.GetInt("id", 0)
		if id < 1 {
			c.Ctx.WriteString(`{"error":true  }`)
			return
		}

		var datas []*m.AppMain
		o.QueryTable("app_main").Filter("id", id).All(&datas)
		fmt.Print(datas)

		rown := len(datas)
		if rown < 1 {
			c.Ctx.WriteString(`{"success":false   }`)
			return
		}
		strrow := ""

		strrow += fmt.Sprint(`{"id":"`, datas[0].Id, `"  `)
		strrow += fmt.Sprint(`, "filename":"`, datas[0].Filename, `"`)
		strrow += fmt.Sprint(`, "content":"`, datas[0].Content, `"`)
		strrow += fmt.Sprint(`,"success":true}`)
		c.Ctx.WriteString(strrow)
		return
	}

}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "index.tpl"
}

func init() {
	mUid = 1
	mUidSer = 1

}
