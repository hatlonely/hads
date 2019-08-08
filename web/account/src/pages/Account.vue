<template>
  <v-app>
    <h-app-bar />
    <h-sider />

    <v-content>
      <v-flex xs12 sm12 md12 lg12 px-12>
        <v-layout align-center justify-center text-center row wrap>
          <v-card flat max-width="900" color="#fafafa">
            <router-view class="view"></router-view>
          </v-card>
        </v-layout>
      </v-flex>
    </v-content>
  </v-app>
</template>

<script>
import HAppBar from "../components/HAppBar";
import HSider from "../components/HSider";
import HFooter from "../components/HFooter";

export default {
  name: "Account",
  components: {
    HAppBar,
    HSider,
    HFooter
  }
};
</script>
