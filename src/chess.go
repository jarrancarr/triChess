package main

import (
	"fmt"
	"net/http"
	"html/template"

	"github.com/jarrancarr/website"
	"github.com/jarrancarr/website/service"
)

var triChess *website.Site

func main() {
	website.ResourceDir = ".."
	setup()

	http.HandleFunc("/js/", website.ServeResource)
	http.HandleFunc("/css/", website.ServeResource)
	http.HandleFunc("/img/", website.ServeResource)
	http.ListenAndServe(":8070", nil)
}

func setup() {
	//website
	triChess = website.CreateSite("chess", "localhost:8070")
	triChess.AddMenu("nav").
		AddItem("My Games", "/games").
		AddItem("Settings", "/settings").
		AddItem("New Game", "/newGame").
		AddItem("Teams", "/teams").
		AddItem("Clubs", "/clubs").
		AddItem("Message", "/message").
		AddItem("Login", "/login").
		Add("nav nav-pills nav-stacked", "", "")

	// services
	acs := website.CreateAccountService()
	triChess.AddService("account", acs)
	mgs := service.CreateMessageService(acs)
	triChess.AddService("message", mgs)

	// template subpages
	triChess.AddPage("", "head", "")
	triChess.AddPage("", "banner", "")

	// pages
	chess := triChess.AddPage("chess", "chess", "/")
	
	scaleX := 30
	scaleY := 15
	offX := 120
	offY := 0
	spaces := 4
	perspective := 2
	chess.Data["triChessBoard"] = append(chess.Data["triChessBoard"], 
		template.HTML(fmt.Sprintf("<circle cx='%d' cy='%d' r='%d' stroke='black' stroke-width='2' fill='#248'/>",
			offX+(2*scaleX+scaleY)*spaces,offY+scaleY*spaces, (2*scaleX+scaleY)*spaces)))
	for y := 0; y<spaces; y++ {
		scaleY += perspective
		for x := 0; x<spaces+1+y; x++ {
			if x>0 {				
				chess.Data["triChessBoard"] = append(chess.Data["triChessBoard"], 
					triangle(offX,offY,scaleX,scaleY,perspective,
					2*x-y,2*y,2*x-y+1,2*y+2,2*x-y+2,2*y,"downTri",0))
			}
			chess.Data["triChessBoard"] = append(chess.Data["triChessBoard"], 
				triangle(offX,offY,scaleX,scaleY,perspective,
				2*x-y+1,2*y+2,2*x-y+2,2*y,2*x-y+3,2*y+2,"upTri",1))
		}
	}
	for y := spaces; y<spaces*2; y++ {
		scaleY += perspective
		for x := 0; x<spaces*3-y; x++ {
			if x>0 {		
				chess.Data["triChessBoard"] = append(chess.Data["triChessBoard"], 
					triangle(offX,offY,scaleX,scaleY,perspective,
					2*x+y+3-spaces*2,2*y+2,2*x+y+2-spaces*2,2*y,2*x+y+1-spaces*2,2*y+2,"upTri",1))
			}	
			chess.Data["triChessBoard"] = append(chess.Data["triChessBoard"], 
				triangle(offX,offY,scaleX,scaleY,perspective,
				2*x+y+2-spaces*2,2*y,2*x+y+3-spaces*2,2*y+2,2*x+y+4-spaces*2,2*y,"downTri",0))			
		}
	}
}

func triangle(offX,offY,scaleX,scaleY,perspective,px1,py1,px2,py2,px3,py3 int, pClass string, up int) template.HTML {
	return template.HTML(fmt.Sprintf(
		"<polygon class='%s' points='%d,%d %d,%d %d,%d' />",pClass,
			offX+px1*(scaleX+scaleY+up*perspective), offY+py1*(scaleY+up*perspective), 
			offX+px2*(scaleX+scaleY+(1-up)*perspective), offY+py2*(scaleY+(1-up)*perspective), 
			offX+px3*(scaleX+scaleY+up*perspective), offY+py3*(scaleY+up*perspective)))
}