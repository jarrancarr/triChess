package main

import (
	"fmt"
	"net/http"
	"html/template"

	"github.com/jarrancarr/website"
	"github.com/jarrancarr/website/service"
	"github.com/jarrancarr/website/html"
)

var triChess *website.Site

type Board struct {
	piece []Piecer
}

type Piecer interface {
	Moves(b *Board) []Move
	GetPiece() *Piece
}
type Piece struct {
	rank, file int
	team bool
}
func (p Piece) GetRank() int { return p.rank }
func (p Piece) GetFile() int { return p.file }
func (p Piece) GetTeam() bool { return p.team }
type Pawn struct {
	p *Piece
	moved bool
}
type Rook struct {
	p *Piece
	moved bool
}
type Knight struct {
	p *Piece
}
type Bishop struct {
	p *Piece
}
type Cannon struct {
	p *Piece
}
type Queen struct {
	p *Piece
}
type King struct {
	p *Piece
	moved bool
}
type Move struct {
	rank, file int
}

func checkMove(b *Board, p *Piece, r, f int, moves []Move) bool {
	who := b.getPieceAt(r,f)
	if who == nil {
		moves = append(moves,Move{r,f})
		return true
	}
	return false
}
func checkAttack(b *Board, p *Piece, r, f int, moves []Move) bool {
	who := b.getPieceAt(p.GetRank()+r,p.GetFile()+f)
	if who != nil && who.GetPiece().GetTeam() == !p.GetTeam() {
		moves = append(moves,Move{p.GetRank()+r,p.GetFile()+f})
		return true
	}
	return false
}
func checkMoveAttack(b *Board, p *Piece, r, f int, moves []Move) {
	checkMove(b, p, r, f, moves)
	checkAttack(b, p, r, f, moves)
}
func (p Pawn) Moves(b *Board) []Move {
	moves := make([]Move,0)
	dir := 1
	if p.p.team { dir = -1 }
	checkMove(b, p.p, dir,0, moves)
	if (!p.moved) { checkMove(b, p.p, 2*dir,0, moves) }
	checkAttack(b, p.p, dir, -1, moves)
	checkAttack(b, p.p, dir, 1, moves)
	return moves
}
func (p Pawn) GetPiece() *Piece { return p.p }
func (r Rook) Moves(b *Board) []Move {
	moves := make([]Move,0)
	rank := 0
	for rank = 1; checkMove(b, r.p, rank, 0, moves); rank += 1 {}
	checkAttack(b,r.p,rank, 0,moves)
	for rank = -1; checkMove(b , r.p, rank, 0, moves); rank -= 1 {}
	checkAttack(b,r.p,rank, 0,moves)
	file := 0
	for file = 1; checkMove(b, r.p, 0, file, moves); file += 1 {}
	checkAttack(b, r.p, 0, file ,moves)
	for file = -1; checkMove(b, r.p, 0, file, moves); file -= 1 {}
	checkAttack(b, r.p, 0, file, moves)
	return moves
}
func (r Rook) GetPiece() *Piece { return r.p }
func (n Knight) Moves(b *Board) []Move {
	moves := make([]Move,0)
	dir := ((n.p.rank + n.p.file) % 2 )*2 - 1
	checkMoveAttack(b, n.p, 2*dir, -1, moves)
	checkMoveAttack(b, n.p, 2*dir, 1, moves)
	checkMoveAttack(b, n.p, 1*dir, -2, moves)
	checkMoveAttack(b, n.p, 1*dir, 2, moves)
	checkMoveAttack(b, n.p, 0, -3, moves)
	checkMoveAttack(b, n.p, 0, 3, moves)
	checkMoveAttack(b, n.p, -1*dir, -4, moves)
	checkMoveAttack(b, n.p, -1*dir, 4, moves)
	checkMoveAttack(b, n.p, -2*dir, -3, moves)
	checkMoveAttack(b, n.p, -2*dir, -1, moves)
	checkMoveAttack(b, n.p, -2*dir, 1, moves)
	checkMoveAttack(b, n.p, -2*dir, 3, moves)
	return moves
}
func (n Knight) GetPiece() *Piece { return n.p }
func (B Bishop) Moves(b *Board) []Move {
	moves := make([]Move,0)
	rank := 0
	for rank = 1; checkMove(b, B.p, rank, 0, moves); rank += 1 {
		checkAttack(b,B.p,rank, 0,moves)
	}
	for rank = -1; checkMove(b , B.p, rank, 0, moves); rank -= 1 {
		checkAttack(b,B.p,rank, 0,moves)
	}
	file := 0
	for file = 1; checkMove(b, B.p, 0, file, moves); file += 1 {
		checkAttack(b, B.p, 0, file ,moves)
	}
	for file = -1; checkMove(b, B.p, 0, file, moves); file -= 1 {
		checkAttack(b, B.p, 0, file, moves)
	}
	return moves
}
func (B Bishop) GetPiece() *Piece { return B.p }
func (q Queen) Moves(b *Board) []Move {
	moves := make([]Move,0)
	rank := 0
	for rank = 1; checkMove(b, q.p, rank, 0, moves); rank += 1 {
		checkAttack(b,q.p,rank, 0,moves)	
	}
	for rank = -1; checkMove(b , q.p, rank, 0, moves); rank -= 1 {
		checkAttack(b,q.p,rank, 0,moves)
	}	
	file := 0
	for file = 1; checkMove(b, q.p, 0, file, moves); file += 1 {
		checkAttack(b, q.p, 0, file ,moves)
	}
	for file = -1; checkMove(b, q.p, 0, file, moves); file -= 1 {
		checkAttack(b, q.p, 0, file, moves)
	}
	for rank = 1; checkMove(b, q.p, rank, 0, moves); rank += 1 {
		checkAttack(b,q.p,rank, 0,moves)
	}
	for rank = -1; checkMove(b , q.p, rank, 0, moves); rank -= 1 {
		checkAttack(b,q.p,rank, 0,moves)
	}
	file = 0
	for file = 1; checkMove(b, q.p, 0, file, moves); file += 1 {
		checkAttack(b, q.p, 0, file ,moves)
	}
	for file = -1; checkMove(b, q.p, 0, file, moves); file -= 1 {
		checkAttack(b, q.p, 0, file, moves)
	}
	return moves
}
func (q Queen) GetPiece() *Piece { return q.p }
func (c Cannon) Moves(b *Board) []Move { // moves like a horse, attacks like the queen, but not on adjacent square
	moves := make([]Move,0)
	dir := ((c.p.rank + c.p.file) % 2 )*2 - 1
	checkMove(b, c.p, 2*dir, -1, moves)
	checkMove(b, c.p, 2*dir, 1, moves)
	checkMove(b, c.p, 1*dir, -2, moves)
	checkMove(b, c.p, 1*dir, 2, moves)
	checkMove(b, c.p, 0, -3, moves)
	checkMove(b, c.p, 0, 3, moves)
	checkMove(b, c.p, -1*dir, -4, moves)
	checkMove(b, c.p, -1*dir, 4, moves)
	checkMove(b, c.p, -2*dir, -3, moves)
	checkMove(b, c.p, -2*dir, -1, moves)
	checkMove(b, c.p, -2*dir, 1, moves)
	checkMove(b, c.p, -2*dir, 3, moves)
	rank := 0
	for rank = 2; b.getPieceAt(c.p.GetRank()+rank,c.p.GetFile())==nil; rank += 1 {
		checkAttack(b,c.p,rank, 0,moves)	
	}
	for rank = -2; b.getPieceAt(c.p.GetRank()+rank,c.p.GetFile())==nil; rank -= 1 {
		checkAttack(b,c.p,rank, 0,moves)
	}	
	file := 0
	for file = 2; b.getPieceAt(c.p.GetRank(),c.p.GetFile()+file)==nil; file += 1 {
		checkAttack(b, c.p, 0, file ,moves)
	}
	for file = -2; b.getPieceAt(c.p.GetRank(),c.p.GetFile()+file)==nil; file -= 1 {
		checkAttack(b, c.p, 0, file, moves)
	}
	for rank = 2; b.getPieceAt(c.p.GetRank()+file,c.p.GetFile()+file)==nil; rank += 1 {
		checkAttack(b,c.p,rank, 0,moves)
	}
	for rank = -2; b.getPieceAt(c.p.GetRank()+file,c.p.GetFile()+file)==nil; rank -= 1 {
		checkAttack(b,c.p,rank, 0,moves)
	}
	for rank = 2; b.getPieceAt(c.p.GetRank()+file,c.p.GetFile()-file)==nil; rank += 1 {
		checkAttack(b,c.p,rank, 0,moves)
	}
	for rank = -2; b.getPieceAt(c.p.GetRank()+file,c.p.GetFile()-file)==nil; rank -= 1 {
		checkAttack(b,c.p,rank, 0,moves)
	}
	return moves
}
func (c Cannon) GetPiece() *Piece { return c.p }

func (b Board) setup() {
	b.piece = make([]Piecer,0)
	b.piece = append(b.piece,Rook{&Piece{1,4,false}, false})
	b.piece = append(b.piece,Rook{&Piece{1,12,false}, false})
	b.piece = append(b.piece,Knight{&Piece{1,5,false}})
	b.piece = append(b.piece,Knight{&Piece{1,11,false}})
	b.piece = append(b.piece,Bishop{&Piece{1,6,false}})
	b.piece = append(b.piece,Bishop{&Piece{1,10,false}})
	b.piece = append(b.piece,Queen{&Piece{1,7,false}})
	b.piece = append(b.piece,Cannon{&Piece{1,8,false}})
	b.piece = append(b.piece,King{&Piece{1,9,false}, false})
	for i := 3 ; i < 13 ; i++ {
		b.piece = append(b.piece,Pawn{&Piece{2,i,false}, false})
		b.piece = append(b.piece,Pawn{&Piece{7,i,true}, false})
	}	
	b.piece = append(b.piece,Rook{&Piece{8,4,true}, false})
	b.piece = append(b.piece,Rook{&Piece{8,12,true}, false})
	b.piece = append(b.piece,Knight{&Piece{8,5,true}})
	b.piece = append(b.piece,Knight{&Piece{8,11,true}})
	b.piece = append(b.piece,Bishop{&Piece{8,6,true}})
	b.piece = append(b.piece,Bishop{&Piece{8,10,true}})
	b.piece = append(b.piece,Queen{&Piece{8,7,true}})
	b.piece = append(b.piece,Cannon{&Piece{8,8,true}})
	b.piece = append(b.piece,King{&Piece{8,9,true}, false})
}

func (b Board) getPieceAt(rank, file int) Piecer {
	for _,p := range(b.piece) {
		if p.GetPiece().rank == rank && p.GetPiece().file == file {
			return p
		}
	}
	return nil
}
func (k King) Moves(b *Board) []Move {
	moves := make([]Move,0)
	dir := ((k.GetPiece().rank + k.GetPiece().file) % 2 )*2 - 1
	checkMoveAttack(b, k.p, dir, -2, moves)
	checkMoveAttack(b, k.p, dir, -1, moves)
	checkMoveAttack(b, k.p, dir, 0, moves)
	checkMoveAttack(b, k.p, dir, 1, moves)
	checkMoveAttack(b, k.p, dir, 2, moves)
	checkMoveAttack(b, k.p, 0, -2, moves)
	checkMoveAttack(b, k.p, 0, -1, moves)
	checkMoveAttack(b, k.p, 0, 1, moves)
	checkMoveAttack(b, k.p, 0, 2, moves)
	checkMoveAttack(b, k.p, dir, -1, moves)
	checkMoveAttack(b, k.p, dir, 0, moves)
	checkMoveAttack(b, k.p, dir, 1, moves)
	return moves
}
func (k King) GetPiece() *Piece { return k.p }

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
	triChess = website.CreateSite("chess", "localhost:8070", "en")
	triChess.Html.Add("nav", "ul", []string{"class::nav nav-pills nav-stacked"}).
		AddTo("nav","li",[]string{"My Games","url::/games"}).
		AddTo("nav","li",[]string{"New Game","url::/settings"}).
		AddTo("nav","li",[]string{"Teams","url::/newGame"}).
		AddTo("nav","li",[]string{"Clubs","url::/teams"}).
		AddTo("nav","li",[]string{"Message","url::/clubs"}).
		AddTo("nav","li",[]string{"Login","url::/login"})

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
	
	boardId := "PONMLKJIHGFEDCBA12345678"
	scaleX := 30
	scaleY := 15
	offX := 140
	offY := 40
	spaces := 4
	perspective := 2
	indexId := 0
	
	chess.Html.Add("triChessBoard", "svg", []string{"xmlns::http://www.w3.org/2000/svg", "height=::720", "width=960"})
	chess.Html.AddTo("triChessBoard", "circle", []string{
		"transform::skewX(12) scale(1,0.66)", 
		fmt.Sprintf("cx::%d",offX+272), 
		fmt.Sprintf("cy::%d",offY+443), 
		fmt.Sprintf("r::%d",470), 
		"stroke::black", "stroke-width::2", "fill::#412"})
	
	for y := 0; y<spaces; y++ {
		scaleY += perspective
		for x := 0; x<spaces+1+y; x++ {
			if x>0 {
				chess.Html.AddTo("triChessBoard", "polygon", triangle(offX,offY,scaleX,scaleY,perspective,
					2*x-y,2*y,2*x-y+1,2*y+2,2*x-y+2,2*y,"downTri", string(boardId[y*2])+"_"+string(boardId[x-y+19]),0))
				chess.Data["board"] = append(chess.Data["board"], fmt.Sprintf(
					`$().ready(function() { document.getElementById('%s').addEventListener('click', function(event) { alert('fda'); $(this).css({fill:#919}); } ); });`, 
					string(boardId[y*2])+"_"+string(boardId[x-y+19])))
					indexId ++
			}
			chess.Html.AddTo("triChessBoard", "polygon", triangle(offX,offY,scaleX,scaleY,perspective,
				2*x-y+1,2*y+2,2*x-y+2,2*y,2*x-y+3,2*y+2,"upTri", string(boardId[y*2+1])+"_"+string(boardId[x-y+19]),1))
					indexId ++
		}
	}
	for y := spaces; y<spaces*2; y++ {
		scaleY += perspective
		for x := 0; x<spaces*3-y; x++ {
			if x>0 {		
				chess.Html.AddTo("triChessBoard", "polygon", triangle(offX,offY,scaleX,scaleY,perspective,
					2*x+y+3-spaces*2,2*y+2,2*x+y+2-spaces*2,2*y,2*x+y+1-spaces*2,2*y+2,"upTri", 
					string(boardId[y*2+1])+"_"+string(boardId[x+15]),1))
					indexId ++
			}	
			chess.Html.AddTo("triChessBoard", "polygon", triangle(offX,offY,scaleX,scaleY,perspective,
					2*x+y+2-spaces*2,2*y,2*x+y+3-spaces*2,2*y+2,2*x+y+4-spaces*2,2*y,"downTri", 
					string(boardId[y*2])+"_"+string(boardId[x+16]),0))
				indexId ++			
		}
	}
	
	chess.Html.Add("blackKing", "g", []string{
		fmt.Sprintf("transform:::translate(%d,%d) scale(1.1)",offX+215,offY-20)),
		"id:::blackKing",
		"style:::fill:none; fill-opacity:1; fill-rule:evenodd; stroke:#000000; stroke-width:1.5; stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:4; stroke-dasharray:none; stroke-opacity:1;" }
	chess.Html.AddTo("blackKing", "path", []string{
		"d:::M 22.5,11.63 L 22.5,6",
		"style:::fill:none; stroke:#000000; stroke-linejoin:miter;",
		"id:::path6570"
	}	
	chess.Html.AddTo("blackKing", "path", []string{
		"d:::M 22.5,25 C 22.5,25 27,17.5 25.5,14.5 C 25.5,14.5 24.5,12 22.5,12 C 20.5,12 19.5,14.5 19.5,14.5 C 18,17.5 22.5,25 22.5,25",
		"style:::fill:#000000;fill-opacity:1; stroke-linecap:butt; stroke-linejoin:miter;"
	}	
	chess.Html.AddTo("blackKing", "path", []string{
		"d:::M 11.5,37 C 17,40.5 27,40.5 32.5,37 L 32.5,30 C 32.5,30 41.5,25.5 38.5,19.5 C 34.5,13 25,16 22.5,23.5 L 22.5,27 L 22.5,23.5 C 19,16 9.5,13 6.5,19.5 C 3.5,25.5 11.5,29.5 11.5,29.5 L 11.5,37 z ",
		"style:::fill:#000000; stroke:#000000;"
	}	
	chess.Html.AddTo("blackKing", "path", []string{
		"d:::M 20,8 L 25,8",
		"style:::fill:none; stroke:#000000; stroke-linejoin:miter;"
	}	
	chess.Html.AddTo("blackKing", "path", []string{
		"d:::M 32,29.5 C 32,29.5 40.5,25.5 38.03,19.85 C 34.15,14 25,18 22.5,24.5 L 22.51,26.6 L 22.5,24.5 C 20,18 9.906,14 6.997,19.85 C 4.5,25.5 11.85,28.85 11.85,28.85",
		"fill:none; stroke:#ffffff;"
	}	
	chess.Html.AddTo("blackKing", "path", []string{
		"d:::M 11.5,30 C 17,27 27,27 32.5,30 M 11.5,33.5 C 17,30.5 27,30.5 32.5,33.5 M 11.5,37 C 17,34 27,34 32.5,37",
		"fill:none; stroke:#ffffff;"
	}
	
	chess.Html.Add("blackQueen", "g", []string{
		"transform:::translate(%d,%d) scale(1.1)",
		"id:::blackQueen", 
		"style:::opacity:1; fill:000000; fill-opacity:1; fill-rule:evenodd; stroke:#000000; stroke-width:1.5; stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:4; stroke-dasharray:none; stroke-opacity:1;"	
	}
	chess.Html.Tag("blackQueen",)

				<g style="fill:#000000; stroke:none;">
					<circle cx="6"    cy="12" r="2.75" />	<circle cx="14"   cy="9"  r="2.75" />	<circle cx="22.5" cy="8"  r="2.75" />	<circle cx="31"   cy="9"  r="2.75" />	<circle cx="39"   cy="12" r="2.75" />
				</g>
					<path d="M 9,26 C 17.5,24.5 30,24.5 36,26 L 38.5,13.5 L 31,25 L 30.7,10.9 L 25.5,24.5 L 22.5,10 L 19.5,24.5 L 14.3,10.9 L 14,25 L 6.5,13.5 L 9,26 z" style="stroke-linecap:butt; stroke:#000000;" />
					<path d="M 9,26 C 9,28 10.5,28 11.5,30 C 12.5,31.5 12.5,31 12,33.5 C 10.5,34.5 10.5,36 10.5,36 C 9,37.5 11,38.5 11,38.5 C 17.5,39.5 27.5,39.5 34,38.5 C 34,38.5 35.5,37.5 34,36 C 34,36 34.5,34.5 33,33.5 C 32.5,31 32.5,31.5 33.5,30 C 34.5,28 36,28 36,26 C 27.5,24.5 17.5,24.5 9,26 z" style="stroke-linecap:butt;" />
					<path d="M 11,38.5 A 35,35 1 0 0 34,38.5" style="fill:none; stroke:#000000; stroke-linecap:butt;" />
					<path d="M 11,29 A 35,35 1 0 1 34,29" style="fill:none; stroke:#ffffff;" />
					<path d="M 12.5,31.5 L 32.5,31.5" style="fill:none; stroke:#ffffff;" />
					<path d="M 11.5,34.5 A 35,35 1 0 0 33.5,34.5" style="fill:none; stroke:#ffffff;" />
					<path d="M 10.5,37.5 A 35,35 1 0 0 34.5,37.5" style="fill:none; stroke:#ffffff;" />
				</g>
			</g>`,offX+313,offY-20)))
	chess.Data["piece"] = append(chess.Data["piece"], 
			template.HTML(fmt.Sprintf(`
				<g transform="translate(%d,%d) scale(1.1)" id="blackCannon"  stroke-linejoin="round" fill:000000; fill-opacity:1; stroke="#000" stroke-width:1.5; stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:4; stroke-dasharray:none; stroke-opacity:1;>
					<g stroke-width="1.00000001" stroke-linecap="butt" fill="#000">
						<path d="M13,20c-1.3,0.9-3.4,1.5-5,1.5-1.64,0-2.74,1-4,2.5,1,0,2.87-0.5,4.5-0.5,1.6,0,3.3,0,4.5-0.5,1.2,0.5,2.9,0.5,4.5,0.5s3.5,0.5,4.5,0.5c-1.3-1.5-2.4-2.5-4-2.5s-3.7-0.6-5-1.5z" transform="scale(1.7307692,1.7307692)"/>
						<path d="M11,4c-2-2-8,4-6,6l5,5s-2,1.5-1,2.5c0,0-1,1,0,2,0.71,0.7,2,0.5,4,0.5s3.3,0.2,4-0.5c1-1,0-2,0-2,1-1-1-2.5-1-2.5,1.2467-0.35064,2.4102-3.5898,1-5z" transform="scale(1.7307692,1.7307692)"/>
					</g>
					<path stroke-linejoin="miter" d="m17.3,26,10.4,0m-12.1,4.33,13.8,0" stroke-linecap="square" stroke-width="2.07692313" fill="none"/>
				</g>`,offX+265,offY-12)))
	for i:=0; i<2; i++ {
		chess.Data["piece"] = append(chess.Data["piece"], 
			template.HTML(fmt.Sprintf(`
				<g transform="translate(%d,%d) scale(1.1)" id="blackRook%d" style="opacity:1; fill:000000; fill-opacity:1; fill-rule:evenodd; stroke:#000000; stroke-width:1.5; stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:4; stroke-dasharray:none; stroke-opacity:1;">
					<path d="M 9,39 L 36,39 L 36,36 L 9,36 L 9,39 z " style="stroke-linecap:butt;" />
					<path d="M 12.5,32 L 14,29.5 L 31,29.5 L 32.5,32 L 12.5,32 z " style="stroke-linecap:butt;" />
					<path d="M 12,36 L 12,32 L 33,32 L 33,36 L 12,36 z " style="stroke-linecap:butt;" />
					<path d="M 14,29.5 L 14,16.5 L 31,16.5 L 31,29.5 L 14,29.5 z " style="stroke-linecap:butt;stroke-linejoin:miter;" />
					<path d="M 14,16.5 L 11,14 L 34,14 L 31,16.5 L 14,16.5 z " style="stroke-linecap:butt;" />
					<path d="M 11,14 L 11,9 L 15,9 L 15,11 L 20,11 L 20,9 L 25,9 L 25,11 L 30,11 L 30,9 L 34,9 L 34,14 L 11,14 z " style="stroke-linecap:butt;" />
					<path d="M 12,35.5 L 33,35.5 L 33,35.5" style="fill:none; stroke:#ffffff; stroke-width:1; stroke-linejoin:miter;" />
					<path d="M 13,31.5 L 32,31.5" style="fill:none; stroke:#ffffff; stroke-width:1; stroke-linejoin:miter;" />
					<path d="M 14,29.5 L 31,29.5" style="fill:none; stroke:#ffffff; stroke-width:1; stroke-linejoin:miter;" />
					<path d="M 14,16.5 L 31,16.5" style="fill:none; stroke:#ffffff; stroke-width:1; stroke-linejoin:miter;" />
					<path d="M 11,14 L 34,14" style="fill:none; stroke:#ffffff; stroke-width:1; stroke-linejoin:miter;" />
				</g>`,offX+68+390*i,offY-12,i+1)))
	}
	for i:=0; i<2; i++ {
		chess.Data["piece"] = append(chess.Data["piece"], 
			template.HTML(fmt.Sprintf(`
				<g transform="translate(%d,%d) scale(1)" id="blackKnight%d" style="opacity:1; fill:none; fill-opacity:1; fill-rule:evenodd; stroke:#000000; stroke-width:1.5; stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:4; stroke-dasharray:none; stroke-opacity:1;">
					<path d="M 22,10 C 32.5,11 38.5,18 38,39 L 15,39 C 15,30 25,32.5 23,18" style="fill:#000000; stroke:#000000;" />
					<path d="M 24,18 C 24.38,20.91 18.45,25.37 16,27 C 13,29 13.18,31.34 11,31 C 9.958,30.06 12.41,27.96 11,28 C 10,28 11.19,29.23 10,30 C 9,30 5.997,31 6,26 C 6,24 12,14 12,14 C 12,14 13.89,12.1 14,10.5 C 13.27,9.506 13.5,8.5 13.5,7.5 C 14.5,6.5 16.5,10 16.5,10 L 18.5,10 C 18.5,10 19.28,8.008 21,7 C 22,7 22,10 22,10" style="fill:#000000; stroke:#000000;" />
					<path d="M 9.5 25.5 A 0.5 0.5 0 1 1 8.5,25.5 A 0.5 0.5 0 1 1 9.5 25.5 z" style="fill:#ffffff; stroke:#ffffff;" />
					<path d="M 15 15.5 A 0.5 1.5 0 1 1  14,15.5 A 0.5 1.5 0 1 1  15 15.5 z" transform="matrix(0.866,0.5,-0.5,0.866,9.693,-5.173)" style="fill:#ffffff; stroke:#ffffff;" />
					<path d="M 24.55,10.4 L 24.1,11.85 L 24.6,12 C 27.75,13 30.25,14.49 32.5,18.75 C 34.75,23.01 35.75,29.06 35.25,39 L 35.2,39.5 L 37.45,39.5 L 37.5,39 C 38,28.94 36.62,22.15 34.25,17.66 C 31.88,13.17 28.46,11.02 25.06,10.5 L 24.55,10.4 z " style="fill:#ffffff; stroke:none;" />
				</g>`,offX+120+288*i,offY-20,i+1)))
	}
	for i:=0; i<2; i++ {
	chess.Data["piece"] = append(chess.Data["piece"], 
		template.HTML(fmt.Sprintf(`
			<g transform="translate(%d,%d) scale(1)" id="blackBissop%d" style="opacity:1; fill:none; fill-rule:evenodd; fill-opacity:1; stroke:#000000; stroke-width:1.5; stroke-linecap:round; stroke-linejoin:round; stroke-miterlimit:4; stroke-dasharray:none; stroke-opacity:1;">
				<g style="fill:#000000; stroke:#000000; stroke-linecap:butt;">
					<path d="M 9,36 C 12.39,35.03 19.11,36.43 22.5,34 C 25.89,36.43 32.61,35.03 36,36 C 36,36 37.65,36.54 39,38 C 38.32,38.97 37.35,38.99 36,38.5 C 32.61,37.53 25.89,38.96 22.5,37.5 C 19.11,38.96 12.39,37.53 9,38.5 C 7.646,38.99 6.677,38.97 6,38 C 7.354,36.06 9,36 9,36 z" />
					<path d="M 15,32 C 17.5,34.5 27.5,34.5 30,32 C 30.5,30.5 30,30 30,30 C 30,27.5 27.5,26 27.5,26 C 33,24.5 33.5,14.5 22.5,10.5 C 11.5,14.5 12,24.5 17.5,26 C 17.5,26 15,27.5 15,30 C 15,30 14.5,30.5 15,32 z" />
					<path d="M 25 8 A 2.5 2.5 0 1 1  20,8 A 2.5 2.5 0 1 1  25 8 z" />
				</g>
				<path d="M 17.5,26 L 27.5,26 M 15,30 L 30,30 M 22.5,15.5 L 22.5,20.5 M 20,18 L 25,18" style="fill:none; stroke:#ffffff; stroke-linejoin:miter;" />
			</g>`,offX+168+193*i,offY-10,i+1)))
	}
	for i:=0 ; i<5; i++ {
		chess.Data["piece"] = append(chess.Data["piece"], 
			template.HTML(fmt.Sprintf(`
				<path transform="translate(%d,%d) scale(1.2)" id="blackPawn%d" d="M 22,9 C 19.79,9 18,10.79 18,13 C 18,13.89 18.29,14.71 18.78,15.38 C 16.83,16.5 15.5,18.59 15.5,21 C 15.5,23.03 16.44,24.84 17.91,26.03 C 14.91,27.09 10.5,31.58 10.5,39.5 L 33.5,39.5 C 33.5,31.58 29.09,27.09 26.09,26.03 C 27.56,24.84 28.5,23.03 28.5,21 C 28.5,18.59 27.17,16.5 25.22,15.38 C 25.71,14.71 26,13.89 26,13 C 26,10.79 24.21,9 22,9 z " style="opacity:1; fill:#000000; fill-opacity:1; fill-rule:nonzero; stroke:#000000; stroke-width:1.5; stroke-linecap:round; stroke-linejoin:miter; stroke-miterlimit:4; stroke-dasharray:none; stroke-opacity:1;" />
				`,offX+75 + i*100,offY+16 , 2*i+2)))
	}	
	for i:=0 ; i<6; i++ {
		chess.Data["piece"] = append(chess.Data["piece"], 
			template.HTML(fmt.Sprintf(`
				<path transform="translate(%d,%d) scale(1.2)" id="blackPawn%d" d="M 22,9 C 19.79,9 18,10.79 18,13 C 18,13.89 18.29,14.71 18.78,15.38 C 16.83,16.5 15.5,18.59 15.5,21 C 15.5,23.03 16.44,24.84 17.91,26.03 C 14.91,27.09 10.5,31.58 10.5,39.5 L 33.5,39.5 C 33.5,31.58 29.09,27.09 26.09,26.03 C 27.56,24.84 28.5,23.03 28.5,21 C 28.5,18.59 27.17,16.5 25.22,15.38 C 25.71,14.71 26,13.89 26,13 C 26,10.79 24.21,9 22,9 z " style="opacity:1; fill:#000000; fill-opacity:1; fill-rule:nonzero; stroke:#000000; stroke-width:1.5; stroke-linecap:round; stroke-linejoin:miter; stroke-miterlimit:4; stroke-dasharray:none; stroke-opacity:1;" />
				`,offX+30 + i*100, offY+26, 2*i+1)))
	}
	for i:=0 ; i<6; i++ {
		chess.Data["piece"] = append(chess.Data["piece"], 
			template.HTML(fmt.Sprintf(`
				<path transform="translate(%d,%d) scale(1.8)" id="whitePawn%d" d="M 22,9 C 19.79,9 18,10.79 18,13 C 18,13.89 18.29,14.71 18.78,15.38 C 16.83,16.5 15.5,18.59 15.5,21 C 15.5,23.03 16.44,24.84 17.91,26.03 C 14.91,27.09 10.5,31.58 10.5,39.5 L 33.5,39.5 C 33.5,31.58 29.09,27.09 26.09,26.03 C 27.56,24.84 28.5,23.03 28.5,21 C 28.5,18.59 27.17,16.5 25.22,15.38 C 25.71,14.71 26,13.89 26,13 C 26,10.79 24.21,9 22,9 z " style="opacity:1; fill:#ffffff; fill-opacity:1; fill-rule:nonzero; stroke:#000000; stroke-width:1.5; stroke-linecap:round; stroke-linejoin:miter; stroke-miterlimit:4; stroke-dasharray:none; stroke-opacity:1;" />
				`, offX+ 20 + i*120, offY+320, 2*i+1)))
	}
	for i:=0 ; i<5; i++ {
		chess.Data["piece"] = append(chess.Data["piece"], 
			template.HTML(fmt.Sprintf(`
				<path transform="translate(%d,%d) scale(1.8)" id="whitePawn%d" d="M 22,9 C 19.79,9 18,10.79 18,13 C 18,13.89 18.29,14.71 18.78,15.38 C 16.83,16.5 15.5,18.59 15.5,21 C 15.5,23.03 16.44,24.84 17.91,26.03 C 14.91,27.09 10.5,31.58 10.5,39.5 L 33.5,39.5 C 33.5,31.58 29.09,27.09 26.09,26.03 C 27.56,24.84 28.5,23.03 28.5,21 C 28.5,18.59 27.17,16.5 25.22,15.38 C 25.71,14.71 26,13.89 26,13 C 26,10.79 24.21,9 22,9 z " style="opacity:1; fill:#ffffff; fill-opacity:1; fill-rule:nonzero; stroke:#000000; stroke-width:1.5; stroke-linecap:round; stroke-linejoin:miter; stroke-miterlimit:4; stroke-dasharray:none; stroke-opacity:1;" />
				`, offX+ 80 + i*120, offY+340, 2*i+2)))
	}
	chess.Data["piece"] = append(chess.Data["piece"], 
			template.HTML(fmt.Sprintf(`
				<g transform="translate(%d,%d) scale(1.8)" id="whiteCannon"  stroke-linejoin="round" stroke="#000" stroke-miterlimit="4" stroke-dasharray="none">
					<g stroke-width="1.00000001" stroke-linecap="butt" fill="#FFF">
						<path d="M13,20c-1.3,0.9-3.4,1.5-5,1.5-1.64,0-2.74,1-4,2.5,1,0,2.87-0.5,4.5-0.5,1.6,0,3.3,0,4.5-0.5,1.2,0.5,2.9,0.5,4.5,0.5s3.5,0.5,4.5,0.5c-1.3-1.5-2.4-2.5-4-2.5s-3.7-0.6-5-1.5z" transform="scale(1.7307692,1.7307692)"/>
						<path d="M11,4c-2-2-8,4-6,6l5,5s-2,1.5-1,2.5c0,0-1,1,0,2,0.71,0.7,2,0.5,4,0.5s3.3,0.2,4-0.5c1-1,0-2,0-2,1-1-1-2.5-1-2.5,1.2467-0.35064,2.4102-3.5898,1-5z" transform="scale(1.7307692,1.7307692)"/>
					</g>
					<path stroke-linejoin="miter" d="m17.3,26,10.4,0m-12.1,4.33,13.8,0" stroke-linecap="square" stroke-width="2.07692313" fill="none"/>
				</g>`,offX+320,offY+410)))
	chess.Data["piece"] = append(chess.Data["piece"], 
		template.HTML(fmt.Sprintf(`
			<g transform="translate(%d,%d) scale(1.8)" id="whiteKing" style="fill:none; fill-opacity:1; fill-rule:evenodd; stroke:#000000; stroke-width:1.5; stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:4; stroke-dasharray:none; stroke-opacity:1;">
				<path d="M 22.5,11.63 L 22.5,6"	style="fill:none; stroke:#000000; stroke-linejoin:miter;" />
				<path d="M 20,8 L 25,8"	style="fill:none; stroke:#000000; stroke-linejoin:miter;" />
				<path d="M 22.5,25 C 22.5,25 27,17.5 25.5,14.5 C 25.5,14.5 24.5,12 22.5,12 C 20.5,12 19.5,14.5 19.5,14.5 C 18,17.5 22.5,25 22.5,25"	style="fill:#ffffff; stroke:#000000; stroke-linecap:butt; stroke-linejoin:miter;" />
				<path d="M 11.5,37 C 17,40.5 27,40.5 32.5,37 L 32.5,30 C 32.5,30 41.5,25.5 38.5,19.5 C 34.5,13 25,16 22.5,23.5 L 22.5,27 L 22.5,23.5 C 19,16 9.5,13 6.5,19.5 C 3.5,25.5 11.5,29.5 11.5,29.5 L 11.5,37 z " style="fill:#ffffff; stroke:#000000;" />
				<path d="M 11.5,30 C 17,27 27,27 32.5,30" style="fill:none; stroke:#000000;" />
				<path d="M 11.5,33.5 C 17,30.5 27,30.5 32.5,33.5" style="fill:none; stroke:#000000;" />
				<path d="M 11.5,37 C 17,34 27,34 32.5,37" style="fill:none; stroke:#000000;" />
			</g>`,offX+270,offY+440)))
	chess.Data["piece"] = append(chess.Data["piece"], 
		template.HTML(fmt.Sprintf(`
			<g transform="translate(%d,%d) scale(1.8)" id="whiteQueen" style="opacity:1; fill:#ffffff; fill-opacity:1; fill-rule:evenodd; stroke:#000000; stroke-width:1.5; stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:4; stroke-dasharray:none; stroke-opacity:1;">
				<path d="M 9 13 A 2 2 0 1 1  5,13 A 2 2 0 1 1  9 13 z" transform="translate(-1,-1)" />
				<path d="M 9 13 A 2 2 0 1 1  5,13 A 2 2 0 1 1  9 13 z" transform="translate(15.5,-5.5)" />
				<path d="M 9 13 A 2 2 0 1 1  5,13 A 2 2 0 1 1  9 13 z" transform="translate(32,-1)" />
				<path d="M 9 13 A 2 2 0 1 1  5,13 A 2 2 0 1 1  9 13 z" transform="translate(7,-4.5)" />
				<path d="M 9 13 A 2 2 0 1 1  5,13 A 2 2 0 1 1  9 13 z" transform="translate(24,-4)" />
				<path d="M 9,26 C 17.5,24.5 30,24.5 36,26 L 38,14 L 31,25 L 31,11 L 25.5,24.5 L 22.5,9.5 L 19.5,24.5 L 14,10.5 L 14,25 L 7,14 L 9,26 z " style="stroke-linecap:butt;" />
				<path d="M 9,26 C 9,28 10.5,28 11.5,30 C 12.5,31.5 12.5,31 12,33.5 C 10.5,34.5 10.5,36 10.5,36 C 9,37.5 11,38.5 11,38.5 C 17.5,39.5 27.5,39.5 34,38.5 C 34,38.5 35.5,37.5 34,36 C 34,36 34.5,34.5 33,33.5 C 32.5,31 32.5,31.5 33.5,30 C 34.5,28 36,28 36,26 C 27.5,24.5 17.5,24.5 9,26 z " style="stroke-linecap:butt;" />
				<path d="M 11.5,30 C 15,29 30,29 33.5,30" style="fill:none;" />
				<path d="M 12,33.5 C 18,32.5 27,32.5 33,33.5" style="fill:none;" />
			</g>`,offX+385,offY+435)))
	for i:=0; i<2; i++ {	
		chess.Data["piece"] = append(chess.Data["piece"], 
			template.HTML(fmt.Sprintf(`
				<g transform="translate(%d,%d) scale(1.8)" id="whiteKnight%d" style="opacity:1; fill:none; fill-opacity:1; fill-rule:evenodd; stroke:#000000; stroke-width:1.5; stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:4; stroke-dasharray:none; stroke-opacity:1;">
					<path d="M 22,10 C 32.5,11 38.5,18 38,39 L 15,39 C 15,30 25,32.5 23,18" style="fill:#ffffff; stroke:#000000;" />
					<path d="M 24,18 C 24.38,20.91 18.45,25.37 16,27 C 13,29 13.18,31.34 11,31 C 9.958,30.06 12.41,27.96 11,28 C 10,28 11.19,29.23 10,30 C 9,30 5.997,31 6,26 C 6,24 12,14 12,14 C 12,14 13.89,12.1 14,10.5 C 13.27,9.506 13.5,8.5 13.5,7.5 C 14.5,6.5 16.5,10 16.5,10 L 18.5,10 C 18.5,10 19.28,8.008 21,7 C 22,7 22,10 22,10" style="fill:#ffffff; stroke:#000000;" />
					<path d="M 9.5 25.5 A 0.5 0.5 0 1 1 8.5,25.5 A 0.5 0.5 0 1 1 9.5 25.5 z" style="fill:#000000; stroke:#000000;" />
					<path d="M 15 15.5 A 0.5 1.5 0 1 1  14,15.5 A 0.5 1.5 0 1 1  15 15.5 z" transform="matrix(0.866,0.5,-0.5,0.866,9.693,-5.173)" style="fill:#000000; stroke:#000000;" />
				</g>`,offX+135+i*380,offY+435, i+1)))
	}
	for i:=0; i<2; i++ {
		chess.Data["piece"] = append(chess.Data["piece"], 
			template.HTML(fmt.Sprintf(`
				<g transform="translate(%d,%d) scale(1.8)" id="whiteRook%d" style="opacity:1; fill:#ffffff; fill-opacity:1; fill-rule:evenodd; stroke:#000000; stroke-width:1.5; stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:4; stroke-dasharray:none; stroke-opacity:1;">
					<path d="M 9,39 L 36,39 L 36,36 L 9,36 L 9,39 z " style="stroke-linecap:butt;" />
					<path d="M 12,36 L 12,32 L 33,32 L 33,36 L 12,36 z " style="stroke-linecap:butt;" />
					<path d="M 11,14 L 11,9 L 15,9 L 15,11 L 20,11 L 20,9 L 25,9 L 25,11 L 30,11 L 30,9 L 34,9 L 34,14" style="stroke-linecap:butt;" />
					<path d="M 34,14 L 31,17 L 14,17 L 11,14" />
					<path d="M 31,17 L 31,29.5 L 14,29.5 L 14,17" style="stroke-linecap:butt; stroke-linejoin:miter;" />
					<path d="M 31,29.5 L 32.5,32 L 12.5,32 L 14,29.5" />
					<path d="M 11,14 L 34,14" style="fill:none; stroke:#000000; stroke-linejoin:miter;" />
				</g>`,offX+80+i*500,offY+410, i+1)))
	}
	for i:=0; i<2; i++ {
		chess.Data["piece"] = append(chess.Data["piece"], 
			template.HTML(fmt.Sprintf(`
				<g transform="translate(%d,%d) scale(1.8)" id="whiteBishop%d" style="opacity:1; fill:none; fill-rule:evenodd; fill-opacity:1; stroke:#000000; stroke-width:1.5; stroke-linecap:round; stroke-linejoin:round; stroke-miterlimit:4; stroke-dasharray:none; stroke-opacity:1;">
					<g style="fill:#ffffff; stroke:#000000; stroke-linecap:butt;"> 
						<path d="M 9,36 C 12.39,35.03 19.11,36.43 22.5,34 C 25.89,36.43 32.61,35.03 36,36 C 36,36 37.65,36.54 39,38 C 38.32,38.97 37.35,38.99 36,38.5 C 32.61,37.53 25.89,38.96 22.5,37.5 C 19.11,38.96 12.39,37.53 9,38.5 C 7.646,38.99 6.677,38.97 6,38 C 7.354,36.06 9,36 9,36 z" />
						<path d="M 15,32 C 17.5,34.5 27.5,34.5 30,32 C 30.5,30.5 30,30 30,30 C 30,27.5 27.5,26 27.5,26 C 33,24.5 33.5,14.5 22.5,10.5 C 11.5,14.5 12,24.5 17.5,26 C 17.5,26 15,27.5 15,30 C 15,30 14.5,30.5 15,32 z" />
						<path d="M 25 8 A 2.5 2.5 0 1 1  20,8 A 2.5 2.5 0 1 1  25 8 z" />
					</g>
					<path d="M 17.5,26 L 27.5,26 M 15,30 L 30,30 M 22.5,15.5 L 22.5,20.5 M 20,18 L 25,18" style="fill:none; stroke:#000000; stroke-linejoin:miter;" />
				</g>`,offX+200+i*250,offY+410, i+1)))
	}
	
	
	clubs := triChess.AddPage("clubs", "clubs", "/clubs")
	clubs.AddAJAXHandler("postStatement", mgs.PostStatement)
	clubs.AddAJAXHandler("createRoom", mgs.CreateRoom)
	clubs.AddParam("roomPanel",`
		var ul = $( "<ul/>", {"class": "my-new-list"}); 
		var obj = JSON.parse(data);	
		$("#rooms").empty(); 
		$("#rooms").append(ul); 
		$.each(obj, function(i,val) { 
			console.log('adding room '+i); 
			room = $('<li />', { } ).appendTo(ul);
			button = $('<button />', { 'class':'accordian', 'text': i + ' - ' + val } ).appendTo(room); 
			$('<div />', { 'class':'panel', 'text': 'occupants:' } ).appendTo(button); 
		}); 
		collapsable();`)
}

func triangle(offX,offY,scaleX,scaleY,perspective,px1,py1,px2,py2,px3,py3 int, id, pClass string, up int) []string {
	return []string{
		fmt.Sprintf("id:::%s", id),
		fmt.Sprintf("class:::%s", pClass),
		fmt.Sprintf("points:::%d,%d %d,%d %d,%d", offX+px1*(scaleX+scaleY+up*perspective), offY+py1*(scaleY+up*perspective), 
			offX+px2*(scaleX+scaleY+(1-up)*perspective), offY+py2*(scaleY+(1-up)*perspective), 
			offX+px3*(scaleX+scaleY+up*perspective), offY+py3*(scaleY+up*perspective))}
}