package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["apiproject/controllers:Basecontroller"] = append(beego.GlobalControllerRouter["apiproject/controllers:Basecontroller"],
        beego.ControllerComments{
            Method: "ChannelRegion",
            Router: "/channel/region",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apiproject/controllers:Basecontroller"] = append(beego.GlobalControllerRouter["apiproject/controllers:Basecontroller"],
        beego.ControllerComments{
            Method: "ChannelType",
            Router: "/channel/type",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apiproject/controllers:CommentController"] = append(beego.GlobalControllerRouter["apiproject/controllers:CommentController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/comment/list",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apiproject/controllers:CommentController"] = append(beego.GlobalControllerRouter["apiproject/controllers:CommentController"],
        beego.ControllerComments{
            Method: "Save",
            Router: "/comment/save",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apiproject/controllers:TopController"] = append(beego.GlobalControllerRouter["apiproject/controllers:TopController"],
        beego.ControllerComments{
            Method: "ChannelTop",
            Router: "/channel/top",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apiproject/controllers:TopController"] = append(beego.GlobalControllerRouter["apiproject/controllers:TopController"],
        beego.ControllerComments{
            Method: "TypeTop",
            Router: "/type/top",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apiproject/controllers:UserController"] = append(beego.GlobalControllerRouter["apiproject/controllers:UserController"],
        beego.ControllerComments{
            Method: "LoginDo",
            Router: "/login/do",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apiproject/controllers:UserController"] = append(beego.GlobalControllerRouter["apiproject/controllers:UserController"],
        beego.ControllerComments{
            Method: "SaveRegister",
            Router: "/register/save",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apiproject/controllers:UserController"] = append(beego.GlobalControllerRouter["apiproject/controllers:UserController"],
        beego.ControllerComments{
            Method: "SendMessageDo",
            Router: "/send/message",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apiproject/controllers:VedioController"] = append(beego.GlobalControllerRouter["apiproject/controllers:VedioController"],
        beego.ControllerComments{
            Method: "ChannelAdvert",
            Router: "/channel/advert",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apiproject/controllers:VedioController"] = append(beego.GlobalControllerRouter["apiproject/controllers:VedioController"],
        beego.ControllerComments{
            Method: "ChannelHotList",
            Router: "/channel/hot",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apiproject/controllers:VedioController"] = append(beego.GlobalControllerRouter["apiproject/controllers:VedioController"],
        beego.ControllerComments{
            Method: "ChannelRegionRecommendList",
            Router: "/channel/recommend/region",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apiproject/controllers:VedioController"] = append(beego.GlobalControllerRouter["apiproject/controllers:VedioController"],
        beego.ControllerComments{
            Method: "ChannelTypeRecommendList",
            Router: "/channel/recommend/type",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apiproject/controllers:VedioController"] = append(beego.GlobalControllerRouter["apiproject/controllers:VedioController"],
        beego.ControllerComments{
            Method: "ChannelVideo",
            Router: "/channel/video",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apiproject/controllers:VedioController"] = append(beego.GlobalControllerRouter["apiproject/controllers:VedioController"],
        beego.ControllerComments{
            Method: "VideoEpisodesList",
            Router: "/video/episodes/list",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apiproject/controllers:VedioController"] = append(beego.GlobalControllerRouter["apiproject/controllers:VedioController"],
        beego.ControllerComments{
            Method: "VideoInfo",
            Router: "/video/info",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
