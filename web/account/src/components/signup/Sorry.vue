<template>
  <v-card width="748" height="511" flat outlined>
    <v-layout row wrap mx-0 text-left px-8>
      <v-flex xs8 pr-8>
        <v-flex mt-8 mb-5 xs12 lg12>
          <v-layout>
            <v-img :src="require('../../assets/img/logo.png')" max-width="70" inline></v-img>
          </v-layout>
        </v-flex>
        <v-flex my-4 xs12 lg12>
          <h2>出了点问题</h2>
        </v-flex>

        <v-flex my-4>
          <v-card class="overflow-y-auto" height="300" flat>
            <div class="body-1">抱歉，出了点问题，请重试</div>
          </v-card>
        </v-flex>

        <v-flex mt-0 mx-0>
          <v-layout align-left justify-center row fill-height text-left>
            <v-flex xs3>
              <v-btn text color="primary" @click="$router.go(-1)" pl-0>取消</v-btn>
            </v-flex>
            <v-flex xs6></v-flex>
            <v-flex xs3>
              <v-btn color="primary" depressed @click="signup">我同意</v-btn>
            </v-flex>
          </v-layout>
        </v-flex>
      </v-flex>

      <v-flex xs4>
        <v-layout align-center justify-center row fill-height mx-0 class="text-center">
          <v-flex xs12>
            <v-flex xs12>
              <v-img :src="require('../../assets/img/signup_phone.svg')" max-width="244" inline></v-img>
            </v-flex>
            <v-flex xs12 px-4>
              <p class="body-1">您的个人信息不会对外公开，并且绝对安全</p>
            </v-flex>
          </v-flex>
        </v-layout>
      </v-flex>
    </v-layout>
  </v-card>
</template>

<script>
const axios = require("axios");

export default {
  methods: {
    async signup() {
      try {
        const response = await axios.post(this.$config.api + "/signup", {
          email: this.$store.state.signup.email,
          phone: this.$store.state.signup.phone,
          password: this.$store.state.signup.password,
          birthday: this.$store.state.signup.birthday,
          gender: this.$store.state.signup.gender,
          firstName: this.$store.state.signup.firstName,
          lastName: this.$store.state.signup.lastName
        });
        console.log(response.data);
        this.$router.push("/");
      } catch (error) {
        this.$router.push("/signup/sorry");
      }
    }
  }
};
</script>
