############# Para ver las slides se debe ejecutar lo siguiente #############

- Para instalar present se tomo como referencia el siguiente link:
  
  - http://halyph.com/blog/2015/05/18/golang-presentation-tool.html

Instructivo:

1-  Instalar GO

2-  go get golang.org/x/net

3-  go install golang.org/x/tools/cmd/present

4-  Posicionarse donde se encuentran las slides:

	- export GOPATH=$HOME/go
	- export GOBIN=$GOPATH/bin
	- $GOBIN/present

Las slides se veran en http://127.0.0.1:3999

### Mas facil
docker run -it --name go-cys -p 3999:3999 -v /home/phazan/proyectos/cys/presentacionGo/:/app -e "EXTERNAL_HOST=<ip/>" mkboudreau/go-present

