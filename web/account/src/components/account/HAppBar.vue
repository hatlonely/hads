<template>
  <v-app-bar app clipped-left elevate-on-scroll class="pb-3" color="#fafafa">
    <v-img :src="require('../../assets/img/logo.png')" max-width="70" tile></v-img>
    <v-toolbar-title>
      &nbsp;
      <span class="body-2">账号</span>
    </v-toolbar-title>
    <v-spacer></v-spacer>
    <v-text-field
      v-if="this.$store.state.account.isSignedIn"
      v-model="search"
      prepend-inner-icon="search"
      :label="'搜索 ' + $config.org + ' 账户'"
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
    <v-btn v-if="this.$store.state.account.email == ''" to="/signin" color="primary" depressed>登录</v-btn>
    <v-btn v-else to="/signin" color="#fafafa" depressed>
      <div v-html="identicon"></div>
    </v-btn>
  </v-app-bar>
</template>

<script>
const jdenticon = require("jdenticon");

export default {
  name: "HAppBar",
  computed: {
    identicon: function() {
      return jdenticon.toSvg(this.$store.state.account.email, 30);
    }
  },
  data: () => ({
    search: ""
  })
};
</script>
