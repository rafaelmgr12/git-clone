<h1 align="center">Git Clone</h1>
<p align = "center"> A Simplify version control system similair to git</p>

<p align="center">
  <a href="#-technology">Technology</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
    <a href="#-project">Project</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-how-to-run">How to Run</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-license">License</a>
</p>

<p align="center">
  <img alt="License" src="https://img.shields.io/static/v1?label=license&message=MIT&color=8257E5&labelColor=000000">
</p>

## Introduction

The objective of this project is to develop a version control system, similar to the renowned tool, Git. However, our goal is to create a simplified version of it. A version control system (VCS) is essential for developers as it allows them to track and manage changes to their codebase over time. With this system in place, developers can work collaboratively, maintain a history of code changes, and easily revert to previous versions if needed.

The following challenge can be found [here](https://app.devgym.com.br/challenges/5b56d4a1-378c-41f0-9c91-7a9577d00671).

### How Git Works: A Brief Overview

Git is a distributed version control system that allows multiple developers to work on the same project simultaneously without interfering with each other's work. Here's a concise overview of how Git operates:

1. **Repository Initialization**: Using the `git init`  command, a new repository is created. This initializes a new `.git` directory where all the information about the repository is stored.
2. **Staging Area**: Before committing changes, files are added to a staging area using the `git add` command. This area holds all the changes that will be included in the next commit.
3. **Commit**: The `git commit` command saves the changes from the staging area to the repository. Each commit has a unique ID, a message describing the changes, and information about the author.
4. **Branching**: Git allows for the creation of branches, which are separate lines of development. This feature enables developers to work on different features or bug fixes without affecting the main codebase.
5. **Merging**: Once the work on a branch is complete, it can be merged back into the main branch using the `git merge` command.
6. **Remote Repositories**: Git supports remote repositories, which are versions of the project hosted on the internet or a network. Developers can push their changes to these repositories or pull changes from them using the `git push` and `git pull` commands, respectively.
7. **Log**: The `git log` command provides a history of all the commits made in the repository, allowing developers to track changes over time.

By understanding these fundamental concepts, one can appreciate the power and flexibility that Git offers to software development teams.

## ✨ Technology

The Project was develop as using the following techs:

- [Go](https://go.dev/)
- [Cobra](github.com/spf13/cobra)

## 💻 Project

In this challenge, we aim to build a command-line interface (CLI) program that can save copies of files and provide feedback on the state of these copies. This tool will be named "fit" (though the name is optional), and it will perform basic version control operations such as initializing a repository, adding files, committing changes, checking the status, and viewing the log of commits.

## 🚀 How to Run

1. Clone the repository
2. Change to the project directory
3. Build the project 
```bash
go build -o fit cmd/main.go
```
### Linux
4. Setting Up Environment Variables:
    * To add the current directory to your PATH, you can edit the .bashrc file (or .zshrc if you're using Zsh) in your home directory:
    
    ```bash
    nano ~/.bashrc
    ```
    
    Add the following line to the end of the file (replace /path/to/directory with the full path to the directory where the fit executable is located):
  ```bash
  export PATH=$PATH:/path/to/directory
  ```
  
### Windows
3. Build the project 
```bash
go build -o fit.exe cmd/main.go
```
4. Setting Up Environment Variables:
    * Right-click on the "This PC" or "Computer" icon on the desktop or in File Explorer and select "Properties".
    * In the left panel, click on "Advanced system settings".
    * Click the "Environment Variables" button at the bottom of the window.
    * In the "System Variables" section, find the "Path" variable and click "Edit".
    * Click "New" and add the full path to the directory where the fit.exe executable is located.
    * Click "OK" to close each of the window

## 📄 License

The projects is under the MIT license. See the file [LICENSE](LICENSE) fore more details

---

## Author

Made with ♥ by Rafael 👋🏻

[![Linkedin Badge](https://img.shields.io/badge/-Rafael-blue?style=flat-square&logo=Linkedin&logoColor=white&link=https://www.linkedin.com/in/tgmarinho/)](https://www.linkedin.com/in/rafael-mgr/)
[![Gmail Badge](https://img.shields.io/badge/-Gmail-red?style=flat-square&link=mailto:nelsonsantosaraujo@hotmail.com)](mailto:ribeirorafaelmatehus@gmail.com)
