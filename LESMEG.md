Gå
==

Gå er ei omsetjing av programmeringsspråket [Go](https://golang.org/) til
nynorsk.

For mange mjukvarenyvinnarar er det ei stor utfordring å skriva mjukvare med
framandspråk medan resten av kvardagen hender på nynorsk. Gå er laga for å
avgrensa denne lasta, så kan ein bruka krefter på edb i staden for på målform.

Gå er kraftig inspirert av [Sjå](https://www.viktigperia.org/wiki/Sj%C3%A5).

Innlegging
----------

Gå leggast inn med dette:

	$ go get github.com/aude/g_/cmd/gå

Høve
----

	$ gå hjelp
	Go på nynorsk.
	
	Høve:
	
		gå omset [katalog...] [fil...]
		gå køyr <fil> [argument...]
	
	Td.:
	
		$ gå omset .
		$ gå omset /src/a /src/b
		$ gå omset framifrå.gå samskap.gå oppslag.gå
		$ gå køyr main.gå
		$ gå køyr main.gå -h
		$ gå køyr main.gå input.txt > output.txt
	

Nyvinningar
-----------

I Gå er framandord byta ut med kjende og sæle uttrykk, td.:

- `heiltal` brukast for å skildre heile tal
- Med `medan` kan ein laga løkkjer (same som `for` i framandspråk)
- `viss` i staden for `if`
- Ein kan introdusera ein `break` ved å skriva `kaffipause`

For heilskap, sku [omset.go](cmd/gå/omset.go).

Målsetjing
----------

Me skyt mot stjernene, og siktar på anerkjenning hjå Noregs EDB-mållag som eit
nytt alternativ til [Sjå, Sjå Skarpt, Sjå Nært, Sjå Langtvekk, osb.](https://www.viktigperia.org/wiki/Nynorsk_programmering).
Me er i kontakt med Jon Aasen og Ivar Fosse for førebels utgreiing.
