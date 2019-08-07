<template>
  <v-app-bar app clipped-left elevate-on-scroll class="pb-3" color="rgba(0, 0, 0, 0)">
    <v-img :src="require('../assets/logo.png')" max-width="70" tile></v-img>
    <v-toolbar-title class="headline">
      &nbsp;
      <span class="headline">账户</span>
    </v-toolbar-title>
    <v-spacer></v-spacer>
    <v-text-field
      v-model="search"
      prepend-inner-icon="search"
      label="Search hads account"
      single-line
      rounded
      clearable
      clear-icon="clear"
      class="mt-9"
      filled
      full-width
    ></v-text-field>
    <v-spacer></v-spacer>
    <v-btn text icon>
      <v-icon>apps</v-icon>
    </v-btn>
    <v-btn href="/login" color="primary" depressed>登录</v-btn>
  </v-app-bar>
</template>

<script>
export default {
  name: "HAppBar",
  data: () => ({
    search: ""
  })
};
</script>