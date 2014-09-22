##GIT HowTo
- **How add to the index changed files already commited?:**

	`git add -u`

- **How add to the index both changed files already commited and new ones?:**

	`git add -A`

- **How commit all already commited but modified files?:**

	`git commit -a`

- **How revert local changes?:**

	`git reset --hard HEAD`

	`git checkout -f`

	`git checkout <filename>`

- **How ignore some files?:**
	
	`add the files/dirs in .gitignore`

- **How merge a branch in master?:**

	`git checkout master`

	`git merge <branch name>`

- **How undo add before commit?:**

	`git reset HEAD <file>...`
	
- **How remove branch?:**

	`git branch -d <branch name> // If merged in the master`

	`git branch -D <branch name> // If not merged in the master`

- **How rename branch?:**
	
	`git branch -m <current name> <new name>`

- **How works with tags?:**

	`git tag -l // to see the tags`

	`git checkout tags/<tag> // to checkout`

	`git tag -a v1.1 -m "Versione 1.1 rilasciata il 12/09/2013"`

	`git pull --tag //not necessary`

	`git push --tag // per mandare i tag al server`
	
- **How see info about remote server?:**
	
	`git remote show origin`

- **How pull all branches?**

	`git push origin newtag`

	`git fetch --all`

	`git checkout master`

	`git rebase origin/master`

	`git checkout sviluppo`

	`git rebase origin/sviluppo`

- **How remove all Untracked files?**

	`git clean -f`

	`git clean -f -d`

##GIT-SVN:
###Useful Links

[https://git.wiki.kernel.org/index.php/Git-svn](https://git.wiki.kernel.org/index.php/Git-svn)

[http://git-scm.com/book/ch8-1.html](http://git-scm.com/book/ch8-1.html)

[http://trac.parrot.org/parrot/wiki/git-svn-tutorial](http://trac.parrot.org/parrot/wiki/git-svn-tutorial)

[http://git.or.cz/course/svn.html](http://git.or.cz/course/svn.html)

###Commands
- **clone**

	`git svn clone -s -r HEAD https://svn.parrot.org/parrot`

	- option `-s` only if it's used the standard layout trunk tags branches'
	- option `-r` is for the revision to start
	- examples:
	
			- git svn clone -s -r HEAD --no-minimize-url https://...
			-
			
- **pull:**

	`git svn rebase`
		
- **commit the changes of a branch in master and then in remote svn:**

	`git rebase master`

	`git branch -M master`

- **commit:**

	`git svn dcommit`
	
- **sequence for commit:**

	`git rebase master`:

	- if conflicts:
		- edit files with conflict
		- use `git add` to mark the conflict as resolved
		- when all conflicts are resolved: `git rebase --continue`
	
	`git branch -M master`

	`git svn rebase`
	
	`git svn dcommit`

	- if conflicts:
		- edit files with conflict
		- use 'git add' to mark the conflict as resolved
		- when all conflicts are resolved:
		
			- `git rebase --continue`
			
			- `git svn dcommit`
	
	`git svn rebase`
	
	notes:

	- "git svn dcommit needs update ERROR from SVN: Item is out of date":
		- "http://andy.delcambre.com/2008/03/04/git-svn-workflow.html"
