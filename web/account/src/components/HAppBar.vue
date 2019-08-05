<template>
  <v-app-bar app clipped-left elevate-on-scroll class="pb-3">
    <v-toolbar-title class="headline">
      <span>HADS</span>&nbsp;
      <span class="font-weight">Account</span>
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
    <v-btn text href="https://github.com/vuetifyjs/vuetify/releases/latest" target="_blank">
      <span class="mr-2">Latest Release</span>
    </v-btn>
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