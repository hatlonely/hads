<template>
  <v-app>
    <h-app-bar />
    <h-sider />

    <v-content>
      <v-flex xs12 sm12 md12 lg10 xl10 px-10>
        <v-layout align-center justify-center text-center row wrap>
          <v-card flat max-width="900" color="#fafafa">
            <transition name="slide-x-transition" mode="out-in">
              <router-view></router-view>
            </transition>
          </v-card>
        </v-layout>
      </v-flex>
    </v-content>
  </v-app>
</template>

<script>
import HAppBar from "../components/account/HAppBar";
import HSider from "../components/account/HSider";

export default {
  name: "Account",
  components: {
    HAppBar,
    HSider
  }
};
</script>
