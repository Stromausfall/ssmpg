title: Kivy - Resist Temptation
date: 2015-04-19 22:22
categories:
- Python

---

[Kivy](http://kivy.org/) is a cross-plattform framework that allows to run python code on Linux, Windows, OS X, Android and iOS. The framework allows simple, stable and quick creation of apps.

The second application I created using the framework after the obligatory 'Hello World',  is a small application that allows to track *distractions*.

The app measures the time until the user is distracted from whatever task she does (when distracted the user has to press the **fail** button). Should the user feel an urge to give in to the temptation of a distraction, but manages to resist it (by consciously deciding to resist) she clicks the **resist** button.

When **fail** is clicked the current score is used as the **high-score** if the time is higher than the previous *high-score*. If **restart** is clicked the current score is simply thrown away.

Two measurements are performed, *time* and how many *minutes between resists*.

[Project Source](http://www.matthias-auer.net/Projects/ResistTemptation/ResistTemptation.zip)

[Android binaries](http://www.matthias-auer.net/Projects/ResistTemptation/ResistTemptation-0.02.apk)

![Resist Temptation](http://www.matthias-auer.net/Projects/ResistTemptation/ResistTemptation.png)
