<template>
  <v-app>
    <h-app-bar />
    <h-sider />

    <v-content>
      <v-layout row>
        <v-flex xs10 px-12>
          <router-view class="view"></router-view>
        </v-flex>
      </v-layout>
    </v-content>

    <h-footer />
  </v-app>
</template>

<style>
a {
  text-decoration: none;
}

.v-input__control {
  min-height: 30px !important;
}
</style>

<script>
import HAppBar from "./components/HAppBar";
import HSider from "./components/HSider";
import HFooter from "./components/HFooter";

export default {
  name: "App",
  components: {
    HAppBar,
    HSider,
    HFooter
  }
};
</script>
