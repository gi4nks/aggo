Start building with Ionic!
Follow these quick steps and start building high quality mobile apps in minutes. For a more indepth overview, watch our Crash Course video, which covers everything else you'd want to know.

Install Ionic

First, install Node.js. Then, install the latest Cordova and Ionic command-line tools. Follow the Android and iOS platform guides to install required platform dependencies. Windows users might want to watch this installation video (or try Vagrant below).

We also have a Vagrant package for an all-in-one setup (experimental).

Note: iOS development requires Mac OS X.

$ npm install -g cordova ionic

Start a project

Create an Ionic project using one of our ready-made app templates, or a blank one to start fresh.

$ ionic start myApp tabs
$ ionic start myApp blank $ ionic start myApp tabs $ ionic start myApp sidemenu

Run it

Ionic apps are based on Cordova, so we can use the Cordova utilities to build, test, and deploy our apps, but we provide simple ways to do the same with the ionic utility (substitute ios for android to build for Android):

$ cd myApp
$ ionic platform add ios
$ ionic build ios
$ ionic emulate ios