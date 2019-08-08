<template>
  <v-app>
    <transition name="slide-x-transition" mode="out-in">
      <router-view></router-view>
    </transition>
  </v-app>
</template>

<style>
@import url("https://fonts.googleapis.com/css?family=Montserrat|Open+Sans|Noto+Sans+SC|Noto+Serif+SC|ZCOOL+XiaoWei&display=swap");
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
  font-family: "Montserrat", sans-serif, "Noto Sans SC", sans-serif;
}

a {
  text-decoration: none;
}
</style>

<script>
export default {
  name: "App"
};
</script>
