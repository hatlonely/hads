<template>
  <v-app>
    <h-app-bar />
    <v-content>
      <v-layout row>
        <v-flex xs2>
          <h-sider />
        </v-flex>
        <v-flex xs8 px-12>
          <h-home-body />
        </v-flex>
      </v-layout>
    </v-content>
  </v-app>
</template>

<script>
import HHomeBody from "./components/HHomeBody";
import HAppBar from "./components/HAppBar";
import HSider from "./components/HSider";

export default {
  name: "App",
  components: {
    HAppBar,
    HHomeBody,
    HSider
  }
};
</script>
