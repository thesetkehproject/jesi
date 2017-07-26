# Jesi
Jesi something that does things.

---
[![Build Status](http://drone.setkeh.com:8000/api/badges/thesetkehproject/jesi/status.svg)](http://drone.setkeh.com:8000/thesetkehproject/jesi)

#### Idea Pad

* Screenshot Api to Store Screenshot in Base64 in a Database.
* Restful Blog Service.
* Social Rest Integrations (Probably Github, Spotify, Fitbit, Twitter, Possibly More)
* Probably More as i think of it.

Im thinking this might turn into some sort of super basic CMS for personal websites that has the
ability to integrate with many services and be highly extensible via rest endpoints.

---

Im using a lot of git flow at work so im going to use this repo as the test bed for integrating it into personal projects.
As it a Nice clean way of managing git projects.

---
Feel Free to fork and contribute to any part of this project including open issues and project goals.

---
#### Project Structure (Constant WIP)

- model: All of the Project Models Live Here.
- router: The Project routers will live here.
- shared: All the shared variable thet the whole project needs will be here.

---
#### Git Flow

The Master branch should always be working code and never committed to.

Develop branch will be all testing code and currently in development but should also never be committed too, Each project milestone will be forked from master upon commencement and no more than one milestone should be in progress at once.

feature branch's: Feature branches are where work should be conducted either in repo or from pull request by contributors once a feature has passed integration testing it will be merged into develop.

This model should help ensure that edge branch (develop) and release (master) should always have mostly/working code.

[Git Flow In Depth Explanation](https://l.facebook.com/l.php?u=https%3A%2F%2Fdatasift.github.io%2Fgitflow%2FIntroducingGitFlow.html&h=ATNOaCBBQVafyBUdykfJS8qEUUeSssmeVArKPYzAj-K858LWSr_FB25tghJu5OMvAzGP3XPhohzVbTaqlS0C3fnYVn92wMYDiHWM0NS8gbdosWdviz8TjmzGJdM4_cArWy4)