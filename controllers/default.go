package controllers

import (
	"crypto/rand"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	m "github.com/xulei8/lifeapp/models"
	"io/ioutil"
	"math/big"
	"os"
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

type AppSerUE struct {
	beego.Controller
}

func (c *AppSerUE) Get() {
	action := c.GetString("action")
	if action == "listimage" {
		start, _ := c.GetInt("start", 0)
		size, _ := c.GetInt("size", 5)
		if size < 2 {
			size = 2
		}

		o := orm.NewOrm()
		var datas []*m.AppMain
		allc, _ := o.QueryTable("app_main").Filter("modname", "upimg").Count()

		o.QueryTable("app_main").Filter("modname", "upimg").Limit(size, start).All(&datas)

		rown := len(datas)
		var rows []string
		i := 0
		for i = 0; i < rown; i++ {
			strrow := fmt.Sprint(`{"url":"` + datas[i].Filename + `","mtime":1400203383}`)
			rows = append(rows, strrow)
		}
		datastr := "[" + strings.Join(rows, ",") + "]"
		rstr := `{"state":"SUCCESS","start":"` + fmt.Sprint(start) + `", "total":` + fmt.Sprint(allc) + `,"list":` + datastr + `}`
		c.Ctx.WriteString(rstr)
		return
	}

	if action == "listfile" {
		start, _ := c.GetInt("start", 0)
		size, _ := c.GetInt("size", 5)
		if size < 2 {
			size = 2
		}

		o := orm.NewOrm()
		var datas []*m.AppMain
		allc, _ := o.QueryTable("app_main").Filter("modname", "files").Count()

		o.QueryTable("app_main").Filter("modname", "files").Limit(size, start).All(&datas)

		rown := len(datas)
		var rows []string
		i := 0
		for i = 0; i < rown; i++ {
			strrow := fmt.Sprint(`{"url":"` + datas[i].Filename + `","mtime":1400203383}`)
			rows = append(rows, strrow)
		}
		datastr := "[" + strings.Join(rows, ",") + "]"
		rstr := `{"state":"SUCCESS","start":"` + fmt.Sprint(start) + `", "total":` + fmt.Sprint(allc) + `,"list":` + datastr + `}`
		c.Ctx.WriteString(rstr)
		return
	}

	if action == "config" {
		fi, err := os.Open("static/bdeditor/php/config.json")
		if err != nil {
			panic(err)
			return
		}
		defer fi.Close()
		str, _ := ioutil.ReadAll(fi)

		c.Ctx.WriteString(string(str))
		return
	}
}

func (c *AppSerUE) Post() {
	if c.GetString("action") == "uploadimage" {
		_, h, err := c.GetFile("upfile")
		if err != nil {
			fmt.Println("getfile err ", err)
		}

		filename := h.Filename[strings.LastIndex(h.Filename, `:`)+1:]
		filetype := h.Filename[strings.LastIndex(h.Filename, `.`):]
		max := big.NewInt(100000)

		randi, _ := rand.Int(rand.Reader, max)

		tn := time.Now()
		timestr := fmt.Sprint(strings.TrimSpace(fmt.Sprint(tn.Year())), strings.TrimSpace(tn.Month().String()))
		timestr += strings.TrimSpace(fmt.Sprint(tn.Day()))
		timestr += strings.TrimSpace(fmt.Sprint(tn.Hour()))
		timestr += strings.TrimSpace(fmt.Sprint(tn.Minute())) + "_"
		path := `static/up/` + timestr + fmt.Sprint(randi) + filetype

		c.SaveToFile("upfile", path)
		o := orm.NewOrm()
		mmd := m.AppMain{}
		mmd.Title = filename
		mmd.Addtime = time.Now()
		mmd.Filename = "/" + path
		mmd.Modname = "upimg"
		o.Insert(&mmd)

		c.Ctx.WriteString(`{"original":"` + filename + `","name":"` + filename + `","url":"/` + path + `","size":"123123","type":"` + filetype + `","state":"SUCCESS"}`)
		return
	}

	if c.GetString("action") == "uploadfile" {
		_, h, err := c.GetFile("upfile")
		if err != nil {
			fmt.Println("getfile err ", err)
		}

		filename := h.Filename[strings.LastIndex(h.Filename, `:`)+1:]
		filetype := h.Filename[strings.LastIndex(h.Filename, `.`):]
		max := big.NewInt(100000)

		randi, _ := rand.Int(rand.Reader, max)

		tn := time.Now()
		timestr := fmt.Sprint(strings.TrimSpace(fmt.Sprint(tn.Year())), strings.TrimSpace(tn.Month().String()))
		timestr += strings.TrimSpace(fmt.Sprint(tn.Day()))
		timestr += strings.TrimSpace(fmt.Sprint(tn.Hour()))
		timestr += strings.TrimSpace(fmt.Sprint(tn.Minute())) + "_"
		nname := timestr + fmt.Sprint(randi) + filetype
		path := `static/up/` + nname

		c.SaveToFile("upfile", path)

		o := orm.NewOrm()
		mmd := m.AppMain{}
		mmd.Title = filename
		mmd.Addtime = time.Now()
		mmd.Filename = "/" + path
		mmd.Modname = "files"
		o.Insert(&mmd)

		c.Ctx.WriteString(`{"original":"` + filename + `","name":"` + nname + `","url":"/` + path + `","size":"99697","type":"` + filetype + `","state":"SUCCESS"}`)
		//c.Ctx.WriteString(`{"original":"demo.jpg","name":"demo.jpg","url":"/` + path + `","size":"99697","type:"` + filetype + `","state":"SUCCESS"}`)
		return
	}

	return

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
			c.Ctx.WriteString(`{"success":true ,"id":"` + newid + `"}`)
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
			strrow := fmt.Sprint(`{"id":"`, datas[i].Id, `","title":"`+datas[i].Title+`" `)
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

		c.Ctx.WriteString(`{"rows":` + fmt.Sprintf("%d", rown) + ` ,"data":` + datastr + `}`)
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
		contentstr := strings.Replace(datas[0].Content, "\"", `\"`, -1)
		strrow += fmt.Sprint(`{"id":"`, datas[0].Id, `"  `)
		strrow += fmt.Sprint(`, "filename":"`, datas[0].Filename, `"`)
		strrow += fmt.Sprint(`, "content":"`, contentstr, `"`)
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
