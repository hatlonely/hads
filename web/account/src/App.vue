<template>
  <v-app>
    <h-app-bar />
    <v-content>
      <v-layout row>
        <v-flex xs2>
          <h-sider />
        </v-flex>
        <v-flex xs8 px-12>
          <h-body />
        </v-flex>
      </v-layout>
    </v-content>
  </v-app>
</template>

<script>
import HBody from "./components/HBody";
import HAppBar from "./components/HAppBar";
import HSider from "./components/HSider";

export default {
  name: "App",
  components: {
    HAppBar,
    HBody,
    HSider
  }
};
</script>
