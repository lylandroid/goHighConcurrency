## …or create a new repository on the command line
    echo "# goPathIris" >> README.md
    git init
    git add README.md（git add .| git add --all）
    git commit -m "first commit"
    git remote add origin git@github.com:lylandroid/goPathIris.git
    git push -u origin master
    
## …or push an existing repository from the command line
    git remote add origin git@github.com:lylandroid/goPathIris.git
    git push -u origin master