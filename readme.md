# gotree
gotreeは対象のディレクトリのファイル/ディレクトリを再帰的に表示するコマンドです。  
windowsでは深さ制限をつけて tree を実行できないので作りました。

# install
````
git clone https://github.com/desktopgame/gotree
cd gotree
go install
````

# how to use
````
gotree -l 2 project
gotree -l 2 -h project
````

# example
````
cd gotree
gotree
>C:\Work\go\desktopgame\gotree>gotree
>.
> go.mod
> gotree.code-workspace
> main.go
> readme.md
````