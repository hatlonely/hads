<template>
  <v-app>
    <h-app-bar />
    <h-sider />

    <v-content>
      <v-flex xs12 sm12 md12 lg10 xl10 px-10>
        <v-layout align-center justify-center text-center row wrap>
          <transition name="slide-x-transition" mode="out-in">
            <router-view></router-view>
          </transition>
        </v-layout>
      </v-flex>
    </v-content>
  </v-app>
</template>

<script>
const axios = require("axios");
import HAppBar from "../components/account/HAppBar";
import HSider from "../components/account/HSider";

export default {
  name: "Account",
  components: {
    HAppBar,
    HSider
  },
  async created() {
    this.loading = true;
    if (!!this.$cookies.get("token")) {
      try {
        const res = await axios.get(this.$config.api + "/getaccount", {
          params: {
            token: this.$cookies.get("token")
          },
          withCredentials: true
        });
        if (res.data.ok) {
          const account = res.data.account;
          this.$store.state.account.firstName = account.firstName;
          this.$store.state.account.lastName = account.lastName;
          this.$store.state.account.email = account.email;
          this.$store.state.account.phone = account.phone;
          this.$store.state.account.birthday = account.birthday;
          this.$store.state.account.gender = account.gender;
        } else {
          this.$router.push("/introduction");
        }
      } catch (error) {
        console.log(error);
        this.$router.push("/introduction");
      }
    } else {
      this.$router.push("/introduction");
    }
    this.loading = false;
  }
};
</script>
