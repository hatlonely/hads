<template>
  <v-app>
    <transition name="slide-x-transition" mode="out-in">
      <router-view></router-view>
    </transition>
  </v-app>
</template>

<style>
a {
  text-decoration: none;
}
</style>

<script>
export default {
  name: "App"
};
</script>
