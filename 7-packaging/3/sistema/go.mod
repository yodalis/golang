module github.com/yodalis/golang/7-packaging/3/sistema

go 1.21.6

replace github.com/yodalis/golang/7-packaging/3/math => ../math 
// Caso tenha alguma situação que você esteja trabalhando com algo seu que ainda n foi publicado, 
// você pode indicar pro go que sempre que você tá falando daquela url, 
// vc está falando também sobre a url relativa que está no replace
// *NÃO É RECOMENDADO DESSA FORMA*

require github.com/yodalis/golang/7-packaging/3/math v0.0.0-00010101000000-000000000000
