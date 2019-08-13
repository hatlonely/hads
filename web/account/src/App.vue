<template>
  <v-app>
    <transition name="slide-x-transition" mode="out-in">
      <router-view></router-view>
    </transition>
  </v-app>
</template>

<style>
@import "./assets/css/font.css";
/* @import url("https://fonts.googleapis.com/css?family=Montserrat|Open+Sans|Noto+Sans+SC|Noto+Serif+SC|ZCOOL+XiaoWei&display=swap"); */

*,
.display-4,
.display-3,
.display-2,
.display-1,
.headline,
.title,
.subtitle-1,
.subtitle-2,
.body-1,
.body-2,
.caption,
.overline {
  /* font-family: "Red Hat Text", "ZCool XiaoWei", sans-serif; */
  /* font-family: "Red Hat Text", "Source Han Sans SC", sans-serif; */
  /* font-family: "Red Hat Text", "Yahei", sans-serif; */
  font-family: sans-serif;
}

a {
  text-decoration: none;
}

footer.v-footer a {
  color: rgba(0, 0, 0, 0.54);
}
</style>
