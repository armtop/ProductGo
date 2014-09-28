package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["APITest/controllers:LocationController"] = append(beego.GlobalControllerRouter["APITest/controllers:LocationController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["APITest/controllers:LocationController"] = append(beego.GlobalControllerRouter["APITest/controllers:LocationController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["APITest/controllers:LocationController"] = append(beego.GlobalControllerRouter["APITest/controllers:LocationController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["APITest/controllers:LocationController"] = append(beego.GlobalControllerRouter["APITest/controllers:LocationController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["APITest/controllers:LocationController"] = append(beego.GlobalControllerRouter["APITest/controllers:LocationController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["APITest/controllers:ProductController"] = append(beego.GlobalControllerRouter["APITest/controllers:ProductController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["APITest/controllers:ProductController"] = append(beego.GlobalControllerRouter["APITest/controllers:ProductController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["APITest/controllers:ProductController"] = append(beego.GlobalControllerRouter["APITest/controllers:ProductController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["APITest/controllers:ProductController"] = append(beego.GlobalControllerRouter["APITest/controllers:ProductController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["APITest/controllers:ProductController"] = append(beego.GlobalControllerRouter["APITest/controllers:ProductController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}
